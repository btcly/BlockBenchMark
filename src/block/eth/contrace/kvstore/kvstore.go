// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kvstore

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

// KvstoreMetaData contains all meta data concerning the Kvstore contract.
var KvstoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610794806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063693ec85e1461003b578063e942b5161461006b575b600080fd5b610055600480360381019061005091906102c1565b610087565b6040516100629190610389565b60405180910390f35b610085600480360381019061008091906103ab565b610137565b005b6060600082604051610099919061045f565b908152602001604051809103902080546100b2906104a5565b80601f01602080910402602001604051908101604052809291908181526020018280546100de906104a5565b801561012b5780601f106101005761010080835404028352916020019161012b565b820191906000526020600020905b81548152906001019060200180831161010e57829003601f168201915b50505050509050919050565b80600083604051610148919061045f565b90815260200160405180910390209081610162919061068c565b505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6101ce82610185565b810181811067ffffffffffffffff821117156101ed576101ec610196565b5b80604052505050565b6000610200610167565b905061020c82826101c5565b919050565b600067ffffffffffffffff82111561022c5761022b610196565b5b61023582610185565b9050602081019050919050565b82818337600083830152505050565b600061026461025f84610211565b6101f6565b9050828152602081018484840111156102805761027f610180565b5b61028b848285610242565b509392505050565b600082601f8301126102a8576102a761017b565b5b81356102b8848260208601610251565b91505092915050565b6000602082840312156102d7576102d6610171565b5b600082013567ffffffffffffffff8111156102f5576102f4610176565b5b61030184828501610293565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610344578082015181840152602081019050610329565b60008484015250505050565b600061035b8261030a565b6103658185610315565b9350610375818560208601610326565b61037e81610185565b840191505092915050565b600060208201905081810360008301526103a38184610350565b905092915050565b600080604083850312156103c2576103c1610171565b5b600083013567ffffffffffffffff8111156103e0576103df610176565b5b6103ec85828601610293565b925050602083013567ffffffffffffffff81111561040d5761040c610176565b5b61041985828601610293565b9150509250929050565b600081905092915050565b60006104398261030a565b6104438185610423565b9350610453818560208601610326565b80840191505092915050565b600061046b828461042e565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806104bd57607f821691505b6020821081036104d0576104cf610476565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026105387fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826104fb565b61054286836104fb565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061058961058461057f8461055a565b610564565b61055a565b9050919050565b6000819050919050565b6105a38361056e565b6105b76105af82610590565b848454610508565b825550505050565b600090565b6105cc6105bf565b6105d781848461059a565b505050565b5b818110156105fb576105f06000826105c4565b6001810190506105dd565b5050565b601f82111561064057610611816104d6565b61061a846104eb565b81016020851015610629578190505b61063d610635856104eb565b8301826105dc565b50505b505050565b600082821c905092915050565b600061066360001984600802610645565b1980831691505092915050565b600061067c8383610652565b9150826002028217905092915050565b6106958261030a565b67ffffffffffffffff8111156106ae576106ad610196565b5b6106b882546104a5565b6106c38282856105ff565b600060209050601f8311600181146106f657600084156106e4578287015190505b6106ee8582610670565b865550610756565b601f198416610704866104d6565b60005b8281101561072c57848901518255600182019150602085019450602081019050610707565b868310156107495784890151610745601f891682610652565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220e8960958d5c74e44af7dd91e679d28bf85b4cddbb0c42ebf36cbc82dc110feee64736f6c63430008100033",
}

// KvstoreABI is the input ABI used to generate the binding from.
// Deprecated: Use KvstoreMetaData.ABI instead.
var KvstoreABI = KvstoreMetaData.ABI

// KvstoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KvstoreMetaData.Bin instead.
var KvstoreBin = KvstoreMetaData.Bin

// DeployKvstore deploys a new Ethereum contract, binding an instance of Kvstore to it.
func DeployKvstore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Kvstore, error) {
	parsed, err := KvstoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KvstoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Kvstore{KvstoreCaller: KvstoreCaller{contract: contract}, KvstoreTransactor: KvstoreTransactor{contract: contract}, KvstoreFilterer: KvstoreFilterer{contract: contract}}, nil
}

