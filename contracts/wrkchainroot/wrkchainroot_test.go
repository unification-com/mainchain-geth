// Copyright (c) 2019 Unification Foundation
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package wrkchainroot

import (
	"context"
	ethereum "github.com/unification-com/mainchain"
	"github.com/unification-com/mainchain/accounts/abi"
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/accounts/abi/bind/backends"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/common/hexutil"
	"github.com/unification-com/mainchain/contracts/wrkchainroot/contract"
	"github.com/unification-com/mainchain/core"
	"github.com/unification-com/mainchain/crypto"
	"math/big"
	"strings"
	"testing"
	"time"
)

var (
	deployerKey, _         = crypto.HexToECDSA("e7b1f342c77881ff425d09b83892cc0683435878cf1d3ff2b6e015ada9378ddf")
	deployerAddr           = crypto.PubkeyToAddress(deployerKey.PublicKey)
	wrkchainOwnerKey, _    = crypto.HexToECDSA("d6acd86a8589afb9f3658626eda542ace08d17432219a1eaf2f7dfeac32adf66")
	wrkchainOwnerAddr      = crypto.PubkeyToAddress(wrkchainOwnerKey.PublicKey)
	accKey1, _             = crypto.HexToECDSA("d7c15922c5371f9536f8b1cef750d5451b01f57b2865278fef6bd65c9a17a3a3")
	accAddr1               = crypto.PubkeyToAddress(accKey1.PublicKey)
	accKey2, _             = crypto.HexToECDSA("e6d90f41cebb11eb6f5f5d7b71a49f3f4309d78ffb3e66eaa6c7af94c7d99d41")
	accAddr2               = crypto.PubkeyToAddress(accKey2.PublicKey)
	accKey3, _             = crypto.HexToECDSA("dc23751c35323b70ad70ec2d08d9be59b74f8491baccc87123a0c0ffdfa24936")
	accAddr3               = crypto.PubkeyToAddress(accKey3.PublicKey)
	wrkchainId             = "MyWRKChain12345" //If modified, also need to regenerate registerData and recordData below
	wrkchainIdBytes32      = common.BytesToHash([]byte(wrkchainId))
	wrkchainAuthAddresses1 = []common.Address{wrkchainOwnerAddr, accAddr1}
	wrkchainGenesisHash1   = [32]byte{'x'}
	genesisBalance         = new(big.Int).SetUint64(1000000000000000000)
)

