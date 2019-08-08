package beacon

import (
	"context"
	ethereum "github.com/unification-com/mainchain"
	"github.com/unification-com/mainchain/accounts/abi"
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/accounts/abi/bind/backends"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/common/hexutil"
	"github.com/unification-com/mainchain/contracts/beacon/contract"
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
	beaconOwnerKey, _    = crypto.HexToECDSA("d6acd86a8589afb9f3658626eda542ace08d17432219a1eaf2f7dfeac32adf66")
	beaconOwnerAddr      = crypto.PubkeyToAddress(beaconOwnerKey.PublicKey)
	accKey1, _             = crypto.HexToECDSA("d7c15922c5371f9536f8b1cef750d5451b01f57b2865278fef6bd65c9a17a3a3")
	accAddr1               = crypto.PubkeyToAddress(accKey1.PublicKey)
	accKey2, _             = crypto.HexToECDSA("e6d90f41cebb11eb6f5f5d7b71a49f3f4309d78ffb3e66eaa6c7af94c7d99d41")
	accAddr2               = crypto.PubkeyToAddress(accKey2.PublicKey)
	accKey3, _             = crypto.HexToECDSA("dc23751c35323b70ad70ec2d08d9be59b74f8491baccc87123a0c0ffdfa24936")
	accAddr3               = crypto.PubkeyToAddress(accKey3.PublicKey)
	beaconId             = "MyBeacon12345"
	beaconIdBytes32      = common.BytesToHash([]byte(beaconId))
	beaconAuthAddresses1 = []common.Address{beaconOwnerAddr, accAddr1}
	genesisBalance         = new(big.Int).SetUint64(1000000000000000000)
)

