package dsg

import (
	"github.com/unification-com/mainchain/common"
	"math/big"
	"sort"
)

type StakedWallet struct {
	Address common.Address
	Staked  *big.Int
}

func getStakedWallets() []StakedWallet {
	// ToDo - get from stateDb, or create snapshot

	stakedWallets := []StakedWallet{
		{common.HexToAddress("0x001A320943d4535e93d31E4A65a6e21C5dF375D7"), big.NewInt(1000)},
		{common.HexToAddress("0x002A956804bAD8DCad148aBFF71515F9B057F7E0"), big.NewInt(100)},
		{common.HexToAddress("0x003ADc30A6f4DB59d2698e3D3029fd1BA68b6B15"), big.NewInt(10)},
		{common.HexToAddress("0x004A435F1D54aA5cc9FCfA0fEB6B8c4a428bbB93"), big.NewInt(1)},
	}
	return stakedWallets
}

// GetValidatorPool is the exported function for getValidatorPool
func GetValidatorPool() []common.Address {
	return getValidatorPool()
}

// getValidatorPool calculates the top staked wallet addresses and returns
// a list of addresses of size common.NumSignersInEpoch
func getValidatorPool() []common.Address {

	validatorPool := make([]common.Address, 0)

	stakedWallets := getStakedWallets()
	sort.Slice(stakedWallets, func(i, j int) bool {
		cmp := stakedWallets[i].Staked.Cmp(stakedWallets[j].Staked)
		//reverse order
		if cmp == 1 {
			return true
		}
		return false
	})

	top := stakedWallets[:common.NumSignersInEpoch]

	for _, t := range top {
		validatorPool = append(validatorPool, t.Address)
	}

	return validatorPool
}

func EVIdFromEtherbase(etherbase common.Address) uint64 {
	stakedWallets := getStakedWallets()

	for index, stakedWallet := range stakedWallets {
		if stakedWallet.Address == etherbase {
			return uint64(index)
		}
	}

	return uint64(0)
}

func getAddressFromSlotNumber(index uint64) common.Address {
	stakedWallets := getStakedWallets()
	if int(index) < len(stakedWallets) {
		return stakedWallets[index].Address
	}
	return common.Address{}
}

// Take the first EV and append it to the end, shifting the slots
func ShiftEVs(stakedWallets []StakedWallet) []StakedWallet {
	x, stakedWallets := stakedWallets[0], stakedWallets[1:]
	stakedWallets = append(stakedWallets, x)
	return stakedWallets
}
