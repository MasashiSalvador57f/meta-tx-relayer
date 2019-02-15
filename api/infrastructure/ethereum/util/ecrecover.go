package util

import "github.com/ethereum/go-ethereum/common"

type ECRecoverProperty struct {
	sigV uint8
	sigR [32]byte
	sigS [32]byte
	destination common.Address
	data []byte
	originalSigner common.Address
}

func NewECRecoveryProperty(
	sigv uint8,
	sigr string,
	sigs string,
	destAddrStr string,
	rawData string,
	signerAddr string,
) *ECRecoverProperty {
	var sigR [32]byte
	var sigS [32]byte
	copy(sigR[:], sigr)
	copy(sigS[:], sigs)

	destination := common.HexToAddress(destAddrStr)
	signer := common.HexToAddress(signerAddr)

	prop := new(ECRecoverProperty)
	prop.sigV = sigv
	prop.sigR = sigR
	prop.sigS = sigS
	prop.destination = destination
	// NOTE this cast is expensive coz it require memory copy TODO: fixit)
	prop.data = []byte(rawData)
	prop.originalSigner = signer

	return prop
}