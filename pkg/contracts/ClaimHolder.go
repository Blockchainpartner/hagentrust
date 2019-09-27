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

// ClaimHolderABI is the input ABI used to generate the binding from.
const ClaimHolderABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimsByTopic\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"topics\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"removeClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"existingTopic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_topic\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_scheme\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"addClaim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"claimRequestId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_scheme\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"changeClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"getClaim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"scheme\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"recipientReview\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_claimId\",\"type\":\"bytes32\"}],\"name\":\"toggleReviewClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"claims\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"scheme\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"recipientReview\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_topic\",\"type\":\"bytes32\"}],\"name\":\"getClaimIdsByTopic\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"claimIds\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTopics\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"createdTopics\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimApprovalToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"claimId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"topic\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scheme\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"ClaimChanged\",\"type\":\"event\"}]"

// ClaimHolderBin is the compiled bytecode used for deploying new contracts.
const ClaimHolderBin = `608060405234801561001057600080fd5b506040516126793803806126798339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506125e5806100946000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063c1b3221c11610071578063c1b3221c14610368578063c9100bcb146104b7578063da0e65a41461068d578063eff0f592146106d3578063f9e9b618146108a9578063fee011771461092c576100b4565b80630e4e4dfa146100b957806318d9adab146101055780634eee424a146101475780635fa970bb1461018d57806370c374a8146101d35780638da5cb5b1461031e575b600080fd5b6100ef600480360360408110156100cf57600080fd5b81019080803590602001909291908035906020019092919050505061098b565b6040518082815260200191505060405180910390f35b6101316004803603602081101561011b57600080fd5b81019080803590602001909291905050506109b9565b6040518082815260200191505060405180910390f35b6101736004803603602081101561015d57600080fd5b81019080803590602001909291905050506109da565b604051808215151515815260200191505060405180910390f35b6101b9600480360360208110156101a357600080fd5b8101908080359060200190929190505050610ef7565b604051808215151515815260200191505060405180910390f35b610308600480360360a08110156101e957600080fd5b8101908080359060200190929190803590602001909291908035906020019064010000000081111561021a57600080fd5b82018360208201111561022c57600080fd5b8035906020019184600183028401116401000000008311171561024e57600080fd5b90919293919293908035906020019064010000000081111561026f57600080fd5b82018360208201111561028157600080fd5b803590602001918460018302840111640100000000831117156102a357600080fd5b9091929391929390803590602001906401000000008111156102c457600080fd5b8201836020820111156102d657600080fd5b803590602001918460018302840111640100000000831117156102f857600080fd5b9091929391929390505050610f17565b6040518082815260200191505060405180910390f35b61032661143f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61049d600480360360a081101561037e57600080fd5b810190808035906020019092919080359060200190929190803590602001906401000000008111156103af57600080fd5b8201836020820111156103c157600080fd5b803590602001918460018302840111640100000000831117156103e357600080fd5b90919293919293908035906020019064010000000081111561040457600080fd5b82018360208201111561041657600080fd5b8035906020019184600183028401116401000000008311171561043857600080fd5b90919293919293908035906020019064010000000081111561045957600080fd5b82018360208201111561046b57600080fd5b8035906020019184600183028401116401000000008311171561048d57600080fd5b9091929391929390505050611464565b604051808215151515815260200191505060405180910390f35b6104e3600480360360208110156104cd57600080fd5b810190808035906020019092919050505061186e565b604051808981526020018881526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200180602001806020018615151515815260200185151515158152602001848103845289818151815260200191508051906020019080838360005b8381101561057d578082015181840152602081019050610562565b50505050905090810190601f1680156105aa5780820380516001836020036101000a031916815260200191505b50848103835288818151815260200191508051906020019080838360005b838110156105e35780820151818401526020810190506105c8565b50505050905090810190601f1680156106105780820380516001836020036101000a031916815260200191505b50848103825287818151815260200191508051906020019080838360005b8381101561064957808201518184015260208101905061062e565b50505050905090810190601f1680156106765780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390f35b6106b9600480360360208110156106a357600080fd5b8101908080359060200190929190505050611ae1565b604051808215151515815260200191505060405180910390f35b6106ff600480360360208110156106e957600080fd5b8101908080359060200190929190505050611e8f565b604051808981526020018881526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200180602001806020018615151515815260200185151515158152602001848103845289818151815260200191508051906020019080838360005b8381101561079957808201518184015260208101905061077e565b50505050905090810190601f1680156107c65780820380516001836020036101000a031916815260200191505b50848103835288818151815260200191508051906020019080838360005b838110156107ff5780820151818401526020810190506107e4565b50505050905090810190601f16801561082c5780820380516001836020036101000a031916815260200191505b50848103825287818151815260200191508051906020019080838360005b8381101561086557808201518184015260208101905061084a565b50505050905090810190601f1680156108925780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390f35b6108d5600480360360208110156108bf57600080fd5b81019080803590602001909291905050506120d9565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156109185780820151818401526020810190506108fd565b505050509050019250505060405180910390f35b610934612286565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561097757808201518184015260208101905061095c565b505050509050019250505060405180910390f35b600460205281600052604060002081815481106109a457fe5b90600052602060002001600091509150505481565b600181815481106109c657fe5b906000526020600020016000915090505481565b6000813373ffffffffffffffffffffffffffffffffffffffff166003600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a97576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061258c6025913960400191505060405180910390fd5b60006003600085815260200190815260200160002060000154905060006003600086815260200190815260200160002060010154905060006003600087815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506060600360008881526020019081526020016000206003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bb45780601f10610b8957610100808354040283529160200191610bb4565b820191906000526020600020905b815481529060010190602001808311610b9757829003601f168201915b505050505090506060600360008981526020019081526020016000206004018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c675780601f10610c3c57610100808354040283529160200191610c67565b820191906000526020600020905b815481529060010190602001808311610c4a57829003601f168201915b505050505090506060600360008a81526020019081526020016000206005018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d1a5780601f10610cef57610100808354040283529160200191610d1a565b820191906000526020600020905b815481529060010190602001808311610cfd57829003601f168201915b505050505090506000600360008b815260200190815260200160002060060160016101000a81548160ff0219169083151502179055508373ffffffffffffffffffffffffffffffffffffffff16868a7fa44ca5fcc7ed253c9105180bf33122376581343724070fdfefe298b5368ecd688887878760405180858152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b83811015610ddc578082015181840152602081019050610dc1565b50505050905090810190601f168015610e095780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b83811015610e42578082015181840152602081019050610e27565b50505050905090810190601f168015610e6f5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015610ea8578082015181840152602081019050610e8d565b50505050905090810190601f168015610ed55780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a46001975050505050505050919050565b60026020528060005260406000206000915054906101000a900460ff1681565b60008061101e88888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061101033308e8b8b604051602001808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018481526020018383808284378083019250505095505050505050604051602081830303815290604052805190602001206122de565b61233690919063ffffffff16565b90503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146110a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180612560602c913960400191505060405180910390fd5b336000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff168b604051602001808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018281526020019350505050604051602081830303815290604052805190602001209150600260008b815260200190815260200160002060009054906101000a900460ff166111da5760018a90806001815401808255809150509060018203906000526020600020016000909192909190915055506001600260008c815260200190815260200160002060006101000a81548160ff0219169083151502179055505b896003600084815260200190815260200160002060000181905550886003600084815260200190815260200160002060010181905550336003600084815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550878760036000858152602001908152602001600020600301919061128a92919061243a565b5085856003600085815260200190815260200160002060040191906112b092919061243a565b5083836003600085815260200190815260200160002060050191906112d69291906124ba565b5060016003600084815260200190815260200160002060060160016101000a81548160ff021916908315150217905550600460008b81526020019081526020016000208290806001815401808255809150509060018203906000526020600020016000909192909190915055503373ffffffffffffffffffffffffffffffffffffffff168a837f3d66b15bc1454b30bbc371d391357b9143a121b652afb35cc9fbc987ba1dee228c8c8c8c8c8c8c6040518088815260200180602001806020018060200184810384528a8a82818152602001925080828437600081840152601f19601f8201169050808301925050508481038352888882818152602001925080828437600081840152601f19601f8201169050808301925050508481038252868682818152602001925080828437600081840152601f19601f8201169050808301925050509a505050505050505050505060405180910390a48191505098975050505050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000883373ffffffffffffffffffffffffffffffffffffffff166003600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611521576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061258c6025913960400191505060405180910390fd5b6000600360008c8152602001908152602001600020905060006116428a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050611634333086600001548d8d604051602001808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014018481526020018383808284378083019250505095505050505050604051602081830303815290604052805190602001206122de565b61233690919063ffffffff16565b90503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146116c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180612560602c913960400191505060405180910390fd5b8a826001018190555089898360030191906116e492919061243a565b5087878360040191906116f892919061243a565b50858583600501919061170c9291906124ba565b5060008260060160006101000a81548160ff02191690831515021790555060018260060160016101000a81548160ff0219169083151502179055508160020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1682600001548d7f42b484727dbcbd10cc3598de8024439705ba74ffda9ada1c3e30d413f18a1bcb8e8e8e8e8e8e8e6040518088815260200180602001806020018060200184810384528a8a82818152602001925080828437600081840152601f19601f8201169050808301925050508481038352888882818152602001925080828437600081840152601f19601f8201169050808301925050508481038252868682818152602001925080828437600081840152601f19601f8201169050808301925050509a505050505050505050505060405180910390a46001935050505098975050505050505050565b600080600060608060606000806000600360008b81526020019081526020016000209050806000015481600101548260020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168360030184600401856005018660060160009054906101000a900460ff168760060160019054906101000a900460ff16848054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156119865780601f1061195b57610100808354040283529160200191611986565b820191906000526020600020905b81548152906001019060200180831161196957829003601f168201915b50505050509450838054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611a225780601f106119f757610100808354040283529160200191611a22565b820191906000526020600020905b815481529060010190602001808311611a0557829003601f168201915b50505050509350828054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611abe5780601f10611a9357610100808354040283529160200191611abe565b820191906000526020600020905b815481529060010190602001808311611aa157829003601f168201915b505050505092509850985098509850985098509850985050919395975091939597565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611ba5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f6d73672e73656e6465722073686f756c6420626520746865206f776e6572000081525060200191505060405180910390fd5b6003600083815260200190815260200160002060060160009054906101000a900460ff16156003600084815260200190815260200160002060060160006101000a81548160ff0219169083151502179055506003600083815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060000154837ff2f05270e751fd6a99aa8facc264d713c60f184ed1bb6c88312a1261582c38e060036000878152602001908152602001600020600101546003600088815260200190815260200160002060030160036000898152602001908152602001600020600401600360008a815260200190815260200160002060050160405180858152602001806020018060200180602001848103845287818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015611d6d5780601f10611d4257610100808354040283529160200191611d6d565b820191906000526020600020905b815481529060010190602001808311611d5057829003601f168201915b5050848103835286818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015611df05780601f10611dc557610100808354040283529160200191611df0565b820191906000526020600020905b815481529060010190602001808311611dd357829003601f168201915b5050848103825285818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015611e735780601f10611e4857610100808354040283529160200191611e73565b820191906000526020600020905b815481529060010190602001808311611e5657829003601f168201915b505097505050505050505060405180910390a460019050919050565b60036020528060005260406000206000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611f6d5780601f10611f4257610100808354040283529160200191611f6d565b820191906000526020600020905b815481529060010190602001808311611f5057829003601f168201915b505050505090806004018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561200b5780601f10611fe05761010080835404028352916020019161200b565b820191906000526020600020905b815481529060010190602001808311611fee57829003601f168201915b505050505090806005018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156120a95780601f1061207e576101008083540402835291602001916120a9565b820191906000526020600020905b81548152906001019060200180831161208c57829003601f168201915b5050505050908060060160009054906101000a900460ff16908060060160019054906101000a900460ff16905088565b6060806004600084815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561213957602002820191906000526020600020905b815481526020019060010190808311612125575b50505050509050606081516040519080825280602002602001820160405280156121725781602001602082028038833980820191505090505b509050600080905060008090505b83518110156121ec576003600085838151811061219957fe5b6020026020010151815260200190815260200160002060060160019054906101000a900460ff16156121df57808383815181106121d257fe5b6020026020010181815250505b8080600101915050612180565b600182019150816040519080825280602002602001820160405280156122215781602001602082028038833980820191505090505b509450600090505b8181101561227d578383828151811061223e57fe5b60200260200101518151811061225057fe5b602002602001015185828151811061226457fe5b6020026020010181815250508080600101915050612229565b50505050919050565b606060018054806020026020016040519081016040528092919081815260200182805480156122d457602002820191906000526020600020905b8154815260200190600101908083116122c0575b5050505050905090565b60008160405160200180807f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250601c01828152602001915050604051602081830303815290604052805190602001209050919050565b6000604182511461234a5760009050612434565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c111561239e5760009350505050612434565b601b8160ff16141580156123b65750601c8160ff1614155b156123c75760009350505050612434565b60018682858560405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015612424573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061247b57803560ff19168380011785556124a9565b828001600101855582156124a9579182015b828111156124a857823582559160200191906001019061248d565b5b5090506124b6919061253a565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106124fb57803560ff1916838001178555612529565b82800160010185558215612529579182015b8281111561252857823582559160200191906001019061250d565b5b509050612536919061253a565b5090565b61255c91905b80821115612558576000816000905550600101612540565b5090565b9056fe436c61696d732073656e64657220646f6573206e6f74207365656d20746f206265206d73672e73656e6465726d73672e73656e6465722073686f756c642062652074686520636c61696d20697373756572a265627a7a72315820825011bdb2ccbc719943c9a43a99548006f6959aa58581b31f9ac3646b854a8c64736f6c634300050b0032`

