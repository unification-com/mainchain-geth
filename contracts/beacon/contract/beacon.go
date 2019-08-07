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

// BeaconABI is the input ABI used to generate the binding from.
const BeaconABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_beaconId\",\"type\":\"bytes32\"},{\"name\":\"_authAddresses\",\"type\":\"address[]\"}],\"name\":\"registerBeacon\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beaconId\",\"type\":\"bytes32\"},{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"recordHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beaconId\",\"type\":\"bytes32\"}],\"name\":\"getLastHash\",\"outputs\":[{\"name\":\"hashNumIdx_\",\"type\":\"uint256\"},{\"name\":\"lastHashId_\",\"type\":\"bytes32\"},{\"name\":\"lastHash_\",\"type\":\"bytes32\"},{\"name\":\"lastSubmissionTime_\",\"type\":\"uint256\"},{\"name\":\"lastSubmittedBy_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beaconId\",\"type\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isAuthorisedAddress\",\"outputs\":[{\"name\":\"isAuthorised_\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beaconId\",\"type\":\"bytes32\"},{\"name\":\"_addressesToRemove\",\"type\":\"address[]\"}],\"name\":\"removeAuthAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beaconId\",\"type\":\"bytes32\"}],\"name\":\"beaconExists\",\"outputs\":[{\"name\":\"beaconExists_\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beaconId\",\"type\":\"bytes32\"},{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeBeaconOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"beaconList\",\"outputs\":[{\"name\":\"isBeacon\",\"type\":\"bool\"},{\"name\":\"numAuthAddresses\",\"type\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"hashNumIdx\",\"type\":\"uint256\"},{\"name\":\"lastHash\",\"type\":\"bytes32\"},{\"name\":\"lastHashId\",\"type\":\"bytes32\"},{\"name\":\"lastSubmissionTime\",\"type\":\"uint256\"},{\"name\":\"lastSubmittedBy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beaconId\",\"type\":\"bytes32\"},{\"name\":\"_newAuthAddresses\",\"type\":\"address[]\"}],\"name\":\"addAuthAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_beaconId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_hashId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_hashNum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_submittedBy\",\"type\":\"address\"}],\"name\":\"RecordHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_beaconId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"RegisterBeacon\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogFallbackFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_beaconId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_authorisedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"AuthoriseNewAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_beaconId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_removedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"RemoveAuthorisedAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_beaconId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_old\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_new\",\"type\":\"address\"}],\"name\":\"BeaconOwnerChanged\",\"type\":\"event\"}]"

// BeaconFuncSigs maps the 4-byte function signature to its string representation.
var BeaconFuncSigs = map[string]string{
	"d4202677": "addAuthAddresses(bytes32,address[])",
	"92e5ff0a": "beaconExists(bytes32)",
	"be510a44": "beaconList(bytes32)",
	"aad6f07d": "changeBeaconOwner(bytes32,address)",
	"55acd07a": "getLastHash(bytes32)",
	"62269f7a": "isAuthorisedAddress(bytes32,address)",
	"17705412": "recordHash(bytes32,bytes32)",
	"1149e8a8": "registerBeacon(bytes32,address[])",
	"6c5664ac": "removeAuthAddresses(bytes32,address[])",
}

