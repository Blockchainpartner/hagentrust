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

// ClaimRegistryABI is the input ABI used to generate the binding from.
const ClaimRegistryABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_topic\",\"type\":\"bytes32\"}],\"name\":\"hasIssuedTopic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"issued\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"}],\"name\":\"hasIssuedForRecipient\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"issued\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_topic\",\"type\":\"bytes32\"}],\"name\":\"getIssuerForRecipientTopic\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"issuers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"getIssuerForRecipient\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"issuers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"}],\"name\":\"getIssuerTopicForRecipient\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_topic\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"setClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"isRecipient\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"alreadyRecipent\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_topic\",\"type\":\"bytes32\"}],\"name\":\"removeClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIssuers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"issuers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_topic\",\"type\":\"bytes32\"}],\"name\":\"getRecipientForIssuerTopic\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecipients\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_topic\",\"type\":\"bytes32\"}],\"name\":\"getClaim\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTopics\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"topics\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"ClaimSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"}],\"name\":\"ClaimRemoved\",\"type\":\"event\"}]"

// ClaimRegistryBin is the compiled bytecode used for deploying new contracts.
const ClaimRegistryBin = `608060405234801561001057600080fd5b50611ca2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063a1c8ca7e1161008c578063d78d610b11610066578063d78d610b14610635578063e1661eff14610694578063e6d970aa1461077b578063fee0117714610862576100ea565b8063a1c8ca7e146104e5578063a594da9514610533578063a7dd744c14610592576100ea565b8063343766d0116100c8578063343766d014610294578063411420d21461032d57806377640f9c146103e65780638c5143ea14610489576100ea565b806303915fac146100ef5780630c131a36146101755780633378b1c1146101f1575b600080fd5b61015b6004803603606081101561010557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506108c1565b604051808215151515815260200191505060405180910390f35b6101d76004803603604081101561018b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061096f565b604051808215151515815260200191505060405180910390f35b61023d6004803603604081101561020757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610986565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610280578082015181840152602081019050610265565b505050509050019250505060405180910390f35b6102d6600480360360208110156102aa57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610af3565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156103195780820151818401526020810190506102fe565b505050509050019250505060405180910390f35b61038f6004803603604081101561034357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d1a565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156103d25780820151818401526020810190506103b7565b505050509050019250505060405180910390f35b610487600480360360608110156103fc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561044357600080fd5b82018360208201111561045557600080fd5b8035906020019184600183028401116401000000008311171561047757600080fd5b9091929391929390505050610f1b565b005b6104cb6004803603602081101561049f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061122b565b604051808215151515815260200191505060405180910390f35b610531600480360360408110156104fb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506112c7565b005b61053b6113be565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561057e578082015181840152602081019050610563565b505050509050019250505060405180910390f35b6105de600480360360408110156105a857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611571565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610621578082015181840152602081019050610606565b505050509050019250505060405180910390f35b61063d611720565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610680578082015181840152602081019050610665565b505050509050019250505060405180910390f35b610700600480360360608110156106aa57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506117ae565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610740578082015181840152602081019050610725565b50505050905090810190601f16801561076d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6107e76004803603606081101561079157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506118de565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561082757808201518184015260208101905061080c565b50505050905090810190601f1680156108545780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61086a6119a8565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156108ad578082015181840152602081019050610892565b505050509050019250505060405180910390f35b6000806000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000208054600181600116156101000203166002900490501190509392505050565b60008061097c8484610d1a565b5111905092915050565b60608061099284610af3565b9050606081516040519080825280602002602001820160405280156109c65781602001602082028038833980820191505090505b509050600080905060008090505b8351811015610a27576109fb878583815181106109ed57fe5b6020026020010151886108c1565b15610a1a5780838381518110610a0d57fe5b6020026020010181815250505b80806001019150506109d4565b60018201915081604051908082528060200260200182016040528015610a5c5781602001602082028038833980820191505090505b509450600090505b81811015610ae65783838281518110610a7957fe5b602002602001015181518110610a8b57fe5b6020026020010151858281518110610a9f57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080600101915050610a64565b8494505050505092915050565b606080600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610bb557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610b6b575b5050505050905060608151604051908082528060200260200182016040528015610bee5781602001602082028038833980820191505090505b509050600080905060008090505b8351811015610c52576000610c2487868481518110610c1757fe5b6020026020010151610d1a565b511115610c455780838381518110610c3857fe5b6020026020010181815250505b8080600101915050610bfc565b60018201915081604051908082528060200260200182016040528015610c875781602001602082028038833980820191505090505b509450600090505b81811015610d115783838281518110610ca457fe5b602002602001015181518110610cb657fe5b6020026020010151858281518110610cca57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080600101915050610c8f565b50505050919050565b606080600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610de357602002820191906000526020600020905b815481526020019060010190808311610dcf575b5050505050905060608151604051908082528060200260200182016040528015610e1c5781602001602082028038833980820191505090505b509050600080905060008090505b8351811015610e8357610e518787868481518110610e4457fe5b60200260200101516108c1565b15610e765780838381518110610e6357fe5b6020026020010181815250506001820191505b8080600101915050610e2a565b81604051908082528060200260200182016040528015610eb25781602001602082028038833980820191505090505b509450600090505b81811015610f0e5783838281518110610ecf57fe5b602002602001015181518110610ee157fe5b6020026020010151858281518110610ef557fe5b6020026020010181815250508080600101915050610eba565b8494505050505092915050565b81816000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008681526020019081526020016000209190610fb6929190611b80565b50610fc08461122b565b61102b5760018490806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b611035843361096f565b6110dd57600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000203390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b6110e88433856108c1565b61119357600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208390806001815401808255809150509060018203906000526020600020016000909192909190915055505b828473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f7c49dba8ead478c5260eb86834965ce378cfaecfd9c899d0078ed9276b6e1ffa858560405180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060405180910390a450505050565b600080600090505b6001805490508110156112c1576001818154811061124d57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156112b457600191506112c1565b8080600101915050611233565b50919050565b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828152602001908152602001600020600061135f9190611c00565b808273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f895b934151a219f5c7dd6868b1397f44e9435ef5969693eed4bdc1c74dff2df160405160405180910390a45050565b606060008090505b60018054905081101561156a576060611415600183815481106113e557fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610af3565b905060008151111561155c57606081518451016040519080825280602002602001820160405280156114565781602001602082028038833980820191505090505b509050600080905060008090505b85518110156114da5785818151811061147957fe5b602002602001015183838151811061148d57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506001820191508080600101915050611464565b600090505b8351811015611555578381815181106114f457fe5b602002602001015183838151811061150857fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060018201915080806001019150506114df565b8295505050505b5080806001019150506113c6565b8191505090565b6060806001805490506040519080825280602002602001820160405280156115a85781602001602082028038833980820191505090505b509050600080905060008090505b60018054905081101561163157611605600182815481106115d357fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1687876108c1565b15611624578083838151811061161757fe5b6020026020010181815250505b80806001019150506115b6565b600182019150816040519080825280602002602001820160405280156116665781602001602082028038833980820191505090505b509350600090505b8181101561171457600183828151811061168457fe5b60200260200101518154811061169657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168482815181106116cd57fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808060010191505061166e565b83935050505092915050565b606060018054806020026020016040519081016040528092919081815260200182805480156117a457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161175a575b5050505050905090565b60606000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156118d05780601f106118a5576101008083540402835291602001916118d0565b820191906000526020600020905b8154815290600101906020018083116118b357829003601f168201915b505050505090509392505050565b600060205282600052604060002060205281600052604060002060205280600052604060002060009250925050508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156119a05780601f10611975576101008083540402835291602001916119a0565b820191906000526020600020905b81548152906001019060200180831161198357829003601f168201915b505050505081565b606060008090505b600180549050811015611b795760606119ff600183815481106119cf57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610af3565b9050600081511115611b6b5760008090505b8151811015611b69576060611a7060018581548110611a2c57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848481518110611a6357fe5b6020026020010151610d1a565b9050600081511115611b5b5760608151865101604051908082528060200260200182016040528015611ab15781602001602082028038833980820191505090505b509050600080905060008090505b8751811015611b0757878181518110611ad457fe5b6020026020010151838381518110611ae857fe5b6020026020010181815250506001820191508080600101915050611abf565b600090505b8551811015611b5457838181518110611b2157fe5b6020026020010151838381518110611b3557fe5b6020026020010181815250506001820191508080600101915050611b0c565b8297505050505b508080600101915050611a11565b505b5080806001019150506119b0565b8191505090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611bc157803560ff1916838001178555611bef565b82800160010185558215611bef579182015b82811115611bee578235825591602001919060010190611bd3565b5b509050611bfc9190611c48565b5090565b50805460018160011615610100020316600290046000825580601f10611c265750611c45565b601f016020900490600052602060002090810190611c449190611c48565b5b50565b611c6a91905b80821115611c66576000816000905550600101611c4e565b5090565b9056fea265627a7a7231582026ffe5d4404f66b8b7615bc823407c82e6cc204b3431606266fc1a1e3edc013f64736f6c634300050b0032`

