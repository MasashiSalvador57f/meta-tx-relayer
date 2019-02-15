package services

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"

	"github.com/MasashiSalvador57f/meta-tx-relayer/api/contracts"
	"github.com/ethereum/go-ethereum/common"
)

// TxRelay implements TxRelayHandler
type TxRelay struct {
	txrContract *contract.TxRelay
	txOpt *bind.TransactOpts
}

// NewTxRelay returns TxRelay instance.
func NewTxRelay(txrContract *contract.TxRelay, txOpt *bind.TransactOpts) TxRelayHandler {
	return &TxRelay{
		txrContract: txrContract,
		txOpt: txOpt,
	}
}

// GetNonce returns localNonce of ...
func (t *TxRelay) GetNonce(addr string) (*big.Int, error) {
	nonce, err := t.txrContract.Nonce(nil, common.HexToAddress(addr))

	return nonce, err
}

// RelayTx relays signed transaction to RelayerContract.
func (t *TxRelay) RelayTx(
	sigV uint8, sigR [32]byte, sigS [32]byte, destination common.Address, data []byte, originalSigner common.Address,
) (*types.Transaction, error) {
	tx, err := t.txrContract.RelayMetaTx(t.txOpt, sigV, sigR, sigS, destination, data, originalSigner)

	return tx, err
}
