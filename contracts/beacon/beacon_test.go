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

func generateBeaconCallMsg(fromAddr common.Address, beaconAddress common.Address, data []byte) ethereum.CallMsg {

	msg := ethereum.CallMsg{}
	msg.From = fromAddr
	msg.To = &beaconAddress
	msg.Gas = 10000000
	msg.GasPrice = big.NewInt(25000)
	msg.Value = big.NewInt(0)
	msg.IsWrkchainRootMessage = true
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