// DeployClaimRegistry deploys a new Ethereum contract, binding an instance of ClaimRegistry to it.
func DeployClaimRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ClaimRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClaimRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClaimRegistry{ClaimRegistryCaller: ClaimRegistryCaller{contract: contract}, ClaimRegistryTransactor: ClaimRegistryTransactor{contract: contract}, ClaimRegistryFilterer: ClaimRegistryFilterer{contract: contract}}, nil
}

// ClaimRegistry is an auto generated Go binding around an Ethereum contract.
type ClaimRegistry struct {
	ClaimRegistryCaller     // Read-only binding to the contract
	ClaimRegistryTransactor // Write-only binding to the contract
	ClaimRegistryFilterer   // Log filterer for contract events
}

// ClaimRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimRegistrySession struct {
	Contract     *ClaimRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimRegistryCallerSession struct {
	Contract *ClaimRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ClaimRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimRegistryTransactorSession struct {
	Contract     *ClaimRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ClaimRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimRegistryRaw struct {
	Contract *ClaimRegistry // Generic contract binding to access the raw methods on
}

// ClaimRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimRegistryCallerRaw struct {
	Contract *ClaimRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimRegistryTransactorRaw struct {
	Contract *ClaimRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimRegistry creates a new instance of ClaimRegistry, bound to a specific deployed contract.
func NewClaimRegistry(address common.Address, backend bind.ContractBackend) (*ClaimRegistry, error) {
	contract, err := bindClaimRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimRegistry{ClaimRegistryCaller: ClaimRegistryCaller{contract: contract}, ClaimRegistryTransactor: ClaimRegistryTransactor{contract: contract}, ClaimRegistryFilterer: ClaimRegistryFilterer{contract: contract}}, nil
}

// NewClaimRegistryCaller creates a new read-only instance of ClaimRegistry, bound to a specific deployed contract.
func NewClaimRegistryCaller(address common.Address, caller bind.ContractCaller) (*ClaimRegistryCaller, error) {
	contract, err := bindClaimRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimRegistryCaller{contract: contract}, nil
}

// NewClaimRegistryTransactor creates a new write-only instance of ClaimRegistry, bound to a specific deployed contract.
func NewClaimRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimRegistryTransactor, error) {
	contract, err := bindClaimRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimRegistryTransactor{contract: contract}, nil
}

// NewClaimRegistryFilterer creates a new log filterer instance of ClaimRegistry, bound to a specific deployed contract.
func NewClaimRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimRegistryFilterer, error) {
	contract, err := bindClaimRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimRegistryFilterer{contract: contract}, nil
}

// bindClaimRegistry binds a generic wrapper to an already deployed contract.
func bindClaimRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimRegistry *ClaimRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClaimRegistry.Contract.ClaimRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimRegistry *ClaimRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimRegistry.Contract.ClaimRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimRegistry *ClaimRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimRegistry.Contract.ClaimRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimRegistry *ClaimRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClaimRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimRegistry *ClaimRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimRegistry *ClaimRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address _recipient, address _issuer, bytes32 _topic) constant returns(bytes data)
func (_ClaimRegistry *ClaimRegistryCaller) GetClaim(opts *bind.CallOpts, _recipient common.Address, _issuer common.Address, _topic [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ClaimRegistry.contract.Call(opts, out, "getClaim", _recipient, _issuer, _topic)
	return *ret0, err
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address _recipient, address _issuer, bytes32 _topic) constant returns(bytes data)
func (_ClaimRegistry *ClaimRegistrySession) GetClaim(_recipient common.Address, _issuer common.Address, _topic [32]byte) ([]byte, error) {
	return _ClaimRegistry.Contract.GetClaim(&_ClaimRegistry.CallOpts, _recipient, _issuer, _topic)
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address _recipient, address _issuer, bytes32 _topic) constant returns(bytes data)
func (_ClaimRegistry *ClaimRegistryCallerSession) GetClaim(_recipient common.Address, _issuer common.Address, _topic [32]byte) ([]byte, error) {
	return _ClaimRegistry.Contract.GetClaim(&_ClaimRegistry.CallOpts, _recipient, _issuer, _topic)
}

// GetIssuerForRecipient is a free data retrieval call binding the contract method 0x343766d0.
//
// Solidity: function getIssuerForRecipient(address _recipient) constant returns(address[] issuers)
func (_ClaimRegistry *ClaimRegistryCaller) GetIssuerForRecipient(opts *bind.CallOpts, _recipient common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ClaimRegistry.contract.Call(opts, out, "getIssuerForRecipient", _recipient)
	return *ret0, err
}

// GetIssuerForRecipient is a free data retrieval call binding the contract method 0x343766d0.
//
// Solidity: function getIssuerForRecipient(address _recipient) constant returns(address[] issuers)
func (_ClaimRegistry *ClaimRegistrySession) GetIssuerForRecipient(_recipient common.Address) ([]common.Address, error) {
	return _ClaimRegistry.Contract.GetIssuerForRecipient(&_ClaimRegistry.CallOpts, _recipient)
}

// GetIssuerForRecipient is a free data retrieval call binding the contract method 0x343766d0.
//
// Solidity: function getIssuerForRecipient(address _recipient) constant returns(address[] issuers)
func (_ClaimRegistry *ClaimRegistryCallerSession) GetIssuerForRecipient(_recipient common.Address) ([]common.Address, error) {
	return _ClaimRegistry.Contract.GetIssuerForRecipient(&_ClaimRegistry.CallOpts, _recipient)
}

// GetIssuerForRecipientTopic is a free data retrieval call binding the contract method 0x3378b1c1.
//
// Solidity: function getIssuerForRecipientTopic(address _recipient, bytes32 _topic) constant returns(address[] issuers)
func (_ClaimRegistry *ClaimRegistryCaller) GetIssuerForRecipientTopic(opts *bind.CallOpts, _recipient common.Address, _topic [32]byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ClaimRegistry.contract.Call(opts, out, "getIssuerForRecipientTopic", _recipient, _topic)
	return *ret0, err
}

// GetIssuerForRecipientTopic is a free data retrieval call binding the contract method 0x3378b1c1.
//
// Solidity: function getIssuerForRecipientTopic(address _recipient, bytes32 _topic) constant returns(address[] issuers)
func (_ClaimRegistry *ClaimRegistrySession) GetIssuerForRecipientTopic(_recipient common.Address, _topic [32]byte) ([]common.Address, error) {
	return _ClaimRegistry.Contract.GetIssuerForRecipientTopic(&_ClaimRegistry.CallOpts, _recipient, _topic)
}

// GetIssuerForRecipientTopic is a free data retrieval call binding the contract method 0x3378b1c1.
//
// Solidity: function getIssuerForRecipientTopic(address _recipient, bytes32 _topic) constant returns(address[] issuers)
func (_ClaimRegistry *ClaimRegistryCallerSession) GetIssuerForRecipientTopic(_recipient common.Address, _topic [32]byte) ([]common.Address, error) {
	return _ClaimRegistry.Contract.GetIssuerForRecipientTopic(&_ClaimRegistry.CallOpts, _recipient, _topic)
}

// GetIssuerTopicForRecipient is a free data retrieval call binding the contract method 0x411420d2.
//
// Solidity: function getIssuerTopicForRecipient(address _recipient, address _issuer) constant returns(bytes32[] topics)
func (_ClaimRegistry *ClaimRegistryCaller) GetIssuerTopicForRecipient(opts *bind.CallOpts, _recipient common.Address, _issuer common.Address) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _ClaimRegistry.contract.Call(opts, out, "getIssuerTopicForRecipient", _recipient, _issuer)
	return *ret0, err
}

