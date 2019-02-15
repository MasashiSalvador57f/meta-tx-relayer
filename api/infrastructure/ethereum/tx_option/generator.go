package txoption

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// TODO cache estimated gas price (maybe?)

const (
	defaultTxValue  = 0
	defaultGasPrice = 20000000000
	defaultGasLimit = 51000
)

type ServerTxOptGenerator struct {
	privateKey     *ecdsa.PrivateKey
	publicKeyECDSA *ecdsa.PublicKey
	addr           common.Address
	// TODO: this might be share variable (considering sync package)
	nonce    *big.Int
	conn     *ethclient.Client
	gasPrice *big.Int
	gasLimit uint64
}

func NewServerTxOptGenerator(privateKeyStr string) (*ServerTxOptGenerator, error) {
	txOptGenerator := new(ServerTxOptGenerator)
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert string to ECDSA private key")
	}

	txOptGenerator.privateKey = privateKey
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.Errorf("failed to create publicKeyECDSA")
	}
	txOptGenerator.publicKeyECDSA = publicKeyECDSA
	txOptGenerator.gasPrice = big.NewInt(defaultGasPrice)
	txOptGenerator.gasLimit = defaultGasLimit

	return txOptGenerator, nil
}

func (s *ServerTxOptGenerator) GenerateNewOption(ctx context.Context) *bind.TransactOpts {
	opt := bind.NewKeyedTransactor(s.privateKey)
	opt.Context = ctx

	// TODO: handle latest pending nonce (with concurrent goroutine)
	opt.Nonce = s.nonce
	opt.GasLimit = s.gasLimit
	opt.GasPrice = s.gasPrice
	opt.Value = big.NewInt(defaultTxValue)

	return opt
}
