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

// DSGContractABI is the input ABI used to generate the binding from.
const DSGContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"x49\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x85\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x14\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x44\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x93\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x59\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x18\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x75\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x24\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x45\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x12\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x78\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x63\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x7\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x40\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x80\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x39\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x62\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x17\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x30\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x23\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x43\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x79\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x46\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x89\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x94\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x61\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x82\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x90\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x70\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x35\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x83\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x25\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x22\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x31\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x47\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x65\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x77\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x84\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x55\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x10\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x71\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x76\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x48\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x51\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x60\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x13\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x21\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x32\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x88\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x42\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x73\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x64\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x91\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x66\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x38\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x16\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x58\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x95\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x72\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x86\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x54\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x81\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x36\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x50\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x29\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x37\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x74\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x27\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x67\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x53\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x15\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x33\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x4\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x5\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rotate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x92\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_ALLOWED_STAKE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x8\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x26\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x69\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x6\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x19\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x57\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x87\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x41\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x11\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x52\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x68\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x34\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x56\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x28\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAllowedStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"_stakeUnstake\",\"type\":\"uint8\"}],\"name\":\"StakingEvent\",\"type\":\"event\"}]"

// DSGContractFuncSigs maps the 4-byte function signature to its string representation.
var DSGContractFuncSigs = map[string]string{
	"df83f579": "MIN_ALLOWED_STAKE()",
	"8ab66a90": "candidates(address)",
	"399080ec": "getStaked(address)",
	"d992818d": "rotate()",
	"3a4b66f1": "stake()",
	"817b1cd2": "totalStaked()",
	"5d3eea91": "unStake(uint256)",
	"343943bd": "x1()",
	"7d38e0ff": "x10()",
	"f25e3949": "x11()",
	"1d0fb81a": "x12()",
	"8bd149b3": "x13()",
	"05f04234": "x14()",
	"c0a1a949": "x15()",
	"a4f930d8": "x16()",
	"37545b00": "x17()",
	"14846dd3": "x18()",
	"eaecb012": "x19()",
	"4ff13571": "x2()",
	"9ffa6a9f": "x20()",
	"8cadd7f6": "x21()",
	"6bc3d2b8": "x22()",
	"3f3f3444": "x23()",
	"17bd117a": "x24()",
	"68937ad5": "x25()",
	"e749773f": "x26()",
	"bc8c3c9e": "x27()",
	"fe1878ef": "x28()",
	"b54d94f8": "x29()",
	"a3b78873": "x3()",
	"39cb8435": "x30()",
	"72c019b3": "x31()",
	"908df6b1": "x32()",
	"c5f720e4": "x33()",
	"fd56a650": "x34()",
	"5e217f42": "x35()",
	"b407a0db": "x36()",
	"b9fbc543": "x37()",
	"a19f482a": "x38()",
	"3176cca2": "x39()",
	"cb6cee8a": "x4()",
	"2ca07afd": "x40()",
	"f24ec5c8": "x41()",
	"943a0169": "x42()",
	"402cc27e": "x43()",
	"0ae490ad": "x44()",
	"1c073086": "x45()",
	"449136cf": "x46()",
	"72e1b839": "x47()",
	"7fda3be2": "x48()",
	"0418043d": "x49()",
	"cd145f65": "x5()",
	"b4feb110": "x50()",
	"88ce3426": "x51()",
	"f304e720": "x52()",
	"bcfbbc46": "x53()",
	"af594904": "x54()",
	"7af6e874": "x55()",
	"fddbd387": "x56()",
	"eb5d92a0": "x57()",
	"a87f6ab5": "x58()",
	"13813beb": "x59()",
	"eaeb89f7": "x6()",
	"89b3bc3a": "x60()",
	"588d1150": "x61()",
	"332bb13b": "x62()",
	"1fc58dae": "x63()",
	"9bf3f61a": "x64()",
	"745be928": "x65()",
	"a11be9af": "x66()",
	"bca63eb4": "x67()",
	"f721734f": "x68()",
	"eac637f2": "x69()",
	"282940a7": "x7()",
	"5d370399": "x70()",
	"7ede4904": "x71()",
	"ac1e43a0": "x72()",
	"9b85e451": "x73()",
	"ba16cefe": "x74()",
	"175d756b": "x75()",
	"7f758a15": "x76()",
	"74ec408e": "x77()",
	"1e0a08db": "x78()",
	"41f01d3c": "x79()",
	"e66e4c3f": "x8()",
	"2e3f53c9": "x80()",
	"b02d7daa": "x81()",
	"5ae05570": "x82()",
	"60c72ae2": "x83()",
	"772eda28": "x84()",
	"04ba000d": "x85()",
	"ac95e388": "x86()",
	"ebc4e3d4": "x87()",
	"92770c96": "x88()",
	"457f2f2e": "x89()",
	"d6efd153": "x9()",
	"5c356ab5": "x90()",
	"9c856be7": "x91()",
	"da9f47ad": "x92()",
	"0f99cd32": "x93()",
	"5880ab1b": "x94()",
	"aa5b743d": "x95()",
}