// GetIssuerTopicForRecipient is a free data retrieval call binding the contract method 0x411420d2.
//
// Solidity: function getIssuerTopicForRecipient(address _recipient, address _issuer) constant returns(bytes32[] topics)
func (_ClaimRegistry *ClaimRegistrySession) GetIssuerTopicForRecipient(_recipient common.Address, _issuer common.Address) ([][32]byte, error) {
	return _ClaimRegistry.Contract.GetIssuerTopicForRecipient(&_ClaimRegistry.CallOpts, _recipient, _issuer)
}

// GetIssuerTopicForRecipient is a free data retrieval call binding the contract method 0x411420d2.
//
// Solidity: function getIssuerTopicForRecipient(address _recipient, address _issuer) constant returns(bytes32[] topics)
func (_ClaimRegistry *ClaimRegistryCallerSession) GetIssuerTopicForRecipient(_recipient common.Address, _issuer common.Address) ([][32]byte, error) {
	return _ClaimRegistry.Contract.GetIssuerTopicForRecipient(&_ClaimRegistry.CallOpts, _recipient, _issuer)
}

// GetIssuers is a free data retrieval call binding the contract method 0xa594da95.
//
// Solidity: function getIssuers() constant returns(address[] issuers)
func (_ClaimRegistry *ClaimRegistryCaller) GetIssuers(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ClaimRegistry.contract.Call(opts, out, "getIssuers")
	return *ret0, err
}

