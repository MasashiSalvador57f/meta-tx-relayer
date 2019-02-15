package requests

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type ECRecoveryRequest struct {
	SigV uint8 `json:"sig_v"`
	SigR string `json:"sig_r"`
	SigS string `json:"sig_s"`
	DestAddress string `json:"dest_address"`
	RawTx string `json:"raw_tx"`
	OriginalSigner string `json:"original_sender"`
}

// Validate checks request format.
func (er *ECRecoveryRequest) Validate() error {
	// TODO: brush up the validiation e.g. byte string should begin with 0x or so...
	if len(er.SigR) != 32 {
		return errors.New("sig_r's length should be 32")
	}
	if len(er.SigS) != 32 {
		return errors.New("sig_s's length should be 32")
	}
	if len(er.DestAddress) < 0 {
		return errors.New("dest_address should be non empty string")
	}
	if len(er.RawTx) < 0 {
		return errors.New("raw_tx should be non empty")
	}
	if len(er.OriginalSigner) < 0 {
		return errors.New("original_signer should be non empty string")
	}

	return nil
}

func (er *ECRecoveryRequest) ToECRecoveryProperty() *ECRecoveryProperty {
	var sigR [32]byte
	var sigS [32]byte
	copy(sigR[:], er.SigR)
	copy(sigS[:], er.SigS)

	destination := common.HexToAddress(er.DestAddress)
	signer := common.HexToAddress(er.OriginalSigner)

	prop := new(ECRecoveryProperty)
	prop.SigV = er.SigV
	prop.SigR = sigR
	prop.SigS = sigS
	prop.Destination = destination
	// NOTE this cast is expensive coz it require memory copy TODO: fixit)
	prop.Data = []byte(er.RawTx)
	prop.OriginalSigner = signer

	return prop
}

type ECRecoveryProperty struct {
	SigV uint8
	SigR [32]byte
	SigS [32]byte
	Destination common.Address
	Data []byte
	OriginalSigner common.Address
}