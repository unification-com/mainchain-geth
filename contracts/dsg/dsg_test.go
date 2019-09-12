package dsg

import (
	"context"
	"crypto/ecdsa"
	ethereum "github.com/unification-com/mainchain"
	"github.com/unification-com/mainchain/accounts/abi"
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/accounts/abi/bind/backends"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/common/hexutil"
	"github.com/unification-com/mainchain/contracts/dsg/contract"
	"github.com/unification-com/mainchain/core"
	"github.com/unification-com/mainchain/crypto"
	"math/big"
	"strings"
	"testing"
	"time"
)

var (
	deployerKey, _ = crypto.HexToECDSA("e7b1f342c77881ff425d09b83892cc0683435878cf1d3ff2b6e015ada9378ddf")
	deployerAddr   = crypto.PubkeyToAddress(deployerKey.PublicKey)
	ev1Key, _      = crypto.HexToECDSA("d7c15922c5371f9536f8b1cef750d5451b01f57b2865278fef6bd65c9a17a3a3")
	ev1Addr        = crypto.PubkeyToAddress(ev1Key.PublicKey)
	ev2Key, _      = crypto.HexToECDSA("e6d90f41cebb11eb6f5f5d7b71a49f3f4309d78ffb3e66eaa6c7af94c7d99d41")
	ev2Addr        = crypto.PubkeyToAddress(ev2Key.PublicKey)
	ev3Key, _      = crypto.HexToECDSA("dc23751c35323b70ad70ec2d08d9be59b74f8491baccc87123a0c0ffdfa24936")
	ev3Addr        = crypto.PubkeyToAddress(ev3Key.PublicKey)

	genesisBalance = new(big.Int).SetUint64(1000000000000000000)
	minStake       = new(big.Int).SetUint64(1000)
	amountToStake  = new(big.Int).SetUint64(100000)
	unstakeMoreThanStake  = new(big.Int).SetUint64(100001)
	stakeTooSmall  = new(big.Int).SetUint64(10)
)

func initBackend(preallocAddresses []common.Address) (*backends.SimulatedBackend, common.Address) {
	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
	}

	for _, a := range preallocAddresses {
		genesisAlloc[a] = core.GenesisAccount{
			Balance: genesisBalance,
		}
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)
	transactOptsDeployer := bind.NewKeyedTransactor(deployerKey)
	dsgAddress, _, _ := DeployDSG(transactOptsDeployer, contractBackend, minStake)
	contractBackend.Commit()

	return contractBackend, dsgAddress
}

func defaultEvTxOpts(value *big.Int, key *ecdsa.PrivateKey) *bind.TransactOpts {
	transactOpts := bind.NewKeyedTransactor(key)
	transactOpts.Value = value
	transactOpts.GasLimit = 10000000
	transactOpts.GasPrice = big.NewInt(25000)

	return transactOpts
}

func packStakeData() ([]byte, error) {
	dsgAbi, err :=  abi.JSON(strings.NewReader(contract.DSGContractABI))
	if err != nil {
		return nil, err
	}
	packed, err := dsgAbi.Pack("stake")
	if err != nil {
		return nil, err
	}
	return packed, nil
}

