// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dsgcontract

import (
	"strings"

	"github.com/unification-com/mainchain/accounts/abi"
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/core/types"
)

// DsgcontractABI is the input ABI used to generate the binding from.
const DsgcontractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"x49\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x85\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x14\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x44\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x93\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x59\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x18\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x75\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x24\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x45\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x12\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x78\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x63\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x7\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x40\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x80\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x39\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x62\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x1\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x17\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x30\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x23\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x43\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x79\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x46\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x89\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x2\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x94\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x61\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x82\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x90\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x70\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x35\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x83\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x25\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x22\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x31\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x47\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x65\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x77\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x84\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x55\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x10\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x71\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x76\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x48\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x51\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x60\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x13\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x21\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x32\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x88\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x42\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x73\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x64\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x91\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x20\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x66\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x38\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x3\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x16\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x58\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x95\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x72\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x86\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x54\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x81\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x36\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x50\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x29\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x37\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x74\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x27\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x67\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x53\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x15\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x33\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x4\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x5\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x9\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rotate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x92\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x8\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x26\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x69\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x6\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x19\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x57\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x87\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x41\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x11\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x52\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x68\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x34\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x56\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"x28\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DsgcontractBin is the compiled bytecode used for deploying new contracts.
const DsgcontractBin = `608060405234801561001057600080fd5b506152d2806100206000396000f300608060405260043610610456576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630418043d1461045b57806304ba000d146104b257806305f04234146105095780630ae490ad146105605780630f99cd32146105b757806313813beb1461060e57806314846dd314610665578063175d756b146106bc57806317bd117a146107135780631c0730861461076a5780631d0fb81a146107c15780631e0a08db146108185780631fc58dae1461086f578063282940a7146108c65780632ca07afd1461091d5780632e3f53c9146109745780633176cca2146109cb578063332bb13b14610a22578063343943bd14610a7957806337545b0014610ad057806339cb843514610b275780633f3f344414610b7e578063402cc27e14610bd557806341f01d3c14610c2c578063449136cf14610c83578063457f2f2e14610cda5780634ff1357114610d315780635880ab1b14610d88578063588d115014610ddf5780635ae0557014610e365780635c356ab514610e8d5780635d37039914610ee45780635e217f4214610f3b57806360c72ae214610f9257806368937ad514610fe95780636bc3d2b81461104057806372c019b31461109757806372e1b839146110ee578063745be9281461114557806374ec408e1461119c578063772eda28146111f35780637af6e8741461124a5780637d38e0ff146112a15780637ede4904146112f85780637f758a151461134f5780637fda3be2146113a657806388ce3426146113fd57806389b3bc3a146114545780638bd149b3146114ab5780638cadd7f614611502578063908df6b11461155957806392770c96146115b0578063943a0169146116075780639b85e4511461165e5780639bf3f61a146116b55780639c856be71461170c5780639ffa6a9f14611763578063a11be9af146117ba578063a19f482a14611811578063a3b7887314611868578063a4f930d8146118bf578063a87f6ab514611916578063aa5b743d1461196d578063ac1e43a0146119c4578063ac95e38814611a1b578063af59490414611a72578063b02d7daa14611ac9578063b407a0db14611b20578063b4feb11014611b77578063b54d94f814611bce578063b9fbc54314611c25578063ba16cefe14611c7c578063bc8c3c9e14611cd3578063bca63eb414611d2a578063bcfbbc4614611d81578063c0a1a94914611dd8578063c5f720e414611e2f578063cb6cee8a14611e86578063cd145f6514611edd578063d6efd15314611f34578063d992818d14611f8b578063da9f47ad14611fa2578063e66e4c3f14611ff9578063e749773f14612050578063eac637f2146120a7578063eaeb89f7146120fe578063eaecb01214612155578063eb5d92a0146121ac578063ebc4e3d414612203578063f24ec5c81461225a578063f25e3949146122b1578063f304e72014612308578063f721734f1461235f578063fd56a650146123b6578063fddbd3871461240d578063fe1878ef14612464575b600080fd5b34801561046757600080fd5b506104706124bb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104be57600080fd5b506104c76124e1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561051557600080fd5b5061051e612507565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561056c57600080fd5b5061057561252d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105c357600080fd5b506105cc612553565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561061a57600080fd5b50610623612579565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561067157600080fd5b5061067a61259f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156106c857600080fd5b506106d16125c5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561071f57600080fd5b506107286125eb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561077657600080fd5b5061077f612611565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156107cd57600080fd5b506107d6612637565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561082457600080fd5b5061082d61265d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561087b57600080fd5b50610884612683565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156108d257600080fd5b506108db6126a9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561092957600080fd5b506109326126cf565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561098057600080fd5b506109896126f5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156109d757600080fd5b506109e061271b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610a2e57600080fd5b50610a37612741565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610a8557600080fd5b50610a8e612767565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610adc57600080fd5b50610ae561278c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610b3357600080fd5b50610b3c6127b2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610b8a57600080fd5b50610b936127d8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610be157600080fd5b50610bea6127fe565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610c3857600080fd5b50610c41612824565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610c8f57600080fd5b50610c9861284a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610ce657600080fd5b50610cef612870565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610d3d57600080fd5b50610d46612896565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610d9457600080fd5b50610d9d6128bc565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610deb57600080fd5b50610df46128e2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610e4257600080fd5b50610e4b612908565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610e9957600080fd5b50610ea261292e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610ef057600080fd5b50610ef9612954565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610f4757600080fd5b50610f5061297a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610f9e57600080fd5b50610fa76129a0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610ff557600080fd5b50610ffe6129c6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561104c57600080fd5b506110556129ec565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156110a357600080fd5b506110ac612a12565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156110fa57600080fd5b50611103612a38565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561115157600080fd5b5061115a612a5e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156111a857600080fd5b506111b1612a84565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156111ff57600080fd5b50611208612aaa565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561125657600080fd5b5061125f612ad0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156112ad57600080fd5b506112b6612af6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561130457600080fd5b5061130d612b1c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561135b57600080fd5b50611364612b42565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156113b257600080fd5b506113bb612b68565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561140957600080fd5b50611412612b8e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561146057600080fd5b50611469612bb4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156114b757600080fd5b506114c0612bda565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561150e57600080fd5b50611517612c00565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561156557600080fd5b5061156e612c26565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156115bc57600080fd5b506115c5612c4c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561161357600080fd5b5061161c612c72565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561166a57600080fd5b50611673612c98565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156116c157600080fd5b506116ca612cbe565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561171857600080fd5b50611721612ce4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561176f57600080fd5b50611778612d0a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156117c657600080fd5b506117cf612d30565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561181d57600080fd5b50611826612d56565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561187457600080fd5b5061187d612d7c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156118cb57600080fd5b506118d4612da2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561192257600080fd5b5061192b612dc8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561197957600080fd5b50611982612dee565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156119d057600080fd5b506119d9612e14565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611a2757600080fd5b50611a30612e3a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611a7e57600080fd5b50611a87612e60565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611ad557600080fd5b50611ade612e86565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611b2c57600080fd5b50611b35612eac565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611b8357600080fd5b50611b8c612ed2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611bda57600080fd5b50611be3612ef8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611c3157600080fd5b50611c3a612f1e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611c8857600080fd5b50611c91612f44565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611cdf57600080fd5b50611ce8612f6a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611d3657600080fd5b50611d3f612f90565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611d8d57600080fd5b50611d96612fb6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611de457600080fd5b50611ded612fdc565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611e3b57600080fd5b50611e44613002565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611e9257600080fd5b50611e9b613028565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611ee957600080fd5b50611ef261304e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611f4057600080fd5b50611f49613074565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015611f9757600080fd5b50611fa061309a565b005b348015611fae57600080fd5b50611fb761506c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561200557600080fd5b5061200e615092565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561205c57600080fd5b506120656150b8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156120b357600080fd5b506120bc6150de565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561210a57600080fd5b50612113615104565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561216157600080fd5b5061216a61512a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156121b857600080fd5b506121c1615150565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561220f57600080fd5b50612218615176565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561226657600080fd5b5061226f61519c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156122bd57600080fd5b506122c66151c2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561231457600080fd5b5061231d6151e8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561236b57600080fd5b5061237461520e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156123c257600080fd5b506123cb615234565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561241957600080fd5b5061242261525a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561247057600080fd5b50612479615280565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b603060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b73f30f951b0426f8bf37ac71967407081358df7a7b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156130e85761506a565b722a956804bad8dcad148abff71515f9b057f7e06000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550723adc30a6f4db59d2698e3d3029fd1ba68b6b15600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550724a435f1d54aa5cc9fcfa0feb6b8c4a428bbb93600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550725aa882b1708f2e032b1e6ac8fb92fb2e8f09e2600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550726a60f879db155cb01dcc70452c0a935fcb15d3600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550727aa694e0fd812b5f3d6d42654c8f8f33b58a8a600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550728a1c6df0be8682dfb87f8d1265f90f1d181647600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550729a2cbe70409e8ab9f2fee621354578916f0118600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073010a7c34c6beb24b2a92febd123d780788cf2235600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073011a6b9f30767ad1266ea9d3307b91470f6e22cd600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073012ab4f14f3be13d4649be31ca4d65947ef95554600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073013a0e389b4ccb117b0809fc1ab4a1b4b3f18f1f600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073014aa8687ca6b10afd7aaf101efa215ab52ca075600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073015ac19e712a6b839ffd14acacbe58151911b916600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073016a4575a058d7f0c73641472f6ba9aa09d9ebc6600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073017a31bf0827b5280bfcf2e2863a80ef4bf03152600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073018a31e3f69e9c30086e03d73654edff06171d96601060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073019afd07eb5a3198597b4845db932254fc11d9b9601160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073020a618a1bb1c751bd9d7c985edd4f3d9c5fc1a0601260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073021a8866723a67828747f0239d3123098568e86e601360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073022ab4c0b02912dd100be27646b594849373e84b601460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073023af3313253a7b2338992042a21479041820b44601560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073024afc0f101527bd62d86c85ab18769b9084af74601660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073025ad5dc2e38368ca489de9bafc99ea845fa93dd601760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073026abcd44b611511f2950fc45a78be446ed21afb601860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073027ac37befdaae0b9fa316a3d1436a5a6256ac16601960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073028ac2879de90937d22544ba893b1ba04e161e19601a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073029af5c22f1375ddaeb41f0ac96893de2259b68b601b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073030a728b7e2ad365d2fdde05c0e536c8f31a5c6f601c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073031a2b49f84bccd3c18f35ddaa07d72970412a67601d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073032aae65c9f7135576f2e58a23b8430b30e00348601e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073033a10b8526e7c8e70d5b41e31f254a060e0ec2d601f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073034a7f2494ecb251cac89933da04dec1709e8565602060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073035a60510dbd42f3a65d58fb990066e9b355adc6602160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073036af29b42d58f44bf7ca73d3a9b18f657794cfe602260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073037aa30cd556dfb4d51a8880790c077737dd95c7602360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073038a0e937609f695a552f1f3ae6c213f287f84aa602460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073039a28d44b5dad2d3d4772da3d2e0d083a18647b602560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073040a3dcfabdf555797feb37bc1e60315753930a1602660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073041a92ec07212f45b4887c800edfd7c1a25c6f81602760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073042a956ec294b85e639a43bea2c93ff8c78bac28602860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073043a6bde9d509712f0d93cf9196e62badf6afaf2602960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073044ac1a25dd8d6e78a272f0e7b7798aae593972c602a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073045aed6663d3b239fe31477521c8ad962570d05f602b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073046aa3f2a00562fea28e354a8b28a37539417973602c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073047a91571a4725efa09981cf4e829fa6670e6119602d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073048a6c75dc7c0cc5df4562458bee72ba3b084195602e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073049a47f51837d54a1c40e91469935c880f7af720602f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073050a7673bace7c43787798b8dd280ddd4a17ef12603060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073051a24373f5ad334e2292b336b497097b702f7b3603160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073052a987fa6fc78a0b2324bcb47db73fda2ec0761603260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073053ae7420dbcc57cb03856c35b9e7ceb764735e9603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073054ab96c0886050a20317ee8cfad593c3a17f5c6603460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073055a1668669a3f1e3602f1c346cf6ea5ead81fe4603560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073056ae2baea2b1572d9293ef8c688d81a84790ca4603660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073057adc1e8ccf640767aa7526cd46b6559b88de0c603760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073058af6bb396153079d6d4cd9bceafc87cf5c97a7603860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073059acc4768532d8d1888ef4d1ae6c5fc55e1285e603960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073060ad483059b2d55542168bd13746f4bc9d90800603a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073061a99d798b34e91f94696f6a704bc0545526964603b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073062a06c8a3ad35dd60bcaaf403cbdd140034e2f9603c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073063ab9229016eca057e026b551bd587f3196ea58603d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073064a975ffd46ee34ac5f43c87d43fd6ef55f8fd3603e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073065af52373eded1a2db013e5fe0c49b34325c3aa603f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073066ad8c53e8b458290a01525da6dafe127eb37b0604060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073067af3fbb6a30c45b4c2855877315b6d528674b1604160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073068afdc63964f3cfd6e61615bfd48ce0c61c292b604260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073069a729be61423a08f1bfa4471ad203c5bc8ef0c604360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073070a3b781c3d5ad9fbc8b616c3d1a683d0bf7abc604460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073071a9130bd8501c80658c2336ae08a30eb723a0a604560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073072a6c85f39e9c97603ff6162d5ef053e112ee05604660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073073a5d1cc708b853ff7fc05cd1eb6ce066a2d337604760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073074a7eec2044540d8c031c817636966e5d96d113604860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073075ad3e9fa874fd60d4d058750c2457f8c40d9af604960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073076ac70c5605893b0cb224b115053e654c2c11bb604a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073077a938f9e52b643d47ef24fb203ea7206573f00604b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073078a6d1ce5abcff8cd7a979b67b193391975e3dd604c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073079a9b6d74b3b515ca141cab5c6789358721358b604d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073080addb3ed8fa2c583b627bdb86d63780623ddd1604e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073081a9af6fad95002971ef7775b547e828223f03c604f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073082a93c3c886429e4ea5d274a96851705c00b7b9605060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073083ab59ce9fce039c1546219855ace5ec472decf605160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073084a5bd013c104229161958359cfd9da63f85513605260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073085ac9744d7afb313edff60a370cf0b97a5dff84605360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073086a902bbc478db69f6b36b54d5d5954f19b893f605460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073087a204b4a5841550181af660a07daf11914b12b605560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073088a4555c3bad6b770fd2b6b6f112d86b0179b3c605660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073089adb67d2aa720bb6d37ac49784ee25b4eb3adf605760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073090a40813e7e679026f10534fdbf10d216c5315b605860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073091ad8b556a910123abdc002de39269cdf198e1d605960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073092a93e3ef9d79df3145487448388f42863e1104605a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073093ac3d0ae17afef0badeea471ec6466511e9c42605b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073094af16bd3700e2ee4709776da4c3334f6163c69605c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073095a2be3334c23eb8d2a7d463027dfd0cdbe21b1605d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550721a320943d4535e93d31e4a65a6e21c5df375d7605e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b565b605b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b605660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b604360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b603760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a723058204cb2a37c1462634ef6a460a70116310a1a48f2c52ef8ce23977278248a3735020029`