// DeployClaimHolder deploys a new Ethereum contract, binding an instance of ClaimHolder to it.
func DeployClaimHolder(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *ClaimHolder, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimHolderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClaimHolderBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClaimHolder{ClaimHolderCaller: ClaimHolderCaller{contract: contract}, ClaimHolderTransactor: ClaimHolderTransactor{contract: contract}, ClaimHolderFilterer: ClaimHolderFilterer{contract: contract}}, nil
}

// ClaimHolder is an auto generated Go binding around an Ethereum contract.
type ClaimHolder struct {
	ClaimHolderCaller     // Read-only binding to the contract
	ClaimHolderTransactor // Write-only binding to the contract
	ClaimHolderFilterer   // Log filterer for contract events
}

// ClaimHolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimHolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimHolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimHolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimHolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimHolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimHolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimHolderSession struct {
	Contract     *ClaimHolder      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimHolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimHolderCallerSession struct {
	Contract *ClaimHolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ClaimHolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimHolderTransactorSession struct {
	Contract     *ClaimHolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ClaimHolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimHolderRaw struct {
	Contract *ClaimHolder // Generic contract binding to access the raw methods on
}

// ClaimHolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimHolderCallerRaw struct {
	Contract *ClaimHolderCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimHolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimHolderTransactorRaw struct {
	Contract *ClaimHolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimHolder creates a new instance of ClaimHolder, bound to a specific deployed contract.
func NewClaimHolder(address common.Address, backend bind.ContractBackend) (*ClaimHolder, error) {
	contract, err := bindClaimHolder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimHolder{ClaimHolderCaller: ClaimHolderCaller{contract: contract}, ClaimHolderTransactor: ClaimHolderTransactor{contract: contract}, ClaimHolderFilterer: ClaimHolderFilterer{contract: contract}}, nil
}

// NewClaimHolderCaller creates a new read-only instance of ClaimHolder, bound to a specific deployed contract.
func NewClaimHolderCaller(address common.Address, caller bind.ContractCaller) (*ClaimHolderCaller, error) {
	contract, err := bindClaimHolder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderCaller{contract: contract}, nil
}

// NewClaimHolderTransactor creates a new write-only instance of ClaimHolder, bound to a specific deployed contract.
func NewClaimHolderTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimHolderTransactor, error) {
	contract, err := bindClaimHolder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderTransactor{contract: contract}, nil
}

// NewClaimHolderFilterer creates a new log filterer instance of ClaimHolder, bound to a specific deployed contract.
func NewClaimHolderFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimHolderFilterer, error) {
	contract, err := bindClaimHolder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderFilterer{contract: contract}, nil
}

// bindClaimHolder binds a generic wrapper to an already deployed contract.
func bindClaimHolder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimHolderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimHolder *ClaimHolderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClaimHolder.Contract.ClaimHolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimHolder *ClaimHolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimHolder.Contract.ClaimHolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimHolder *ClaimHolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimHolder.Contract.ClaimHolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimHolder *ClaimHolderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ClaimHolder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimHolder *ClaimHolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimHolder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimHolder *ClaimHolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimHolder.Contract.contract.Transact(opts, method, params...)
}

