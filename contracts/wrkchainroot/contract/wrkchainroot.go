// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/unification-com/mainchain"
	"github.com/unification-com/mainchain/accounts/abi"
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/core/types"
	"github.com/unification-com/mainchain/event"
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820893fd575257bb2e11d237797dd122e4d1936aaf78c4e0afe9376a92a5b5e215464736f6c634300050a0032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// WRKChainRootABI is the input ABI used to generate the binding from.
const WRKChainRootABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"bytes32\"},{\"name\":\"_authAddresses\",\"type\":\"address[]\"},{\"name\":\"_genesisHash\",\"type\":\"bytes32\"}],\"name\":\"registerWrkChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"bytes32\"},{\"name\":\"_height\",\"type\":\"uint256\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"name\":\"_receiptRoot\",\"type\":\"bytes32\"},{\"name\":\"_txRoot\",\"type\":\"bytes32\"},{\"name\":\"_stateRoot\",\"type\":\"bytes32\"}],\"name\":\"recordHeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"bytes32\"}],\"name\":\"getNumBlocksSubmitted\",\"outputs\":[{\"name\":\"numBlocksSubmitted_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"bytes32\"}],\"name\":\"getGenesis\",\"outputs\":[{\"name\":\"genesisHash_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isAuthorisedAddress\",\"outputs\":[{\"name\":\"isAuthorised_\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"bytes32\"},{\"name\":\"_addressesToRemove\",\"type\":\"address[]\"}],\"name\":\"removeAuthAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"bytes32\"}],\"name\":\"getLastBlock\",\"outputs\":[{\"name\":\"lastHash_\",\"type\":\"bytes32\"},{\"name\":\"lastParentHash_\",\"type\":\"bytes32\"},{\"name\":\"lastReceiptRoot_\",\"type\":\"bytes32\"},{\"name\":\"lastTxRoot_\",\"type\":\"bytes32\"},{\"name\":\"lastStateRoot_\",\"type\":\"bytes32\"},{\"name\":\"lastSubmissionTime_\",\"type\":\"uint256\"},{\"name\":\"lastSubmittedBy_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"wrkchainList\",\"outputs\":[{\"name\":\"genesisHash\",\"type\":\"bytes32\"},{\"name\":\"isWrkchain\",\"type\":\"bool\"},{\"name\":\"numAuthAddresses\",\"type\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"numBlocksSubmitted\",\"type\":\"uint256\"},{\"name\":\"numHistoricalBlocksSubmitted\",\"type\":\"uint256\"},{\"name\":\"lastBlock\",\"type\":\"uint256\"},{\"name\":\"lastHash\",\"type\":\"bytes32\"},{\"name\":\"lastParentHash\",\"type\":\"bytes32\"},{\"name\":\"lastReceiptRoot\",\"type\":\"bytes32\"},{\"name\":\"lastTxRoot\",\"type\":\"bytes32\"},{\"name\":\"lastStateRoot\",\"type\":\"bytes32\"},{\"name\":\"lastSubmissionTime\",\"type\":\"uint256\"},{\"name\":\"lastSubmittedBy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"bytes32\"},{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeWrkchainOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"bytes32\"}],\"name\":\"getLastRecordedBlockNum\",\"outputs\":[{\"name\":\"lastBlock_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"bytes32\"}],\"name\":\"chainExists\",\"outputs\":[{\"name\":\"chainExists_\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"bytes32\"},{\"name\":\"_newAuthAddresses\",\"type\":\"address[]\"}],\"name\":\"addAuthAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_height\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_receiptRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_txRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_submittedBy\",\"type\":\"address\"}],\"name\":\"RecordHeader\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_height\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_receiptRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_txRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_submittedBy\",\"type\":\"address\"}],\"name\":\"RecordHeaderHistorical\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_genesisHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"RegisterWrkChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogFallbackFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_authorisedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"AuthoriseNewAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_removedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"RemoveAuthorisedAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_old\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_new\",\"type\":\"address\"}],\"name\":\"WRKChainOwnerChanged\",\"type\":\"event\"}]"

// WRKChainRootFuncSigs maps the 4-byte function signature to its string representation.
var WRKChainRootFuncSigs = map[string]string{
	"d4202677": "addAuthAddresses(bytes32,address[])",
	"ca804ccd": "chainExists(bytes32)",
	"a51af476": "changeWrkchainOwner(bytes32,address)",
	"278f1daf": "getGenesis(bytes32)",
	"8b3e5e06": "getLastBlock(bytes32)",
	"c4d8561e": "getLastRecordedBlockNum(bytes32)",
	"0b2b1087": "getNumBlocksSubmitted(bytes32)",
	"62269f7a": "isAuthorisedAddress(bytes32,address)",
	"01242bc5": "recordHeader(bytes32,uint256,bytes32,bytes32,bytes32,bytes32,bytes32)",
	"004d6c96": "registerWrkChain(bytes32,address[],bytes32)",
	"6c5664ac": "removeAuthAddresses(bytes32,address[])",
	"8c75c5d1": "wrkchainList(bytes32)",
}