// DSGContractBin is the compiled bytecode used for deploying new contracts.
var DSGContractBin = "0x608060405234801561001057600080fd5b506040516124ba3803806124ba8339818101604052602081101561003357600080fd5b5051806100a157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f6d696e20616c6c6f776564207374616b65206d757374206265203e2030000000604482015290519081900360640190fd5b600055612407806100b36000396000f3fe6080604052600436106105c55760003560e01c806389b3bc3a116102f3578063ba16cefe1161019b578063e749773f116100e7578063f24ec5c8116100a0578063f721734f1161007a578063f721734f14610ecb578063fd56a65014610ee0578063fddbd38714610ef5578063fe1878ef14610f0a576105c5565b8063f24ec5c814610e8c578063f25e394914610ea1578063f304e72014610eb6576105c5565b8063e749773f14610e0e578063eac637f214610e23578063eaeb89f714610e38578063eaecb01214610e4d578063eb5d92a014610e62578063ebc4e3d414610e77576105c5565b8063cb6cee8a11610154578063d992818d1161012e578063d992818d14610dba578063da9f47ad14610dcf578063df83f57914610de4578063e66e4c3f14610df9576105c5565b8063cb6cee8a14610d7b578063cd145f6514610d90578063d6efd15314610da5576105c5565b8063ba16cefe14610cfd578063bc8c3c9e14610d12578063bca63eb414610d27578063bcfbbc4614610d3c578063c0a1a94914610d51578063c5f720e414610d66576105c5565b8063a19f482a1161025a578063ac95e38811610213578063b407a0db116101ed578063b407a0db14610ca9578063b4feb11014610cbe578063b54d94f814610cd3578063b9fbc54314610ce8576105c5565b8063ac95e38814610c6a578063af59490414610c7f578063b02d7daa14610c94576105c5565b8063a19f482a14610bec578063a3b7887314610c01578063a4f930d814610c16578063a87f6ab514610c2b578063aa5b743d14610c40578063ac1e43a014610c55576105c5565b8063943a0169116102ac578063943a016914610b6e5780639b85e45114610b835780639bf3f61a14610b985780639c856be714610bad5780639ffa6a9f14610bc2578063a11be9af14610bd7576105c5565b806389b3bc3a14610ad25780638ab66a9014610ae75780638bd149b314610b1a5780638cadd7f614610b2f578063908df6b114610b4457806392770c9614610b59576105c5565b806341f01d3c116104715780636bc3d2b8116103bd5780637af6e874116103765780637f758a15116103505780637f758a1514610a7e5780637fda3be214610a93578063817b1cd214610aa857806388ce342614610abd576105c5565b80637af6e87414610a3f5780637d38e0ff14610a545780637ede490414610a69576105c5565b80636bc3d2b8146109c157806372c019b3146109d657806372e1b839146109eb578063745be92814610a0057806374ec408e14610a15578063772eda2814610a2a576105c5565b80635ae055701161042a5780635d3eea91116104045780635d3eea91146109585780635e217f421461098257806360c72ae21461099757806368937ad5146109ac576105c5565b80635ae05570146109195780635c356ab51461092e5780635d37039914610943576105c5565b806341f01d3c1461089b578063449136cf146108b0578063457f2f2e146108c55780634ff13571146108da5780635880ab1b146108ef578063588d115014610904576105c5565b80631fc58dae11610530578063343943bd116104e957806339cb8435116104c357806339cb8435146108545780633a4b66f1146108695780633f3f344414610871578063402cc27e14610886576105c5565b8063343943bd146107e557806337545b00146107fa578063399080ec1461080f576105c5565b80631fc58dae14610767578063282940a71461077c5780632ca07afd146107915780632e3f53c9146107a65780633176cca2146107bb578063332bb13b146107d0576105c5565b806314846dd31161058257806314846dd3146106e9578063175d756b146106fe57806317bd117a146107135780631c073086146107285780631d0fb81a1461073d5780631e0a08db14610752576105c5565b80630418043d1461064f57806304ba000d1461068057806305f04234146106955780630ae490ad146106aa5780630f99cd32146106bf57806313813beb146106d4575b341561064d57336000908152600260205260409020546105eb903463ffffffff610f1f16565b3360009081526002602052604090205560015461060e903463ffffffff610f1f16565b600190815560408051348152426020820152815133927f3585198d07e828811559e83f38a04e95c328af4edd28122306b10475843fdd35928290030190a35b005b34801561065b57600080fd5b50610664610f80565b604080516001600160a01b039092168252519081900360200190f35b34801561068c57600080fd5b50610664610f8f565b3480156106a157600080fd5b50610664610f9e565b3480156106b657600080fd5b50610664610fad565b3480156106cb57600080fd5b50610664610fbc565b3480156106e057600080fd5b50610664610fcb565b3480156106f557600080fd5b50610664610fda565b34801561070a57600080fd5b50610664610fe9565b34801561071f57600080fd5b50610664610ff8565b34801561073457600080fd5b50610664611007565b34801561074957600080fd5b50610664611016565b34801561075e57600080fd5b50610664611025565b34801561077357600080fd5b50610664611034565b34801561078857600080fd5b50610664611043565b34801561079d57600080fd5b50610664611052565b3480156107b257600080fd5b50610664611061565b3480156107c757600080fd5b50610664611070565b3480156107dc57600080fd5b5061066461107f565b3480156107f157600080fd5b5061066461108e565b34801561080657600080fd5b5061066461109d565b34801561081b57600080fd5b506108426004803603602081101561083257600080fd5b50356001600160a01b03166110ac565b60408051918252519081900360200190f35b34801561086057600080fd5b506106646110c7565b61064d6110d6565b34801561087d57600080fd5b506106646111b0565b34801561089257600080fd5b506106646111bf565b3480156108a757600080fd5b506106646111ce565b3480156108bc57600080fd5b506106646111dd565b3480156108d157600080fd5b506106646111ec565b3480156108e657600080fd5b506106646111fb565b3480156108fb57600080fd5b5061066461120a565b34801561091057600080fd5b50610664611219565b34801561092557600080fd5b50610664611228565b34801561093a57600080fd5b50610664611237565b34801561094f57600080fd5b50610664611246565b34801561096457600080fd5b5061064d6004803603602081101561097b57600080fd5b5035611255565b34801561098e57600080fd5b506106646113ef565b3480156109a357600080fd5b506106646113fe565b3480156109b857600080fd5b5061066461140d565b3480156109cd57600080fd5b5061066461141c565b3480156109e257600080fd5b5061066461142b565b3480156109f757600080fd5b5061066461143a565b348015610a0c57600080fd5b50610664611449565b348015610a2157600080fd5b50610664611458565b348015610a3657600080fd5b50610664611467565b348015610a4b57600080fd5b50610664611476565b348015610a6057600080fd5b50610664611485565b348015610a7557600080fd5b50610664611494565b348015610a8a57600080fd5b506106646114a3565b348015610a9f57600080fd5b506106646114b2565b348015610ab457600080fd5b506108426114c1565b348015610ac957600080fd5b506106646114c7565b348015610ade57600080fd5b506106646114d6565b348015610af357600080fd5b5061084260048036036020811015610b0a57600080fd5b50356001600160a01b03166114e5565b348015610b2657600080fd5b506106646114f7565b348015610b3b57600080fd5b50610664611506565b348015610b5057600080fd5b50610664611515565b348015610b6557600080fd5b50610664611524565b348015610b7a57600080fd5b50610664611533565b348015610b8f57600080fd5b50610664611542565b348015610ba457600080fd5b50610664611551565b348015610bb957600080fd5b50610664611560565b348015610bce57600080fd5b5061066461156f565b348015610be357600080fd5b5061066461157e565b348015610bf857600080fd5b5061066461158d565b348015610c0d57600080fd5b5061066461159c565b348015610c2257600080fd5b506106646115ab565b348015610c3757600080fd5b506106646115ba565b348015610c4c57600080fd5b506106646115c9565b348015610c6157600080fd5b506106646115d8565b348015610c7657600080fd5b506106646115e7565b348015610c8b57600080fd5b506106646115f6565b348015610ca057600080fd5b50610664611605565b348015610cb557600080fd5b50610664611614565b348015610cca57600080fd5b50610664611623565b348015610cdf57600080fd5b50610664611632565b348015610cf457600080fd5b50610664611641565b348015610d0957600080fd5b50610664611650565b348015610d1e57600080fd5b5061066461165f565b348015610d3357600080fd5b5061066461166e565b348015610d4857600080fd5b5061066461167d565b348015610d5d57600080fd5b5061066461168c565b348015610d7257600080fd5b5061066461169b565b348015610d8757600080fd5b506106646116aa565b348015610d9c57600080fd5b506106646116b9565b348015610db157600080fd5b506106646116c8565b348015610dc657600080fd5b5061064d6116d7565b348015610ddb57600080fd5b5061066461221f565b348015610df057600080fd5b5061084261222e565b348015610e0557600080fd5b50610664612234565b348015610e1a57600080fd5b50610664612243565b348015610e2f57600080fd5b50610664612252565b348015610e4457600080fd5b50610664612261565b348015610e5957600080fd5b50610664612270565b348015610e6e57600080fd5b5061066461227f565b348015610e8357600080fd5b5061066461228e565b348015610e9857600080fd5b5061066461229d565b348015610ead57600080fd5b506106646122ac565b348015610ec257600080fd5b506106646122bb565b348015610ed757600080fd5b506106646122ca565b348015610eec57600080fd5b506106646122d9565b348015610f0157600080fd5b506106646122e8565b348015610f1657600080fd5b506106646122f7565b600082820183811015610f79576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6033546001600160a01b031681565b6057546001600160a01b031681565b6010546001600160a01b031681565b602e546001600160a01b031681565b605f546001600160a01b031681565b603d546001600160a01b031681565b6014546001600160a01b031681565b604d546001600160a01b031681565b601a546001600160a01b031681565b602f546001600160a01b031681565b600e546001600160a01b031681565b6050546001600160a01b031681565b6041546001600160a01b031681565b6009546001600160a01b031681565b602a546001600160a01b031681565b6052546001600160a01b031681565b6029546001600160a01b031681565b6040546001600160a01b031681565b6003546001600160a01b031681565b6013546001600160a01b031681565b6001600160a01b031660009081526002602052604090205490565b6020546001600160a01b031681565b60005434101561112d576040805162461bcd60e51b815260206004820152601e60248201527f6d696e696d756d207374616b696e6720616d6f756e74206e6f74206d65740000604482015290519081900360640190fd5b3360009081526002602052604090205461114d903463ffffffff610f1f16565b33600090815260026020526040902055600154611170903463ffffffff610f1f16565b600190815560408051348152426020820152815133927f3585198d07e828811559e83f38a04e95c328af4edd28122306b10475843fdd35928290030190a3565b6019546001600160a01b031681565b602d546001600160a01b031681565b6051546001600160a01b031681565b6030546001600160a01b031681565b605b546001600160a01b031681565b6004546001600160a01b031681565b6060546001600160a01b031681565b603f546001600160a01b031681565b6054546001600160a01b031681565b605c546001600160a01b031681565b6048546001600160a01b031681565b600081116112aa576040805162461bcd60e51b815260206004820152601d60248201527f616d6f756e7420746f20756e7374616b65206d757374206265203e2030000000604482015290519081900360640190fd5b336000908152600260205260409020548111156112f85760405162461bcd60e51b815260040180806020018281038252603a815260200180612399603a913960400191505060405180910390fd5b6001548111156113395760405162461bcd60e51b81526004018080602001828103825260358152602001806123646035913960400191505060405180910390fd5b33600090815260026020526040902054611359908263ffffffff61230616565b3360009081526002602052604090205560015461137c908263ffffffff61230616565b60015560408051828152426020820152815160009233927f3585198d07e828811559e83f38a04e95c328af4edd28122306b10475843fdd35929081900390910190a3604051339082156108fc029083906000818181858888f193505050501580156113eb573d6000803e3d6000fd5b5050565b6025546001600160a01b031681565b6055546001600160a01b031681565b601b546001600160a01b031681565b6018546001600160a01b031681565b6021546001600160a01b031681565b6031546001600160a01b031681565b6043546001600160a01b031681565b604f546001600160a01b031681565b6056546001600160a01b031681565b6039546001600160a01b031681565b600c546001600160a01b031681565b6049546001600160a01b031681565b604e546001600160a01b031681565b6032546001600160a01b031681565b60015481565b6035546001600160a01b031681565b603e546001600160a01b031681565b60026020526000908152604090205481565b600f546001600160a01b031681565b6017546001600160a01b031681565b6022546001600160a01b031681565b605a546001600160a01b031681565b602c546001600160a01b031681565b604b546001600160a01b031681565b6042546001600160a01b031681565b605d546001600160a01b031681565b6016546001600160a01b031681565b6044546001600160a01b031681565b6028546001600160a01b031681565b6005546001600160a01b031681565b6012546001600160a01b031681565b603c546001600160a01b031681565b6061546001600160a01b031681565b604a546001600160a01b031681565b6058546001600160a01b031681565b6038546001600160a01b031681565b6053546001600160a01b031681565b6026546001600160a01b031681565b6034546001600160a01b031681565b601f546001600160a01b031681565b6027546001600160a01b031681565b604c546001600160a01b031681565b601d546001600160a01b031681565b6045546001600160a01b031681565b6037546001600160a01b031681565b6011546001600160a01b031681565b6023546001600160a01b031681565b6006546001600160a01b031681565b6007546001600160a01b031681565b600b546001600160a01b031681565b73f30f951b0426f8bf37ac71967407081358df7a7b33146116f75761221d565b600380546001600160a01b0319908116722a956804bad8dcad148abff71515f9b057f7e017909155600480548216723adc30a6f4db59d2698e3d3029fd1ba68b6b15179055600580548216724a435f1d54aa5cc9fcfa0feb6b8c4a428bbb93179055600680548216725aa882b1708f2e032b1e6ac8fb92fb2e8f09e2179055600780548216726a60f879db155cb01dcc70452c0a935fcb15d3179055600880548216727aa694e0fd812b5f3d6d42654c8f8f33b58a8a179055600980548216728a1c6df0be8682dfb87f8d1265f90f1d181647179055600a80548216729a2cbe70409e8ab9f2fee621354578916f0118179055600b8054821673010a7c34c6beb24b2a92febd123d780788cf2235179055600c8054821673011a6b9f30767ad1266ea9d3307b91470f6e22cd179055600d8054821673012ab4f14f3be13d4649be31ca4d65947ef95554179055600e8054821673013a0e389b4ccb117b0809fc1ab4a1b4b3f18f1f179055600f8054821673014aa8687ca6b10afd7aaf101efa215ab52ca07517905560108054821673015ac19e712a6b839ffd14acacbe58151911b91617905560118054821673016a4575a058d7f0c73641472f6ba9aa09d9ebc617905560128054821673017a31bf0827b5280bfcf2e2863a80ef4bf0315217905560138054821673018a31e3f69e9c30086e03d73654edff06171d9617905560148054821673019afd07eb5a3198597b4845db932254fc11d9b917905560158054821673020a618a1bb1c751bd9d7c985edd4f3d9c5fc1a017905560168054821673021a8866723a67828747f0239d3123098568e86e17905560178054821673022ab4c0b02912dd100be27646b594849373e84b17905560188054821673023af3313253a7b2338992042a21479041820b4417905560198054821673024afc0f101527bd62d86c85ab18769b9084af74179055601a8054821673025ad5dc2e38368ca489de9bafc99ea845fa93dd179055601b8054821673026abcd44b611511f2950fc45a78be446ed21afb179055601c8054821673027ac37befdaae0b9fa316a3d1436a5a6256ac16179055601d8054821673028ac2879de90937d22544ba893b1ba04e161e19179055601e8054821673029af5c22f1375ddaeb41f0ac96893de2259b68b179055601f8054821673030a728b7e2ad365d2fdde05c0e536c8f31a5c6f17905560208054821673031a2b49f84bccd3c18f35ddaa07d72970412a6717905560218054821673032aae65c9f7135576f2e58a23b8430b30e0034817905560228054821673033a10b8526e7c8e70d5b41e31f254a060e0ec2d17905560238054821673034a7f2494ecb251cac89933da04dec1709e856517905560248054821673035a60510dbd42f3a65d58fb990066e9b355adc617905560258054821673036af29b42d58f44bf7ca73d3a9b18f657794cfe17905560268054821673037aa30cd556dfb4d51a8880790c077737dd95c717905560278054821673038a0e937609f695a552f1f3ae6c213f287f84aa17905560288054821673039a28d44b5dad2d3d4772da3d2e0d083a18647b17905560298054821673040a3dcfabdf555797feb37bc1e60315753930a1179055602a8054821673041a92ec07212f45b4887c800edfd7c1a25c6f81179055602b8054821673042a956ec294b85e639a43bea2c93ff8c78bac28179055602c8054821673043a6bde9d509712f0d93cf9196e62badf6afaf2179055602d8054821673044ac1a25dd8d6e78a272f0e7b7798aae593972c179055602e8054821673045aed6663d3b239fe31477521c8ad962570d05f179055602f8054821673046aa3f2a00562fea28e354a8b28a3753941797317905560308054821673047a91571a4725efa09981cf4e829fa6670e611917905560318054821673048a6c75dc7c0cc5df4562458bee72ba3b08419517905560328054821673049a47f51837d54a1c40e91469935c880f7af72017905560338054821673050a7673bace7c43787798b8dd280ddd4a17ef1217905560348054821673051a24373f5ad334e2292b336b497097b702f7b317905560358054821673052a987fa6fc78a0b2324bcb47db73fda2ec076117905560368054821673053ae7420dbcc57cb03856c35b9e7ceb764735e917905560378054821673054ab96c0886050a20317ee8cfad593c3a17f5c617905560388054821673055a1668669a3f1e3602f1c346cf6ea5ead81fe417905560398054821673056ae2baea2b1572d9293ef8c688d81a84790ca4179055603a8054821673057adc1e8ccf640767aa7526cd46b6559b88de0c179055603b8054821673058af6bb396153079d6d4cd9bceafc87cf5c97a7179055603c8054821673059acc4768532d8d1888ef4d1ae6c5fc55e1285e179055603d8054821673060ad483059b2d55542168bd13746f4bc9d90800179055603e8054821673061a99d798b34e91f94696f6a704bc0545526964179055603f8054821673062a06c8a3ad35dd60bcaaf403cbdd140034e2f917905560408054821673063ab9229016eca057e026b551bd587f3196ea5817905560418054821673064a975ffd46ee34ac5f43c87d43fd6ef55f8fd317905560428054821673065af52373eded1a2db013e5fe0c49b34325c3aa17905560438054821673066ad8c53e8b458290a01525da6dafe127eb37b017905560448054821673067af3fbb6a30c45b4c2855877315b6d528674b117905560458054821673068afdc63964f3cfd6e61615bfd48ce0c61c292b17905560468054821673069a729be61423a08f1bfa4471ad203c5bc8ef0c17905560478054821673070a3b781c3d5ad9fbc8b616c3d1a683d0bf7abc17905560488054821673071a9130bd8501c80658c2336ae08a30eb723a0a17905560498054821673072a6c85f39e9c97603ff6162d5ef053e112ee05179055604a8054821673073a5d1cc708b853ff7fc05cd1eb6ce066a2d337179055604b8054821673074a7eec2044540d8c031c817636966e5d96d113179055604c8054821673075ad3e9fa874fd60d4d058750c2457f8c40d9af179055604d8054821673076ac70c5605893b0cb224b115053e654c2c11bb179055604e8054821673077a938f9e52b643d47ef24fb203ea7206573f00179055604f8054821673078a6d1ce5abcff8cd7a979b67b193391975e3dd17905560508054821673079a9b6d74b3b515ca141cab5c6789358721358b17905560518054821673080addb3ed8fa2c583b627bdb86d63780623ddd117905560528054821673081a9af6fad95002971ef7775b547e828223f03c17905560538054821673082a93c3c886429e4ea5d274a96851705c00b7b917905560548054821673083ab59ce9fce039c1546219855ace5ec472decf17905560558054821673084a5bd013c104229161958359cfd9da63f8551317905560568054821673085ac9744d7afb313edff60a370cf0b97a5dff8417905560578054821673086a902bbc478db69f6b36b54d5d5954f19b893f17905560588054821673087a204b4a5841550181af660a07daf11914b12b17905560598054821673088a4555c3bad6b770fd2b6b6f112d86b0179b3c179055605a8054821673089adb67d2aa720bb6d37ac49784ee25b4eb3adf179055605b8054821673090a40813e7e679026f10534fdbf10d216c5315b179055605c8054821673091ad8b556a910123abdc002de39269cdf198e1d179055605d8054821673092a93e3ef9d79df3145487448388f42863e1104179055605e8054821673093ac3d0ae17afef0badeea471ec6466511e9c42179055605f8054821673094af16bd3700e2ee4709776da4c3334f6163c6917905560608054821673095a2be3334c23eb8d2a7d463027dfd0cdbe21b117905560618054909116721a320943d4535e93d31e4a65a6e21c5df375d71790555b565b605e546001600160a01b031681565b60005481565b600a546001600160a01b031681565b601c546001600160a01b031681565b6047546001600160a01b031681565b6008546001600160a01b031681565b6015546001600160a01b031681565b603b546001600160a01b031681565b6059546001600160a01b031681565b602b546001600160a01b031681565b600d546001600160a01b031681565b6036546001600160a01b031681565b6046546001600160a01b031681565b6024546001600160a01b031681565b603a546001600160a01b031681565b601e546001600160a01b031681565b60008282111561235d576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe756e7374616b65206661696c6564202d2063616e6e6f7420756e7374616b65206d6f7265207468616e20746f74616c5374616b6564756e7374616b65206661696c6564202d20617474656d7074696e6720746f20756e7374616b65206d6f7265207468616e206973207374616b6564a265627a7a72315820fb1b94b98df364947ab8b568c7f43f849c3ca360f416d8bf6ff975ed98ee793464736f6c634300050b0032"

