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

package beacon

import (
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/contracts/beacon/contract"
)

type Beacon struct {
	*contract.BeaconSession
	contractBackend bind.ContractBackend
}

func NewBeacon(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*Beacon, error) {
	beaconContract, err := contract.NewBeacon(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &Beacon{
		&contract.BeaconSession{
			Contract:     beaconContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}


func DeployBeacon(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *Beacon, error) {
	beaconAddr, _, _, err := contract.DeployBeacon(transactOpts, contractBackend)
	if err != nil {
		return beaconAddr, nil, err
	}

	beacon, err := NewBeacon(transactOpts, beaconAddr, contractBackend)
	if err != nil {
		return beaconAddr, nil, err
	}

	return beaconAddr, beacon, nil
}