func packRegData(beaconId common.Hash, authAddresses []common.Address) ([]byte, error) {
	beaconAbi, err :=  abi.JSON(strings.NewReader(contract.BeaconABI))
	if err != nil {
		return nil, err
	}
	packed, err := beaconAbi.Pack("registerBeacon", beaconId, authAddresses)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

func packRecordData(beaconId common.Hash, hash [32]byte) ([]byte, error) {
	beaconAbi, err :=  abi.JSON(strings.NewReader(contract.BeaconABI))
	if err != nil {
		return nil, err
	}
	packed, err := beaconAbi.Pack("recordHash", beaconId, hash)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

func packAddAuthAddresses(chainId common.Hash, authAddresses []common.Address) ([]byte, error) {
	beaconAbi, err :=  abi.JSON(strings.NewReader(contract.BeaconABI))
	if err != nil {
		return nil, err
	}
	packed, err := beaconAbi.Pack("addAuthAddresses", chainId, authAddresses)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

func packRemoveAuthAddresses(chainId common.Hash, authAddresses []common.Address) ([]byte, error) {
	beaconAbi, err :=  abi.JSON(strings.NewReader(contract.BeaconABI))
	if err != nil {
		return nil, err
	}
	packed, err := beaconAbi.Pack("removeAuthAddresses", chainId, authAddresses)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

func generateBeaconCallMsg(fromAddr common.Address, beaconAddress common.Address, data []byte) ethereum.CallMsg {

	msg := ethereum.CallMsg{}
	msg.From = fromAddr
	msg.To = &beaconAddress
	msg.Gas = 10000000
	msg.GasPrice = big.NewInt(25000)
	msg.Value = big.NewInt(0)
	msg.IsWrkchainBeaconMessage = true
	msg.Data = data

	return msg
}

func TestBeaconDeploy(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)
	transactOpts := bind.NewKeyedTransactor(deployerKey)

	beaconAddress, _, err := DeployBeacon(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't deploy beacon: %v", err)
	}
	contractBackend.Commit()
	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	code, err := contractBackend.CodeAt(ctx, beaconAddress, nil)
	if err != nil {
		t.Fatalf("can't get beacon code: %v", err)
	}

	t.Log("contract code", hexutil.Encode(code))
	t.Log("contract address", beaconAddress.Hex())
}

func TestRegisterBeacon(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		beaconOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsBeaconOwner := bind.NewKeyedTransactor(beaconOwnerKey)

	beaconAddress, _, err  := DeployBeacon(transactOptsDeployer, contractBackend)
	contractBackend.Commit()

	t.Log("contract address", beaconAddress.Hex())

	msgData, err := packRegData(beaconIdBytes32, beaconAuthAddresses1)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}

	msg := generateBeaconCallMsg(beaconOwnerAddr, beaconAddress, msgData)

	// run the contract call, and process the result
	res, err := contractBackend.PendingCallContract(context.Background(), msg)
	if len(res) > 0 {
		t.Log("res string", string(res))
		t.Fatalf("Failed to register")
	}

	// beaconOwnerAddr registers a Beacon for real
	transactOptsBeaconOwner.Value =  big.NewInt(0)
	beaconOwner, _ := NewBeacon(transactOptsBeaconOwner, beaconAddress, contractBackend)

	t.Log("beaconIdBytes32", beaconIdBytes32)
	t.Log("beaconAuthAddresses1", beaconAuthAddresses1)
	tx, err := beaconOwner.RegisterBeacon(beaconIdBytes32, beaconAuthAddresses1)

	t.Log("RegisterBeacon Tx", tx)
	if err != nil {
		t.Fatalf("can't register beacon: %v", err)
	}
	contractBackend.Commit()

	t.Log("RegisterBeacon Tx", tx)
	t.Log("RegisterBeacon Tx.Data()", common.Bytes2Hex(tx.Data()))

	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())

	t.Log("receipt", receipt)

	if receipt.Status != 1 {
		t.Fatalf("Receipt status == false")
	}

	regged, err := beaconOwner.BeaconExists(common.BytesToHash([]byte(beaconId)))

	if err != nil {
		t.Fatalf("BeaconExists error: %v", err)
	}

	t.Log("regged", regged)

	if !regged {
		t.Fatalf("BeaconExists should be true. Got %v", regged)
	}
}


func TestCanRegisterBeaconOnlyOnce(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		beaconOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsBeaconOwner := bind.NewKeyedTransactor(beaconOwnerKey)

	beaconAddress, _, err  := DeployBeacon(transactOptsDeployer, contractBackend)
	contractBackend.Commit()
	msgData, err := packRegData(beaconIdBytes32, beaconAuthAddresses1)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}
	msg := generateBeaconCallMsg(beaconOwnerAddr, beaconAddress, msgData)

	// run the contract call, and process the result
	res, err := contractBackend.PendingCallContract(context.Background(), msg)
	if len(res) > 0 {
		t.Log("res string", string(res))
		t.Fatalf("Failed to register")
	}

	// beacon owner registers a Beacon for real
	transactOptsBeaconOwner.Value =  big.NewInt(0)
	beaconOwner, _ := NewBeacon(transactOptsBeaconOwner, beaconAddress, contractBackend)
	tx, err := beaconOwner.RegisterBeacon(beaconIdBytes32, beaconAuthAddresses1)

	if err != nil {
		t.Fatalf("can't register beacon: %v", err)
	}
	contractBackend.Commit()
	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())
	if receipt.Status != 1 {
		t.Fatalf("Receipt status == false")
	}

	// Attempt to register same Beacon ID a second time. Both should fail
	//check pending call first
	res, err = contractBackend.PendingCallContract(context.Background(), msg)
	if len(res) == 0 {
		t.Fatalf("should not be able to register the same chain id more than once")
	}

	t.Log("res", string(res))

	// try sending a Tx
	tx, err = beaconOwner.RegisterBeacon(beaconIdBytes32, beaconAuthAddresses1)

	if err == nil {
		t.Fatalf("tx sucess but should not be able to register the same chain id more than once")
	}

	t.Log("err", err)

}

func TestRegisterBeaconNullBeaconId(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		beaconOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsBeaconOwner := bind.NewKeyedTransactor(beaconOwnerKey)

	beaconAddress, _, err  := DeployBeacon(transactOptsDeployer, contractBackend)
	contractBackend.Commit()

	msgData, err := packRegData(common.BytesToHash([]byte("")), beaconAuthAddresses1)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}
	msg := generateBeaconCallMsg(beaconOwnerAddr, beaconAddress, msgData)
	// run the contract call, and process the result
	res, err := contractBackend.PendingCallContract(context.Background(), msg)
	t.Log("res string", string(res))

	//should fail
	if len(res) == 0 {
		t.Fatalf("Should not allow null BeaconID registration")
	}

	// beacon owner registers a Beacon for real
	beaconOwner, _ := NewBeacon(transactOptsBeaconOwner, beaconAddress, contractBackend)

	_, err = beaconOwner.RegisterBeacon(common.BytesToHash([]byte("")), beaconAuthAddresses1)

	if err == nil {
		t.Fatalf("registering Beacon with null chainid should fail")
	}

	t.Log("err:", err)
}

