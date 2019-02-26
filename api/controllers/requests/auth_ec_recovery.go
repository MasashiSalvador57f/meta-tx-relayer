package requests

import "github.com/pkg/errors"

type AuthECRecoveryRequest struct {
	ECRecoveryRequest
}

// Validation returns error on fail.
func (er *AuthECRecoveryRequest) Validate() error {
	// TODO: brush up the validiation e.g. byte string should begin with 0x or so...
	if len(er.SigR) != 32 {
		return errors.New("sig_r's length should be 32")
	}
	if len(er.SigS) != 32 {
		return errors.New("sig_s's length should be 32")
	}
	if len(er.Data) < 0 {
		return errors.New("raw_tx should be non empty")
	}
	if len(er.OriginalSigner) < 0 {
		return errors.New("original_signer should be non empty string")
	}

	return nil
}
