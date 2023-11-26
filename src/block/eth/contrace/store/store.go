// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"vers\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"data_contract\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getItem\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionContract\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version_\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000f6d38038062000f6d8339818101604052810190620000379190620001e3565b80600190816200004891906200047f565b505062000566565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620000b9826200006e565b810181811067ffffffffffffffff82111715620000db57620000da6200007f565b5b80604052505050565b6000620000f062000050565b9050620000fe8282620000ae565b919050565b600067ffffffffffffffff8211156200012157620001206200007f565b5b6200012c826200006e565b9050602081019050919050565b60005b83811015620001595780820151818401526020810190506200013c565b60008484015250505050565b60006200017c620001768462000103565b620000e4565b9050828152602081018484840111156200019b576200019a62000069565b5b620001a884828562000139565b509392505050565b600082601f830112620001c857620001c762000064565b5b8151620001da84826020860162000165565b91505092915050565b600060208284031215620001fc57620001fb6200005a565b5b600082015167ffffffffffffffff8111156200021d576200021c6200005f565b5b6200022b84828501620001b0565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200028757607f821691505b6020821081036200029d576200029c6200023f565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003077fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620002c8565b620003138683620002c8565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003606200035a62000354846200032b565b62000335565b6200032b565b9050919050565b6000819050919050565b6200037c836200033f565b620003946200038b8262000367565b848454620002d5565b825550505050565b600090565b620003ab6200039c565b620003b881848462000371565b505050565b5b81811015620003e057620003d4600082620003a1565b600181019050620003be565b5050565b601f8211156200042f57620003f981620002a3565b6200040484620002b8565b8101602085101562000414578190505b6200042c6200042385620002b8565b830182620003bd565b50505b505050565b600082821c905092915050565b6000620004546000198460080262000434565b1980831691505092915050565b60006200046f838362000441565b9150826002028217905092915050565b6200048a8262000234565b67ffffffffffffffff811115620004a657620004a56200007f565b5b620004b282546200026e565b620004bf828285620003e4565b600060209050601f831160018114620004f75760008415620004e2578287015190505b620004ee858262000461565b8655506200055e565b601f1984166200050786620002a3565b60005b8281101562000531578489015182556001820191506020850194506020810190506200050a565b868310156200055157848901516200054d601f89168262000441565b8355505b6001600288020188555050505b505050505050565b6109f780620005766000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80637a4c982e1461005c578063a882d86614610078578063b6010fcd14610096578063dd1bb1fc146100c6578063fb4b5411146100e4575b600080fd5b61007660048036038101906100719190610524565b610114565b005b610080610144565b60405161008d919061061b565b60405180910390f35b6100b060048036038101906100ab919061063d565b6101d6565b6040516100bd919061061b565b60405180910390f35b6100ce610286565b6040516100db919061061b565b60405180910390f35b6100fe60048036038101906100f9919061063d565b610314565b60405161010b919061061b565b60405180910390f35b8060008360405161012591906106c2565b9081526020016040518091039020908161013f91906108ef565b505050565b60606001805461015390610708565b80601f016020809104026020016040519081016040528092919081815260200182805461017f90610708565b80156101cc5780601f106101a1576101008083540402835291602001916101cc565b820191906000526020600020905b8154815290600101906020018083116101af57829003601f168201915b5050505050905090565b60606000826040516101e891906106c2565b9081526020016040518091039020805461020190610708565b80601f016020809104026020016040519081016040528092919081815260200182805461022d90610708565b801561027a5780601f1061024f5761010080835404028352916020019161027a565b820191906000526020600020905b81548152906001019060200180831161025d57829003601f168201915b50505050509050919050565b6001805461029390610708565b80601f01602080910402602001604051908101604052809291908181526020018280546102bf90610708565b801561030c5780601f106102e15761010080835404028352916020019161030c565b820191906000526020600020905b8154815290600101906020018083116102ef57829003601f168201915b505050505081565b600081805160208101820180518482526020830160208501208183528095505050505050600091509050805461034990610708565b80601f016020809104026020016040519081016040528092919081815260200182805461037590610708565b80156103c25780601f10610397576101008083540402835291602001916103c2565b820191906000526020600020905b8154815290600101906020018083116103a557829003601f168201915b505050505081565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610431826103e8565b810181811067ffffffffffffffff821117156104505761044f6103f9565b5b80604052505050565b60006104636103ca565b905061046f8282610428565b919050565b600067ffffffffffffffff82111561048f5761048e6103f9565b5b610498826103e8565b9050602081019050919050565b82818337600083830152505050565b60006104c76104c284610474565b610459565b9050828152602081018484840111156104e3576104e26103e3565b5b6104ee8482856104a5565b509392505050565b600082601f83011261050b5761050a6103de565b5b813561051b8482602086016104b4565b91505092915050565b6000806040838503121561053b5761053a6103d4565b5b600083013567ffffffffffffffff811115610559576105586103d9565b5b610565858286016104f6565b925050602083013567ffffffffffffffff811115610586576105856103d9565b5b610592858286016104f6565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b838110156105d65780820151818401526020810190506105bb565b60008484015250505050565b60006105ed8261059c565b6105f781856105a7565b93506106078185602086016105b8565b610610816103e8565b840191505092915050565b6000602082019050818103600083015261063581846105e2565b905092915050565b600060208284031215610653576106526103d4565b5b600082013567ffffffffffffffff811115610671576106706103d9565b5b61067d848285016104f6565b91505092915050565b600081905092915050565b600061069c8261059c565b6106a68185610686565b93506106b68185602086016105b8565b80840191505092915050565b60006106ce8284610691565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061072057607f821691505b602082108103610733576107326106d9565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261079b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261075e565b6107a5868361075e565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006107ec6107e76107e2846107bd565b6107c7565b6107bd565b9050919050565b6000819050919050565b610806836107d1565b61081a610812826107f3565b84845461076b565b825550505050565b600090565b61082f610822565b61083a8184846107fd565b505050565b5b8181101561085e57610853600082610827565b600181019050610840565b5050565b601f8211156108a35761087481610739565b61087d8461074e565b8101602085101561088c578190505b6108a06108988561074e565b83018261083f565b50505b505050565b600082821c905092915050565b60006108c6600019846008026108a8565b1980831691505092915050565b60006108df83836108b5565b9150826002028217905092915050565b6108f88261059c565b67ffffffffffffffff811115610911576109106103f9565b5b61091b8254610708565b610926828285610862565b600060209050601f8311600181146109595760008415610947578287015190505b61095185826108d3565b8655506109b9565b601f19841661096786610739565b60005b8281101561098f5784890151825560018201915060208501945060208101905061096a565b868310156109ac57848901516109a8601f8916826108b5565b8355505b6001600288020188555050505b50505050505056fea26469706673582212206f8fcd81a6e618903e396cbb51c563e88a566d0bf01cb32ce78112135a92113264736f6c63430008100033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, vers string) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend, vers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// DataContract is a free data retrieval call binding the contract method 0xfb4b5411.
//
// Solidity: function data_contract(string ) view returns(string)
func (_Store *StoreCaller) DataContract(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "data_contract", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DataContract is a free data retrieval call binding the contract method 0xfb4b5411.
//
// Solidity: function data_contract(string ) view returns(string)
func (_Store *StoreSession) DataContract(arg0 string) (string, error) {
	return _Store.Contract.DataContract(&_Store.CallOpts, arg0)
}

// DataContract is a free data retrieval call binding the contract method 0xfb4b5411.
//
// Solidity: function data_contract(string ) view returns(string)
func (_Store *StoreCallerSession) DataContract(arg0 string) (string, error) {
	return _Store.Contract.DataContract(&_Store.CallOpts, arg0)
}

// GetItem is a free data retrieval call binding the contract method 0xb6010fcd.
//
// Solidity: function getItem(string key) view returns(string)
func (_Store *StoreCaller) GetItem(opts *bind.CallOpts, key string) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getItem", key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetItem is a free data retrieval call binding the contract method 0xb6010fcd.
//
// Solidity: function getItem(string key) view returns(string)
func (_Store *StoreSession) GetItem(key string) (string, error) {
	return _Store.Contract.GetItem(&_Store.CallOpts, key)
}

// GetItem is a free data retrieval call binding the contract method 0xb6010fcd.
//
// Solidity: function getItem(string key) view returns(string)
func (_Store *StoreCallerSession) GetItem(key string) (string, error) {
	return _Store.Contract.GetItem(&_Store.CallOpts, key)
}

// VersionContract is a free data retrieval call binding the contract method 0xa882d866.
//
// Solidity: function versionContract() view returns(string)
func (_Store *StoreCaller) VersionContract(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "versionContract")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionContract is a free data retrieval call binding the contract method 0xa882d866.
//
// Solidity: function versionContract() view returns(string)
func (_Store *StoreSession) VersionContract() (string, error) {
	return _Store.Contract.VersionContract(&_Store.CallOpts)
}

// VersionContract is a free data retrieval call binding the contract method 0xa882d866.
//
// Solidity: function versionContract() view returns(string)
func (_Store *StoreCallerSession) VersionContract() (string, error) {
	return _Store.Contract.VersionContract(&_Store.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xdd1bb1fc.
//
// Solidity: function version_() view returns(string)
func (_Store *StoreCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "version_")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0xdd1bb1fc.
//
// Solidity: function version_() view returns(string)
func (_Store *StoreSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xdd1bb1fc.
//
// Solidity: function version_() view returns(string)
func (_Store *StoreCallerSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// SetItem is a paid mutator transaction binding the contract method 0x7a4c982e.
//
// Solidity: function setItem(string key, string value) returns()
func (_Store *StoreTransactor) SetItem(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setItem", key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0x7a4c982e.
//
// Solidity: function setItem(string key, string value) returns()
func (_Store *StoreSession) SetItem(key string, value string) (*types.Transaction, error) {
	return _Store.Contract.SetItem(&_Store.TransactOpts, key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0x7a4c982e.
//
// Solidity: function setItem(string key, string value) returns()
func (_Store *StoreTransactorSession) SetItem(key string, value string) (*types.Transaction, error) {
	return _Store.Contract.SetItem(&_Store.TransactOpts, key, value)
}