func packRegData(chainId common.Hash, authAddresses []common.Address, genesisHash [32]byte) ([]byte, error) {
	wrkchainRootAbi, err :=  abi.JSON(strings.NewReader(contract.WRKChainRootABI))
	if err != nil {
		return nil, err
	}
	packed, err := wrkchainRootAbi.Pack("registerWrkChain", chainId, authAddresses, genesisHash)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

func packRecordData(chainId common.Hash,
	height *big.Int,
	hash [32]byte,
	parentHash [32]byte,
	receiptRoot [32]byte,
	txRoot [32]byte,
	stateRoot [32]byte) ([]byte, error) {

	wrkchainRootAbi, err :=  abi.JSON(strings.NewReader(contract.WRKChainRootABI))
	if err != nil {
		return nil, err
	}
	packed, err := wrkchainRootAbi.Pack("recordHeader", chainId, height, hash, parentHash, receiptRoot, txRoot, stateRoot)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

func packAddAuthAddresses(chainId common.Hash, authAddresses []common.Address) ([]byte, error) {
	wrkchainRootAbi, err :=  abi.JSON(strings.NewReader(contract.WRKChainRootABI))
	if err != nil {
		return nil, err
	}
	packed, err := wrkchainRootAbi.Pack("addAuthAddresses", chainId, authAddresses)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

func packRemoveAuthAddresses(chainId common.Hash, authAddresses []common.Address) ([]byte, error) {
	wrkchainRootAbi, err :=  abi.JSON(strings.NewReader(contract.WRKChainRootABI))
	if err != nil {
		return nil, err
	}
	packed, err := wrkchainRootAbi.Pack("removeAuthAddresses", chainId, authAddresses)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

func generateWrkchainRootCallMsg(fromAddr common.Address, wrkchainRootAddress common.Address, data []byte) ethereum.CallMsg {

	msg := ethereum.CallMsg{}
	msg.From = fromAddr
	msg.To = &wrkchainRootAddress
	msg.Gas = 10000000
	msg.GasPrice = big.NewInt(25000)
	msg.Value = big.NewInt(0)
	msg.IsWrkchainBeaconMessage = true
	msg.Data = data

	return msg
}

func TestWrkchainRootDeploy(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOpts := bind.NewKeyedTransactor(deployerKey)

	wrkchainRootAddress, _, err := DeployWrkchainRoot(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't deploy wrkchainroot: %v", err)
	}
	contractBackend.Commit()
	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	code, err := contractBackend.CodeAt(ctx, wrkchainRootAddress, nil)
	if err != nil {
		t.Fatalf("can't get wrkchainroot code: %v", err)
	}

	t.Log("contract code", hexutil.Encode(code))
	t.Log("contract address", wrkchainRootAddress.Hex())
}

func TestRegisterWrkchain(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)

	wrkchainRootAddress, _, err  := DeployWrkchainRoot(transactOptsDeployer, contractBackend)
	contractBackend.Commit()

	t.Log("contract address", wrkchainRootAddress.Hex())

	msgData, err := packRegData(wrkchainIdBytes32, wrkchainAuthAddresses1, wrkchainGenesisHash1)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}

	msg := generateWrkchainRootCallMsg(wrkchainOwnerAddr, wrkchainRootAddress, msgData)

	// run the contract call, and process the result
	res, err := contractBackend.PendingCallContract(context.Background(), msg)
	if len(res) > 0 {
		t.Log("res string", string(res))
		t.Fatalf("Failed to register")
	}

	// wrkchain owner registers a WRKChain for real
	transactOptsWrkchainOwner.Value =  big.NewInt(0)
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	t.Log("wrkchainIdBytes32", wrkchainIdBytes32)
	t.Log("wrkchainAuthAddresses1", wrkchainAuthAddresses1)
	t.Log("wrkchainGenesisHash1", wrkchainGenesisHash1)
	tx, err := wrkchainOwner.RegisterWrkChain(wrkchainIdBytes32, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	t.Log("RegisterWrkChain Tx", tx)
	if err != nil {
		t.Fatalf("can't register wrkchainroot: %v", err)
	}
	contractBackend.Commit()

	t.Log("RegisterWrkChain Tx", tx)
	t.Log("RegisterWrkChain Tx.Data()", common.Bytes2Hex(tx.Data()))

	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())

	t.Log("receipt", receipt)

	if receipt.Status != 1 {
		t.Fatalf("Receipt status == false")
	}

	regged, err := wrkchainOwner.ChainExists(common.BytesToHash([]byte(wrkchainId)))

	if err != nil {
		t.Fatalf("ChainExists error: %v", err)
	}

	t.Log("regged", regged)

	if !regged {
		t.Fatalf("ChainExists should be true. Got %v", regged)
	}
}

func TestCanRegisterWRKChainOnlyOnce(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)

	wrkchainRootAddress, _, err  := DeployWrkchainRoot(transactOptsDeployer, contractBackend)
	contractBackend.Commit()
	msgData, err := packRegData(wrkchainIdBytes32, wrkchainAuthAddresses1, wrkchainGenesisHash1)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}
	msg := generateWrkchainRootCallMsg(wrkchainOwnerAddr, wrkchainRootAddress, msgData)

	// run the contract call, and process the result
	res, err := contractBackend.PendingCallContract(context.Background(), msg)
	if len(res) > 0 {
		t.Log("res string", string(res))
		t.Fatalf("Failed to register")
	}

	// wrkchain owner registers a WRKChain for real
	transactOptsWrkchainOwner.Value =  big.NewInt(0)
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)
	tx, err := wrkchainOwner.RegisterWrkChain(wrkchainIdBytes32, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	if err != nil {
		t.Fatalf("can't register wrkchainroot: %v", err)
	}
	contractBackend.Commit()
	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())
	if receipt.Status != 1 {
		t.Fatalf("Receipt status == false")
	}

    // Attempt to register same Chain ID a second time. Both should fail
    //check pending call first
	res, err = contractBackend.PendingCallContract(context.Background(), msg)
	if len(res) == 0 {
		t.Fatalf("should not be able to register the same chain id more than once")
	}

	t.Log("res", string(res))

	// try sending a Tx
	tx, err = wrkchainOwner.RegisterWrkChain(wrkchainIdBytes32, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	if err == nil {
		t.Fatalf("tx sucess but should not be able to register the same chain id more than once")
	}

	t.Log("err", err)

}