func packUnStakeData(amount *big.Int) ([]byte, error) {
	dsgAbi, err :=  abi.JSON(strings.NewReader(contract.DSGContractABI))
	if err != nil {
		return nil, err
	}
	packed, err := dsgAbi.Pack("unStake", amount)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

func generateDSGCallMsg(fromAddr common.Address, dsgAddress common.Address, value *big.Int, data []byte) ethereum.CallMsg {

	msg := ethereum.CallMsg{}
	msg.From = fromAddr
	msg.To = &dsgAddress
	msg.Gas = 10000000
	msg.GasPrice = big.NewInt(25000)
	msg.Value = value
	msg.Data = data

	return msg
}

func TestDSGDeploy(t *testing.T) {

	genesisAlloc := map[common.Address]core.GenesisAccount{
		deployerAddr: {
			Balance: genesisBalance,
		},
	}

	contractBackend := backends.NewSimulatedBackend(genesisAlloc, 100000000)

	transactOpts := bind.NewKeyedTransactor(deployerKey)

	dsgAddress, _, err := DeployDSG(transactOpts, contractBackend, minStake)
	if err != nil {
		t.Fatalf("can't deploy DSG: %v", err)
	}
	contractBackend.Commit()
	d := time.Now().Add(1000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	code, err := contractBackend.CodeAt(ctx, dsgAddress, nil)
	if err != nil {
		t.Fatalf("can't get wrkchainroot code: %v", err)
	}

	t.Log("contract code", hexutil.Encode(code))
	t.Log("contract address", dsgAddress.Hex())
}

func TestStake(t *testing.T) {

	contractBackend, dsgAddress := initBackend([]common.Address{ev1Addr, ev2Addr})

	transactOptsEv1 := defaultEvTxOpts(amountToStake, ev1Key)

	stakeData, err := packStakeData()

	if err != nil {
		t.Fatalf("error packing data: %v", err)
	}

	msg := generateDSGCallMsg(ev1Addr, dsgAddress, amountToStake, stakeData)

	// dry-run the contract call, and process the EVM result
	res, _ := contractBackend.PendingCallContract(context.Background(), msg)
	if len(res) > 0 {
		t.Log("res string", string(res))
		t.Fatalf("Failed to stake")
	}

	// EV1 stakes for real
	ev1, _ := NewDSG(transactOptsEv1, dsgAddress, contractBackend)

	tx, err := ev1.Stake()

	t.Log("Stake Tx", tx)
	if err != nil {
		t.Fatalf("can't stake: %v", err)
	}
	contractBackend.Commit()

	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())

	t.Log("receipt", receipt)

	if receipt.Status != 1 {
		t.Fatalf("Receipt status == false")
	}

	staked, err := ev1.GetStaked(ev1Addr)

	if err != nil {
		t.Fatalf("GetStaked error: %v", err)
	}

	t.Log("staked", staked.String())

	if staked.Cmp(amountToStake) != 0 {
		t.Fatalf("staked should be %v. Got %v",amountToStake.String(), staked.String())
	}
}

func TestStakeTooSmall(t *testing.T) {
	contractBackend, dsgAddress := initBackend([]common.Address{ev1Addr, ev2Addr})
	transactOptsEv1 := defaultEvTxOpts(amountToStake, ev1Key)

	stakeData, err := packStakeData()

	if err != nil {
		t.Fatalf("error packing data: %v", err)
	}

	msg := generateDSGCallMsg(ev1Addr, dsgAddress, stakeTooSmall, stakeData)

	// dry-run the contract call, and process the EVM result
	res, _ := contractBackend.PendingCallContract(context.Background(), msg)
	t.Log("res string", string(res))

	if len(res) == 0 {
		t.Fatalf("should have failed with stake too small error")
	}

	// EV1 stakes for real with too small stake amount
	transactOptsEv1.Value = stakeTooSmall
	ev1, _ := NewDSG(transactOptsEv1, dsgAddress, contractBackend)

	tx, err := ev1.Stake()

	t.Log("Stake Tx", tx)
	t.Log("Stake Tx err", err)

	contractBackend.Commit()

	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())

	t.Log("receipt", receipt)

	if receipt.Status != 0 {
		t.Fatalf("Receipt status == true. Should have failed")
	}

}