// GetIssuers is a free data retrieval call binding the contract method 0xa594da95.
//
// Solidity: function getIssuers() constant returns(address[] issuers)
func (_ClaimRegistry *ClaimRegistrySession) GetIssuers() ([]common.Address, error) {
	return _ClaimRegistry.Contract.GetIssuers(&_ClaimRegistry.CallOpts)
}

// GetIssuers is a free data retrieval call binding the contract method 0xa594da95.
//
// Solidity: function getIssuers() constant returns(address[] issuers)
func (_ClaimRegistry *ClaimRegistryCallerSession) GetIssuers() ([]common.Address, error) {
	return _ClaimRegistry.Contract.GetIssuers(&_ClaimRegistry.CallOpts)
}

// GetRecipientForIssuerTopic is a free data retrieval call binding the contract method 0xa7dd744c.
//
// Solidity: function getRecipientForIssuerTopic(address _issuer, bytes32 _topic) constant returns(address[] recipients)
func (_ClaimRegistry *ClaimRegistryCaller) GetRecipientForIssuerTopic(opts *bind.CallOpts, _issuer common.Address, _topic [32]byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ClaimRegistry.contract.Call(opts, out, "getRecipientForIssuerTopic", _issuer, _topic)
	return *ret0, err
}

// GetRecipientForIssuerTopic is a free data retrieval call binding the contract method 0xa7dd744c.
//
// Solidity: function getRecipientForIssuerTopic(address _issuer, bytes32 _topic) constant returns(address[] recipients)
func (_ClaimRegistry *ClaimRegistrySession) GetRecipientForIssuerTopic(_issuer common.Address, _topic [32]byte) ([]common.Address, error) {
	return _ClaimRegistry.Contract.GetRecipientForIssuerTopic(&_ClaimRegistry.CallOpts, _issuer, _topic)
}