// DeployDSGContract deploys a new Ethereum contract, binding an instance of DSGContract to it.
func DeployDSGContract(auth *bind.TransactOpts, backend bind.ContractBackend, _minAllowedStake *big.Int) (common.Address, *types.Transaction, *DSGContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSGContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSGContractBin), backend, _minAllowedStake)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSGContract{DSGContractCaller: DSGContractCaller{contract: contract}, DSGContractTransactor: DSGContractTransactor{contract: contract}, DSGContractFilterer: DSGContractFilterer{contract: contract}}, nil
}

// DSGContract is an auto generated Go binding around an Ethereum contract.
type DSGContract struct {
	DSGContractCaller     // Read-only binding to the contract
	DSGContractTransactor // Write-only binding to the contract
	DSGContractFilterer   // Log filterer for contract events
}

// DSGContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSGContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSGContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSGContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSGContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSGContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSGContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSGContractSession struct {
	Contract     *DSGContract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSGContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSGContractCallerSession struct {
	Contract *DSGContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DSGContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSGContractTransactorSession struct {
	Contract     *DSGContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DSGContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSGContractRaw struct {
	Contract *DSGContract // Generic contract binding to access the raw methods on
}

// DSGContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSGContractCallerRaw struct {
	Contract *DSGContractCaller // Generic read-only contract binding to access the raw methods on
}

// DSGContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSGContractTransactorRaw struct {
	Contract *DSGContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSGContract creates a new instance of DSGContract, bound to a specific deployed contract.
func NewDSGContract(address common.Address, backend bind.ContractBackend) (*DSGContract, error) {
	contract, err := bindDSGContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSGContract{DSGContractCaller: DSGContractCaller{contract: contract}, DSGContractTransactor: DSGContractTransactor{contract: contract}, DSGContractFilterer: DSGContractFilterer{contract: contract}}, nil
}

// NewDSGContractCaller creates a new read-only instance of DSGContract, bound to a specific deployed contract.
func NewDSGContractCaller(address common.Address, caller bind.ContractCaller) (*DSGContractCaller, error) {
	contract, err := bindDSGContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSGContractCaller{contract: contract}, nil
}

// NewDSGContractTransactor creates a new write-only instance of DSGContract, bound to a specific deployed contract.
func NewDSGContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DSGContractTransactor, error) {
	contract, err := bindDSGContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSGContractTransactor{contract: contract}, nil
}

// NewDSGContractFilterer creates a new log filterer instance of DSGContract, bound to a specific deployed contract.
func NewDSGContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DSGContractFilterer, error) {
	contract, err := bindDSGContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSGContractFilterer{contract: contract}, nil
}

// bindDSGContract binds a generic wrapper to an already deployed contract.
func bindDSGContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSGContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSGContract *DSGContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSGContract.Contract.DSGContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSGContract *DSGContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSGContract.Contract.DSGContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSGContract *DSGContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSGContract.Contract.DSGContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSGContract *DSGContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSGContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSGContract *DSGContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSGContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSGContract *DSGContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSGContract.Contract.contract.Transact(opts, method, params...)
}

// MINALLOWEDSTAKE is a free data retrieval call binding the contract method 0xdf83f579.
//
// Solidity: function MIN_ALLOWED_STAKE() constant returns(uint256)
func (_DSGContract *DSGContractCaller) MINALLOWEDSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "MIN_ALLOWED_STAKE")
	return *ret0, err
}

// MINALLOWEDSTAKE is a free data retrieval call binding the contract method 0xdf83f579.
//
// Solidity: function MIN_ALLOWED_STAKE() constant returns(uint256)
func (_DSGContract *DSGContractSession) MINALLOWEDSTAKE() (*big.Int, error) {
	return _DSGContract.Contract.MINALLOWEDSTAKE(&_DSGContract.CallOpts)
}

// MINALLOWEDSTAKE is a free data retrieval call binding the contract method 0xdf83f579.
//
// Solidity: function MIN_ALLOWED_STAKE() constant returns(uint256)
func (_DSGContract *DSGContractCallerSession) MINALLOWEDSTAKE() (*big.Int, error) {
	return _DSGContract.Contract.MINALLOWEDSTAKE(&_DSGContract.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x8ab66a90.
//
// Solidity: function candidates(address ) constant returns(uint256)
func (_DSGContract *DSGContractCaller) Candidates(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "candidates", arg0)
	return *ret0, err
}

// Candidates is a free data retrieval call binding the contract method 0x8ab66a90.
//
// Solidity: function candidates(address ) constant returns(uint256)
func (_DSGContract *DSGContractSession) Candidates(arg0 common.Address) (*big.Int, error) {
	return _DSGContract.Contract.Candidates(&_DSGContract.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x8ab66a90.
//
// Solidity: function candidates(address ) constant returns(uint256)
func (_DSGContract *DSGContractCallerSession) Candidates(arg0 common.Address) (*big.Int, error) {
	return _DSGContract.Contract.Candidates(&_DSGContract.CallOpts, arg0)
}

// GetStaked is a free data retrieval call binding the contract method 0x399080ec.
//
// Solidity: function getStaked(address _staker) constant returns(uint256 _amount)
func (_DSGContract *DSGContractCaller) GetStaked(opts *bind.CallOpts, _staker common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "getStaked", _staker)
	return *ret0, err
}

// GetStaked is a free data retrieval call binding the contract method 0x399080ec.
//
// Solidity: function getStaked(address _staker) constant returns(uint256 _amount)
func (_DSGContract *DSGContractSession) GetStaked(_staker common.Address) (*big.Int, error) {
	return _DSGContract.Contract.GetStaked(&_DSGContract.CallOpts, _staker)
}

// GetStaked is a free data retrieval call binding the contract method 0x399080ec.
//
// Solidity: function getStaked(address _staker) constant returns(uint256 _amount)
func (_DSGContract *DSGContractCallerSession) GetStaked(_staker common.Address) (*big.Int, error) {
	return _DSGContract.Contract.GetStaked(&_DSGContract.CallOpts, _staker)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_DSGContract *DSGContractCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "totalStaked")
	return *ret0, err
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_DSGContract *DSGContractSession) TotalStaked() (*big.Int, error) {
	return _DSGContract.Contract.TotalStaked(&_DSGContract.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_DSGContract *DSGContractCallerSession) TotalStaked() (*big.Int, error) {
	return _DSGContract.Contract.TotalStaked(&_DSGContract.CallOpts)
}

// X1 is a free data retrieval call binding the contract method 0x343943bd.
//
// Solidity: function x1() constant returns(address)
func (_DSGContract *DSGContractCaller) X1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x1")
	return *ret0, err
}

// X1 is a free data retrieval call binding the contract method 0x343943bd.
//
// Solidity: function x1() constant returns(address)
func (_DSGContract *DSGContractSession) X1() (common.Address, error) {
	return _DSGContract.Contract.X1(&_DSGContract.CallOpts)
}

// X1 is a free data retrieval call binding the contract method 0x343943bd.
//
// Solidity: function x1() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X1() (common.Address, error) {
	return _DSGContract.Contract.X1(&_DSGContract.CallOpts)
}

// X10 is a free data retrieval call binding the contract method 0x7d38e0ff.
//
// Solidity: function x10() constant returns(address)
func (_DSGContract *DSGContractCaller) X10(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x10")
	return *ret0, err
}

// X10 is a free data retrieval call binding the contract method 0x7d38e0ff.
//
// Solidity: function x10() constant returns(address)
func (_DSGContract *DSGContractSession) X10() (common.Address, error) {
	return _DSGContract.Contract.X10(&_DSGContract.CallOpts)
}

// X10 is a free data retrieval call binding the contract method 0x7d38e0ff.
//
// Solidity: function x10() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X10() (common.Address, error) {
	return _DSGContract.Contract.X10(&_DSGContract.CallOpts)
}

// X11 is a free data retrieval call binding the contract method 0xf25e3949.
//
// Solidity: function x11() constant returns(address)
func (_DSGContract *DSGContractCaller) X11(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x11")
	return *ret0, err
}

// X11 is a free data retrieval call binding the contract method 0xf25e3949.
//
// Solidity: function x11() constant returns(address)
func (_DSGContract *DSGContractSession) X11() (common.Address, error) {
	return _DSGContract.Contract.X11(&_DSGContract.CallOpts)
}

// X11 is a free data retrieval call binding the contract method 0xf25e3949.
//
// Solidity: function x11() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X11() (common.Address, error) {
	return _DSGContract.Contract.X11(&_DSGContract.CallOpts)
}

// X12 is a free data retrieval call binding the contract method 0x1d0fb81a.
//
// Solidity: function x12() constant returns(address)
func (_DSGContract *DSGContractCaller) X12(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x12")
	return *ret0, err
}

// X12 is a free data retrieval call binding the contract method 0x1d0fb81a.
//
// Solidity: function x12() constant returns(address)
func (_DSGContract *DSGContractSession) X12() (common.Address, error) {
	return _DSGContract.Contract.X12(&_DSGContract.CallOpts)
}

// X12 is a free data retrieval call binding the contract method 0x1d0fb81a.
//
// Solidity: function x12() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X12() (common.Address, error) {
	return _DSGContract.Contract.X12(&_DSGContract.CallOpts)
}

// X13 is a free data retrieval call binding the contract method 0x8bd149b3.
//
// Solidity: function x13() constant returns(address)
func (_DSGContract *DSGContractCaller) X13(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x13")
	return *ret0, err
}

// X13 is a free data retrieval call binding the contract method 0x8bd149b3.
//
// Solidity: function x13() constant returns(address)
func (_DSGContract *DSGContractSession) X13() (common.Address, error) {
	return _DSGContract.Contract.X13(&_DSGContract.CallOpts)
}

// X13 is a free data retrieval call binding the contract method 0x8bd149b3.
//
// Solidity: function x13() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X13() (common.Address, error) {
	return _DSGContract.Contract.X13(&_DSGContract.CallOpts)
}

// X14 is a free data retrieval call binding the contract method 0x05f04234.
//
// Solidity: function x14() constant returns(address)
func (_DSGContract *DSGContractCaller) X14(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x14")
	return *ret0, err
}

// X14 is a free data retrieval call binding the contract method 0x05f04234.
//
// Solidity: function x14() constant returns(address)
func (_DSGContract *DSGContractSession) X14() (common.Address, error) {
	return _DSGContract.Contract.X14(&_DSGContract.CallOpts)
}

// X14 is a free data retrieval call binding the contract method 0x05f04234.
//
// Solidity: function x14() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X14() (common.Address, error) {
	return _DSGContract.Contract.X14(&_DSGContract.CallOpts)
}

// X15 is a free data retrieval call binding the contract method 0xc0a1a949.
//
// Solidity: function x15() constant returns(address)
func (_DSGContract *DSGContractCaller) X15(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x15")
	return *ret0, err
}

// X15 is a free data retrieval call binding the contract method 0xc0a1a949.
//
// Solidity: function x15() constant returns(address)
func (_DSGContract *DSGContractSession) X15() (common.Address, error) {
	return _DSGContract.Contract.X15(&_DSGContract.CallOpts)
}

// X15 is a free data retrieval call binding the contract method 0xc0a1a949.
//
// Solidity: function x15() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X15() (common.Address, error) {
	return _DSGContract.Contract.X15(&_DSGContract.CallOpts)
}

// X16 is a free data retrieval call binding the contract method 0xa4f930d8.
//
// Solidity: function x16() constant returns(address)
func (_DSGContract *DSGContractCaller) X16(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x16")
	return *ret0, err
}

// X16 is a free data retrieval call binding the contract method 0xa4f930d8.
//
// Solidity: function x16() constant returns(address)
func (_DSGContract *DSGContractSession) X16() (common.Address, error) {
	return _DSGContract.Contract.X16(&_DSGContract.CallOpts)
}

// X16 is a free data retrieval call binding the contract method 0xa4f930d8.
//
// Solidity: function x16() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X16() (common.Address, error) {
	return _DSGContract.Contract.X16(&_DSGContract.CallOpts)
}

// X17 is a free data retrieval call binding the contract method 0x37545b00.
//
// Solidity: function x17() constant returns(address)
func (_DSGContract *DSGContractCaller) X17(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x17")
	return *ret0, err
}

// X17 is a free data retrieval call binding the contract method 0x37545b00.
//
// Solidity: function x17() constant returns(address)
func (_DSGContract *DSGContractSession) X17() (common.Address, error) {
	return _DSGContract.Contract.X17(&_DSGContract.CallOpts)
}

// X17 is a free data retrieval call binding the contract method 0x37545b00.
//
// Solidity: function x17() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X17() (common.Address, error) {
	return _DSGContract.Contract.X17(&_DSGContract.CallOpts)
}