// DeployDsgcontract deploys a new Ethereum contract, binding an instance of Dsgcontract to it.
func DeployDsgcontract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dsgcontract, error) {
	parsed, err := abi.JSON(strings.NewReader(DsgcontractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DsgcontractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dsgcontract{DsgcontractCaller: DsgcontractCaller{contract: contract}, DsgcontractTransactor: DsgcontractTransactor{contract: contract}, DsgcontractFilterer: DsgcontractFilterer{contract: contract}}, nil
}

// Dsgcontract is an auto generated Go binding around an Ethereum contract.
type Dsgcontract struct {
	DsgcontractCaller     // Read-only binding to the contract
	DsgcontractTransactor // Write-only binding to the contract
	DsgcontractFilterer   // Log filterer for contract events
}

// DsgcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DsgcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsgcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DsgcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsgcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DsgcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsgcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DsgcontractSession struct {
	Contract     *Dsgcontract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DsgcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DsgcontractCallerSession struct {
	Contract *DsgcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DsgcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DsgcontractTransactorSession struct {
	Contract     *DsgcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DsgcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DsgcontractRaw struct {
	Contract *Dsgcontract // Generic contract binding to access the raw methods on
}

// DsgcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DsgcontractCallerRaw struct {
	Contract *DsgcontractCaller // Generic read-only contract binding to access the raw methods on
}

// DsgcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DsgcontractTransactorRaw struct {
	Contract *DsgcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDsgcontract creates a new instance of Dsgcontract, bound to a specific deployed contract.
func NewDsgcontract(address common.Address, backend bind.ContractBackend) (*Dsgcontract, error) {
	contract, err := bindDsgcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dsgcontract{DsgcontractCaller: DsgcontractCaller{contract: contract}, DsgcontractTransactor: DsgcontractTransactor{contract: contract}, DsgcontractFilterer: DsgcontractFilterer{contract: contract}}, nil
}

// NewDsgcontractCaller creates a new read-only instance of Dsgcontract, bound to a specific deployed contract.
func NewDsgcontractCaller(address common.Address, caller bind.ContractCaller) (*DsgcontractCaller, error) {
	contract, err := bindDsgcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DsgcontractCaller{contract: contract}, nil
}

// NewDsgcontractTransactor creates a new write-only instance of Dsgcontract, bound to a specific deployed contract.
func NewDsgcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*DsgcontractTransactor, error) {
	contract, err := bindDsgcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DsgcontractTransactor{contract: contract}, nil
}

// NewDsgcontractFilterer creates a new log filterer instance of Dsgcontract, bound to a specific deployed contract.
func NewDsgcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*DsgcontractFilterer, error) {
	contract, err := bindDsgcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DsgcontractFilterer{contract: contract}, nil
}

// bindDsgcontract binds a generic wrapper to an already deployed contract.
func bindDsgcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DsgcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dsgcontract *DsgcontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dsgcontract.Contract.DsgcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dsgcontract *DsgcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dsgcontract.Contract.DsgcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dsgcontract *DsgcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dsgcontract.Contract.DsgcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dsgcontract *DsgcontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dsgcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dsgcontract *DsgcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dsgcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dsgcontract *DsgcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dsgcontract.Contract.contract.Transact(opts, method, params...)
}

// X1 is a free data retrieval call binding the contract method 0x343943bd.
//
// Solidity: function x1() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x1")
	return *ret0, err
}

// X1 is a free data retrieval call binding the contract method 0x343943bd.
//
// Solidity: function x1() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X1() (common.Address, error) {
	return _Dsgcontract.Contract.X1(&_Dsgcontract.CallOpts)
}

// X1 is a free data retrieval call binding the contract method 0x343943bd.
//
// Solidity: function x1() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X1() (common.Address, error) {
	return _Dsgcontract.Contract.X1(&_Dsgcontract.CallOpts)
}

// X10 is a free data retrieval call binding the contract method 0x7d38e0ff.
//
// Solidity: function x10() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X10(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x10")
	return *ret0, err
}