// WRKChainRootBin is the compiled bytecode used for deploying new contracts.
var WRKChainRootBin = "0x608060405234801561001057600080fd5b50611efe806100206000396000f3fe6080604052600436106100a65760003560e01c80638b3e5e06116100645780638b3e5e06146103515780638c75c5d1146103bc578063a51af4761461045d578063c4d8561e14610496578063ca804ccd146104c0578063d4202677146104ea576100a6565b80624d6c96146100de57806301242bc5146101995780630b2b1087146101e7578063278f1daf1461022357806362269f7a1461024d5780636c5664ac1461029a575b60408051348152905133917fb1db2589d999730e8faec7dccf62c4d28a7f6ef1095b46c9a744dcd46d87130e919081900360200190a2005b3480156100ea57600080fd5b506101976004803603606081101561010157600080fd5b8135919081019060408101602082013564010000000081111561012357600080fd5b82018360208201111561013557600080fd5b8035906020019184602083028401116401000000008311171561015757600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050913592506105a1915050565b005b3480156101a557600080fd5b50610197600480360360e08110156101bc57600080fd5b5080359060208101359060408101359060608101359060808101359060a08101359060c00135610a40565b3480156101f357600080fd5b506102116004803603602081101561020a57600080fd5b5035610def565b60408051918252519081900360200190f35b34801561022f57600080fd5b506102116004803603602081101561024657600080fd5b5035610ea3565b34801561025957600080fd5b506102866004803603604081101561027057600080fd5b50803590602001356001600160a01b0316610f54565b604080519115158252519081900360200190f35b3480156102a657600080fd5b50610197600480360360408110156102bd57600080fd5b813591908101906040810160208201356401000000008111156102df57600080fd5b8201836020820111156102f157600080fd5b8035906020019184602083028401116401000000008311171561031357600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061101f945050505050565b34801561035d57600080fd5b5061037b6004803603602081101561037457600080fd5b50356113d1565b604080519788526020880196909652868601949094526060860192909252608085015260a08401526001600160a01b031660c0830152519081900360e00190f35b3480156103c857600080fd5b506103e6600480360360208110156103df57600080fd5b50356114c5565b604080519e8f529c151560208f01528d8d019b909b526001600160a01b03998a1660608e015260808d019890985260a08c019690965260c08b019490945260e08a0192909252610100890152610120880152610140870152610160860152610180850152166101a083015251908190036101c00190f35b34801561046957600080fd5b506101976004803603604081101561048057600080fd5b50803590602001356001600160a01b0316611543565b3480156104a257600080fd5b50610211600480360360208110156104b957600080fd5b5035611801565b3480156104cc57600080fd5b50610286600480360360208110156104e357600080fd5b50356118b5565b3480156104f657600080fd5b506101976004803603604081101561050d57600080fd5b8135919081019060408101602082013564010000000081111561052f57600080fd5b82018360208201111561054157600080fd5b8035906020019184602083028401116401000000008311171561056357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611918945050505050565b82806105e8576040805162461bcd60e51b815260206004820152601160248201527010da185a5b881251081c995c5d5a5c9959607a1b604482015290519081900360640190fd5b60008481526020819052604090206001015460ff161561064f576040805162461bcd60e51b815260206004820152601760248201527f436861696e20494420616c726561647920657869737473000000000000000000604482015290519081900360640190fd5b81610699576040805162461bcd60e51b815260206004820152601560248201527411d95b995cda5cc81a185cda081c995c5d5a5c9959605a1b604482015290519081900360640190fd5b60008351116106d95760405162461bcd60e51b8152600401808060200182810382526025815260200180611e0d6025913960400191505060405180910390fd5b600a8351111561071a5760405162461bcd60e51b8152600401808060200182810382526027815260200180611db46027913960400191505060405180910390fd5b604051806101c0016040528083815260200160011515815260200184518152602001336001600160a01b031681526020016000815260200160008152602001600081526020018381526020016000801b81526020016000801b81526020016000801b81526020016000801b8152602001428152602001336001600160a01b03168152506000808681526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff0219169083151502179055506040820151816003015560608201518160040160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506080820151816005015560a0820151816006015560c0820151816007015560e08201518160080155610100820151816009015561012082015181600a015561014082015181600b015561016082015181600c015561018082015181600d01556101a082015181600e0160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550905050837fa70d58903899f893c484418f077825c24e4d2508e35cf7f19ec5b0aa30ef43d48342604051808381526020018281526020019250505060405180910390a260005b83518110156109ad576001600080878152602001908152602001600020600201600086848151811061091157fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550336001600160a01b031685600080516020611e7883398151915286848151811061097657fe5b602002602001015160405180826001600160a01b03166001600160a01b0316815260200191505060405180910390a36001016108e3565b5060008481526020818152604080832033845260020190915290205460ff16610a3a57600084815260208181526040808320338452600281018352908320805460ff191660019081179091558784529290915260030154610a139163ffffffff611cd516565b50604080513380825291518691600080516020611e78833981519152919081900360200190a35b50505050565b600087815260208190526040902060010154879060ff16610aa1576040805162461bcd60e51b815260206004820152601660248201527510da185a5b925108191bd95cc81b9bdd08195e1a5cdd60521b604482015290519081900360640190fd5b60008181526020818152604080832033845260020190915290205460ff16151560011480610ae857506000818152602081905260409020600401546001600160a01b031633145b610b2a576040805162461bcd60e51b815260206004820152600e60248201526d1b9bdd08185d5d1a1bdc9a5cd95960921b604482015290519081900360640190fd5b8780610b71576040805162461bcd60e51b815260206004820152601160248201527010da185a5b881251081c995c5d5a5c9959607a1b604482015290519081900360640190fd5b60008811610bbe576040805162461bcd60e51b8152602060048201526015602482015274109b1bd8dac81b9d5b58995c881c995c5d5a5c9959605a1b604482015290519081900360640190fd5b86610c00576040805162461bcd60e51b815260206004820152600d60248201526c1a185cda081c995c5d5a5c9959609a1b604482015290519081900360640190fd5b60008981526020819052604090206001015460ff16610c54576040805162461bcd60e51b81526020600482015260176024820152600080516020611d94833981519152604482015290519081900360640190fd5b60008981526020819052604090206007810154891115610d525760008a815260208190526040902060050154610c9190600163ffffffff611cd516565b60008b815260208181526040918290206005810193909355600783018c9055600883018b9055600983018a9055600a8301899055600b8301889055600c830187905542600d8401819055600e90930180546001600160a01b0319163390811790915582518b81529182018a905281830189905260608201889052608082019390935260a08101929092525189918b918d917f22f683acdb61cbd91a27907cbd5a4c5a12268ab76c79b5af59511aa03f0fe2af919081900360c00190a4610de3565b60008a815260208190526040902060060154610d7590600163ffffffff611cd516565b60008b81526020818152604091829020600601929092558051898152918201889052818101879052606082018690524260808301523360a08301525189918b918d917f599336707b966c6560001ed107c04f461a648d8c67977e6c7fbd346655890eb2919081900360c00190a45b50505050505050505050565b60008180610e38576040805162461bcd60e51b815260206004820152601160248201527010da185a5b881251081c995c5d5a5c9959607a1b604482015290519081900360640190fd5b60008381526020819052604090206001015460ff16610e8c576040805162461bcd60e51b81526020600482015260176024820152600080516020611d94833981519152604482015290519081900360640190fd5b505060009081526020819052604090206005015490565b60008180610eec576040805162461bcd60e51b815260206004820152601160248201527010da185a5b881251081c995c5d5a5c9959607a1b604482015290519081900360640190fd5b60008381526020819052604090206001015460ff16610f40576040805162461bcd60e51b81526020600482015260176024820152600080516020611d94833981519152604482015290519081900360640190fd5b505060009081526020819052604090205490565b60008280610f9d576040805162461bcd60e51b815260206004820152601160248201527010da185a5b881251081c995c5d5a5c9959607a1b604482015290519081900360640190fd5b60008481526020819052604090206001015460ff16610ff1576040805162461bcd60e51b81526020600482015260176024820152600080516020611d94833981519152604482015290519081900360640190fd5b50506000918252602082815260408084206001600160a01b0390931684526002909201905290205460ff1690565b600082815260208190526040902060010154829060ff16611080576040805162461bcd60e51b815260206004820152601660248201527510da185a5b925108191bd95cc81b9bdd08195e1a5cdd60521b604482015290519081900360640190fd5b6000818152602081905260409020600401546001600160a01b031633146110e6576040805162461bcd60e51b81526020600482015260156024820152742cb7ba9030b932903737ba103a34329037bbb732b960591b604482015290519081900360640190fd5b828061112d576040805162461bcd60e51b815260206004820152601160248201527010da185a5b881251081c995c5d5a5c9959607a1b604482015290519081900360640190fd5b600083511161117d576040805162461bcd60e51b8152602060048201526017602482015276105d5d1a081859191c995cdcd95cc81c995c5d5a5c9959604a1b604482015290519081900360640190fd5b600a835111156111be5760405162461bcd60e51b8152600401808060200182810382526023815260200180611e556023913960400191505060405180910390fd5b60008481526020819052604090206001015460ff16611212576040805162461bcd60e51b81526020600482015260176024820152600080516020611d94833981519152604482015290519081900360640190fd5b6000848152602081905260409020600301548351106112625760405162461bcd60e51b8152600401808060200182810382526032815260200180611e986032913960400191505060405180910390fd5b60005b83518110156113ca5760006001600160a01b031684828151811061128557fe5b60200260200101516001600160a01b031614156112e9576040805162461bcd60e51b815260206004820152601d60248201527f63616e6e6f7420617574686f72697365207a65726f2061646472657373000000604482015290519081900360640190fd5b600080868152602001908152602001600020600201600085838151811061130c57fe5b6020908102919091018101516001600160a01b031682528181019290925260409081016000908120805460ff191690558781529182905290206003015461135a90600163ffffffff611d3616565b50336001600160a01b0316857f1ab6ecd37e23ab30318669bb4642f1cdfba82f23405ce8dfba3648fc9b4744e386848151811061139357fe5b602002602001015160405180826001600160a01b03166001600160a01b0316815260200191505060405180910390a3600101611265565b5050505050565b60008080808080808780611420576040805162461bcd60e51b815260206004820152601160248201527010da185a5b881251081c995c5d5a5c9959607a1b604482015290519081900360640190fd5b60008981526020819052604090206001015460ff16611474576040805162461bcd60e51b81526020600482015260176024820152600080516020611d94833981519152604482015290519081900360640190fd5b5050506000958652505050602083905250506040902060088101546009820154600a830154600b840154600c850154600d860154600e9096015494969395929491939092916001600160a01b031690565b6000602081905290815260409020805460018201546003830154600484015460058501546006860154600787015460088801546009890154600a8a0154600b8b0154600c8c0154600d8d0154600e909d01549b9c60ff909b169b999a6001600160a01b03998a169a989997989697959694959394929391929091168e565b600082815260208190526040902060010154829060ff166115a4576040805162461bcd60e51b815260206004820152601660248201527510da185a5b925108191bd95cc81b9bdd08195e1a5cdd60521b604482015290519081900360640190fd5b6000818152602081905260409020600401546001600160a01b0316331461160a576040805162461bcd60e51b81526020600482015260156024820152742cb7ba9030b932903737ba103a34329037bbb732b960591b604482015290519081900360640190fd5b8280611651576040805162461bcd60e51b815260206004820152601160248201527010da185a5b881251081c995c5d5a5c9959607a1b604482015290519081900360640190fd5b6001600160a01b0383166116965760405162461bcd60e51b8152600401808060200182810382526032815260200180611ddb6032913960400191505060405180910390fd5b60008481526020819052604090206001015460ff166116ea576040805162461bcd60e51b81526020600482015260176024820152600080516020611d94833981519152604482015290519081900360640190fd5b600084815260208181526040918290206004810180546001600160a01b0319166001600160a01b0388169081179091558351338152928301528251909287927fa59f2eb1d0ae2221e2f9854e3bbc9ff4e5edd43b124d0ac97b6d2d83125d1d19929081900390910190a26000858152602081815260408083206001600160a01b038816845260020190915290205460ff166113ca576000858152602081815260408083206001600160a01b0388168452600281018352908320805460ff1916600190811790915588845292909152600301546117cb9163ffffffff611cd516565b50604080516001600160a01b0386168152905133918791600080516020611e788339815191529181900360200190a35050505050565b6000818061184a576040805162461bcd60e51b815260206004820152601160248201527010da185a5b881251081c995c5d5a5c9959607a1b604482015290519081900360640190fd5b60008381526020819052604090206001015460ff1661189e576040805162461bcd60e51b81526020600482015260176024820152600080516020611d94833981519152604482015290519081900360640190fd5b505060009081526020819052604090206007015490565b600081806118fe576040805162461bcd60e51b815260206004820152601160248201527010da185a5b881251081c995c5d5a5c9959607a1b604482015290519081900360640190fd5b505060009081526020819052604090206001015460ff1690565b600082815260208190526040902060010154829060ff16611979576040805162461bcd60e51b815260206004820152601660248201527510da185a5b925108191bd95cc81b9bdd08195e1a5cdd60521b604482015290519081900360640190fd5b6000818152602081905260409020600401546001600160a01b031633146119df576040805162461bcd60e51b81526020600482015260156024820152742cb7ba9030b932903737ba103a34329037bbb732b960591b604482015290519081900360640190fd5b8280611a26576040805162461bcd60e51b815260206004820152601160248201527010da185a5b881251081c995c5d5a5c9959607a1b604482015290519081900360640190fd5b6000835111611a76576040805162461bcd60e51b8152602060048201526017602482015276105d5d1a081859191c995cdcd95cc81c995c5d5a5c9959604a1b604482015290519081900360640190fd5b600a83511115611ab75760405162461bcd60e51b8152600401808060200182810382526023815260200180611e556023913960400191505060405180910390fd5b60008481526020819052604090206001015460ff16611b0b576040805162461bcd60e51b81526020600482015260176024820152600080516020611d94833981519152604482015290519081900360640190fd5b82516000858152602081905260408120600301549091611b31919063ffffffff611cd516565b90506064811115611b735760405162461bcd60e51b8152600401808060200182810382526023815260200180611e326023913960400191505060405180910390fd5b60005b8451811015611ccd5760006001600160a01b0316858281518110611b9657fe5b60200260200101516001600160a01b03161415611bfa576040805162461bcd60e51b815260206004820152601d60248201527f63616e6e6f7420617574686f72697365207a65726f2061646472657373000000604482015290519081900360640190fd5b60016000808881526020019081526020016000206002016000878481518110611c1f57fe5b6020908102919091018101516001600160a01b031682528181019290925260409081016000908120805460ff19169415159490941790935588835290829052902060030154611c6f906001611cd5565b50336001600160a01b031686600080516020611e78833981519152878481518110611c9657fe5b602002602001015160405180826001600160a01b03166001600160a01b0316815260200191505060405180910390a3600101611b76565b505050505050565b600082820183811015611d2f576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600082821115611d8d576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe436861696e20494420646f6573206e6f742065786973740000000000000000004d6178696d756d20313020417574686f72697365642061646472657373657320616c6c6f7765644e6577206f776e65722072657175697265642c20616e642073686f756c64206265206e6f6e2d7a65726f2061646472657373496e697469616c20417574686f7269736564206164647265737365732072657175697265644d617820313020617574686f72697365642061646472657373657320616c6c6f7765644d6178207375626d697373696f6e2031302061646472657373657320616c6c6f776564fc23324c4d8b370287cc3e6d2db8892f41de294bd67a7e3c8ecb88742081aa2d6174206c65617374206f6e652072656d61696e696e6720617574686f72697365642061646472657373207265717569726564a265627a7a72305820d1b81f5772614aa49e7cab633fbf3f0c7536308d25835f840a9e1b8f338cb6bb64736f6c634300050a0032"

