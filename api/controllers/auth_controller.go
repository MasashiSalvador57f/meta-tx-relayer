package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MasashiSalvador57f/meta-tx-relayer/api/constants"
	contract "github.com/MasashiSalvador57f/meta-tx-relayer/api/contracts"
	"github.com/go-chi/chi"

	"github.com/MasashiSalvador57f/meta-tx-relayer/api/controllers/repsonses"
	"github.com/MasashiSalvador57f/meta-tx-relayer/api/controllers/requests"
	"github.com/MasashiSalvador57f/meta-tx-relayer/api/services"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return new(AuthController)
}

func (ac *AuthController) IssueRawMessage(w http.ResponseWriter, r *http.Request) {
	addr := chi.URLParam(r, "address")
	msgIssuer := services.NewOneTimeMessageHandler()
	msg := msgIssuer.IssueMsg(30*24*time.Hour, "seed")

	signValidatorContract := r.Context().Value(constants.ContextKeyAuthContract).(*contract.SignValidator)
	signValidationHandler := services.NewSignValidationHandler(signValidatorContract)
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

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprint(err.Error())))

		return
	}
}

func PostSignedMessage(w http.ResponseWriter, r *http.Request) {
	var er requests.ECRecoveryRequest
	err := json.NewDecoder(r.Body).Decode(&er)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to decode json body %v", err)))
	}

	erp := er.ToECRecoveryProperty()

}