// BeaconBin is the compiled bytecode used for deploying new contracts.
var BeaconBin = "0x608060405234801561001057600080fd5b5061198a806100206000396000f3fe6080604052600436106100865760003560e01c80636c5664ac116100595780636c5664ac1461025257806392e5ff0a14610309578063aad6f07d14610333578063be510a441461036c578063d4202677146103e257610086565b80631149e8a8146100be578063177054121461017757806355acd07a146101a757806362269f7a14610205575b60408051348152905133917fb1db2589d999730e8faec7dccf62c4d28a7f6ef1095b46c9a744dcd46d87130e919081900360200190a2005b3480156100ca57600080fd5b50610175600480360360408110156100e157600080fd5b8135919081019060408101602082013564010000000081111561010357600080fd5b82018360208201111561011557600080fd5b8035906020019184602083028401116401000000008311171561013757600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610499945050505050565b005b34801561018357600080fd5b506101756004803603604081101561019a57600080fd5b508035906020013561080c565b3480156101b357600080fd5b506101d1600480360360208110156101ca57600080fd5b5035610acc565b6040805195865260208601949094528484019290925260608401526001600160a01b03166080830152519081900360a00190f35b34801561021157600080fd5b5061023e6004803603604081101561022857600080fd5b50803590602001356001600160a01b0316610bbb565b604080519115158252519081900360200190f35b34801561025e57600080fd5b506101756004803603604081101561027557600080fd5b8135919081019060408101602082013564010000000081111561029757600080fd5b8201836020820111156102a957600080fd5b803590602001918460208302840111640100000000831117156102cb57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610c90945050505050565b34801561031557600080fd5b5061023e6004803603602081101561032c57600080fd5b503561104a565b34801561033f57600080fd5b506101756004803603604081101561035657600080fd5b50803590602001356001600160a01b03166110ab565b34801561037857600080fd5b506103966004803603602081101561038f57600080fd5b503561136a565b60408051981515895260208901979097526001600160a01b03958616888801526060880194909452608087019290925260a086015260c08501521660e083015251908190036101000190f35b3480156103ee57600080fd5b506101756004803603604081101561040557600080fd5b8135919081019060408101602082013564010000000081111561042757600080fd5b82018360208201111561043957600080fd5b8035906020019184602083028401116401000000008311171561045b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506113bc945050505050565b81806104e1576040805162461bcd60e51b815260206004820152601260248201527110995858dbdb881251081c995c5d5a5c995960721b604482015290519081900360640190fd5b60008381526020819052604090205460ff1615610545576040805162461bcd60e51b815260206004820152601860248201527f426561636f6e20494420616c7265616479206578697374730000000000000000604482015290519081900360640190fd5b60008251116105855760405162461bcd60e51b81526004018080602001828103825260258152602001806118996025913960400191505060405180910390fd5b600a825111156105c65760405162461bcd60e51b81526004018080602001828103825260278152602001806118406027913960400191505060405180910390fd5b60408051610100810182526001815283516020808301918252338385019081526000606085018181526080860182815260a0870183815260c0880184815260e089018581528d8652858852948a90209851895460ff19169015151789559651600289015593516003880180546001600160a01b03199081166001600160a01b039384161790915592516004890155905160058801559251600687015593516007860155516008909401805490931693169290921790558151428152915185927f7f2313114e1b6cdbb36da39db143ad3325e2db40cf5ef22f44b39d49ff8fbba792908290030190a260005b825181101561077b57600160008086815260200190815260200160002060010160008584815181106106df57fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff021916908315150217905550336001600160a01b03168460008051602061190483398151915285848151811061074457fe5b602002602001015160405180826001600160a01b03166001600160a01b0316815260200191505060405180910390a36001016106b1565b5060008381526020818152604080832033845260010190915290205460ff166108075760008381526020818152604080832033845260018181018452918420805460ff191683179055868452929091526002909101546107e09163ffffffff61178116565b50604080513380825291518591600080516020611904833981519152919081900360200190a35b505050565b600082815260208190526040902054829060ff1661086b576040805162461bcd60e51b815260206004820152601760248201527610995858dbdb925108191bd95cc81b9bdd08195e1a5cdd604a1b604482015290519081900360640190fd5b60008181526020818152604080832033845260019081019092529091205460ff16151514806108b357506000818152602081905260409020600301546001600160a01b031633145b6108f5576040805162461bcd60e51b815260206004820152600e60248201526d1b9bdd08185d5d1a1bdc9a5cd95960921b604482015290519081900360640190fd5b828061093d576040805162461bcd60e51b815260206004820152601260248201527110995858dbdb881251081c995c5d5a5c995960721b604482015290519081900360640190fd5b8261097f576040805162461bcd60e51b815260206004820152600d60248201526c1a185cda081c995c5d5a5c9959609a1b604482015290519081900360640190fd5b60008481526020819052604090205460ff166109e2576040805162461bcd60e51b815260206004820152601860248201527f426561636f6e20494420646f6573206e6f742065786973740000000000000000604482015290519081900360640190fd5b600084815260208190526040812060040154429190610a0890600163ffffffff61178116565b6040805160208082018a90528183018490526060808301879052835180840390910181526080830180855281519183019190912060008c8152928390529184902060048101869055600581018b9055600681018390556007810188905560080180546001600160a01b031916339081179091559085905260a0830187905260c083015291519293509091879183918a917f2c43852a4b9af2daed696ea8fd4d00f6865532a47d2cb798161dc91695362622919081900360e00190a450505050505050565b6000808080808580610b1a576040805162461bcd60e51b815260206004820152601260248201527110995858dbdb881251081c995c5d5a5c995960721b604482015290519081900360640190fd5b60008781526020819052604090205460ff16610b77576040805162461bcd60e51b815260206004820152601760248201527610da185a5b88125108191bd95cc81b9bdd08195e1a5cdd604a1b604482015290519081900360640190fd5b5050506000938452505050602081905260409020600481015460068201546005830154600784015460089094015492949193909290916001600160a01b0390911690565b60008280610c05576040805162461bcd60e51b815260206004820152601260248201527110995858dbdb881251081c995c5d5a5c995960721b604482015290519081900360640190fd5b60008481526020819052604090205460ff16610c62576040805162461bcd60e51b815260206004820152601760248201527610da185a5b88125108191bd95cc81b9bdd08195e1a5cdd604a1b604482015290519081900360640190fd5b50506000918252602082815260408084206001600160a01b0390931684526001909201905290205460ff1690565b600082815260208190526040902054829060ff16610cef576040805162461bcd60e51b815260206004820152601760248201527610995858dbdb925108191bd95cc81b9bdd08195e1a5cdd604a1b604482015290519081900360640190fd5b6000818152602081905260409020600301546001600160a01b03163314610d55576040805162461bcd60e51b81526020600482015260156024820152742cb7ba9030b932903737ba103a34329037bbb732b960591b604482015290519081900360640190fd5b8280610d9d576040805162461bcd60e51b815260206004820152601260248201527110995858dbdb881251081c995c5d5a5c995960721b604482015290519081900360640190fd5b6000835111610ded576040805162461bcd60e51b8152602060048201526017602482015276105d5d1a081859191c995cdcd95cc81c995c5d5a5c9959604a1b604482015290519081900360640190fd5b600a83511115610e2e5760405162461bcd60e51b81526004018080602001828103825260238152602001806118e16023913960400191505060405180910390fd5b60008481526020819052604090205460ff16610e8b576040805162461bcd60e51b815260206004820152601760248201527610da185a5b88125108191bd95cc81b9bdd08195e1a5cdd604a1b604482015290519081900360640190fd5b600084815260208190526040902060020154835110610edb5760405162461bcd60e51b81526004018080602001828103825260328152602001806119246032913960400191505060405180910390fd5b60005b83518110156110435760006001600160a01b0316848281518110610efe57fe5b60200260200101516001600160a01b03161415610f62576040805162461bcd60e51b815260206004820152601d60248201527f63616e6e6f7420617574686f72697365207a65726f2061646472657373000000604482015290519081900360640190fd5b6000808681526020019081526020016000206001016000858381518110610f8557fe5b6020908102919091018101516001600160a01b031682528181019290925260409081016000908120805460ff1916905587815291829052902060020154610fd390600163ffffffff6117e216565b50336001600160a01b0316857f1ab6ecd37e23ab30318669bb4642f1cdfba82f23405ce8dfba3648fc9b4744e386848151811061100c57fe5b602002602001015160405180826001600160a01b03166001600160a01b0316815260200191505060405180910390a3600101610ede565b5050505050565b60008180611094576040805162461bcd60e51b815260206004820152601260248201527110995858dbdb881251081c995c5d5a5c995960721b604482015290519081900360640190fd5b505060009081526020819052604090205460ff1690565b600082815260208190526040902054829060ff1661110a576040805162461bcd60e51b815260206004820152601760248201527610995858dbdb925108191bd95cc81b9bdd08195e1a5cdd604a1b604482015290519081900360640190fd5b6000818152602081905260409020600301546001600160a01b03163314611170576040805162461bcd60e51b81526020600482015260156024820152742cb7ba9030b932903737ba103a34329037bbb732b960591b604482015290519081900360640190fd5b82806111b8576040805162461bcd60e51b815260206004820152601260248201527110995858dbdb881251081c995c5d5a5c995960721b604482015290519081900360640190fd5b6001600160a01b0383166111fd5760405162461bcd60e51b81526004018080602001828103825260328152602001806118676032913960400191505060405180910390fd5b60008481526020819052604090205460ff1661125a576040805162461bcd60e51b815260206004820152601760248201527610da185a5b88125108191bd95cc81b9bdd08195e1a5cdd604a1b604482015290519081900360640190fd5b6000848152602081815260409182902060030180546001600160a01b0319166001600160a01b038716908117909155825133815291820152815186927fad85d1ac5a61bdac309481e2184b8d7ee6584b7a9cd6621ec55d1a44ae72c30c928290030190a26000848152602081815260408083206001600160a01b038716845260010190915290205460ff16611364576000848152602081815260408083206001600160a01b038716845260018181018452918420805460ff191683179055878452929091526002909101546113349163ffffffff61178116565b50604080516001600160a01b03851681529051339186916000805160206119048339815191529181900360200190a35b50505050565b6000602081905290815260409020805460028201546003830154600484015460058501546006860154600787015460089097015460ff9096169694956001600160a01b03948516959394929391921688565b600082815260208190526040902054829060ff1661141b576040805162461bcd60e51b815260206004820152601760248201527610995858dbdb925108191bd95cc81b9bdd08195e1a5cdd604a1b604482015290519081900360640190fd5b6000818152602081905260409020600301546001600160a01b03163314611481576040805162461bcd60e51b81526020600482015260156024820152742cb7ba9030b932903737ba103a34329037bbb732b960591b604482015290519081900360640190fd5b82806114c9576040805162461bcd60e51b815260206004820152601260248201527110995858dbdb881251081c995c5d5a5c995960721b604482015290519081900360640190fd5b6000835111611519576040805162461bcd60e51b8152602060048201526017602482015276105d5d1a081859191c995cdcd95cc81c995c5d5a5c9959604a1b604482015290519081900360640190fd5b600a8351111561155a5760405162461bcd60e51b81526004018080602001828103825260238152602001806118e16023913960400191505060405180910390fd5b60008481526020819052604090205460ff166115b7576040805162461bcd60e51b815260206004820152601760248201527610da185a5b88125108191bd95cc81b9bdd08195e1a5cdd604a1b604482015290519081900360640190fd5b825160008581526020819052604081206002015490916115dd919063ffffffff61178116565b9050606481111561161f5760405162461bcd60e51b81526004018080602001828103825260238152602001806118be6023913960400191505060405180910390fd5b60005b84518110156117795760006001600160a01b031685828151811061164257fe5b60200260200101516001600160a01b031614156116a6576040805162461bcd60e51b815260206004820152601d60248201527f63616e6e6f7420617574686f72697365207a65726f2061646472657373000000604482015290519081900360640190fd5b600160008088815260200190815260200160002060010160008784815181106116cb57fe5b6020908102919091018101516001600160a01b031682528181019290925260409081016000908120805460ff1916941515949094179093558883529082905290206002015461171b906001611781565b50336001600160a01b03168660008051602061190483398151915287848151811061174257fe5b602002602001015160405180826001600160a01b03166001600160a01b0316815260200191505060405180910390a3600101611622565b505050505050565b6000828201838110156117db576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600082821115611839576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe4d6178696d756d20313020417574686f72697365642061646472657373657320616c6c6f7765644e6577206f776e65722072657175697265642c20616e642073686f756c64206265206e6f6e2d7a65726f2061646472657373496e697469616c20417574686f7269736564206164647265737365732072657175697265644d617820313020617574686f72697365642061646472657373657320616c6c6f7765644d6178207375626d697373696f6e2031302061646472657373657320616c6c6f776564fc23324c4d8b370287cc3e6d2db8892f41de294bd67a7e3c8ecb88742081aa2d6174206c65617374206f6e652072656d61696e696e6720617574686f72697365642061646472657373207265717569726564a265627a7a723058203f5a18bd6fdd81dc73de655fa439e976203bdd421835a487fdf1545988cebf1864736f6c634300050a0032"