// DeployWRKChainRoot deploys a new Ethereum contract, binding an instance of WRKChainRoot to it.
func DeployWRKChainRoot(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WRKChainRoot, error) {
	parsed, err := abi.JSON(strings.NewReader(WRKChainRootABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WRKChainRootBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WRKChainRoot{WRKChainRootCaller: WRKChainRootCaller{contract: contract}, WRKChainRootTransactor: WRKChainRootTransactor{contract: contract}, WRKChainRootFilterer: WRKChainRootFilterer{contract: contract}}, nil
}

// WRKChainRoot is an auto generated Go binding around an Ethereum contract.
type WRKChainRoot struct {
	WRKChainRootCaller     // Read-only binding to the contract
	WRKChainRootTransactor // Write-only binding to the contract
	WRKChainRootFilterer   // Log filterer for contract events
}

// WRKChainRootCaller is an auto generated read-only Go binding around an Ethereum contract.
type WRKChainRootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WRKChainRootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WRKChainRootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WRKChainRootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WRKChainRootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WRKChainRootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WRKChainRootSession struct {
	Contract     *WRKChainRoot     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WRKChainRootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WRKChainRootCallerSession struct {
	Contract *WRKChainRootCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WRKChainRootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WRKChainRootTransactorSession struct {
	Contract     *WRKChainRootTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WRKChainRootRaw is an auto generated low-level Go binding around an Ethereum contract.
type WRKChainRootRaw struct {
	Contract *WRKChainRoot // Generic contract binding to access the raw methods on
}

// WRKChainRootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WRKChainRootCallerRaw struct {
	Contract *WRKChainRootCaller // Generic read-only contract binding to access the raw methods on
}

// WRKChainRootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WRKChainRootTransactorRaw struct {
	Contract *WRKChainRootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWRKChainRoot creates a new instance of WRKChainRoot, bound to a specific deployed contract.
func NewWRKChainRoot(address common.Address, backend bind.ContractBackend) (*WRKChainRoot, error) {
	contract, err := bindWRKChainRoot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WRKChainRoot{WRKChainRootCaller: WRKChainRootCaller{contract: contract}, WRKChainRootTransactor: WRKChainRootTransactor{contract: contract}, WRKChainRootFilterer: WRKChainRootFilterer{contract: contract}}, nil
}

// NewWRKChainRootCaller creates a new read-only instance of WRKChainRoot, bound to a specific deployed contract.
func NewWRKChainRootCaller(address common.Address, caller bind.ContractCaller) (*WRKChainRootCaller, error) {
	contract, err := bindWRKChainRoot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootCaller{contract: contract}, nil
}

// NewWRKChainRootTransactor creates a new write-only instance of WRKChainRoot, bound to a specific deployed contract.
func NewWRKChainRootTransactor(address common.Address, transactor bind.ContractTransactor) (*WRKChainRootTransactor, error) {
	contract, err := bindWRKChainRoot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootTransactor{contract: contract}, nil
}

// NewWRKChainRootFilterer creates a new log filterer instance of WRKChainRoot, bound to a specific deployed contract.
func NewWRKChainRootFilterer(address common.Address, filterer bind.ContractFilterer) (*WRKChainRootFilterer, error) {
	contract, err := bindWRKChainRoot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootFilterer{contract: contract}, nil
}

// bindWRKChainRoot binds a generic wrapper to an already deployed contract.
func bindWRKChainRoot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WRKChainRootABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WRKChainRoot *WRKChainRootRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WRKChainRoot.Contract.WRKChainRootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WRKChainRoot *WRKChainRootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.WRKChainRootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WRKChainRoot *WRKChainRootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.WRKChainRootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WRKChainRoot *WRKChainRootCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WRKChainRoot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WRKChainRoot *WRKChainRootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WRKChainRoot *WRKChainRootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.contract.Transact(opts, method, params...)
}

// ChainExists is a free data retrieval call binding the contract method 0xca804ccd.
//
// Solidity: function chainExists(bytes32 _chainId) constant returns(bool chainExists_)
func (_WRKChainRoot *WRKChainRootCaller) ChainExists(opts *bind.CallOpts, _chainId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "chainExists", _chainId)
	return *ret0, err
}

// ChainExists is a free data retrieval call binding the contract method 0xca804ccd.
//
// Solidity: function chainExists(bytes32 _chainId) constant returns(bool chainExists_)
func (_WRKChainRoot *WRKChainRootSession) ChainExists(_chainId [32]byte) (bool, error) {
	return _WRKChainRoot.Contract.ChainExists(&_WRKChainRoot.CallOpts, _chainId)
}

// ChainExists is a free data retrieval call binding the contract method 0xca804ccd.
//
// Solidity: function chainExists(bytes32 _chainId) constant returns(bool chainExists_)
func (_WRKChainRoot *WRKChainRootCallerSession) ChainExists(_chainId [32]byte) (bool, error) {
	return _WRKChainRoot.Contract.ChainExists(&_WRKChainRoot.CallOpts, _chainId)
}

// GetGenesis is a free data retrieval call binding the contract method 0x278f1daf.
//
// Solidity: function getGenesis(bytes32 _chainId) constant returns(bytes32 genesisHash_)
func (_WRKChainRoot *WRKChainRootCaller) GetGenesis(opts *bind.CallOpts, _chainId [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "getGenesis", _chainId)
	return *ret0, err
}

// GetGenesis is a free data retrieval call binding the contract method 0x278f1daf.
//
// Solidity: function getGenesis(bytes32 _chainId) constant returns(bytes32 genesisHash_)
func (_WRKChainRoot *WRKChainRootSession) GetGenesis(_chainId [32]byte) ([32]byte, error) {
	return _WRKChainRoot.Contract.GetGenesis(&_WRKChainRoot.CallOpts, _chainId)
}

// GetGenesis is a free data retrieval call binding the contract method 0x278f1daf.
//
// Solidity: function getGenesis(bytes32 _chainId) constant returns(bytes32 genesisHash_)
func (_WRKChainRoot *WRKChainRootCallerSession) GetGenesis(_chainId [32]byte) ([32]byte, error) {
	return _WRKChainRoot.Contract.GetGenesis(&_WRKChainRoot.CallOpts, _chainId)
}

// GetLastBlock is a free data retrieval call binding the contract method 0x8b3e5e06.
//
// Solidity: function getLastBlock(bytes32 _chainId) constant returns(bytes32 lastHash_, bytes32 lastParentHash_, bytes32 lastReceiptRoot_, bytes32 lastTxRoot_, bytes32 lastStateRoot_, uint256 lastSubmissionTime_, address lastSubmittedBy_)
func (_WRKChainRoot *WRKChainRootCaller) GetLastBlock(opts *bind.CallOpts, _chainId [32]byte) (struct {
	LastHash           [32]byte
	LastParentHash     [32]byte
	LastReceiptRoot    [32]byte
	LastTxRoot         [32]byte
	LastStateRoot      [32]byte
	LastSubmissionTime *big.Int
	LastSubmittedBy    common.Address
}, error) {
	ret := new(struct {
		LastHash           [32]byte
		LastParentHash     [32]byte
		LastReceiptRoot    [32]byte
		LastTxRoot         [32]byte
		LastStateRoot      [32]byte
		LastSubmissionTime *big.Int
		LastSubmittedBy    common.Address
	})
	out := ret
	err := _WRKChainRoot.contract.Call(opts, out, "getLastBlock", _chainId)
	return *ret, err
}

// GetLastBlock is a free data retrieval call binding the contract method 0x8b3e5e06.
//
// Solidity: function getLastBlock(bytes32 _chainId) constant returns(bytes32 lastHash_, bytes32 lastParentHash_, bytes32 lastReceiptRoot_, bytes32 lastTxRoot_, bytes32 lastStateRoot_, uint256 lastSubmissionTime_, address lastSubmittedBy_)
func (_WRKChainRoot *WRKChainRootSession) GetLastBlock(_chainId [32]byte) (struct {
	LastHash           [32]byte
	LastParentHash     [32]byte
	LastReceiptRoot    [32]byte
	LastTxRoot         [32]byte
	LastStateRoot      [32]byte
	LastSubmissionTime *big.Int
	LastSubmittedBy    common.Address
}, error) {
	return _WRKChainRoot.Contract.GetLastBlock(&_WRKChainRoot.CallOpts, _chainId)
}

// GetLastBlock is a free data retrieval call binding the contract method 0x8b3e5e06.
//
// Solidity: function getLastBlock(bytes32 _chainId) constant returns(bytes32 lastHash_, bytes32 lastParentHash_, bytes32 lastReceiptRoot_, bytes32 lastTxRoot_, bytes32 lastStateRoot_, uint256 lastSubmissionTime_, address lastSubmittedBy_)
func (_WRKChainRoot *WRKChainRootCallerSession) GetLastBlock(_chainId [32]byte) (struct {
	LastHash           [32]byte
	LastParentHash     [32]byte
	LastReceiptRoot    [32]byte
	LastTxRoot         [32]byte
	LastStateRoot      [32]byte
	LastSubmissionTime *big.Int
	LastSubmittedBy    common.Address
}, error) {
	return _WRKChainRoot.Contract.GetLastBlock(&_WRKChainRoot.CallOpts, _chainId)
}

// GetLastRecordedBlockNum is a free data retrieval call binding the contract method 0xc4d8561e.
//
// Solidity: function getLastRecordedBlockNum(bytes32 _chainId) constant returns(uint256 lastBlock_)
func (_WRKChainRoot *WRKChainRootCaller) GetLastRecordedBlockNum(opts *bind.CallOpts, _chainId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "getLastRecordedBlockNum", _chainId)
	return *ret0, err
}

// GetLastRecordedBlockNum is a free data retrieval call binding the contract method 0xc4d8561e.
//
// Solidity: function getLastRecordedBlockNum(bytes32 _chainId) constant returns(uint256 lastBlock_)
func (_WRKChainRoot *WRKChainRootSession) GetLastRecordedBlockNum(_chainId [32]byte) (*big.Int, error) {
	return _WRKChainRoot.Contract.GetLastRecordedBlockNum(&_WRKChainRoot.CallOpts, _chainId)
}

// GetLastRecordedBlockNum is a free data retrieval call binding the contract method 0xc4d8561e.
//
// Solidity: function getLastRecordedBlockNum(bytes32 _chainId) constant returns(uint256 lastBlock_)
func (_WRKChainRoot *WRKChainRootCallerSession) GetLastRecordedBlockNum(_chainId [32]byte) (*big.Int, error) {
	return _WRKChainRoot.Contract.GetLastRecordedBlockNum(&_WRKChainRoot.CallOpts, _chainId)
}

// GetNumBlocksSubmitted is a free data retrieval call binding the contract method 0x0b2b1087.
//
// Solidity: function getNumBlocksSubmitted(bytes32 _chainId) constant returns(uint256 numBlocksSubmitted_)
func (_WRKChainRoot *WRKChainRootCaller) GetNumBlocksSubmitted(opts *bind.CallOpts, _chainId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "getNumBlocksSubmitted", _chainId)
	return *ret0, err
}

