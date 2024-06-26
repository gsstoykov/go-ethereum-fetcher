// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// SimplePersonInfoContractMetaData contains all meta data concerning the SimplePersonInfoContract contract.
var SimplePersonInfoContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"personIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAge\",\"type\":\"uint256\"}],\"name\":\"PersonInfoUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_personIndex\",\"type\":\"uint256\"}],\"name\":\"getPersonInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPersonsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"persons\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"age\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_age\",\"type\":\"uint256\"}],\"name\":\"setPersonInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50610b7e8061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806333f3b2a41461004e5780638f97cff01461006a578063a2f9eac614610088578063d336ac80146100b9575b5f80fd5b6100686004803603810190610063919061057f565b6100ea565b005b61007261021e565b60405161007f91906105e8565b60405180910390f35b6100a2600480360381019061009d9190610601565b610229565b6040516100b092919061068c565b60405180910390f35b6100d360048036038101906100ce9190610601565b6102dd565b6040516100e192919061068c565b60405180910390f35b5f82511161012d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161012490610704565b60405180910390fd5b5f811161016f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101669061076c565b60405180910390fd5b5f60405180604001604052808481526020018381525090505f81908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f0190816101c29190610984565b5060208201518160010155505060015f805490506101e09190610a80565b7f96fb71ab58332a1b713976cc33c58781380e987f4cf4f8b2ef62be13218fec32848460405161021192919061068c565b60405180910390a2505050565b5f8080549050905090565b5f8181548110610237575f80fd5b905f5260205f2090600202015f91509050805f018054610256906107b7565b80601f0160208091040260200160405190810160405280929190818152602001828054610282906107b7565b80156102cd5780601f106102a4576101008083540402835291602001916102cd565b820191905f5260205f20905b8154815290600101906020018083116102b057829003601f168201915b5050505050908060010154905082565b60605f80805490508310610326576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031d90610afd565b60405180910390fd5b5f80848154811061033a57610339610b1b565b5b905f5260205f2090600202016040518060400160405290815f82018054610360906107b7565b80601f016020809104026020016040519081016040528092919081815260200182805461038c906107b7565b80156103d75780601f106103ae576101008083540402835291602001916103d7565b820191905f5260205f20905b8154815290600101906020018083116103ba57829003601f168201915b505050505081526020016001820154815250509050805f015181602001519250925050915091565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61045e82610418565b810181811067ffffffffffffffff8211171561047d5761047c610428565b5b80604052505050565b5f61048f6103ff565b905061049b8282610455565b919050565b5f67ffffffffffffffff8211156104ba576104b9610428565b5b6104c382610418565b9050602081019050919050565b828183375f83830152505050565b5f6104f06104eb846104a0565b610486565b90508281526020810184848401111561050c5761050b610414565b5b6105178482856104d0565b509392505050565b5f82601f83011261053357610532610410565b5b81356105438482602086016104de565b91505092915050565b5f819050919050565b61055e8161054c565b8114610568575f80fd5b50565b5f8135905061057981610555565b92915050565b5f806040838503121561059557610594610408565b5b5f83013567ffffffffffffffff8111156105b2576105b161040c565b5b6105be8582860161051f565b92505060206105cf8582860161056b565b9150509250929050565b6105e28161054c565b82525050565b5f6020820190506105fb5f8301846105d9565b92915050565b5f6020828403121561061657610615610408565b5b5f6106238482850161056b565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61065e8261062c565b6106688185610636565b9350610678818560208601610646565b61068181610418565b840191505092915050565b5f6040820190508181035f8301526106a48185610654565b90506106b360208301846105d9565b9392505050565b7f4e616d652073686f756c64206e6f7420626520656d70747900000000000000005f82015250565b5f6106ee601883610636565b91506106f9826106ba565b602082019050919050565b5f6020820190508181035f83015261071b816106e2565b9050919050565b7f4167652073686f756c642062652067726561746572207468616e2030000000005f82015250565b5f610756601c83610636565b915061076182610722565b602082019050919050565b5f6020820190508181035f8301526107838161074a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806107ce57607f821691505b6020821081036107e1576107e061078a565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026108437fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610808565b61084d8683610808565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61088861088361087e8461054c565b610865565b61054c565b9050919050565b5f819050919050565b6108a18361086e565b6108b56108ad8261088f565b848454610814565b825550505050565b5f90565b6108c96108bd565b6108d4818484610898565b505050565b5b818110156108f7576108ec5f826108c1565b6001810190506108da565b5050565b601f82111561093c5761090d816107e7565b610916846107f9565b81016020851015610925578190505b610939610931856107f9565b8301826108d9565b50505b505050565b5f82821c905092915050565b5f61095c5f1984600802610941565b1980831691505092915050565b5f610974838361094d565b9150826002028217905092915050565b61098d8261062c565b67ffffffffffffffff8111156109a6576109a5610428565b5b6109b082546107b7565b6109bb8282856108fb565b5f60209050601f8311600181146109ec575f84156109da578287015190505b6109e48582610969565b865550610a4b565b601f1984166109fa866107e7565b5f5b82811015610a21578489015182556001820191506020850194506020810190506109fc565b86831015610a3e5784890151610a3a601f89168261094d565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610a8a8261054c565b9150610a958361054c565b9250828203905081811115610aad57610aac610a53565b5b92915050565b7f506572736f6e20696e646578206f7574206f6620626f756e64730000000000005f82015250565b5f610ae7601a83610636565b9150610af282610ab3565b602082019050919050565b5f6020820190508181035f830152610b1481610adb565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea2646970667358221220e3e028107f906bf294cef5d6ae53e56f7226538e287a618f9bed3cbf05f8654664736f6c63430008190033",
}

// SimplePersonInfoContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SimplePersonInfoContractMetaData.ABI instead.
var SimplePersonInfoContractABI = SimplePersonInfoContractMetaData.ABI

// SimplePersonInfoContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimplePersonInfoContractMetaData.Bin instead.
var SimplePersonInfoContractBin = SimplePersonInfoContractMetaData.Bin

// DeploySimplePersonInfoContract deploys a new Ethereum contract, binding an instance of SimplePersonInfoContract to it.
func DeploySimplePersonInfoContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimplePersonInfoContract, error) {
	parsed, err := SimplePersonInfoContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimplePersonInfoContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimplePersonInfoContract{SimplePersonInfoContractCaller: SimplePersonInfoContractCaller{contract: contract}, SimplePersonInfoContractTransactor: SimplePersonInfoContractTransactor{contract: contract}, SimplePersonInfoContractFilterer: SimplePersonInfoContractFilterer{contract: contract}}, nil
}

// SimplePersonInfoContract is an auto generated Go binding around an Ethereum contract.
type SimplePersonInfoContract struct {
	SimplePersonInfoContractCaller     // Read-only binding to the contract
	SimplePersonInfoContractTransactor // Write-only binding to the contract
	SimplePersonInfoContractFilterer   // Log filterer for contract events
}

// SimplePersonInfoContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimplePersonInfoContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplePersonInfoContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimplePersonInfoContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplePersonInfoContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimplePersonInfoContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplePersonInfoContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimplePersonInfoContractSession struct {
	Contract     *SimplePersonInfoContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SimplePersonInfoContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimplePersonInfoContractCallerSession struct {
	Contract *SimplePersonInfoContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// SimplePersonInfoContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimplePersonInfoContractTransactorSession struct {
	Contract     *SimplePersonInfoContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// SimplePersonInfoContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimplePersonInfoContractRaw struct {
	Contract *SimplePersonInfoContract // Generic contract binding to access the raw methods on
}

// SimplePersonInfoContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimplePersonInfoContractCallerRaw struct {
	Contract *SimplePersonInfoContractCaller // Generic read-only contract binding to access the raw methods on
}

// SimplePersonInfoContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimplePersonInfoContractTransactorRaw struct {
	Contract *SimplePersonInfoContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplePersonInfoContract creates a new instance of SimplePersonInfoContract, bound to a specific deployed contract.
func NewSimplePersonInfoContract(address common.Address, backend bind.ContractBackend) (*SimplePersonInfoContract, error) {
	contract, err := bindSimplePersonInfoContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimplePersonInfoContract{SimplePersonInfoContractCaller: SimplePersonInfoContractCaller{contract: contract}, SimplePersonInfoContractTransactor: SimplePersonInfoContractTransactor{contract: contract}, SimplePersonInfoContractFilterer: SimplePersonInfoContractFilterer{contract: contract}}, nil
}

// NewSimplePersonInfoContractCaller creates a new read-only instance of SimplePersonInfoContract, bound to a specific deployed contract.
func NewSimplePersonInfoContractCaller(address common.Address, caller bind.ContractCaller) (*SimplePersonInfoContractCaller, error) {
	contract, err := bindSimplePersonInfoContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimplePersonInfoContractCaller{contract: contract}, nil
}

// NewSimplePersonInfoContractTransactor creates a new write-only instance of SimplePersonInfoContract, bound to a specific deployed contract.
func NewSimplePersonInfoContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SimplePersonInfoContractTransactor, error) {
	contract, err := bindSimplePersonInfoContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimplePersonInfoContractTransactor{contract: contract}, nil
}

// NewSimplePersonInfoContractFilterer creates a new log filterer instance of SimplePersonInfoContract, bound to a specific deployed contract.
func NewSimplePersonInfoContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SimplePersonInfoContractFilterer, error) {
	contract, err := bindSimplePersonInfoContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimplePersonInfoContractFilterer{contract: contract}, nil
}

// bindSimplePersonInfoContract binds a generic wrapper to an already deployed contract.
func bindSimplePersonInfoContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimplePersonInfoContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimplePersonInfoContract *SimplePersonInfoContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimplePersonInfoContract.Contract.SimplePersonInfoContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimplePersonInfoContract *SimplePersonInfoContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimplePersonInfoContract.Contract.SimplePersonInfoContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimplePersonInfoContract *SimplePersonInfoContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimplePersonInfoContract.Contract.SimplePersonInfoContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimplePersonInfoContract *SimplePersonInfoContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimplePersonInfoContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimplePersonInfoContract *SimplePersonInfoContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimplePersonInfoContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimplePersonInfoContract *SimplePersonInfoContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimplePersonInfoContract.Contract.contract.Transact(opts, method, params...)
}

