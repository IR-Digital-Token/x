// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vat

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
	_ = abi.ConvertType
)

// VatMetaData contains all meta data concerning the Vat contract.
var VatMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Cage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"}],\"name\":\"Fess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Flap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"era\",\"type\":\"uint256\"}],\"name\":\"Flog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Flop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Flux\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"rate\",\"type\":\"int256\"}],\"name\":\"Fold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"Fork\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"Frob\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"Grab\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"Heal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Hope\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"Kiss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"Move\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Nope\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"wad\",\"type\":\"int256\"}],\"name\":\"Slip\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"Suck\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Line\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"can\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"flux\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"rate\",\"type\":\"int256\"}],\"name\":\"fold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"fork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"frob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"grab\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"heal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"hope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"Art\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"line\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dust\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"nope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"wad\",\"type\":\"int256\"}],\"name\":\"slip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"suck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"urns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ink\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"art\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VatABI is the input ABI used to generate the binding from.
// Deprecated: Use VatMetaData.ABI instead.
var VatABI = VatMetaData.ABI

// Vat is an auto generated Go binding around an Ethereum contract.
type Vat struct {
	VatCaller     // Read-only binding to the contract
	VatTransactor // Write-only binding to the contract
	VatFilterer   // Log filterer for contract events
}

// VatCaller is an auto generated read-only Go binding around an Ethereum contract.
type VatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VatSession struct {
	Contract     *Vat              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VatCallerSession struct {
	Contract *VatCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VatTransactorSession struct {
	Contract     *VatTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VatRaw is an auto generated low-level Go binding around an Ethereum contract.
type VatRaw struct {
	Contract *Vat // Generic contract binding to access the raw methods on
}

// VatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VatCallerRaw struct {
	Contract *VatCaller // Generic read-only contract binding to access the raw methods on
}

// VatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VatTransactorRaw struct {
	Contract *VatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVat creates a new instance of Vat, bound to a specific deployed contract.
func NewVat(address common.Address, backend bind.ContractBackend) (*Vat, error) {
	contract, err := bindVat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vat{VatCaller: VatCaller{contract: contract}, VatTransactor: VatTransactor{contract: contract}, VatFilterer: VatFilterer{contract: contract}}, nil
}

// NewVatCaller creates a new read-only instance of Vat, bound to a specific deployed contract.
func NewVatCaller(address common.Address, caller bind.ContractCaller) (*VatCaller, error) {
	contract, err := bindVat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VatCaller{contract: contract}, nil
}

// NewVatTransactor creates a new write-only instance of Vat, bound to a specific deployed contract.
func NewVatTransactor(address common.Address, transactor bind.ContractTransactor) (*VatTransactor, error) {
	contract, err := bindVat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VatTransactor{contract: contract}, nil
}

// NewVatFilterer creates a new log filterer instance of Vat, bound to a specific deployed contract.
func NewVatFilterer(address common.Address, filterer bind.ContractFilterer) (*VatFilterer, error) {
	contract, err := bindVat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VatFilterer{contract: contract}, nil
}

// bindVat binds a generic wrapper to an already deployed contract.
func bindVat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VatMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vat *VatRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vat.Contract.VatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vat *VatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vat.Contract.VatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vat *VatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vat.Contract.VatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vat *VatCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vat *VatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vat *VatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vat.Contract.contract.Transact(opts, method, params...)
}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() view returns(uint256)
func (_Vat *VatCaller) Line(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "Line")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() view returns(uint256)
func (_Vat *VatSession) Line() (*big.Int, error) {
	return _Vat.Contract.Line(&_Vat.CallOpts)
}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() view returns(uint256)
func (_Vat *VatCallerSession) Line() (*big.Int, error) {
	return _Vat.Contract.Line(&_Vat.CallOpts)
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_Vat *VatCaller) Can(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "can", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_Vat *VatSession) Can(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vat.Contract.Can(&_Vat.CallOpts, arg0, arg1)
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_Vat *VatCallerSession) Can(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vat.Contract.Can(&_Vat.CallOpts, arg0, arg1)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_Vat *VatCaller) Debt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "debt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_Vat *VatSession) Debt() (*big.Int, error) {
	return _Vat.Contract.Debt(&_Vat.CallOpts)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_Vat *VatCallerSession) Debt() (*big.Int, error) {
	return _Vat.Contract.Debt(&_Vat.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) view returns(uint256)
func (_Vat *VatCaller) Gem(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "gem", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) view returns(uint256)
func (_Vat *VatSession) Gem(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _Vat.Contract.Gem(&_Vat.CallOpts, arg0, arg1)
}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) view returns(uint256)
func (_Vat *VatCallerSession) Gem(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _Vat.Contract.Gem(&_Vat.CallOpts, arg0, arg1)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_Vat *VatCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Art  *big.Int
		Rate *big.Int
		Spot *big.Int
		Line *big.Int
		Dust *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Art = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Rate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Spot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Line = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Dust = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_Vat *VatSession) Ilks(arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	return _Vat.Contract.Ilks(&_Vat.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_Vat *VatCallerSession) Ilks(arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	return _Vat.Contract.Ilks(&_Vat.CallOpts, arg0)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vat *VatCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vat *VatSession) Live() (*big.Int, error) {
	return _Vat.Contract.Live(&_Vat.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vat *VatCallerSession) Live() (*big.Int, error) {
	return _Vat.Contract.Live(&_Vat.CallOpts)
}

// Sim is a free data retrieval call binding the contract method 0x3171ae31.
//
// Solidity: function sim(address ) view returns(uint256)
func (_Vat *VatCaller) Sim(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "sim", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sim is a free data retrieval call binding the contract method 0x3171ae31.
//
// Solidity: function sim(address ) view returns(uint256)
func (_Vat *VatSession) Sim(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Sim(&_Vat.CallOpts, arg0)
}

// Sim is a free data retrieval call binding the contract method 0x3171ae31.
//
// Solidity: function sim(address ) view returns(uint256)
func (_Vat *VatCallerSession) Sim(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Sim(&_Vat.CallOpts, arg0)
}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) view returns(uint256)
func (_Vat *VatCaller) Sin(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "sin", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) view returns(uint256)
func (_Vat *VatSession) Sin(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Sin(&_Vat.CallOpts, arg0)
}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) view returns(uint256)
func (_Vat *VatCallerSession) Sin(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Sin(&_Vat.CallOpts, arg0)
}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) view returns(uint256 ink, uint256 art)
func (_Vat *VatCaller) Urns(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "urns", arg0, arg1)

	outstruct := new(struct {
		Ink *big.Int
		Art *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ink = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Art = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) view returns(uint256 ink, uint256 art)
func (_Vat *VatSession) Urns(arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	return _Vat.Contract.Urns(&_Vat.CallOpts, arg0, arg1)
}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) view returns(uint256 ink, uint256 art)
func (_Vat *VatCallerSession) Urns(arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	return _Vat.Contract.Urns(&_Vat.CallOpts, arg0, arg1)
}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() view returns(uint256)
func (_Vat *VatCaller) Vice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "vice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() view returns(uint256)
func (_Vat *VatSession) Vice() (*big.Int, error) {
	return _Vat.Contract.Vice(&_Vat.CallOpts)
}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() view returns(uint256)
func (_Vat *VatCallerSession) Vice() (*big.Int, error) {
	return _Vat.Contract.Vice(&_Vat.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vat *VatCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vat *VatSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Wards(&_Vat.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vat *VatCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Wards(&_Vat.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vat *VatTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vat *VatSession) Cage() (*types.Transaction, error) {
	return _Vat.Contract.Cage(&_Vat.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vat *VatTransactorSession) Cage() (*types.Transaction, error) {
	return _Vat.Contract.Cage(&_Vat.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vat *VatTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vat *VatSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Deny(&_Vat.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vat *VatTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Deny(&_Vat.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Vat *VatTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Vat *VatSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.File(&_Vat.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Vat *VatTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.File(&_Vat.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vat *VatTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vat *VatSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.File0(&_Vat.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vat *VatTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.File0(&_Vat.TransactOpts, what, data)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_Vat *VatTransactor) Flux(opts *bind.TransactOpts, ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "flux", ilk, src, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_Vat *VatSession) Flux(ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Flux(&_Vat.TransactOpts, ilk, src, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_Vat *VatTransactorSession) Flux(ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Flux(&_Vat.TransactOpts, ilk, src, dst, wad)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_Vat *VatTransactor) Fold(opts *bind.TransactOpts, i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "fold", i, u, rate)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_Vat *VatSession) Fold(i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Fold(&_Vat.TransactOpts, i, u, rate)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_Vat *VatTransactorSession) Fold(i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Fold(&_Vat.TransactOpts, i, u, rate)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_Vat *VatTransactor) Fork(opts *bind.TransactOpts, ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "fork", ilk, src, dst, dink, dart)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_Vat *VatSession) Fork(ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Fork(&_Vat.TransactOpts, ilk, src, dst, dink, dart)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_Vat *VatTransactorSession) Fork(ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Fork(&_Vat.TransactOpts, ilk, src, dst, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatTransactor) Frob(opts *bind.TransactOpts, i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "frob", i, u, v, w, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatSession) Frob(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Frob(&_Vat.TransactOpts, i, u, v, w, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatTransactorSession) Frob(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Frob(&_Vat.TransactOpts, i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatTransactor) Grab(opts *bind.TransactOpts, i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "grab", i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatSession) Grab(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Grab(&_Vat.TransactOpts, i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatTransactorSession) Grab(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Grab(&_Vat.TransactOpts, i, u, v, w, dink, dart)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vat *VatTransactor) Heal(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "heal", rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vat *VatSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Heal(&_Vat.TransactOpts, rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vat *VatTransactorSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Heal(&_Vat.TransactOpts, rad)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_Vat *VatTransactor) Hope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "hope", usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_Vat *VatSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Hope(&_Vat.TransactOpts, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_Vat *VatTransactorSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Hope(&_Vat.TransactOpts, usr)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Vat *VatTransactor) Init(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "init", ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Vat *VatSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Vat.Contract.Init(&_Vat.TransactOpts, ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Vat *VatTransactorSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Vat.Contract.Init(&_Vat.TransactOpts, ilk)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_Vat *VatTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "move", src, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_Vat *VatSession) Move(src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Move(&_Vat.TransactOpts, src, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_Vat *VatTransactorSession) Move(src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Move(&_Vat.TransactOpts, src, dst, rad)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_Vat *VatTransactor) Nope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "nope", usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_Vat *VatSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Nope(&_Vat.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_Vat *VatTransactorSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Nope(&_Vat.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vat *VatTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vat *VatSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Rely(&_Vat.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vat *VatTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Rely(&_Vat.TransactOpts, usr)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_Vat *VatTransactor) Slip(opts *bind.TransactOpts, ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "slip", ilk, usr, wad)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_Vat *VatSession) Slip(ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Slip(&_Vat.TransactOpts, ilk, usr, wad)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_Vat *VatTransactorSession) Slip(ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Slip(&_Vat.TransactOpts, ilk, usr, wad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_Vat *VatTransactor) Suck(opts *bind.TransactOpts, u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "suck", u, v, rad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_Vat *VatSession) Suck(u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Suck(&_Vat.TransactOpts, u, v, rad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_Vat *VatTransactorSession) Suck(u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Suck(&_Vat.TransactOpts, u, v, rad)
}

// VatCageIterator is returned from FilterCage and is used to iterate over the raw logs and unpacked data for Cage events raised by the Vat contract.
type VatCageIterator struct {
	Event *VatCage // Event containing the contract specifics and raw log

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
func (it *VatCageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatCage)
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
		it.Event = new(VatCage)
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
func (it *VatCageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatCageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatCage represents a Cage event raised by the Vat contract.
type VatCage struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCage is a free log retrieval operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Vat *VatFilterer) FilterCage(opts *bind.FilterOpts) (*VatCageIterator, error) {

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return &VatCageIterator{contract: _Vat.contract, event: "Cage", logs: logs, sub: sub}, nil
}

// WatchCage is a free log subscription operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Vat *VatFilterer) WatchCage(opts *bind.WatchOpts, sink chan<- *VatCage) (event.Subscription, error) {

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatCage)
				if err := _Vat.contract.UnpackLog(event, "Cage", log); err != nil {
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

// ParseCage is a log parse operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_Vat *VatFilterer) ParseCage(log types.Log) (*VatCage, error) {
	event := new(VatCage)
	if err := _Vat.contract.UnpackLog(event, "Cage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Vat contract.
type VatDenyIterator struct {
	Event *VatDeny // Event containing the contract specifics and raw log

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
func (it *VatDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatDeny)
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
		it.Event = new(VatDeny)
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
func (it *VatDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatDeny represents a Deny event raised by the Vat contract.
type VatDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Vat *VatFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*VatDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &VatDenyIterator{contract: _Vat.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Vat *VatFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *VatDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatDeny)
				if err := _Vat.contract.UnpackLog(event, "Deny", log); err != nil {
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

// ParseDeny is a log parse operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Vat *VatFilterer) ParseDeny(log types.Log) (*VatDeny, error) {
	event := new(VatDeny)
	if err := _Vat.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatFessIterator is returned from FilterFess and is used to iterate over the raw logs and unpacked data for Fess events raised by the Vat contract.
type VatFessIterator struct {
	Event *VatFess // Event containing the contract specifics and raw log

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
func (it *VatFessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatFess)
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
		it.Event = new(VatFess)
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
func (it *VatFessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatFessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatFess represents a Fess event raised by the Vat contract.
type VatFess struct {
	Tab *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFess is a free log retrieval operation binding the contract event 0x7a3f1a1ebf14b193365bc7468b58eb3b80ae1638635424aae4eec386da2f02ba.
//
// Solidity: event Fess(uint256 indexed tab)
func (_Vat *VatFilterer) FilterFess(opts *bind.FilterOpts, tab []*big.Int) (*VatFessIterator, error) {

	var tabRule []interface{}
	for _, tabItem := range tab {
		tabRule = append(tabRule, tabItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Fess", tabRule)
	if err != nil {
		return nil, err
	}
	return &VatFessIterator{contract: _Vat.contract, event: "Fess", logs: logs, sub: sub}, nil
}

// WatchFess is a free log subscription operation binding the contract event 0x7a3f1a1ebf14b193365bc7468b58eb3b80ae1638635424aae4eec386da2f02ba.
//
// Solidity: event Fess(uint256 indexed tab)
func (_Vat *VatFilterer) WatchFess(opts *bind.WatchOpts, sink chan<- *VatFess, tab []*big.Int) (event.Subscription, error) {

	var tabRule []interface{}
	for _, tabItem := range tab {
		tabRule = append(tabRule, tabItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Fess", tabRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatFess)
				if err := _Vat.contract.UnpackLog(event, "Fess", log); err != nil {
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

// ParseFess is a log parse operation binding the contract event 0x7a3f1a1ebf14b193365bc7468b58eb3b80ae1638635424aae4eec386da2f02ba.
//
// Solidity: event Fess(uint256 indexed tab)
func (_Vat *VatFilterer) ParseFess(log types.Log) (*VatFess, error) {
	event := new(VatFess)
	if err := _Vat.contract.UnpackLog(event, "Fess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the Vat contract.
type VatFileIterator struct {
	Event *VatFile // Event containing the contract specifics and raw log

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
func (it *VatFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatFile)
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
		it.Event = new(VatFile)
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
func (it *VatFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatFile represents a File event raised by the Vat contract.
type VatFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Vat *VatFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*VatFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &VatFileIterator{contract: _Vat.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Vat *VatFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *VatFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatFile)
				if err := _Vat.contract.UnpackLog(event, "File", log); err != nil {
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

// ParseFile is a log parse operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Vat *VatFilterer) ParseFile(log types.Log) (*VatFile, error) {
	event := new(VatFile)
	if err := _Vat.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatFile0Iterator is returned from FilterFile0 and is used to iterate over the raw logs and unpacked data for File0 events raised by the Vat contract.
type VatFile0Iterator struct {
	Event *VatFile0 // Event containing the contract specifics and raw log

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
func (it *VatFile0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatFile0)
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
		it.Event = new(VatFile0)
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
func (it *VatFile0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatFile0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatFile0 represents a File0 event raised by the Vat contract.
type VatFile0 struct {
	Ilk  [32]byte
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile0 is a free log retrieval operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 indexed data)
func (_Vat *VatFilterer) FilterFile0(opts *bind.FilterOpts, ilk [][32]byte, what [][32]byte, data []*big.Int) (*VatFile0Iterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "File0", ilkRule, whatRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &VatFile0Iterator{contract: _Vat.contract, event: "File0", logs: logs, sub: sub}, nil
}

// WatchFile0 is a free log subscription operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 indexed data)
func (_Vat *VatFilterer) WatchFile0(opts *bind.WatchOpts, sink chan<- *VatFile0, ilk [][32]byte, what [][32]byte, data []*big.Int) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "File0", ilkRule, whatRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatFile0)
				if err := _Vat.contract.UnpackLog(event, "File0", log); err != nil {
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

// ParseFile0 is a log parse operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 indexed ilk, bytes32 indexed what, uint256 indexed data)
func (_Vat *VatFilterer) ParseFile0(log types.Log) (*VatFile0, error) {
	event := new(VatFile0)
	if err := _Vat.contract.UnpackLog(event, "File0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatFlapIterator is returned from FilterFlap and is used to iterate over the raw logs and unpacked data for Flap events raised by the Vat contract.
type VatFlapIterator struct {
	Event *VatFlap // Event containing the contract specifics and raw log

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
func (it *VatFlapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatFlap)
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
		it.Event = new(VatFlap)
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
func (it *VatFlapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatFlapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatFlap represents a Flap event raised by the Vat contract.
type VatFlap struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFlap is a free log retrieval operation binding the contract event 0x4c8a914f94eb278fed11934ac386da1b691d8c2e8a30291487140881038d13ce.
//
// Solidity: event Flap(uint256 indexed id)
func (_Vat *VatFilterer) FilterFlap(opts *bind.FilterOpts, id []*big.Int) (*VatFlapIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Flap", idRule)
	if err != nil {
		return nil, err
	}
	return &VatFlapIterator{contract: _Vat.contract, event: "Flap", logs: logs, sub: sub}, nil
}

// WatchFlap is a free log subscription operation binding the contract event 0x4c8a914f94eb278fed11934ac386da1b691d8c2e8a30291487140881038d13ce.
//
// Solidity: event Flap(uint256 indexed id)
func (_Vat *VatFilterer) WatchFlap(opts *bind.WatchOpts, sink chan<- *VatFlap, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Flap", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatFlap)
				if err := _Vat.contract.UnpackLog(event, "Flap", log); err != nil {
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

// ParseFlap is a log parse operation binding the contract event 0x4c8a914f94eb278fed11934ac386da1b691d8c2e8a30291487140881038d13ce.
//
// Solidity: event Flap(uint256 indexed id)
func (_Vat *VatFilterer) ParseFlap(log types.Log) (*VatFlap, error) {
	event := new(VatFlap)
	if err := _Vat.contract.UnpackLog(event, "Flap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatFlogIterator is returned from FilterFlog and is used to iterate over the raw logs and unpacked data for Flog events raised by the Vat contract.
type VatFlogIterator struct {
	Event *VatFlog // Event containing the contract specifics and raw log

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
func (it *VatFlogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatFlog)
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
		it.Event = new(VatFlog)
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
func (it *VatFlogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatFlogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatFlog represents a Flog event raised by the Vat contract.
type VatFlog struct {
	Era *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFlog is a free log retrieval operation binding the contract event 0x5aa14c9b66239d17e56d0724b7e90d8d82f28fcbdfb0d39e60614bd1d01dc753.
//
// Solidity: event Flog(uint256 indexed era)
func (_Vat *VatFilterer) FilterFlog(opts *bind.FilterOpts, era []*big.Int) (*VatFlogIterator, error) {

	var eraRule []interface{}
	for _, eraItem := range era {
		eraRule = append(eraRule, eraItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Flog", eraRule)
	if err != nil {
		return nil, err
	}
	return &VatFlogIterator{contract: _Vat.contract, event: "Flog", logs: logs, sub: sub}, nil
}

// WatchFlog is a free log subscription operation binding the contract event 0x5aa14c9b66239d17e56d0724b7e90d8d82f28fcbdfb0d39e60614bd1d01dc753.
//
// Solidity: event Flog(uint256 indexed era)
func (_Vat *VatFilterer) WatchFlog(opts *bind.WatchOpts, sink chan<- *VatFlog, era []*big.Int) (event.Subscription, error) {

	var eraRule []interface{}
	for _, eraItem := range era {
		eraRule = append(eraRule, eraItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Flog", eraRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatFlog)
				if err := _Vat.contract.UnpackLog(event, "Flog", log); err != nil {
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

// ParseFlog is a log parse operation binding the contract event 0x5aa14c9b66239d17e56d0724b7e90d8d82f28fcbdfb0d39e60614bd1d01dc753.
//
// Solidity: event Flog(uint256 indexed era)
func (_Vat *VatFilterer) ParseFlog(log types.Log) (*VatFlog, error) {
	event := new(VatFlog)
	if err := _Vat.contract.UnpackLog(event, "Flog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatFlopIterator is returned from FilterFlop and is used to iterate over the raw logs and unpacked data for Flop events raised by the Vat contract.
type VatFlopIterator struct {
	Event *VatFlop // Event containing the contract specifics and raw log

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
func (it *VatFlopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatFlop)
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
		it.Event = new(VatFlop)
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
func (it *VatFlopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatFlopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatFlop represents a Flop event raised by the Vat contract.
type VatFlop struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFlop is a free log retrieval operation binding the contract event 0xfa3d951cbf852d2a9cc2dfc9fc6b57914afbf57597ecf432c403ed2d74124b2c.
//
// Solidity: event Flop(uint256 indexed id)
func (_Vat *VatFilterer) FilterFlop(opts *bind.FilterOpts, id []*big.Int) (*VatFlopIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Flop", idRule)
	if err != nil {
		return nil, err
	}
	return &VatFlopIterator{contract: _Vat.contract, event: "Flop", logs: logs, sub: sub}, nil
}

// WatchFlop is a free log subscription operation binding the contract event 0xfa3d951cbf852d2a9cc2dfc9fc6b57914afbf57597ecf432c403ed2d74124b2c.
//
// Solidity: event Flop(uint256 indexed id)
func (_Vat *VatFilterer) WatchFlop(opts *bind.WatchOpts, sink chan<- *VatFlop, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Flop", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatFlop)
				if err := _Vat.contract.UnpackLog(event, "Flop", log); err != nil {
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

// ParseFlop is a log parse operation binding the contract event 0xfa3d951cbf852d2a9cc2dfc9fc6b57914afbf57597ecf432c403ed2d74124b2c.
//
// Solidity: event Flop(uint256 indexed id)
func (_Vat *VatFilterer) ParseFlop(log types.Log) (*VatFlop, error) {
	event := new(VatFlop)
	if err := _Vat.contract.UnpackLog(event, "Flop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatFluxIterator is returned from FilterFlux and is used to iterate over the raw logs and unpacked data for Flux events raised by the Vat contract.
type VatFluxIterator struct {
	Event *VatFlux // Event containing the contract specifics and raw log

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
func (it *VatFluxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatFlux)
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
		it.Event = new(VatFlux)
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
func (it *VatFluxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatFluxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatFlux represents a Flux event raised by the Vat contract.
type VatFlux struct {
	Ilk [32]byte
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFlux is a free log retrieval operation binding the contract event 0x5718eae79ffb8b6c98c497e5029a903705cf6a33a17aaab32de7fe198d8d8a0d.
//
// Solidity: event Flux(bytes32 indexed ilk, address indexed src, address indexed dst, uint256 wad)
func (_Vat *VatFilterer) FilterFlux(opts *bind.FilterOpts, ilk [][32]byte, src []common.Address, dst []common.Address) (*VatFluxIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Flux", ilkRule, srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &VatFluxIterator{contract: _Vat.contract, event: "Flux", logs: logs, sub: sub}, nil
}

// WatchFlux is a free log subscription operation binding the contract event 0x5718eae79ffb8b6c98c497e5029a903705cf6a33a17aaab32de7fe198d8d8a0d.
//
// Solidity: event Flux(bytes32 indexed ilk, address indexed src, address indexed dst, uint256 wad)
func (_Vat *VatFilterer) WatchFlux(opts *bind.WatchOpts, sink chan<- *VatFlux, ilk [][32]byte, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Flux", ilkRule, srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatFlux)
				if err := _Vat.contract.UnpackLog(event, "Flux", log); err != nil {
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

// ParseFlux is a log parse operation binding the contract event 0x5718eae79ffb8b6c98c497e5029a903705cf6a33a17aaab32de7fe198d8d8a0d.
//
// Solidity: event Flux(bytes32 indexed ilk, address indexed src, address indexed dst, uint256 wad)
func (_Vat *VatFilterer) ParseFlux(log types.Log) (*VatFlux, error) {
	event := new(VatFlux)
	if err := _Vat.contract.UnpackLog(event, "Flux", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatFoldIterator is returned from FilterFold and is used to iterate over the raw logs and unpacked data for Fold events raised by the Vat contract.
type VatFoldIterator struct {
	Event *VatFold // Event containing the contract specifics and raw log

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
func (it *VatFoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatFold)
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
		it.Event = new(VatFold)
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
func (it *VatFoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatFoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatFold represents a Fold event raised by the Vat contract.
type VatFold struct {
	I    [32]byte
	U    common.Address
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFold is a free log retrieval operation binding the contract event 0x8e03d1ac49b6d4e90dd1c4e641ecc5ca76b7c07a487690b6897c0e5e374b19d2.
//
// Solidity: event Fold(bytes32 indexed i, address indexed u, int256 indexed rate)
func (_Vat *VatFilterer) FilterFold(opts *bind.FilterOpts, i [][32]byte, u []common.Address, rate []*big.Int) (*VatFoldIterator, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}
	var uRule []interface{}
	for _, uItem := range u {
		uRule = append(uRule, uItem)
	}
	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Fold", iRule, uRule, rateRule)
	if err != nil {
		return nil, err
	}
	return &VatFoldIterator{contract: _Vat.contract, event: "Fold", logs: logs, sub: sub}, nil
}

// WatchFold is a free log subscription operation binding the contract event 0x8e03d1ac49b6d4e90dd1c4e641ecc5ca76b7c07a487690b6897c0e5e374b19d2.
//
// Solidity: event Fold(bytes32 indexed i, address indexed u, int256 indexed rate)
func (_Vat *VatFilterer) WatchFold(opts *bind.WatchOpts, sink chan<- *VatFold, i [][32]byte, u []common.Address, rate []*big.Int) (event.Subscription, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}
	var uRule []interface{}
	for _, uItem := range u {
		uRule = append(uRule, uItem)
	}
	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Fold", iRule, uRule, rateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatFold)
				if err := _Vat.contract.UnpackLog(event, "Fold", log); err != nil {
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

// ParseFold is a log parse operation binding the contract event 0x8e03d1ac49b6d4e90dd1c4e641ecc5ca76b7c07a487690b6897c0e5e374b19d2.
//
// Solidity: event Fold(bytes32 indexed i, address indexed u, int256 indexed rate)
func (_Vat *VatFilterer) ParseFold(log types.Log) (*VatFold, error) {
	event := new(VatFold)
	if err := _Vat.contract.UnpackLog(event, "Fold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatForkIterator is returned from FilterFork and is used to iterate over the raw logs and unpacked data for Fork events raised by the Vat contract.
type VatForkIterator struct {
	Event *VatFork // Event containing the contract specifics and raw log

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
func (it *VatForkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatFork)
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
		it.Event = new(VatFork)
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
func (it *VatForkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatForkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatFork represents a Fork event raised by the Vat contract.
type VatFork struct {
	Ilk  [32]byte
	Src  common.Address
	Dst  common.Address
	Dink *big.Int
	Dart *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFork is a free log retrieval operation binding the contract event 0x4b67161d2a4293b296b2f023c52ea4214353fa4f03e58973572faa00097dbd1e.
//
// Solidity: event Fork(bytes32 indexed ilk, address indexed src, address indexed dst, int256 dink, int256 dart)
func (_Vat *VatFilterer) FilterFork(opts *bind.FilterOpts, ilk [][32]byte, src []common.Address, dst []common.Address) (*VatForkIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Fork", ilkRule, srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &VatForkIterator{contract: _Vat.contract, event: "Fork", logs: logs, sub: sub}, nil
}

// WatchFork is a free log subscription operation binding the contract event 0x4b67161d2a4293b296b2f023c52ea4214353fa4f03e58973572faa00097dbd1e.
//
// Solidity: event Fork(bytes32 indexed ilk, address indexed src, address indexed dst, int256 dink, int256 dart)
func (_Vat *VatFilterer) WatchFork(opts *bind.WatchOpts, sink chan<- *VatFork, ilk [][32]byte, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Fork", ilkRule, srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatFork)
				if err := _Vat.contract.UnpackLog(event, "Fork", log); err != nil {
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

// ParseFork is a log parse operation binding the contract event 0x4b67161d2a4293b296b2f023c52ea4214353fa4f03e58973572faa00097dbd1e.
//
// Solidity: event Fork(bytes32 indexed ilk, address indexed src, address indexed dst, int256 dink, int256 dart)
func (_Vat *VatFilterer) ParseFork(log types.Log) (*VatFork, error) {
	event := new(VatFork)
	if err := _Vat.contract.UnpackLog(event, "Fork", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatFrobIterator is returned from FilterFrob and is used to iterate over the raw logs and unpacked data for Frob events raised by the Vat contract.
type VatFrobIterator struct {
	Event *VatFrob // Event containing the contract specifics and raw log

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
func (it *VatFrobIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatFrob)
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
		it.Event = new(VatFrob)
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
func (it *VatFrobIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatFrobIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatFrob represents a Frob event raised by the Vat contract.
type VatFrob struct {
	I    [32]byte
	U    common.Address
	V    common.Address
	W    common.Address
	Dink *big.Int
	Dart *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFrob is a free log retrieval operation binding the contract event 0xe37707842c8387f7c3c357f1d6c5bf57084e681573bdda024fae70cf0ecde80e.
//
// Solidity: event Frob(bytes32 i, address indexed u, address indexed v, address indexed w, int256 dink, int256 dart)
func (_Vat *VatFilterer) FilterFrob(opts *bind.FilterOpts, u []common.Address, v []common.Address, w []common.Address) (*VatFrobIterator, error) {

	var uRule []interface{}
	for _, uItem := range u {
		uRule = append(uRule, uItem)
	}
	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}
	var wRule []interface{}
	for _, wItem := range w {
		wRule = append(wRule, wItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Frob", uRule, vRule, wRule)
	if err != nil {
		return nil, err
	}
	return &VatFrobIterator{contract: _Vat.contract, event: "Frob", logs: logs, sub: sub}, nil
}

// WatchFrob is a free log subscription operation binding the contract event 0xe37707842c8387f7c3c357f1d6c5bf57084e681573bdda024fae70cf0ecde80e.
//
// Solidity: event Frob(bytes32 i, address indexed u, address indexed v, address indexed w, int256 dink, int256 dart)
func (_Vat *VatFilterer) WatchFrob(opts *bind.WatchOpts, sink chan<- *VatFrob, u []common.Address, v []common.Address, w []common.Address) (event.Subscription, error) {

	var uRule []interface{}
	for _, uItem := range u {
		uRule = append(uRule, uItem)
	}
	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}
	var wRule []interface{}
	for _, wItem := range w {
		wRule = append(wRule, wItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Frob", uRule, vRule, wRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatFrob)
				if err := _Vat.contract.UnpackLog(event, "Frob", log); err != nil {
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

// ParseFrob is a log parse operation binding the contract event 0xe37707842c8387f7c3c357f1d6c5bf57084e681573bdda024fae70cf0ecde80e.
//
// Solidity: event Frob(bytes32 i, address indexed u, address indexed v, address indexed w, int256 dink, int256 dart)
func (_Vat *VatFilterer) ParseFrob(log types.Log) (*VatFrob, error) {
	event := new(VatFrob)
	if err := _Vat.contract.UnpackLog(event, "Frob", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatGrabIterator is returned from FilterGrab and is used to iterate over the raw logs and unpacked data for Grab events raised by the Vat contract.
type VatGrabIterator struct {
	Event *VatGrab // Event containing the contract specifics and raw log

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
func (it *VatGrabIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatGrab)
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
		it.Event = new(VatGrab)
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
func (it *VatGrabIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatGrabIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatGrab represents a Grab event raised by the Vat contract.
type VatGrab struct {
	I    [32]byte
	U    common.Address
	V    common.Address
	W    common.Address
	Dink *big.Int
	Dart *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGrab is a free log retrieval operation binding the contract event 0x1b2837fd40844c96cf39e52acaae7902fb74257fe20b1b7df5458b97d896c636.
//
// Solidity: event Grab(bytes32 i, address indexed u, address indexed v, address indexed w, int256 dink, int256 dart)
func (_Vat *VatFilterer) FilterGrab(opts *bind.FilterOpts, u []common.Address, v []common.Address, w []common.Address) (*VatGrabIterator, error) {

	var uRule []interface{}
	for _, uItem := range u {
		uRule = append(uRule, uItem)
	}
	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}
	var wRule []interface{}
	for _, wItem := range w {
		wRule = append(wRule, wItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Grab", uRule, vRule, wRule)
	if err != nil {
		return nil, err
	}
	return &VatGrabIterator{contract: _Vat.contract, event: "Grab", logs: logs, sub: sub}, nil
}

// WatchGrab is a free log subscription operation binding the contract event 0x1b2837fd40844c96cf39e52acaae7902fb74257fe20b1b7df5458b97d896c636.
//
// Solidity: event Grab(bytes32 i, address indexed u, address indexed v, address indexed w, int256 dink, int256 dart)
func (_Vat *VatFilterer) WatchGrab(opts *bind.WatchOpts, sink chan<- *VatGrab, u []common.Address, v []common.Address, w []common.Address) (event.Subscription, error) {

	var uRule []interface{}
	for _, uItem := range u {
		uRule = append(uRule, uItem)
	}
	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}
	var wRule []interface{}
	for _, wItem := range w {
		wRule = append(wRule, wItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Grab", uRule, vRule, wRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatGrab)
				if err := _Vat.contract.UnpackLog(event, "Grab", log); err != nil {
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

// ParseGrab is a log parse operation binding the contract event 0x1b2837fd40844c96cf39e52acaae7902fb74257fe20b1b7df5458b97d896c636.
//
// Solidity: event Grab(bytes32 i, address indexed u, address indexed v, address indexed w, int256 dink, int256 dart)
func (_Vat *VatFilterer) ParseGrab(log types.Log) (*VatGrab, error) {
	event := new(VatGrab)
	if err := _Vat.contract.UnpackLog(event, "Grab", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatHealIterator is returned from FilterHeal and is used to iterate over the raw logs and unpacked data for Heal events raised by the Vat contract.
type VatHealIterator struct {
	Event *VatHeal // Event containing the contract specifics and raw log

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
func (it *VatHealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatHeal)
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
		it.Event = new(VatHeal)
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
func (it *VatHealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatHealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatHeal represents a Heal event raised by the Vat contract.
type VatHeal struct {
	Rad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterHeal is a free log retrieval operation binding the contract event 0x917d6982889419f491488c036c2e6abe788b07222064ab462158ec64ca2c4db7.
//
// Solidity: event Heal(uint256 indexed rad)
func (_Vat *VatFilterer) FilterHeal(opts *bind.FilterOpts, rad []*big.Int) (*VatHealIterator, error) {

	var radRule []interface{}
	for _, radItem := range rad {
		radRule = append(radRule, radItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Heal", radRule)
	if err != nil {
		return nil, err
	}
	return &VatHealIterator{contract: _Vat.contract, event: "Heal", logs: logs, sub: sub}, nil
}

// WatchHeal is a free log subscription operation binding the contract event 0x917d6982889419f491488c036c2e6abe788b07222064ab462158ec64ca2c4db7.
//
// Solidity: event Heal(uint256 indexed rad)
func (_Vat *VatFilterer) WatchHeal(opts *bind.WatchOpts, sink chan<- *VatHeal, rad []*big.Int) (event.Subscription, error) {

	var radRule []interface{}
	for _, radItem := range rad {
		radRule = append(radRule, radItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Heal", radRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatHeal)
				if err := _Vat.contract.UnpackLog(event, "Heal", log); err != nil {
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

// ParseHeal is a log parse operation binding the contract event 0x917d6982889419f491488c036c2e6abe788b07222064ab462158ec64ca2c4db7.
//
// Solidity: event Heal(uint256 indexed rad)
func (_Vat *VatFilterer) ParseHeal(log types.Log) (*VatHeal, error) {
	event := new(VatHeal)
	if err := _Vat.contract.UnpackLog(event, "Heal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatHopeIterator is returned from FilterHope and is used to iterate over the raw logs and unpacked data for Hope events raised by the Vat contract.
type VatHopeIterator struct {
	Event *VatHope // Event containing the contract specifics and raw log

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
func (it *VatHopeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatHope)
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
		it.Event = new(VatHope)
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
func (it *VatHopeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatHopeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatHope represents a Hope event raised by the Vat contract.
type VatHope struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterHope is a free log retrieval operation binding the contract event 0x3a21b662999d3fc0ceca48751a22bf61a806dcf3631e136271f02f7cb981fd43.
//
// Solidity: event Hope(address indexed usr)
func (_Vat *VatFilterer) FilterHope(opts *bind.FilterOpts, usr []common.Address) (*VatHopeIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Hope", usrRule)
	if err != nil {
		return nil, err
	}
	return &VatHopeIterator{contract: _Vat.contract, event: "Hope", logs: logs, sub: sub}, nil
}

// WatchHope is a free log subscription operation binding the contract event 0x3a21b662999d3fc0ceca48751a22bf61a806dcf3631e136271f02f7cb981fd43.
//
// Solidity: event Hope(address indexed usr)
func (_Vat *VatFilterer) WatchHope(opts *bind.WatchOpts, sink chan<- *VatHope, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Hope", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatHope)
				if err := _Vat.contract.UnpackLog(event, "Hope", log); err != nil {
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

// ParseHope is a log parse operation binding the contract event 0x3a21b662999d3fc0ceca48751a22bf61a806dcf3631e136271f02f7cb981fd43.
//
// Solidity: event Hope(address indexed usr)
func (_Vat *VatFilterer) ParseHope(log types.Log) (*VatHope, error) {
	event := new(VatHope)
	if err := _Vat.contract.UnpackLog(event, "Hope", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatInitIterator is returned from FilterInit and is used to iterate over the raw logs and unpacked data for Init events raised by the Vat contract.
type VatInitIterator struct {
	Event *VatInit // Event containing the contract specifics and raw log

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
func (it *VatInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatInit)
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
		it.Event = new(VatInit)
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
func (it *VatInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatInit represents a Init event raised by the Vat contract.
type VatInit struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInit is a free log retrieval operation binding the contract event 0xeeb45f27c5b399a603237b10d4803743d494bfc24c3a004cadb716c41033a555.
//
// Solidity: event Init(bytes32 indexed ilk)
func (_Vat *VatFilterer) FilterInit(opts *bind.FilterOpts, ilk [][32]byte) (*VatInitIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Init", ilkRule)
	if err != nil {
		return nil, err
	}
	return &VatInitIterator{contract: _Vat.contract, event: "Init", logs: logs, sub: sub}, nil
}

// WatchInit is a free log subscription operation binding the contract event 0xeeb45f27c5b399a603237b10d4803743d494bfc24c3a004cadb716c41033a555.
//
// Solidity: event Init(bytes32 indexed ilk)
func (_Vat *VatFilterer) WatchInit(opts *bind.WatchOpts, sink chan<- *VatInit, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Init", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatInit)
				if err := _Vat.contract.UnpackLog(event, "Init", log); err != nil {
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

// ParseInit is a log parse operation binding the contract event 0xeeb45f27c5b399a603237b10d4803743d494bfc24c3a004cadb716c41033a555.
//
// Solidity: event Init(bytes32 indexed ilk)
func (_Vat *VatFilterer) ParseInit(log types.Log) (*VatInit, error) {
	event := new(VatInit)
	if err := _Vat.contract.UnpackLog(event, "Init", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatKissIterator is returned from FilterKiss and is used to iterate over the raw logs and unpacked data for Kiss events raised by the Vat contract.
type VatKissIterator struct {
	Event *VatKiss // Event containing the contract specifics and raw log

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
func (it *VatKissIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatKiss)
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
		it.Event = new(VatKiss)
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
func (it *VatKissIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatKissIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatKiss represents a Kiss event raised by the Vat contract.
type VatKiss struct {
	Rad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKiss is a free log retrieval operation binding the contract event 0xdf1d0254f949dd4607095c8a45ed43a96d548776dbb1d6e8347513d07b109e9b.
//
// Solidity: event Kiss(uint256 indexed rad)
func (_Vat *VatFilterer) FilterKiss(opts *bind.FilterOpts, rad []*big.Int) (*VatKissIterator, error) {

	var radRule []interface{}
	for _, radItem := range rad {
		radRule = append(radRule, radItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Kiss", radRule)
	if err != nil {
		return nil, err
	}
	return &VatKissIterator{contract: _Vat.contract, event: "Kiss", logs: logs, sub: sub}, nil
}

// WatchKiss is a free log subscription operation binding the contract event 0xdf1d0254f949dd4607095c8a45ed43a96d548776dbb1d6e8347513d07b109e9b.
//
// Solidity: event Kiss(uint256 indexed rad)
func (_Vat *VatFilterer) WatchKiss(opts *bind.WatchOpts, sink chan<- *VatKiss, rad []*big.Int) (event.Subscription, error) {

	var radRule []interface{}
	for _, radItem := range rad {
		radRule = append(radRule, radItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Kiss", radRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatKiss)
				if err := _Vat.contract.UnpackLog(event, "Kiss", log); err != nil {
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

// ParseKiss is a log parse operation binding the contract event 0xdf1d0254f949dd4607095c8a45ed43a96d548776dbb1d6e8347513d07b109e9b.
//
// Solidity: event Kiss(uint256 indexed rad)
func (_Vat *VatFilterer) ParseKiss(log types.Log) (*VatKiss, error) {
	event := new(VatKiss)
	if err := _Vat.contract.UnpackLog(event, "Kiss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatMoveIterator is returned from FilterMove and is used to iterate over the raw logs and unpacked data for Move events raised by the Vat contract.
type VatMoveIterator struct {
	Event *VatMove // Event containing the contract specifics and raw log

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
func (it *VatMoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatMove)
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
		it.Event = new(VatMove)
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
func (it *VatMoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatMoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatMove represents a Move event raised by the Vat contract.
type VatMove struct {
	Src common.Address
	Dst common.Address
	Rad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMove is a free log retrieval operation binding the contract event 0xdeb3a6837278f6e9914a507e4d73f08e841d8fca434fb97d4307b3b0d3d6b105.
//
// Solidity: event Move(address indexed src, address indexed dst, uint256 rad)
func (_Vat *VatFilterer) FilterMove(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*VatMoveIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Move", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &VatMoveIterator{contract: _Vat.contract, event: "Move", logs: logs, sub: sub}, nil
}

// WatchMove is a free log subscription operation binding the contract event 0xdeb3a6837278f6e9914a507e4d73f08e841d8fca434fb97d4307b3b0d3d6b105.
//
// Solidity: event Move(address indexed src, address indexed dst, uint256 rad)
func (_Vat *VatFilterer) WatchMove(opts *bind.WatchOpts, sink chan<- *VatMove, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Move", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatMove)
				if err := _Vat.contract.UnpackLog(event, "Move", log); err != nil {
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

// ParseMove is a log parse operation binding the contract event 0xdeb3a6837278f6e9914a507e4d73f08e841d8fca434fb97d4307b3b0d3d6b105.
//
// Solidity: event Move(address indexed src, address indexed dst, uint256 rad)
func (_Vat *VatFilterer) ParseMove(log types.Log) (*VatMove, error) {
	event := new(VatMove)
	if err := _Vat.contract.UnpackLog(event, "Move", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatNopeIterator is returned from FilterNope and is used to iterate over the raw logs and unpacked data for Nope events raised by the Vat contract.
type VatNopeIterator struct {
	Event *VatNope // Event containing the contract specifics and raw log

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
func (it *VatNopeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatNope)
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
		it.Event = new(VatNope)
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
func (it *VatNopeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatNopeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatNope represents a Nope event raised by the Vat contract.
type VatNope struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNope is a free log retrieval operation binding the contract event 0x9cd85b2ca76a06c46be663a514e012af1aea8954b0e53f42146cd9b1ebb21ebc.
//
// Solidity: event Nope(address indexed usr)
func (_Vat *VatFilterer) FilterNope(opts *bind.FilterOpts, usr []common.Address) (*VatNopeIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Nope", usrRule)
	if err != nil {
		return nil, err
	}
	return &VatNopeIterator{contract: _Vat.contract, event: "Nope", logs: logs, sub: sub}, nil
}

// WatchNope is a free log subscription operation binding the contract event 0x9cd85b2ca76a06c46be663a514e012af1aea8954b0e53f42146cd9b1ebb21ebc.
//
// Solidity: event Nope(address indexed usr)
func (_Vat *VatFilterer) WatchNope(opts *bind.WatchOpts, sink chan<- *VatNope, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Nope", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatNope)
				if err := _Vat.contract.UnpackLog(event, "Nope", log); err != nil {
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

// ParseNope is a log parse operation binding the contract event 0x9cd85b2ca76a06c46be663a514e012af1aea8954b0e53f42146cd9b1ebb21ebc.
//
// Solidity: event Nope(address indexed usr)
func (_Vat *VatFilterer) ParseNope(log types.Log) (*VatNope, error) {
	event := new(VatNope)
	if err := _Vat.contract.UnpackLog(event, "Nope", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Vat contract.
type VatRelyIterator struct {
	Event *VatRely // Event containing the contract specifics and raw log

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
func (it *VatRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatRely)
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
		it.Event = new(VatRely)
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
func (it *VatRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatRely represents a Rely event raised by the Vat contract.
type VatRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Vat *VatFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*VatRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &VatRelyIterator{contract: _Vat.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Vat *VatFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *VatRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatRely)
				if err := _Vat.contract.UnpackLog(event, "Rely", log); err != nil {
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

// ParseRely is a log parse operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Vat *VatFilterer) ParseRely(log types.Log) (*VatRely, error) {
	event := new(VatRely)
	if err := _Vat.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatSlipIterator is returned from FilterSlip and is used to iterate over the raw logs and unpacked data for Slip events raised by the Vat contract.
type VatSlipIterator struct {
	Event *VatSlip // Event containing the contract specifics and raw log

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
func (it *VatSlipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatSlip)
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
		it.Event = new(VatSlip)
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
func (it *VatSlipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatSlipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatSlip represents a Slip event raised by the Vat contract.
type VatSlip struct {
	Ilk [32]byte
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSlip is a free log retrieval operation binding the contract event 0x0d5f62756a04d37a9bb68fd20b97c7c6a03e96ab87385a99f99c2463157dba4e.
//
// Solidity: event Slip(bytes32 indexed ilk, address indexed usr, int256 indexed wad)
func (_Vat *VatFilterer) FilterSlip(opts *bind.FilterOpts, ilk [][32]byte, usr []common.Address, wad []*big.Int) (*VatSlipIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var wadRule []interface{}
	for _, wadItem := range wad {
		wadRule = append(wadRule, wadItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Slip", ilkRule, usrRule, wadRule)
	if err != nil {
		return nil, err
	}
	return &VatSlipIterator{contract: _Vat.contract, event: "Slip", logs: logs, sub: sub}, nil
}

// WatchSlip is a free log subscription operation binding the contract event 0x0d5f62756a04d37a9bb68fd20b97c7c6a03e96ab87385a99f99c2463157dba4e.
//
// Solidity: event Slip(bytes32 indexed ilk, address indexed usr, int256 indexed wad)
func (_Vat *VatFilterer) WatchSlip(opts *bind.WatchOpts, sink chan<- *VatSlip, ilk [][32]byte, usr []common.Address, wad []*big.Int) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var wadRule []interface{}
	for _, wadItem := range wad {
		wadRule = append(wadRule, wadItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Slip", ilkRule, usrRule, wadRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatSlip)
				if err := _Vat.contract.UnpackLog(event, "Slip", log); err != nil {
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

// ParseSlip is a log parse operation binding the contract event 0x0d5f62756a04d37a9bb68fd20b97c7c6a03e96ab87385a99f99c2463157dba4e.
//
// Solidity: event Slip(bytes32 indexed ilk, address indexed usr, int256 indexed wad)
func (_Vat *VatFilterer) ParseSlip(log types.Log) (*VatSlip, error) {
	event := new(VatSlip)
	if err := _Vat.contract.UnpackLog(event, "Slip", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatSuckIterator is returned from FilterSuck and is used to iterate over the raw logs and unpacked data for Suck events raised by the Vat contract.
type VatSuckIterator struct {
	Event *VatSuck // Event containing the contract specifics and raw log

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
func (it *VatSuckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatSuck)
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
		it.Event = new(VatSuck)
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
func (it *VatSuckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatSuckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatSuck represents a Suck event raised by the Vat contract.
type VatSuck struct {
	U   common.Address
	V   common.Address
	Rad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSuck is a free log retrieval operation binding the contract event 0x02d16dda43fd89f02e33ce23ecf0251cdc426807cc72ae74d37e8d3681dae7e5.
//
// Solidity: event Suck(address indexed u, address indexed v, uint256 indexed rad)
func (_Vat *VatFilterer) FilterSuck(opts *bind.FilterOpts, u []common.Address, v []common.Address, rad []*big.Int) (*VatSuckIterator, error) {

	var uRule []interface{}
	for _, uItem := range u {
		uRule = append(uRule, uItem)
	}
	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}
	var radRule []interface{}
	for _, radItem := range rad {
		radRule = append(radRule, radItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Suck", uRule, vRule, radRule)
	if err != nil {
		return nil, err
	}
	return &VatSuckIterator{contract: _Vat.contract, event: "Suck", logs: logs, sub: sub}, nil
}

// WatchSuck is a free log subscription operation binding the contract event 0x02d16dda43fd89f02e33ce23ecf0251cdc426807cc72ae74d37e8d3681dae7e5.
//
// Solidity: event Suck(address indexed u, address indexed v, uint256 indexed rad)
func (_Vat *VatFilterer) WatchSuck(opts *bind.WatchOpts, sink chan<- *VatSuck, u []common.Address, v []common.Address, rad []*big.Int) (event.Subscription, error) {

	var uRule []interface{}
	for _, uItem := range u {
		uRule = append(uRule, uItem)
	}
	var vRule []interface{}
	for _, vItem := range v {
		vRule = append(vRule, vItem)
	}
	var radRule []interface{}
	for _, radItem := range rad {
		radRule = append(radRule, radItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Suck", uRule, vRule, radRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatSuck)
				if err := _Vat.contract.UnpackLog(event, "Suck", log); err != nil {
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

// ParseSuck is a log parse operation binding the contract event 0x02d16dda43fd89f02e33ce23ecf0251cdc426807cc72ae74d37e8d3681dae7e5.
//
// Solidity: event Suck(address indexed u, address indexed v, uint256 indexed rad)
func (_Vat *VatFilterer) ParseSuck(log types.Log) (*VatSuck, error) {
	event := new(VatSuck)
	if err := _Vat.contract.UnpackLog(event, "Suck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