// GetRecipientForIssuerTopic is a free data retrieval call binding the contract method 0xa7dd744c.
//
// Solidity: function getRecipientForIssuerTopic(address _issuer, bytes32 _topic) constant returns(address[] recipients)
func (_ClaimRegistry *ClaimRegistryCallerSession) GetRecipientForIssuerTopic(_issuer common.Address, _topic [32]byte) ([]common.Address, error) {
	return _ClaimRegistry.Contract.GetRecipientForIssuerTopic(&_ClaimRegistry.CallOpts, _issuer, _topic)
}

// GetRecipients is a free data retrieval call binding the contract method 0xd78d610b.
//
// Solidity: function getRecipients() constant returns(address[] recipients)
func (_ClaimRegistry *ClaimRegistryCaller) GetRecipients(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ClaimRegistry.contract.Call(opts, out, "getRecipients")
	return *ret0, err
}

// GetRecipients is a free data retrieval call binding the contract method 0xd78d610b.
//
// Solidity: function getRecipients() constant returns(address[] recipients)
func (_ClaimRegistry *ClaimRegistrySession) GetRecipients() ([]common.Address, error) {
	return _ClaimRegistry.Contract.GetRecipients(&_ClaimRegistry.CallOpts)
}

// GetRecipients is a free data retrieval call binding the contract method 0xd78d610b.
//
// Solidity: function getRecipients() constant returns(address[] recipients)
func (_ClaimRegistry *ClaimRegistryCallerSession) GetRecipients() ([]common.Address, error) {
	return _ClaimRegistry.Contract.GetRecipients(&_ClaimRegistry.CallOpts)
}

// GetTopics is a free data retrieval call binding the contract method 0xfee01177.
//
// Solidity: function getTopics() constant returns(bytes32[] topics)
func (_ClaimRegistry *ClaimRegistryCaller) GetTopics(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _ClaimRegistry.contract.Call(opts, out, "getTopics")
	return *ret0, err
}

// GetTopics is a free data retrieval call binding the contract method 0xfee01177.
//
// Solidity: function getTopics() constant returns(bytes32[] topics)
func (_ClaimRegistry *ClaimRegistrySession) GetTopics() ([][32]byte, error) {
	return _ClaimRegistry.Contract.GetTopics(&_ClaimRegistry.CallOpts)
}