// Kvstore is an auto generated Go binding around an Ethereum contract.
type Kvstore struct {
	KvstoreCaller     // Read-only binding to the contract
	KvstoreTransactor // Write-only binding to the contract
	KvstoreFilterer   // Log filterer for contract events
}

// KvstoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type KvstoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KvstoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KvstoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KvstoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KvstoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KvstoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KvstoreSession struct {
	Contract     *Kvstore          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KvstoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KvstoreCallerSession struct {
	Contract *KvstoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// KvstoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KvstoreTransactorSession struct {
	Contract     *KvstoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KvstoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type KvstoreRaw struct {
	Contract *Kvstore // Generic contract binding to access the raw methods on
}

// KvstoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KvstoreCallerRaw struct {
	Contract *KvstoreCaller // Generic read-only contract binding to access the raw methods on
}

// KvstoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KvstoreTransactorRaw struct {
	Contract *KvstoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKvstore creates a new instance of Kvstore, bound to a specific deployed contract.
func NewKvstore(address common.Address, backend bind.ContractBackend) (*Kvstore, error) {
	contract, err := bindKvstore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kvstore{KvstoreCaller: KvstoreCaller{contract: contract}, KvstoreTransactor: KvstoreTransactor{contract: contract}, KvstoreFilterer: KvstoreFilterer{contract: contract}}, nil
}

// NewKvstoreCaller creates a new read-only instance of Kvstore, bound to a specific deployed contract.
func NewKvstoreCaller(address common.Address, caller bind.ContractCaller) (*KvstoreCaller, error) {
	contract, err := bindKvstore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KvstoreCaller{contract: contract}, nil
}

// NewKvstoreTransactor creates a new write-only instance of Kvstore, bound to a specific deployed contract.
func NewKvstoreTransactor(address common.Address, transactor bind.ContractTransactor) (*KvstoreTransactor, error) {
	contract, err := bindKvstore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KvstoreTransactor{contract: contract}, nil
}

// NewKvstoreFilterer creates a new log filterer instance of Kvstore, bound to a specific deployed contract.
func NewKvstoreFilterer(address common.Address, filterer bind.ContractFilterer) (*KvstoreFilterer, error) {
	contract, err := bindKvstore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KvstoreFilterer{contract: contract}, nil
}

// bindKvstore binds a generic wrapper to an already deployed contract.
func bindKvstore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KvstoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kvstore *KvstoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kvstore.Contract.KvstoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kvstore *KvstoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kvstore.Contract.KvstoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kvstore *KvstoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kvstore.Contract.KvstoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kvstore *KvstoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kvstore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kvstore *KvstoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kvstore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kvstore *KvstoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kvstore.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) view returns(string)
func (_Kvstore *KvstoreCaller) Get(opts *bind.CallOpts, key string) (string, error) {
	var out []interface{}
	err := _Kvstore.contract.Call(opts, &out, "get", key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) view returns(string)
func (_Kvstore *KvstoreSession) Get(key string) (string, error) {
	return _Kvstore.Contract.Get(&_Kvstore.CallOpts, key)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) view returns(string)
func (_Kvstore *KvstoreCallerSession) Get(key string) (string, error) {
	return _Kvstore.Contract.Get(&_Kvstore.CallOpts, key)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_Kvstore *KvstoreTransactor) Set(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _Kvstore.contract.Transact(opts, "set", key, value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_Kvstore *KvstoreSession) Set(key string, value string) (*types.Transaction, error) {
	return _Kvstore.Contract.Set(&_Kvstore.TransactOpts, key, value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_Kvstore *KvstoreTransactorSession) Set(key string, value string) (*types.Transaction, error) {
	return _Kvstore.Contract.Set(&_Kvstore.TransactOpts, key, value)
}