func TestRegisterWrkchainNullChainId(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)

	wrkchainRootAddress, _, err  := DeployWrkchainRoot(transactOptsDeployer, contractBackend)
	contractBackend.Commit()

	msgData, err := packRegData(common.BytesToHash([]byte("")), wrkchainAuthAddresses1, wrkchainGenesisHash1)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}
	msg := generateWrkchainRootCallMsg(wrkchainOwnerAddr, wrkchainRootAddress, msgData)
	// run the contract call, and process the result
	res, err := contractBackend.PendingCallContract(context.Background(), msg)
	t.Log("res string", string(res))

	//should fail
	if len(res) == 0 {
		t.Fatalf("Should not allow null ChainID registration")
	}

	// wrkchain owner registers a WRKChain for real
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	_, err = wrkchainOwner.RegisterWrkChain(common.BytesToHash([]byte("")), wrkchainAuthAddresses1, wrkchainGenesisHash1)

	if err == nil {
		t.Fatalf("registering WRKChain with null chainid should faile")
	}

	t.Log("err:", err)
}

func TestCanOnlyAuthTenAddresses(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)

	wrkchainRootAddress, _, err  := DeployWrkchainRoot(transactOptsDeployer, contractBackend)
	contractBackend.Commit()

	authAddr := []common.Address{
		accAddr1,
		accAddr2,
		accAddr3,
		common.HexToAddress("0xa13A71dfe5cD57F9b67ec6A54AD2Ae7537D7fc3b"),
		common.HexToAddress("0xb47FD09F1d379Ce2c9BFF59D668cf0B7b994a2B7"),
		common.HexToAddress("0xcA29F1275470DE81DE6E1Bb53a55228Da676E752"),
		common.HexToAddress("0xd71aD3263666e03004479B30CA840a26eaC5b763"),
		common.HexToAddress("0xd333016e7fcf9eDd4F1468493c10cb463131CAf4"),
		common.HexToAddress("0xe76D1b02C95C11DB13f733C07Fce5e173b322B2f"),
		common.HexToAddress("0xf30F951b0426f8Bf37ac71967407081358DF7a7B"),
		common.HexToAddress("0xeeee680c1D5337c364ab6eaf7828ab4A36522cfb"),
	}
	msgData, err := packRegData(wrkchainIdBytes32, authAddr, wrkchainGenesisHash1)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}

	msg := generateWrkchainRootCallMsg(wrkchainOwnerAddr, wrkchainRootAddress, msgData)
	// run the contract call, and process the result
	res, err := contractBackend.PendingCallContract(context.Background(), msg)
	t.Log("res string", string(res))

	//should fail
	if len(res) == 0 {
		t.Fatalf("Should not allow registration with more than 10 auth addresses")
	}

	msgData, err = packAddAuthAddresses(wrkchainIdBytes32, authAddr)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}

	msg = generateWrkchainRootCallMsg(wrkchainOwnerAddr, wrkchainRootAddress, msgData)
	// run the contract call, and process the result
	res, err = contractBackend.PendingCallContract(context.Background(), msg)
	t.Log("res string", string(res))

	//should fail
	if len(res) == 0 {
		t.Fatalf("Should not allow registration with more than 10 auth addresses")
	}
}

