// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethSusyContract

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

// EthSusyContractABI is the input ABI used to generate the binding from.
const EthSusyContractABI = "[{\"inputs\":[{\"internalType\":\"address[5]\",\"name\":\"newAdmins\",\"type\":\"address[5]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"target\",\"type\":\"string\"}],\"name\":\"NewRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Return\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumSupersymmetry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"StatusChanged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"enumSupersymmetry.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumSupersymmetry.Type\",\"name\":\"rType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"target\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createBurnRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"sender\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"createMintRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[5]\",\"name\":\"v\",\"type\":\"uint8[5]\"},{\"internalType\":\"bytes32[5]\",\"name\":\"r\",\"type\":\"bytes32[5]\"},{\"internalType\":\"bytes32[5]\",\"name\":\"s\",\"type\":\"bytes32[5]\"},{\"internalType\":\"uint8\",\"name\":\"intStatus\",\"type\":\"uint8\"}],\"name\":\"changeStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthSusyContract is an auto generated Go binding around an Ethereum contract.
type EthSusyContract struct {
	EthSusyContractCaller     // Read-only binding to the contract
	EthSusyContractTransactor // Write-only binding to the contract
	EthSusyContractFilterer   // Log filterer for contract events
}

// EthSusyContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthSusyContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSusyContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthSusyContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSusyContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthSusyContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSusyContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthSusyContractSession struct {
	Contract     *EthSusyContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthSusyContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthSusyContractCallerSession struct {
	Contract *EthSusyContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EthSusyContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthSusyContractTransactorSession struct {
	Contract     *EthSusyContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EthSusyContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthSusyContractRaw struct {
	Contract *EthSusyContract // Generic contract binding to access the raw methods on
}

// EthSusyContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthSusyContractCallerRaw struct {
	Contract *EthSusyContractCaller // Generic read-only contract binding to access the raw methods on
}

// EthSusyContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthSusyContractTransactorRaw struct {
	Contract *EthSusyContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthSusyContract creates a new instance of EthSusyContract, bound to a specific deployed contract.
func NewEthSusyContract(address common.Address, backend bind.ContractBackend) (*EthSusyContract, error) {
	contract, err := bindEthSusyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthSusyContract{EthSusyContractCaller: EthSusyContractCaller{contract: contract}, EthSusyContractTransactor: EthSusyContractTransactor{contract: contract}, EthSusyContractFilterer: EthSusyContractFilterer{contract: contract}}, nil
}

// NewEthSusyContractCaller creates a new read-only instance of EthSusyContract, bound to a specific deployed contract.
func NewEthSusyContractCaller(address common.Address, caller bind.ContractCaller) (*EthSusyContractCaller, error) {
	contract, err := bindEthSusyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthSusyContractCaller{contract: contract}, nil
}

// NewEthSusyContractTransactor creates a new write-only instance of EthSusyContract, bound to a specific deployed contract.
func NewEthSusyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*EthSusyContractTransactor, error) {
	contract, err := bindEthSusyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthSusyContractTransactor{contract: contract}, nil
}

// NewEthSusyContractFilterer creates a new log filterer instance of EthSusyContract, bound to a specific deployed contract.
func NewEthSusyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*EthSusyContractFilterer, error) {
	contract, err := bindEthSusyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthSusyContractFilterer{contract: contract}, nil
}

// bindEthSusyContract binds a generic wrapper to an already deployed contract.
func bindEthSusyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthSusyContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthSusyContract *EthSusyContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthSusyContract.Contract.EthSusyContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthSusyContract *EthSusyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSusyContract.Contract.EthSusyContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthSusyContract *EthSusyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthSusyContract.Contract.EthSusyContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthSusyContract *EthSusyContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthSusyContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthSusyContract *EthSusyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthSusyContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthSusyContract *EthSusyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthSusyContract.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x14bfd6d0.
//
// Solidity: function admins(uint256 ) constant returns(address)
func (_EthSusyContract *EthSusyContractCaller) Admins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthSusyContract.contract.Call(opts, out, "admins", arg0)
	return *ret0, err
}