// GetPersonInfo is a free data retrieval call binding the contract method 0xd336ac80.
//
// Solidity: function getPersonInfo(uint256 _personIndex) view returns(string, uint256)
func (_SimplePersonInfoContract *SimplePersonInfoContractCaller) GetPersonInfo(opts *bind.CallOpts, _personIndex *big.Int) (string, *big.Int, error) {
	var out []interface{}
	err := _SimplePersonInfoContract.contract.Call(opts, &out, "getPersonInfo", _personIndex)

	if err != nil {
		return *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPersonInfo is a free data retrieval call binding the contract method 0xd336ac80.
//
// Solidity: function getPersonInfo(uint256 _personIndex) view returns(string, uint256)
func (_SimplePersonInfoContract *SimplePersonInfoContractSession) GetPersonInfo(_personIndex *big.Int) (string, *big.Int, error) {
	return _SimplePersonInfoContract.Contract.GetPersonInfo(&_SimplePersonInfoContract.CallOpts, _personIndex)
}

// GetPersonInfo is a free data retrieval call binding the contract method 0xd336ac80.
//
// Solidity: function getPersonInfo(uint256 _personIndex) view returns(string, uint256)
func (_SimplePersonInfoContract *SimplePersonInfoContractCallerSession) GetPersonInfo(_personIndex *big.Int) (string, *big.Int, error) {
	return _SimplePersonInfoContract.Contract.GetPersonInfo(&_SimplePersonInfoContract.CallOpts, _personIndex)
}

// GetPersonsCount is a free data retrieval call binding the contract method 0x8f97cff0.
//
// Solidity: function getPersonsCount() view returns(uint256)
func (_SimplePersonInfoContract *SimplePersonInfoContractCaller) GetPersonsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimplePersonInfoContract.contract.Call(opts, &out, "getPersonsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPersonsCount is a free data retrieval call binding the contract method 0x8f97cff0.
//
// Solidity: function getPersonsCount() view returns(uint256)
func (_SimplePersonInfoContract *SimplePersonInfoContractSession) GetPersonsCount() (*big.Int, error) {
	return _SimplePersonInfoContract.Contract.GetPersonsCount(&_SimplePersonInfoContract.CallOpts)
}

// GetPersonsCount is a free data retrieval call binding the contract method 0x8f97cff0.
//
// Solidity: function getPersonsCount() view returns(uint256)
func (_SimplePersonInfoContract *SimplePersonInfoContractCallerSession) GetPersonsCount() (*big.Int, error) {
	return _SimplePersonInfoContract.Contract.GetPersonsCount(&_SimplePersonInfoContract.CallOpts)
}

// Persons is a free data retrieval call binding the contract method 0xa2f9eac6.
//
// Solidity: function persons(uint256 ) view returns(string name, uint256 age)
func (_SimplePersonInfoContract *SimplePersonInfoContractCaller) Persons(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name string
	Age  *big.Int
}, error) {
	var out []interface{}
	err := _SimplePersonInfoContract.contract.Call(opts, &out, "persons", arg0)

	outstruct := new(struct {
		Name string
		Age  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Age = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Persons is a free data retrieval call binding the contract method 0xa2f9eac6.
//
// Solidity: function persons(uint256 ) view returns(string name, uint256 age)
func (_SimplePersonInfoContract *SimplePersonInfoContractSession) Persons(arg0 *big.Int) (struct {
	Name string
	Age  *big.Int
}, error) {
	return _SimplePersonInfoContract.Contract.Persons(&_SimplePersonInfoContract.CallOpts, arg0)
}

// Persons is a free data retrieval call binding the contract method 0xa2f9eac6.
//
// Solidity: function persons(uint256 ) view returns(string name, uint256 age)
func (_SimplePersonInfoContract *SimplePersonInfoContractCallerSession) Persons(arg0 *big.Int) (struct {
	Name string
	Age  *big.Int
}, error) {
	return _SimplePersonInfoContract.Contract.Persons(&_SimplePersonInfoContract.CallOpts, arg0)
}

// SetPersonInfo is a paid mutator transaction binding the contract method 0x33f3b2a4.
//
// Solidity: function setPersonInfo(string _name, uint256 _age) returns()
func (_SimplePersonInfoContract *SimplePersonInfoContractTransactor) SetPersonInfo(opts *bind.TransactOpts, _name string, _age *big.Int) (*types.Transaction, error) {
	return _SimplePersonInfoContract.contract.Transact(opts, "setPersonInfo", _name, _age)
}

// SetPersonInfo is a paid mutator transaction binding the contract method 0x33f3b2a4.
//
// Solidity: function setPersonInfo(string _name, uint256 _age) returns()
func (_SimplePersonInfoContract *SimplePersonInfoContractSession) SetPersonInfo(_name string, _age *big.Int) (*types.Transaction, error) {
	return _SimplePersonInfoContract.Contract.SetPersonInfo(&_SimplePersonInfoContract.TransactOpts, _name, _age)
}

// SetPersonInfo is a paid mutator transaction binding the contract method 0x33f3b2a4.
//
// Solidity: function setPersonInfo(string _name, uint256 _age) returns()
func (_SimplePersonInfoContract *SimplePersonInfoContractTransactorSession) SetPersonInfo(_name string, _age *big.Int) (*types.Transaction, error) {
	return _SimplePersonInfoContract.Contract.SetPersonInfo(&_SimplePersonInfoContract.TransactOpts, _name, _age)
}

// SimplePersonInfoContractPersonInfoUpdatedIterator is returned from FilterPersonInfoUpdated and is used to iterate over the raw logs and unpacked data for PersonInfoUpdated events raised by the SimplePersonInfoContract contract.
type SimplePersonInfoContractPersonInfoUpdatedIterator struct {
	Event *SimplePersonInfoContractPersonInfoUpdated // Event containing the contract specifics and raw log

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
func (it *SimplePersonInfoContractPersonInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplePersonInfoContractPersonInfoUpdated)
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
		it.Event = new(SimplePersonInfoContractPersonInfoUpdated)
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
func (it *SimplePersonInfoContractPersonInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplePersonInfoContractPersonInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplePersonInfoContractPersonInfoUpdated represents a PersonInfoUpdated event raised by the SimplePersonInfoContract contract.
type SimplePersonInfoContractPersonInfoUpdated struct {
	PersonIndex *big.Int
	NewName     string
	NewAge      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPersonInfoUpdated is a free log retrieval operation binding the contract event 0x96fb71ab58332a1b713976cc33c58781380e987f4cf4f8b2ef62be13218fec32.
//
// Solidity: event PersonInfoUpdated(uint256 indexed personIndex, string newName, uint256 newAge)
func (_SimplePersonInfoContract *SimplePersonInfoContractFilterer) FilterPersonInfoUpdated(opts *bind.FilterOpts, personIndex []*big.Int) (*SimplePersonInfoContractPersonInfoUpdatedIterator, error) {

	var personIndexRule []interface{}
	for _, personIndexItem := range personIndex {
		personIndexRule = append(personIndexRule, personIndexItem)
	}

	logs, sub, err := _SimplePersonInfoContract.contract.FilterLogs(opts, "PersonInfoUpdated", personIndexRule)
	if err != nil {
		return nil, err
	}
	return &SimplePersonInfoContractPersonInfoUpdatedIterator{contract: _SimplePersonInfoContract.contract, event: "PersonInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPersonInfoUpdated is a free log subscription operation binding the contract event 0x96fb71ab58332a1b713976cc33c58781380e987f4cf4f8b2ef62be13218fec32.
//
// Solidity: event PersonInfoUpdated(uint256 indexed personIndex, string newName, uint256 newAge)
func (_SimplePersonInfoContract *SimplePersonInfoContractFilterer) WatchPersonInfoUpdated(opts *bind.WatchOpts, sink chan<- *SimplePersonInfoContractPersonInfoUpdated, personIndex []*big.Int) (event.Subscription, error) {

	var personIndexRule []interface{}
	for _, personIndexItem := range personIndex {
		personIndexRule = append(personIndexRule, personIndexItem)
	}

	logs, sub, err := _SimplePersonInfoContract.contract.WatchLogs(opts, "PersonInfoUpdated", personIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplePersonInfoContractPersonInfoUpdated)
				if err := _SimplePersonInfoContract.contract.UnpackLog(event, "PersonInfoUpdated", log); err != nil {
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

// ParsePersonInfoUpdated is a log parse operation binding the contract event 0x96fb71ab58332a1b713976cc33c58781380e987f4cf4f8b2ef62be13218fec32.
//
// Solidity: event PersonInfoUpdated(uint256 indexed personIndex, string newName, uint256 newAge)
func (_SimplePersonInfoContract *SimplePersonInfoContractFilterer) ParsePersonInfoUpdated(log types.Log) (*SimplePersonInfoContractPersonInfoUpdated, error) {
	event := new(SimplePersonInfoContractPersonInfoUpdated)
	if err := _SimplePersonInfoContract.contract.UnpackLog(event, "PersonInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
