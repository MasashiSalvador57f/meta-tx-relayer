package services

import (
	"math/big"

	contract "github.com/MasashiSalvador57f/meta-tx-relayer/api/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type SignValidationHandler struct {
	signValidatorContract *contract.SignValidator
	txOpt                 *bind.TransactOpts
}

// SignValidationHandler is ...
func NewSignValidationHandler(c *contract.SignValidator) SignValidator {
	return &SignValidationHandler{
		signValidatorContract: c,
	}
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
