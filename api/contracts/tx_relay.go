// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TxRelayABI is the input ABI used to generate the binding from.
const TxRelayABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x70ae92d2\"},{\"constant\":false,\"inputs\":[{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"originalSigner\",\"type\":\"address\"}],\"name\":\"relayMetaTx\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xc3f44c0a\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xe43252d7\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x8ab1d681\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x678f3ef3\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x5cadd08d\"}]"

// TxRelay is an auto generated Go binding around an Ethereum contract.
type TxRelay struct {
	TxRelayCaller     // Read-only binding to the contract
	TxRelayTransactor // Write-only binding to the contract
	TxRelayFilterer   // Log filterer for contract events
}

// TxRelayCaller is an auto generated read-only Go binding around an Ethereum contract.
type TxRelayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxRelayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TxRelayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxRelayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TxRelayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxRelaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TxRelaySession struct {
	Contract     *TxRelay          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TxRelayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TxRelayCallerSession struct {
	Contract *TxRelayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TxRelayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TxRelayTransactorSession struct {
	Contract     *TxRelayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TxRelayRaw is an auto generated low-level Go binding around an Ethereum contract.
type TxRelayRaw struct {
	Contract *TxRelay // Generic contract binding to access the raw methods on
}

// TxRelayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TxRelayCallerRaw struct {
	Contract *TxRelayCaller // Generic read-only contract binding to access the raw methods on
}

// TxRelayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TxRelayTransactorRaw struct {
	Contract *TxRelayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTxRelay creates a new instance of TxRelay, bound to a specific deployed contract.
func NewTxRelay(address common.Address, backend bind.ContractBackend) (*TxRelay, error) {
	contract, err := bindTxRelay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TxRelay{TxRelayCaller: TxRelayCaller{contract: contract}, TxRelayTransactor: TxRelayTransactor{contract: contract}, TxRelayFilterer: TxRelayFilterer{contract: contract}}, nil
}

// NewTxRelayCaller creates a new read-only instance of TxRelay, bound to a specific deployed contract.
func NewTxRelayCaller(address common.Address, caller bind.ContractCaller) (*TxRelayCaller, error) {
	contract, err := bindTxRelay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TxRelayCaller{contract: contract}, nil
}

// NewTxRelayTransactor creates a new write-only instance of TxRelay, bound to a specific deployed contract.
func NewTxRelayTransactor(address common.Address, transactor bind.ContractTransactor) (*TxRelayTransactor, error) {
	contract, err := bindTxRelay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TxRelayTransactor{contract: contract}, nil
}

// NewTxRelayFilterer creates a new log filterer instance of TxRelay, bound to a specific deployed contract.
func NewTxRelayFilterer(address common.Address, filterer bind.ContractFilterer) (*TxRelayFilterer, error) {
	contract, err := bindTxRelay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TxRelayFilterer{contract: contract}, nil
}

// bindTxRelay binds a generic wrapper to an already deployed contract.
func bindTxRelay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TxRelayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxRelay *TxRelayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TxRelay.Contract.TxRelayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxRelay *TxRelayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxRelay.Contract.TxRelayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxRelay *TxRelayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxRelay.Contract.TxRelayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxRelay *TxRelayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TxRelay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxRelay *TxRelayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxRelay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxRelay *TxRelayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxRelay.Contract.contract.Transact(opts, method, params...)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) constant returns(uint256)
func (_TxRelay *TxRelayCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TxRelay.contract.Call(opts, out, "nonce", arg0)
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) constant returns(uint256)
func (_TxRelay *TxRelaySession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _TxRelay.Contract.Nonce(&_TxRelay.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) constant returns(uint256)
func (_TxRelay *TxRelayCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _TxRelay.Contract.Nonce(&_TxRelay.CallOpts, arg0)
}

// AddDestination is a paid mutator transaction binding the contract method 0x678f3ef3.
//
// Solidity: function addDestination(address addr) returns()
func (_TxRelay *TxRelayTransactor) AddDestination(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TxRelay.contract.Transact(opts, "addDestination", addr)
}

// AddDestination is a paid mutator transaction binding the contract method 0x678f3ef3.
//
// Solidity: function addDestination(address addr) returns()
func (_TxRelay *TxRelaySession) AddDestination(addr common.Address) (*types.Transaction, error) {
	return _TxRelay.Contract.AddDestination(&_TxRelay.TransactOpts, addr)
}

// AddDestination is a paid mutator transaction binding the contract method 0x678f3ef3.
//
// Solidity: function addDestination(address addr) returns()
func (_TxRelay *TxRelayTransactorSession) AddDestination(addr common.Address) (*types.Transaction, error) {
	return _TxRelay.Contract.AddDestination(&_TxRelay.TransactOpts, addr)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address addr) returns()
func (_TxRelay *TxRelayTransactor) AddToWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TxRelay.contract.Transact(opts, "addToWhitelist", addr)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address addr) returns()
func (_TxRelay *TxRelaySession) AddToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _TxRelay.Contract.AddToWhitelist(&_TxRelay.TransactOpts, addr)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address addr) returns()
func (_TxRelay *TxRelayTransactorSession) AddToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _TxRelay.Contract.AddToWhitelist(&_TxRelay.TransactOpts, addr)
}