// X18 is a free data retrieval call binding the contract method 0x14846dd3.
//
// Solidity: function x18() constant returns(address)
func (_DSGContract *DSGContractCaller) X18(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x18")
	return *ret0, err
}

// X18 is a free data retrieval call binding the contract method 0x14846dd3.
//
// Solidity: function x18() constant returns(address)
func (_DSGContract *DSGContractSession) X18() (common.Address, error) {
	return _DSGContract.Contract.X18(&_DSGContract.CallOpts)
}

// X18 is a free data retrieval call binding the contract method 0x14846dd3.
//
// Solidity: function x18() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X18() (common.Address, error) {
	return _DSGContract.Contract.X18(&_DSGContract.CallOpts)
}

// X19 is a free data retrieval call binding the contract method 0xeaecb012.
//
// Solidity: function x19() constant returns(address)
func (_DSGContract *DSGContractCaller) X19(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x19")
	return *ret0, err
}

// X19 is a free data retrieval call binding the contract method 0xeaecb012.
//
// Solidity: function x19() constant returns(address)
func (_DSGContract *DSGContractSession) X19() (common.Address, error) {
	return _DSGContract.Contract.X19(&_DSGContract.CallOpts)
}

// X19 is a free data retrieval call binding the contract method 0xeaecb012.
//
// Solidity: function x19() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X19() (common.Address, error) {
	return _DSGContract.Contract.X19(&_DSGContract.CallOpts)
}

// X2 is a free data retrieval call binding the contract method 0x4ff13571.
//
// Solidity: function x2() constant returns(address)
func (_DSGContract *DSGContractCaller) X2(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x2")
	return *ret0, err
}

// X2 is a free data retrieval call binding the contract method 0x4ff13571.
//
// Solidity: function x2() constant returns(address)
func (_DSGContract *DSGContractSession) X2() (common.Address, error) {
	return _DSGContract.Contract.X2(&_DSGContract.CallOpts)
}

// X2 is a free data retrieval call binding the contract method 0x4ff13571.
//
// Solidity: function x2() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X2() (common.Address, error) {
	return _DSGContract.Contract.X2(&_DSGContract.CallOpts)
}

// X20 is a free data retrieval call binding the contract method 0x9ffa6a9f.
//
// Solidity: function x20() constant returns(address)
func (_DSGContract *DSGContractCaller) X20(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x20")
	return *ret0, err
}

// X20 is a free data retrieval call binding the contract method 0x9ffa6a9f.
//
// Solidity: function x20() constant returns(address)
func (_DSGContract *DSGContractSession) X20() (common.Address, error) {
	return _DSGContract.Contract.X20(&_DSGContract.CallOpts)
}

// X20 is a free data retrieval call binding the contract method 0x9ffa6a9f.
//
// Solidity: function x20() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X20() (common.Address, error) {
	return _DSGContract.Contract.X20(&_DSGContract.CallOpts)
}

// X21 is a free data retrieval call binding the contract method 0x8cadd7f6.
//
// Solidity: function x21() constant returns(address)
func (_DSGContract *DSGContractCaller) X21(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x21")
	return *ret0, err
}

// X21 is a free data retrieval call binding the contract method 0x8cadd7f6.
//
// Solidity: function x21() constant returns(address)
func (_DSGContract *DSGContractSession) X21() (common.Address, error) {
	return _DSGContract.Contract.X21(&_DSGContract.CallOpts)
}

// X21 is a free data retrieval call binding the contract method 0x8cadd7f6.
//
// Solidity: function x21() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X21() (common.Address, error) {
	return _DSGContract.Contract.X21(&_DSGContract.CallOpts)
}

// X22 is a free data retrieval call binding the contract method 0x6bc3d2b8.
//
// Solidity: function x22() constant returns(address)
func (_DSGContract *DSGContractCaller) X22(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x22")
	return *ret0, err
}

// X22 is a free data retrieval call binding the contract method 0x6bc3d2b8.
//
// Solidity: function x22() constant returns(address)
func (_DSGContract *DSGContractSession) X22() (common.Address, error) {
	return _DSGContract.Contract.X22(&_DSGContract.CallOpts)
}

// X22 is a free data retrieval call binding the contract method 0x6bc3d2b8.
//
// Solidity: function x22() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X22() (common.Address, error) {
	return _DSGContract.Contract.X22(&_DSGContract.CallOpts)
}

// X23 is a free data retrieval call binding the contract method 0x3f3f3444.
//
// Solidity: function x23() constant returns(address)
func (_DSGContract *DSGContractCaller) X23(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x23")
	return *ret0, err
}

// X23 is a free data retrieval call binding the contract method 0x3f3f3444.
//
// Solidity: function x23() constant returns(address)
func (_DSGContract *DSGContractSession) X23() (common.Address, error) {
	return _DSGContract.Contract.X23(&_DSGContract.CallOpts)
}

// X23 is a free data retrieval call binding the contract method 0x3f3f3444.
//
// Solidity: function x23() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X23() (common.Address, error) {
	return _DSGContract.Contract.X23(&_DSGContract.CallOpts)
}

// X24 is a free data retrieval call binding the contract method 0x17bd117a.
//
// Solidity: function x24() constant returns(address)
func (_DSGContract *DSGContractCaller) X24(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x24")
	return *ret0, err
}

// X24 is a free data retrieval call binding the contract method 0x17bd117a.
//
// Solidity: function x24() constant returns(address)
func (_DSGContract *DSGContractSession) X24() (common.Address, error) {
	return _DSGContract.Contract.X24(&_DSGContract.CallOpts)
}

// X24 is a free data retrieval call binding the contract method 0x17bd117a.
//
// Solidity: function x24() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X24() (common.Address, error) {
	return _DSGContract.Contract.X24(&_DSGContract.CallOpts)
}

// X25 is a free data retrieval call binding the contract method 0x68937ad5.
//
// Solidity: function x25() constant returns(address)
func (_DSGContract *DSGContractCaller) X25(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x25")
	return *ret0, err
}

// X25 is a free data retrieval call binding the contract method 0x68937ad5.
//
// Solidity: function x25() constant returns(address)
func (_DSGContract *DSGContractSession) X25() (common.Address, error) {
	return _DSGContract.Contract.X25(&_DSGContract.CallOpts)
}

// X25 is a free data retrieval call binding the contract method 0x68937ad5.
//
// Solidity: function x25() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X25() (common.Address, error) {
	return _DSGContract.Contract.X25(&_DSGContract.CallOpts)
}

// X26 is a free data retrieval call binding the contract method 0xe749773f.
//
// Solidity: function x26() constant returns(address)
func (_DSGContract *DSGContractCaller) X26(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x26")
	return *ret0, err
}

// X26 is a free data retrieval call binding the contract method 0xe749773f.
//
// Solidity: function x26() constant returns(address)
func (_DSGContract *DSGContractSession) X26() (common.Address, error) {
	return _DSGContract.Contract.X26(&_DSGContract.CallOpts)
}

// X26 is a free data retrieval call binding the contract method 0xe749773f.
//
// Solidity: function x26() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X26() (common.Address, error) {
	return _DSGContract.Contract.X26(&_DSGContract.CallOpts)
}

// X27 is a free data retrieval call binding the contract method 0xbc8c3c9e.
//
// Solidity: function x27() constant returns(address)
func (_DSGContract *DSGContractCaller) X27(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x27")
	return *ret0, err
}

// X27 is a free data retrieval call binding the contract method 0xbc8c3c9e.
//
// Solidity: function x27() constant returns(address)
func (_DSGContract *DSGContractSession) X27() (common.Address, error) {
	return _DSGContract.Contract.X27(&_DSGContract.CallOpts)
}

// X27 is a free data retrieval call binding the contract method 0xbc8c3c9e.
//
// Solidity: function x27() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X27() (common.Address, error) {
	return _DSGContract.Contract.X27(&_DSGContract.CallOpts)
}

// X28 is a free data retrieval call binding the contract method 0xfe1878ef.
//
// Solidity: function x28() constant returns(address)
func (_DSGContract *DSGContractCaller) X28(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x28")
	return *ret0, err
}

// X28 is a free data retrieval call binding the contract method 0xfe1878ef.
//
// Solidity: function x28() constant returns(address)
func (_DSGContract *DSGContractSession) X28() (common.Address, error) {
	return _DSGContract.Contract.X28(&_DSGContract.CallOpts)
}

// X28 is a free data retrieval call binding the contract method 0xfe1878ef.
//
// Solidity: function x28() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X28() (common.Address, error) {
	return _DSGContract.Contract.X28(&_DSGContract.CallOpts)
}

// X29 is a free data retrieval call binding the contract method 0xb54d94f8.
//
// Solidity: function x29() constant returns(address)
func (_DSGContract *DSGContractCaller) X29(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x29")
	return *ret0, err
}

// X29 is a free data retrieval call binding the contract method 0xb54d94f8.
//
// Solidity: function x29() constant returns(address)
func (_DSGContract *DSGContractSession) X29() (common.Address, error) {
	return _DSGContract.Contract.X29(&_DSGContract.CallOpts)
}

// X29 is a free data retrieval call binding the contract method 0xb54d94f8.
//
// Solidity: function x29() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X29() (common.Address, error) {
	return _DSGContract.Contract.X29(&_DSGContract.CallOpts)
}

// X3 is a free data retrieval call binding the contract method 0xa3b78873.
//
// Solidity: function x3() constant returns(address)
func (_DSGContract *DSGContractCaller) X3(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x3")
	return *ret0, err
}

// X3 is a free data retrieval call binding the contract method 0xa3b78873.
//
// Solidity: function x3() constant returns(address)
func (_DSGContract *DSGContractSession) X3() (common.Address, error) {
	return _DSGContract.Contract.X3(&_DSGContract.CallOpts)
}

// X3 is a free data retrieval call binding the contract method 0xa3b78873.
//
// Solidity: function x3() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X3() (common.Address, error) {
	return _DSGContract.Contract.X3(&_DSGContract.CallOpts)
}

// X30 is a free data retrieval call binding the contract method 0x39cb8435.
//
// Solidity: function x30() constant returns(address)
func (_DSGContract *DSGContractCaller) X30(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x30")
	return *ret0, err
}

// X30 is a free data retrieval call binding the contract method 0x39cb8435.
//
// Solidity: function x30() constant returns(address)
func (_DSGContract *DSGContractSession) X30() (common.Address, error) {
	return _DSGContract.Contract.X30(&_DSGContract.CallOpts)
}

// X30 is a free data retrieval call binding the contract method 0x39cb8435.
//
// Solidity: function x30() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X30() (common.Address, error) {
	return _DSGContract.Contract.X30(&_DSGContract.CallOpts)
}

// X31 is a free data retrieval call binding the contract method 0x72c019b3.
//
// Solidity: function x31() constant returns(address)
func (_DSGContract *DSGContractCaller) X31(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x31")
	return *ret0, err
}

// X31 is a free data retrieval call binding the contract method 0x72c019b3.
//
// Solidity: function x31() constant returns(address)
func (_DSGContract *DSGContractSession) X31() (common.Address, error) {
	return _DSGContract.Contract.X31(&_DSGContract.CallOpts)
}

// X31 is a free data retrieval call binding the contract method 0x72c019b3.
//
// Solidity: function x31() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X31() (common.Address, error) {
	return _DSGContract.Contract.X31(&_DSGContract.CallOpts)
}

// X32 is a free data retrieval call binding the contract method 0x908df6b1.
//
// Solidity: function x32() constant returns(address)
func (_DSGContract *DSGContractCaller) X32(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x32")
	return *ret0, err
}

// X32 is a free data retrieval call binding the contract method 0x908df6b1.
//
// Solidity: function x32() constant returns(address)
func (_DSGContract *DSGContractSession) X32() (common.Address, error) {
	return _DSGContract.Contract.X32(&_DSGContract.CallOpts)
}

// X32 is a free data retrieval call binding the contract method 0x908df6b1.
//
// Solidity: function x32() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X32() (common.Address, error) {
	return _DSGContract.Contract.X32(&_DSGContract.CallOpts)
}

// X33 is a free data retrieval call binding the contract method 0xc5f720e4.
//
// Solidity: function x33() constant returns(address)
func (_DSGContract *DSGContractCaller) X33(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x33")
	return *ret0, err
}

// X33 is a free data retrieval call binding the contract method 0xc5f720e4.
//
// Solidity: function x33() constant returns(address)
func (_DSGContract *DSGContractSession) X33() (common.Address, error) {
	return _DSGContract.Contract.X33(&_DSGContract.CallOpts)
}

// X33 is a free data retrieval call binding the contract method 0xc5f720e4.
//
// Solidity: function x33() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X33() (common.Address, error) {
	return _DSGContract.Contract.X33(&_DSGContract.CallOpts)
}

// X34 is a free data retrieval call binding the contract method 0xfd56a650.
//
// Solidity: function x34() constant returns(address)
func (_DSGContract *DSGContractCaller) X34(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x34")
	return *ret0, err
}

// X34 is a free data retrieval call binding the contract method 0xfd56a650.
//
// Solidity: function x34() constant returns(address)
func (_DSGContract *DSGContractSession) X34() (common.Address, error) {
	return _DSGContract.Contract.X34(&_DSGContract.CallOpts)
}

// X34 is a free data retrieval call binding the contract method 0xfd56a650.
//
// Solidity: function x34() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X34() (common.Address, error) {
	return _DSGContract.Contract.X34(&_DSGContract.CallOpts)
}