// GetTopics is a free data retrieval call binding the contract method 0xfee01177.
//
// Solidity: function getTopics() constant returns(bytes32[] topics)
func (_ClaimRegistry *ClaimRegistryCallerSession) GetTopics() ([][32]byte, error) {
	return _ClaimRegistry.Contract.GetTopics(&_ClaimRegistry.CallOpts)
}

// HasIssuedForRecipient is a free data retrieval call binding the contract method 0x0c131a36.
//
// Solidity: function hasIssuedForRecipient(address _recipient, address _issuer) constant returns(bool issued)
func (_ClaimRegistry *ClaimRegistryCaller) HasIssuedForRecipient(opts *bind.CallOpts, _recipient common.Address, _issuer common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ClaimRegistry.contract.Call(opts, out, "hasIssuedForRecipient", _recipient, _issuer)
	return *ret0, err
}

// HasIssuedForRecipient is a free data retrieval call binding the contract method 0x0c131a36.
//
// Solidity: function hasIssuedForRecipient(address _recipient, address _issuer) constant returns(bool issued)
func (_ClaimRegistry *ClaimRegistrySession) HasIssuedForRecipient(_recipient common.Address, _issuer common.Address) (bool, error) {
	return _ClaimRegistry.Contract.HasIssuedForRecipient(&_ClaimRegistry.CallOpts, _recipient, _issuer)
}

// HasIssuedForRecipient is a free data retrieval call binding the contract method 0x0c131a36.
//
// Solidity: function hasIssuedForRecipient(address _recipient, address _issuer) constant returns(bool issued)
func (_ClaimRegistry *ClaimRegistryCallerSession) HasIssuedForRecipient(_recipient common.Address, _issuer common.Address) (bool, error) {
	return _ClaimRegistry.Contract.HasIssuedForRecipient(&_ClaimRegistry.CallOpts, _recipient, _issuer)
}

