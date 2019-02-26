package services

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type SignValidator interface {
	GetNonce(addr string) *big.Int
	ValidateSignature(sigV uint8, sigR [32]byte, sigS [32]byte, data []byte, originalSigner common.Address) (*types.Transaction, error)
}
