package repsonses

// RawAuthMessage is ...
type RawAuthMessage struct {
	MessageToSign string `json:"message_to_sign"`
	AuthNonce     int64  `json:"auth_nonce"`
}