// X35 is a free data retrieval call binding the contract method 0x5e217f42.
//
// Solidity: function x35() constant returns(address)
func (_DSGContract *DSGContractCaller) X35(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x35")
	return *ret0, err
}

// X35 is a free data retrieval call binding the contract method 0x5e217f42.
//
// Solidity: function x35() constant returns(address)
func (_DSGContract *DSGContractSession) X35() (common.Address, error) {
	return _DSGContract.Contract.X35(&_DSGContract.CallOpts)
}

// X35 is a free data retrieval call binding the contract method 0x5e217f42.
//
// Solidity: function x35() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X35() (common.Address, error) {
	return _DSGContract.Contract.X35(&_DSGContract.CallOpts)
}

// X36 is a free data retrieval call binding the contract method 0xb407a0db.
//
// Solidity: function x36() constant returns(address)
func (_DSGContract *DSGContractCaller) X36(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x36")
	return *ret0, err
}

// X36 is a free data retrieval call binding the contract method 0xb407a0db.
//
// Solidity: function x36() constant returns(address)
func (_DSGContract *DSGContractSession) X36() (common.Address, error) {
	return _DSGContract.Contract.X36(&_DSGContract.CallOpts)
}

// X36 is a free data retrieval call binding the contract method 0xb407a0db.
//
// Solidity: function x36() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X36() (common.Address, error) {
	return _DSGContract.Contract.X36(&_DSGContract.CallOpts)
}

// X37 is a free data retrieval call binding the contract method 0xb9fbc543.
//
// Solidity: function x37() constant returns(address)
func (_DSGContract *DSGContractCaller) X37(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x37")
	return *ret0, err
}

// X37 is a free data retrieval call binding the contract method 0xb9fbc543.
//
// Solidity: function x37() constant returns(address)
func (_DSGContract *DSGContractSession) X37() (common.Address, error) {
	return _DSGContract.Contract.X37(&_DSGContract.CallOpts)
}

// X37 is a free data retrieval call binding the contract method 0xb9fbc543.
//
// Solidity: function x37() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X37() (common.Address, error) {
	return _DSGContract.Contract.X37(&_DSGContract.CallOpts)
}

// X38 is a free data retrieval call binding the contract method 0xa19f482a.
//
// Solidity: function x38() constant returns(address)
func (_DSGContract *DSGContractCaller) X38(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x38")
	return *ret0, err
}

// X38 is a free data retrieval call binding the contract method 0xa19f482a.
//
// Solidity: function x38() constant returns(address)
func (_DSGContract *DSGContractSession) X38() (common.Address, error) {
	return _DSGContract.Contract.X38(&_DSGContract.CallOpts)
}

// X38 is a free data retrieval call binding the contract method 0xa19f482a.
//
// Solidity: function x38() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X38() (common.Address, error) {
	return _DSGContract.Contract.X38(&_DSGContract.CallOpts)
}

// X39 is a free data retrieval call binding the contract method 0x3176cca2.
//
// Solidity: function x39() constant returns(address)
func (_DSGContract *DSGContractCaller) X39(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x39")
	return *ret0, err
}

// X39 is a free data retrieval call binding the contract method 0x3176cca2.
//
// Solidity: function x39() constant returns(address)
func (_DSGContract *DSGContractSession) X39() (common.Address, error) {
	return _DSGContract.Contract.X39(&_DSGContract.CallOpts)
}

// X39 is a free data retrieval call binding the contract method 0x3176cca2.
//
// Solidity: function x39() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X39() (common.Address, error) {
	return _DSGContract.Contract.X39(&_DSGContract.CallOpts)
}

// X4 is a free data retrieval call binding the contract method 0xcb6cee8a.
//
// Solidity: function x4() constant returns(address)
func (_DSGContract *DSGContractCaller) X4(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x4")
	return *ret0, err
}

// X4 is a free data retrieval call binding the contract method 0xcb6cee8a.
//
// Solidity: function x4() constant returns(address)
func (_DSGContract *DSGContractSession) X4() (common.Address, error) {
	return _DSGContract.Contract.X4(&_DSGContract.CallOpts)
}

// X4 is a free data retrieval call binding the contract method 0xcb6cee8a.
//
// Solidity: function x4() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X4() (common.Address, error) {
	return _DSGContract.Contract.X4(&_DSGContract.CallOpts)
}

// X40 is a free data retrieval call binding the contract method 0x2ca07afd.
//
// Solidity: function x40() constant returns(address)
func (_DSGContract *DSGContractCaller) X40(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x40")
	return *ret0, err
}

// X40 is a free data retrieval call binding the contract method 0x2ca07afd.
//
// Solidity: function x40() constant returns(address)
func (_DSGContract *DSGContractSession) X40() (common.Address, error) {
	return _DSGContract.Contract.X40(&_DSGContract.CallOpts)
}

// X40 is a free data retrieval call binding the contract method 0x2ca07afd.
//
// Solidity: function x40() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X40() (common.Address, error) {
	return _DSGContract.Contract.X40(&_DSGContract.CallOpts)
}

// X41 is a free data retrieval call binding the contract method 0xf24ec5c8.
//
// Solidity: function x41() constant returns(address)
func (_DSGContract *DSGContractCaller) X41(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x41")
	return *ret0, err
}

// X41 is a free data retrieval call binding the contract method 0xf24ec5c8.
//
// Solidity: function x41() constant returns(address)
func (_DSGContract *DSGContractSession) X41() (common.Address, error) {
	return _DSGContract.Contract.X41(&_DSGContract.CallOpts)
}

// X41 is a free data retrieval call binding the contract method 0xf24ec5c8.
//
// Solidity: function x41() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X41() (common.Address, error) {
	return _DSGContract.Contract.X41(&_DSGContract.CallOpts)
}

// X42 is a free data retrieval call binding the contract method 0x943a0169.
//
// Solidity: function x42() constant returns(address)
func (_DSGContract *DSGContractCaller) X42(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x42")
	return *ret0, err
}

// X42 is a free data retrieval call binding the contract method 0x943a0169.
//
// Solidity: function x42() constant returns(address)
func (_DSGContract *DSGContractSession) X42() (common.Address, error) {
	return _DSGContract.Contract.X42(&_DSGContract.CallOpts)
}

// X42 is a free data retrieval call binding the contract method 0x943a0169.
//
// Solidity: function x42() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X42() (common.Address, error) {
	return _DSGContract.Contract.X42(&_DSGContract.CallOpts)
}

// X43 is a free data retrieval call binding the contract method 0x402cc27e.
//
// Solidity: function x43() constant returns(address)
func (_DSGContract *DSGContractCaller) X43(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x43")
	return *ret0, err
}

// X43 is a free data retrieval call binding the contract method 0x402cc27e.
//
// Solidity: function x43() constant returns(address)
func (_DSGContract *DSGContractSession) X43() (common.Address, error) {
	return _DSGContract.Contract.X43(&_DSGContract.CallOpts)
}

// X43 is a free data retrieval call binding the contract method 0x402cc27e.
//
// Solidity: function x43() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X43() (common.Address, error) {
	return _DSGContract.Contract.X43(&_DSGContract.CallOpts)
}

// X44 is a free data retrieval call binding the contract method 0x0ae490ad.
//
// Solidity: function x44() constant returns(address)
func (_DSGContract *DSGContractCaller) X44(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x44")
	return *ret0, err
}

// X44 is a free data retrieval call binding the contract method 0x0ae490ad.
//
// Solidity: function x44() constant returns(address)
func (_DSGContract *DSGContractSession) X44() (common.Address, error) {
	return _DSGContract.Contract.X44(&_DSGContract.CallOpts)
}

// X44 is a free data retrieval call binding the contract method 0x0ae490ad.
//
// Solidity: function x44() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X44() (common.Address, error) {
	return _DSGContract.Contract.X44(&_DSGContract.CallOpts)
}

// X45 is a free data retrieval call binding the contract method 0x1c073086.
//
// Solidity: function x45() constant returns(address)
func (_DSGContract *DSGContractCaller) X45(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x45")
	return *ret0, err
}

// X45 is a free data retrieval call binding the contract method 0x1c073086.
//
// Solidity: function x45() constant returns(address)
func (_DSGContract *DSGContractSession) X45() (common.Address, error) {
	return _DSGContract.Contract.X45(&_DSGContract.CallOpts)
}

// X45 is a free data retrieval call binding the contract method 0x1c073086.
//
// Solidity: function x45() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X45() (common.Address, error) {
	return _DSGContract.Contract.X45(&_DSGContract.CallOpts)
}

// X46 is a free data retrieval call binding the contract method 0x449136cf.
//
// Solidity: function x46() constant returns(address)
func (_DSGContract *DSGContractCaller) X46(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x46")
	return *ret0, err
}

// X46 is a free data retrieval call binding the contract method 0x449136cf.
//
// Solidity: function x46() constant returns(address)
func (_DSGContract *DSGContractSession) X46() (common.Address, error) {
	return _DSGContract.Contract.X46(&_DSGContract.CallOpts)
}

// X46 is a free data retrieval call binding the contract method 0x449136cf.
//
// Solidity: function x46() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X46() (common.Address, error) {
	return _DSGContract.Contract.X46(&_DSGContract.CallOpts)
}

// X47 is a free data retrieval call binding the contract method 0x72e1b839.
//
// Solidity: function x47() constant returns(address)
func (_DSGContract *DSGContractCaller) X47(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x47")
	return *ret0, err
}

// X47 is a free data retrieval call binding the contract method 0x72e1b839.
//
// Solidity: function x47() constant returns(address)
func (_DSGContract *DSGContractSession) X47() (common.Address, error) {
	return _DSGContract.Contract.X47(&_DSGContract.CallOpts)
}

// X47 is a free data retrieval call binding the contract method 0x72e1b839.
//
// Solidity: function x47() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X47() (common.Address, error) {
	return _DSGContract.Contract.X47(&_DSGContract.CallOpts)
}

// X48 is a free data retrieval call binding the contract method 0x7fda3be2.
//
// Solidity: function x48() constant returns(address)
func (_DSGContract *DSGContractCaller) X48(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x48")
	return *ret0, err
}

// X48 is a free data retrieval call binding the contract method 0x7fda3be2.
//
// Solidity: function x48() constant returns(address)
func (_DSGContract *DSGContractSession) X48() (common.Address, error) {
	return _DSGContract.Contract.X48(&_DSGContract.CallOpts)
}

// X48 is a free data retrieval call binding the contract method 0x7fda3be2.
//
// Solidity: function x48() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X48() (common.Address, error) {
	return _DSGContract.Contract.X48(&_DSGContract.CallOpts)
}

// X49 is a free data retrieval call binding the contract method 0x0418043d.
//
// Solidity: function x49() constant returns(address)
func (_DSGContract *DSGContractCaller) X49(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x49")
	return *ret0, err
}

// X49 is a free data retrieval call binding the contract method 0x0418043d.
//
// Solidity: function x49() constant returns(address)
func (_DSGContract *DSGContractSession) X49() (common.Address, error) {
	return _DSGContract.Contract.X49(&_DSGContract.CallOpts)
}

// X49 is a free data retrieval call binding the contract method 0x0418043d.
//
// Solidity: function x49() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X49() (common.Address, error) {
	return _DSGContract.Contract.X49(&_DSGContract.CallOpts)
}

// X5 is a free data retrieval call binding the contract method 0xcd145f65.
//
// Solidity: function x5() constant returns(address)
func (_DSGContract *DSGContractCaller) X5(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x5")
	return *ret0, err
}

// X5 is a free data retrieval call binding the contract method 0xcd145f65.
//
// Solidity: function x5() constant returns(address)
func (_DSGContract *DSGContractSession) X5() (common.Address, error) {
	return _DSGContract.Contract.X5(&_DSGContract.CallOpts)
}

// X5 is a free data retrieval call binding the contract method 0xcd145f65.
//
// Solidity: function x5() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X5() (common.Address, error) {
	return _DSGContract.Contract.X5(&_DSGContract.CallOpts)
}

// X50 is a free data retrieval call binding the contract method 0xb4feb110.
//
// Solidity: function x50() constant returns(address)
func (_DSGContract *DSGContractCaller) X50(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x50")
	return *ret0, err
}

// X50 is a free data retrieval call binding the contract method 0xb4feb110.
//
// Solidity: function x50() constant returns(address)
func (_DSGContract *DSGContractSession) X50() (common.Address, error) {
	return _DSGContract.Contract.X50(&_DSGContract.CallOpts)
}

// X50 is a free data retrieval call binding the contract method 0xb4feb110.
//
// Solidity: function x50() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X50() (common.Address, error) {
	return _DSGContract.Contract.X50(&_DSGContract.CallOpts)
}

// X51 is a free data retrieval call binding the contract method 0x88ce3426.
//
// Solidity: function x51() constant returns(address)
func (_DSGContract *DSGContractCaller) X51(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x51")
	return *ret0, err
}

// X51 is a free data retrieval call binding the contract method 0x88ce3426.
//
// Solidity: function x51() constant returns(address)
func (_DSGContract *DSGContractSession) X51() (common.Address, error) {
	return _DSGContract.Contract.X51(&_DSGContract.CallOpts)
}

