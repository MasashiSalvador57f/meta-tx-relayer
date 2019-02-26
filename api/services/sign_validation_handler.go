package services

import (
	"github.com/MasashiSalvador57f/meta-tx-relayer/api/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type SignValidationHandler struct {
	signValidatorContract *contract.SignValidator
	txOpt *types.Transaction
}

// SignValidationHandler is ...
func NewSignValidationHanlder() SignValidator {
	return new(SignValidationHandler)
}

func (s *SignValidationHandler) GetNonce(addrStr string) (*big.Int, error) {
	addr := common.HexToAddress(addrStr)
	nonce, err := s.signValidatorContract.Nonce(nil, addr)

	return nonce, err
}

func (s *SignValidationHandler) ValidateSignature(sigV uint8, sigR [32]byte, sigS [32]byte, data []byte, originalSigner common.Address) (*types.Transaction, error) {
	tx, err := s.signValidatorContract.ValidateSignature(s.txOpt, sigV, sigR, sigS, data, originalSigner)
	return tx, err
}