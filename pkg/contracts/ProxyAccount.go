// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ProxyAccountABI is the input ABI used to generate the binding from.
const ProxyAccountABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operationType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"store\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"DataChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ExecutedCall\",\"type\":\"event\"}]"

// ProxyAccountBin is the compiled bytecode used for deploying new contracts.
const ProxyAccountBin = `608060405234801561001057600080fd5b50604051610e78380380610e788339818101604052602081101561003357600080fd5b81019080805190602001909291905050506100538161008260201b60201c565b6000808060001b8152602001908152602001600020908051906020019061007b9291906100da565b505061017f565b606081604051602001808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014019150506040516020818303038152906040529050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011b57805160ff1916838001178555610149565b82800160010185558215610149579182015b8281111561014857825182559160200191906001019061012d565b5b509050610156919061015a565b5090565b61017c91905b80821115610178576000816000905550600101610160565b5090565b90565b610cea8061018e6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806344c028fe1461005c57806354f6127f14610109578063654cf88c146101b05780637f23690c14610257578063a6f9dae1146102da575b600080fd5b6101076004803603608081101561007257600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156100c357600080fd5b8201836020820111156100d557600080fd5b803590602001918460018302840111640100000000831117156100f757600080fd5b909192939192939050505061031e565b005b6101356004803603602081101561011f57600080fd5b81019080803590602001909291905050506105bc565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561017557808201518184015260208101905061015a565b50505050905090810190601f1680156101a25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101dc600480360360208110156101c657600080fd5b8101908080359060200190929190505050610670565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561021c578082015181840152602081019050610201565b50505050905090810190601f1680156102495780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102d86004803603604081101561026d57600080fd5b81019080803590602001909291908035906020019064010000000081111561029457600080fd5b8201836020820111156102a657600080fd5b803590602001918460018302840111640100000000831117156102c857600080fd5b9091929391929390505050610720565b005b61031c600480360360208110156102f057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610901565b005b6103d46000808060001b81526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103ca5780601f1061039f576101008083540402835291602001916103ca565b820191906000526020600020905b8154815290600101906020018083116103ad57829003601f168201915b5050505050610acf565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610474576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6f6e6c792d6f776e65722d616c6c6f776564000000000000000000000000000081525060200191505060405180910390fd5b60008514156104d2576104cc848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610ae0565b506105b5565b60018514156105af57600061052a83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610b26565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561056657600080fd5b8073ffffffffffffffffffffffffffffffffffffffff167fcf78cf0d6f3d8371e1075c69c492ab4ec5d8cf23a1a239b6a51a1d00be7ca31260405160405180910390a2506105b4565b600080fd5b5b5050505050565b60606000808381526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106645780601f1061063957610100808354040283529160200191610664565b820191906000526020600020905b81548152906001019060200180831161064757829003601f168201915b50505050509050919050565b60006020528060005260406000206000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107185780601f106106ed57610100808354040283529160200191610718565b820191906000526020600020905b8154815290600101906020018083116106fb57829003601f168201915b505050505081565b6107d66000808060001b81526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107cc5780601f106107a1576101008083540402835291602001916107cc565b820191906000526020600020905b8154815290600101906020018083116107af57829003601f168201915b5050505050610acf565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610876576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6f6e6c792d6f776e65722d616c6c6f776564000000000000000000000000000081525060200191505060405180910390fd5b81816000808681526020019081526020016000209190610897929190610b90565b50827fece574603820d07bc9b91f2a932baadf4628aabcb8afba49776529c14a6104b2838360405180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060405180910390a2505050565b6109b76000808060001b81526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109ad5780601f10610982576101008083540402835291602001916109ad565b820191906000526020600020905b81548152906001019060200180831161099057829003601f168201915b5050505050610acf565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a57576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f6f6e6c792d6f776e65722d616c6c6f776564000000000000000000000000000081525060200191505060405180910390fd5b610a6081610b38565b6000808060001b81526020019081526020016000209080519060200190610a88929190610c10565b508073ffffffffffffffffffffffffffffffffffffffff167fa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf3660405160405180910390a250565b600060148201519050809050919050565b60008082519050604051602084016000828483898b6187965a03f193508360008114610b135760018114610b1757610b1b565b8383fd5b8383f35b505050509392505050565b60008151602083016000f09050919050565b606081604051602001808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014019150506040516020818303038152906040529050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610bd157803560ff1916838001178555610bff565b82800160010185558215610bff579182015b82811115610bfe578235825591602001919060010190610be3565b5b509050610c0c9190610c90565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c5157805160ff1916838001178555610c7f565b82800160010185558215610c7f579182015b82811115610c7e578251825591602001919060010190610c63565b5b509050610c8c9190610c90565b5090565b610cb291905b80821115610cae576000816000905550600101610c96565b5090565b9056fea265627a7a72315820c803174814a645fc2da4bd3ba3c280dd415bccd1652d6ad073f80f56bfa68e2964736f6c634300050b0032`

