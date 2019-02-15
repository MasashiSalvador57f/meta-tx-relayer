package services

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// TxRelayHandler is ...
type TxRelayHandler interface {
	GetNonce(string) (*big.Int, error)
	RelayTx(sigV uint8, sigR [32]byte, sigS [32]byte, destination common.Address, data []byte, originalSigner common.Address) (*types.Transaction, error)
}
