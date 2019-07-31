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
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820a3968129a6b83f89d1789cb5e64d378309030066b0fdbe2b5e0b4a239f4247210029`

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
const WRKChainRootABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_addressesToRemove\",\"type\":\"address[]\"}],\"name\":\"removeAuthAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_newAuthAddresses\",\"type\":\"address[]\"}],\"name\":\"addAuthAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_authAddresses\",\"type\":\"address[]\"},{\"name\":\"_genesisHash\",\"type\":\"bytes32\"}],\"name\":\"registerWrkChain\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getNumBlocksSubmitted\",\"outputs\":[{\"name\":\"numBlocksSubmitted_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wrkchainList\",\"outputs\":[{\"name\":\"lastBlock\",\"type\":\"uint256\"},{\"name\":\"genesisHash\",\"type\":\"bytes32\"},{\"name\":\"isWrkchain\",\"type\":\"bool\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"numBlocksSubmitted\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_height\",\"type\":\"uint256\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"name\":\"_receiptRoot\",\"type\":\"bytes32\"},{\"name\":\"_txRoot\",\"type\":\"bytes32\"},{\"name\":\"_stateRoot\",\"type\":\"bytes32\"},{\"name\":\"_blockSigner\",\"type\":\"address\"}],\"name\":\"recordHeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isAuthorisedAddress\",\"outputs\":[{\"name\":\"_isAuthorised\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getGenesis\",\"outputs\":[{\"name\":\"genesisHash_\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"},{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeWrkchainOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getLastRecordedBlockNum\",\"outputs\":[{\"name\":\"lastBlock_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_wrkchainRegDeposit\",\"type\":\"uint256\"},{\"name\":\"_minBlocksForDepositReturn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_height\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_receiptRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_txRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_blockSigner\",\"type\":\"address\"}],\"name\":\"RecordHeader\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_genesisHash\",\"type\":\"bytes32\"}],\"name\":\"RegisterWrkChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogFallbackFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WRKChainDepositRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_authorisedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"AuthoriseNewAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_removedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"RemoveAuthorisedAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_old\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_new\",\"type\":\"address\"}],\"name\":\"WRKChainOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WRKChainDepositPaid\",\"type\":\"event\"}]"

// WRKChainRootBin is the compiled bytecode used for deploying new contracts.
const WRKChainRootBin = `0x608060405234801561001057600080fd5b50604051604080611acb8339810160405280516020909101516000821161009857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4465706f73697420616d6f756e742073686f756c64206265203e203000000000604482015290519081900360640190fd5b6000811161010757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4d696e20626c6f636b732073686f756c64206265203e20300000000000000000604482015290519081900360640190fd5b6000919091556001556119ac8061011f6000396000f3006080604052600436106100a35763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025c30a681146100db5780631c52cff3146101375780635be4e228146101915780635eefa77f146101ab5780638021982a146101d5578063a78c498814610221578063ac23db0c14610257578063cfdf3e841461028f578063eec5d0b2146102a7578063f4da14f6146102cb575b60408051348152905133917fb1db2589d999730e8faec7dccf62c4d28a7f6ef1095b46c9a744dcd46d87130e919081900360200190a2005b3480156100e757600080fd5b50604080516020600460248035828101358481028087018601909752808652610135968435963696604495919490910192918291850190849080828437509497506102e39650505050505050565b005b34801561014357600080fd5b50604080516020600460248035828101358481028087018601909752808652610135968435963696604495919490910192918291850190849080828437509497506106689650505050505050565b6101356004803590602480359081019101356044356109f4565b3480156101b757600080fd5b506101c3600435610e54565b60408051918252519081900360200190f35b3480156101e157600080fd5b506101ed600435610f0a565b60408051958652602086019490945291151584840152600160a060020a031660608401526080830152519081900360a00190f35b34801561022d57600080fd5b5061013560043560243560443560643560843560a43560c435600160a060020a0360e43516610f4a565b34801561026357600080fd5b5061027b600435600160a060020a0360243516611308565b604080519115158252519081900360200190f35b34801561029b57600080fd5b506101c36004356113d7565b3480156102b357600080fd5b50610135600435600160a060020a036024351661148d565b3480156102d757600080fd5b506101c36004356117a7565b600082815260046020526040812060020154839060ff16151561033e576040805160e560020a62461bcd0281526020600482015260166024820152600080516020611961833981519152604482015290519081900360640190fd5b60008181526004602081905260409091200154600160a060020a031633146103b0576040805160e560020a62461bcd02815260206004820152601560248201527f596f7520617265206e6f7420746865206f776e65720000000000000000000000604482015290519081900360640190fd5b600084116103f6576040805160e560020a62461bcd0281526020600482015260116024820152600080516020611921833981519152604482015290519081900360640190fd5b825160001061044f576040805160e560020a62461bcd02815260206004820152601760248201527f4175746820616464726573736573207265717569726564000000000000000000604482015290519081900360640190fd5b8251600a10156104cf576040805160e560020a62461bcd02815260206004820152602360248201527f4d6178207375626d697373696f6e2031302061646472657373657320616c6c6f60448201527f7765640000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008481526004602052604090206002015460ff161515610528576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611941833981519152604482015290519081900360640190fd5b600091505b825182101561066257825160009084908490811061054757fe5b60209081029091010151600160a060020a031614156105b0576040805160e560020a62461bcd02815260206004820152601d60248201527f63616e6e6f7420617574686f72697365207a65726f2061646472657373000000604482015290519081900360640190fd5b60008481526004602052604081208451600390910191908590859081106105d357fe5b6020908102909101810151600160a060020a03168252810191909152604001600020805460ff191690558251339085907f73d5cc6b165e5939d279249c551c64c6a67b23d2df045acd47a75971ceca4a1f9086908690811061063157fe5b602090810290910181015160408051600160a060020a039092168252519081900390910190a360019091019061052d565b50505050565b600082815260046020526040812060020154839060ff1615156106c3576040805160e560020a62461bcd0281526020600482015260166024820152600080516020611961833981519152604482015290519081900360640190fd5b60008181526004602081905260409091200154600160a060020a03163314610735576040805160e560020a62461bcd02815260206004820152601560248201527f596f7520617265206e6f7420746865206f776e65720000000000000000000000604482015290519081900360640190fd5b6000841161077b576040805160e560020a62461bcd0281526020600482015260116024820152600080516020611921833981519152604482015290519081900360640190fd5b82516000106107d4576040805160e560020a62461bcd02815260206004820152601760248201527f4175746820616464726573736573207265717569726564000000000000000000604482015290519081900360640190fd5b8251600a1015610854576040805160e560020a62461bcd02815260206004820152602360248201527f4d6178207375626d697373696f6e2031302061646472657373657320616c6c6f60448201527f7765640000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008481526004602052604090206002015460ff1615156108ad576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611941833981519152604482015290519081900360640190fd5b600091505b82518210156106625782516000908490849081106108cc57fe5b60209081029091010151600160a060020a03161415610935576040805160e560020a62461bcd02815260206004820152601d60248201527f63616e6e6f7420617574686f72697365207a65726f2061646472657373000000604482015290519081900360640190fd5b6001600460008681526020019081526020016000206003016000858581518110151561095d57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff19169115159190911790558251339085907f4cf08eb13c8e5bd0af6b9ce939523be90d1494cf5001a7b594b785467090b53a908690869081106109c357fe5b602090810290910181015160408051600160a060020a039092168252519081900390910190a36001909101906108b2565b60008481526004602052604081206002015460ff1615610a5e576040805160e560020a62461bcd02815260206004820152601760248201527f436861696e20494420616c726561647920657869737473000000000000000000604482015290519081900360640190fd5b60005434148015610a6f5750600034115b1515610ac5576040805160e560020a62461bcd02815260206004820152601060248201527f4465706f73697420726571756972656400000000000000000000000000000000604482015290519081900360640190fd5b60008511610b0b576040805160e560020a62461bcd0281526020600482015260116024820152600080516020611921833981519152604482015290519081900360640190fd5b60008311610b89576040805160e560020a62461bcd02815260206004820152602560248201527f496e697469616c20417574686f7269736564206164647265737365732072657160448201527f7569726564000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600a831115610c08576040805160e560020a62461bcd02815260206004820152602f60248201527f4d6178696d756d20313020496e697469616c20417574686f726973656420616460448201527f6472657373657320616c6c6f7765640000000000000000000000000000000000606482015290519081900360840190fd5b6040805160a0810182526000808252602080830186815260018486018181523360608701908152608087018681528d8752600480875296899020975188559351928701929092555160028601805460ff191691151591909117905551928401805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039094169390931790925590516005909201919091558151848152915187927fd49920535e715987d1fbf16411df1f3f3347d62e741006463c13550661ce4b6992908290030190a25060005b82811015610d89576000858152600460205260408120600191600390910190868685818110610cfe57fe5b60209081029290920135600160a060020a0316835250810191909152604001600020805460ff191691151591909117905533857f4cf08eb13c8e5bd0af6b9ce939523be90d1494cf5001a7b594b785467090b53a868685818110610d5e57fe5b6040805160209283029490940135600160a060020a03168452519283900301919050a3600101610cd3565b6000858152600460209081526040808320338085526003909101835292819020805460ff191660011790558051838152905188927f4cf08eb13c8e5bd0af6b9ce939523be90d1494cf5001a7b594b785467090b53a928290030190a3600254610df8903463ffffffff61185a16565b60025533600081815260036020908152604080832089845282529182902034908190558251908152915188927f8da0661955dba0c91f23c6cc86d131f83a23d8e4f7569a808fa145630e47863e92908290030190a35050505050565b6000808211610e9b576040805160e560020a62461bcd0281526020600482015260116024820152600080516020611921833981519152604482015290519081900360640190fd5b60008281526004602052604090206002015460ff161515610ef4576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611941833981519152604482015290519081900360640190fd5b5060009081526004602052604090206005015490565b6004602081905260009182526040909120805460018201546002830154938301546005909301549193909260ff90911691600160a060020a039091169085565b600088815260046020526040812060020154899060ff161515610fa5576040805160e560020a62461bcd0281526020600482015260166024820152600080516020611961833981519152604482015290519081900360640190fd5b600081815260046020908152604080832033845260030190915290205460ff16151560011480610fef575060008181526004602081905260409091200154600160a060020a031633145b1515610ffa57600080fd5b60008a11611040576040805160e560020a62461bcd0281526020600482015260116024820152600080516020611921833981519152604482015290519081900360640190fd5b60008911611098576040805160e560020a62461bcd02815260206004820152601560248201527f426c6f636b206e756d6265722072657175697265640000000000000000000000604482015290519081900360640190fd5b60008a81526004602052604090206002015460ff1615156110f1576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611941833981519152604482015290519081900360640190fd5b60008a81526004602052604090208054909250891161115a576040805160e560020a62461bcd02815260206004820152601d60248201527f426c6f636b2068656164657220616c7265616479207265636f72646564000000604482015290519081900360640190fd5b60008a81526004602052604090208981556005015461118090600163ffffffff61185a16565b60008b81526004602052604090206005018190556001541480156111d35750600080548b825260046020818152604080852090920154600160a060020a03168452600381528184208e8552905290912054145b1561129f576004820154600160a060020a031660009081526003602090815260408083208d84529091528120819055546002546112159163ffffffff6118be16565b60025560048201546000546040805191825251600160a060020a03909216918c917f5fdfd7834e3e713189c97ffe1bc13c31b4e5003b84f2b73cdf2055cfe49bf647919081900360200190a3600482015460008054604051600160a060020a039093169281156108fc0292818181858888f1935050505015801561129d573d6000803e3d6000fd5b505b604080518881526020810188905280820187905260608101869052600160a060020a0385166080820152905189918b918d917f943621e8fb786306cd7b16a4467604107bbf3b0b04b536d39a245ea68c1062ef919081900360a00190a450505050505050505050565b600080831161134f576040805160e560020a62461bcd0281526020600482015260116024820152600080516020611921833981519152604482015290519081900360640190fd5b60008381526004602052604090206002015460ff1615156113a8576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611941833981519152604482015290519081900360640190fd5b506000918252600460209081526040808420600160a060020a0390931684526003909201905290205460ff1690565b600080821161141e576040805160e560020a62461bcd0281526020600482015260116024820152600080516020611921833981519152604482015290519081900360640190fd5b60008281526004602052604090206002015460ff161515611477576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611941833981519152604482015290519081900360640190fd5b5060009081526004602052604090206001015490565b600082815260046020526040812060020154839060ff1615156114e8576040805160e560020a62461bcd0281526020600482015260166024820152600080516020611961833981519152604482015290519081900360640190fd5b60008181526004602081905260409091200154600160a060020a0316331461155a576040805160e560020a62461bcd02815260206004820152601560248201527f596f7520617265206e6f7420746865206f776e65720000000000000000000000604482015290519081900360640190fd5b600084116115a0576040805160e560020a62461bcd0281526020600482015260116024820152600080516020611921833981519152604482015290519081900360640190fd5b600160a060020a0383161515611626576040805160e560020a62461bcd02815260206004820152603260248201527f4e6577206f776e65722072657175697265642c20616e642073686f756c64206260448201527f65206e6f6e2d7a65726f20616464726573730000000000000000000000000000606482015290519081900360840190fd5b60008481526004602052604090206002015460ff16151561167f576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611941833981519152604482015290519081900360640190fd5b6000848152600460205260409020600154600582015491935010611713576040805160e560020a62461bcd02815260206004820152603860248201527f4f776e65722063616e6e6f74206265206368616e676564206265666f7265206460448201527f65706f73697420686173206265656e2072657475726e65640000000000000000606482015290519081900360840190fd5b6000848152600460208181526040808420928301805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0389169081179091558085526003909301825292839020805460ff19166001179055825133815290810191909152815186927ff4fb8ce4b6ba19f99d4309750d2d0bf6ddb2eab6e3ff6f1d2bf27dea39fdb067928290030190a250505050565b60008082116117ee576040805160e560020a62461bcd0281526020600482015260116024820152600080516020611921833981519152604482015290519081900360640190fd5b60008281526004602052604090206002015460ff161515611847576040805160e560020a62461bcd0281526020600482015260176024820152600080516020611941833981519152604482015290519081900360640190fd5b5060009081526004602052604090205490565b6000828201838110156118b7576040805160e560020a62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b60008083831115611919576040805160e560020a62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50509003905600436861696e204944207265717569726564000000000000000000000000000000436861696e20494420646f6573206e6f74206578697374000000000000000000436861696e494420646f6573206e6f7420657869737400000000000000000000a165627a7a7230582021b58a75ab3e8c5f38b771136ddf60fe37b25cb9b4c41e8b8e9fa6cc92a27b980029`

// DeployWRKChainRoot deploys a new Ethereum contract, binding an instance of WRKChainRoot to it.
func DeployWRKChainRoot(auth *bind.TransactOpts, backend bind.ContractBackend, _wrkchainRegDeposit *big.Int, _minBlocksForDepositReturn *big.Int) (common.Address, *types.Transaction, *WRKChainRoot, error) {
	parsed, err := abi.JSON(strings.NewReader(WRKChainRootABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WRKChainRootBin), backend, _wrkchainRegDeposit, _minBlocksForDepositReturn)
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

// GetGenesis is a free data retrieval call binding the contract method 0xcfdf3e84.
//
// Solidity: function getGenesis(uint256 _chainId) constant returns(bytes32 genesisHash_)
func (_WRKChainRoot *WRKChainRootCaller) GetGenesis(opts *bind.CallOpts, _chainId *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "getGenesis", _chainId)
	return *ret0, err
}

// GetGenesis is a free data retrieval call binding the contract method 0xcfdf3e84.
//
// Solidity: function getGenesis(uint256 _chainId) constant returns(bytes32 genesisHash_)
func (_WRKChainRoot *WRKChainRootSession) GetGenesis(_chainId *big.Int) ([32]byte, error) {
	return _WRKChainRoot.Contract.GetGenesis(&_WRKChainRoot.CallOpts, _chainId)
}

// GetGenesis is a free data retrieval call binding the contract method 0xcfdf3e84.
//
// Solidity: function getGenesis(uint256 _chainId) constant returns(bytes32 genesisHash_)
func (_WRKChainRoot *WRKChainRootCallerSession) GetGenesis(_chainId *big.Int) ([32]byte, error) {
	return _WRKChainRoot.Contract.GetGenesis(&_WRKChainRoot.CallOpts, _chainId)
}

// GetLastRecordedBlockNum is a free data retrieval call binding the contract method 0xf4da14f6.
//
// Solidity: function getLastRecordedBlockNum(uint256 _chainId) constant returns(uint256 lastBlock_)
func (_WRKChainRoot *WRKChainRootCaller) GetLastRecordedBlockNum(opts *bind.CallOpts, _chainId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "getLastRecordedBlockNum", _chainId)
	return *ret0, err
}

// GetLastRecordedBlockNum is a free data retrieval call binding the contract method 0xf4da14f6.
//
// Solidity: function getLastRecordedBlockNum(uint256 _chainId) constant returns(uint256 lastBlock_)
func (_WRKChainRoot *WRKChainRootSession) GetLastRecordedBlockNum(_chainId *big.Int) (*big.Int, error) {
	return _WRKChainRoot.Contract.GetLastRecordedBlockNum(&_WRKChainRoot.CallOpts, _chainId)
}

// GetLastRecordedBlockNum is a free data retrieval call binding the contract method 0xf4da14f6.
//
// Solidity: function getLastRecordedBlockNum(uint256 _chainId) constant returns(uint256 lastBlock_)
func (_WRKChainRoot *WRKChainRootCallerSession) GetLastRecordedBlockNum(_chainId *big.Int) (*big.Int, error) {
	return _WRKChainRoot.Contract.GetLastRecordedBlockNum(&_WRKChainRoot.CallOpts, _chainId)
}

// GetNumBlocksSubmitted is a free data retrieval call binding the contract method 0x5eefa77f.
//
// Solidity: function getNumBlocksSubmitted(uint256 _chainId) constant returns(uint256 numBlocksSubmitted_)
func (_WRKChainRoot *WRKChainRootCaller) GetNumBlocksSubmitted(opts *bind.CallOpts, _chainId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "getNumBlocksSubmitted", _chainId)
	return *ret0, err
}

// GetNumBlocksSubmitted is a free data retrieval call binding the contract method 0x5eefa77f.
//
// Solidity: function getNumBlocksSubmitted(uint256 _chainId) constant returns(uint256 numBlocksSubmitted_)
func (_WRKChainRoot *WRKChainRootSession) GetNumBlocksSubmitted(_chainId *big.Int) (*big.Int, error) {
	return _WRKChainRoot.Contract.GetNumBlocksSubmitted(&_WRKChainRoot.CallOpts, _chainId)
}

// GetNumBlocksSubmitted is a free data retrieval call binding the contract method 0x5eefa77f.
//
// Solidity: function getNumBlocksSubmitted(uint256 _chainId) constant returns(uint256 numBlocksSubmitted_)
func (_WRKChainRoot *WRKChainRootCallerSession) GetNumBlocksSubmitted(_chainId *big.Int) (*big.Int, error) {
	return _WRKChainRoot.Contract.GetNumBlocksSubmitted(&_WRKChainRoot.CallOpts, _chainId)
}

// IsAuthorisedAddress is a free data retrieval call binding the contract method 0xac23db0c.
//
// Solidity: function isAuthorisedAddress(uint256 _chainId, address _address) constant returns(bool _isAuthorised)
func (_WRKChainRoot *WRKChainRootCaller) IsAuthorisedAddress(opts *bind.CallOpts, _chainId *big.Int, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WRKChainRoot.contract.Call(opts, out, "isAuthorisedAddress", _chainId, _address)
	return *ret0, err
}

// IsAuthorisedAddress is a free data retrieval call binding the contract method 0xac23db0c.
//
// Solidity: function isAuthorisedAddress(uint256 _chainId, address _address) constant returns(bool _isAuthorised)
func (_WRKChainRoot *WRKChainRootSession) IsAuthorisedAddress(_chainId *big.Int, _address common.Address) (bool, error) {
	return _WRKChainRoot.Contract.IsAuthorisedAddress(&_WRKChainRoot.CallOpts, _chainId, _address)
}

// IsAuthorisedAddress is a free data retrieval call binding the contract method 0xac23db0c.
//
// Solidity: function isAuthorisedAddress(uint256 _chainId, address _address) constant returns(bool _isAuthorised)
func (_WRKChainRoot *WRKChainRootCallerSession) IsAuthorisedAddress(_chainId *big.Int, _address common.Address) (bool, error) {
	return _WRKChainRoot.Contract.IsAuthorisedAddress(&_WRKChainRoot.CallOpts, _chainId, _address)
}

// WrkchainList is a free data retrieval call binding the contract method 0x8021982a.
//
// Solidity: function wrkchainList(uint256 ) constant returns(uint256 lastBlock, bytes32 genesisHash, bool isWrkchain, address owner, uint256 numBlocksSubmitted)
func (_WRKChainRoot *WRKChainRootCaller) WrkchainList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LastBlock          *big.Int
	GenesisHash        [32]byte
	IsWrkchain         bool
	Owner              common.Address
	NumBlocksSubmitted *big.Int
}, error) {
	ret := new(struct {
		LastBlock          *big.Int
		GenesisHash        [32]byte
		IsWrkchain         bool
		Owner              common.Address
		NumBlocksSubmitted *big.Int
	})
	out := ret
	err := _WRKChainRoot.contract.Call(opts, out, "wrkchainList", arg0)
	return *ret, err
}

// WrkchainList is a free data retrieval call binding the contract method 0x8021982a.
//
// Solidity: function wrkchainList(uint256 ) constant returns(uint256 lastBlock, bytes32 genesisHash, bool isWrkchain, address owner, uint256 numBlocksSubmitted)
func (_WRKChainRoot *WRKChainRootSession) WrkchainList(arg0 *big.Int) (struct {
	LastBlock          *big.Int
	GenesisHash        [32]byte
	IsWrkchain         bool
	Owner              common.Address
	NumBlocksSubmitted *big.Int
}, error) {
	return _WRKChainRoot.Contract.WrkchainList(&_WRKChainRoot.CallOpts, arg0)
}

// WrkchainList is a free data retrieval call binding the contract method 0x8021982a.
//
// Solidity: function wrkchainList(uint256 ) constant returns(uint256 lastBlock, bytes32 genesisHash, bool isWrkchain, address owner, uint256 numBlocksSubmitted)
func (_WRKChainRoot *WRKChainRootCallerSession) WrkchainList(arg0 *big.Int) (struct {
	LastBlock          *big.Int
	GenesisHash        [32]byte
	IsWrkchain         bool
	Owner              common.Address
	NumBlocksSubmitted *big.Int
}, error) {
	return _WRKChainRoot.Contract.WrkchainList(&_WRKChainRoot.CallOpts, arg0)
}

// AddAuthAddresses is a paid mutator transaction binding the contract method 0x1c52cff3.
//
// Solidity: function addAuthAddresses(uint256 _chainId, address[] _newAuthAddresses) returns()
func (_WRKChainRoot *WRKChainRootTransactor) AddAuthAddresses(opts *bind.TransactOpts, _chainId *big.Int, _newAuthAddresses []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "addAuthAddresses", _chainId, _newAuthAddresses)
}

// AddAuthAddresses is a paid mutator transaction binding the contract method 0x1c52cff3.
//
// Solidity: function addAuthAddresses(uint256 _chainId, address[] _newAuthAddresses) returns()
func (_WRKChainRoot *WRKChainRootSession) AddAuthAddresses(_chainId *big.Int, _newAuthAddresses []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.AddAuthAddresses(&_WRKChainRoot.TransactOpts, _chainId, _newAuthAddresses)
}

// AddAuthAddresses is a paid mutator transaction binding the contract method 0x1c52cff3.
//
// Solidity: function addAuthAddresses(uint256 _chainId, address[] _newAuthAddresses) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) AddAuthAddresses(_chainId *big.Int, _newAuthAddresses []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.AddAuthAddresses(&_WRKChainRoot.TransactOpts, _chainId, _newAuthAddresses)
}

// ChangeWrkchainOwner is a paid mutator transaction binding the contract method 0xeec5d0b2.
//
// Solidity: function changeWrkchainOwner(uint256 _chainId, address _newOwner) returns()
func (_WRKChainRoot *WRKChainRootTransactor) ChangeWrkchainOwner(opts *bind.TransactOpts, _chainId *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "changeWrkchainOwner", _chainId, _newOwner)
}

// ChangeWrkchainOwner is a paid mutator transaction binding the contract method 0xeec5d0b2.
//
// Solidity: function changeWrkchainOwner(uint256 _chainId, address _newOwner) returns()
func (_WRKChainRoot *WRKChainRootSession) ChangeWrkchainOwner(_chainId *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.ChangeWrkchainOwner(&_WRKChainRoot.TransactOpts, _chainId, _newOwner)
}

// ChangeWrkchainOwner is a paid mutator transaction binding the contract method 0xeec5d0b2.
//
// Solidity: function changeWrkchainOwner(uint256 _chainId, address _newOwner) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) ChangeWrkchainOwner(_chainId *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.ChangeWrkchainOwner(&_WRKChainRoot.TransactOpts, _chainId, _newOwner)
}

// RecordHeader is a paid mutator transaction binding the contract method 0xa78c4988.
//
// Solidity: function recordHeader(uint256 _chainId, uint256 _height, bytes32 _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, address _blockSigner) returns()
func (_WRKChainRoot *WRKChainRootTransactor) RecordHeader(opts *bind.TransactOpts, _chainId *big.Int, _height *big.Int, _hash [32]byte, _parentHash [32]byte, _receiptRoot [32]byte, _txRoot [32]byte, _stateRoot [32]byte, _blockSigner common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "recordHeader", _chainId, _height, _hash, _parentHash, _receiptRoot, _txRoot, _stateRoot, _blockSigner)
}

// RecordHeader is a paid mutator transaction binding the contract method 0xa78c4988.
//
// Solidity: function recordHeader(uint256 _chainId, uint256 _height, bytes32 _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, address _blockSigner) returns()
func (_WRKChainRoot *WRKChainRootSession) RecordHeader(_chainId *big.Int, _height *big.Int, _hash [32]byte, _parentHash [32]byte, _receiptRoot [32]byte, _txRoot [32]byte, _stateRoot [32]byte, _blockSigner common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RecordHeader(&_WRKChainRoot.TransactOpts, _chainId, _height, _hash, _parentHash, _receiptRoot, _txRoot, _stateRoot, _blockSigner)
}

// RecordHeader is a paid mutator transaction binding the contract method 0xa78c4988.
//
// Solidity: function recordHeader(uint256 _chainId, uint256 _height, bytes32 _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, address _blockSigner) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) RecordHeader(_chainId *big.Int, _height *big.Int, _hash [32]byte, _parentHash [32]byte, _receiptRoot [32]byte, _txRoot [32]byte, _stateRoot [32]byte, _blockSigner common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RecordHeader(&_WRKChainRoot.TransactOpts, _chainId, _height, _hash, _parentHash, _receiptRoot, _txRoot, _stateRoot, _blockSigner)
}

// RegisterWrkChain is a paid mutator transaction binding the contract method 0x5be4e228.
//
// Solidity: function registerWrkChain(uint256 _chainId, address[] _authAddresses, bytes32 _genesisHash) returns()
func (_WRKChainRoot *WRKChainRootTransactor) RegisterWrkChain(opts *bind.TransactOpts, _chainId *big.Int, _authAddresses []common.Address, _genesisHash [32]byte) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "registerWrkChain", _chainId, _authAddresses, _genesisHash)
}

// RegisterWrkChain is a paid mutator transaction binding the contract method 0x5be4e228.
//
// Solidity: function registerWrkChain(uint256 _chainId, address[] _authAddresses, bytes32 _genesisHash) returns()
func (_WRKChainRoot *WRKChainRootSession) RegisterWrkChain(_chainId *big.Int, _authAddresses []common.Address, _genesisHash [32]byte) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RegisterWrkChain(&_WRKChainRoot.TransactOpts, _chainId, _authAddresses, _genesisHash)
}

// RegisterWrkChain is a paid mutator transaction binding the contract method 0x5be4e228.
//
// Solidity: function registerWrkChain(uint256 _chainId, address[] _authAddresses, bytes32 _genesisHash) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) RegisterWrkChain(_chainId *big.Int, _authAddresses []common.Address, _genesisHash [32]byte) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RegisterWrkChain(&_WRKChainRoot.TransactOpts, _chainId, _authAddresses, _genesisHash)
}

// RemoveAuthAddresses is a paid mutator transaction binding the contract method 0x025c30a6.
//
// Solidity: function removeAuthAddresses(uint256 _chainId, address[] _addressesToRemove) returns()
func (_WRKChainRoot *WRKChainRootTransactor) RemoveAuthAddresses(opts *bind.TransactOpts, _chainId *big.Int, _addressesToRemove []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.contract.Transact(opts, "removeAuthAddresses", _chainId, _addressesToRemove)
}

// RemoveAuthAddresses is a paid mutator transaction binding the contract method 0x025c30a6.
//
// Solidity: function removeAuthAddresses(uint256 _chainId, address[] _addressesToRemove) returns()
func (_WRKChainRoot *WRKChainRootSession) RemoveAuthAddresses(_chainId *big.Int, _addressesToRemove []common.Address) (*types.Transaction, error) {
	return _WRKChainRoot.Contract.RemoveAuthAddresses(&_WRKChainRoot.TransactOpts, _chainId, _addressesToRemove)
}

// RemoveAuthAddresses is a paid mutator transaction binding the contract method 0x025c30a6.
//
// Solidity: function removeAuthAddresses(uint256 _chainId, address[] _addressesToRemove) returns()
func (_WRKChainRoot *WRKChainRootTransactorSession) RemoveAuthAddresses(_chainId *big.Int, _addressesToRemove []common.Address) (*types.Transaction, error) {
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
	ChainId      *big.Int
	AuthorisedBy common.Address
	Address      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuthoriseNewAddress is a free log retrieval operation binding the contract event 0x4cf08eb13c8e5bd0af6b9ce939523be90d1494cf5001a7b594b785467090b53a.
//
// Solidity: event AuthoriseNewAddress(uint256 indexed _chainId, address indexed _authorisedBy, address _address)
func (_WRKChainRoot *WRKChainRootFilterer) FilterAuthoriseNewAddress(opts *bind.FilterOpts, _chainId []*big.Int, _authorisedBy []common.Address) (*WRKChainRootAuthoriseNewAddressIterator, error) {

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

// WatchAuthoriseNewAddress is a free log subscription operation binding the contract event 0x4cf08eb13c8e5bd0af6b9ce939523be90d1494cf5001a7b594b785467090b53a.
//
// Solidity: event AuthoriseNewAddress(uint256 indexed _chainId, address indexed _authorisedBy, address _address)
func (_WRKChainRoot *WRKChainRootFilterer) WatchAuthoriseNewAddress(opts *bind.WatchOpts, sink chan<- *WRKChainRootAuthoriseNewAddress, _chainId []*big.Int, _authorisedBy []common.Address) (event.Subscription, error) {

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
	ChainId     *big.Int
	Height      *big.Int
	Hash        [32]byte
	ParentHash  [32]byte
	ReceiptRoot [32]byte
	TxRoot      [32]byte
	StateRoot   [32]byte
	BlockSigner common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRecordHeader is a free log retrieval operation binding the contract event 0x943621e8fb786306cd7b16a4467604107bbf3b0b04b536d39a245ea68c1062ef.
//
// Solidity: event RecordHeader(uint256 indexed _chainId, uint256 indexed _height, bytes32 indexed _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, address _blockSigner)
func (_WRKChainRoot *WRKChainRootFilterer) FilterRecordHeader(opts *bind.FilterOpts, _chainId []*big.Int, _height []*big.Int, _hash [][32]byte) (*WRKChainRootRecordHeaderIterator, error) {

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

// WatchRecordHeader is a free log subscription operation binding the contract event 0x943621e8fb786306cd7b16a4467604107bbf3b0b04b536d39a245ea68c1062ef.
//
// Solidity: event RecordHeader(uint256 indexed _chainId, uint256 indexed _height, bytes32 indexed _hash, bytes32 _parentHash, bytes32 _receiptRoot, bytes32 _txRoot, bytes32 _stateRoot, address _blockSigner)
func (_WRKChainRoot *WRKChainRootFilterer) WatchRecordHeader(opts *bind.WatchOpts, sink chan<- *WRKChainRootRecordHeader, _chainId []*big.Int, _height []*big.Int, _hash [][32]byte) (event.Subscription, error) {

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
	ChainId     *big.Int
	GenesisHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegisterWrkChain is a free log retrieval operation binding the contract event 0xd49920535e715987d1fbf16411df1f3f3347d62e741006463c13550661ce4b69.
//
// Solidity: event RegisterWrkChain(uint256 indexed _chainId, bytes32 _genesisHash)
func (_WRKChainRoot *WRKChainRootFilterer) FilterRegisterWrkChain(opts *bind.FilterOpts, _chainId []*big.Int) (*WRKChainRootRegisterWrkChainIterator, error) {

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

// WatchRegisterWrkChain is a free log subscription operation binding the contract event 0xd49920535e715987d1fbf16411df1f3f3347d62e741006463c13550661ce4b69.
//
// Solidity: event RegisterWrkChain(uint256 indexed _chainId, bytes32 _genesisHash)
func (_WRKChainRoot *WRKChainRootFilterer) WatchRegisterWrkChain(opts *bind.WatchOpts, sink chan<- *WRKChainRootRegisterWrkChain, _chainId []*big.Int) (event.Subscription, error) {

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
	ChainId   *big.Int
	RemovedBy common.Address
	Address   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemoveAuthorisedAddress is a free log retrieval operation binding the contract event 0x73d5cc6b165e5939d279249c551c64c6a67b23d2df045acd47a75971ceca4a1f.
//
// Solidity: event RemoveAuthorisedAddress(uint256 indexed _chainId, address indexed _removedBy, address _address)
func (_WRKChainRoot *WRKChainRootFilterer) FilterRemoveAuthorisedAddress(opts *bind.FilterOpts, _chainId []*big.Int, _removedBy []common.Address) (*WRKChainRootRemoveAuthorisedAddressIterator, error) {

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

// WatchRemoveAuthorisedAddress is a free log subscription operation binding the contract event 0x73d5cc6b165e5939d279249c551c64c6a67b23d2df045acd47a75971ceca4a1f.
//
// Solidity: event RemoveAuthorisedAddress(uint256 indexed _chainId, address indexed _removedBy, address _address)
func (_WRKChainRoot *WRKChainRootFilterer) WatchRemoveAuthorisedAddress(opts *bind.WatchOpts, sink chan<- *WRKChainRootRemoveAuthorisedAddress, _chainId []*big.Int, _removedBy []common.Address) (event.Subscription, error) {

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

// WRKChainRootWRKChainDepositPaidIterator is returned from FilterWRKChainDepositPaid and is used to iterate over the raw logs and unpacked data for WRKChainDepositPaid events raised by the WRKChainRoot contract.
type WRKChainRootWRKChainDepositPaidIterator struct {
	Event *WRKChainRootWRKChainDepositPaid // Event containing the contract specifics and raw log

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
func (it *WRKChainRootWRKChainDepositPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootWRKChainDepositPaid)
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
		it.Event = new(WRKChainRootWRKChainDepositPaid)
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
func (it *WRKChainRootWRKChainDepositPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootWRKChainDepositPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootWRKChainDepositPaid represents a WRKChainDepositPaid event raised by the WRKChainRoot contract.
type WRKChainRootWRKChainDepositPaid struct {
	ChainId *big.Int
	Owner   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWRKChainDepositPaid is a free log retrieval operation binding the contract event 0x8da0661955dba0c91f23c6cc86d131f83a23d8e4f7569a808fa145630e47863e.
//
// Solidity: event WRKChainDepositPaid(uint256 indexed _chainId, address indexed _owner, uint256 _amount)
func (_WRKChainRoot *WRKChainRootFilterer) FilterWRKChainDepositPaid(opts *bind.FilterOpts, _chainId []*big.Int, _owner []common.Address) (*WRKChainRootWRKChainDepositPaidIterator, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "WRKChainDepositPaid", _chainIdRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootWRKChainDepositPaidIterator{contract: _WRKChainRoot.contract, event: "WRKChainDepositPaid", logs: logs, sub: sub}, nil
}

// WatchWRKChainDepositPaid is a free log subscription operation binding the contract event 0x8da0661955dba0c91f23c6cc86d131f83a23d8e4f7569a808fa145630e47863e.
//
// Solidity: event WRKChainDepositPaid(uint256 indexed _chainId, address indexed _owner, uint256 _amount)
func (_WRKChainRoot *WRKChainRootFilterer) WatchWRKChainDepositPaid(opts *bind.WatchOpts, sink chan<- *WRKChainRootWRKChainDepositPaid, _chainId []*big.Int, _owner []common.Address) (event.Subscription, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "WRKChainDepositPaid", _chainIdRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootWRKChainDepositPaid)
				if err := _WRKChainRoot.contract.UnpackLog(event, "WRKChainDepositPaid", log); err != nil {
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

// WRKChainRootWRKChainDepositRefundIterator is returned from FilterWRKChainDepositRefund and is used to iterate over the raw logs and unpacked data for WRKChainDepositRefund events raised by the WRKChainRoot contract.
type WRKChainRootWRKChainDepositRefundIterator struct {
	Event *WRKChainRootWRKChainDepositRefund // Event containing the contract specifics and raw log

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
func (it *WRKChainRootWRKChainDepositRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRKChainRootWRKChainDepositRefund)
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
		it.Event = new(WRKChainRootWRKChainDepositRefund)
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
func (it *WRKChainRootWRKChainDepositRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRKChainRootWRKChainDepositRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRKChainRootWRKChainDepositRefund represents a WRKChainDepositRefund event raised by the WRKChainRoot contract.
type WRKChainRootWRKChainDepositRefund struct {
	ChainId *big.Int
	Owner   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWRKChainDepositRefund is a free log retrieval operation binding the contract event 0x5fdfd7834e3e713189c97ffe1bc13c31b4e5003b84f2b73cdf2055cfe49bf647.
//
// Solidity: event WRKChainDepositRefund(uint256 indexed _chainId, address indexed _owner, uint256 _amount)
func (_WRKChainRoot *WRKChainRootFilterer) FilterWRKChainDepositRefund(opts *bind.FilterOpts, _chainId []*big.Int, _owner []common.Address) (*WRKChainRootWRKChainDepositRefundIterator, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _WRKChainRoot.contract.FilterLogs(opts, "WRKChainDepositRefund", _chainIdRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return &WRKChainRootWRKChainDepositRefundIterator{contract: _WRKChainRoot.contract, event: "WRKChainDepositRefund", logs: logs, sub: sub}, nil
}

// WatchWRKChainDepositRefund is a free log subscription operation binding the contract event 0x5fdfd7834e3e713189c97ffe1bc13c31b4e5003b84f2b73cdf2055cfe49bf647.
//
// Solidity: event WRKChainDepositRefund(uint256 indexed _chainId, address indexed _owner, uint256 _amount)
func (_WRKChainRoot *WRKChainRootFilterer) WatchWRKChainDepositRefund(opts *bind.WatchOpts, sink chan<- *WRKChainRootWRKChainDepositRefund, _chainId []*big.Int, _owner []common.Address) (event.Subscription, error) {

	var _chainIdRule []interface{}
	for _, _chainIdItem := range _chainId {
		_chainIdRule = append(_chainIdRule, _chainIdItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _WRKChainRoot.contract.WatchLogs(opts, "WRKChainDepositRefund", _chainIdRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRKChainRootWRKChainDepositRefund)
				if err := _WRKChainRoot.contract.UnpackLog(event, "WRKChainDepositRefund", log); err != nil {
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
	ChainId *big.Int
	Old     common.Address
	New     common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWRKChainOwnerChanged is a free log retrieval operation binding the contract event 0xf4fb8ce4b6ba19f99d4309750d2d0bf6ddb2eab6e3ff6f1d2bf27dea39fdb067.
//
// Solidity: event WRKChainOwnerChanged(uint256 indexed _chainId, address _old, address _new)
func (_WRKChainRoot *WRKChainRootFilterer) FilterWRKChainOwnerChanged(opts *bind.FilterOpts, _chainId []*big.Int) (*WRKChainRootWRKChainOwnerChangedIterator, error) {

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

// WatchWRKChainOwnerChanged is a free log subscription operation binding the contract event 0xf4fb8ce4b6ba19f99d4309750d2d0bf6ddb2eab6e3ff6f1d2bf27dea39fdb067.
//
// Solidity: event WRKChainOwnerChanged(uint256 indexed _chainId, address _old, address _new)
func (_WRKChainRoot *WRKChainRootFilterer) WatchWRKChainOwnerChanged(opts *bind.WatchOpts, sink chan<- *WRKChainRootWRKChainOwnerChanged, _chainId []*big.Int) (event.Subscription, error) {

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