func TestCanOnlyAuthTenAddresses(t *testing.T) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		beaconOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)

	beaconAddress, _, err  := DeployBeacon(transactOptsDeployer, contractBackend)
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
	msgData, err := packRegData(beaconIdBytes32, authAddr)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}

	msg := generateBeaconCallMsg(beaconOwnerAddr, beaconAddress, msgData)
	// run the contract call, and process the result
	res, err := contractBackend.PendingCallContract(context.Background(), msg)
	t.Log("res string", string(res))

	//should fail
	if len(res) == 0 {
		t.Fatalf("Should not allow registration with more than 10 auth addresses")
	}

	msgData, err = packAddAuthAddresses(beaconIdBytes32, authAddr)
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}

	msg = generateBeaconCallMsg(beaconOwnerAddr, beaconAddress, msgData)
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
		beaconOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsBeaconOwner := bind.NewKeyedTransactor(beaconOwnerKey)
	transactOptsBeaconOwner.GasLimit = 2400000
	transactOptsBeaconOwner.GasPrice = big.NewInt(250000)

	beaconAddress, _, _  := DeployBeacon(transactOptsDeployer, contractBackend)

	contractBackend.Commit()

	// beacon owner registers a Beacon
	transactOptsBeaconOwner.Value =  big.NewInt(0)
	beaconOwner, _ := NewBeacon(transactOptsBeaconOwner, beaconAddress, contractBackend)

	_, _ = beaconOwner.RegisterBeacon(beaconIdBytes32, beaconAuthAddresses1)

	contractBackend.Commit()

	msgData, err := packRecordData(beaconIdBytes32, [32]byte{'a'})
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}
	// record some hashes
	msg := generateBeaconCallMsg(beaconOwnerAddr, beaconAddress, msgData)

	// dry run the contract call, and process the VM result
	res, _ := contractBackend.PendingCallContract(context.Background(), msg)
	if len(res) != 0 {
		t.Log("res", string(res))
		t.Fatalf("Should have recorded hashes")
	}

	// reset Value to 0 as we're recording
	transactOptsBeaconOwner.Value = big.NewInt(0)
	beaconOwner, _ = NewBeacon(transactOptsBeaconOwner, beaconAddress, contractBackend)

	tx, err := beaconOwner.RecordHash(beaconIdBytes32, [32]byte{'a'})

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

	lastData, err := beaconOwner.GetLastHash(beaconIdBytes32)

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
		beaconOwnerAddr: {
			Balance: genesisBalance,
		},
		accAddr2: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsBeaconOwner := bind.NewKeyedTransactor(beaconOwnerKey)

	transactOptsAcc2 := bind.NewKeyedTransactor(accKey2)
	transactOptsAcc2.GasLimit = 4700000

	beaconAddress, _, _  := DeployBeacon(transactOptsDeployer, contractBackend)

	contractBackend.Commit()

	// beacon owner registers a Beacon
	transactOptsBeaconOwner.Value =  big.NewInt(0)
	beaconOwner, _ := NewBeacon(transactOptsBeaconOwner, beaconAddress, contractBackend)

	_, _ = beaconOwner.RegisterBeacon(beaconIdBytes32, beaconAuthAddresses1)

	contractBackend.Commit()

	msgData, err := packRecordData(beaconIdBytes32, [32]byte{'a'})
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}
	// accAddr2 should not be able to record hashes at this point
	msg := generateBeaconCallMsg(accAddr2, beaconAddress, msgData)

	// dry run the contract call, and process the VM result
	res, _ := contractBackend.PendingCallContract(context.Background(), msg)

	if len(res) == 0 {
		t.Fatalf("accAddr2 should not be able to record hashes")
	}

	t.Log("msg.res", string(res))

	acc2, _ := NewBeacon(transactOptsAcc2, beaconAddress, contractBackend)

	tx, err := acc2.RecordHash(beaconIdBytes32, [32]byte{'a'})

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
		beaconOwnerAddr: {
			Balance: genesisBalance,
		},
		accAddr2: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsBeaconOwner := bind.NewKeyedTransactor(beaconOwnerKey)
	transactOptsAcc2 := bind.NewKeyedTransactor(accKey2)

	beaconAddress, _, _  := DeployBeacon(transactOptsDeployer, contractBackend)

	contractBackend.Commit()

	// beacon owner registers a Beacon
	transactOptsBeaconOwner.Value =  big.NewInt(0)
	beaconOwner, _ := NewBeacon(transactOptsBeaconOwner, beaconAddress, contractBackend)

	_, _ = beaconOwner.RegisterBeacon(beaconIdBytes32, beaconAuthAddresses1)

	contractBackend.Commit()

	//Beacon owner adds new Auth
	transactOptsBeaconOwner.Value = new(big.Int).SetUint64(0)
	beaconOwner, _ = NewBeacon(transactOptsBeaconOwner, beaconAddress, contractBackend)
	_, err := beaconOwner.AddAuthAddresses(beaconIdBytes32, []common.Address{accAddr2})

	if err != nil {
		t.Fatalf("Failed to authorise new address: %v", err)
	}
	contractBackend.Commit()
	msgData, err := packRecordData(beaconIdBytes32, [32]byte{'a'})
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}
	// accAddr2 should be able to record hashes now
	msg := generateBeaconCallMsg(accAddr2, beaconAddress, msgData)

	// dry run the contract call, and process the VM result
	res, _ := contractBackend.PendingCallContract(context.Background(), msg)

	if len(res) != 0 {
		t.Log("msg.res", string(res))
		t.Fatalf("accAddr2 should be able to record hashes now")
	}

	acc2, _ := NewBeacon(transactOptsAcc2, beaconAddress, contractBackend)

	tx, err := acc2.RecordHash(beaconIdBytes32, [32]byte{'a'})

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
		beaconOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)

	transactOptsBeaconOwner := bind.NewKeyedTransactor(beaconOwnerKey)
	transactOptsBeaconOwner.GasLimit = 240000
	transactOptsBeaconOwner.GasPrice = big.NewInt(2500)
	transactOptsBeaconOwner.Value = big.NewInt(0)

	beaconAddress, _, _  := DeployBeacon(transactOptsDeployer, contractBackend)

	contractBackend.Commit()

	msgData, err := packRecordData(beaconIdBytes32, [32]byte{'a'})
	if err != nil {
		t.Fatalf("Failed to pack data: %v", err)
	}

	msg := generateBeaconCallMsg(beaconOwnerAddr, beaconAddress, msgData)

	// run the contract call, and process the result
	res, err := contractBackend.PendingCallContract(context.Background(), msg)

	t.Log("msg", msg)
	t.Log("res", res)
	t.Log("res string", string(res))
	t.Log("err", err)

	if len(res) == 0 {
		t.Fatalf("should have failed and reverted")
	}

	beaconOwner, _ := NewBeacon(transactOptsBeaconOwner, beaconAddress, contractBackend)

	tx, err := beaconOwner.RecordHash(beaconIdBytes32, [32]byte{'a'})

	if err != nil {
		t.Fatalf("error running beaconOwner.RecordHeader: %v", err)
	}

	contractBackend.Commit()

	t.Log("RecordHeader Tx", tx)

	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())

	t.Log("receipt", receipt)

	if receipt.Status != 0 {
		t.Fatalf("Receipt status == true. Should have failed")
	}
}

func TestBeaconExists(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
		beaconOwnerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	transactOptsBeaconOwner := bind.NewKeyedTransactor(beaconOwnerKey)

	beaconAddress, _, _  := DeployBeacon(transactOptsDeployer, contractBackend)
	contractBackend.Commit()

	transactOptsBeaconOwner.Value =  big.NewInt(0)
	beaconOwner, _ := NewBeacon(transactOptsBeaconOwner, beaconAddress, contractBackend)

	_, _ = beaconOwner.RegisterBeacon(beaconIdBytes32, beaconAuthAddresses1)
	contractBackend.Commit()

	regged, err := beaconOwner.BeaconExists(common.BytesToHash([]byte(beaconId)))

	if err != nil {
		t.Fatalf("BeaconExists error: %v", err)
	}

	if !regged {
		t.Fatalf("BeaconExists for %v should be true. Got %v", beaconId, regged)
	}

	chainIdNotExist := "notexist"
	regged, _ = beaconOwner.BeaconExists(common.BytesToHash([]byte(chainIdNotExist)))

	if regged {
		t.Fatalf("BeaconExists for %v should be false. Got %v", chainIdNotExist, regged)
	}
}