// Admins is a free data retrieval call binding the contract method 0x14bfd6d0.
//
// Solidity: function admins(uint256 ) constant returns(address)
func (_EthSusyContract *EthSusyContractSession) Admins(arg0 *big.Int) (common.Address, error) {
	return _EthSusyContract.Contract.Admins(&_EthSusyContract.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x14bfd6d0.
//
// Solidity: function admins(uint256 ) constant returns(address)
func (_EthSusyContract *EthSusyContractCallerSession) Admins(arg0 *big.Int) (common.Address, error) {
	return _EthSusyContract.Contract.Admins(&_EthSusyContract.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) constant returns(uint8 status, uint8 rType, address owner, string target, uint256 tokenAmount)
func (_EthSusyContract *EthSusyContractCaller) Requests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status      uint8
	RType       uint8
	Owner       common.Address
	Target      string
	TokenAmount *big.Int
}, error) {
	ret := new(struct {
		Status      uint8
		RType       uint8
		Owner       common.Address
		Target      string
		TokenAmount *big.Int
	})
	out := ret
	err := _EthSusyContract.contract.Call(opts, out, "requests", arg0)
	return *ret, err
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) constant returns(uint8 status, uint8 rType, address owner, string target, uint256 tokenAmount)
func (_EthSusyContract *EthSusyContractSession) Requests(arg0 [32]byte) (struct {
	Status      uint8
	RType       uint8
	Owner       common.Address
	Target      string
	TokenAmount *big.Int
}, error) {
	return _EthSusyContract.Contract.Requests(&_EthSusyContract.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) constant returns(uint8 status, uint8 rType, address owner, string target, uint256 tokenAmount)
func (_EthSusyContract *EthSusyContractCallerSession) Requests(arg0 [32]byte) (struct {
	Status      uint8
	RType       uint8
	Owner       common.Address
	Target      string
	TokenAmount *big.Int
}, error) {
	return _EthSusyContract.Contract.Requests(&_EthSusyContract.CallOpts, arg0)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_EthSusyContract *EthSusyContractCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthSusyContract.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_EthSusyContract *EthSusyContractSession) TokenAddress() (common.Address, error) {
	return _EthSusyContract.Contract.TokenAddress(&_EthSusyContract.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_EthSusyContract *EthSusyContractCallerSession) TokenAddress() (common.Address, error) {
	return _EthSusyContract.Contract.TokenAddress(&_EthSusyContract.CallOpts)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xa8413510.
//
// Solidity: function changeStatus(bytes32 requestHash, uint8[5] v, bytes32[5] r, bytes32[5] s, uint8 intStatus) returns()
func (_EthSusyContract *EthSusyContractTransactor) ChangeStatus(opts *bind.TransactOpts, requestHash [32]byte, v [5]uint8, r [5][32]byte, s [5][32]byte, intStatus uint8) (*types.Transaction, error) {
	return _EthSusyContract.contract.Transact(opts, "changeStatus", requestHash, v, r, s, intStatus)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xa8413510.
//
// Solidity: function changeStatus(bytes32 requestHash, uint8[5] v, bytes32[5] r, bytes32[5] s, uint8 intStatus) returns()
func (_EthSusyContract *EthSusyContractSession) ChangeStatus(requestHash [32]byte, v [5]uint8, r [5][32]byte, s [5][32]byte, intStatus uint8) (*types.Transaction, error) {
	return _EthSusyContract.Contract.ChangeStatus(&_EthSusyContract.TransactOpts, requestHash, v, r, s, intStatus)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xa8413510.
//
// Solidity: function changeStatus(bytes32 requestHash, uint8[5] v, bytes32[5] r, bytes32[5] s, uint8 intStatus) returns()
func (_EthSusyContract *EthSusyContractTransactorSession) ChangeStatus(requestHash [32]byte, v [5]uint8, r [5][32]byte, s [5][32]byte, intStatus uint8) (*types.Transaction, error) {
	return _EthSusyContract.Contract.ChangeStatus(&_EthSusyContract.TransactOpts, requestHash, v, r, s, intStatus)
}

// CreateBurnRequest is a paid mutator transaction binding the contract method 0x983d6d9d.
//
// Solidity: function createBurnRequest(string recipient, uint256 amount) returns(bytes32)
func (_EthSusyContract *EthSusyContractTransactor) CreateBurnRequest(opts *bind.TransactOpts, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _EthSusyContract.contract.Transact(opts, "createBurnRequest", recipient, amount)
}

// CreateBurnRequest is a paid mutator transaction binding the contract method 0x983d6d9d.
//
// Solidity: function createBurnRequest(string recipient, uint256 amount) returns(bytes32)
func (_EthSusyContract *EthSusyContractSession) CreateBurnRequest(recipient string, amount *big.Int) (*types.Transaction, error) {
	return _EthSusyContract.Contract.CreateBurnRequest(&_EthSusyContract.TransactOpts, recipient, amount)
}

// CreateBurnRequest is a paid mutator transaction binding the contract method 0x983d6d9d.
//
// Solidity: function createBurnRequest(string recipient, uint256 amount) returns(bytes32)
func (_EthSusyContract *EthSusyContractTransactorSession) CreateBurnRequest(recipient string, amount *big.Int) (*types.Transaction, error) {
	return _EthSusyContract.Contract.CreateBurnRequest(&_EthSusyContract.TransactOpts, recipient, amount)
}

// CreateMintRequest is a paid mutator transaction binding the contract method 0xcfec5ce8.
//
// Solidity: function createMintRequest(string sender, uint256 value) returns(bytes32)
func (_EthSusyContract *EthSusyContractTransactor) CreateMintRequest(opts *bind.TransactOpts, sender string, value *big.Int) (*types.Transaction, error) {
	return _EthSusyContract.contract.Transact(opts, "createMintRequest", sender, value)
}

// CreateMintRequest is a paid mutator transaction binding the contract method 0xcfec5ce8.
//
// Solidity: function createMintRequest(string sender, uint256 value) returns(bytes32)
func (_EthSusyContract *EthSusyContractSession) CreateMintRequest(sender string, value *big.Int) (*types.Transaction, error) {
	return _EthSusyContract.Contract.CreateMintRequest(&_EthSusyContract.TransactOpts, sender, value)
}

// CreateMintRequest is a paid mutator transaction binding the contract method 0xcfec5ce8.
//
// Solidity: function createMintRequest(string sender, uint256 value) returns(bytes32)
func (_EthSusyContract *EthSusyContractTransactorSession) CreateMintRequest(sender string, value *big.Int) (*types.Transaction, error) {
	return _EthSusyContract.Contract.CreateMintRequest(&_EthSusyContract.TransactOpts, sender, value)
}

// EthSusyContractMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the EthSusyContract contract.
type EthSusyContractMintIterator struct {
	Event *EthSusyContractMint // Event containing the contract specifics and raw log

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
func (it *EthSusyContractMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthSusyContractMint)
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
		it.Event = new(EthSusyContractMint)
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
func (it *EthSusyContractMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthSusyContractMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthSusyContractMint represents a Mint event raised by the EthSusyContract contract.
type EthSusyContractMint struct {
	RequestHash [32]byte
	Owner       common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xf0bcd414edc501835c6c541058514fe45a43fa62a41671c9f62317c9889fcf7f.
//
// Solidity: event Mint(bytes32 requestHash, address owner, uint256 amount)
func (_EthSusyContract *EthSusyContractFilterer) FilterMint(opts *bind.FilterOpts) (*EthSusyContractMintIterator, error) {

	logs, sub, err := _EthSusyContract.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &EthSusyContractMintIterator{contract: _EthSusyContract.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xf0bcd414edc501835c6c541058514fe45a43fa62a41671c9f62317c9889fcf7f.
//
// Solidity: event Mint(bytes32 requestHash, address owner, uint256 amount)
func (_EthSusyContract *EthSusyContractFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *EthSusyContractMint) (event.Subscription, error) {

	logs, sub, err := _EthSusyContract.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthSusyContractMint)
				if err := _EthSusyContract.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xf0bcd414edc501835c6c541058514fe45a43fa62a41671c9f62317c9889fcf7f.
//
// Solidity: event Mint(bytes32 requestHash, address owner, uint256 amount)
func (_EthSusyContract *EthSusyContractFilterer) ParseMint(log types.Log) (*EthSusyContractMint, error) {
	event := new(EthSusyContractMint)
	if err := _EthSusyContract.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthSusyContractNewRequestIterator is returned from FilterNewRequest and is used to iterate over the raw logs and unpacked data for NewRequest events raised by the EthSusyContract contract.
type EthSusyContractNewRequestIterator struct {
	Event *EthSusyContractNewRequest // Event containing the contract specifics and raw log

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
func (it *EthSusyContractNewRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthSusyContractNewRequest)
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
		it.Event = new(EthSusyContractNewRequest)
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
func (it *EthSusyContractNewRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthSusyContractNewRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthSusyContractNewRequest represents a NewRequest event raised by the EthSusyContract contract.
type EthSusyContractNewRequest struct {
	RequestHash [32]byte
	Owner       common.Address
	Target      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewRequest is a free log retrieval operation binding the contract event 0xe5048cb72529de5404e374d7fdbda8558caeb479015b507903292cf51c115896.
//
// Solidity: event NewRequest(bytes32 requestHash, address owner, string target)
func (_EthSusyContract *EthSusyContractFilterer) FilterNewRequest(opts *bind.FilterOpts) (*EthSusyContractNewRequestIterator, error) {

	logs, sub, err := _EthSusyContract.contract.FilterLogs(opts, "NewRequest")
	if err != nil {
		return nil, err
	}
	return &EthSusyContractNewRequestIterator{contract: _EthSusyContract.contract, event: "NewRequest", logs: logs, sub: sub}, nil
}

// WatchNewRequest is a free log subscription operation binding the contract event 0xe5048cb72529de5404e374d7fdbda8558caeb479015b507903292cf51c115896.
//
// Solidity: event NewRequest(bytes32 requestHash, address owner, string target)
func (_EthSusyContract *EthSusyContractFilterer) WatchNewRequest(opts *bind.WatchOpts, sink chan<- *EthSusyContractNewRequest) (event.Subscription, error) {

	logs, sub, err := _EthSusyContract.contract.WatchLogs(opts, "NewRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthSusyContractNewRequest)
				if err := _EthSusyContract.contract.UnpackLog(event, "NewRequest", log); err != nil {
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

// ParseNewRequest is a log parse operation binding the contract event 0xe5048cb72529de5404e374d7fdbda8558caeb479015b507903292cf51c115896.
//
// Solidity: event NewRequest(bytes32 requestHash, address owner, string target)
func (_EthSusyContract *EthSusyContractFilterer) ParseNewRequest(log types.Log) (*EthSusyContractNewRequest, error) {
	event := new(EthSusyContractNewRequest)
	if err := _EthSusyContract.contract.UnpackLog(event, "NewRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthSusyContractReturnIterator is returned from FilterReturn and is used to iterate over the raw logs and unpacked data for Return events raised by the EthSusyContract contract.
type EthSusyContractReturnIterator struct {
	Event *EthSusyContractReturn // Event containing the contract specifics and raw log

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
func (it *EthSusyContractReturnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthSusyContractReturn)
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
		it.Event = new(EthSusyContractReturn)
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
func (it *EthSusyContractReturnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthSusyContractReturnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthSusyContractReturn represents a Return event raised by the EthSusyContract contract.
type EthSusyContractReturn struct {
	RequestHash [32]byte
	Owner       common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReturn is a free log retrieval operation binding the contract event 0x86bd878938275851d348ab8cd7e09276a8abb4b70910d2e9e913a38815997313.
//
// Solidity: event Return(bytes32 requestHash, address owner, uint256 amount)
func (_EthSusyContract *EthSusyContractFilterer) FilterReturn(opts *bind.FilterOpts) (*EthSusyContractReturnIterator, error) {

	logs, sub, err := _EthSusyContract.contract.FilterLogs(opts, "Return")
	if err != nil {
		return nil, err
	}
	return &EthSusyContractReturnIterator{contract: _EthSusyContract.contract, event: "Return", logs: logs, sub: sub}, nil
}

// WatchReturn is a free log subscription operation binding the contract event 0x86bd878938275851d348ab8cd7e09276a8abb4b70910d2e9e913a38815997313.
//
// Solidity: event Return(bytes32 requestHash, address owner, uint256 amount)
func (_EthSusyContract *EthSusyContractFilterer) WatchReturn(opts *bind.WatchOpts, sink chan<- *EthSusyContractReturn) (event.Subscription, error) {

	logs, sub, err := _EthSusyContract.contract.WatchLogs(opts, "Return")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthSusyContractReturn)
				if err := _EthSusyContract.contract.UnpackLog(event, "Return", log); err != nil {
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

// ParseReturn is a log parse operation binding the contract event 0x86bd878938275851d348ab8cd7e09276a8abb4b70910d2e9e913a38815997313.
//
// Solidity: event Return(bytes32 requestHash, address owner, uint256 amount)
func (_EthSusyContract *EthSusyContractFilterer) ParseReturn(log types.Log) (*EthSusyContractReturn, error) {
	event := new(EthSusyContractReturn)
	if err := _EthSusyContract.contract.UnpackLog(event, "Return", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthSusyContractStatusChangedIterator is returned from FilterStatusChanged and is used to iterate over the raw logs and unpacked data for StatusChanged events raised by the EthSusyContract contract.
type EthSusyContractStatusChangedIterator struct {
	Event *EthSusyContractStatusChanged // Event containing the contract specifics and raw log

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
func (it *EthSusyContractStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthSusyContractStatusChanged)
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
		it.Event = new(EthSusyContractStatusChanged)
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
func (it *EthSusyContractStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthSusyContractStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthSusyContractStatusChanged represents a StatusChanged event raised by the EthSusyContract contract.
type EthSusyContractStatusChanged struct {
	RequestHash [32]byte
	Status      uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStatusChanged is a free log retrieval operation binding the contract event 0x22ce671cd3f328f71ef109d846af764df2d16bb8eb664c4f41065656c6e9962c.
//
// Solidity: event StatusChanged(bytes32 requestHash, uint8 status)
func (_EthSusyContract *EthSusyContractFilterer) FilterStatusChanged(opts *bind.FilterOpts) (*EthSusyContractStatusChangedIterator, error) {

	logs, sub, err := _EthSusyContract.contract.FilterLogs(opts, "StatusChanged")
	if err != nil {
		return nil, err
	}
	return &EthSusyContractStatusChangedIterator{contract: _EthSusyContract.contract, event: "StatusChanged", logs: logs, sub: sub}, nil
}

// WatchStatusChanged is a free log subscription operation binding the contract event 0x22ce671cd3f328f71ef109d846af764df2d16bb8eb664c4f41065656c6e9962c.
//
// Solidity: event StatusChanged(bytes32 requestHash, uint8 status)
func (_EthSusyContract *EthSusyContractFilterer) WatchStatusChanged(opts *bind.WatchOpts, sink chan<- *EthSusyContractStatusChanged) (event.Subscription, error) {

	logs, sub, err := _EthSusyContract.contract.WatchLogs(opts, "StatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthSusyContractStatusChanged)
				if err := _EthSusyContract.contract.UnpackLog(event, "StatusChanged", log); err != nil {
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

// ParseStatusChanged is a log parse operation binding the contract event 0x22ce671cd3f328f71ef109d846af764df2d16bb8eb664c4f41065656c6e9962c.
//
// Solidity: event StatusChanged(bytes32 requestHash, uint8 status)
func (_EthSusyContract *EthSusyContractFilterer) ParseStatusChanged(log types.Log) (*EthSusyContractStatusChanged, error) {
	event := new(EthSusyContractStatusChanged)
	if err := _EthSusyContract.contract.UnpackLog(event, "StatusChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