// Claims is a free data retrieval call binding the contract method 0xeff0f592.
//
// Solidity: function claims(bytes32 ) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes signature, bytes data, string uri, bool recipientReview, bool isValid)
func (_ClaimHolder *ClaimHolderCaller) Claims(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Topic           [32]byte
	Scheme          *big.Int
	Issuer          common.Address
	Signature       []byte
	Data            []byte
	Uri             string
	RecipientReview bool
	IsValid         bool
}, error) {
	ret := new(struct {
		Topic           [32]byte
		Scheme          *big.Int
		Issuer          common.Address
		Signature       []byte
		Data            []byte
		Uri             string
		RecipientReview bool
		IsValid         bool
	})
	out := ret
	err := _ClaimHolder.contract.Call(opts, out, "claims", arg0)
	return *ret, err
}

// Claims is a free data retrieval call binding the contract method 0xeff0f592.
//
// Solidity: function claims(bytes32 ) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes signature, bytes data, string uri, bool recipientReview, bool isValid)
func (_ClaimHolder *ClaimHolderSession) Claims(arg0 [32]byte) (struct {
	Topic           [32]byte
	Scheme          *big.Int
	Issuer          common.Address
	Signature       []byte
	Data            []byte
	Uri             string
	RecipientReview bool
	IsValid         bool
}, error) {
	return _ClaimHolder.Contract.Claims(&_ClaimHolder.CallOpts, arg0)
}

