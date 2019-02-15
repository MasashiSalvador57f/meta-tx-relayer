package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	txoption "github.com/MasashiSalvador57f/meta-tx-relayer/api/infrastructure/ethereum/tx_option"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"

	"github.com/MasashiSalvador57f/meta-tx-relayer/api/constants"
	contract "github.com/MasashiSalvador57f/meta-tx-relayer/api/contracts"
	"github.com/MasashiSalvador57f/meta-tx-relayer/api/controllers"
)

func main() {
	rawTx := ""
	fmt.Println(rawTx)
	txRelayAddressStr := os.Getenv("TX_RELAY_ADDRESS")
	authContractAddressStr := os.Getenv("AUTH_CONTRACT_ADDRESS")
	ethreumNodeEndpoint := os.Getenv("ETH_NODE_ENDPOINT")
	serverPrivateKeyStr := os.Getenv("SERVER_PRIVATE_KEY")

	port := os.Getenv("PORT")
	if len(port) <= 0 {
		port = "3000"
	}
	if len(txRelayAddressStr) <= 0 {
		log.Fatal("txRelayAddress is required in env variable")
	}

	conn, err := ethclient.Dial(ethreumNodeEndpoint)
	if err != nil {
		log.Fatalf("failed to connect ethereum node %s", err)
	}

	txOptGenerator, err := txoption.NewServerTxOptGenerator(serverPrivateKeyStr)
	if err != nil {
		log.Fatalf("failed to create tx opt generator %v", err)
	}

	txRelayAddress := common.HexToAddress(txRelayAddressStr[1 : len(txRelayAddressStr)-1])
	txRelay, err := contract.NewTxRelay(txRelayAddress, conn)

	if err != nil {
		log.Fatalf("error in creating txRelayContract instance address: %v, conn : %v", txRelayAddress, conn)
	}

	authContractAddress := common.HexToAddress(authContractAddressStr[1 : len(authContractAddressStr)-1])
	authContract, err := contract.NewSignValidator(authContractAddress, conn)
	if err != nil {
		log.Fatalf("error in creating signValidator contract")
	}

	txRelayController := controllers.NewTxRelayController(txOptGenerator)
	r := chi.NewRouter()
	r.Route("/tx_relay/nonce", func(r chi.Router) {
		r.Use(txRelayCtx(txRelay))
		r.Get("/", txRelayController.GetTxRelayNonce)
	})

	r.Route("/tx_relay/meta_tx", func(r chi.Router) {
		r.Use(txRelayCtx(txRelay))
		r.Post("/", txRelayController.PostMetaTx)
	})

	authController := controllers.NewAuthController(txOptGenerator)
	r.Route("/auth", func(r chi.Router) {
		r.Use(authCtx(authContract))
		r.Get("/raw_message", authController.IssueRawMessage)
		r.Post("/token", authController.PostSignedMessage)
	})

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("error starting gateway-api : %s", err.Error())
	}
}

func txRelayCtx(txRelay *contract.TxRelay) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), constants.ContextKeyTxRelayContract, txRelay)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func authCtx(authContract *contract.SignValidator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), constants.ContextKeyAuthContract, authContract)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