func TestUnstake(t *testing.T) {

	contractBackend, dsgAddress := initBackend([]common.Address{ev1Addr, ev2Addr})

	transactOptsEv1 := defaultEvTxOpts(amountToStake, ev1Key)

	// EV1 stakes for real
	ev1, _ := NewDSG(transactOptsEv1, dsgAddress, contractBackend)

	_, _ = ev1.Stake()

	contractBackend.Commit()

	// quick check to see it staked OK
	staked, err := ev1.GetStaked(ev1Addr)

	if staked.Cmp(amountToStake) != 0 {
		t.Fatalf("staked should be %v. Got %v",amountToStake.String(), staked.String())
	}

	// dry run unstake contract call, and process the EVM result
	unStakeData, err := packUnStakeData(amountToStake)
	if err != nil {
		t.Fatalf("error packing data: %v", err)
	}
	msg := generateDSGCallMsg(ev1Addr, dsgAddress, big.NewInt(0), unStakeData)
	res, _ := contractBackend.PendingCallContract(context.Background(), msg)
	if len(res) > 0 {
		t.Log("res string", string(res))
		t.Fatalf("Failed to unstake")
	}

    // EV1 unstakes for real
	transactOptsEv1.Value = big.NewInt(0)
	ev1, _ = NewDSG(transactOptsEv1, dsgAddress, contractBackend)

	tx, err := ev1.UnStake(amountToStake)

	t.Log("UnStake Tx", tx)
	if err != nil {
		t.Fatalf("can't unstake: %v", err)
	}
	contractBackend.Commit()

	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())

	t.Log("receipt", receipt)

	if receipt.Status != 1 {
		t.Fatalf("Receipt status == false")
	}

	staked, err = ev1.GetStaked(ev1Addr)

	if err != nil {
		t.Fatalf("GetStaked error: %v", err)
	}

	t.Log("staked", staked.String())

	if staked.Cmp(big.NewInt(0)) != 0 {
		t.Fatalf("staked should be 0. Got %v", staked.String())
	}
}

func TestUnstakeTooMuch(t *testing.T) {

	contractBackend, dsgAddress := initBackend([]common.Address{ev1Addr})

	transactOptsEv1 := defaultEvTxOpts(amountToStake, ev1Key)

	// EV1 stakes for real
	ev1, _ := NewDSG(transactOptsEv1, dsgAddress, contractBackend)

	_, _ = ev1.Stake()

	contractBackend.Commit()

	// quick check to see it staked OK
	staked, err := ev1.GetStaked(ev1Addr)
	if staked.Cmp(amountToStake) != 0 {
		t.Fatalf("staked should be %v. Got %v",amountToStake.String(), staked.String())
	}

	// dry run unstake contract call, and process the EVM result
	unStakeData, err := packUnStakeData(unstakeMoreThanStake)
	if err != nil {
		t.Fatalf("error packing data: %v", err)
	}
	msg := generateDSGCallMsg(ev1Addr, dsgAddress, big.NewInt(0), unStakeData)
	res, _ := contractBackend.PendingCallContract(context.Background(), msg)
	t.Log("res string", string(res))

	if len(res) == 0 {
		t.Fatalf("should have failed to unstake")
	}

	// EV1 unstakes for real
	transactOptsEv1.Value = big.NewInt(0)
	ev1, _ = NewDSG(transactOptsEv1, dsgAddress, contractBackend)

	tx, err := ev1.UnStake(unstakeMoreThanStake)

	t.Log("Stake Tx", tx)
	t.Log("Stake Tx err", err)

	contractBackend.Commit()

	receipt, _ := contractBackend.TransactionReceipt(context.Background(), tx.Hash())

	t.Log("receipt", receipt)

	if receipt.Status != 0 {
		t.Fatalf("Receipt status == true. Should have failed")
	}

	staked, err = ev1.GetStaked(ev1Addr)

	t.Log("staked", staked.String())

	if staked.Cmp(amountToStake) != 0 {
		t.Fatalf("staked should still be %v. Got %v", amountToStake.String(), staked.String())
	}
}