// X51 is a free data retrieval call binding the contract method 0x88ce3426.
//
// Solidity: function x51() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X51() (common.Address, error) {
	return _DSGContract.Contract.X51(&_DSGContract.CallOpts)
}

// X52 is a free data retrieval call binding the contract method 0xf304e720.
//
// Solidity: function x52() constant returns(address)
func (_DSGContract *DSGContractCaller) X52(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x52")
	return *ret0, err
}

// X52 is a free data retrieval call binding the contract method 0xf304e720.
//
// Solidity: function x52() constant returns(address)
func (_DSGContract *DSGContractSession) X52() (common.Address, error) {
	return _DSGContract.Contract.X52(&_DSGContract.CallOpts)
}

// X52 is a free data retrieval call binding the contract method 0xf304e720.
//
// Solidity: function x52() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X52() (common.Address, error) {
	return _DSGContract.Contract.X52(&_DSGContract.CallOpts)
}

// X53 is a free data retrieval call binding the contract method 0xbcfbbc46.
//
// Solidity: function x53() constant returns(address)
func (_DSGContract *DSGContractCaller) X53(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x53")
	return *ret0, err
}

// X53 is a free data retrieval call binding the contract method 0xbcfbbc46.
//
// Solidity: function x53() constant returns(address)
func (_DSGContract *DSGContractSession) X53() (common.Address, error) {
	return _DSGContract.Contract.X53(&_DSGContract.CallOpts)
}

// X53 is a free data retrieval call binding the contract method 0xbcfbbc46.
//
// Solidity: function x53() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X53() (common.Address, error) {
	return _DSGContract.Contract.X53(&_DSGContract.CallOpts)
}

// X54 is a free data retrieval call binding the contract method 0xaf594904.
//
// Solidity: function x54() constant returns(address)
func (_DSGContract *DSGContractCaller) X54(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x54")
	return *ret0, err
}

// X54 is a free data retrieval call binding the contract method 0xaf594904.
//
// Solidity: function x54() constant returns(address)
func (_DSGContract *DSGContractSession) X54() (common.Address, error) {
	return _DSGContract.Contract.X54(&_DSGContract.CallOpts)
}

// X54 is a free data retrieval call binding the contract method 0xaf594904.
//
// Solidity: function x54() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X54() (common.Address, error) {
	return _DSGContract.Contract.X54(&_DSGContract.CallOpts)
}

// X55 is a free data retrieval call binding the contract method 0x7af6e874.
//
// Solidity: function x55() constant returns(address)
func (_DSGContract *DSGContractCaller) X55(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x55")
	return *ret0, err
}

// X55 is a free data retrieval call binding the contract method 0x7af6e874.
//
// Solidity: function x55() constant returns(address)
func (_DSGContract *DSGContractSession) X55() (common.Address, error) {
	return _DSGContract.Contract.X55(&_DSGContract.CallOpts)
}

// X55 is a free data retrieval call binding the contract method 0x7af6e874.
//
// Solidity: function x55() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X55() (common.Address, error) {
	return _DSGContract.Contract.X55(&_DSGContract.CallOpts)
}

// X56 is a free data retrieval call binding the contract method 0xfddbd387.
//
// Solidity: function x56() constant returns(address)
func (_DSGContract *DSGContractCaller) X56(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x56")
	return *ret0, err
}

// X56 is a free data retrieval call binding the contract method 0xfddbd387.
//
// Solidity: function x56() constant returns(address)
func (_DSGContract *DSGContractSession) X56() (common.Address, error) {
	return _DSGContract.Contract.X56(&_DSGContract.CallOpts)
}

// X56 is a free data retrieval call binding the contract method 0xfddbd387.
//
// Solidity: function x56() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X56() (common.Address, error) {
	return _DSGContract.Contract.X56(&_DSGContract.CallOpts)
}

// X57 is a free data retrieval call binding the contract method 0xeb5d92a0.
//
// Solidity: function x57() constant returns(address)
func (_DSGContract *DSGContractCaller) X57(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x57")
	return *ret0, err
}

// X57 is a free data retrieval call binding the contract method 0xeb5d92a0.
//
// Solidity: function x57() constant returns(address)
func (_DSGContract *DSGContractSession) X57() (common.Address, error) {
	return _DSGContract.Contract.X57(&_DSGContract.CallOpts)
}

// X57 is a free data retrieval call binding the contract method 0xeb5d92a0.
//
// Solidity: function x57() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X57() (common.Address, error) {
	return _DSGContract.Contract.X57(&_DSGContract.CallOpts)
}

// X58 is a free data retrieval call binding the contract method 0xa87f6ab5.
//
// Solidity: function x58() constant returns(address)
func (_DSGContract *DSGContractCaller) X58(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x58")
	return *ret0, err
}

// X58 is a free data retrieval call binding the contract method 0xa87f6ab5.
//
// Solidity: function x58() constant returns(address)
func (_DSGContract *DSGContractSession) X58() (common.Address, error) {
	return _DSGContract.Contract.X58(&_DSGContract.CallOpts)
}

// X58 is a free data retrieval call binding the contract method 0xa87f6ab5.
//
// Solidity: function x58() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X58() (common.Address, error) {
	return _DSGContract.Contract.X58(&_DSGContract.CallOpts)
}

// X59 is a free data retrieval call binding the contract method 0x13813beb.
//
// Solidity: function x59() constant returns(address)
func (_DSGContract *DSGContractCaller) X59(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x59")
	return *ret0, err
}

// X59 is a free data retrieval call binding the contract method 0x13813beb.
//
// Solidity: function x59() constant returns(address)
func (_DSGContract *DSGContractSession) X59() (common.Address, error) {
	return _DSGContract.Contract.X59(&_DSGContract.CallOpts)
}

// X59 is a free data retrieval call binding the contract method 0x13813beb.
//
// Solidity: function x59() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X59() (common.Address, error) {
	return _DSGContract.Contract.X59(&_DSGContract.CallOpts)
}

// X6 is a free data retrieval call binding the contract method 0xeaeb89f7.
//
// Solidity: function x6() constant returns(address)
func (_DSGContract *DSGContractCaller) X6(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x6")
	return *ret0, err
}

// X6 is a free data retrieval call binding the contract method 0xeaeb89f7.
//
// Solidity: function x6() constant returns(address)
func (_DSGContract *DSGContractSession) X6() (common.Address, error) {
	return _DSGContract.Contract.X6(&_DSGContract.CallOpts)
}

// X6 is a free data retrieval call binding the contract method 0xeaeb89f7.
//
// Solidity: function x6() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X6() (common.Address, error) {
	return _DSGContract.Contract.X6(&_DSGContract.CallOpts)
}

// X60 is a free data retrieval call binding the contract method 0x89b3bc3a.
//
// Solidity: function x60() constant returns(address)
func (_DSGContract *DSGContractCaller) X60(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x60")
	return *ret0, err
}

// X60 is a free data retrieval call binding the contract method 0x89b3bc3a.
//
// Solidity: function x60() constant returns(address)
func (_DSGContract *DSGContractSession) X60() (common.Address, error) {
	return _DSGContract.Contract.X60(&_DSGContract.CallOpts)
}

// X60 is a free data retrieval call binding the contract method 0x89b3bc3a.
//
// Solidity: function x60() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X60() (common.Address, error) {
	return _DSGContract.Contract.X60(&_DSGContract.CallOpts)
}

// X61 is a free data retrieval call binding the contract method 0x588d1150.
//
// Solidity: function x61() constant returns(address)
func (_DSGContract *DSGContractCaller) X61(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x61")
	return *ret0, err
}

// X61 is a free data retrieval call binding the contract method 0x588d1150.
//
// Solidity: function x61() constant returns(address)
func (_DSGContract *DSGContractSession) X61() (common.Address, error) {
	return _DSGContract.Contract.X61(&_DSGContract.CallOpts)
}

// X61 is a free data retrieval call binding the contract method 0x588d1150.
//
// Solidity: function x61() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X61() (common.Address, error) {
	return _DSGContract.Contract.X61(&_DSGContract.CallOpts)
}

// X62 is a free data retrieval call binding the contract method 0x332bb13b.
//
// Solidity: function x62() constant returns(address)
func (_DSGContract *DSGContractCaller) X62(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x62")
	return *ret0, err
}

// X62 is a free data retrieval call binding the contract method 0x332bb13b.
//
// Solidity: function x62() constant returns(address)
func (_DSGContract *DSGContractSession) X62() (common.Address, error) {
	return _DSGContract.Contract.X62(&_DSGContract.CallOpts)
}

// X62 is a free data retrieval call binding the contract method 0x332bb13b.
//
// Solidity: function x62() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X62() (common.Address, error) {
	return _DSGContract.Contract.X62(&_DSGContract.CallOpts)
}

// X63 is a free data retrieval call binding the contract method 0x1fc58dae.
//
// Solidity: function x63() constant returns(address)
func (_DSGContract *DSGContractCaller) X63(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x63")
	return *ret0, err
}

// X63 is a free data retrieval call binding the contract method 0x1fc58dae.
//
// Solidity: function x63() constant returns(address)
func (_DSGContract *DSGContractSession) X63() (common.Address, error) {
	return _DSGContract.Contract.X63(&_DSGContract.CallOpts)
}

// X63 is a free data retrieval call binding the contract method 0x1fc58dae.
//
// Solidity: function x63() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X63() (common.Address, error) {
	return _DSGContract.Contract.X63(&_DSGContract.CallOpts)
}

// X64 is a free data retrieval call binding the contract method 0x9bf3f61a.
//
// Solidity: function x64() constant returns(address)
func (_DSGContract *DSGContractCaller) X64(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x64")
	return *ret0, err
}

// X64 is a free data retrieval call binding the contract method 0x9bf3f61a.
//
// Solidity: function x64() constant returns(address)
func (_DSGContract *DSGContractSession) X64() (common.Address, error) {
	return _DSGContract.Contract.X64(&_DSGContract.CallOpts)
}

// X64 is a free data retrieval call binding the contract method 0x9bf3f61a.
//
// Solidity: function x64() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X64() (common.Address, error) {
	return _DSGContract.Contract.X64(&_DSGContract.CallOpts)
}

// X65 is a free data retrieval call binding the contract method 0x745be928.
//
// Solidity: function x65() constant returns(address)
func (_DSGContract *DSGContractCaller) X65(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x65")
	return *ret0, err
}

// X65 is a free data retrieval call binding the contract method 0x745be928.
//
// Solidity: function x65() constant returns(address)
func (_DSGContract *DSGContractSession) X65() (common.Address, error) {
	return _DSGContract.Contract.X65(&_DSGContract.CallOpts)
}

// X65 is a free data retrieval call binding the contract method 0x745be928.
//
// Solidity: function x65() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X65() (common.Address, error) {
	return _DSGContract.Contract.X65(&_DSGContract.CallOpts)
}

// X66 is a free data retrieval call binding the contract method 0xa11be9af.
//
// Solidity: function x66() constant returns(address)
func (_DSGContract *DSGContractCaller) X66(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x66")
	return *ret0, err
}

// X66 is a free data retrieval call binding the contract method 0xa11be9af.
//
// Solidity: function x66() constant returns(address)
func (_DSGContract *DSGContractSession) X66() (common.Address, error) {
	return _DSGContract.Contract.X66(&_DSGContract.CallOpts)
}

// X66 is a free data retrieval call binding the contract method 0xa11be9af.
//
// Solidity: function x66() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X66() (common.Address, error) {
	return _DSGContract.Contract.X66(&_DSGContract.CallOpts)
}

// X67 is a free data retrieval call binding the contract method 0xbca63eb4.
//
// Solidity: function x67() constant returns(address)
func (_DSGContract *DSGContractCaller) X67(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x67")
	return *ret0, err
}

// X67 is a free data retrieval call binding the contract method 0xbca63eb4.
//
// Solidity: function x67() constant returns(address)
func (_DSGContract *DSGContractSession) X67() (common.Address, error) {
	return _DSGContract.Contract.X67(&_DSGContract.CallOpts)
}

// X67 is a free data retrieval call binding the contract method 0xbca63eb4.
//
// Solidity: function x67() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X67() (common.Address, error) {
	return _DSGContract.Contract.X67(&_DSGContract.CallOpts)
}

// X68 is a free data retrieval call binding the contract method 0xf721734f.
//
// Solidity: function x68() constant returns(address)
func (_DSGContract *DSGContractCaller) X68(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x68")
	return *ret0, err
}

// X68 is a free data retrieval call binding the contract method 0xf721734f.
//
// Solidity: function x68() constant returns(address)
func (_DSGContract *DSGContractSession) X68() (common.Address, error) {
	return _DSGContract.Contract.X68(&_DSGContract.CallOpts)
}

// X68 is a free data retrieval call binding the contract method 0xf721734f.
//
// Solidity: function x68() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X68() (common.Address, error) {
	return _DSGContract.Contract.X68(&_DSGContract.CallOpts)
}

// X69 is a free data retrieval call binding the contract method 0xeac637f2.
//
// Solidity: function x69() constant returns(address)
func (_DSGContract *DSGContractCaller) X69(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x69")
	return *ret0, err
}

// X69 is a free data retrieval call binding the contract method 0xeac637f2.
//
// Solidity: function x69() constant returns(address)
func (_DSGContract *DSGContractSession) X69() (common.Address, error) {
	return _DSGContract.Contract.X69(&_DSGContract.CallOpts)
}

