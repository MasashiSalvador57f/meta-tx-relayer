package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MasashiSalvador57f/meta-tx-relayer/api/constants"
	"github.com/MasashiSalvador57f/meta-tx-relayer/api/contracts"
	"github.com/MasashiSalvador57f/meta-tx-relayer/api/controllers/requests"
	"github.com/MasashiSalvador57f/meta-tx-relayer/api/infrastructure/ethereum/tx_option"

	"github.com/go-chi/chi"

	"github.com/MasashiSalvador57f/meta-tx-relayer/api/services"
)

type TxRelayController struct {
	optGenerator *txoption.ServerTxOptGenerator
}

func NewTxRelayController(optGenerator *txoption.ServerTxOptGenerator) *TxRelayController {
	return &TxRelayController{
		optGenerator: optGenerator,
	}
}

// GetTxRelayNonce returns local version number stored in TxRelayContract.
func (t *TxRelayController) GetTxRelayNonce(w http.ResponseWriter, r *http.Request) {
	addr := chi.URLParam(r, "address")

	txRelayContract, ok := r.Context().Value(constants.ContextKeyTxRelayContract).(*contract.TxRelay)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to unwrap tx relay contract instance in the context"))

		return
	}

	txOpt := t.optGenerator.GenerateNewOption(r.Context())
	txRelayer := services.NewTxRelay(txRelayContract, txOpt)

	nonce, err := txRelayer.GetNonce(addr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", nonce)))
}

// PostMetaTx is ...
func (t *TxRelayController) PostMetaTx(w http.ResponseWriter, r *http.Request) {
	var er requests.ECRecoveryRequest
	err := json.NewDecoder(r.Body).Decode(&er)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to decode json body %v", err)))
	}
	//
	//err = er.Validate()
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	w.Write([]byte(fmt.Sprintf("request format is wrong : %v", err)))
	//}

	txRelayContract, ok := r.Context().Value(constants.ContextKeyTxRelayContract).(*contract.TxRelay)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to unwrap tx relay contract instance in the context"))

		return
	}
	txOpt := t.optGenerator.GenerateNewOption(r.Context())
	txRelayer := services.NewTxRelay(txRelayContract, txOpt)

	erp := er.ToECRecoveryProperty()
	tx, err := txRelayer.RelayTx(
		erp.SigV, erp.SigR, erp.SigS,
		erp.Destination, erp.Data, erp.OriginalSigner,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to call RelayTx : %v", err)))

		return
	}

	err = json.NewEncoder(w).Encode(tx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