func TestAttemptMultipleStakesInDifferentBlocks(t *testing.T) {
	contractBackend, dsgAddress := initBackend([]common.Address{ev1Addr})
	transactOptsEv1 := defaultEvTxOpts(amountToStake, ev1Key)

	// EV1 first stake in this block
	ev1, _ := NewDSG(transactOptsEv1, dsgAddress, contractBackend)
	tx1, _ := ev1.Stake()

	contractBackend.Commit()

	receipt1, _ := contractBackend.TransactionReceipt(context.Background(), tx1.Hash())

	// Second attempt
	tx2, _ := ev1.Stake()

	t.Log("Stake Tx 2", tx2)

	contractBackend.Commit()

	if receipt1.Status != 1 {
		t.Fatal("Tx1 should have been successful")
	}

	receipt2, _ := contractBackend.TransactionReceipt(context.Background(), tx2.Hash())

	if receipt2.Status != 1 {
		t.Fatal("Tx2 should have been successful")
	}

	staked, _ := ev1.GetStaked(ev1Addr)

	t.Log("staked", staked.String())

	newStakedTotal := big.NewInt(0)
	newStakedTotal.Mul(amountToStake, big.NewInt(2))

	if staked.Cmp(newStakedTotal) != 0 {
		t.Fatalf("staked should be %v. Got %v", newStakedTotal.String(), staked.String())
	}
}

func TestAttemptMultipleStakesInSameBlock(t *testing.T) {
	contractBackend, dsgAddress := initBackend([]common.Address{ev1Addr})
	transactOptsEv1 := defaultEvTxOpts(amountToStake, ev1Key)

	// EV1 first stake in this block
	ev1, _ := NewDSG(transactOptsEv1, dsgAddress, contractBackend)
	tx1, _ := ev1.Stake()

	t.Log("Stake Tx 1", tx1)

	// Second attempt in same block. Should fail. Note - not committing the block between Txs
	tx2, _ := ev1.Stake()

	t.Log("Stake Tx 2", tx2)

	contractBackend.Commit()

	receipt1, _ := contractBackend.TransactionReceipt(context.Background(), tx1.Hash())

	if receipt1.Status != 1 {
		t.Fatal("Tx1 should have been successful")
	}

	receipt2, _ := contractBackend.TransactionReceipt(context.Background(), tx2.Hash())

	if receipt2.Status != 0 {
		t.Fatal("Tx2 should have failed")
	}

	staked, _ := ev1.GetStaked(ev1Addr)

	t.Log("staked", staked.String())

	if staked.Cmp(amountToStake) != 0 {
		t.Fatalf("staked should still be %v. Got %v", amountToStake.String(), staked.String())
	}
}

func TestAttemptStakeUnstakeInSameBlock(t *testing.T) {
	contractBackend, dsgAddress := initBackend([]common.Address{ev1Addr})
	transactOptsEv1 := defaultEvTxOpts(amountToStake, ev1Key)

	// EV1 first stake in this block
	ev1, _ := NewDSG(transactOptsEv1, dsgAddress, contractBackend)
	tx1, _ := ev1.Stake()

	// Attempt to unstake in same block. Should fail. Note - not committing the block between Txs
	transactOptsEv1.Value = big.NewInt(0)
	ev1, _ = NewDSG(transactOptsEv1, dsgAddress, contractBackend)
	tx2, _ := ev1.UnStake(amountToStake)

	contractBackend.Commit()

	receipt1, _ := contractBackend.TransactionReceipt(context.Background(), tx1.Hash())

	if receipt1.Status != 1 {
		t.Fatal("Tx1 should have been successful")
	}

	receipt2, _ := contractBackend.TransactionReceipt(context.Background(), tx2.Hash())

	if receipt2.Status != 0 {
		t.Fatal("Tx2 should have failed")
	}

	staked, _ := ev1.GetStaked(ev1Addr)

	t.Log("staked", staked.String())

	if staked.Cmp(amountToStake) != 0 {
		t.Fatalf("staked should still be %v. Got %v", amountToStake.String(), staked.String())
	}
}