// DeployProxyAccount deploys a new Ethereum contract, binding an instance of ProxyAccount to it.
func DeployProxyAccount(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *ProxyAccount, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyAccountABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProxyAccountBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProxyAccount{ProxyAccountCaller: ProxyAccountCaller{contract: contract}, ProxyAccountTransactor: ProxyAccountTransactor{contract: contract}, ProxyAccountFilterer: ProxyAccountFilterer{contract: contract}}, nil
}

// ProxyAccount is an auto generated Go binding around an Ethereum contract.
type ProxyAccount struct {
	ProxyAccountCaller     // Read-only binding to the contract
	ProxyAccountTransactor // Write-only binding to the contract
	ProxyAccountFilterer   // Log filterer for contract events
}

// ProxyAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyAccountSession struct {
	Contract     *ProxyAccount     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyAccountCallerSession struct {
	Contract *ProxyAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ProxyAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyAccountTransactorSession struct {
	Contract     *ProxyAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ProxyAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyAccountRaw struct {
	Contract *ProxyAccount // Generic contract binding to access the raw methods on
}

// ProxyAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyAccountCallerRaw struct {
	Contract *ProxyAccountCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyAccountTransactorRaw struct {
	Contract *ProxyAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyAccount creates a new instance of ProxyAccount, bound to a specific deployed contract.
func NewProxyAccount(address common.Address, backend bind.ContractBackend) (*ProxyAccount, error) {
	contract, err := bindProxyAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyAccount{ProxyAccountCaller: ProxyAccountCaller{contract: contract}, ProxyAccountTransactor: ProxyAccountTransactor{contract: contract}, ProxyAccountFilterer: ProxyAccountFilterer{contract: contract}}, nil
}

// NewProxyAccountCaller creates a new read-only instance of ProxyAccount, bound to a specific deployed contract.
func NewProxyAccountCaller(address common.Address, caller bind.ContractCaller) (*ProxyAccountCaller, error) {
	contract, err := bindProxyAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAccountCaller{contract: contract}, nil
}

// NewProxyAccountTransactor creates a new write-only instance of ProxyAccount, bound to a specific deployed contract.
func NewProxyAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyAccountTransactor, error) {
	contract, err := bindProxyAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAccountTransactor{contract: contract}, nil
}

// NewProxyAccountFilterer creates a new log filterer instance of ProxyAccount, bound to a specific deployed contract.
func NewProxyAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyAccountFilterer, error) {
	contract, err := bindProxyAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyAccountFilterer{contract: contract}, nil
}

// bindProxyAccount binds a generic wrapper to an already deployed contract.
func bindProxyAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyAccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAccount *ProxyAccountRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProxyAccount.Contract.ProxyAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAccount *ProxyAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAccount.Contract.ProxyAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAccount *ProxyAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAccount.Contract.ProxyAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAccount *ProxyAccountCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProxyAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAccount *ProxyAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAccount *ProxyAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAccount.Contract.contract.Transact(opts, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 _key) constant returns(bytes _value)
func (_ProxyAccount *ProxyAccountCaller) GetData(opts *bind.CallOpts, _key [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ProxyAccount.contract.Call(opts, out, "getData", _key)
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 _key) constant returns(bytes _value)
func (_ProxyAccount *ProxyAccountSession) GetData(_key [32]byte) ([]byte, error) {
	return _ProxyAccount.Contract.GetData(&_ProxyAccount.CallOpts, _key)
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 _key) constant returns(bytes _value)
func (_ProxyAccount *ProxyAccountCallerSession) GetData(_key [32]byte) ([]byte, error) {
	return _ProxyAccount.Contract.GetData(&_ProxyAccount.CallOpts, _key)
}

// Store is a free data retrieval call binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 ) constant returns(bytes)
func (_ProxyAccount *ProxyAccountCaller) Store(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ProxyAccount.contract.Call(opts, out, "store", arg0)
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 ) constant returns(bytes)
func (_ProxyAccount *ProxyAccountSession) Store(arg0 [32]byte) ([]byte, error) {
	return _ProxyAccount.Contract.Store(&_ProxyAccount.CallOpts, arg0)
}

// Store is a free data retrieval call binding the contract method 0x654cf88c.
//
// Solidity: function store(bytes32 ) constant returns(bytes)
func (_ProxyAccount *ProxyAccountCallerSession) Store(arg0 [32]byte) ([]byte, error) {
	return _ProxyAccount.Contract.Store(&_ProxyAccount.CallOpts, arg0)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_ProxyAccount *ProxyAccountTransactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _ProxyAccount.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_ProxyAccount *ProxyAccountSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _ProxyAccount.Contract.ChangeOwner(&_ProxyAccount.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_ProxyAccount *ProxyAccountTransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _ProxyAccount.Contract.ChangeOwner(&_ProxyAccount.TransactOpts, _owner)
}

// Execute is a paid mutator transaction binding the contract method 0x44c028fe.
//
// Solidity: function execute(uint256 _operationType, address _to, uint256 _value, bytes _data) returns()
func (_ProxyAccount *ProxyAccountTransactor) Execute(opts *bind.TransactOpts, _operationType *big.Int, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ProxyAccount.contract.Transact(opts, "execute", _operationType, _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x44c028fe.
//
// Solidity: function execute(uint256 _operationType, address _to, uint256 _value, bytes _data) returns()
func (_ProxyAccount *ProxyAccountSession) Execute(_operationType *big.Int, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ProxyAccount.Contract.Execute(&_ProxyAccount.TransactOpts, _operationType, _to, _value, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x44c028fe.
//
// Solidity: function execute(uint256 _operationType, address _to, uint256 _value, bytes _data) returns()
func (_ProxyAccount *ProxyAccountTransactorSession) Execute(_operationType *big.Int, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ProxyAccount.Contract.Execute(&_ProxyAccount.TransactOpts, _operationType, _to, _value, _data)
}

// SetData is a paid mutator transaction binding the contract method 0x7f23690c.
//
// Solidity: function setData(bytes32 _key, bytes _value) returns()
func (_ProxyAccount *ProxyAccountTransactor) SetData(opts *bind.TransactOpts, _key [32]byte, _value []byte) (*types.Transaction, error) {
	return _ProxyAccount.contract.Transact(opts, "setData", _key, _value)
}

// SetData is a paid mutator transaction binding the contract method 0x7f23690c.
//
// Solidity: function setData(bytes32 _key, bytes _value) returns()
func (_ProxyAccount *ProxyAccountSession) SetData(_key [32]byte, _value []byte) (*types.Transaction, error) {
	return _ProxyAccount.Contract.SetData(&_ProxyAccount.TransactOpts, _key, _value)
}

// SetData is a paid mutator transaction binding the contract method 0x7f23690c.
//
// Solidity: function setData(bytes32 _key, bytes _value) returns()
func (_ProxyAccount *ProxyAccountTransactorSession) SetData(_key [32]byte, _value []byte) (*types.Transaction, error) {
	return _ProxyAccount.Contract.SetData(&_ProxyAccount.TransactOpts, _key, _value)
}

// ProxyAccountContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the ProxyAccount contract.
type ProxyAccountContractCreatedIterator struct {
	Event *ProxyAccountContractCreated // Event containing the contract specifics and raw log

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
func (it *ProxyAccountContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAccountContractCreated)
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
		it.Event = new(ProxyAccountContractCreated)
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
func (it *ProxyAccountContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAccountContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAccountContractCreated represents a ContractCreated event raised by the ProxyAccount contract.
type ProxyAccountContractCreated struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0xcf78cf0d6f3d8371e1075c69c492ab4ec5d8cf23a1a239b6a51a1d00be7ca312.
//
// Solidity: event ContractCreated(address indexed contractAddress)
func (_ProxyAccount *ProxyAccountFilterer) FilterContractCreated(opts *bind.FilterOpts, contractAddress []common.Address) (*ProxyAccountContractCreatedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ProxyAccount.contract.FilterLogs(opts, "ContractCreated", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ProxyAccountContractCreatedIterator{contract: _ProxyAccount.contract, event: "ContractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0xcf78cf0d6f3d8371e1075c69c492ab4ec5d8cf23a1a239b6a51a1d00be7ca312.
//
// Solidity: event ContractCreated(address indexed contractAddress)
func (_ProxyAccount *ProxyAccountFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *ProxyAccountContractCreated, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ProxyAccount.contract.WatchLogs(opts, "ContractCreated", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAccountContractCreated)
				if err := _ProxyAccount.contract.UnpackLog(event, "ContractCreated", log); err != nil {
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

// ProxyAccountDataChangedIterator is returned from FilterDataChanged and is used to iterate over the raw logs and unpacked data for DataChanged events raised by the ProxyAccount contract.
type ProxyAccountDataChangedIterator struct {
	Event *ProxyAccountDataChanged // Event containing the contract specifics and raw log

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
func (it *ProxyAccountDataChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAccountDataChanged)
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
		it.Event = new(ProxyAccountDataChanged)
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
func (it *ProxyAccountDataChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAccountDataChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAccountDataChanged represents a DataChanged event raised by the ProxyAccount contract.
type ProxyAccountDataChanged struct {
	Key   [32]byte
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDataChanged is a free log retrieval operation binding the contract event 0xece574603820d07bc9b91f2a932baadf4628aabcb8afba49776529c14a6104b2.
//
// Solidity: event DataChanged(bytes32 indexed key, bytes value)
func (_ProxyAccount *ProxyAccountFilterer) FilterDataChanged(opts *bind.FilterOpts, key [][32]byte) (*ProxyAccountDataChangedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _ProxyAccount.contract.FilterLogs(opts, "DataChanged", keyRule)
	if err != nil {
		return nil, err
	}
	return &ProxyAccountDataChangedIterator{contract: _ProxyAccount.contract, event: "DataChanged", logs: logs, sub: sub}, nil
}

// WatchDataChanged is a free log subscription operation binding the contract event 0xece574603820d07bc9b91f2a932baadf4628aabcb8afba49776529c14a6104b2.
//
// Solidity: event DataChanged(bytes32 indexed key, bytes value)
func (_ProxyAccount *ProxyAccountFilterer) WatchDataChanged(opts *bind.WatchOpts, sink chan<- *ProxyAccountDataChanged, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _ProxyAccount.contract.WatchLogs(opts, "DataChanged", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAccountDataChanged)
				if err := _ProxyAccount.contract.UnpackLog(event, "DataChanged", log); err != nil {
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

// ProxyAccountExecutedCallIterator is returned from FilterExecutedCall and is used to iterate over the raw logs and unpacked data for ExecutedCall events raised by the ProxyAccount contract.
type ProxyAccountExecutedCallIterator struct {
	Event *ProxyAccountExecutedCall // Event containing the contract specifics and raw log

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
func (it *ProxyAccountExecutedCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAccountExecutedCall)
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
		it.Event = new(ProxyAccountExecutedCall)
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
func (it *ProxyAccountExecutedCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAccountExecutedCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAccountExecutedCall represents a ExecutedCall event raised by the ProxyAccount contract.
type ProxyAccountExecutedCall struct {
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExecutedCall is a free log retrieval operation binding the contract event 0x9c24516d129fb91aea31c4413f5db0a249820a0b2b83c12b3c6260f1733ca5d2.
//
// Solidity: event ExecutedCall(address indexed to, uint256 indexed value, bytes data)
func (_ProxyAccount *ProxyAccountFilterer) FilterExecutedCall(opts *bind.FilterOpts, to []common.Address, value []*big.Int) (*ProxyAccountExecutedCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _ProxyAccount.contract.FilterLogs(opts, "ExecutedCall", toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &ProxyAccountExecutedCallIterator{contract: _ProxyAccount.contract, event: "ExecutedCall", logs: logs, sub: sub}, nil
}

// WatchExecutedCall is a free log subscription operation binding the contract event 0x9c24516d129fb91aea31c4413f5db0a249820a0b2b83c12b3c6260f1733ca5d2.
//
// Solidity: event ExecutedCall(address indexed to, uint256 indexed value, bytes data)
func (_ProxyAccount *ProxyAccountFilterer) WatchExecutedCall(opts *bind.WatchOpts, sink chan<- *ProxyAccountExecutedCall, to []common.Address, value []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _ProxyAccount.contract.WatchLogs(opts, "ExecutedCall", toRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAccountExecutedCall)
				if err := _ProxyAccount.contract.UnpackLog(event, "ExecutedCall", log); err != nil {
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

// ProxyAccountOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the ProxyAccount contract.
type ProxyAccountOwnerChangedIterator struct {
	Event *ProxyAccountOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ProxyAccountOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAccountOwnerChanged)
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
		it.Event = new(ProxyAccountOwnerChanged)
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
func (it *ProxyAccountOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAccountOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAccountOwnerChanged represents a OwnerChanged event raised by the ProxyAccount contract.
type ProxyAccountOwnerChanged struct {
	OwnerAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address indexed ownerAddress)
func (_ProxyAccount *ProxyAccountFilterer) FilterOwnerChanged(opts *bind.FilterOpts, ownerAddress []common.Address) (*ProxyAccountOwnerChangedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _ProxyAccount.contract.FilterLogs(opts, "OwnerChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ProxyAccountOwnerChangedIterator{contract: _ProxyAccount.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address indexed ownerAddress)
func (_ProxyAccount *ProxyAccountFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *ProxyAccountOwnerChanged, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _ProxyAccount.contract.WatchLogs(opts, "OwnerChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAccountOwnerChanged)
				if err := _ProxyAccount.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