// GetNumBlocksSubmitted is a free data retrieval call binding the contract method 0x0b2b1087.
//
// Solidity: function getNumBlocksSubmitted(bytes32 _chainId) constant returns(uint256 numBlocksSubmitted_)
func (_WRKChainRoot *WRKChainRootSession) GetNumBlocksSubmitted(_chainId [32]byte) (*big.Int, error) {
	return _WRKChainRoot.Contract.GetNumBlocksSubmitted(&_WRKChainRoot.CallOpts, _chainId)
}

// GetNumBlocksSubmitted is a free data retrieval call binding the contract method 0x0b2b1087.
//
// Solidity: function getNumBlocksSubmitted(bytes32 _chainId) constant returns(uint256 numBlocksSubmitted_)
func (_WRKChainRoot *WRKChainRootCallerSession) GetNumBlocksSubmitted(_chainId [32]byte) (*big.Int, error) {
	return _WRKChainRoot.Contract.GetNumBlocksSubmitted(&_WRKChainRoot.CallOpts, _chainId)
}

// IsAuthorisedAddress is a free data retrieval call binding the contract method 0x62269f7a.
//
// Solidity: function isAuthorisedAddress(bytes32 _chainId, address _address) constant returns(bool isAuthorised_)
func (_WRKChainRoot *WRKChainRootCaller) IsAuthorisedAddress(opts *bind.CallOpts, _chainId [32]byte, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "isAuthorisedAddress", _chainId, _address)
	return *ret0, err
}

// IsAuthorisedAddress is a free data retrieval call binding the contract method 0x62269f7a.
//
// Solidity: function isAuthorisedAddress(bytes32 _chainId, address _address) constant returns(bool isAuthorised_)
func (_WRKChainRoot *WRKChainRootSession) IsAuthorisedAddress(_chainId [32]byte, _address common.Address) (bool, error) {
	return _WRKChainRoot.Contract.IsAuthorisedAddress(&_WRKChainRoot.CallOpts, _chainId, _address)
}

// IsAuthorisedAddress is a free data retrieval call binding the contract method 0x62269f7a.
//
// Solidity: function isAuthorisedAddress(bytes32 _chainId, address _address) constant returns(bool isAuthorised_)
func (_WRKChainRoot *WRKChainRootCallerSession) IsAuthorisedAddress(_chainId [32]byte, _address common.Address) (bool, error) {
	return _WRKChainRoot.Contract.IsAuthorisedAddress(&_WRKChainRoot.CallOpts, _chainId, _address)
}