// X10 is a free data retrieval call binding the contract method 0x7d38e0ff.
//
// Solidity: function x10() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X10() (common.Address, error) {
	return _Dsgcontract.Contract.X10(&_Dsgcontract.CallOpts)
}

// X10 is a free data retrieval call binding the contract method 0x7d38e0ff.
//
// Solidity: function x10() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X10() (common.Address, error) {
	return _Dsgcontract.Contract.X10(&_Dsgcontract.CallOpts)
}

// X11 is a free data retrieval call binding the contract method 0xf25e3949.
//
// Solidity: function x11() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X11(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x11")
	return *ret0, err
}

// X11 is a free data retrieval call binding the contract method 0xf25e3949.
//
// Solidity: function x11() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X11() (common.Address, error) {
	return _Dsgcontract.Contract.X11(&_Dsgcontract.CallOpts)
}

// X11 is a free data retrieval call binding the contract method 0xf25e3949.
//
// Solidity: function x11() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X11() (common.Address, error) {
	return _Dsgcontract.Contract.X11(&_Dsgcontract.CallOpts)
}

// X12 is a free data retrieval call binding the contract method 0x1d0fb81a.
//
// Solidity: function x12() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X12(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x12")
	return *ret0, err
}

// X12 is a free data retrieval call binding the contract method 0x1d0fb81a.
//
// Solidity: function x12() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X12() (common.Address, error) {
	return _Dsgcontract.Contract.X12(&_Dsgcontract.CallOpts)
}

// X12 is a free data retrieval call binding the contract method 0x1d0fb81a.
//
// Solidity: function x12() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X12() (common.Address, error) {
	return _Dsgcontract.Contract.X12(&_Dsgcontract.CallOpts)
}

// X13 is a free data retrieval call binding the contract method 0x8bd149b3.
//
// Solidity: function x13() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X13(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x13")
	return *ret0, err
}

// X13 is a free data retrieval call binding the contract method 0x8bd149b3.
//
// Solidity: function x13() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X13() (common.Address, error) {
	return _Dsgcontract.Contract.X13(&_Dsgcontract.CallOpts)
}

// X13 is a free data retrieval call binding the contract method 0x8bd149b3.
//
// Solidity: function x13() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X13() (common.Address, error) {
	return _Dsgcontract.Contract.X13(&_Dsgcontract.CallOpts)
}

// X14 is a free data retrieval call binding the contract method 0x05f04234.
//
// Solidity: function x14() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X14(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x14")
	return *ret0, err
}

// X14 is a free data retrieval call binding the contract method 0x05f04234.
//
// Solidity: function x14() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X14() (common.Address, error) {
	return _Dsgcontract.Contract.X14(&_Dsgcontract.CallOpts)
}

// X14 is a free data retrieval call binding the contract method 0x05f04234.
//
// Solidity: function x14() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X14() (common.Address, error) {
	return _Dsgcontract.Contract.X14(&_Dsgcontract.CallOpts)
}

// X15 is a free data retrieval call binding the contract method 0xc0a1a949.
//
// Solidity: function x15() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X15(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x15")
	return *ret0, err
}

// X15 is a free data retrieval call binding the contract method 0xc0a1a949.
//
// Solidity: function x15() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X15() (common.Address, error) {
	return _Dsgcontract.Contract.X15(&_Dsgcontract.CallOpts)
}

// X15 is a free data retrieval call binding the contract method 0xc0a1a949.
//
// Solidity: function x15() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X15() (common.Address, error) {
	return _Dsgcontract.Contract.X15(&_Dsgcontract.CallOpts)
}

// X16 is a free data retrieval call binding the contract method 0xa4f930d8.
//
// Solidity: function x16() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X16(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x16")
	return *ret0, err
}

// X16 is a free data retrieval call binding the contract method 0xa4f930d8.
//
// Solidity: function x16() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X16() (common.Address, error) {
	return _Dsgcontract.Contract.X16(&_Dsgcontract.CallOpts)
}

// X16 is a free data retrieval call binding the contract method 0xa4f930d8.
//
// Solidity: function x16() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X16() (common.Address, error) {
	return _Dsgcontract.Contract.X16(&_Dsgcontract.CallOpts)
}

// X17 is a free data retrieval call binding the contract method 0x37545b00.
//
// Solidity: function x17() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X17(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x17")
	return *ret0, err
}

// X17 is a free data retrieval call binding the contract method 0x37545b00.
//
// Solidity: function x17() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X17() (common.Address, error) {
	return _Dsgcontract.Contract.X17(&_Dsgcontract.CallOpts)
}

// X17 is a free data retrieval call binding the contract method 0x37545b00.
//
// Solidity: function x17() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X17() (common.Address, error) {
	return _Dsgcontract.Contract.X17(&_Dsgcontract.CallOpts)
}

// X18 is a free data retrieval call binding the contract method 0x14846dd3.
//
// Solidity: function x18() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X18(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x18")
	return *ret0, err
}

// X18 is a free data retrieval call binding the contract method 0x14846dd3.
//
// Solidity: function x18() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X18() (common.Address, error) {
	return _Dsgcontract.Contract.X18(&_Dsgcontract.CallOpts)
}

// X18 is a free data retrieval call binding the contract method 0x14846dd3.
//
// Solidity: function x18() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X18() (common.Address, error) {
	return _Dsgcontract.Contract.X18(&_Dsgcontract.CallOpts)
}

// X19 is a free data retrieval call binding the contract method 0xeaecb012.
//
// Solidity: function x19() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X19(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x19")
	return *ret0, err
}

// X19 is a free data retrieval call binding the contract method 0xeaecb012.
//
// Solidity: function x19() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X19() (common.Address, error) {
	return _Dsgcontract.Contract.X19(&_Dsgcontract.CallOpts)
}

// X19 is a free data retrieval call binding the contract method 0xeaecb012.
//
// Solidity: function x19() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X19() (common.Address, error) {
	return _Dsgcontract.Contract.X19(&_Dsgcontract.CallOpts)
}

// X2 is a free data retrieval call binding the contract method 0x4ff13571.
//
// Solidity: function x2() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X2(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x2")
	return *ret0, err
}

// X2 is a free data retrieval call binding the contract method 0x4ff13571.
//
// Solidity: function x2() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X2() (common.Address, error) {
	return _Dsgcontract.Contract.X2(&_Dsgcontract.CallOpts)
}

// X2 is a free data retrieval call binding the contract method 0x4ff13571.
//
// Solidity: function x2() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X2() (common.Address, error) {
	return _Dsgcontract.Contract.X2(&_Dsgcontract.CallOpts)
}

// X20 is a free data retrieval call binding the contract method 0x9ffa6a9f.
//
// Solidity: function x20() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X20(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x20")
	return *ret0, err
}

// X20 is a free data retrieval call binding the contract method 0x9ffa6a9f.
//
// Solidity: function x20() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X20() (common.Address, error) {
	return _Dsgcontract.Contract.X20(&_Dsgcontract.CallOpts)
}

// X20 is a free data retrieval call binding the contract method 0x9ffa6a9f.
//
// Solidity: function x20() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X20() (common.Address, error) {
	return _Dsgcontract.Contract.X20(&_Dsgcontract.CallOpts)
}

// X21 is a free data retrieval call binding the contract method 0x8cadd7f6.
//
// Solidity: function x21() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X21(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x21")
	return *ret0, err
}

// X21 is a free data retrieval call binding the contract method 0x8cadd7f6.
//
// Solidity: function x21() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X21() (common.Address, error) {
	return _Dsgcontract.Contract.X21(&_Dsgcontract.CallOpts)
}

// X21 is a free data retrieval call binding the contract method 0x8cadd7f6.
//
// Solidity: function x21() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X21() (common.Address, error) {
	return _Dsgcontract.Contract.X21(&_Dsgcontract.CallOpts)
}

// X22 is a free data retrieval call binding the contract method 0x6bc3d2b8.
//
// Solidity: function x22() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X22(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x22")
	return *ret0, err
}

// X22 is a free data retrieval call binding the contract method 0x6bc3d2b8.
//
// Solidity: function x22() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X22() (common.Address, error) {
	return _Dsgcontract.Contract.X22(&_Dsgcontract.CallOpts)
}

// X22 is a free data retrieval call binding the contract method 0x6bc3d2b8.
//
// Solidity: function x22() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X22() (common.Address, error) {
	return _Dsgcontract.Contract.X22(&_Dsgcontract.CallOpts)
}

// X23 is a free data retrieval call binding the contract method 0x3f3f3444.
//
// Solidity: function x23() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X23(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x23")
	return *ret0, err
}

// X23 is a free data retrieval call binding the contract method 0x3f3f3444.
//
// Solidity: function x23() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X23() (common.Address, error) {
	return _Dsgcontract.Contract.X23(&_Dsgcontract.CallOpts)
}

// X23 is a free data retrieval call binding the contract method 0x3f3f3444.
//
// Solidity: function x23() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X23() (common.Address, error) {
	return _Dsgcontract.Contract.X23(&_Dsgcontract.CallOpts)
}

// X24 is a free data retrieval call binding the contract method 0x17bd117a.
//
// Solidity: function x24() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X24(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x24")
	return *ret0, err
}

// X24 is a free data retrieval call binding the contract method 0x17bd117a.
//
// Solidity: function x24() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X24() (common.Address, error) {
	return _Dsgcontract.Contract.X24(&_Dsgcontract.CallOpts)
}

// X24 is a free data retrieval call binding the contract method 0x17bd117a.
//
// Solidity: function x24() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X24() (common.Address, error) {
	return _Dsgcontract.Contract.X24(&_Dsgcontract.CallOpts)
}