// RelayMetaTx is a paid mutator transaction binding the contract method 0xc3f44c0a.
//
// Solidity: function relayMetaTx(uint8 sigV, bytes32 sigR, bytes32 sigS, address destination, bytes data, address originalSigner) returns()
func (_TxRelay *TxRelayTransactor) RelayMetaTx(opts *bind.TransactOpts, sigV uint8, sigR [32]byte, sigS [32]byte, destination common.Address, data []byte, originalSigner common.Address) (*types.Transaction, error) {
	return _TxRelay.contract.Transact(opts, "relayMetaTx", sigV, sigR, sigS, destination, data, originalSigner)
}

// RelayMetaTx is a paid mutator transaction binding the contract method 0xc3f44c0a.
//
// Solidity: function relayMetaTx(uint8 sigV, bytes32 sigR, bytes32 sigS, address destination, bytes data, address originalSigner) returns()
func (_TxRelay *TxRelaySession) RelayMetaTx(sigV uint8, sigR [32]byte, sigS [32]byte, destination common.Address, data []byte, originalSigner common.Address) (*types.Transaction, error) {
	return _TxRelay.Contract.RelayMetaTx(&_TxRelay.TransactOpts, sigV, sigR, sigS, destination, data, originalSigner)
}

// RelayMetaTx is a paid mutator transaction binding the contract method 0xc3f44c0a.
//
// Solidity: function relayMetaTx(uint8 sigV, bytes32 sigR, bytes32 sigS, address destination, bytes data, address originalSigner) returns()
func (_TxRelay *TxRelayTransactorSession) RelayMetaTx(sigV uint8, sigR [32]byte, sigS [32]byte, destination common.Address, data []byte, originalSigner common.Address) (*types.Transaction, error) {
	return _TxRelay.Contract.RelayMetaTx(&_TxRelay.TransactOpts, sigV, sigR, sigS, destination, data, originalSigner)
}

// RemoveDestination is a paid mutator transaction binding the contract method 0x5cadd08d.
//
// Solidity: function removeDestination(address addr) returns()
func (_TxRelay *TxRelayTransactor) RemoveDestination(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TxRelay.contract.Transact(opts, "removeDestination", addr)
}

// RemoveDestination is a paid mutator transaction binding the contract method 0x5cadd08d.
//
// Solidity: function removeDestination(address addr) returns()
func (_TxRelay *TxRelaySession) RemoveDestination(addr common.Address) (*types.Transaction, error) {
	return _TxRelay.Contract.RemoveDestination(&_TxRelay.TransactOpts, addr)
}

// RemoveDestination is a paid mutator transaction binding the contract method 0x5cadd08d.
//
// Solidity: function removeDestination(address addr) returns()
func (_TxRelay *TxRelayTransactorSession) RemoveDestination(addr common.Address) (*types.Transaction, error) {
	return _TxRelay.Contract.RemoveDestination(&_TxRelay.TransactOpts, addr)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address addr) returns()
func (_TxRelay *TxRelayTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TxRelay.contract.Transact(opts, "removeFromWhitelist", addr)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address addr) returns()
func (_TxRelay *TxRelaySession) RemoveFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _TxRelay.Contract.RemoveFromWhitelist(&_TxRelay.TransactOpts, addr)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address addr) returns()
func (_TxRelay *TxRelayTransactorSession) RemoveFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _TxRelay.Contract.RemoveFromWhitelist(&_TxRelay.TransactOpts, addr)
}
