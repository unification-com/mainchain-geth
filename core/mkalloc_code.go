// Copyright 2019 The Unification Authors
// This file is part of the Unification mainchain library.
//
// The Unification mainchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Unification mainchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Unification mainchain library. If not, see <http://www.gnu.org/licenses/>.

// +build none

/*

   The mkalloc_code tool creates the genesis allocation constants in genesis_alloc.go
   It outputs a const declaration that contains an RLP-encoded list of (address, balance, code, storage) tuples.

       go run mkalloc_code.go genesis.json

*/
package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"

	"github.com/unification-com/mainchain/core"
	"github.com/unification-com/mainchain/rlp"
)

type allocItem struct{ Addr, Balance *big.Int
	                   Code []byte
                       Storage []byte
}

type allocList []allocItem

func (a allocList) Len() int           { return len(a) }
func (a allocList) Less(i, j int) bool { return a[i].Addr.Cmp(a[j].Addr) < 0 }
func (a allocList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func makelist(g *core.Genesis) allocList {
	a := make(allocList, 0, len(g.Alloc))
	for addr, account := range g.Alloc {
		bigAddr := new(big.Int).SetBytes(addr.Bytes())
		encBuf := new(bytes.Buffer)
		err := gob.NewEncoder(encBuf).Encode(account.Storage)
		if err != nil {
			panic(err)
		}
		a = append(a, allocItem{bigAddr, account.Balance, account.Code, encBuf.Bytes()})
	}
	sort.Sort(a)
	return a
}

func makealloc(g *core.Genesis) string {
	a := makelist(g)
	data, err := rlp.EncodeToBytes(a)
	if err != nil {
		panic(err)
	}
	return strconv.QuoteToASCII(string(data))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: mkalloc genesis.json")
		os.Exit(1)
	}

	g := new(core.Genesis)
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(file).Decode(g); err != nil {
		panic(err)
	}
	allocData := makealloc(g)
	fmt.Println("const allocData =", allocData)
}