// X25 is a free data retrieval call binding the contract method 0x68937ad5.
//
// Solidity: function x25() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X25(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x25")
	return *ret0, err
}

// X25 is a free data retrieval call binding the contract method 0x68937ad5.
//
// Solidity: function x25() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X25() (common.Address, error) {
	return _Dsgcontract.Contract.X25(&_Dsgcontract.CallOpts)
}

// X25 is a free data retrieval call binding the contract method 0x68937ad5.
//
// Solidity: function x25() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X25() (common.Address, error) {
	return _Dsgcontract.Contract.X25(&_Dsgcontract.CallOpts)
}

// X26 is a free data retrieval call binding the contract method 0xe749773f.
//
// Solidity: function x26() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X26(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x26")
	return *ret0, err
}

// X26 is a free data retrieval call binding the contract method 0xe749773f.
//
// Solidity: function x26() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X26() (common.Address, error) {
	return _Dsgcontract.Contract.X26(&_Dsgcontract.CallOpts)
}

// X26 is a free data retrieval call binding the contract method 0xe749773f.
//
// Solidity: function x26() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X26() (common.Address, error) {
	return _Dsgcontract.Contract.X26(&_Dsgcontract.CallOpts)
}

// X27 is a free data retrieval call binding the contract method 0xbc8c3c9e.
//
// Solidity: function x27() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X27(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x27")
	return *ret0, err
}

// X27 is a free data retrieval call binding the contract method 0xbc8c3c9e.
//
// Solidity: function x27() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X27() (common.Address, error) {
	return _Dsgcontract.Contract.X27(&_Dsgcontract.CallOpts)
}

// X27 is a free data retrieval call binding the contract method 0xbc8c3c9e.
//
// Solidity: function x27() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X27() (common.Address, error) {
	return _Dsgcontract.Contract.X27(&_Dsgcontract.CallOpts)
}

// X28 is a free data retrieval call binding the contract method 0xfe1878ef.
//
// Solidity: function x28() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X28(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x28")
	return *ret0, err
}

// X28 is a free data retrieval call binding the contract method 0xfe1878ef.
//
// Solidity: function x28() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X28() (common.Address, error) {
	return _Dsgcontract.Contract.X28(&_Dsgcontract.CallOpts)
}

// X28 is a free data retrieval call binding the contract method 0xfe1878ef.
//
// Solidity: function x28() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X28() (common.Address, error) {
	return _Dsgcontract.Contract.X28(&_Dsgcontract.CallOpts)
}

// X29 is a free data retrieval call binding the contract method 0xb54d94f8.
//
// Solidity: function x29() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X29(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x29")
	return *ret0, err
}

// X29 is a free data retrieval call binding the contract method 0xb54d94f8.
//
// Solidity: function x29() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X29() (common.Address, error) {
	return _Dsgcontract.Contract.X29(&_Dsgcontract.CallOpts)
}

// X29 is a free data retrieval call binding the contract method 0xb54d94f8.
//
// Solidity: function x29() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X29() (common.Address, error) {
	return _Dsgcontract.Contract.X29(&_Dsgcontract.CallOpts)
}

// X3 is a free data retrieval call binding the contract method 0xa3b78873.
//
// Solidity: function x3() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X3(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x3")
	return *ret0, err
}

// X3 is a free data retrieval call binding the contract method 0xa3b78873.
//
// Solidity: function x3() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X3() (common.Address, error) {
	return _Dsgcontract.Contract.X3(&_Dsgcontract.CallOpts)
}

// X3 is a free data retrieval call binding the contract method 0xa3b78873.
//
// Solidity: function x3() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X3() (common.Address, error) {
	return _Dsgcontract.Contract.X3(&_Dsgcontract.CallOpts)
}

// X30 is a free data retrieval call binding the contract method 0x39cb8435.
//
// Solidity: function x30() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X30(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x30")
	return *ret0, err
}

// X30 is a free data retrieval call binding the contract method 0x39cb8435.
//
// Solidity: function x30() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X30() (common.Address, error) {
	return _Dsgcontract.Contract.X30(&_Dsgcontract.CallOpts)
}

// X30 is a free data retrieval call binding the contract method 0x39cb8435.
//
// Solidity: function x30() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X30() (common.Address, error) {
	return _Dsgcontract.Contract.X30(&_Dsgcontract.CallOpts)
}

// X31 is a free data retrieval call binding the contract method 0x72c019b3.
//
// Solidity: function x31() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X31(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x31")
	return *ret0, err
}

// X31 is a free data retrieval call binding the contract method 0x72c019b3.
//
// Solidity: function x31() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X31() (common.Address, error) {
	return _Dsgcontract.Contract.X31(&_Dsgcontract.CallOpts)
}

// X31 is a free data retrieval call binding the contract method 0x72c019b3.
//
// Solidity: function x31() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X31() (common.Address, error) {
	return _Dsgcontract.Contract.X31(&_Dsgcontract.CallOpts)
}

// X32 is a free data retrieval call binding the contract method 0x908df6b1.
//
// Solidity: function x32() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X32(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x32")
	return *ret0, err
}

// X32 is a free data retrieval call binding the contract method 0x908df6b1.
//
// Solidity: function x32() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X32() (common.Address, error) {
	return _Dsgcontract.Contract.X32(&_Dsgcontract.CallOpts)
}

// X32 is a free data retrieval call binding the contract method 0x908df6b1.
//
// Solidity: function x32() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X32() (common.Address, error) {
	return _Dsgcontract.Contract.X32(&_Dsgcontract.CallOpts)
}

// X33 is a free data retrieval call binding the contract method 0xc5f720e4.
//
// Solidity: function x33() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X33(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x33")
	return *ret0, err
}

// X33 is a free data retrieval call binding the contract method 0xc5f720e4.
//
// Solidity: function x33() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X33() (common.Address, error) {
	return _Dsgcontract.Contract.X33(&_Dsgcontract.CallOpts)
}

// X33 is a free data retrieval call binding the contract method 0xc5f720e4.
//
// Solidity: function x33() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X33() (common.Address, error) {
	return _Dsgcontract.Contract.X33(&_Dsgcontract.CallOpts)
}

// X34 is a free data retrieval call binding the contract method 0xfd56a650.
//
// Solidity: function x34() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X34(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x34")
	return *ret0, err
}

// X34 is a free data retrieval call binding the contract method 0xfd56a650.
//
// Solidity: function x34() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X34() (common.Address, error) {
	return _Dsgcontract.Contract.X34(&_Dsgcontract.CallOpts)
}

// X34 is a free data retrieval call binding the contract method 0xfd56a650.
//
// Solidity: function x34() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X34() (common.Address, error) {
	return _Dsgcontract.Contract.X34(&_Dsgcontract.CallOpts)
}

// X35 is a free data retrieval call binding the contract method 0x5e217f42.
//
// Solidity: function x35() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X35(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x35")
	return *ret0, err
}

// X35 is a free data retrieval call binding the contract method 0x5e217f42.
//
// Solidity: function x35() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X35() (common.Address, error) {
	return _Dsgcontract.Contract.X35(&_Dsgcontract.CallOpts)
}

// X35 is a free data retrieval call binding the contract method 0x5e217f42.
//
// Solidity: function x35() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X35() (common.Address, error) {
	return _Dsgcontract.Contract.X35(&_Dsgcontract.CallOpts)
}

// X36 is a free data retrieval call binding the contract method 0xb407a0db.
//
// Solidity: function x36() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X36(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x36")
	return *ret0, err
}

// X36 is a free data retrieval call binding the contract method 0xb407a0db.
//
// Solidity: function x36() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X36() (common.Address, error) {
	return _Dsgcontract.Contract.X36(&_Dsgcontract.CallOpts)
}

// X36 is a free data retrieval call binding the contract method 0xb407a0db.
//
// Solidity: function x36() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X36() (common.Address, error) {
	return _Dsgcontract.Contract.X36(&_Dsgcontract.CallOpts)
}

// X37 is a free data retrieval call binding the contract method 0xb9fbc543.
//
// Solidity: function x37() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X37(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x37")
	return *ret0, err
}

// X37 is a free data retrieval call binding the contract method 0xb9fbc543.
//
// Solidity: function x37() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X37() (common.Address, error) {
	return _Dsgcontract.Contract.X37(&_Dsgcontract.CallOpts)
}

// X37 is a free data retrieval call binding the contract method 0xb9fbc543.
//
// Solidity: function x37() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X37() (common.Address, error) {
	return _Dsgcontract.Contract.X37(&_Dsgcontract.CallOpts)
}

// X38 is a free data retrieval call binding the contract method 0xa19f482a.
//
// Solidity: function x38() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X38(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x38")
	return *ret0, err
}

// X38 is a free data retrieval call binding the contract method 0xa19f482a.
//
// Solidity: function x38() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X38() (common.Address, error) {
	return _Dsgcontract.Contract.X38(&_Dsgcontract.CallOpts)
}

// X38 is a free data retrieval call binding the contract method 0xa19f482a.
//
// Solidity: function x38() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X38() (common.Address, error) {
	return _Dsgcontract.Contract.X38(&_Dsgcontract.CallOpts)
}

// X39 is a free data retrieval call binding the contract method 0x3176cca2.
//
// Solidity: function x39() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X39(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x39")
	return *ret0, err
}

// X39 is a free data retrieval call binding the contract method 0x3176cca2.
//
// Solidity: function x39() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X39() (common.Address, error) {
	return _Dsgcontract.Contract.X39(&_Dsgcontract.CallOpts)
}

// X39 is a free data retrieval call binding the contract method 0x3176cca2.
//
// Solidity: function x39() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X39() (common.Address, error) {
	return _Dsgcontract.Contract.X39(&_Dsgcontract.CallOpts)
}