// WrkchainList is a free data retrieval call binding the contract method 0x8c75c5d1.
//
// Solidity: function wrkchainList(bytes32 ) constant returns(bytes32 genesisHash, bool isWrkchain, uint256 numAuthAddresses, address owner, uint256 numBlocksSubmitted, uint256 numHistoricalBlocksSubmitted, uint256 lastBlock, bytes32 lastHash, bytes32 lastParentHash, bytes32 lastReceiptRoot, bytes32 lastTxRoot, bytes32 lastStateRoot, uint256 lastSubmissionTime, address lastSubmittedBy)
func (_WRKChainRoot *WRKChainRootCaller) WrkchainList(opts *bind.CallOpts, arg0 [32]byte) (struct {
	GenesisHash                  [32]byte
	IsWrkchain                   bool
	NumAuthAddresses             *big.Int
	Owner                        common.Address
	NumBlocksSubmitted           *big.Int
	NumHistoricalBlocksSubmitted *big.Int
	LastBlock                    *big.Int
	LastHash                     [32]byte
	LastParentHash               [32]byte
	LastReceiptRoot              [32]byte
	LastTxRoot                   [32]byte
	LastStateRoot                [32]byte
	LastSubmissionTime           *big.Int
	LastSubmittedBy              common.Address
}, error) {
	ret := new(struct {
		GenesisHash                  [32]byte
		IsWrkchain                   bool
		NumAuthAddresses             *big.Int
		Owner                        common.Address
		NumBlocksSubmitted           *big.Int
		NumHistoricalBlocksSubmitted *big.Int
		LastBlock                    *big.Int
		LastHash                     [32]byte
		LastParentHash               [32]byte
		LastReceiptRoot              [32]byte
		LastTxRoot                   [32]byte
		LastStateRoot                [32]byte
		LastSubmissionTime           *big.Int
		LastSubmittedBy              common.Address
	})
	out := ret
	err := _WRKChainRoot.contract.Call(opts, out, "wrkchainList", arg0)
	return *ret, err
}

// WrkchainList is a free data retrieval call binding the contract method 0x8c75c5d1.
//
// Solidity: function wrkchainList(bytes32 ) constant returns(bytes32 genesisHash, bool isWrkchain, uint256 numAuthAddresses, address owner, uint256 numBlocksSubmitted, uint256 numHistoricalBlocksSubmitted, uint256 lastBlock, bytes32 lastHash, bytes32 lastParentHash, bytes32 lastReceiptRoot, bytes32 lastTxRoot, bytes32 lastStateRoot, uint256 lastSubmissionTime, address lastSubmittedBy)
func (_WRKChainRoot *WRKChainRootSession) WrkchainList(arg0 [32]byte) (struct {
	GenesisHash                  [32]byte
	IsWrkchain                   bool
	NumAuthAddresses             *big.Int
	Owner                        common.Address
	NumBlocksSubmitted           *big.Int
	NumHistoricalBlocksSubmitted *big.Int
	LastBlock                    *big.Int
	LastHash                     [32]byte
	LastParentHash               [32]byte
	LastReceiptRoot              [32]byte
	LastTxRoot                   [32]byte
	LastStateRoot                [32]byte
	LastSubmissionTime           *big.Int
	LastSubmittedBy              common.Address
}, error) {
	return _WRKChainRoot.Contract.WrkchainList(&_WRKChainRoot.CallOpts, arg0)
}

// WrkchainList is a free data retrieval call binding the contract method 0x8c75c5d1.
//
// Solidity: function wrkchainList(bytes32 ) constant returns(bytes32 genesisHash, bool isWrkchain, uint256 numAuthAddresses, address owner, uint256 numBlocksSubmitted, uint256 numHistoricalBlocksSubmitted, uint256 lastBlock, bytes32 lastHash, bytes32 lastParentHash, bytes32 lastReceiptRoot, bytes32 lastTxRoot, bytes32 lastStateRoot, uint256 lastSubmissionTime, address lastSubmittedBy)
func (_WRKChainRoot *WRKChainRootCallerSession) WrkchainList(arg0 [32]byte) (struct {
	GenesisHash                  [32]byte
	IsWrkchain                   bool
	NumAuthAddresses             *big.Int
	Owner                        common.Address
	NumBlocksSubmitted           *big.Int
	NumHistoricalBlocksSubmitted *big.Int
	LastBlock                    *big.Int
	LastHash                     [32]byte
	LastParentHash               [32]byte
	LastReceiptRoot              [32]byte
	LastTxRoot                   [32]byte
	LastStateRoot                [32]byte
	LastSubmissionTime           *big.Int
	LastSubmittedBy              common.Address
}, error) {
	return _WRKChainRoot.Contract.WrkchainList(&_WRKChainRoot.CallOpts, arg0)
}

// AddAuthAddresses is a paid mutator transaction binding the contract method 0xd4202677.
//
// Solidity: function addAuthAddresses(bytes32 _chainId, address[] _newAuthAddresses) returns()
func (_WRKChainRoot *WRKChainRootTransactor) AddAuthAddresses(opts *bind.TransactOpts, _chainId [32]byte, _newAuthAddresses []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "addAuthAddresses", _chainId, _newAuthAddresses)
}

// AddAuthAddresses is a paid mutator transaction binding the contract method 0xd4202677.
//
// Solidity: function addAuthAddresses(bytes32 _chainId, address[] _newAuthAddresses) returns()
func (_WRKChainRoot *WRKChainRootSession) AddAuthAddresses(_chainId [32]byte, _newAuthAddresses []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.AddAuthAddresses(&_WRKChainRoot.TransactOpts, _chainId, _newAuthAddresses)
}

// AddAuthAddresses is a paid mutator transaction binding the contract method 0xd4202677.
//
// Solidity: function addAuthAddresses(bytes32 _chainId, address[] _newAuthAddresses) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) AddAuthAddresses(_chainId [32]byte, _newAuthAddresses []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.AddAuthAddresses(&_WRKChainRoot.TransactOpts, _chainId, _newAuthAddresses)
}

// ChangeWrkchainOwner is a paid mutator transaction binding the contract method 0xa51af476.
//
// Solidity: function changeWrkchainOwner(bytes32 _chainId, address _newOwner) returns()
func (_WRKChainRoot *WRKChainRootTransactor) ChangeWrkchainOwner(opts *bind.TransactOpts, _chainId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "changeWrkchainOwner", _chainId, _newOwner)
}

// ChangeWrkchainOwner is a paid mutator transaction binding the contract method 0xa51af476.
//
// Solidity: function changeWrkchainOwner(bytes32 _chainId, address _newOwner) returns()
func (_WRKChainRoot *WRKChainRootSession) ChangeWrkchainOwner(_chainId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.ChangeWrkchainOwner(&_WRKChainRoot.TransactOpts, _chainId, _newOwner)
}

// ChangeWrkchainOwner is a paid mutator transaction binding the contract method 0xa51af476.
//
// Solidity: function changeWrkchainOwner(bytes32 _chainId, address _newOwner) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) ChangeWrkchainOwner(_chainId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.ChangeWrkchainOwner(&_WRKChainRoot.TransactOpts, _chainId, _newOwner)
}

// RecordHeader is a paid mutator transaction binding the contract method 0x01242bc5.
//
// Solidity: function recordHeader(bytes32 _chainId, uint256 _height, bytes32 _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot) returns()
func (_WRKChainRoot *WRKChainRootTransactor) RecordHeader(opts *bind.TransactOpts, _chainId [32]byte, _height *big.Int, _hash [32]byte, _parentHash [32]byte, _receiptRoot [32]byte, _txRoot [32]byte, _stateRoot [32]byte) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "recordHeader", _chainId, _height, _hash, _parentHash, _receiptRoot, _txRoot, _stateRoot)
}

// RecordHeader is a paid mutator transaction binding the contract method 0x01242bc5.
//
// Solidity: function recordHeader(bytes32 _chainId, uint256 _height, bytes32 _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot) returns()
func (_WRKChainRoot *WRKChainRootSession) RecordHeader(_chainId [32]byte, _height *big.Int, _hash [32]byte, _parentHash [32]byte, _receiptRoot [32]byte, _txRoot [32]byte, _stateRoot [32]byte) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RecordHeader(&_WRKChainRoot.TransactOpts, _chainId, _height, _hash, _parentHash, _receiptRoot, _txRoot, _stateRoot)
}

// RecordHeader is a paid mutator transaction binding the contract method 0x01242bc5.
//
// Solidity: function recordHeader(bytes32 _chainId, uint256 _height, bytes32 _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) RecordHeader(_chainId [32]byte, _height *big.Int, _hash [32]byte, _parentHash [32]byte, _receiptRoot [32]byte, _txRoot [32]byte, _stateRoot [32]byte) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RecordHeader(&_WRKChainRoot.TransactOpts, _chainId, _height, _hash, _parentHash, _receiptRoot, _txRoot, _stateRoot)
}

