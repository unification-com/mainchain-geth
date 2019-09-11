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

package dsg

import (
	"github.com/unification-com/mainchain/accounts/abi/bind"
	"github.com/unification-com/mainchain/common"
	dsgcontract "github.com/unification-com/mainchain/contracts/dsg/contract"
	"math/big"
)

type DSG struct {
	*dsgcontract.DSGContractSession
	contractBackend bind.ContractBackend
}


func NewDSG(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*DSG, error) {
	dsg, err := dsgcontract.NewDSGContract(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &DSG{
		&dsgcontract.DSGContractSession{
			Contract:     dsg,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}


func DeployDSG(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, minAllowedStake *big.Int) (common.Address, *DSG, error) {
	dsgAddr, _, _, err := dsgcontract.DeployDSGContract(transactOpts, contractBackend, minAllowedStake)
	if err != nil {
		return dsgAddr, nil, err
	}

	dsg, err := NewDSG(transactOpts, dsgAddr, contractBackend)
	if err != nil {
		return dsgAddr, nil, err
	}

	return dsgAddr, dsg, nil
}