// X4 is a free data retrieval call binding the contract method 0xcb6cee8a.
//
// Solidity: function x4() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X4(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x4")
	return *ret0, err
}

// X4 is a free data retrieval call binding the contract method 0xcb6cee8a.
//
// Solidity: function x4() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X4() (common.Address, error) {
	return _Dsgcontract.Contract.X4(&_Dsgcontract.CallOpts)
}

// X4 is a free data retrieval call binding the contract method 0xcb6cee8a.
//
// Solidity: function x4() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X4() (common.Address, error) {
	return _Dsgcontract.Contract.X4(&_Dsgcontract.CallOpts)
}

// X40 is a free data retrieval call binding the contract method 0x2ca07afd.
//
// Solidity: function x40() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X40(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x40")
	return *ret0, err
}

// X40 is a free data retrieval call binding the contract method 0x2ca07afd.
//
// Solidity: function x40() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X40() (common.Address, error) {
	return _Dsgcontract.Contract.X40(&_Dsgcontract.CallOpts)
}

// X40 is a free data retrieval call binding the contract method 0x2ca07afd.
//
// Solidity: function x40() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X40() (common.Address, error) {
	return _Dsgcontract.Contract.X40(&_Dsgcontract.CallOpts)
}

// X41 is a free data retrieval call binding the contract method 0xf24ec5c8.
//
// Solidity: function x41() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X41(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x41")
	return *ret0, err
}

// X41 is a free data retrieval call binding the contract method 0xf24ec5c8.
//
// Solidity: function x41() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X41() (common.Address, error) {
	return _Dsgcontract.Contract.X41(&_Dsgcontract.CallOpts)
}

// X41 is a free data retrieval call binding the contract method 0xf24ec5c8.
//
// Solidity: function x41() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X41() (common.Address, error) {
	return _Dsgcontract.Contract.X41(&_Dsgcontract.CallOpts)
}

// X42 is a free data retrieval call binding the contract method 0x943a0169.
//
// Solidity: function x42() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X42(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x42")
	return *ret0, err
}

// X42 is a free data retrieval call binding the contract method 0x943a0169.
//
// Solidity: function x42() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X42() (common.Address, error) {
	return _Dsgcontract.Contract.X42(&_Dsgcontract.CallOpts)
}

// X42 is a free data retrieval call binding the contract method 0x943a0169.
//
// Solidity: function x42() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X42() (common.Address, error) {
	return _Dsgcontract.Contract.X42(&_Dsgcontract.CallOpts)
}

// X43 is a free data retrieval call binding the contract method 0x402cc27e.
//
// Solidity: function x43() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X43(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x43")
	return *ret0, err
}

// X43 is a free data retrieval call binding the contract method 0x402cc27e.
//
// Solidity: function x43() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X43() (common.Address, error) {
	return _Dsgcontract.Contract.X43(&_Dsgcontract.CallOpts)
}

// X43 is a free data retrieval call binding the contract method 0x402cc27e.
//
// Solidity: function x43() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X43() (common.Address, error) {
	return _Dsgcontract.Contract.X43(&_Dsgcontract.CallOpts)
}

// X44 is a free data retrieval call binding the contract method 0x0ae490ad.
//
// Solidity: function x44() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X44(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x44")
	return *ret0, err
}

// X44 is a free data retrieval call binding the contract method 0x0ae490ad.
//
// Solidity: function x44() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X44() (common.Address, error) {
	return _Dsgcontract.Contract.X44(&_Dsgcontract.CallOpts)
}

// X44 is a free data retrieval call binding the contract method 0x0ae490ad.
//
// Solidity: function x44() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X44() (common.Address, error) {
	return _Dsgcontract.Contract.X44(&_Dsgcontract.CallOpts)
}

// X45 is a free data retrieval call binding the contract method 0x1c073086.
//
// Solidity: function x45() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X45(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x45")
	return *ret0, err
}

// X45 is a free data retrieval call binding the contract method 0x1c073086.
//
// Solidity: function x45() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X45() (common.Address, error) {
	return _Dsgcontract.Contract.X45(&_Dsgcontract.CallOpts)
}

// X45 is a free data retrieval call binding the contract method 0x1c073086.
//
// Solidity: function x45() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X45() (common.Address, error) {
	return _Dsgcontract.Contract.X45(&_Dsgcontract.CallOpts)
}

// X46 is a free data retrieval call binding the contract method 0x449136cf.
//
// Solidity: function x46() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X46(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x46")
	return *ret0, err
}

// X46 is a free data retrieval call binding the contract method 0x449136cf.
//
// Solidity: function x46() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X46() (common.Address, error) {
	return _Dsgcontract.Contract.X46(&_Dsgcontract.CallOpts)
}

// X46 is a free data retrieval call binding the contract method 0x449136cf.
//
// Solidity: function x46() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X46() (common.Address, error) {
	return _Dsgcontract.Contract.X46(&_Dsgcontract.CallOpts)
}

// X47 is a free data retrieval call binding the contract method 0x72e1b839.
//
// Solidity: function x47() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X47(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x47")
	return *ret0, err
}

// X47 is a free data retrieval call binding the contract method 0x72e1b839.
//
// Solidity: function x47() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X47() (common.Address, error) {
	return _Dsgcontract.Contract.X47(&_Dsgcontract.CallOpts)
}

// X47 is a free data retrieval call binding the contract method 0x72e1b839.
//
// Solidity: function x47() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X47() (common.Address, error) {
	return _Dsgcontract.Contract.X47(&_Dsgcontract.CallOpts)
}

// X48 is a free data retrieval call binding the contract method 0x7fda3be2.
//
// Solidity: function x48() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X48(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x48")
	return *ret0, err
}

// X48 is a free data retrieval call binding the contract method 0x7fda3be2.
//
// Solidity: function x48() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X48() (common.Address, error) {
	return _Dsgcontract.Contract.X48(&_Dsgcontract.CallOpts)
}

// X48 is a free data retrieval call binding the contract method 0x7fda3be2.
//
// Solidity: function x48() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X48() (common.Address, error) {
	return _Dsgcontract.Contract.X48(&_Dsgcontract.CallOpts)
}

// X49 is a free data retrieval call binding the contract method 0x0418043d.
//
// Solidity: function x49() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X49(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x49")
	return *ret0, err
}

// X49 is a free data retrieval call binding the contract method 0x0418043d.
//
// Solidity: function x49() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X49() (common.Address, error) {
	return _Dsgcontract.Contract.X49(&_Dsgcontract.CallOpts)
}

// X49 is a free data retrieval call binding the contract method 0x0418043d.
//
// Solidity: function x49() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X49() (common.Address, error) {
	return _Dsgcontract.Contract.X49(&_Dsgcontract.CallOpts)
}

// X5 is a free data retrieval call binding the contract method 0xcd145f65.
//
// Solidity: function x5() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X5(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x5")
	return *ret0, err
}

// X5 is a free data retrieval call binding the contract method 0xcd145f65.
//
// Solidity: function x5() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X5() (common.Address, error) {
	return _Dsgcontract.Contract.X5(&_Dsgcontract.CallOpts)
}

// X5 is a free data retrieval call binding the contract method 0xcd145f65.
//
// Solidity: function x5() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X5() (common.Address, error) {
	return _Dsgcontract.Contract.X5(&_Dsgcontract.CallOpts)
}

// X50 is a free data retrieval call binding the contract method 0xb4feb110.
//
// Solidity: function x50() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X50(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x50")
	return *ret0, err
}

// X50 is a free data retrieval call binding the contract method 0xb4feb110.
//
// Solidity: function x50() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X50() (common.Address, error) {
	return _Dsgcontract.Contract.X50(&_Dsgcontract.CallOpts)
}

// X50 is a free data retrieval call binding the contract method 0xb4feb110.
//
// Solidity: function x50() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X50() (common.Address, error) {
	return _Dsgcontract.Contract.X50(&_Dsgcontract.CallOpts)
}

// X51 is a free data retrieval call binding the contract method 0x88ce3426.
//
// Solidity: function x51() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X51(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x51")
	return *ret0, err
}

// X51 is a free data retrieval call binding the contract method 0x88ce3426.
//
// Solidity: function x51() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X51() (common.Address, error) {
	return _Dsgcontract.Contract.X51(&_Dsgcontract.CallOpts)
}

// X51 is a free data retrieval call binding the contract method 0x88ce3426.
//
// Solidity: function x51() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X51() (common.Address, error) {
	return _Dsgcontract.Contract.X51(&_Dsgcontract.CallOpts)
}

// X52 is a free data retrieval call binding the contract method 0xf304e720.
//
// Solidity: function x52() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X52(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x52")
	return *ret0, err
}

// X52 is a free data retrieval call binding the contract method 0xf304e720.
//
// Solidity: function x52() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X52() (common.Address, error) {
	return _Dsgcontract.Contract.X52(&_Dsgcontract.CallOpts)
}

// X52 is a free data retrieval call binding the contract method 0xf304e720.
//
// Solidity: function x52() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X52() (common.Address, error) {
	return _Dsgcontract.Contract.X52(&_Dsgcontract.CallOpts)
}

// X53 is a free data retrieval call binding the contract method 0xbcfbbc46.
//
// Solidity: function x53() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X53(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x53")
	return *ret0, err
}

// X53 is a free data retrieval call binding the contract method 0xbcfbbc46.
//
// Solidity: function x53() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X53() (common.Address, error) {
	return _Dsgcontract.Contract.X53(&_Dsgcontract.CallOpts)
}

// X53 is a free data retrieval call binding the contract method 0xbcfbbc46.
//
// Solidity: function x53() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X53() (common.Address, error) {
	return _Dsgcontract.Contract.X53(&_Dsgcontract.CallOpts)
}