// Claims is a free data retrieval call binding the contract method 0xeff0f592.
//
// Solidity: function claims(bytes32 ) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes signature, bytes data, string uri, bool recipientReview, bool isValid)
func (_ClaimHolder *ClaimHolderCallerSession) Claims(arg0 [32]byte) (struct {
	Topic           [32]byte
	Scheme          *big.Int
	Issuer          common.Address
	Signature       []byte
	Data            []byte
	Uri             string
	RecipientReview bool
	IsValid         bool
}, error) {
	return _ClaimHolder.Contract.Claims(&_ClaimHolder.CallOpts, arg0)
}

// ClaimsByTopic is a free data retrieval call binding the contract method 0x0e4e4dfa.
//
// Solidity: function claimsByTopic(bytes32 , uint256 ) constant returns(bytes32)
func (_ClaimHolder *ClaimHolderCaller) ClaimsByTopic(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ClaimHolder.contract.Call(opts, out, "claimsByTopic", arg0, arg1)
	return *ret0, err
}

// ClaimsByTopic is a free data retrieval call binding the contract method 0x0e4e4dfa.
//
// Solidity: function claimsByTopic(bytes32 , uint256 ) constant returns(bytes32)
func (_ClaimHolder *ClaimHolderSession) ClaimsByTopic(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _ClaimHolder.Contract.ClaimsByTopic(&_ClaimHolder.CallOpts, arg0, arg1)
}

// ClaimsByTopic is a free data retrieval call binding the contract method 0x0e4e4dfa.
//
// Solidity: function claimsByTopic(bytes32 , uint256 ) constant returns(bytes32)
func (_ClaimHolder *ClaimHolderCallerSession) ClaimsByTopic(arg0 [32]byte, arg1 *big.Int) ([32]byte, error) {
	return _ClaimHolder.Contract.ClaimsByTopic(&_ClaimHolder.CallOpts, arg0, arg1)
}