func TestWritingHashes(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)
	transactOptsWrkchainOwner.GasLimit = 2400000
	transactOptsWrkchainOwner.GasPrice = big.NewInt(250000)

	wrkchainRootAddress, _, _  := DeployWrkchainRoot(transactOptsDeployer, contractBackend)

	contractBackend.Commit()

	// wrkchain owner registers a WRKChain
	transactOptsWrkchainOwner.Value =  big.NewInt(0)
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	_, _ = wrkchainOwner.RegisterWrkChain(wrkchainIdBytes32, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	contractBackend.Commit()

	msgData, err := packRecordData(wrkchainIdBytes32,
		new(big.Int).SetUint64(1),
		[32]byte{'a'},
		[32]byte{'b'},
		[32]byte{'c'},
		[32]byte{'d'},
		[32]byte{'e'},
	)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}
	// record some hashes
	msg := generateWrkchainRootCallMsg(wrkchainOwnerAddr, wrkchainRootAddress, msgData)

	// dry run the contract call, and process the VM result
	res, _ := contractBackend.PendingCallContract(context.Background(), msg)
	if len(res) != 0 {
		t.Log("res", string(res))
		t.Fatalf("Should have recorded hashes")
	}

	// reset Value to 0 as we're recording
	transactOptsWrkchainOwner.Value = big.NewInt(0)
	wrkchainOwner, _ = NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	tx, err := wrkchainOwner.RecordHeader(
		wrkchainIdBytes32,
		new(big.Int).SetUint64(1),
		[32]byte{'a'},
		[32]byte{'b'},
		[32]byte{'c'},
		[32]byte{'d'},
		[32]byte{'e'},
	)

	if err != nil {
		t.Fatalf("can't record header hashes: %v", err)
	}

	contractBackend.Commit()

	t.Log("RecordHeader Tx", tx)
	t.Log("RecordHeader Tx.Data()", common.Bytes2Hex(tx.Data()))

	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())

	t.Log("receipt", receipt)

	if receipt.Status != 1 {
		t.Fatalf("Receipt status == false")
	}

	lastData, err := wrkchainOwner.GetLastBlock(wrkchainIdBytes32)

	if err != nil {
		t.Log(err)
	}

	t.Log("lastData", lastData)
}

func TestOnlyAuthCanWriteHashes(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
		accAddr2: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)

	transactOptsAcc2 := bind.NewKeyedTransactor(accKey2)
	transactOptsAcc2.GasLimit = 4700000

	wrkchainRootAddress, _, _  := DeployWrkchainRoot(transactOptsDeployer, contractBackend)

	contractBackend.Commit()

	// wrkchain owner registers a WRKChain
	transactOptsWrkchainOwner.Value =  big.NewInt(0)
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	_, _ = wrkchainOwner.RegisterWrkChain(wrkchainIdBytes32, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	contractBackend.Commit()

	msgData, err := packRecordData(wrkchainIdBytes32,
		new(big.Int).SetUint64(1),
		[32]byte{'a'},
		[32]byte{'b'},
		[32]byte{'c'},
		[32]byte{'d'},
		[32]byte{'e'},
	)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}
	// accAddr2 should not be able to record hashes at this point
	msg := generateWrkchainRootCallMsg(accAddr2, wrkchainRootAddress, msgData)

	// dry run the contract call, and process the VM result
	res, _ := contractBackend.PendingCallContract(context.Background(), msg)

	if len(res) == 0 {
		t.Fatalf("accAddr2 should not be able to record hashes")
	}

	t.Log("msg.res", string(res))

	acc2, _ := NewWrkchainRoot(transactOptsAcc2, wrkchainRootAddress, contractBackend)

	tx, err := acc2.RecordHeader(
		wrkchainIdBytes32,
		new(big.Int).SetUint64(1),
		[32]byte{'a'},
		[32]byte{'b'},
		[32]byte{'c'},
		[32]byte{'d'},
		[32]byte{'e'},
	)

	if err != nil {
		t.Log("tx", tx)
		t.Fatalf("error running acc2.RecordHeader: %v", err)
	}

	contractBackend.Commit()

	t.Log("RecordHeader Tx", tx)

	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())

	t.Log("receipt", receipt)

	if receipt.Status != 0 {
		t.Fatalf("Receipt status == true. accAddr2 should not be able to record hashes")
	}
}