// RegisterWrkChain is a paid mutator transaction binding the contract method 0x004d6c96.
//
// Solidity: function registerWrkChain(bytes32 _chainId, address[] _authAddresses, bytes32 _genesisHash) returns()
func (_WRKChainRoot *WRKChainRootTransactor) RegisterWrkChain(opts *bind.TransactOpts, _chainId [32]byte, _authAddresses []common.Address, _genesisHash [32]byte) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "registerWrkChain", _chainId, _authAddresses, _genesisHash)
}

// RegisterWrkChain is a paid mutator transaction binding the contract method 0x004d6c96.
//
// Solidity: function registerWrkChain(bytes32 _chainId, address[] _authAddresses, bytes32 _genesisHash) returns()
func (_WRKChainRoot *WRKChainRootSession) RegisterWrkChain(_chainId [32]byte, _authAddresses []common.Address, _genesisHash [32]byte) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RegisterWrkChain(&_WRKChainRoot.TransactOpts, _chainId, _authAddresses, _genesisHash)
}

// RegisterWrkChain is a paid mutator transaction binding the contract method 0x004d6c96.
//
// Solidity: function registerWrkChain(bytes32 _chainId, address[] _authAddresses, bytes32 _genesisHash) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) RegisterWrkChain(_chainId [32]byte, _authAddresses []common.Address, _genesisHash [32]byte) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RegisterWrkChain(&_WRKChainRoot.TransactOpts, _chainId, _authAddresses, _genesisHash)
}

// RemoveAuthAddresses is a paid mutator transaction binding the contract method 0x6c5664ac.
//
// Solidity: function removeAuthAddresses(bytes32 _chainId, address[] _addressesToRemove) returns()
func (_WRKChainRoot *WRKChainRootTransactor) RemoveAuthAddresses(opts *bind.TransactOpts, _chainId [32]byte, _addressesToRemove []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "removeAuthAddresses", _chainId, _addressesToRemove)
}

// RemoveAuthAddresses is a paid mutator transaction binding the contract method 0x6c5664ac.
//
// Solidity: function removeAuthAddresses(bytes32 _chainId, address[] _addressesToRemove) returns()
func (_WRKChainRoot *WRKChainRootSession) RemoveAuthAddresses(_chainId [32]byte, _addressesToRemove []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RemoveAuthAddresses(&_WRKChainRoot.TransactOpts, _chainId, _addressesToRemove)
}

// RemoveAuthAddresses is a paid mutator transaction binding the contract method 0x6c5664ac.
//
// Solidity: function removeAuthAddresses(bytes32 _chainId, address[] _addressesToRemove) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) RemoveAuthAddresses(_chainId [32]byte, _addressesToRemove []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RemoveAuthAddresses(&_WRKChainRoot.TransactOpts, _chainId, _addressesToRemove)
}