// ExistingTopic is a free data retrieval call binding the contract method 0x5fa970bb.
//
// Solidity: function existingTopic(bytes32 ) constant returns(bool)
func (_ClaimHolder *ClaimHolderCaller) ExistingTopic(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ClaimHolder.contract.Call(opts, out, "existingTopic", arg0)
	return *ret0, err
}

// ExistingTopic is a free data retrieval call binding the contract method 0x5fa970bb.
//
// Solidity: function existingTopic(bytes32 ) constant returns(bool)
func (_ClaimHolder *ClaimHolderSession) ExistingTopic(arg0 [32]byte) (bool, error) {
	return _ClaimHolder.Contract.ExistingTopic(&_ClaimHolder.CallOpts, arg0)
}

// ExistingTopic is a free data retrieval call binding the contract method 0x5fa970bb.
//
// Solidity: function existingTopic(bytes32 ) constant returns(bool)
func (_ClaimHolder *ClaimHolderCallerSession) ExistingTopic(arg0 [32]byte) (bool, error) {
	return _ClaimHolder.Contract.ExistingTopic(&_ClaimHolder.CallOpts, arg0)
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes signature, bytes data, string uri, bool recipientReview, bool isValid)
func (_ClaimHolder *ClaimHolderCaller) GetClaim(opts *bind.CallOpts, _claimId [32]byte) (struct {
	Topic           [32]byte
	Scheme          *big.Int
	Issuer          common.Address
	Signature       []byte
	Data            []byte
	Uri             string
	RecipientReview bool
	IsValid         bool
}, error) {
	ret := new(struct {
		Topic           [32]byte
		Scheme          *big.Int
		Issuer          common.Address
		Signature       []byte
		Data            []byte
		Uri             string
		RecipientReview bool
		IsValid         bool
	})
	out := ret
	err := _ClaimHolder.contract.Call(opts, out, "getClaim", _claimId)
	return *ret, err
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes signature, bytes data, string uri, bool recipientReview, bool isValid)
func (_ClaimHolder *ClaimHolderSession) GetClaim(_claimId [32]byte) (struct {
	Topic           [32]byte
	Scheme          *big.Int
	Issuer          common.Address
	Signature       []byte
	Data            []byte
	Uri             string
	RecipientReview bool
	IsValid         bool
}, error) {
	return _ClaimHolder.Contract.GetClaim(&_ClaimHolder.CallOpts, _claimId)
}

// GetClaim is a free data retrieval call binding the contract method 0xc9100bcb.
//
// Solidity: function getClaim(bytes32 _claimId) constant returns(bytes32 topic, uint256 scheme, address issuer, bytes signature, bytes data, string uri, bool recipientReview, bool isValid)
func (_ClaimHolder *ClaimHolderCallerSession) GetClaim(_claimId [32]byte) (struct {
	Topic           [32]byte
	Scheme          *big.Int
	Issuer          common.Address
	Signature       []byte
	Data            []byte
	Uri             string
	RecipientReview bool
	IsValid         bool
}, error) {
	return _ClaimHolder.Contract.GetClaim(&_ClaimHolder.CallOpts, _claimId)
}

