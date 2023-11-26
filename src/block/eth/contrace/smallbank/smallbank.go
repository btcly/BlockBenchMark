// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smallbank

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SmallbankMetaData contains all meta data concerning the Smallbank contract.
var SmallbankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg0\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"arg1\",\"type\":\"string\"}],\"name\":\"almagate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg0\",\"type\":\"string\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg0\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"arg1\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"arg2\",\"type\":\"uint256\"}],\"name\":\"sendPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg0\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"arg1\",\"type\":\"uint256\"}],\"name\":\"updateBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg0\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"arg1\",\"type\":\"uint256\"}],\"name\":\"updateSaving\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg0\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"arg1\",\"type\":\"uint256\"}],\"name\":\"writeCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610911806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630b488b37146100675780630be8374d146100835780633a51d2461461009f578063870187eb146100cf578063901d706f146100eb578063ca30543514610107575b600080fd5b610081600480360381019061007c91906105ea565b610123565b005b61009d600480360381019061009891906105ea565b61017e565b005b6100b960048036038101906100b49190610646565b610250565b6040516100c6919061069e565b60405180910390f35b6100e960048036038101906100e491906105ea565b6102ac565b005b610105600480360381019061010091906106b9565b610308565b005b610121600480360381019061011c9190610731565b6103a5565b005b60008083604051610134919061082d565b9081526020016040518091039020549050600082905080826101569190610873565b600085604051610166919061082d565b90815260200160405180910390208190555050505050565b6000600183604051610190919061082d565b9081526020016040518091039020549050600080846040516101b2919061082d565b9081526020016040518091039020549050600083905081836101d49190610873565b81101561021a57600181846101e991906108a7565b6101f391906108a7565b600186604051610203919061082d565b908152602001604051809103902081905550610249565b808361022691906108a7565b600186604051610236919061082d565b9081526020016040518091039020819055505b5050505050565b600080600083604051610263919061082d565b90815260200160405180910390205490506000600184604051610286919061082d565b908152602001604051809103902054905080826102a39190610873565b92505050919050565b60006001836040516102be919061082d565b9081526020016040518091039020549050600082905080826102e09190610873565b6001856040516102f0919061082d565b90815260200160405180910390208190555050505050565b60008083604051610319919061082d565b9081526020016040518091039020549050600060018360405161033c919061082d565b9081526020016040518091039020549050600060018560405161035f919061082d565b908152602001604051809103902081905550808261037d9190610873565b60008460405161038d919061082d565b90815260200160405180910390208190555050505050565b60006001846040516103b7919061082d565b908152602001604051809103902054905060006001846040516103da919061082d565b9081526020016040518091039020549050600083905080836103fc91906108a7565b9250808261040a9190610873565b91508260018760405161041d919061082d565b90815260200160405180910390208190555081600186604051610440919061082d565b908152602001604051809103902081905550505050505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6104c182610478565b810181811067ffffffffffffffff821117156104e0576104df610489565b5b80604052505050565b60006104f361045a565b90506104ff82826104b8565b919050565b600067ffffffffffffffff82111561051f5761051e610489565b5b61052882610478565b9050602081019050919050565b82818337600083830152505050565b600061055761055284610504565b6104e9565b90508281526020810184848401111561057357610572610473565b5b61057e848285610535565b509392505050565b600082601f83011261059b5761059a61046e565b5b81356105ab848260208601610544565b91505092915050565b6000819050919050565b6105c7816105b4565b81146105d257600080fd5b50565b6000813590506105e4816105be565b92915050565b6000806040838503121561060157610600610464565b5b600083013567ffffffffffffffff81111561061f5761061e610469565b5b61062b85828601610586565b925050602061063c858286016105d5565b9150509250929050565b60006020828403121561065c5761065b610464565b5b600082013567ffffffffffffffff81111561067a57610679610469565b5b61068684828501610586565b91505092915050565b610698816105b4565b82525050565b60006020820190506106b3600083018461068f565b92915050565b600080604083850312156106d0576106cf610464565b5b600083013567ffffffffffffffff8111156106ee576106ed610469565b5b6106fa85828601610586565b925050602083013567ffffffffffffffff81111561071b5761071a610469565b5b61072785828601610586565b9150509250929050565b60008060006060848603121561074a57610749610464565b5b600084013567ffffffffffffffff81111561076857610767610469565b5b61077486828701610586565b935050602084013567ffffffffffffffff81111561079557610794610469565b5b6107a186828701610586565b92505060406107b2868287016105d5565b9150509250925092565b600081519050919050565b600081905092915050565b60005b838110156107f05780820151818401526020810190506107d5565b60008484015250505050565b6000610807826107bc565b61081181856107c7565b93506108218185602086016107d2565b80840191505092915050565b600061083982846107fc565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061087e826105b4565b9150610889836105b4565b92508282019050808211156108a1576108a0610844565b5b92915050565b60006108b2826105b4565b91506108bd836105b4565b92508282039050818111156108d5576108d4610844565b5b9291505056fea26469706673582212202f244d50610cd76c0c504503ae44310e9aa43b2f1ba117617d012e708662588e64736f6c63430008100033",
}