// X54 is a free data retrieval call binding the contract method 0xaf594904.
//
// Solidity: function x54() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X54(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x54")
	return *ret0, err
}

// X54 is a free data retrieval call binding the contract method 0xaf594904.
//
// Solidity: function x54() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X54() (common.Address, error) {
	return _Dsgcontract.Contract.X54(&_Dsgcontract.CallOpts)
}

// X54 is a free data retrieval call binding the contract method 0xaf594904.
//
// Solidity: function x54() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X54() (common.Address, error) {
	return _Dsgcontract.Contract.X54(&_Dsgcontract.CallOpts)
}

// X55 is a free data retrieval call binding the contract method 0x7af6e874.
//
// Solidity: function x55() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X55(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x55")
	return *ret0, err
}

// X55 is a free data retrieval call binding the contract method 0x7af6e874.
//
// Solidity: function x55() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X55() (common.Address, error) {
	return _Dsgcontract.Contract.X55(&_Dsgcontract.CallOpts)
}

// X55 is a free data retrieval call binding the contract method 0x7af6e874.
//
// Solidity: function x55() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X55() (common.Address, error) {
	return _Dsgcontract.Contract.X55(&_Dsgcontract.CallOpts)
}

// X56 is a free data retrieval call binding the contract method 0xfddbd387.
//
// Solidity: function x56() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X56(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x56")
	return *ret0, err
}

// X56 is a free data retrieval call binding the contract method 0xfddbd387.
//
// Solidity: function x56() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X56() (common.Address, error) {
	return _Dsgcontract.Contract.X56(&_Dsgcontract.CallOpts)
}

// X56 is a free data retrieval call binding the contract method 0xfddbd387.
//
// Solidity: function x56() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X56() (common.Address, error) {
	return _Dsgcontract.Contract.X56(&_Dsgcontract.CallOpts)
}

// X57 is a free data retrieval call binding the contract method 0xeb5d92a0.
//
// Solidity: function x57() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X57(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x57")
	return *ret0, err
}

// X57 is a free data retrieval call binding the contract method 0xeb5d92a0.
//
// Solidity: function x57() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X57() (common.Address, error) {
	return _Dsgcontract.Contract.X57(&_Dsgcontract.CallOpts)
}

// X57 is a free data retrieval call binding the contract method 0xeb5d92a0.
//
// Solidity: function x57() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X57() (common.Address, error) {
	return _Dsgcontract.Contract.X57(&_Dsgcontract.CallOpts)
}

// X58 is a free data retrieval call binding the contract method 0xa87f6ab5.
//
// Solidity: function x58() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X58(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x58")
	return *ret0, err
}

// X58 is a free data retrieval call binding the contract method 0xa87f6ab5.
//
// Solidity: function x58() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X58() (common.Address, error) {
	return _Dsgcontract.Contract.X58(&_Dsgcontract.CallOpts)
}

// X58 is a free data retrieval call binding the contract method 0xa87f6ab5.
//
// Solidity: function x58() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X58() (common.Address, error) {
	return _Dsgcontract.Contract.X58(&_Dsgcontract.CallOpts)
}

// X59 is a free data retrieval call binding the contract method 0x13813beb.
//
// Solidity: function x59() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X59(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x59")
	return *ret0, err
}

// X59 is a free data retrieval call binding the contract method 0x13813beb.
//
// Solidity: function x59() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X59() (common.Address, error) {
	return _Dsgcontract.Contract.X59(&_Dsgcontract.CallOpts)
}

// X59 is a free data retrieval call binding the contract method 0x13813beb.
//
// Solidity: function x59() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X59() (common.Address, error) {
	return _Dsgcontract.Contract.X59(&_Dsgcontract.CallOpts)
}

// X6 is a free data retrieval call binding the contract method 0xeaeb89f7.
//
// Solidity: function x6() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X6(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x6")
	return *ret0, err
}

// X6 is a free data retrieval call binding the contract method 0xeaeb89f7.
//
// Solidity: function x6() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X6() (common.Address, error) {
	return _Dsgcontract.Contract.X6(&_Dsgcontract.CallOpts)
}

// X6 is a free data retrieval call binding the contract method 0xeaeb89f7.
//
// Solidity: function x6() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X6() (common.Address, error) {
	return _Dsgcontract.Contract.X6(&_Dsgcontract.CallOpts)
}

// X60 is a free data retrieval call binding the contract method 0x89b3bc3a.
//
// Solidity: function x60() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X60(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x60")
	return *ret0, err
}

// X60 is a free data retrieval call binding the contract method 0x89b3bc3a.
//
// Solidity: function x60() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X60() (common.Address, error) {
	return _Dsgcontract.Contract.X60(&_Dsgcontract.CallOpts)
}

// X60 is a free data retrieval call binding the contract method 0x89b3bc3a.
//
// Solidity: function x60() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X60() (common.Address, error) {
	return _Dsgcontract.Contract.X60(&_Dsgcontract.CallOpts)
}

// X61 is a free data retrieval call binding the contract method 0x588d1150.
//
// Solidity: function x61() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X61(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x61")
	return *ret0, err
}

// X61 is a free data retrieval call binding the contract method 0x588d1150.
//
// Solidity: function x61() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X61() (common.Address, error) {
	return _Dsgcontract.Contract.X61(&_Dsgcontract.CallOpts)
}

// X61 is a free data retrieval call binding the contract method 0x588d1150.
//
// Solidity: function x61() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X61() (common.Address, error) {
	return _Dsgcontract.Contract.X61(&_Dsgcontract.CallOpts)
}

// X62 is a free data retrieval call binding the contract method 0x332bb13b.
//
// Solidity: function x62() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X62(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x62")
	return *ret0, err
}

// X62 is a free data retrieval call binding the contract method 0x332bb13b.
//
// Solidity: function x62() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X62() (common.Address, error) {
	return _Dsgcontract.Contract.X62(&_Dsgcontract.CallOpts)
}

// X62 is a free data retrieval call binding the contract method 0x332bb13b.
//
// Solidity: function x62() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X62() (common.Address, error) {
	return _Dsgcontract.Contract.X62(&_Dsgcontract.CallOpts)
}

// X63 is a free data retrieval call binding the contract method 0x1fc58dae.
//
// Solidity: function x63() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X63(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x63")
	return *ret0, err
}

// X63 is a free data retrieval call binding the contract method 0x1fc58dae.
//
// Solidity: function x63() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X63() (common.Address, error) {
	return _Dsgcontract.Contract.X63(&_Dsgcontract.CallOpts)
}

// X63 is a free data retrieval call binding the contract method 0x1fc58dae.
//
// Solidity: function x63() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X63() (common.Address, error) {
	return _Dsgcontract.Contract.X63(&_Dsgcontract.CallOpts)
}

// X64 is a free data retrieval call binding the contract method 0x9bf3f61a.
//
// Solidity: function x64() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X64(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x64")
	return *ret0, err
}

// X64 is a free data retrieval call binding the contract method 0x9bf3f61a.
//
// Solidity: function x64() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X64() (common.Address, error) {
	return _Dsgcontract.Contract.X64(&_Dsgcontract.CallOpts)
}

// X64 is a free data retrieval call binding the contract method 0x9bf3f61a.
//
// Solidity: function x64() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X64() (common.Address, error) {
	return _Dsgcontract.Contract.X64(&_Dsgcontract.CallOpts)
}

// X65 is a free data retrieval call binding the contract method 0x745be928.
//
// Solidity: function x65() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X65(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x65")
	return *ret0, err
}

// X65 is a free data retrieval call binding the contract method 0x745be928.
//
// Solidity: function x65() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X65() (common.Address, error) {
	return _Dsgcontract.Contract.X65(&_Dsgcontract.CallOpts)
}

// X65 is a free data retrieval call binding the contract method 0x745be928.
//
// Solidity: function x65() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X65() (common.Address, error) {
	return _Dsgcontract.Contract.X65(&_Dsgcontract.CallOpts)
}

// X66 is a free data retrieval call binding the contract method 0xa11be9af.
//
// Solidity: function x66() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X66(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x66")
	return *ret0, err
}

// X66 is a free data retrieval call binding the contract method 0xa11be9af.
//
// Solidity: function x66() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X66() (common.Address, error) {
	return _Dsgcontract.Contract.X66(&_Dsgcontract.CallOpts)
}

// X66 is a free data retrieval call binding the contract method 0xa11be9af.
//
// Solidity: function x66() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X66() (common.Address, error) {
	return _Dsgcontract.Contract.X66(&_Dsgcontract.CallOpts)
}

// X67 is a free data retrieval call binding the contract method 0xbca63eb4.
//
// Solidity: function x67() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X67(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x67")
	return *ret0, err
}

// X67 is a free data retrieval call binding the contract method 0xbca63eb4.
//
// Solidity: function x67() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X67() (common.Address, error) {
	return _Dsgcontract.Contract.X67(&_Dsgcontract.CallOpts)
}

// X67 is a free data retrieval call binding the contract method 0xbca63eb4.
//
// Solidity: function x67() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X67() (common.Address, error) {
	return _Dsgcontract.Contract.X67(&_Dsgcontract.CallOpts)
}

// X68 is a free data retrieval call binding the contract method 0xf721734f.
//
// Solidity: function x68() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X68(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x68")
	return *ret0, err
}

// X68 is a free data retrieval call binding the contract method 0xf721734f.
//
// Solidity: function x68() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X68() (common.Address, error) {
	return _Dsgcontract.Contract.X68(&_Dsgcontract.CallOpts)
}

// X68 is a free data retrieval call binding the contract method 0xf721734f.
//
// Solidity: function x68() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X68() (common.Address, error) {
	return _Dsgcontract.Contract.X68(&_Dsgcontract.CallOpts)
}