func TestNewAuthCanWrite(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
		accAddr2: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)
	transactOptsAcc2 := bind.NewKeyedTransactor(accKey2)

	wrkchainRootAddress, _, _  := DeployWrkchainRoot(transactOptsDeployer, contractBackend)

	contractBackend.Commit()

	// wrkchain owner registers a WRKChain
	transactOptsWrkchainOwner.Value =  big.NewInt(0)
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	_, _ = wrkchainOwner.RegisterWrkChain(wrkchainIdBytes32, wrkchainAuthAddresses1, wrkchainGenesisHash1)

	contractBackend.Commit()

	//WRKChain owner adds new Auth
	transactOptsWrkchainOwner.Value = new(big.Int).SetUint64(0)
	wrkchainOwner, _ = NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)
	_, err := wrkchainOwner.AddAuthAddresses(wrkchainIdBytes32, []common.Address{accAddr2})

	if err != nil {
		t.Fatalf("Failed to authorise new address: %v", err)
	}
	contractBackend.Commit()
	msgData, err := packRecordData(wrkchainIdBytes32,
		new(big.Int).SetUint64(1),
		[32]byte{'a'},
		[32]byte{'b'},
		[32]byte{'c'},
		[32]byte{'d'},
		[32]byte{'e'},
	)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}
	// accAddr2 should be able to record hashes now
	msg := generateWrkchainRootCallMsg(accAddr2, wrkchainRootAddress, msgData)

	// dry run the contract call, and process the VM result
	res, _ := contractBackend.PendingCallContract(context.Background(), msg)

	if len(res) != 0 {
		t.Log("msg.res", string(res))
		t.Fatalf("accAddr2 should be able to record hashes now")
	}

	acc2, _ := NewWrkchainRoot(transactOptsAcc2, wrkchainRootAddress, contractBackend)

	tx, err := acc2.RecordHeader(
		wrkchainIdBytes32,
		new(big.Int).SetUint64(1),
		[32]byte{'a'},
		[32]byte{'b'},
		[32]byte{'c'},
		[32]byte{'d'},
		[32]byte{'e'},
	)

	if err != nil {
		t.Fatalf("error running acc2.RecordHeader: %v", err)
	}

	contractBackend.Commit()

	t.Log("RecordHeader Tx", tx)

	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())

	t.Log("receipt", receipt)

	if receipt.Status != 1 {
		t.Fatalf("Receipt status == false. accAddr2 should be able to record hashes now")
	}
}

func TestRecordWithoutRegisterFail (t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)

	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)
	transactOptsWrkchainOwner.GasLimit = 240000
	transactOptsWrkchainOwner.GasPrice = big.NewInt(2500)
	transactOptsWrkchainOwner.Value = big.NewInt(0)

	wrkchainRootAddress, _, _  := DeployWrkchainRoot(transactOptsDeployer, contractBackend)

	contractBackend.Commit()

	msgData, err := packRecordData(wrkchainIdBytes32,
		new(big.Int).SetUint64(1),
		[32]byte{'a'},
		[32]byte{'b'},
		[32]byte{'c'},
		[32]byte{'d'},
		[32]byte{'e'},
	)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}

	msg := generateWrkchainRootCallMsg(wrkchainOwnerAddr, wrkchainRootAddress, msgData)

	// run the contract call, and process the result
	res, err := contractBackend.PendingCallContract(context.Background(), msg)

	t.Log("msg", msg)
	t.Log("res", res)
	t.Log("res string", string(res))
	t.Log("err", err)

	if len(res) == 0 {
		t.Fatalf("should have failed and reverted")
	}

	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	tx, err := wrkchainOwner.RecordHeader(
		wrkchainIdBytes32,
		new(big.Int).SetUint64(1),
		[32]byte{'a'},
		[32]byte{'b'},
		[32]byte{'c'},
		[32]byte{'d'},
		[32]byte{'e'},
	)

	if err != nil {
		t.Fatalf("error running wrkchainOwner.RecordHeader: %v", err)
	}

	contractBackend.Commit()

	t.Log("RecordHeader Tx", tx)

	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())

	t.Log("receipt", receipt)

	if receipt.Status != 0 {
		t.Fatalf("Receipt status == true. Should have failed")
	}
}

func TestChainExists(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		wrkchainOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsWrkchainOwner := bind.NewKeyedTransactor(wrkchainOwnerKey)

	wrkchainRootAddress, _, _  := DeployWrkchainRoot(transactOptsDeployer, contractBackend)
	contractBackend.Commit()

	transactOptsWrkchainOwner.Value =  big.NewInt(0)
	wrkchainOwner, _ := NewWrkchainRoot(transactOptsWrkchainOwner, wrkchainRootAddress, contractBackend)

	_, _ = wrkchainOwner.RegisterWrkChain(wrkchainIdBytes32, wrkchainAuthAddresses1, wrkchainGenesisHash1)
	contractBackend.Commit()

	regged, err := wrkchainOwner.ChainExists(common.BytesToHash([]byte(wrkchainId)))

	if err != nil {
		t.Fatalf("ChainExists error: %v", err)
	}

	if !regged {
		t.Fatalf("ChainExists for %v should be true. Got %v", wrkchainId, regged)
	}

	chainIdNotExist := "notexist"
	regged, _ = wrkchainOwner.ChainExists(common.BytesToHash([]byte(chainIdNotExist)))

	if regged {
		t.Fatalf("ChainExists for %v should be false. Got %v", chainIdNotExist, regged)
	}
}