// DeployBeacon deploys a new Ethereum contract, binding an instance of Beacon to it.
func DeployBeacon(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Beacon, error) {
	parsed, err := abi.JSON(strings.NewReader(BeaconABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BeaconBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Beacon{BeaconCaller: BeaconCaller{contract: contract}, BeaconTransactor: BeaconTransactor{contract: contract}, BeaconFilterer: BeaconFilterer{contract: contract}}, nil
}

// Beacon is an auto generated Go binding around an Ethereum contract.
type Beacon struct {
	BeaconCaller     // Read-only binding to the contract
	BeaconTransactor // Write-only binding to the contract
	BeaconFilterer   // Log filterer for contract events
}

// BeaconCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeaconCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeaconTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeaconFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeaconSession struct {
	Contract     *Beacon           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeaconCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeaconCallerSession struct {
	Contract *BeaconCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BeaconTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeaconTransactorSession struct {
	Contract     *BeaconTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeaconRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeaconRaw struct {
	Contract *Beacon // Generic contract binding to access the raw methods on
}

// BeaconCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeaconCallerRaw struct {
	Contract *BeaconCaller // Generic read-only contract binding to access the raw methods on
}

// BeaconTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeaconTransactorRaw struct {
	Contract *BeaconTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeacon creates a new instance of Beacon, bound to a specific deployed contract.
func NewBeacon(address common.Address, backend bind.ContractBackend) (*Beacon, error) {
	contract, err := bindBeacon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Beacon{BeaconCaller: BeaconCaller{contract: contract}, BeaconTransactor: BeaconTransactor{contract: contract}, BeaconFilterer: BeaconFilterer{contract: contract}}, nil
}

// NewBeaconCaller creates a new read-only instance of Beacon, bound to a specific deployed contract.
func NewBeaconCaller(address common.Address, caller bind.ContractCaller) (*BeaconCaller, error) {
	contract, err := bindBeacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconCaller{contract: contract}, nil
}

// NewBeaconTransactor creates a new write-only instance of Beacon, bound to a specific deployed contract.
func NewBeaconTransactor(address common.Address, transactor bind.ContractTransactor) (*BeaconTransactor, error) {
	contract, err := bindBeacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconTransactor{contract: contract}, nil
}

// NewBeaconFilterer creates a new log filterer instance of Beacon, bound to a specific deployed contract.
func NewBeaconFilterer(address common.Address, filterer bind.ContractFilterer) (*BeaconFilterer, error) {
	contract, err := bindBeacon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeaconFilterer{contract: contract}, nil
}

// bindBeacon binds a generic wrapper to an already deployed contract.
func bindBeacon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BeaconABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Beacon *BeaconRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Beacon.Contract.BeaconCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Beacon *BeaconRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Beacon.Contract.BeaconTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Beacon *BeaconRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Beacon.Contract.BeaconTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Beacon *BeaconCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Beacon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Beacon *BeaconTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Beacon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Beacon *BeaconTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Beacon.Contract.contract.Transact(opts, method, params...)
}

// BeaconExists is a free data retrieval call binding the contract method 0x92e5ff0a.
//
// Solidity: function beaconExists(bytes32 _beaconId) constant returns(bool beaconExists_)
func (_Beacon *BeaconCaller) BeaconExists(opts *bind.CallOpts, _beaconId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Beacon.contract.Call(opts, out, "beaconExists", _beaconId)
	return *ret0, err
}

// BeaconExists is a free data retrieval call binding the contract method 0x92e5ff0a.
//
// Solidity: function beaconExists(bytes32 _beaconId) constant returns(bool beaconExists_)
func (_Beacon *BeaconSession) BeaconExists(_beaconId [32]byte) (bool, error) {
	return _Beacon.Contract.BeaconExists(&_Beacon.CallOpts, _beaconId)
}

// BeaconExists is a free data retrieval call binding the contract method 0x92e5ff0a.
//
// Solidity: function beaconExists(bytes32 _beaconId) constant returns(bool beaconExists_)
func (_Beacon *BeaconCallerSession) BeaconExists(_beaconId [32]byte) (bool, error) {
	return _Beacon.Contract.BeaconExists(&_Beacon.CallOpts, _beaconId)
}

// BeaconList is a free data retrieval call binding the contract method 0xbe510a44.
//
// Solidity: function beaconList(bytes32 ) constant returns(bool isBeacon, uint256 numAuthAddresses, address owner, uint256 hashNumIdx, bytes32 lastHash, bytes32 lastHashId, uint256 lastSubmissionTime, address lastSubmittedBy)
func (_Beacon *BeaconCaller) BeaconList(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsBeacon           bool
	NumAuthAddresses   *big.Int
	Owner              common.Address
	HashNumIdx         *big.Int
	LastHash           [32]byte
	LastHashId         [32]byte
	LastSubmissionTime *big.Int
	LastSubmittedBy    common.Address
}, error) {
	ret := new(struct {
		IsBeacon           bool
		NumAuthAddresses   *big.Int
		Owner              common.Address
		HashNumIdx         *big.Int
		LastHash           [32]byte
		LastHashId         [32]byte
		LastSubmissionTime *big.Int
		LastSubmittedBy    common.Address
	})
	out := ret
	err := _Beacon.contract.Call(opts, out, "beaconList", arg0)
	return *ret, err
}

// BeaconList is a free data retrieval call binding the contract method 0xbe510a44.
//
// Solidity: function beaconList(bytes32 ) constant returns(bool isBeacon, uint256 numAuthAddresses, address owner, uint256 hashNumIdx, bytes32 lastHash, bytes32 lastHashId, uint256 lastSubmissionTime, address lastSubmittedBy)
func (_Beacon *BeaconSession) BeaconList(arg0 [32]byte) (struct {
	IsBeacon           bool
	NumAuthAddresses   *big.Int
	Owner              common.Address
	HashNumIdx         *big.Int
	LastHash           [32]byte
	LastHashId         [32]byte
	LastSubmissionTime *big.Int
	LastSubmittedBy    common.Address
}, error) {
	return _Beacon.Contract.BeaconList(&_Beacon.CallOpts, arg0)
}

// BeaconList is a free data retrieval call binding the contract method 0xbe510a44.
//
// Solidity: function beaconList(bytes32 ) constant returns(bool isBeacon, uint256 numAuthAddresses, address owner, uint256 hashNumIdx, bytes32 lastHash, bytes32 lastHashId, uint256 lastSubmissionTime, address lastSubmittedBy)
func (_Beacon *BeaconCallerSession) BeaconList(arg0 [32]byte) (struct {
	IsBeacon           bool
	NumAuthAddresses   *big.Int
	Owner              common.Address
	HashNumIdx         *big.Int
	LastHash           [32]byte
	LastHashId         [32]byte
	LastSubmissionTime *big.Int
	LastSubmittedBy    common.Address
}, error) {
	return _Beacon.Contract.BeaconList(&_Beacon.CallOpts, arg0)
}

// GetLastHash is a free data retrieval call binding the contract method 0x55acd07a.
//
// Solidity: function getLastHash(bytes32 _beaconId) constant returns(uint256 hashNumIdx_, bytes32 lastHashId_, bytes32 lastHash_, uint256 lastSubmissionTime_, address lastSubmittedBy_)
func (_Beacon *BeaconCaller) GetLastHash(opts *bind.CallOpts, _beaconId [32]byte) (struct {
	HashNumIdx         *big.Int
	LastHashId         [32]byte
	LastHash           [32]byte
	LastSubmissionTime *big.Int
	LastSubmittedBy    common.Address
}, error) {
	ret := new(struct {
		HashNumIdx         *big.Int
		LastHashId         [32]byte
		LastHash           [32]byte
		LastSubmissionTime *big.Int
		LastSubmittedBy    common.Address
	})
	out := ret
	err := _Beacon.contract.Call(opts, out, "getLastHash", _beaconId)
	return *ret, err
}

// GetLastHash is a free data retrieval call binding the contract method 0x55acd07a.
//
// Solidity: function getLastHash(bytes32 _beaconId) constant returns(uint256 hashNumIdx_, bytes32 lastHashId_, bytes32 lastHash_, uint256 lastSubmissionTime_, address lastSubmittedBy_)
func (_Beacon *BeaconSession) GetLastHash(_beaconId [32]byte) (struct {
	HashNumIdx         *big.Int
	LastHashId         [32]byte
	LastHash           [32]byte
	LastSubmissionTime *big.Int
	LastSubmittedBy    common.Address
}, error) {
	return _Beacon.Contract.GetLastHash(&_Beacon.CallOpts, _beaconId)
}

// GetLastHash is a free data retrieval call binding the contract method 0x55acd07a.
//
// Solidity: function getLastHash(bytes32 _beaconId) constant returns(uint256 hashNumIdx_, bytes32 lastHashId_, bytes32 lastHash_, uint256 lastSubmissionTime_, address lastSubmittedBy_)
func (_Beacon *BeaconCallerSession) GetLastHash(_beaconId [32]byte) (struct {
	HashNumIdx         *big.Int
	LastHashId         [32]byte
	LastHash           [32]byte
	LastSubmissionTime *big.Int
	LastSubmittedBy    common.Address
}, error) {
	return _Beacon.Contract.GetLastHash(&_Beacon.CallOpts, _beaconId)
}

// IsAuthorisedAddress is a free data retrieval call binding the contract method 0x62269f7a.
//
// Solidity: function isAuthorisedAddress(bytes32 _beaconId, address _address) constant returns(bool isAuthorised_)
func (_Beacon *BeaconCaller) IsAuthorisedAddress(opts *bind.CallOpts, _beaconId [32]byte, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Beacon.contract.Call(opts, out, "isAuthorisedAddress", _beaconId, _address)
	return *ret0, err
}

// IsAuthorisedAddress is a free data retrieval call binding the contract method 0x62269f7a.
//
// Solidity: function isAuthorisedAddress(bytes32 _beaconId, address _address) constant returns(bool isAuthorised_)
func (_Beacon *BeaconSession) IsAuthorisedAddress(_beaconId [32]byte, _address common.Address) (bool, error) {
	return _Beacon.Contract.IsAuthorisedAddress(&_Beacon.CallOpts, _beaconId, _address)
}

// IsAuthorisedAddress is a free data retrieval call binding the contract method 0x62269f7a.
//
// Solidity: function isAuthorisedAddress(bytes32 _beaconId, address _address) constant returns(bool isAuthorised_)
func (_Beacon *BeaconCallerSession) IsAuthorisedAddress(_beaconId [32]byte, _address common.Address) (bool, error) {
	return _Beacon.Contract.IsAuthorisedAddress(&_Beacon.CallOpts, _beaconId, _address)
}

// AddAuthAddresses is a paid mutator transaction binding the contract method 0xd4202677.
//
// Solidity: function addAuthAddresses(bytes32 _beaconId, address[] _newAuthAddresses) returns()
func (_Beacon *BeaconTransactor) AddAuthAddresses(opts *bind.TransactOpts, _beaconId [32]byte, _newAuthAddresses []common.Address) (*types.Transaction, error) {
	return _Beacon.contract.Transact(opts, "addAuthAddresses", _beaconId, _newAuthAddresses)
}

// AddAuthAddresses is a paid mutator transaction binding the contract method 0xd4202677.
//
// Solidity: function addAuthAddresses(bytes32 _beaconId, address[] _newAuthAddresses) returns()
func (_Beacon *BeaconSession) AddAuthAddresses(_beaconId [32]byte, _newAuthAddresses []common.Address) (*types.Transaction, error) {
	return _Beacon.Contract.AddAuthAddresses(&_Beacon.TransactOpts, _beaconId, _newAuthAddresses)
}

// AddAuthAddresses is a paid mutator transaction binding the contract method 0xd4202677.
//
// Solidity: function addAuthAddresses(bytes32 _beaconId, address[] _newAuthAddresses) returns()
func (_Beacon *BeaconTransactorSession) AddAuthAddresses(_beaconId [32]byte, _newAuthAddresses []common.Address) (*types.Transaction, error) {
	return _Beacon.Contract.AddAuthAddresses(&_Beacon.TransactOpts, _beaconId, _newAuthAddresses)
}

// ChangeBeaconOwner is a paid mutator transaction binding the contract method 0xaad6f07d.
//
// Solidity: function changeBeaconOwner(bytes32 _beaconId, address _newOwner) returns()
func (_Beacon *BeaconTransactor) ChangeBeaconOwner(opts *bind.TransactOpts, _beaconId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _Beacon.contract.Transact(opts, "changeBeaconOwner", _beaconId, _newOwner)
}

// ChangeBeaconOwner is a paid mutator transaction binding the contract method 0xaad6f07d.
//
// Solidity: function changeBeaconOwner(bytes32 _beaconId, address _newOwner) returns()
func (_Beacon *BeaconSession) ChangeBeaconOwner(_beaconId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _Beacon.Contract.ChangeBeaconOwner(&_Beacon.TransactOpts, _beaconId, _newOwner)
}

// ChangeBeaconOwner is a paid mutator transaction binding the contract method 0xaad6f07d.
//
// Solidity: function changeBeaconOwner(bytes32 _beaconId, address _newOwner) returns()
func (_Beacon *BeaconTransactorSession) ChangeBeaconOwner(_beaconId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _Beacon.Contract.ChangeBeaconOwner(&_Beacon.TransactOpts, _beaconId, _newOwner)
}

// RecordHash is a paid mutator transaction binding the contract method 0x17705412.
//
// Solidity: function recordHash(bytes32 _beaconId, bytes32 _hash) returns()
func (_Beacon *BeaconTransactor) RecordHash(opts *bind.TransactOpts, _beaconId [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _Beacon.contract.Transact(opts, "recordHash", _beaconId, _hash)
}

// RecordHash is a paid mutator transaction binding the contract method 0x17705412.
//
// Solidity: function recordHash(bytes32 _beaconId, bytes32 _hash) returns()
func (_Beacon *BeaconSession) RecordHash(_beaconId [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _Beacon.Contract.RecordHash(&_Beacon.TransactOpts, _beaconId, _hash)
}

// RecordHash is a paid mutator transaction binding the contract method 0x17705412.
//
// Solidity: function recordHash(bytes32 _beaconId, bytes32 _hash) returns()
func (_Beacon *BeaconTransactorSession) RecordHash(_beaconId [32]byte, _hash [32]byte) (*types.Transaction, error) {
	return _Beacon.Contract.RecordHash(&_Beacon.TransactOpts, _beaconId, _hash)
}

// RegisterBeacon is a paid mutator transaction binding the contract method 0x1149e8a8.
//
// Solidity: function registerBeacon(bytes32 _beaconId, address[] _authAddresses) returns()
func (_Beacon *BeaconTransactor) RegisterBeacon(opts *bind.TransactOpts, _beaconId [32]byte, _authAddresses []common.Address) (*types.Transaction, error) {
	return _Beacon.contract.Transact(opts, "registerBeacon", _beaconId, _authAddresses)
}

// RegisterBeacon is a paid mutator transaction binding the contract method 0x1149e8a8.
//
// Solidity: function registerBeacon(bytes32 _beaconId, address[] _authAddresses) returns()
func (_Beacon *BeaconSession) RegisterBeacon(_beaconId [32]byte, _authAddresses []common.Address) (*types.Transaction, error) {
	return _Beacon.Contract.RegisterBeacon(&_Beacon.TransactOpts, _beaconId, _authAddresses)
}

// RegisterBeacon is a paid mutator transaction binding the contract method 0x1149e8a8.
//
// Solidity: function registerBeacon(bytes32 _beaconId, address[] _authAddresses) returns()
func (_Beacon *BeaconTransactorSession) RegisterBeacon(_beaconId [32]byte, _authAddresses []common.Address) (*types.Transaction, error) {
	return _Beacon.Contract.RegisterBeacon(&_Beacon.TransactOpts, _beaconId, _authAddresses)
}

// RemoveAuthAddresses is a paid mutator transaction binding the contract method 0x6c5664ac.
//
// Solidity: function removeAuthAddresses(bytes32 _beaconId, address[] _addressesToRemove) returns()
func (_Beacon *BeaconTransactor) RemoveAuthAddresses(opts *bind.TransactOpts, _beaconId [32]byte, _addressesToRemove []common.Address) (*types.Transaction, error) {
	return _Beacon.contract.Transact(opts, "removeAuthAddresses", _beaconId, _addressesToRemove)
}

// RemoveAuthAddresses is a paid mutator transaction binding the contract method 0x6c5664ac.
//
// Solidity: function removeAuthAddresses(bytes32 _beaconId, address[] _addressesToRemove) returns()
func (_Beacon *BeaconSession) RemoveAuthAddresses(_beaconId [32]byte, _addressesToRemove []common.Address) (*types.Transaction, error) {
	return _Beacon.Contract.RemoveAuthAddresses(&_Beacon.TransactOpts, _beaconId, _addressesToRemove)
}

// RemoveAuthAddresses is a paid mutator transaction binding the contract method 0x6c5664ac.
//
// Solidity: function removeAuthAddresses(bytes32 _beaconId, address[] _addressesToRemove) returns()
func (_Beacon *BeaconTransactorSession) RemoveAuthAddresses(_beaconId [32]byte, _addressesToRemove []common.Address) (*types.Transaction, error) {
	return _Beacon.Contract.RemoveAuthAddresses(&_Beacon.TransactOpts, _beaconId, _addressesToRemove)
}

// BeaconAuthoriseNewAddressIterator is returned from FilterAuthoriseNewAddress and is used to iterate over the raw logs and unpacked data for AuthoriseNewAddress events raised by the Beacon contract.
type BeaconAuthoriseNewAddressIterator struct {
	Event *BeaconAuthoriseNewAddress // Event containing the contract specifics and raw log

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
func (it *BeaconAuthoriseNewAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconAuthoriseNewAddress)
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
		it.Event = new(BeaconAuthoriseNewAddress)
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
func (it *BeaconAuthoriseNewAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconAuthoriseNewAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconAuthoriseNewAddress represents a AuthoriseNewAddress event raised by the Beacon contract.
type BeaconAuthoriseNewAddress struct {
	BeaconId     [32]byte
	AuthorisedBy common.Address
	Address      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuthoriseNewAddress is a free log retrieval operation binding the contract event 0xfc23324c4d8b370287cc3e6d2db8892f41de294bd67a7e3c8ecb88742081aa2d.
//
// Solidity: event AuthoriseNewAddress(bytes32 indexed _beaconId, address indexed _authorisedBy, address _address)
func (_Beacon *BeaconFilterer) FilterAuthoriseNewAddress(opts *bind.FilterOpts, _beaconId [][32]byte, _authorisedBy []common.Address) (*BeaconAuthoriseNewAddressIterator, error) {

	var _beaconIdRule []interface{}
	for _, _beaconIdItem := range _beaconId {
		_beaconIdRule = append(_beaconIdRule, _beaconIdItem)
	}
	var _authorisedByRule []interface{}
	for _, _authorisedByItem := range _authorisedBy {
		_authorisedByRule = append(_authorisedByRule, _authorisedByItem)
	}

	logs, sub, err := _Beacon.contract.FilterLogs(opts, "AuthoriseNewAddress", _beaconIdRule, _authorisedByRule)
	if err != nil {
		return nil, err
	}
	return &BeaconAuthoriseNewAddressIterator{contract: _Beacon.contract, event: "AuthoriseNewAddress", logs: logs, sub: sub}, nil
}

// WatchAuthoriseNewAddress is a free log subscription operation binding the contract event 0xfc23324c4d8b370287cc3e6d2db8892f41de294bd67a7e3c8ecb88742081aa2d.
//
// Solidity: event AuthoriseNewAddress(bytes32 indexed _beaconId, address indexed _authorisedBy, address _address)
func (_Beacon *BeaconFilterer) WatchAuthoriseNewAddress(opts *bind.WatchOpts, sink chan<- *BeaconAuthoriseNewAddress, _beaconId [][32]byte, _authorisedBy []common.Address) (event.Subscription, error) {

	var _beaconIdRule []interface{}
	for _, _beaconIdItem := range _beaconId {
		_beaconIdRule = append(_beaconIdRule, _beaconIdItem)
	}
	var _authorisedByRule []interface{}
	for _, _authorisedByItem := range _authorisedBy {
		_authorisedByRule = append(_authorisedByRule, _authorisedByItem)
	}

	logs, sub, err := _Beacon.contract.WatchLogs(opts, "AuthoriseNewAddress", _beaconIdRule, _authorisedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconAuthoriseNewAddress)
				if err := _Beacon.contract.UnpackLog(event, "AuthoriseNewAddress", log); err != nil {
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
// Solidity: event AuthoriseNewAddress(bytes32 indexed _beaconId, address indexed _authorisedBy, address _address)
func (_Beacon *BeaconFilterer) ParseAuthoriseNewAddress(log types.Log) (*BeaconAuthoriseNewAddress, error) {
	event := new(BeaconAuthoriseNewAddress)
	if err := _Beacon.contract.UnpackLog(event, "AuthoriseNewAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BeaconBeaconOwnerChangedIterator is returned from FilterBeaconOwnerChanged and is used to iterate over the raw logs and unpacked data for BeaconOwnerChanged events raised by the Beacon contract.
type BeaconBeaconOwnerChangedIterator struct {
	Event *BeaconBeaconOwnerChanged // Event containing the contract specifics and raw log

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
func (it *BeaconBeaconOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconBeaconOwnerChanged)
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
		it.Event = new(BeaconBeaconOwnerChanged)
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
func (it *BeaconBeaconOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconBeaconOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconBeaconOwnerChanged represents a BeaconOwnerChanged event raised by the Beacon contract.
type BeaconBeaconOwnerChanged struct {
	BeaconId [32]byte
	Old      common.Address
	New      common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBeaconOwnerChanged is a free log retrieval operation binding the contract event 0xad85d1ac5a61bdac309481e2184b8d7ee6584b7a9cd6621ec55d1a44ae72c30c.
//
// Solidity: event BeaconOwnerChanged(bytes32 indexed _beaconId, address _old, address _new)
func (_Beacon *BeaconFilterer) FilterBeaconOwnerChanged(opts *bind.FilterOpts, _beaconId [][32]byte) (*BeaconBeaconOwnerChangedIterator, error) {

	var _beaconIdRule []interface{}
	for _, _beaconIdItem := range _beaconId {
		_beaconIdRule = append(_beaconIdRule, _beaconIdItem)
	}

	logs, sub, err := _Beacon.contract.FilterLogs(opts, "BeaconOwnerChanged", _beaconIdRule)
	if err != nil {
		return nil, err
	}
	return &BeaconBeaconOwnerChangedIterator{contract: _Beacon.contract, event: "BeaconOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchBeaconOwnerChanged is a free log subscription operation binding the contract event 0xad85d1ac5a61bdac309481e2184b8d7ee6584b7a9cd6621ec55d1a44ae72c30c.
//
// Solidity: event BeaconOwnerChanged(bytes32 indexed _beaconId, address _old, address _new)
func (_Beacon *BeaconFilterer) WatchBeaconOwnerChanged(opts *bind.WatchOpts, sink chan<- *BeaconBeaconOwnerChanged, _beaconId [][32]byte) (event.Subscription, error) {

	var _beaconIdRule []interface{}
	for _, _beaconIdItem := range _beaconId {
		_beaconIdRule = append(_beaconIdRule, _beaconIdItem)
	}

	logs, sub, err := _Beacon.contract.WatchLogs(opts, "BeaconOwnerChanged", _beaconIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconBeaconOwnerChanged)
				if err := _Beacon.contract.UnpackLog(event, "BeaconOwnerChanged", log); err != nil {
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

// ParseBeaconOwnerChanged is a log parse operation binding the contract event 0xad85d1ac5a61bdac309481e2184b8d7ee6584b7a9cd6621ec55d1a44ae72c30c.
//
// Solidity: event BeaconOwnerChanged(bytes32 indexed _beaconId, address _old, address _new)
func (_Beacon *BeaconFilterer) ParseBeaconOwnerChanged(log types.Log) (*BeaconBeaconOwnerChanged, error) {
	event := new(BeaconBeaconOwnerChanged)
	if err := _Beacon.contract.UnpackLog(event, "BeaconOwnerChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BeaconLogFallbackFunctionCalledIterator is returned from FilterLogFallbackFunctionCalled and is used to iterate over the raw logs and unpacked data for LogFallbackFunctionCalled events raised by the Beacon contract.
type BeaconLogFallbackFunctionCalledIterator struct {
	Event *BeaconLogFallbackFunctionCalled // Event containing the contract specifics and raw log

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
func (it *BeaconLogFallbackFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconLogFallbackFunctionCalled)
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
		it.Event = new(BeaconLogFallbackFunctionCalled)
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
func (it *BeaconLogFallbackFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconLogFallbackFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconLogFallbackFunctionCalled represents a LogFallbackFunctionCalled event raised by the Beacon contract.
type BeaconLogFallbackFunctionCalled struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogFallbackFunctionCalled is a free log retrieval operation binding the contract event 0xb1db2589d999730e8faec7dccf62c4d28a7f6ef1095b46c9a744dcd46d87130e.
//
// Solidity: event LogFallbackFunctionCalled(address indexed _from, uint256 _amount)
func (_Beacon *BeaconFilterer) FilterLogFallbackFunctionCalled(opts *bind.FilterOpts, _from []common.Address) (*BeaconLogFallbackFunctionCalledIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Beacon.contract.FilterLogs(opts, "LogFallbackFunctionCalled", _fromRule)
	if err != nil {
		return nil, err
	}
	return &BeaconLogFallbackFunctionCalledIterator{contract: _Beacon.contract, event: "LogFallbackFunctionCalled", logs: logs, sub: sub}, nil
}

// WatchLogFallbackFunctionCalled is a free log subscription operation binding the contract event 0xb1db2589d999730e8faec7dccf62c4d28a7f6ef1095b46c9a744dcd46d87130e.
//
// Solidity: event LogFallbackFunctionCalled(address indexed _from, uint256 _amount)
func (_Beacon *BeaconFilterer) WatchLogFallbackFunctionCalled(opts *bind.WatchOpts, sink chan<- *BeaconLogFallbackFunctionCalled, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Beacon.contract.WatchLogs(opts, "LogFallbackFunctionCalled", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconLogFallbackFunctionCalled)
				if err := _Beacon.contract.UnpackLog(event, "LogFallbackFunctionCalled", log); err != nil {
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
func (_Beacon *BeaconFilterer) ParseLogFallbackFunctionCalled(log types.Log) (*BeaconLogFallbackFunctionCalled, error) {
	event := new(BeaconLogFallbackFunctionCalled)
	if err := _Beacon.contract.UnpackLog(event, "LogFallbackFunctionCalled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BeaconRecordHashIterator is returned from FilterRecordHash and is used to iterate over the raw logs and unpacked data for RecordHash events raised by the Beacon contract.
type BeaconRecordHashIterator struct {
	Event *BeaconRecordHash // Event containing the contract specifics and raw log

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
func (it *BeaconRecordHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconRecordHash)
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
		it.Event = new(BeaconRecordHash)
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
func (it *BeaconRecordHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconRecordHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconRecordHash represents a RecordHash event raised by the Beacon contract.
type BeaconRecordHash struct {
	BeaconId    [32]byte
	HashId      [32]byte
	Hash        [32]byte
	HashNum     *big.Int
	Timestamp   *big.Int
	SubmittedBy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRecordHash is a free log retrieval operation binding the contract event 0x2c43852a4b9af2daed696ea8fd4d00f6865532a47d2cb798161dc91695362622.
//
// Solidity: event RecordHash(bytes32 indexed _beaconId, bytes32 indexed _hashId, bytes32 indexed _hash, uint256 _hashNum, uint256 _timestamp, address _submittedBy)
func (_Beacon *BeaconFilterer) FilterRecordHash(opts *bind.FilterOpts, _beaconId [][32]byte, _hashId [][32]byte, _hash [][32]byte) (*BeaconRecordHashIterator, error) {

	var _beaconIdRule []interface{}
	for _, _beaconIdItem := range _beaconId {
		_beaconIdRule = append(_beaconIdRule, _beaconIdItem)
	}
	var _hashIdRule []interface{}
	for _, _hashIdItem := range _hashId {
		_hashIdRule = append(_hashIdRule, _hashIdItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _Beacon.contract.FilterLogs(opts, "RecordHash", _beaconIdRule, _hashIdRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return &BeaconRecordHashIterator{contract: _Beacon.contract, event: "RecordHash", logs: logs, sub: sub}, nil
}

// WatchRecordHash is a free log subscription operation binding the contract event 0x2c43852a4b9af2daed696ea8fd4d00f6865532a47d2cb798161dc91695362622.
//
// Solidity: event RecordHash(bytes32 indexed _beaconId, bytes32 indexed _hashId, bytes32 indexed _hash, uint256 _hashNum, uint256 _timestamp, address _submittedBy)
func (_Beacon *BeaconFilterer) WatchRecordHash(opts *bind.WatchOpts, sink chan<- *BeaconRecordHash, _beaconId [][32]byte, _hashId [][32]byte, _hash [][32]byte) (event.Subscription, error) {

	var _beaconIdRule []interface{}
	for _, _beaconIdItem := range _beaconId {
		_beaconIdRule = append(_beaconIdRule, _beaconIdItem)
	}
	var _hashIdRule []interface{}
	for _, _hashIdItem := range _hashId {
		_hashIdRule = append(_hashIdRule, _hashIdItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _Beacon.contract.WatchLogs(opts, "RecordHash", _beaconIdRule, _hashIdRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconRecordHash)
				if err := _Beacon.contract.UnpackLog(event, "RecordHash", log); err != nil {
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

// ParseRecordHash is a log parse operation binding the contract event 0x2c43852a4b9af2daed696ea8fd4d00f6865532a47d2cb798161dc91695362622.
//
// Solidity: event RecordHash(bytes32 indexed _beaconId, bytes32 indexed _hashId, bytes32 indexed _hash, uint256 _hashNum, uint256 _timestamp, address _submittedBy)
func (_Beacon *BeaconFilterer) ParseRecordHash(log types.Log) (*BeaconRecordHash, error) {
	event := new(BeaconRecordHash)
	if err := _Beacon.contract.UnpackLog(event, "RecordHash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BeaconRegisterBeaconIterator is returned from FilterRegisterBeacon and is used to iterate over the raw logs and unpacked data for RegisterBeacon events raised by the Beacon contract.
type BeaconRegisterBeaconIterator struct {
	Event *BeaconRegisterBeacon // Event containing the contract specifics and raw log

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
func (it *BeaconRegisterBeaconIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconRegisterBeacon)
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
		it.Event = new(BeaconRegisterBeacon)
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
func (it *BeaconRegisterBeaconIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconRegisterBeaconIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconRegisterBeacon represents a RegisterBeacon event raised by the Beacon contract.
type BeaconRegisterBeacon struct {
	BeaconId  [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRegisterBeacon is a free log retrieval operation binding the contract event 0x7f2313114e1b6cdbb36da39db143ad3325e2db40cf5ef22f44b39d49ff8fbba7.
//
// Solidity: event RegisterBeacon(bytes32 indexed _beaconId, uint256 _timestamp)
func (_Beacon *BeaconFilterer) FilterRegisterBeacon(opts *bind.FilterOpts, _beaconId [][32]byte) (*BeaconRegisterBeaconIterator, error) {

	var _beaconIdRule []interface{}
	for _, _beaconIdItem := range _beaconId {
		_beaconIdRule = append(_beaconIdRule, _beaconIdItem)
	}

	logs, sub, err := _Beacon.contract.FilterLogs(opts, "RegisterBeacon", _beaconIdRule)
	if err != nil {
		return nil, err
	}
	return &BeaconRegisterBeaconIterator{contract: _Beacon.contract, event: "RegisterBeacon", logs: logs, sub: sub}, nil
}

// WatchRegisterBeacon is a free log subscription operation binding the contract event 0x7f2313114e1b6cdbb36da39db143ad3325e2db40cf5ef22f44b39d49ff8fbba7.
//
// Solidity: event RegisterBeacon(bytes32 indexed _beaconId, uint256 _timestamp)
func (_Beacon *BeaconFilterer) WatchRegisterBeacon(opts *bind.WatchOpts, sink chan<- *BeaconRegisterBeacon, _beaconId [][32]byte) (event.Subscription, error) {

	var _beaconIdRule []interface{}
	for _, _beaconIdItem := range _beaconId {
		_beaconIdRule = append(_beaconIdRule, _beaconIdItem)
	}

	logs, sub, err := _Beacon.contract.WatchLogs(opts, "RegisterBeacon", _beaconIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconRegisterBeacon)
				if err := _Beacon.contract.UnpackLog(event, "RegisterBeacon", log); err != nil {
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

// ParseRegisterBeacon is a log parse operation binding the contract event 0x7f2313114e1b6cdbb36da39db143ad3325e2db40cf5ef22f44b39d49ff8fbba7.
//
// Solidity: event RegisterBeacon(bytes32 indexed _beaconId, uint256 _timestamp)
func (_Beacon *BeaconFilterer) ParseRegisterBeacon(log types.Log) (*BeaconRegisterBeacon, error) {
	event := new(BeaconRegisterBeacon)
	if err := _Beacon.contract.UnpackLog(event, "RegisterBeacon", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BeaconRemoveAuthorisedAddressIterator is returned from FilterRemoveAuthorisedAddress and is used to iterate over the raw logs and unpacked data for RemoveAuthorisedAddress events raised by the Beacon contract.
type BeaconRemoveAuthorisedAddressIterator struct {
	Event *BeaconRemoveAuthorisedAddress // Event containing the contract specifics and raw log

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
func (it *BeaconRemoveAuthorisedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconRemoveAuthorisedAddress)
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
		it.Event = new(BeaconRemoveAuthorisedAddress)
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
func (it *BeaconRemoveAuthorisedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconRemoveAuthorisedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconRemoveAuthorisedAddress represents a RemoveAuthorisedAddress event raised by the Beacon contract.
type BeaconRemoveAuthorisedAddress struct {
	BeaconId  [32]byte
	RemovedBy common.Address
	Address   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemoveAuthorisedAddress is a free log retrieval operation binding the contract event 0x1ab6ecd37e23ab30318669bb4642f1cdfba82f23405ce8dfba3648fc9b4744e3.
//
// Solidity: event RemoveAuthorisedAddress(bytes32 indexed _beaconId, address indexed _removedBy, address _address)
func (_Beacon *BeaconFilterer) FilterRemoveAuthorisedAddress(opts *bind.FilterOpts, _beaconId [][32]byte, _removedBy []common.Address) (*BeaconRemoveAuthorisedAddressIterator, error) {

	var _beaconIdRule []interface{}
	for _, _beaconIdItem := range _beaconId {
		_beaconIdRule = append(_beaconIdRule, _beaconIdItem)
	}
	var _removedByRule []interface{}
	for _, _removedByItem := range _removedBy {
		_removedByRule = append(_removedByRule, _removedByItem)
	}

	logs, sub, err := _Beacon.contract.FilterLogs(opts, "RemoveAuthorisedAddress", _beaconIdRule, _removedByRule)
	if err != nil {
		return nil, err
	}
	return &BeaconRemoveAuthorisedAddressIterator{contract: _Beacon.contract, event: "RemoveAuthorisedAddress", logs: logs, sub: sub}, nil
}

// WatchRemoveAuthorisedAddress is a free log subscription operation binding the contract event 0x1ab6ecd37e23ab30318669bb4642f1cdfba82f23405ce8dfba3648fc9b4744e3.
//
// Solidity: event RemoveAuthorisedAddress(bytes32 indexed _beaconId, address indexed _removedBy, address _address)
func (_Beacon *BeaconFilterer) WatchRemoveAuthorisedAddress(opts *bind.WatchOpts, sink chan<- *BeaconRemoveAuthorisedAddress, _beaconId [][32]byte, _removedBy []common.Address) (event.Subscription, error) {

	var _beaconIdRule []interface{}
	for _, _beaconIdItem := range _beaconId {
		_beaconIdRule = append(_beaconIdRule, _beaconIdItem)
	}
	var _removedByRule []interface{}
	for _, _removedByItem := range _removedBy {
		_removedByRule = append(_removedByRule, _removedByItem)
	}

	logs, sub, err := _Beacon.contract.WatchLogs(opts, "RemoveAuthorisedAddress", _beaconIdRule, _removedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconRemoveAuthorisedAddress)
				if err := _Beacon.contract.UnpackLog(event, "RemoveAuthorisedAddress", log); err != nil {
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
// Solidity: event RemoveAuthorisedAddress(bytes32 indexed _beaconId, address indexed _removedBy, address _address)
func (_Beacon *BeaconFilterer) ParseRemoveAuthorisedAddress(log types.Log) (*BeaconRemoveAuthorisedAddress, error) {
	event := new(BeaconRemoveAuthorisedAddress)
	if err := _Beacon.contract.UnpackLog(event, "RemoveAuthorisedAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

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