// X69 is a free data retrieval call binding the contract method 0xeac637f2.
//
// Solidity: function x69() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X69(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x69")
	return *ret0, err
}

// X69 is a free data retrieval call binding the contract method 0xeac637f2.
//
// Solidity: function x69() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X69() (common.Address, error) {
	return _Dsgcontract.Contract.X69(&_Dsgcontract.CallOpts)
}

// X69 is a free data retrieval call binding the contract method 0xeac637f2.
//
// Solidity: function x69() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X69() (common.Address, error) {
	return _Dsgcontract.Contract.X69(&_Dsgcontract.CallOpts)
}

// X7 is a free data retrieval call binding the contract method 0x282940a7.
//
// Solidity: function x7() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X7(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x7")
	return *ret0, err
}

// X7 is a free data retrieval call binding the contract method 0x282940a7.
//
// Solidity: function x7() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X7() (common.Address, error) {
	return _Dsgcontract.Contract.X7(&_Dsgcontract.CallOpts)
}

// X7 is a free data retrieval call binding the contract method 0x282940a7.
//
// Solidity: function x7() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X7() (common.Address, error) {
	return _Dsgcontract.Contract.X7(&_Dsgcontract.CallOpts)
}

// X70 is a free data retrieval call binding the contract method 0x5d370399.
//
// Solidity: function x70() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X70(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x70")
	return *ret0, err
}

// X70 is a free data retrieval call binding the contract method 0x5d370399.
//
// Solidity: function x70() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X70() (common.Address, error) {
	return _Dsgcontract.Contract.X70(&_Dsgcontract.CallOpts)
}

// X70 is a free data retrieval call binding the contract method 0x5d370399.
//
// Solidity: function x70() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X70() (common.Address, error) {
	return _Dsgcontract.Contract.X70(&_Dsgcontract.CallOpts)
}

// X71 is a free data retrieval call binding the contract method 0x7ede4904.
//
// Solidity: function x71() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X71(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x71")
	return *ret0, err
}

// X71 is a free data retrieval call binding the contract method 0x7ede4904.
//
// Solidity: function x71() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X71() (common.Address, error) {
	return _Dsgcontract.Contract.X71(&_Dsgcontract.CallOpts)
}

// X71 is a free data retrieval call binding the contract method 0x7ede4904.
//
// Solidity: function x71() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X71() (common.Address, error) {
	return _Dsgcontract.Contract.X71(&_Dsgcontract.CallOpts)
}

// X72 is a free data retrieval call binding the contract method 0xac1e43a0.
//
// Solidity: function x72() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X72(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x72")
	return *ret0, err
}

// X72 is a free data retrieval call binding the contract method 0xac1e43a0.
//
// Solidity: function x72() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X72() (common.Address, error) {
	return _Dsgcontract.Contract.X72(&_Dsgcontract.CallOpts)
}

// X72 is a free data retrieval call binding the contract method 0xac1e43a0.
//
// Solidity: function x72() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X72() (common.Address, error) {
	return _Dsgcontract.Contract.X72(&_Dsgcontract.CallOpts)
}

// X73 is a free data retrieval call binding the contract method 0x9b85e451.
//
// Solidity: function x73() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X73(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x73")
	return *ret0, err
}

// X73 is a free data retrieval call binding the contract method 0x9b85e451.
//
// Solidity: function x73() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X73() (common.Address, error) {
	return _Dsgcontract.Contract.X73(&_Dsgcontract.CallOpts)
}

// X73 is a free data retrieval call binding the contract method 0x9b85e451.
//
// Solidity: function x73() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X73() (common.Address, error) {
	return _Dsgcontract.Contract.X73(&_Dsgcontract.CallOpts)
}

// X74 is a free data retrieval call binding the contract method 0xba16cefe.
//
// Solidity: function x74() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X74(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x74")
	return *ret0, err
}

// X74 is a free data retrieval call binding the contract method 0xba16cefe.
//
// Solidity: function x74() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X74() (common.Address, error) {
	return _Dsgcontract.Contract.X74(&_Dsgcontract.CallOpts)
}

// X74 is a free data retrieval call binding the contract method 0xba16cefe.
//
// Solidity: function x74() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X74() (common.Address, error) {
	return _Dsgcontract.Contract.X74(&_Dsgcontract.CallOpts)
}

// X75 is a free data retrieval call binding the contract method 0x175d756b.
//
// Solidity: function x75() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X75(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x75")
	return *ret0, err
}

// X75 is a free data retrieval call binding the contract method 0x175d756b.
//
// Solidity: function x75() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X75() (common.Address, error) {
	return _Dsgcontract.Contract.X75(&_Dsgcontract.CallOpts)
}

// X75 is a free data retrieval call binding the contract method 0x175d756b.
//
// Solidity: function x75() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X75() (common.Address, error) {
	return _Dsgcontract.Contract.X75(&_Dsgcontract.CallOpts)
}

// X76 is a free data retrieval call binding the contract method 0x7f758a15.
//
// Solidity: function x76() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X76(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x76")
	return *ret0, err
}

// X76 is a free data retrieval call binding the contract method 0x7f758a15.
//
// Solidity: function x76() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X76() (common.Address, error) {
	return _Dsgcontract.Contract.X76(&_Dsgcontract.CallOpts)
}

// X76 is a free data retrieval call binding the contract method 0x7f758a15.
//
// Solidity: function x76() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X76() (common.Address, error) {
	return _Dsgcontract.Contract.X76(&_Dsgcontract.CallOpts)
}

// X77 is a free data retrieval call binding the contract method 0x74ec408e.
//
// Solidity: function x77() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X77(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x77")
	return *ret0, err
}

// X77 is a free data retrieval call binding the contract method 0x74ec408e.
//
// Solidity: function x77() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X77() (common.Address, error) {
	return _Dsgcontract.Contract.X77(&_Dsgcontract.CallOpts)
}

// X77 is a free data retrieval call binding the contract method 0x74ec408e.
//
// Solidity: function x77() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X77() (common.Address, error) {
	return _Dsgcontract.Contract.X77(&_Dsgcontract.CallOpts)
}

// X78 is a free data retrieval call binding the contract method 0x1e0a08db.
//
// Solidity: function x78() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X78(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x78")
	return *ret0, err
}

// X78 is a free data retrieval call binding the contract method 0x1e0a08db.
//
// Solidity: function x78() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X78() (common.Address, error) {
	return _Dsgcontract.Contract.X78(&_Dsgcontract.CallOpts)
}

// X78 is a free data retrieval call binding the contract method 0x1e0a08db.
//
// Solidity: function x78() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X78() (common.Address, error) {
	return _Dsgcontract.Contract.X78(&_Dsgcontract.CallOpts)
}

// X79 is a free data retrieval call binding the contract method 0x41f01d3c.
//
// Solidity: function x79() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X79(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x79")
	return *ret0, err
}

// X79 is a free data retrieval call binding the contract method 0x41f01d3c.
//
// Solidity: function x79() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X79() (common.Address, error) {
	return _Dsgcontract.Contract.X79(&_Dsgcontract.CallOpts)
}

// X79 is a free data retrieval call binding the contract method 0x41f01d3c.
//
// Solidity: function x79() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X79() (common.Address, error) {
	return _Dsgcontract.Contract.X79(&_Dsgcontract.CallOpts)
}

// X8 is a free data retrieval call binding the contract method 0xe66e4c3f.
//
// Solidity: function x8() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X8(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x8")
	return *ret0, err
}

// X8 is a free data retrieval call binding the contract method 0xe66e4c3f.
//
// Solidity: function x8() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X8() (common.Address, error) {
	return _Dsgcontract.Contract.X8(&_Dsgcontract.CallOpts)
}

// X8 is a free data retrieval call binding the contract method 0xe66e4c3f.
//
// Solidity: function x8() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X8() (common.Address, error) {
	return _Dsgcontract.Contract.X8(&_Dsgcontract.CallOpts)
}

// X80 is a free data retrieval call binding the contract method 0x2e3f53c9.
//
// Solidity: function x80() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X80(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x80")
	return *ret0, err
}

// X80 is a free data retrieval call binding the contract method 0x2e3f53c9.
//
// Solidity: function x80() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X80() (common.Address, error) {
	return _Dsgcontract.Contract.X80(&_Dsgcontract.CallOpts)
}

// X80 is a free data retrieval call binding the contract method 0x2e3f53c9.
//
// Solidity: function x80() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X80() (common.Address, error) {
	return _Dsgcontract.Contract.X80(&_Dsgcontract.CallOpts)
}

// X81 is a free data retrieval call binding the contract method 0xb02d7daa.
//
// Solidity: function x81() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X81(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x81")
	return *ret0, err
}

// X81 is a free data retrieval call binding the contract method 0xb02d7daa.
//
// Solidity: function x81() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X81() (common.Address, error) {
	return _Dsgcontract.Contract.X81(&_Dsgcontract.CallOpts)
}

// X81 is a free data retrieval call binding the contract method 0xb02d7daa.
//
// Solidity: function x81() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X81() (common.Address, error) {
	return _Dsgcontract.Contract.X81(&_Dsgcontract.CallOpts)
}

// X82 is a free data retrieval call binding the contract method 0x5ae05570.
//
// Solidity: function x82() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X82(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x82")
	return *ret0, err
}

// X82 is a free data retrieval call binding the contract method 0x5ae05570.
//
// Solidity: function x82() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X82() (common.Address, error) {
	return _Dsgcontract.Contract.X82(&_Dsgcontract.CallOpts)
}

// X82 is a free data retrieval call binding the contract method 0x5ae05570.
//
// Solidity: function x82() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X82() (common.Address, error) {
	return _Dsgcontract.Contract.X82(&_Dsgcontract.CallOpts)
}