// WRKChainRootAuthoriseNewAddressIterator is returned from FilterAuthoriseNewAddress and is used to iterate over the raw logs and unpacked data for AuthoriseNewAddress events raised by the WRKChainRoot contract.
type WRKChainRootAuthoriseNewAddressIterator struct {
	Event *WRKChainRootAuthoriseNewAddress // Event containing the contract specifics and raw log

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
func (it *WRKChainRootAuthoriseNewAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootAuthoriseNewAddress)
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
		it.Event = new(WRKChainRootAuthoriseNewAddress)
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
func (it *WRKChainRootAuthoriseNewAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootAuthoriseNewAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootAuthoriseNewAddress represents a AuthoriseNewAddress event raised by the WRKChainRoot contract.
type WRKChainRootAuthoriseNewAddress struct {
	ChainId      [32]byte
	AuthorisedBy common.Address
	Address      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuthoriseNewAddress is a free log retrieval operation binding the contract event 0xfc23324c4d8b370287cc3e6d2db8892f41de294bd67a7e3c8ecb88742081aa2d.
//
// Solidity: event AuthoriseNewAddress(bytes32 indexed _chainId, address indexed _authorisedBy, address _address)
func (_WRKChainRoot *WRKChainRootFilterer) FilterAuthoriseNewAddress(opts *bind.FilterOpts, _chainId [][32]byte, _authorisedBy []common.Address) (*WRKChainRootAuthoriseNewAddressIterator, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}
	var _authorisedByRule []interface{}
	for _, _authorisedByItem := range _authorisedBy {
		_authorisedByRule = append(_authorisedByRule, _authorisedByItem)
	}

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "AuthoriseNewAddress", _chainIdRule, _authorisedByRule)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootAuthoriseNewAddressIterator{contract: _WRKChainRoot.contract, event: "AuthoriseNewAddress", logs: logs, sub: sub}, nil
}

// WatchAuthoriseNewAddress is a free log subscription operation binding the contract event 0xfc23324c4d8b370287cc3e6d2db8892f41de294bd67a7e3c8ecb88742081aa2d.
//
// Solidity: event AuthoriseNewAddress(bytes32 indexed _chainId, address indexed _authorisedBy, address _address)
func (_WRKChainRoot *WRKChainRootFilterer) WatchAuthoriseNewAddress(opts *bind.WatchOpts, sink chan<- *WRKChainRootAuthoriseNewAddress, _chainId [][32]byte, _authorisedBy []common.Address) (event.Subscription, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}
	var _authorisedByRule []interface{}
	for _, _authorisedByItem := range _authorisedBy {
		_authorisedByRule = append(_authorisedByRule, _authorisedByItem)
	}

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "AuthoriseNewAddress", _chainIdRule, _authorisedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootAuthoriseNewAddress)
				if err := _WRKChainRoot.contract.UnpackLog(event, "AuthoriseNewAddress", log); err != nil {
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

// ParseAuthoriseNewAddress is a log parse operation binding the contract event 0xfc23324c4d8b370287cc3e6d2db8892f41de294bd67a7e3c8ecb88742081aa2d.
//
// Solidity: event AuthoriseNewAddress(bytes32 indexed _chainId, address indexed _authorisedBy, address _address)
func (_WRKChainRoot *WRKChainRootFilterer) ParseAuthoriseNewAddress(log types.Log) (*WRKChainRootAuthoriseNewAddress, error) {
	event := new(WRKChainRootAuthoriseNewAddress)
	if err := _WRKChainRoot.contract.UnpackLog(event, "AuthoriseNewAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WRKChainRootLogFallbackFunctionCalledIterator is returned from FilterLogFallbackFunctionCalled and is used to iterate over the raw logs and unpacked data for LogFallbackFunctionCalled events raised by the WRKChainRoot contract.
type WRKChainRootLogFallbackFunctionCalledIterator struct {
	Event *WRKChainRootLogFallbackFunctionCalled // Event containing the contract specifics and raw log

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
func (it *WRKChainRootLogFallbackFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootLogFallbackFunctionCalled)
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
		it.Event = new(WRKChainRootLogFallbackFunctionCalled)
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
func (it *WRKChainRootLogFallbackFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootLogFallbackFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootLogFallbackFunctionCalled represents a LogFallbackFunctionCalled event raised by the WRKChainRoot contract.
type WRKChainRootLogFallbackFunctionCalled struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogFallbackFunctionCalled is a free log retrieval operation binding the contract event 0xb1db2589d999730e8faec7dccf62c4d28a7f6ef1095b46c9a744dcd46d87130e.
//
// Solidity: event LogFallbackFunctionCalled(address indexed _from, uint256 _amount)
func (_WRKChainRoot *WRKChainRootFilterer) FilterLogFallbackFunctionCalled(opts *bind.FilterOpts, _from []common.Address) (*WRKChainRootLogFallbackFunctionCalledIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "LogFallbackFunctionCalled", _fromRule)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootLogFallbackFunctionCalledIterator{contract: _WRKChainRoot.contract, event: "LogFallbackFunctionCalled", logs: logs, sub: sub}, nil
}

// WatchLogFallbackFunctionCalled is a free log subscription operation binding the contract event 0xb1db2589d999730e8faec7dccf62c4d28a7f6ef1095b46c9a744dcd46d87130e.
//
// Solidity: event LogFallbackFunctionCalled(address indexed _from, uint256 _amount)
func (_WRKChainRoot *WRKChainRootFilterer) WatchLogFallbackFunctionCalled(opts *bind.WatchOpts, sink chan<- *WRKChainRootLogFallbackFunctionCalled, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "LogFallbackFunctionCalled", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootLogFallbackFunctionCalled)
				if err := _WRKChainRoot.contract.UnpackLog(event, "LogFallbackFunctionCalled", log); err != nil {
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

// ParseLogFallbackFunctionCalled is a log parse operation binding the contract event 0xb1db2589d999730e8faec7dccf62c4d28a7f6ef1095b46c9a744dcd46d87130e.
//
// Solidity: event LogFallbackFunctionCalled(address indexed _from, uint256 _amount)
func (_WRKChainRoot *WRKChainRootFilterer) ParseLogFallbackFunctionCalled(log types.Log) (*WRKChainRootLogFallbackFunctionCalled, error) {
	event := new(WRKChainRootLogFallbackFunctionCalled)
	if err := _WRKChainRoot.contract.UnpackLog(event, "LogFallbackFunctionCalled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WRKChainRootRecordHeaderIterator is returned from FilterRecordHeader and is used to iterate over the raw logs and unpacked data for RecordHeader events raised by the WRKChainRoot contract.
type WRKChainRootRecordHeaderIterator struct {
	Event *WRKChainRootRecordHeader // Event containing the contract specifics and raw log

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
func (it *WRKChainRootRecordHeaderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootRecordHeader)
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
		it.Event = new(WRKChainRootRecordHeader)
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
func (it *WRKChainRootRecordHeaderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootRecordHeaderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootRecordHeader represents a RecordHeader event raised by the WRKChainRoot contract.
type WRKChainRootRecordHeader struct {
	ChainId     [32]byte
	Height      *big.Int
	Hash        [32]byte
	ParentHash  [32]byte
	ReceiptRoot [32]byte
	TxRoot      [32]byte
	StateRoot   [32]byte
	Timestamp   *big.Int
	SubmittedBy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRecordHeader is a free log retrieval operation binding the contract event 0x22f683acdb61cbd91a27907cbd5a4c5a12268ab76c79b5af59511aa03f0fe2af.
//
// Solidity: event RecordHeader(bytes32 indexed _chainId, uint256 indexed _height, bytes32 indexed _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, uint256 _timestamp, address _submittedBy)
func (_WRKChainRoot *WRKChainRootFilterer) FilterRecordHeader(opts *bind.FilterOpts, _chainId [][32]byte, _height []*big.Int, _hash [][32]byte) (*WRKChainRootRecordHeaderIterator, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}
	var _heightRule []interface{}
	for _, _heightItem := range _height {
		_heightRule = append(_heightRule, _heightItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "RecordHeader", _chainIdRule, _heightRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootRecordHeaderIterator{contract: _WRKChainRoot.contract, event: "RecordHeader", logs: logs, sub: sub}, nil
}

// WatchRecordHeader is a free log subscription operation binding the contract event 0x22f683acdb61cbd91a27907cbd5a4c5a12268ab76c79b5af59511aa03f0fe2af.
//
// Solidity: event RecordHeader(bytes32 indexed _chainId, uint256 indexed _height, bytes32 indexed _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, uint256 _timestamp, address _submittedBy)
func (_WRKChainRoot *WRKChainRootFilterer) WatchRecordHeader(opts *bind.WatchOpts, sink chan<- *WRKChainRootRecordHeader, _chainId [][32]byte, _height []*big.Int, _hash [][32]byte) (event.Subscription, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}
	var _heightRule []interface{}
	for _, _heightItem := range _height {
		_heightRule = append(_heightRule, _heightItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "RecordHeader", _chainIdRule, _heightRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootRecordHeader)
				if err := _WRKChainRoot.contract.UnpackLog(event, "RecordHeader", log); err != nil {
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

// ParseRecordHeader is a log parse operation binding the contract event 0x22f683acdb61cbd91a27907cbd5a4c5a12268ab76c79b5af59511aa03f0fe2af.
//
// Solidity: event RecordHeader(bytes32 indexed _chainId, uint256 indexed _height, bytes32 indexed _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, uint256 _timestamp, address _submittedBy)
func (_WRKChainRoot *WRKChainRootFilterer) ParseRecordHeader(log types.Log) (*WRKChainRootRecordHeader, error) {
	event := new(WRKChainRootRecordHeader)
	if err := _WRKChainRoot.contract.UnpackLog(event, "RecordHeader", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WRKChainRootRecordHeaderHistoricalIterator is returned from FilterRecordHeaderHistorical and is used to iterate over the raw logs and unpacked data for RecordHeaderHistorical events raised by the WRKChainRoot contract.
type WRKChainRootRecordHeaderHistoricalIterator struct {
	Event *WRKChainRootRecordHeaderHistorical // Event containing the contract specifics and raw log

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
func (it *WRKChainRootRecordHeaderHistoricalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootRecordHeaderHistorical)
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
		it.Event = new(WRKChainRootRecordHeaderHistorical)
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
func (it *WRKChainRootRecordHeaderHistoricalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootRecordHeaderHistoricalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootRecordHeaderHistorical represents a RecordHeaderHistorical event raised by the WRKChainRoot contract.
type WRKChainRootRecordHeaderHistorical struct {
	ChainId     [32]byte
	Height      *big.Int
	Hash        [32]byte
	ParentHash  [32]byte
	ReceiptRoot [32]byte
	TxRoot      [32]byte
	StateRoot   [32]byte
	Timestamp   *big.Int
	SubmittedBy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRecordHeaderHistorical is a free log retrieval operation binding the contract event 0x599336707b966c6560001ed107c04f461a648d8c67977e6c7fbd346655890eb2.
//
// Solidity: event RecordHeaderHistorical(bytes32 indexed _chainId, uint256 indexed _height, bytes32 indexed _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, uint256 _timestamp, address _submittedBy)
func (_WRKChainRoot *WRKChainRootFilterer) FilterRecordHeaderHistorical(opts *bind.FilterOpts, _chainId [][32]byte, _height []*big.Int, _hash [][32]byte) (*WRKChainRootRecordHeaderHistoricalIterator, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}
	var _heightRule []interface{}
	for _, _heightItem := range _height {
		_heightRule = append(_heightRule, _heightItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "RecordHeaderHistorical", _chainIdRule, _heightRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootRecordHeaderHistoricalIterator{contract: _WRKChainRoot.contract, event: "RecordHeaderHistorical", logs: logs, sub: sub}, nil
}

// WatchRecordHeaderHistorical is a free log subscription operation binding the contract event 0x599336707b966c6560001ed107c04f461a648d8c67977e6c7fbd346655890eb2.
//
// Solidity: event RecordHeaderHistorical(bytes32 indexed _chainId, uint256 indexed _height, bytes32 indexed _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, uint256 _timestamp, address _submittedBy)
func (_WRKChainRoot *WRKChainRootFilterer) WatchRecordHeaderHistorical(opts *bind.WatchOpts, sink chan<- *WRKChainRootRecordHeaderHistorical, _chainId [][32]byte, _height []*big.Int, _hash [][32]byte) (event.Subscription, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}
	var _heightRule []interface{}
	for _, _heightItem := range _height {
		_heightRule = append(_heightRule, _heightItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "RecordHeaderHistorical", _chainIdRule, _heightRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootRecordHeaderHistorical)
				if err := _WRKChainRoot.contract.UnpackLog(event, "RecordHeaderHistorical", log); err != nil {
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

// ParseRecordHeaderHistorical is a log parse operation binding the contract event 0x599336707b966c6560001ed107c04f461a648d8c67977e6c7fbd346655890eb2.
//
// Solidity: event RecordHeaderHistorical(bytes32 indexed _chainId, uint256 indexed _height, bytes32 indexed _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, uint256 _timestamp, address _submittedBy)
func (_WRKChainRoot *WRKChainRootFilterer) ParseRecordHeaderHistorical(log types.Log) (*WRKChainRootRecordHeaderHistorical, error) {
	event := new(WRKChainRootRecordHeaderHistorical)
	if err := _WRKChainRoot.contract.UnpackLog(event, "RecordHeaderHistorical", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WRKChainRootRegisterWrkChainIterator is returned from FilterRegisterWrkChain and is used to iterate over the raw logs and unpacked data for RegisterWrkChain events raised by the WRKChainRoot contract.
type WRKChainRootRegisterWrkChainIterator struct {
	Event *WRKChainRootRegisterWrkChain // Event containing the contract specifics and raw log

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
func (it *WRKChainRootRegisterWrkChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootRegisterWrkChain)
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
		it.Event = new(WRKChainRootRegisterWrkChain)
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
func (it *WRKChainRootRegisterWrkChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootRegisterWrkChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootRegisterWrkChain represents a RegisterWrkChain event raised by the WRKChainRoot contract.
type WRKChainRootRegisterWrkChain struct {
	ChainId     [32]byte
	GenesisHash [32]byte
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegisterWrkChain is a free log retrieval operation binding the contract event 0xa70d58903899f893c484418f077825c24e4d2508e35cf7f19ec5b0aa30ef43d4.
//
// Solidity: event RegisterWrkChain(bytes32 indexed _chainId, bytes32 _genesisHash, uint256 _timestamp)
func (_WRKChainRoot *WRKChainRootFilterer) FilterRegisterWrkChain(opts *bind.FilterOpts, _chainId [][32]byte) (*WRKChainRootRegisterWrkChainIterator, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "RegisterWrkChain", _chainIdRule)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootRegisterWrkChainIterator{contract: _WRKChainRoot.contract, event: "RegisterWrkChain", logs: logs, sub: sub}, nil
}

// WatchRegisterWrkChain is a free log subscription operation binding the contract event 0xa70d58903899f893c484418f077825c24e4d2508e35cf7f19ec5b0aa30ef43d4.
//
// Solidity: event RegisterWrkChain(bytes32 indexed _chainId, bytes32 _genesisHash, uint256 _timestamp)
func (_WRKChainRoot *WRKChainRootFilterer) WatchRegisterWrkChain(opts *bind.WatchOpts, sink chan<- *WRKChainRootRegisterWrkChain, _chainId [][32]byte) (event.Subscription, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "RegisterWrkChain", _chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootRegisterWrkChain)
				if err := _WRKChainRoot.contract.UnpackLog(event, "RegisterWrkChain", log); err != nil {
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

// ParseRegisterWrkChain is a log parse operation binding the contract event 0xa70d58903899f893c484418f077825c24e4d2508e35cf7f19ec5b0aa30ef43d4.
//
// Solidity: event RegisterWrkChain(bytes32 indexed _chainId, bytes32 _genesisHash, uint256 _timestamp)
func (_WRKChainRoot *WRKChainRootFilterer) ParseRegisterWrkChain(log types.Log) (*WRKChainRootRegisterWrkChain, error) {
	event := new(WRKChainRootRegisterWrkChain)
	if err := _WRKChainRoot.contract.UnpackLog(event, "RegisterWrkChain", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WRKChainRootRemoveAuthorisedAddressIterator is returned from FilterRemoveAuthorisedAddress and is used to iterate over the raw logs and unpacked data for RemoveAuthorisedAddress events raised by the WRKChainRoot contract.
type WRKChainRootRemoveAuthorisedAddressIterator struct {
	Event *WRKChainRootRemoveAuthorisedAddress // Event containing the contract specifics and raw log

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
func (it *WRKChainRootRemoveAuthorisedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootRemoveAuthorisedAddress)
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
		it.Event = new(WRKChainRootRemoveAuthorisedAddress)
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
func (it *WRKChainRootRemoveAuthorisedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootRemoveAuthorisedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootRemoveAuthorisedAddress represents a RemoveAuthorisedAddress event raised by the WRKChainRoot contract.
type WRKChainRootRemoveAuthorisedAddress struct {
	ChainId   [32]byte
	RemovedBy common.Address
	Address   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemoveAuthorisedAddress is a free log retrieval operation binding the contract event 0x1ab6ecd37e23ab30318669bb4642f1cdfba82f23405ce8dfba3648fc9b4744e3.
//
// Solidity: event RemoveAuthorisedAddress(bytes32 indexed _chainId, address indexed _removedBy, address _address)
func (_WRKChainRoot *WRKChainRootFilterer) FilterRemoveAuthorisedAddress(opts *bind.FilterOpts, _chainId [][32]byte, _removedBy []common.Address) (*WRKChainRootRemoveAuthorisedAddressIterator, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}
	var _removedByRule []interface{}
	for _, _removedByItem := range _removedBy {
		_removedByRule = append(_removedByRule, _removedByItem)
	}

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "RemoveAuthorisedAddress", _chainIdRule, _removedByRule)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootRemoveAuthorisedAddressIterator{contract: _WRKChainRoot.contract, event: "RemoveAuthorisedAddress", logs: logs, sub: sub}, nil
}

// WatchRemoveAuthorisedAddress is a free log subscription operation binding the contract event 0x1ab6ecd37e23ab30318669bb4642f1cdfba82f23405ce8dfba3648fc9b4744e3.
//
// Solidity: event RemoveAuthorisedAddress(bytes32 indexed _chainId, address indexed _removedBy, address _address)
func (_WRKChainRoot *WRKChainRootFilterer) WatchRemoveAuthorisedAddress(opts *bind.WatchOpts, sink chan<- *WRKChainRootRemoveAuthorisedAddress, _chainId [][32]byte, _removedBy []common.Address) (event.Subscription, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}
	var _removedByRule []interface{}
	for _, _removedByItem := range _removedBy {
		_removedByRule = append(_removedByRule, _removedByItem)
	}

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "RemoveAuthorisedAddress", _chainIdRule, _removedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootRemoveAuthorisedAddress)
				if err := _WRKChainRoot.contract.UnpackLog(event, "RemoveAuthorisedAddress", log); err != nil {
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

// ParseRemoveAuthorisedAddress is a log parse operation binding the contract event 0x1ab6ecd37e23ab30318669bb4642f1cdfba82f23405ce8dfba3648fc9b4744e3.
//
// Solidity: event RemoveAuthorisedAddress(bytes32 indexed _chainId, address indexed _removedBy, address _address)
func (_WRKChainRoot *WRKChainRootFilterer) ParseRemoveAuthorisedAddress(log types.Log) (*WRKChainRootRemoveAuthorisedAddress, error) {
	event := new(WRKChainRootRemoveAuthorisedAddress)
	if err := _WRKChainRoot.contract.UnpackLog(event, "RemoveAuthorisedAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WRKChainRootWRKChainOwnerChangedIterator is returned from FilterWRKChainOwnerChanged and is used to iterate over the raw logs and unpacked data for WRKChainOwnerChanged events raised by the WRKChainRoot contract.
type WRKChainRootWRKChainOwnerChangedIterator struct {
	Event *WRKChainRootWRKChainOwnerChanged // Event containing the contract specifics and raw log

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
func (it *WRKChainRootWRKChainOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootWRKChainOwnerChanged)
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
		it.Event = new(WRKChainRootWRKChainOwnerChanged)
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
func (it *WRKChainRootWRKChainOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootWRKChainOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootWRKChainOwnerChanged represents a WRKChainOwnerChanged event raised by the WRKChainRoot contract.
type WRKChainRootWRKChainOwnerChanged struct {
	ChainId [32]byte
	Old     common.Address
	New     common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWRKChainOwnerChanged is a free log retrieval operation binding the contract event 0xa59f2eb1d0ae2221e2f9854e3bbc9ff4e5edd43b124d0ac97b6d2d83125d1d19.
//
// Solidity: event WRKChainOwnerChanged(bytes32 indexed _chainId, address _old, address _new)
func (_WRKChainRoot *WRKChainRootFilterer) FilterWRKChainOwnerChanged(opts *bind.FilterOpts, _chainId [][32]byte) (*WRKChainRootWRKChainOwnerChangedIterator, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "WRKChainOwnerChanged", _chainIdRule)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootWRKChainOwnerChangedIterator{contract: _WRKChainRoot.contract, event: "WRKChainOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchWRKChainOwnerChanged is a free log subscription operation binding the contract event 0xa59f2eb1d0ae2221e2f9854e3bbc9ff4e5edd43b124d0ac97b6d2d83125d1d19.
//
// Solidity: event WRKChainOwnerChanged(bytes32 indexed _chainId, address _old, address _new)
func (_WRKChainRoot *WRKChainRootFilterer) WatchWRKChainOwnerChanged(opts *bind.WatchOpts, sink chan<- *WRKChainRootWRKChainOwnerChanged, _chainId [][32]byte) (event.Subscription, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "WRKChainOwnerChanged", _chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootWRKChainOwnerChanged)
				if err := _WRKChainRoot.contract.UnpackLog(event, "WRKChainOwnerChanged", log); err != nil {
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

// ParseWRKChainOwnerChanged is a log parse operation binding the contract event 0xa59f2eb1d0ae2221e2f9854e3bbc9ff4e5edd43b124d0ac97b6d2d83125d1d19.
//
// Solidity: event WRKChainOwnerChanged(bytes32 indexed _chainId, address _old, address _new)
func (_WRKChainRoot *WRKChainRootFilterer) ParseWRKChainOwnerChanged(log types.Log) (*WRKChainRootWRKChainOwnerChanged, error) {
	event := new(WRKChainRootWRKChainOwnerChanged)
	if err := _WRKChainRoot.contract.UnpackLog(event, "WRKChainOwnerChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