// HasIssuedTopic is a free data retrieval call binding the contract method 0x03915fac.
//
// Solidity: function hasIssuedTopic(address _recipient, address _issuer, bytes32 _topic) constant returns(bool issued)
func (_ClaimRegistry *ClaimRegistryCaller) HasIssuedTopic(opts *bind.CallOpts, _recipient common.Address, _issuer common.Address, _topic [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ClaimRegistry.contract.Call(opts, out, "hasIssuedTopic", _recipient, _issuer, _topic)
	return *ret0, err
}

// HasIssuedTopic is a free data retrieval call binding the contract method 0x03915fac.
//
// Solidity: function hasIssuedTopic(address _recipient, address _issuer, bytes32 _topic) constant returns(bool issued)
func (_ClaimRegistry *ClaimRegistrySession) HasIssuedTopic(_recipient common.Address, _issuer common.Address, _topic [32]byte) (bool, error) {
	return _ClaimRegistry.Contract.HasIssuedTopic(&_ClaimRegistry.CallOpts, _recipient, _issuer, _topic)
}

// HasIssuedTopic is a free data retrieval call binding the contract method 0x03915fac.
//
// Solidity: function hasIssuedTopic(address _recipient, address _issuer, bytes32 _topic) constant returns(bool issued)
func (_ClaimRegistry *ClaimRegistryCallerSession) HasIssuedTopic(_recipient common.Address, _issuer common.Address, _topic [32]byte) (bool, error) {
	return _ClaimRegistry.Contract.HasIssuedTopic(&_ClaimRegistry.CallOpts, _recipient, _issuer, _topic)
}

// IsRecipient is a free data retrieval call binding the contract method 0x8c5143ea.
//
// Solidity: function isRecipient(address _recipient) constant returns(bool alreadyRecipent)
func (_ClaimRegistry *ClaimRegistryCaller) IsRecipient(opts *bind.CallOpts, _recipient common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ClaimRegistry.contract.Call(opts, out, "isRecipient", _recipient)
	return *ret0, err
}

// IsRecipient is a free data retrieval call binding the contract method 0x8c5143ea.
//
// Solidity: function isRecipient(address _recipient) constant returns(bool alreadyRecipent)
func (_ClaimRegistry *ClaimRegistrySession) IsRecipient(_recipient common.Address) (bool, error) {
	return _ClaimRegistry.Contract.IsRecipient(&_ClaimRegistry.CallOpts, _recipient)
}

// IsRecipient is a free data retrieval call binding the contract method 0x8c5143ea.
//
// Solidity: function isRecipient(address _recipient) constant returns(bool alreadyRecipent)
func (_ClaimRegistry *ClaimRegistryCallerSession) IsRecipient(_recipient common.Address) (bool, error) {
	return _ClaimRegistry.Contract.IsRecipient(&_ClaimRegistry.CallOpts, _recipient)
}

// Registry is a free data retrieval call binding the contract method 0xe6d970aa.
//
// Solidity: function registry(address , address , bytes32 ) constant returns(bytes)
func (_ClaimRegistry *ClaimRegistryCaller) Registry(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ClaimRegistry.contract.Call(opts, out, "registry", arg0, arg1, arg2)
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0xe6d970aa.
//
// Solidity: function registry(address , address , bytes32 ) constant returns(bytes)
func (_ClaimRegistry *ClaimRegistrySession) Registry(arg0 common.Address, arg1 common.Address, arg2 [32]byte) ([]byte, error) {
	return _ClaimRegistry.Contract.Registry(&_ClaimRegistry.CallOpts, arg0, arg1, arg2)
}

// Registry is a free data retrieval call binding the contract method 0xe6d970aa.
//
// Solidity: function registry(address , address , bytes32 ) constant returns(bytes)
func (_ClaimRegistry *ClaimRegistryCallerSession) Registry(arg0 common.Address, arg1 common.Address, arg2 [32]byte) ([]byte, error) {
	return _ClaimRegistry.Contract.Registry(&_ClaimRegistry.CallOpts, arg0, arg1, arg2)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xa1c8ca7e.
//
// Solidity: function removeClaim(address _recipient, bytes32 _topic) returns()
func (_ClaimRegistry *ClaimRegistryTransactor) RemoveClaim(opts *bind.TransactOpts, _recipient common.Address, _topic [32]byte) (*types.Transaction, error) {
	return _ClaimRegistry.contract.Transact(opts, "removeClaim", _recipient, _topic)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xa1c8ca7e.
//
// Solidity: function removeClaim(address _recipient, bytes32 _topic) returns()
func (_ClaimRegistry *ClaimRegistrySession) RemoveClaim(_recipient common.Address, _topic [32]byte) (*types.Transaction, error) {
	return _ClaimRegistry.Contract.RemoveClaim(&_ClaimRegistry.TransactOpts, _recipient, _topic)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xa1c8ca7e.
//
// Solidity: function removeClaim(address _recipient, bytes32 _topic) returns()
func (_ClaimRegistry *ClaimRegistryTransactorSession) RemoveClaim(_recipient common.Address, _topic [32]byte) (*types.Transaction, error) {
	return _ClaimRegistry.Contract.RemoveClaim(&_ClaimRegistry.TransactOpts, _recipient, _topic)
}

// SetClaim is a paid mutator transaction binding the contract method 0x77640f9c.
//
// Solidity: function setClaim(address _recipient, bytes32 _topic, bytes _data) returns()
func (_ClaimRegistry *ClaimRegistryTransactor) SetClaim(opts *bind.TransactOpts, _recipient common.Address, _topic [32]byte, _data []byte) (*types.Transaction, error) {
	return _ClaimRegistry.contract.Transact(opts, "setClaim", _recipient, _topic, _data)
}

// SetClaim is a paid mutator transaction binding the contract method 0x77640f9c.
//
// Solidity: function setClaim(address _recipient, bytes32 _topic, bytes _data) returns()
func (_ClaimRegistry *ClaimRegistrySession) SetClaim(_recipient common.Address, _topic [32]byte, _data []byte) (*types.Transaction, error) {
	return _ClaimRegistry.Contract.SetClaim(&_ClaimRegistry.TransactOpts, _recipient, _topic, _data)
}

// SetClaim is a paid mutator transaction binding the contract method 0x77640f9c.
//
// Solidity: function setClaim(address _recipient, bytes32 _topic, bytes _data) returns()
func (_ClaimRegistry *ClaimRegistryTransactorSession) SetClaim(_recipient common.Address, _topic [32]byte, _data []byte) (*types.Transaction, error) {
	return _ClaimRegistry.Contract.SetClaim(&_ClaimRegistry.TransactOpts, _recipient, _topic, _data)
}

// ClaimRegistryClaimRemovedIterator is returned from FilterClaimRemoved and is used to iterate over the raw logs and unpacked data for ClaimRemoved events raised by the ClaimRegistry contract.
type ClaimRegistryClaimRemovedIterator struct {
	Event *ClaimRegistryClaimRemoved // Event containing the contract specifics and raw log

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
func (it *ClaimRegistryClaimRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimRegistryClaimRemoved)
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
		it.Event = new(ClaimRegistryClaimRemoved)
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
func (it *ClaimRegistryClaimRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimRegistryClaimRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimRegistryClaimRemoved represents a ClaimRemoved event raised by the ClaimRegistry contract.
type ClaimRegistryClaimRemoved struct {
	Issuer    common.Address
	Recipient common.Address
	Topic     [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimRemoved is a free log retrieval operation binding the contract event 0x895b934151a219f5c7dd6868b1397f44e9435ef5969693eed4bdc1c74dff2df1.
//
// Solidity: event ClaimRemoved(address indexed issuer, address indexed recipient, bytes32 indexed topic)
func (_ClaimRegistry *ClaimRegistryFilterer) FilterClaimRemoved(opts *bind.FilterOpts, issuer []common.Address, recipient []common.Address, topic [][32]byte) (*ClaimRegistryClaimRemovedIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	logs, sub, err := _ClaimRegistry.contract.FilterLogs(opts, "ClaimRemoved", issuerRule, recipientRule, topicRule)
	if err != nil {
		return nil, err
	}
	return &ClaimRegistryClaimRemovedIterator{contract: _ClaimRegistry.contract, event: "ClaimRemoved", logs: logs, sub: sub}, nil
}

// WatchClaimRemoved is a free log subscription operation binding the contract event 0x895b934151a219f5c7dd6868b1397f44e9435ef5969693eed4bdc1c74dff2df1.
//
// Solidity: event ClaimRemoved(address indexed issuer, address indexed recipient, bytes32 indexed topic)
func (_ClaimRegistry *ClaimRegistryFilterer) WatchClaimRemoved(opts *bind.WatchOpts, sink chan<- *ClaimRegistryClaimRemoved, issuer []common.Address, recipient []common.Address, topic [][32]byte) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	logs, sub, err := _ClaimRegistry.contract.WatchLogs(opts, "ClaimRemoved", issuerRule, recipientRule, topicRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimRegistryClaimRemoved)
				if err := _ClaimRegistry.contract.UnpackLog(event, "ClaimRemoved", log); err != nil {
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

// ClaimRegistryClaimSetIterator is returned from FilterClaimSet and is used to iterate over the raw logs and unpacked data for ClaimSet events raised by the ClaimRegistry contract.
type ClaimRegistryClaimSetIterator struct {
	Event *ClaimRegistryClaimSet // Event containing the contract specifics and raw log

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
func (it *ClaimRegistryClaimSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimRegistryClaimSet)
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
		it.Event = new(ClaimRegistryClaimSet)
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
func (it *ClaimRegistryClaimSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimRegistryClaimSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimRegistryClaimSet represents a ClaimSet event raised by the ClaimRegistry contract.
type ClaimRegistryClaimSet struct {
	Issuer    common.Address
	Recipient common.Address
	Topic     [32]byte
	Value     []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimSet is a free log retrieval operation binding the contract event 0x7c49dba8ead478c5260eb86834965ce378cfaecfd9c899d0078ed9276b6e1ffa.
//
// Solidity: event ClaimSet(address indexed issuer, address indexed recipient, bytes32 indexed topic, bytes value)
func (_ClaimRegistry *ClaimRegistryFilterer) FilterClaimSet(opts *bind.FilterOpts, issuer []common.Address, recipient []common.Address, topic [][32]byte) (*ClaimRegistryClaimSetIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	logs, sub, err := _ClaimRegistry.contract.FilterLogs(opts, "ClaimSet", issuerRule, recipientRule, topicRule)
	if err != nil {
		return nil, err
	}
	return &ClaimRegistryClaimSetIterator{contract: _ClaimRegistry.contract, event: "ClaimSet", logs: logs, sub: sub}, nil
}

// WatchClaimSet is a free log subscription operation binding the contract event 0x7c49dba8ead478c5260eb86834965ce378cfaecfd9c899d0078ed9276b6e1ffa.
//
// Solidity: event ClaimSet(address indexed issuer, address indexed recipient, bytes32 indexed topic, bytes value)
func (_ClaimRegistry *ClaimRegistryFilterer) WatchClaimSet(opts *bind.WatchOpts, sink chan<- *ClaimRegistryClaimSet, issuer []common.Address, recipient []common.Address, topic [][32]byte) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	logs, sub, err := _ClaimRegistry.contract.WatchLogs(opts, "ClaimSet", issuerRule, recipientRule, topicRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimRegistryClaimSet)
				if err := _ClaimRegistry.contract.UnpackLog(event, "ClaimSet", log); err != nil {
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