// X83 is a free data retrieval call binding the contract method 0x60c72ae2.
//
// Solidity: function x83() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X83(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x83")
	return *ret0, err
}

// X83 is a free data retrieval call binding the contract method 0x60c72ae2.
//
// Solidity: function x83() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X83() (common.Address, error) {
	return _Dsgcontract.Contract.X83(&_Dsgcontract.CallOpts)
}

// X83 is a free data retrieval call binding the contract method 0x60c72ae2.
//
// Solidity: function x83() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X83() (common.Address, error) {
	return _Dsgcontract.Contract.X83(&_Dsgcontract.CallOpts)
}

// X84 is a free data retrieval call binding the contract method 0x772eda28.
//
// Solidity: function x84() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X84(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x84")
	return *ret0, err
}

// X84 is a free data retrieval call binding the contract method 0x772eda28.
//
// Solidity: function x84() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X84() (common.Address, error) {
	return _Dsgcontract.Contract.X84(&_Dsgcontract.CallOpts)
}

// X84 is a free data retrieval call binding the contract method 0x772eda28.
//
// Solidity: function x84() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X84() (common.Address, error) {
	return _Dsgcontract.Contract.X84(&_Dsgcontract.CallOpts)
}

// X85 is a free data retrieval call binding the contract method 0x04ba000d.
//
// Solidity: function x85() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X85(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x85")
	return *ret0, err
}

// X85 is a free data retrieval call binding the contract method 0x04ba000d.
//
// Solidity: function x85() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X85() (common.Address, error) {
	return _Dsgcontract.Contract.X85(&_Dsgcontract.CallOpts)
}

// X85 is a free data retrieval call binding the contract method 0x04ba000d.
//
// Solidity: function x85() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X85() (common.Address, error) {
	return _Dsgcontract.Contract.X85(&_Dsgcontract.CallOpts)
}

// X86 is a free data retrieval call binding the contract method 0xac95e388.
//
// Solidity: function x86() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X86(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x86")
	return *ret0, err
}

// X86 is a free data retrieval call binding the contract method 0xac95e388.
//
// Solidity: function x86() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X86() (common.Address, error) {
	return _Dsgcontract.Contract.X86(&_Dsgcontract.CallOpts)
}

// X86 is a free data retrieval call binding the contract method 0xac95e388.
//
// Solidity: function x86() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X86() (common.Address, error) {
	return _Dsgcontract.Contract.X86(&_Dsgcontract.CallOpts)
}

// X87 is a free data retrieval call binding the contract method 0xebc4e3d4.
//
// Solidity: function x87() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X87(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x87")
	return *ret0, err
}

// X87 is a free data retrieval call binding the contract method 0xebc4e3d4.
//
// Solidity: function x87() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X87() (common.Address, error) {
	return _Dsgcontract.Contract.X87(&_Dsgcontract.CallOpts)
}

// X87 is a free data retrieval call binding the contract method 0xebc4e3d4.
//
// Solidity: function x87() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X87() (common.Address, error) {
	return _Dsgcontract.Contract.X87(&_Dsgcontract.CallOpts)
}

// X88 is a free data retrieval call binding the contract method 0x92770c96.
//
// Solidity: function x88() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X88(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x88")
	return *ret0, err
}

// X88 is a free data retrieval call binding the contract method 0x92770c96.
//
// Solidity: function x88() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X88() (common.Address, error) {
	return _Dsgcontract.Contract.X88(&_Dsgcontract.CallOpts)
}

// X88 is a free data retrieval call binding the contract method 0x92770c96.
//
// Solidity: function x88() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X88() (common.Address, error) {
	return _Dsgcontract.Contract.X88(&_Dsgcontract.CallOpts)
}

// X89 is a free data retrieval call binding the contract method 0x457f2f2e.
//
// Solidity: function x89() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X89(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x89")
	return *ret0, err
}

// X89 is a free data retrieval call binding the contract method 0x457f2f2e.
//
// Solidity: function x89() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X89() (common.Address, error) {
	return _Dsgcontract.Contract.X89(&_Dsgcontract.CallOpts)
}

// X89 is a free data retrieval call binding the contract method 0x457f2f2e.
//
// Solidity: function x89() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X89() (common.Address, error) {
	return _Dsgcontract.Contract.X89(&_Dsgcontract.CallOpts)
}

// X9 is a free data retrieval call binding the contract method 0xd6efd153.
//
// Solidity: function x9() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X9(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x9")
	return *ret0, err
}

// X9 is a free data retrieval call binding the contract method 0xd6efd153.
//
// Solidity: function x9() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X9() (common.Address, error) {
	return _Dsgcontract.Contract.X9(&_Dsgcontract.CallOpts)
}

// X9 is a free data retrieval call binding the contract method 0xd6efd153.
//
// Solidity: function x9() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X9() (common.Address, error) {
	return _Dsgcontract.Contract.X9(&_Dsgcontract.CallOpts)
}

// X90 is a free data retrieval call binding the contract method 0x5c356ab5.
//
// Solidity: function x90() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X90(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x90")
	return *ret0, err
}

// X90 is a free data retrieval call binding the contract method 0x5c356ab5.
//
// Solidity: function x90() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X90() (common.Address, error) {
	return _Dsgcontract.Contract.X90(&_Dsgcontract.CallOpts)
}

// X90 is a free data retrieval call binding the contract method 0x5c356ab5.
//
// Solidity: function x90() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X90() (common.Address, error) {
	return _Dsgcontract.Contract.X90(&_Dsgcontract.CallOpts)
}

// X91 is a free data retrieval call binding the contract method 0x9c856be7.
//
// Solidity: function x91() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X91(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x91")
	return *ret0, err
}

// X91 is a free data retrieval call binding the contract method 0x9c856be7.
//
// Solidity: function x91() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X91() (common.Address, error) {
	return _Dsgcontract.Contract.X91(&_Dsgcontract.CallOpts)
}

// X91 is a free data retrieval call binding the contract method 0x9c856be7.
//
// Solidity: function x91() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X91() (common.Address, error) {
	return _Dsgcontract.Contract.X91(&_Dsgcontract.CallOpts)
}

// X92 is a free data retrieval call binding the contract method 0xda9f47ad.
//
// Solidity: function x92() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X92(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x92")
	return *ret0, err
}

// X92 is a free data retrieval call binding the contract method 0xda9f47ad.
//
// Solidity: function x92() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X92() (common.Address, error) {
	return _Dsgcontract.Contract.X92(&_Dsgcontract.CallOpts)
}

// X92 is a free data retrieval call binding the contract method 0xda9f47ad.
//
// Solidity: function x92() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X92() (common.Address, error) {
	return _Dsgcontract.Contract.X92(&_Dsgcontract.CallOpts)
}

// X93 is a free data retrieval call binding the contract method 0x0f99cd32.
//
// Solidity: function x93() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X93(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x93")
	return *ret0, err
}

// X93 is a free data retrieval call binding the contract method 0x0f99cd32.
//
// Solidity: function x93() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X93() (common.Address, error) {
	return _Dsgcontract.Contract.X93(&_Dsgcontract.CallOpts)
}

// X93 is a free data retrieval call binding the contract method 0x0f99cd32.
//
// Solidity: function x93() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X93() (common.Address, error) {
	return _Dsgcontract.Contract.X93(&_Dsgcontract.CallOpts)
}

// X94 is a free data retrieval call binding the contract method 0x5880ab1b.
//
// Solidity: function x94() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X94(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x94")
	return *ret0, err
}

// X94 is a free data retrieval call binding the contract method 0x5880ab1b.
//
// Solidity: function x94() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X94() (common.Address, error) {
	return _Dsgcontract.Contract.X94(&_Dsgcontract.CallOpts)
}

// X94 is a free data retrieval call binding the contract method 0x5880ab1b.
//
// Solidity: function x94() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X94() (common.Address, error) {
	return _Dsgcontract.Contract.X94(&_Dsgcontract.CallOpts)
}

// X95 is a free data retrieval call binding the contract method 0xaa5b743d.
//
// Solidity: function x95() constant returns(address)
func (_Dsgcontract *DsgcontractCaller) X95(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dsgcontract.contract.Call(opts, out, "x95")
	return *ret0, err
}

// X95 is a free data retrieval call binding the contract method 0xaa5b743d.
//
// Solidity: function x95() constant returns(address)
func (_Dsgcontract *DsgcontractSession) X95() (common.Address, error) {
	return _Dsgcontract.Contract.X95(&_Dsgcontract.CallOpts)
}

// X95 is a free data retrieval call binding the contract method 0xaa5b743d.
//
// Solidity: function x95() constant returns(address)
func (_Dsgcontract *DsgcontractCallerSession) X95() (common.Address, error) {
	return _Dsgcontract.Contract.X95(&_Dsgcontract.CallOpts)
}

// Rotate is a paid mutator transaction binding the contract method 0xd992818d.
//
// Solidity: function rotate() returns()
func (_Dsgcontract *DsgcontractTransactor) Rotate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dsgcontract.contract.Transact(opts, "rotate")
}

// Rotate is a paid mutator transaction binding the contract method 0xd992818d.
//
// Solidity: function rotate() returns()
func (_Dsgcontract *DsgcontractSession) Rotate() (*types.Transaction, error) {
	return _Dsgcontract.Contract.Rotate(&_Dsgcontract.TransactOpts)
}

// Rotate is a paid mutator transaction binding the contract method 0xd992818d.
//
// Solidity: function rotate() returns()
func (_Dsgcontract *DsgcontractTransactorSession) Rotate() (*types.Transaction, error) {
	return _Dsgcontract.Contract.Rotate(&_Dsgcontract.TransactOpts)
}