// X69 is a free data retrieval call binding the contract method 0xeac637f2.
//
// Solidity: function x69() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X69() (common.Address, error) {
	return _DSGContract.Contract.X69(&_DSGContract.CallOpts)
}

// X7 is a free data retrieval call binding the contract method 0x282940a7.
//
// Solidity: function x7() constant returns(address)
func (_DSGContract *DSGContractCaller) X7(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x7")
	return *ret0, err
}

// X7 is a free data retrieval call binding the contract method 0x282940a7.
//
// Solidity: function x7() constant returns(address)
func (_DSGContract *DSGContractSession) X7() (common.Address, error) {
	return _DSGContract.Contract.X7(&_DSGContract.CallOpts)
}

// X7 is a free data retrieval call binding the contract method 0x282940a7.
//
// Solidity: function x7() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X7() (common.Address, error) {
	return _DSGContract.Contract.X7(&_DSGContract.CallOpts)
}

// X70 is a free data retrieval call binding the contract method 0x5d370399.
//
// Solidity: function x70() constant returns(address)
func (_DSGContract *DSGContractCaller) X70(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x70")
	return *ret0, err
}

// X70 is a free data retrieval call binding the contract method 0x5d370399.
//
// Solidity: function x70() constant returns(address)
func (_DSGContract *DSGContractSession) X70() (common.Address, error) {
	return _DSGContract.Contract.X70(&_DSGContract.CallOpts)
}

// X70 is a free data retrieval call binding the contract method 0x5d370399.
//
// Solidity: function x70() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X70() (common.Address, error) {
	return _DSGContract.Contract.X70(&_DSGContract.CallOpts)
}

// X71 is a free data retrieval call binding the contract method 0x7ede4904.
//
// Solidity: function x71() constant returns(address)
func (_DSGContract *DSGContractCaller) X71(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x71")
	return *ret0, err
}

// X71 is a free data retrieval call binding the contract method 0x7ede4904.
//
// Solidity: function x71() constant returns(address)
func (_DSGContract *DSGContractSession) X71() (common.Address, error) {
	return _DSGContract.Contract.X71(&_DSGContract.CallOpts)
}

// X71 is a free data retrieval call binding the contract method 0x7ede4904.
//
// Solidity: function x71() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X71() (common.Address, error) {
	return _DSGContract.Contract.X71(&_DSGContract.CallOpts)
}

// X72 is a free data retrieval call binding the contract method 0xac1e43a0.
//
// Solidity: function x72() constant returns(address)
func (_DSGContract *DSGContractCaller) X72(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x72")
	return *ret0, err
}

// X72 is a free data retrieval call binding the contract method 0xac1e43a0.
//
// Solidity: function x72() constant returns(address)
func (_DSGContract *DSGContractSession) X72() (common.Address, error) {
	return _DSGContract.Contract.X72(&_DSGContract.CallOpts)
}

// X72 is a free data retrieval call binding the contract method 0xac1e43a0.
//
// Solidity: function x72() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X72() (common.Address, error) {
	return _DSGContract.Contract.X72(&_DSGContract.CallOpts)
}

// X73 is a free data retrieval call binding the contract method 0x9b85e451.
//
// Solidity: function x73() constant returns(address)
func (_DSGContract *DSGContractCaller) X73(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x73")
	return *ret0, err
}

// X73 is a free data retrieval call binding the contract method 0x9b85e451.
//
// Solidity: function x73() constant returns(address)
func (_DSGContract *DSGContractSession) X73() (common.Address, error) {
	return _DSGContract.Contract.X73(&_DSGContract.CallOpts)
}

// X73 is a free data retrieval call binding the contract method 0x9b85e451.
//
// Solidity: function x73() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X73() (common.Address, error) {
	return _DSGContract.Contract.X73(&_DSGContract.CallOpts)
}

// X74 is a free data retrieval call binding the contract method 0xba16cefe.
//
// Solidity: function x74() constant returns(address)
func (_DSGContract *DSGContractCaller) X74(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x74")
	return *ret0, err
}

// X74 is a free data retrieval call binding the contract method 0xba16cefe.
//
// Solidity: function x74() constant returns(address)
func (_DSGContract *DSGContractSession) X74() (common.Address, error) {
	return _DSGContract.Contract.X74(&_DSGContract.CallOpts)
}

// X74 is a free data retrieval call binding the contract method 0xba16cefe.
//
// Solidity: function x74() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X74() (common.Address, error) {
	return _DSGContract.Contract.X74(&_DSGContract.CallOpts)
}

// X75 is a free data retrieval call binding the contract method 0x175d756b.
//
// Solidity: function x75() constant returns(address)
func (_DSGContract *DSGContractCaller) X75(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x75")
	return *ret0, err
}

// X75 is a free data retrieval call binding the contract method 0x175d756b.
//
// Solidity: function x75() constant returns(address)
func (_DSGContract *DSGContractSession) X75() (common.Address, error) {
	return _DSGContract.Contract.X75(&_DSGContract.CallOpts)
}

// X75 is a free data retrieval call binding the contract method 0x175d756b.
//
// Solidity: function x75() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X75() (common.Address, error) {
	return _DSGContract.Contract.X75(&_DSGContract.CallOpts)
}

// X76 is a free data retrieval call binding the contract method 0x7f758a15.
//
// Solidity: function x76() constant returns(address)
func (_DSGContract *DSGContractCaller) X76(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x76")
	return *ret0, err
}

// X76 is a free data retrieval call binding the contract method 0x7f758a15.
//
// Solidity: function x76() constant returns(address)
func (_DSGContract *DSGContractSession) X76() (common.Address, error) {
	return _DSGContract.Contract.X76(&_DSGContract.CallOpts)
}

// X76 is a free data retrieval call binding the contract method 0x7f758a15.
//
// Solidity: function x76() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X76() (common.Address, error) {
	return _DSGContract.Contract.X76(&_DSGContract.CallOpts)
}

// X77 is a free data retrieval call binding the contract method 0x74ec408e.
//
// Solidity: function x77() constant returns(address)
func (_DSGContract *DSGContractCaller) X77(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x77")
	return *ret0, err
}

// X77 is a free data retrieval call binding the contract method 0x74ec408e.
//
// Solidity: function x77() constant returns(address)
func (_DSGContract *DSGContractSession) X77() (common.Address, error) {
	return _DSGContract.Contract.X77(&_DSGContract.CallOpts)
}

// X77 is a free data retrieval call binding the contract method 0x74ec408e.
//
// Solidity: function x77() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X77() (common.Address, error) {
	return _DSGContract.Contract.X77(&_DSGContract.CallOpts)
}

// X78 is a free data retrieval call binding the contract method 0x1e0a08db.
//
// Solidity: function x78() constant returns(address)
func (_DSGContract *DSGContractCaller) X78(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x78")
	return *ret0, err
}

// X78 is a free data retrieval call binding the contract method 0x1e0a08db.
//
// Solidity: function x78() constant returns(address)
func (_DSGContract *DSGContractSession) X78() (common.Address, error) {
	return _DSGContract.Contract.X78(&_DSGContract.CallOpts)
}

// X78 is a free data retrieval call binding the contract method 0x1e0a08db.
//
// Solidity: function x78() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X78() (common.Address, error) {
	return _DSGContract.Contract.X78(&_DSGContract.CallOpts)
}

// X79 is a free data retrieval call binding the contract method 0x41f01d3c.
//
// Solidity: function x79() constant returns(address)
func (_DSGContract *DSGContractCaller) X79(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x79")
	return *ret0, err
}

// X79 is a free data retrieval call binding the contract method 0x41f01d3c.
//
// Solidity: function x79() constant returns(address)
func (_DSGContract *DSGContractSession) X79() (common.Address, error) {
	return _DSGContract.Contract.X79(&_DSGContract.CallOpts)
}

// X79 is a free data retrieval call binding the contract method 0x41f01d3c.
//
// Solidity: function x79() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X79() (common.Address, error) {
	return _DSGContract.Contract.X79(&_DSGContract.CallOpts)
}

// X8 is a free data retrieval call binding the contract method 0xe66e4c3f.
//
// Solidity: function x8() constant returns(address)
func (_DSGContract *DSGContractCaller) X8(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x8")
	return *ret0, err
}

// X8 is a free data retrieval call binding the contract method 0xe66e4c3f.
//
// Solidity: function x8() constant returns(address)
func (_DSGContract *DSGContractSession) X8() (common.Address, error) {
	return _DSGContract.Contract.X8(&_DSGContract.CallOpts)
}

// X8 is a free data retrieval call binding the contract method 0xe66e4c3f.
//
// Solidity: function x8() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X8() (common.Address, error) {
	return _DSGContract.Contract.X8(&_DSGContract.CallOpts)
}

// X80 is a free data retrieval call binding the contract method 0x2e3f53c9.
//
// Solidity: function x80() constant returns(address)
func (_DSGContract *DSGContractCaller) X80(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x80")
	return *ret0, err
}

// X80 is a free data retrieval call binding the contract method 0x2e3f53c9.
//
// Solidity: function x80() constant returns(address)
func (_DSGContract *DSGContractSession) X80() (common.Address, error) {
	return _DSGContract.Contract.X80(&_DSGContract.CallOpts)
}

// X80 is a free data retrieval call binding the contract method 0x2e3f53c9.
//
// Solidity: function x80() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X80() (common.Address, error) {
	return _DSGContract.Contract.X80(&_DSGContract.CallOpts)
}

// X81 is a free data retrieval call binding the contract method 0xb02d7daa.
//
// Solidity: function x81() constant returns(address)
func (_DSGContract *DSGContractCaller) X81(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x81")
	return *ret0, err
}

// X81 is a free data retrieval call binding the contract method 0xb02d7daa.
//
// Solidity: function x81() constant returns(address)
func (_DSGContract *DSGContractSession) X81() (common.Address, error) {
	return _DSGContract.Contract.X81(&_DSGContract.CallOpts)
}

// X81 is a free data retrieval call binding the contract method 0xb02d7daa.
//
// Solidity: function x81() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X81() (common.Address, error) {
	return _DSGContract.Contract.X81(&_DSGContract.CallOpts)
}

// X82 is a free data retrieval call binding the contract method 0x5ae05570.
//
// Solidity: function x82() constant returns(address)
func (_DSGContract *DSGContractCaller) X82(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x82")
	return *ret0, err
}

// X82 is a free data retrieval call binding the contract method 0x5ae05570.
//
// Solidity: function x82() constant returns(address)
func (_DSGContract *DSGContractSession) X82() (common.Address, error) {
	return _DSGContract.Contract.X82(&_DSGContract.CallOpts)
}

// X82 is a free data retrieval call binding the contract method 0x5ae05570.
//
// Solidity: function x82() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X82() (common.Address, error) {
	return _DSGContract.Contract.X82(&_DSGContract.CallOpts)
}

// X83 is a free data retrieval call binding the contract method 0x60c72ae2.
//
// Solidity: function x83() constant returns(address)
func (_DSGContract *DSGContractCaller) X83(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x83")
	return *ret0, err
}

// X83 is a free data retrieval call binding the contract method 0x60c72ae2.
//
// Solidity: function x83() constant returns(address)
func (_DSGContract *DSGContractSession) X83() (common.Address, error) {
	return _DSGContract.Contract.X83(&_DSGContract.CallOpts)
}

// X83 is a free data retrieval call binding the contract method 0x60c72ae2.
//
// Solidity: function x83() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X83() (common.Address, error) {
	return _DSGContract.Contract.X83(&_DSGContract.CallOpts)
}

// X84 is a free data retrieval call binding the contract method 0x772eda28.
//
// Solidity: function x84() constant returns(address)
func (_DSGContract *DSGContractCaller) X84(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x84")
	return *ret0, err
}

// X84 is a free data retrieval call binding the contract method 0x772eda28.
//
// Solidity: function x84() constant returns(address)
func (_DSGContract *DSGContractSession) X84() (common.Address, error) {
	return _DSGContract.Contract.X84(&_DSGContract.CallOpts)
}

// X84 is a free data retrieval call binding the contract method 0x772eda28.
//
// Solidity: function x84() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X84() (common.Address, error) {
	return _DSGContract.Contract.X84(&_DSGContract.CallOpts)
}

// X85 is a free data retrieval call binding the contract method 0x04ba000d.
//
// Solidity: function x85() constant returns(address)
func (_DSGContract *DSGContractCaller) X85(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x85")
	return *ret0, err
}

// X85 is a free data retrieval call binding the contract method 0x04ba000d.
//
// Solidity: function x85() constant returns(address)
func (_DSGContract *DSGContractSession) X85() (common.Address, error) {
	return _DSGContract.Contract.X85(&_DSGContract.CallOpts)
}

// X85 is a free data retrieval call binding the contract method 0x04ba000d.
//
// Solidity: function x85() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X85() (common.Address, error) {
	return _DSGContract.Contract.X85(&_DSGContract.CallOpts)
}

// X86 is a free data retrieval call binding the contract method 0xac95e388.
//
// Solidity: function x86() constant returns(address)
func (_DSGContract *DSGContractCaller) X86(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x86")
	return *ret0, err
}

// X86 is a free data retrieval call binding the contract method 0xac95e388.
//
// Solidity: function x86() constant returns(address)
func (_DSGContract *DSGContractSession) X86() (common.Address, error) {
	return _DSGContract.Contract.X86(&_DSGContract.CallOpts)
}