// SmallbankABI is the input ABI used to generate the binding from.
// Deprecated: Use SmallbankMetaData.ABI instead.
var SmallbankABI = SmallbankMetaData.ABI

// SmallbankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SmallbankMetaData.Bin instead.
var SmallbankBin = SmallbankMetaData.Bin

// DeploySmallbank deploys a new Ethereum contract, binding an instance of Smallbank to it.
func DeploySmallbank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Smallbank, error) {
	parsed, err := SmallbankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SmallbankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Smallbank{SmallbankCaller: SmallbankCaller{contract: contract}, SmallbankTransactor: SmallbankTransactor{contract: contract}, SmallbankFilterer: SmallbankFilterer{contract: contract}}, nil
}

// Smallbank is an auto generated Go binding around an Ethereum contract.
type Smallbank struct {
	SmallbankCaller     // Read-only binding to the contract
	SmallbankTransactor // Write-only binding to the contract
	SmallbankFilterer   // Log filterer for contract events
}

// SmallbankCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmallbankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmallbankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmallbankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmallbankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmallbankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmallbankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmallbankSession struct {
	Contract     *Smallbank        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmallbankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmallbankCallerSession struct {
	Contract *SmallbankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SmallbankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmallbankTransactorSession struct {
	Contract     *SmallbankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SmallbankRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmallbankRaw struct {
	Contract *Smallbank // Generic contract binding to access the raw methods on
}

// SmallbankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmallbankCallerRaw struct {
	Contract *SmallbankCaller // Generic read-only contract binding to access the raw methods on
}

// SmallbankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmallbankTransactorRaw struct {
	Contract *SmallbankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmallbank creates a new instance of Smallbank, bound to a specific deployed contract.
func NewSmallbank(address common.Address, backend bind.ContractBackend) (*Smallbank, error) {
	contract, err := bindSmallbank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smallbank{SmallbankCaller: SmallbankCaller{contract: contract}, SmallbankTransactor: SmallbankTransactor{contract: contract}, SmallbankFilterer: SmallbankFilterer{contract: contract}}, nil
}

// NewSmallbankCaller creates a new read-only instance of Smallbank, bound to a specific deployed contract.
func NewSmallbankCaller(address common.Address, caller bind.ContractCaller) (*SmallbankCaller, error) {
	contract, err := bindSmallbank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmallbankCaller{contract: contract}, nil
}

// NewSmallbankTransactor creates a new write-only instance of Smallbank, bound to a specific deployed contract.
func NewSmallbankTransactor(address common.Address, transactor bind.ContractTransactor) (*SmallbankTransactor, error) {
	contract, err := bindSmallbank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmallbankTransactor{contract: contract}, nil
}

// NewSmallbankFilterer creates a new log filterer instance of Smallbank, bound to a specific deployed contract.
func NewSmallbankFilterer(address common.Address, filterer bind.ContractFilterer) (*SmallbankFilterer, error) {
	contract, err := bindSmallbank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmallbankFilterer{contract: contract}, nil
}

// bindSmallbank binds a generic wrapper to an already deployed contract.
func bindSmallbank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmallbankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smallbank *SmallbankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smallbank.Contract.SmallbankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smallbank *SmallbankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smallbank.Contract.SmallbankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smallbank *SmallbankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smallbank.Contract.SmallbankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smallbank *SmallbankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smallbank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smallbank *SmallbankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smallbank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smallbank *SmallbankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smallbank.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string arg0) view returns(uint256 balance)
func (_Smallbank *SmallbankCaller) GetBalance(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Smallbank.contract.Call(opts, &out, "getBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string arg0) view returns(uint256 balance)
func (_Smallbank *SmallbankSession) GetBalance(arg0 string) (*big.Int, error) {
	return _Smallbank.Contract.GetBalance(&_Smallbank.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string arg0) view returns(uint256 balance)
func (_Smallbank *SmallbankCallerSession) GetBalance(arg0 string) (*big.Int, error) {
	return _Smallbank.Contract.GetBalance(&_Smallbank.CallOpts, arg0)
}

// Almagate is a paid mutator transaction binding the contract method 0x901d706f.
//
// Solidity: function almagate(string arg0, string arg1) returns()
func (_Smallbank *SmallbankTransactor) Almagate(opts *bind.TransactOpts, arg0 string, arg1 string) (*types.Transaction, error) {
	return _Smallbank.contract.Transact(opts, "almagate", arg0, arg1)
}

// Almagate is a paid mutator transaction binding the contract method 0x901d706f.
//
// Solidity: function almagate(string arg0, string arg1) returns()
func (_Smallbank *SmallbankSession) Almagate(arg0 string, arg1 string) (*types.Transaction, error) {
	return _Smallbank.Contract.Almagate(&_Smallbank.TransactOpts, arg0, arg1)
}

// Almagate is a paid mutator transaction binding the contract method 0x901d706f.
//
// Solidity: function almagate(string arg0, string arg1) returns()
func (_Smallbank *SmallbankTransactorSession) Almagate(arg0 string, arg1 string) (*types.Transaction, error) {
	return _Smallbank.Contract.Almagate(&_Smallbank.TransactOpts, arg0, arg1)
}

// SendPayment is a paid mutator transaction binding the contract method 0xca305435.
//
// Solidity: function sendPayment(string arg0, string arg1, uint256 arg2) returns()
func (_Smallbank *SmallbankTransactor) SendPayment(opts *bind.TransactOpts, arg0 string, arg1 string, arg2 *big.Int) (*types.Transaction, error) {
	return _Smallbank.contract.Transact(opts, "sendPayment", arg0, arg1, arg2)
}

// SendPayment is a paid mutator transaction binding the contract method 0xca305435.
//
// Solidity: function sendPayment(string arg0, string arg1, uint256 arg2) returns()
func (_Smallbank *SmallbankSession) SendPayment(arg0 string, arg1 string, arg2 *big.Int) (*types.Transaction, error) {
	return _Smallbank.Contract.SendPayment(&_Smallbank.TransactOpts, arg0, arg1, arg2)
}

// SendPayment is a paid mutator transaction binding the contract method 0xca305435.
//
// Solidity: function sendPayment(string arg0, string arg1, uint256 arg2) returns()
func (_Smallbank *SmallbankTransactorSession) SendPayment(arg0 string, arg1 string, arg2 *big.Int) (*types.Transaction, error) {
	return _Smallbank.Contract.SendPayment(&_Smallbank.TransactOpts, arg0, arg1, arg2)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x870187eb.
//
// Solidity: function updateBalance(string arg0, uint256 arg1) returns()
func (_Smallbank *SmallbankTransactor) UpdateBalance(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Smallbank.contract.Transact(opts, "updateBalance", arg0, arg1)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x870187eb.
//
// Solidity: function updateBalance(string arg0, uint256 arg1) returns()
func (_Smallbank *SmallbankSession) UpdateBalance(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Smallbank.Contract.UpdateBalance(&_Smallbank.TransactOpts, arg0, arg1)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x870187eb.
//
// Solidity: function updateBalance(string arg0, uint256 arg1) returns()
func (_Smallbank *SmallbankTransactorSession) UpdateBalance(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Smallbank.Contract.UpdateBalance(&_Smallbank.TransactOpts, arg0, arg1)
}

// UpdateSaving is a paid mutator transaction binding the contract method 0x0b488b37.
//
// Solidity: function updateSaving(string arg0, uint256 arg1) returns()
func (_Smallbank *SmallbankTransactor) UpdateSaving(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Smallbank.contract.Transact(opts, "updateSaving", arg0, arg1)
}

// UpdateSaving is a paid mutator transaction binding the contract method 0x0b488b37.
//
// Solidity: function updateSaving(string arg0, uint256 arg1) returns()
func (_Smallbank *SmallbankSession) UpdateSaving(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Smallbank.Contract.UpdateSaving(&_Smallbank.TransactOpts, arg0, arg1)
}

// UpdateSaving is a paid mutator transaction binding the contract method 0x0b488b37.
//
// Solidity: function updateSaving(string arg0, uint256 arg1) returns()
func (_Smallbank *SmallbankTransactorSession) UpdateSaving(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Smallbank.Contract.UpdateSaving(&_Smallbank.TransactOpts, arg0, arg1)
}

// WriteCheck is a paid mutator transaction binding the contract method 0x0be8374d.
//
// Solidity: function writeCheck(string arg0, uint256 arg1) returns()
func (_Smallbank *SmallbankTransactor) WriteCheck(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Smallbank.contract.Transact(opts, "writeCheck", arg0, arg1)
}

// WriteCheck is a paid mutator transaction binding the contract method 0x0be8374d.
//
// Solidity: function writeCheck(string arg0, uint256 arg1) returns()
func (_Smallbank *SmallbankSession) WriteCheck(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Smallbank.Contract.WriteCheck(&_Smallbank.TransactOpts, arg0, arg1)
}

// WriteCheck is a paid mutator transaction binding the contract method 0x0be8374d.
//
// Solidity: function writeCheck(string arg0, uint256 arg1) returns()
func (_Smallbank *SmallbankTransactorSession) WriteCheck(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Smallbank.Contract.WriteCheck(&_Smallbank.TransactOpts, arg0, arg1)
}
