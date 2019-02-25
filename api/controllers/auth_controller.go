package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/MasashiSalvador57f/meta-tx-relayer/api/controllers/repsonses"
	"github.com/MasashiSalvador57f/meta-tx-relayer/api/services"
	"net/http"
	"time"
)

type AuthController struct {}

func NewAuthController() *AuthController {
	return new(AuthController)
}

func (ac *AuthController) IssueRawMessage(w http.ResponseWriter, r *http.Request) {
	msgIssuer := services.NewOneTimeMessageHandler()
	msg := msgIssuer.IssueMsg(30 * 24 * time.Hour, "seed")

	res := repsonses.RawAuthMessage{
		MessageToSign: msg,
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprint(err.Error())))

		return
	}
}

func PostSignedMessage(w http.ResponseWriter, r *http.Request) {

}