package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	txoption "github.com/MasashiSalvador57f/meta-tx-relayer/api/infrastructure/ethereum/tx_option"

	"github.com/MasashiSalvador57f/meta-tx-relayer/api/constants"
	contract "github.com/MasashiSalvador57f/meta-tx-relayer/api/contracts"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"

	"github.com/MasashiSalvador57f/meta-tx-relayer/api/controllers/repsonses"
	"github.com/MasashiSalvador57f/meta-tx-relayer/api/controllers/requests"
	"github.com/MasashiSalvador57f/meta-tx-relayer/api/services"
)

type AuthController struct {
	optGenerator *txoption.ServerTxOptGenerator
}

func NewAuthController(optGenerator *txoption.ServerTxOptGenerator) *AuthController {
	return &AuthController{
		optGenerator,
	}
}

func (ac *AuthController) IssueRawMessage(w http.ResponseWriter, r *http.Request) {
	addr := chi.URLParam(r, "address")
	msgIssuer := services.NewOneTimeMessageHandler()
	msg := msgIssuer.IssueMsg(30*24*time.Hour, "seed")

	signValidatorContract := r.Context().Value(constants.ContextKeyAuthContract).(*contract.SignValidator)
	signValidationHandler := services.NewSignValidationHandler(signValidatorContract, nil)
	nonce, err := signValidationHandler.GetNonce(addr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	res := repsonses.RawAuthMessage{
		MessageToSign: msg,
		AuthNonce:     nonce.Int64(),
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprint(err.Error())))

		return
	}
}

func (ac *AuthController) PostSignedMessage(w http.ResponseWriter, r *http.Request) {
	var er requests.ECRecoveryRequest
	err := json.NewDecoder(r.Body).Decode(&er)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to decode json body %v", err)))
	}

	// NOTE validation is skipped for a moment.
	// TODO : comment in
	//err = er.Validate()
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	w.Write([]byte(fmt.Sprint(err.Error())))
	//}

	erp := er.ToECRecoveryProperty()

	txOpt := ac.optGenerator.GenerateNewOption(r.Context())
	signValidatorContract := r.Context().Value(constants.ContextKeyAuthContract).(*contract.SignValidator)
	signValidationHandler := services.NewSignValidationHandler(signValidatorContract, txOpt)

	tx, err := signValidationHandler.ValidateSignature(
		erp.SigV,
		erp.SigR,
		erp.SigS,
		erp.Data,
		erp.OriginalSigner,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(err.Error())))

		return
	}

	spew.Dump(tx)
	w.WriteHeader(http.StatusCreated)
}
