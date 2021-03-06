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

// SignValidatorABI is the input ABI used to generate the binding from.
const SignValidatorABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x70ae92d2\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"address\"}],\"name\":\"Log\",\"type\":\"event\",\"signature\":\"0xb8a00d6d8ca1be30bfec34d8f97e55f0f0fd9eeb7fb46e030516363d4cfe1ad6\"},{\"constant\":false,\"inputs\":[{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"originalSigner\",\"type\":\"address\"}],\"name\":\"validateSignature\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x3fa908c9\"}]"

// SignValidator is an auto generated Go binding around an Ethereum contract.
type SignValidator struct {
	SignValidatorCaller     // Read-only binding to the contract
	SignValidatorTransactor // Write-only binding to the contract
	SignValidatorFilterer   // Log filterer for contract events
}

// SignValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignValidatorSession struct {
	Contract     *SignValidator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignValidatorCallerSession struct {
	Contract *SignValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SignValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignValidatorTransactorSession struct {
	Contract     *SignValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SignValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignValidatorRaw struct {
	Contract *SignValidator // Generic contract binding to access the raw methods on
}

// SignValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignValidatorCallerRaw struct {
	Contract *SignValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// SignValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignValidatorTransactorRaw struct {
	Contract *SignValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignValidator creates a new instance of SignValidator, bound to a specific deployed contract.
func NewSignValidator(address common.Address, backend bind.ContractBackend) (*SignValidator, error) {
	contract, err := bindSignValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignValidator{SignValidatorCaller: SignValidatorCaller{contract: contract}, SignValidatorTransactor: SignValidatorTransactor{contract: contract}, SignValidatorFilterer: SignValidatorFilterer{contract: contract}}, nil
}

// NewSignValidatorCaller creates a new read-only instance of SignValidator, bound to a specific deployed contract.
func NewSignValidatorCaller(address common.Address, caller bind.ContractCaller) (*SignValidatorCaller, error) {
	contract, err := bindSignValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignValidatorCaller{contract: contract}, nil
}

// NewSignValidatorTransactor creates a new write-only instance of SignValidator, bound to a specific deployed contract.
func NewSignValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*SignValidatorTransactor, error) {
	contract, err := bindSignValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignValidatorTransactor{contract: contract}, nil
}

// NewSignValidatorFilterer creates a new log filterer instance of SignValidator, bound to a specific deployed contract.
func NewSignValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*SignValidatorFilterer, error) {
	contract, err := bindSignValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignValidatorFilterer{contract: contract}, nil
}

// bindSignValidator binds a generic wrapper to an already deployed contract.
func bindSignValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignValidator *SignValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SignValidator.Contract.SignValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignValidator *SignValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignValidator.Contract.SignValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignValidator *SignValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignValidator.Contract.SignValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignValidator *SignValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SignValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignValidator *SignValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignValidator *SignValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignValidator.Contract.contract.Transact(opts, method, params...)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) constant returns(uint256)
func (_SignValidator *SignValidatorCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SignValidator.contract.Call(opts, out, "nonce", arg0)
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) constant returns(uint256)
func (_SignValidator *SignValidatorSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _SignValidator.Contract.Nonce(&_SignValidator.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) constant returns(uint256)
func (_SignValidator *SignValidatorCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _SignValidator.Contract.Nonce(&_SignValidator.CallOpts, arg0)
}

// ValidateSignature is a paid mutator transaction binding the contract method 0x3fa908c9.
//
// Solidity: function validateSignature(uint8 sigV, bytes32 sigR, bytes32 sigS, bytes data, address originalSigner) returns()
func (_SignValidator *SignValidatorTransactor) ValidateSignature(opts *bind.TransactOpts, sigV uint8, sigR [32]byte, sigS [32]byte, data []byte, originalSigner common.Address) (*types.Transaction, error) {
	return _SignValidator.contract.Transact(opts, "validateSignature", sigV, sigR, sigS, data, originalSigner)
}

// ValidateSignature is a paid mutator transaction binding the contract method 0x3fa908c9.
//
// Solidity: function validateSignature(uint8 sigV, bytes32 sigR, bytes32 sigS, bytes data, address originalSigner) returns()
func (_SignValidator *SignValidatorSession) ValidateSignature(sigV uint8, sigR [32]byte, sigS [32]byte, data []byte, originalSigner common.Address) (*types.Transaction, error) {
	return _SignValidator.Contract.ValidateSignature(&_SignValidator.TransactOpts, sigV, sigR, sigS, data, originalSigner)
}

// ValidateSignature is a paid mutator transaction binding the contract method 0x3fa908c9.
//
// Solidity: function validateSignature(uint8 sigV, bytes32 sigR, bytes32 sigS, bytes data, address originalSigner) returns()
func (_SignValidator *SignValidatorTransactorSession) ValidateSignature(sigV uint8, sigR [32]byte, sigS [32]byte, data []byte, originalSigner common.Address) (*types.Transaction, error) {
	return _SignValidator.Contract.ValidateSignature(&_SignValidator.TransactOpts, sigV, sigR, sigS, data, originalSigner)
}

// SignValidatorLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the SignValidator contract.
type SignValidatorLogIterator struct {
	Event *SignValidatorLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SignValidatorLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignValidatorLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SignValidatorLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SignValidatorLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignValidatorLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignValidatorLog represents a Log event raised by the SignValidator contract.
type SignValidatorLog struct {
	common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0xb8a00d6d8ca1be30bfec34d8f97e55f0f0fd9eeb7fb46e030516363d4cfe1ad6.
//
// Solidity: event Log(address )
func (_SignValidator *SignValidatorFilterer) FilterLog(opts *bind.FilterOpts) (*SignValidatorLogIterator, error) {

	logs, sub, err := _SignValidator.contract.FilterLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return &SignValidatorLogIterator{contract: _SignValidator.contract, event: "Log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0xb8a00d6d8ca1be30bfec34d8f97e55f0f0fd9eeb7fb46e030516363d4cfe1ad6.
//
// Solidity: event Log(address )
func (_SignValidator *SignValidatorFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *SignValidatorLog) (event.Subscription, error) {

	logs, sub, err := _SignValidator.contract.WatchLogs(opts, "Log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignValidatorLog)
				if err := _SignValidator.contract.UnpackLog(event, "Log", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