// X86 is a free data retrieval call binding the contract method 0xac95e388.
//
// Solidity: function x86() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X86() (common.Address, error) {
	return _DSGContract.Contract.X86(&_DSGContract.CallOpts)
}

// X87 is a free data retrieval call binding the contract method 0xebc4e3d4.
//
// Solidity: function x87() constant returns(address)
func (_DSGContract *DSGContractCaller) X87(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x87")
	return *ret0, err
}

// X87 is a free data retrieval call binding the contract method 0xebc4e3d4.
//
// Solidity: function x87() constant returns(address)
func (_DSGContract *DSGContractSession) X87() (common.Address, error) {
	return _DSGContract.Contract.X87(&_DSGContract.CallOpts)
}

// X87 is a free data retrieval call binding the contract method 0xebc4e3d4.
//
// Solidity: function x87() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X87() (common.Address, error) {
	return _DSGContract.Contract.X87(&_DSGContract.CallOpts)
}

// X88 is a free data retrieval call binding the contract method 0x92770c96.
//
// Solidity: function x88() constant returns(address)
func (_DSGContract *DSGContractCaller) X88(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x88")
	return *ret0, err
}

// X88 is a free data retrieval call binding the contract method 0x92770c96.
//
// Solidity: function x88() constant returns(address)
func (_DSGContract *DSGContractSession) X88() (common.Address, error) {
	return _DSGContract.Contract.X88(&_DSGContract.CallOpts)
}

// X88 is a free data retrieval call binding the contract method 0x92770c96.
//
// Solidity: function x88() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X88() (common.Address, error) {
	return _DSGContract.Contract.X88(&_DSGContract.CallOpts)
}

// X89 is a free data retrieval call binding the contract method 0x457f2f2e.
//
// Solidity: function x89() constant returns(address)
func (_DSGContract *DSGContractCaller) X89(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x89")
	return *ret0, err
}

// X89 is a free data retrieval call binding the contract method 0x457f2f2e.
//
// Solidity: function x89() constant returns(address)
func (_DSGContract *DSGContractSession) X89() (common.Address, error) {
	return _DSGContract.Contract.X89(&_DSGContract.CallOpts)
}

// X89 is a free data retrieval call binding the contract method 0x457f2f2e.
//
// Solidity: function x89() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X89() (common.Address, error) {
	return _DSGContract.Contract.X89(&_DSGContract.CallOpts)
}

// X9 is a free data retrieval call binding the contract method 0xd6efd153.
//
// Solidity: function x9() constant returns(address)
func (_DSGContract *DSGContractCaller) X9(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x9")
	return *ret0, err
}

// X9 is a free data retrieval call binding the contract method 0xd6efd153.
//
// Solidity: function x9() constant returns(address)
func (_DSGContract *DSGContractSession) X9() (common.Address, error) {
	return _DSGContract.Contract.X9(&_DSGContract.CallOpts)
}

// X9 is a free data retrieval call binding the contract method 0xd6efd153.
//
// Solidity: function x9() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X9() (common.Address, error) {
	return _DSGContract.Contract.X9(&_DSGContract.CallOpts)
}

// X90 is a free data retrieval call binding the contract method 0x5c356ab5.
//
// Solidity: function x90() constant returns(address)
func (_DSGContract *DSGContractCaller) X90(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x90")
	return *ret0, err
}

// X90 is a free data retrieval call binding the contract method 0x5c356ab5.
//
// Solidity: function x90() constant returns(address)
func (_DSGContract *DSGContractSession) X90() (common.Address, error) {
	return _DSGContract.Contract.X90(&_DSGContract.CallOpts)
}

// X90 is a free data retrieval call binding the contract method 0x5c356ab5.
//
// Solidity: function x90() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X90() (common.Address, error) {
	return _DSGContract.Contract.X90(&_DSGContract.CallOpts)
}

// X91 is a free data retrieval call binding the contract method 0x9c856be7.
//
// Solidity: function x91() constant returns(address)
func (_DSGContract *DSGContractCaller) X91(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x91")
	return *ret0, err
}

// X91 is a free data retrieval call binding the contract method 0x9c856be7.
//
// Solidity: function x91() constant returns(address)
func (_DSGContract *DSGContractSession) X91() (common.Address, error) {
	return _DSGContract.Contract.X91(&_DSGContract.CallOpts)
}

// X91 is a free data retrieval call binding the contract method 0x9c856be7.
//
// Solidity: function x91() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X91() (common.Address, error) {
	return _DSGContract.Contract.X91(&_DSGContract.CallOpts)
}

// X92 is a free data retrieval call binding the contract method 0xda9f47ad.
//
// Solidity: function x92() constant returns(address)
func (_DSGContract *DSGContractCaller) X92(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x92")
	return *ret0, err
}

// X92 is a free data retrieval call binding the contract method 0xda9f47ad.
//
// Solidity: function x92() constant returns(address)
func (_DSGContract *DSGContractSession) X92() (common.Address, error) {
	return _DSGContract.Contract.X92(&_DSGContract.CallOpts)
}

// X92 is a free data retrieval call binding the contract method 0xda9f47ad.
//
// Solidity: function x92() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X92() (common.Address, error) {
	return _DSGContract.Contract.X92(&_DSGContract.CallOpts)
}

// X93 is a free data retrieval call binding the contract method 0x0f99cd32.
//
// Solidity: function x93() constant returns(address)
func (_DSGContract *DSGContractCaller) X93(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x93")
	return *ret0, err
}

// X93 is a free data retrieval call binding the contract method 0x0f99cd32.
//
// Solidity: function x93() constant returns(address)
func (_DSGContract *DSGContractSession) X93() (common.Address, error) {
	return _DSGContract.Contract.X93(&_DSGContract.CallOpts)
}

// X93 is a free data retrieval call binding the contract method 0x0f99cd32.
//
// Solidity: function x93() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X93() (common.Address, error) {
	return _DSGContract.Contract.X93(&_DSGContract.CallOpts)
}

// X94 is a free data retrieval call binding the contract method 0x5880ab1b.
//
// Solidity: function x94() constant returns(address)
func (_DSGContract *DSGContractCaller) X94(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x94")
	return *ret0, err
}

// X94 is a free data retrieval call binding the contract method 0x5880ab1b.
//
// Solidity: function x94() constant returns(address)
func (_DSGContract *DSGContractSession) X94() (common.Address, error) {
	return _DSGContract.Contract.X94(&_DSGContract.CallOpts)
}

// X94 is a free data retrieval call binding the contract method 0x5880ab1b.
//
// Solidity: function x94() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X94() (common.Address, error) {
	return _DSGContract.Contract.X94(&_DSGContract.CallOpts)
}

// X95 is a free data retrieval call binding the contract method 0xaa5b743d.
//
// Solidity: function x95() constant returns(address)
func (_DSGContract *DSGContractCaller) X95(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DSGContract.contract.Call(opts, out, "x95")
	return *ret0, err
}

// X95 is a free data retrieval call binding the contract method 0xaa5b743d.
//
// Solidity: function x95() constant returns(address)
func (_DSGContract *DSGContractSession) X95() (common.Address, error) {
	return _DSGContract.Contract.X95(&_DSGContract.CallOpts)
}

// X95 is a free data retrieval call binding the contract method 0xaa5b743d.
//
// Solidity: function x95() constant returns(address)
func (_DSGContract *DSGContractCallerSession) X95() (common.Address, error) {
	return _DSGContract.Contract.X95(&_DSGContract.CallOpts)
}

// Rotate is a paid mutator transaction binding the contract method 0xd992818d.
//
// Solidity: function rotate() returns()
func (_DSGContract *DSGContractTransactor) Rotate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSGContract.contract.Transact(opts, "rotate")
}

// Rotate is a paid mutator transaction binding the contract method 0xd992818d.
//
// Solidity: function rotate() returns()
func (_DSGContract *DSGContractSession) Rotate() (*types.Transaction, error) {
	return _DSGContract.Contract.Rotate(&_DSGContract.TransactOpts)
}

// Rotate is a paid mutator transaction binding the contract method 0xd992818d.
//
// Solidity: function rotate() returns()
func (_DSGContract *DSGContractTransactorSession) Rotate() (*types.Transaction, error) {
	return _DSGContract.Contract.Rotate(&_DSGContract.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_DSGContract *DSGContractTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSGContract.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_DSGContract *DSGContractSession) Stake() (*types.Transaction, error) {
	return _DSGContract.Contract.Stake(&_DSGContract.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_DSGContract *DSGContractTransactorSession) Stake() (*types.Transaction, error) {
	return _DSGContract.Contract.Stake(&_DSGContract.TransactOpts)
}

// UnStake is a paid mutator transaction binding the contract method 0x5d3eea91.
//
// Solidity: function unStake(uint256 _amount) returns()
func (_DSGContract *DSGContractTransactor) UnStake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _DSGContract.contract.Transact(opts, "unStake", _amount)
}

// UnStake is a paid mutator transaction binding the contract method 0x5d3eea91.
//
// Solidity: function unStake(uint256 _amount) returns()
func (_DSGContract *DSGContractSession) UnStake(_amount *big.Int) (*types.Transaction, error) {
	return _DSGContract.Contract.UnStake(&_DSGContract.TransactOpts, _amount)
}

// UnStake is a paid mutator transaction binding the contract method 0x5d3eea91.
//
// Solidity: function unStake(uint256 _amount) returns()
func (_DSGContract *DSGContractTransactorSession) UnStake(_amount *big.Int) (*types.Transaction, error) {
	return _DSGContract.Contract.UnStake(&_DSGContract.TransactOpts, _amount)
}

// DSGContractStakingEventIterator is returned from FilterStakingEvent and is used to iterate over the raw logs and unpacked data for StakingEvent events raised by the DSGContract contract.
type DSGContractStakingEventIterator struct {
	Event *DSGContractStakingEvent // Event containing the contract specifics and raw log

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
func (it *DSGContractStakingEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSGContractStakingEvent)
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
		it.Event = new(DSGContractStakingEvent)
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
func (it *DSGContractStakingEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSGContractStakingEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSGContractStakingEvent represents a StakingEvent event raised by the DSGContract contract.
type DSGContractStakingEvent struct {
	Candidate    common.Address
	Amount       *big.Int
	Timestamp    *big.Int
	StakeUnstake uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakingEvent is a free log retrieval operation binding the contract event 0x3585198d07e828811559e83f38a04e95c328af4edd28122306b10475843fdd35.
//
// Solidity: event StakingEvent(address indexed _candidate, uint256 _amount, uint256 _timestamp, uint8 indexed _stakeUnstake)
func (_DSGContract *DSGContractFilterer) FilterStakingEvent(opts *bind.FilterOpts, _candidate []common.Address, _stakeUnstake []uint8) (*DSGContractStakingEventIterator, error) {

	var _candidateRule []interface{}
	for _, _candidateItem := range _candidate {
		_candidateRule = append(_candidateRule, _candidateItem)
	}

	var _stakeUnstakeRule []interface{}
	for _, _stakeUnstakeItem := range _stakeUnstake {
		_stakeUnstakeRule = append(_stakeUnstakeRule, _stakeUnstakeItem)
	}

	logs, sub, err := _DSGContract.contract.FilterLogs(opts, "StakingEvent", _candidateRule, _stakeUnstakeRule)
	if err != nil {
		return nil, err
	}
	return &DSGContractStakingEventIterator{contract: _DSGContract.contract, event: "StakingEvent", logs: logs, sub: sub}, nil
}

// WatchStakingEvent is a free log subscription operation binding the contract event 0x3585198d07e828811559e83f38a04e95c328af4edd28122306b10475843fdd35.
//
// Solidity: event StakingEvent(address indexed _candidate, uint256 _amount, uint256 _timestamp, uint8 indexed _stakeUnstake)
func (_DSGContract *DSGContractFilterer) WatchStakingEvent(opts *bind.WatchOpts, sink chan<- *DSGContractStakingEvent, _candidate []common.Address, _stakeUnstake []uint8) (event.Subscription, error) {

	var _candidateRule []interface{}
	for _, _candidateItem := range _candidate {
		_candidateRule = append(_candidateRule, _candidateItem)
	}

	var _stakeUnstakeRule []interface{}
	for _, _stakeUnstakeItem := range _stakeUnstake {
		_stakeUnstakeRule = append(_stakeUnstakeRule, _stakeUnstakeItem)
	}

	logs, sub, err := _DSGContract.contract.WatchLogs(opts, "StakingEvent", _candidateRule, _stakeUnstakeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSGContractStakingEvent)
				if err := _DSGContract.contract.UnpackLog(event, "StakingEvent", log); err != nil {
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

// ParseStakingEvent is a log parse operation binding the contract event 0x3585198d07e828811559e83f38a04e95c328af4edd28122306b10475843fdd35.
//
// Solidity: event StakingEvent(address indexed _candidate, uint256 _amount, uint256 _timestamp, uint8 indexed _stakeUnstake)
func (_DSGContract *DSGContractFilterer) ParseStakingEvent(log types.Log) (*DSGContractStakingEvent, error) {
	event := new(DSGContractStakingEvent)
	if err := _DSGContract.contract.UnpackLog(event, "StakingEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820b99a416f0ba19df87f1465d1d6459a59b68e158d4adfe5e348db27482d91543964736f6c634300050b0032"

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