// GetClaimIdsByTopic is a free data retrieval call binding the contract method 0xf9e9b618.
//
// Solidity: function getClaimIdsByTopic(bytes32 _topic) constant returns(bytes32[] claimIds)
func (_ClaimHolder *ClaimHolderCaller) GetClaimIdsByTopic(opts *bind.CallOpts, _topic [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _ClaimHolder.contract.Call(opts, out, "getClaimIdsByTopic", _topic)
	return *ret0, err
}

// GetClaimIdsByTopic is a free data retrieval call binding the contract method 0xf9e9b618.
//
// Solidity: function getClaimIdsByTopic(bytes32 _topic) constant returns(bytes32[] claimIds)
func (_ClaimHolder *ClaimHolderSession) GetClaimIdsByTopic(_topic [32]byte) ([][32]byte, error) {
	return _ClaimHolder.Contract.GetClaimIdsByTopic(&_ClaimHolder.CallOpts, _topic)
}

// GetClaimIdsByTopic is a free data retrieval call binding the contract method 0xf9e9b618.
//
// Solidity: function getClaimIdsByTopic(bytes32 _topic) constant returns(bytes32[] claimIds)
func (_ClaimHolder *ClaimHolderCallerSession) GetClaimIdsByTopic(_topic [32]byte) ([][32]byte, error) {
	return _ClaimHolder.Contract.GetClaimIdsByTopic(&_ClaimHolder.CallOpts, _topic)
}

// GetTopics is a free data retrieval call binding the contract method 0xfee01177.
//
// Solidity: function getTopics() constant returns(bytes32[] createdTopics)
func (_ClaimHolder *ClaimHolderCaller) GetTopics(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _ClaimHolder.contract.Call(opts, out, "getTopics")
	return *ret0, err
}

// GetTopics is a free data retrieval call binding the contract method 0xfee01177.
//
// Solidity: function getTopics() constant returns(bytes32[] createdTopics)
func (_ClaimHolder *ClaimHolderSession) GetTopics() ([][32]byte, error) {
	return _ClaimHolder.Contract.GetTopics(&_ClaimHolder.CallOpts)
}

// GetTopics is a free data retrieval call binding the contract method 0xfee01177.
//
// Solidity: function getTopics() constant returns(bytes32[] createdTopics)
func (_ClaimHolder *ClaimHolderCallerSession) GetTopics() ([][32]byte, error) {
	return _ClaimHolder.Contract.GetTopics(&_ClaimHolder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClaimHolder *ClaimHolderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ClaimHolder.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClaimHolder *ClaimHolderSession) Owner() (common.Address, error) {
	return _ClaimHolder.Contract.Owner(&_ClaimHolder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ClaimHolder *ClaimHolderCallerSession) Owner() (common.Address, error) {
	return _ClaimHolder.Contract.Owner(&_ClaimHolder.CallOpts)
}

// Topics is a free data retrieval call binding the contract method 0x18d9adab.
//
// Solidity: function topics(uint256 ) constant returns(bytes32)
func (_ClaimHolder *ClaimHolderCaller) Topics(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ClaimHolder.contract.Call(opts, out, "topics", arg0)
	return *ret0, err
}

// Topics is a free data retrieval call binding the contract method 0x18d9adab.
//
// Solidity: function topics(uint256 ) constant returns(bytes32)
func (_ClaimHolder *ClaimHolderSession) Topics(arg0 *big.Int) ([32]byte, error) {
	return _ClaimHolder.Contract.Topics(&_ClaimHolder.CallOpts, arg0)
}

// Topics is a free data retrieval call binding the contract method 0x18d9adab.
//
// Solidity: function topics(uint256 ) constant returns(bytes32)
func (_ClaimHolder *ClaimHolderCallerSession) Topics(arg0 *big.Int) ([32]byte, error) {
	return _ClaimHolder.Contract.Topics(&_ClaimHolder.CallOpts, arg0)
}

// AddClaim is a paid mutator transaction binding the contract method 0x70c374a8.
//
// Solidity: function addClaim(bytes32 _topic, uint256 _scheme, bytes _signature, bytes _data, string _uri) returns(bytes32 claimRequestId)
func (_ClaimHolder *ClaimHolderTransactor) AddClaim(opts *bind.TransactOpts, _topic [32]byte, _scheme *big.Int, _signature []byte, _data []byte, _uri string) (*types.Transaction, error) {
	return _ClaimHolder.contract.Transact(opts, "addClaim", _topic, _scheme, _signature, _data, _uri)
}

// AddClaim is a paid mutator transaction binding the contract method 0x70c374a8.
//
// Solidity: function addClaim(bytes32 _topic, uint256 _scheme, bytes _signature, bytes _data, string _uri) returns(bytes32 claimRequestId)
func (_ClaimHolder *ClaimHolderSession) AddClaim(_topic [32]byte, _scheme *big.Int, _signature []byte, _data []byte, _uri string) (*types.Transaction, error) {
	return _ClaimHolder.Contract.AddClaim(&_ClaimHolder.TransactOpts, _topic, _scheme, _signature, _data, _uri)
}

// AddClaim is a paid mutator transaction binding the contract method 0x70c374a8.
//
// Solidity: function addClaim(bytes32 _topic, uint256 _scheme, bytes _signature, bytes _data, string _uri) returns(bytes32 claimRequestId)
func (_ClaimHolder *ClaimHolderTransactorSession) AddClaim(_topic [32]byte, _scheme *big.Int, _signature []byte, _data []byte, _uri string) (*types.Transaction, error) {
	return _ClaimHolder.Contract.AddClaim(&_ClaimHolder.TransactOpts, _topic, _scheme, _signature, _data, _uri)
}

// ChangeClaim is a paid mutator transaction binding the contract method 0xc1b3221c.
//
// Solidity: function changeClaim(bytes32 _claimId, uint256 _scheme, bytes _signature, bytes _data, string _uri) returns(bool success)
func (_ClaimHolder *ClaimHolderTransactor) ChangeClaim(opts *bind.TransactOpts, _claimId [32]byte, _scheme *big.Int, _signature []byte, _data []byte, _uri string) (*types.Transaction, error) {
	return _ClaimHolder.contract.Transact(opts, "changeClaim", _claimId, _scheme, _signature, _data, _uri)
}

// ChangeClaim is a paid mutator transaction binding the contract method 0xc1b3221c.
//
// Solidity: function changeClaim(bytes32 _claimId, uint256 _scheme, bytes _signature, bytes _data, string _uri) returns(bool success)
func (_ClaimHolder *ClaimHolderSession) ChangeClaim(_claimId [32]byte, _scheme *big.Int, _signature []byte, _data []byte, _uri string) (*types.Transaction, error) {
	return _ClaimHolder.Contract.ChangeClaim(&_ClaimHolder.TransactOpts, _claimId, _scheme, _signature, _data, _uri)
}

// ChangeClaim is a paid mutator transaction binding the contract method 0xc1b3221c.
//
// Solidity: function changeClaim(bytes32 _claimId, uint256 _scheme, bytes _signature, bytes _data, string _uri) returns(bool success)
func (_ClaimHolder *ClaimHolderTransactorSession) ChangeClaim(_claimId [32]byte, _scheme *big.Int, _signature []byte, _data []byte, _uri string) (*types.Transaction, error) {
	return _ClaimHolder.Contract.ChangeClaim(&_ClaimHolder.TransactOpts, _claimId, _scheme, _signature, _data, _uri)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_ClaimHolder *ClaimHolderTransactor) RemoveClaim(opts *bind.TransactOpts, _claimId [32]byte) (*types.Transaction, error) {
	return _ClaimHolder.contract.Transact(opts, "removeClaim", _claimId)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_ClaimHolder *ClaimHolderSession) RemoveClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _ClaimHolder.Contract.RemoveClaim(&_ClaimHolder.TransactOpts, _claimId)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0x4eee424a.
//
// Solidity: function removeClaim(bytes32 _claimId) returns(bool success)
func (_ClaimHolder *ClaimHolderTransactorSession) RemoveClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _ClaimHolder.Contract.RemoveClaim(&_ClaimHolder.TransactOpts, _claimId)
}

// ToggleReviewClaim is a paid mutator transaction binding the contract method 0xda0e65a4.
//
// Solidity: function toggleReviewClaim(bytes32 _claimId) returns(bool success)
func (_ClaimHolder *ClaimHolderTransactor) ToggleReviewClaim(opts *bind.TransactOpts, _claimId [32]byte) (*types.Transaction, error) {
	return _ClaimHolder.contract.Transact(opts, "toggleReviewClaim", _claimId)
}

// ToggleReviewClaim is a paid mutator transaction binding the contract method 0xda0e65a4.
//
// Solidity: function toggleReviewClaim(bytes32 _claimId) returns(bool success)
func (_ClaimHolder *ClaimHolderSession) ToggleReviewClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _ClaimHolder.Contract.ToggleReviewClaim(&_ClaimHolder.TransactOpts, _claimId)
}

// ToggleReviewClaim is a paid mutator transaction binding the contract method 0xda0e65a4.
//
// Solidity: function toggleReviewClaim(bytes32 _claimId) returns(bool success)
func (_ClaimHolder *ClaimHolderTransactorSession) ToggleReviewClaim(_claimId [32]byte) (*types.Transaction, error) {
	return _ClaimHolder.Contract.ToggleReviewClaim(&_ClaimHolder.TransactOpts, _claimId)
}

// ClaimHolderClaimAddedIterator is returned from FilterClaimAdded and is used to iterate over the raw logs and unpacked data for ClaimAdded events raised by the ClaimHolder contract.
type ClaimHolderClaimAddedIterator struct {
	Event *ClaimHolderClaimAdded // Event containing the contract specifics and raw log

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
func (it *ClaimHolderClaimAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimHolderClaimAdded)
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
		it.Event = new(ClaimHolderClaimAdded)
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
func (it *ClaimHolderClaimAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimHolderClaimAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimHolderClaimAdded represents a ClaimAdded event raised by the ClaimHolder contract.
type ClaimHolderClaimAdded struct {
	ClaimId   [32]byte
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature []byte
	Data      []byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimAdded is a free log retrieval operation binding the contract event 0x3d66b15bc1454b30bbc371d391357b9143a121b652afb35cc9fbc987ba1dee22.
//
// Solidity: event ClaimAdded(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) FilterClaimAdded(opts *bind.FilterOpts, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (*ClaimHolderClaimAddedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.FilterLogs(opts, "ClaimAdded", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderClaimAddedIterator{contract: _ClaimHolder.contract, event: "ClaimAdded", logs: logs, sub: sub}, nil
}

// WatchClaimAdded is a free log subscription operation binding the contract event 0x3d66b15bc1454b30bbc371d391357b9143a121b652afb35cc9fbc987ba1dee22.
//
// Solidity: event ClaimAdded(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) WatchClaimAdded(opts *bind.WatchOpts, sink chan<- *ClaimHolderClaimAdded, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.WatchLogs(opts, "ClaimAdded", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimHolderClaimAdded)
				if err := _ClaimHolder.contract.UnpackLog(event, "ClaimAdded", log); err != nil {
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

// ClaimHolderClaimApprovalToggledIterator is returned from FilterClaimApprovalToggled and is used to iterate over the raw logs and unpacked data for ClaimApprovalToggled events raised by the ClaimHolder contract.
type ClaimHolderClaimApprovalToggledIterator struct {
	Event *ClaimHolderClaimApprovalToggled // Event containing the contract specifics and raw log

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
func (it *ClaimHolderClaimApprovalToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimHolderClaimApprovalToggled)
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
		it.Event = new(ClaimHolderClaimApprovalToggled)
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
func (it *ClaimHolderClaimApprovalToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimHolderClaimApprovalToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimHolderClaimApprovalToggled represents a ClaimApprovalToggled event raised by the ClaimHolder contract.
type ClaimHolderClaimApprovalToggled struct {
	ClaimId   [32]byte
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature []byte
	Data      []byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimApprovalToggled is a free log retrieval operation binding the contract event 0xf2f05270e751fd6a99aa8facc264d713c60f184ed1bb6c88312a1261582c38e0.
//
// Solidity: event ClaimApprovalToggled(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) FilterClaimApprovalToggled(opts *bind.FilterOpts, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (*ClaimHolderClaimApprovalToggledIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.FilterLogs(opts, "ClaimApprovalToggled", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderClaimApprovalToggledIterator{contract: _ClaimHolder.contract, event: "ClaimApprovalToggled", logs: logs, sub: sub}, nil
}

// WatchClaimApprovalToggled is a free log subscription operation binding the contract event 0xf2f05270e751fd6a99aa8facc264d713c60f184ed1bb6c88312a1261582c38e0.
//
// Solidity: event ClaimApprovalToggled(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) WatchClaimApprovalToggled(opts *bind.WatchOpts, sink chan<- *ClaimHolderClaimApprovalToggled, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.WatchLogs(opts, "ClaimApprovalToggled", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimHolderClaimApprovalToggled)
				if err := _ClaimHolder.contract.UnpackLog(event, "ClaimApprovalToggled", log); err != nil {
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

// ClaimHolderClaimChangedIterator is returned from FilterClaimChanged and is used to iterate over the raw logs and unpacked data for ClaimChanged events raised by the ClaimHolder contract.
type ClaimHolderClaimChangedIterator struct {
	Event *ClaimHolderClaimChanged // Event containing the contract specifics and raw log

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
func (it *ClaimHolderClaimChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimHolderClaimChanged)
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
		it.Event = new(ClaimHolderClaimChanged)
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
func (it *ClaimHolderClaimChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimHolderClaimChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimHolderClaimChanged represents a ClaimChanged event raised by the ClaimHolder contract.
type ClaimHolderClaimChanged struct {
	ClaimId   [32]byte
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature []byte
	Data      []byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimChanged is a free log retrieval operation binding the contract event 0x42b484727dbcbd10cc3598de8024439705ba74ffda9ada1c3e30d413f18a1bcb.
//
// Solidity: event ClaimChanged(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) FilterClaimChanged(opts *bind.FilterOpts, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (*ClaimHolderClaimChangedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.FilterLogs(opts, "ClaimChanged", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderClaimChangedIterator{contract: _ClaimHolder.contract, event: "ClaimChanged", logs: logs, sub: sub}, nil
}

// WatchClaimChanged is a free log subscription operation binding the contract event 0x42b484727dbcbd10cc3598de8024439705ba74ffda9ada1c3e30d413f18a1bcb.
//
// Solidity: event ClaimChanged(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) WatchClaimChanged(opts *bind.WatchOpts, sink chan<- *ClaimHolderClaimChanged, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.WatchLogs(opts, "ClaimChanged", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimHolderClaimChanged)
				if err := _ClaimHolder.contract.UnpackLog(event, "ClaimChanged", log); err != nil {
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

// ClaimHolderClaimRemovedIterator is returned from FilterClaimRemoved and is used to iterate over the raw logs and unpacked data for ClaimRemoved events raised by the ClaimHolder contract.
type ClaimHolderClaimRemovedIterator struct {
	Event *ClaimHolderClaimRemoved // Event containing the contract specifics and raw log

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
func (it *ClaimHolderClaimRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimHolderClaimRemoved)
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
		it.Event = new(ClaimHolderClaimRemoved)
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
func (it *ClaimHolderClaimRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimHolderClaimRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimHolderClaimRemoved represents a ClaimRemoved event raised by the ClaimHolder contract.
type ClaimHolderClaimRemoved struct {
	ClaimId   [32]byte
	Topic     [32]byte
	Scheme    *big.Int
	Issuer    common.Address
	Signature []byte
	Data      []byte
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimRemoved is a free log retrieval operation binding the contract event 0xa44ca5fcc7ed253c9105180bf33122376581343724070fdfefe298b5368ecd68.
//
// Solidity: event ClaimRemoved(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) FilterClaimRemoved(opts *bind.FilterOpts, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (*ClaimHolderClaimRemovedIterator, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.FilterLogs(opts, "ClaimRemoved", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimHolderClaimRemovedIterator{contract: _ClaimHolder.contract, event: "ClaimRemoved", logs: logs, sub: sub}, nil
}

// WatchClaimRemoved is a free log subscription operation binding the contract event 0xa44ca5fcc7ed253c9105180bf33122376581343724070fdfefe298b5368ecd68.
//
// Solidity: event ClaimRemoved(bytes32 indexed claimId, bytes32 indexed topic, uint256 scheme, address indexed issuer, bytes signature, bytes data, string uri)
func (_ClaimHolder *ClaimHolderFilterer) WatchClaimRemoved(opts *bind.WatchOpts, sink chan<- *ClaimHolderClaimRemoved, claimId [][32]byte, topic [][32]byte, issuer []common.Address) (event.Subscription, error) {

	var claimIdRule []interface{}
	for _, claimIdItem := range claimId {
		claimIdRule = append(claimIdRule, claimIdItem)
	}
	var topicRule []interface{}
	for _, topicItem := range topic {
		topicRule = append(topicRule, topicItem)
	}

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _ClaimHolder.contract.WatchLogs(opts, "ClaimRemoved", claimIdRule, topicRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimHolderClaimRemoved)
				if err := _ClaimHolder.contract.UnpackLog(event, "ClaimRemoved", log); err != nil {
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
