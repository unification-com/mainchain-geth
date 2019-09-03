// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package state

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/unification-com/mainchain/common"
	"github.com/unification-com/mainchain/core/rawdb"
	"github.com/unification-com/mainchain/crypto"
	"github.com/unification-com/mainchain/ethdb"
	checker "gopkg.in/check.v1"
)

type StateSuite struct {
	db    ethdb.Database
	state *StateDB
}

var _ = checker.Suite(&StateSuite{})

var toAddr = common.BytesToAddress

func (s *StateSuite) TestDump(c *checker.C) {
	// generate a few entries
	obj1 := s.state.GetOrNewStateObject(toAddr([]byte{0x01}))
	obj1.AddBalance(big.NewInt(22))
	obj1.AddLockedAmount(big.NewInt(11))
	obj2 := s.state.GetOrNewStateObject(toAddr([]byte{0x01, 0x02}))
	obj2.SetCode(crypto.Keccak256Hash([]byte{3, 3, 3, 3, 3, 3, 3}), []byte{3, 3, 3, 3, 3, 3, 3})
	obj3 := s.state.GetOrNewStateObject(toAddr([]byte{0x02}))
	obj3.SetBalance(big.NewInt(44))
	obj3.SetLockedAmount(big.NewInt(44))

	// write some of them to the trie
	s.state.updateStateObject(obj1)
	s.state.updateStateObject(obj2)
	s.state.Commit(false)

	// check that dump contains the state objects that are in trie
	got := string(s.state.Dump(false, false, true))
	want := `{
    "root": "5e416d59f7f2dc5fbccbebf05b781601bf5c96ad7fdbcda38ae79bd53e228628",
    "accounts": {
        "0x0000000000000000000000000000000000000001": {
            "balance": "22",
            "nonce": 0,
            "root": "56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
            "codeHash": "c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",
            "locked": true,
            "lockedAmount": "11"
        },
        "0x0000000000000000000000000000000000000002": {
            "balance": "44",
            "nonce": 0,
            "root": "56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
            "codeHash": "c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",
            "locked": true,
            "lockedAmount": "44"
        },
        "0x0000000000000000000000000000000000000102": {
            "balance": "0",
            "nonce": 0,
            "root": "56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
            "codeHash": "87874902497a5bb968da31a2998d8f22e949d1ef6214bcdedd8bae24cca4b9e3",
            "code": "03030303030303",
            "locked": false,
            "lockedAmount": "0"
        }
    }
}`
	if got != want {
		c.Errorf("dump mismatch:\ngot: %s\nwant: %s\n", got, want)
	}
}

func (s *StateSuite) SetUpTest(c *checker.C) {
	s.db = rawdb.NewMemoryDatabase()
	s.state, _ = New(common.Hash{}, NewDatabase(s.db))
}

func (s *StateSuite) TestNull(c *checker.C) {
	address := common.HexToAddress("0x823140710bf13990e4500136726d8b55")
	s.state.CreateAccount(address)
	//value := common.FromHex("0x823140710bf13990e4500136726d8b55")
	var value common.Hash

	s.state.SetState(address, common.Hash{}, value)
	s.state.Commit(false)

	if value := s.state.GetState(address, common.Hash{}); value != (common.Hash{}) {
		c.Errorf("expected empty current value, got %x", value)
	}
	if value := s.state.GetCommittedState(address, common.Hash{}); value != (common.Hash{}) {
		c.Errorf("expected empty committed value, got %x", value)
	}
}

func (s *StateSuite) TestSnapshot(c *checker.C) {
	stateobjaddr := toAddr([]byte("aa"))
	var storageaddr common.Hash
	data1 := common.BytesToHash([]byte{42})
	data2 := common.BytesToHash([]byte{43})

	// snapshot the genesis state
	genesis := s.state.Snapshot()

	// set initial state object value
	s.state.SetState(stateobjaddr, storageaddr, data1)
	snapshot := s.state.Snapshot()

	// set a new state object value, revert it and ensure correct content
	s.state.SetState(stateobjaddr, storageaddr, data2)
	s.state.RevertToSnapshot(snapshot)

	c.Assert(s.state.GetState(stateobjaddr, storageaddr), checker.DeepEquals, data1)
	c.Assert(s.state.GetCommittedState(stateobjaddr, storageaddr), checker.DeepEquals, common.Hash{})

	// revert up to the genesis state and ensure correct content
	s.state.RevertToSnapshot(genesis)
	c.Assert(s.state.GetState(stateobjaddr, storageaddr), checker.DeepEquals, common.Hash{})
	c.Assert(s.state.GetCommittedState(stateobjaddr, storageaddr), checker.DeepEquals, common.Hash{})
}

func (s *StateSuite) TestSnapshotEmpty(c *checker.C) {
	s.state.RevertToSnapshot(s.state.Snapshot())
}

// use testing instead of checker because checker does not support
// printing/logging in tests (-check.vv does not work)
func TestSnapshot2(t *testing.T) {
	state, _ := New(common.Hash{}, NewDatabase(rawdb.NewMemoryDatabase()))

	stateobjaddr0 := toAddr([]byte("so0"))
	stateobjaddr1 := toAddr([]byte("so1"))
	var storageaddr common.Hash

	data0 := common.BytesToHash([]byte{17})
	data1 := common.BytesToHash([]byte{18})

	state.SetState(stateobjaddr0, storageaddr, data0)
	state.SetState(stateobjaddr1, storageaddr, data1)

	// db, trie are already non-empty values
	so0 := state.getStateObject(stateobjaddr0)
	so0.SetBalance(big.NewInt(42))
	so0.SetNonce(43)
	so0.SetCode(crypto.Keccak256Hash([]byte{'c', 'a', 'f', 'e'}), []byte{'c', 'a', 'f', 'e'})
	so0.SetLockedAmount(big.NewInt(0))
	so0.suicided = false
	so0.deleted = false
	state.setStateObject(so0)

	root, _ := state.Commit(false)
	state.Reset(root)

	// and one with deleted == true
	so1 := state.getStateObject(stateobjaddr1)
	so1.SetBalance(big.NewInt(52))
	so1.SetNonce(53)
	so1.SetCode(crypto.Keccak256Hash([]byte{'c', 'a', 'f', 'e', '2'}), []byte{'c', 'a', 'f', 'e', '2'})
	so1.SetLockedAmount(big.NewInt(30))
	so1.suicided = true
	so1.deleted = true
	state.setStateObject(so1)

	so1 = state.getStateObject(stateobjaddr1)
	if so1 != nil {
		t.Fatalf("deleted object not nil when getting")
	}

	snapshot := state.Snapshot()
	state.RevertToSnapshot(snapshot)

	so0Restored := state.getStateObject(stateobjaddr0)
	// Update lazily-loaded values before comparing.
	so0Restored.GetState(state.db, storageaddr)
	so0Restored.Code(state.db)
	// non-deleted is equal (restored)
	compareStateObjects(so0Restored, so0, t)

	// deleted should be nil, both before and after restore of state copy
	so1Restored := state.getStateObject(stateobjaddr1)
	if so1Restored != nil {
		t.Fatalf("deleted object not nil after restoring snapshot: %+v", so1Restored)
	}
}

func TestSimpleUndLocking(t *testing.T) {

	// Create a random state test
	stateObj, _ := New(common.Hash{}, NewDatabase(rawdb.NewMemoryDatabase()))

	// Add locked amount, and check it's locked
	for i := byte(0); i < 255; i++ {
		obj := stateObj.GetOrNewStateObject(common.BytesToAddress([]byte{i}))
		obj.AddBalance(big.NewInt(int64(i + 1)))
		obj.AddLockedAmount(big.NewInt(int64(i + 1)))
		stateObj.updateStateObject(obj)
	}
	stateObj.Finalise(false)

	for i := byte(0); i < 255; i++ {
		obj := stateObj.GetOrNewStateObject(common.BytesToAddress([]byte{i}))

		if want := big.NewInt(int64(i + 1)); obj.LockedAmount().Cmp(want) != 0 {
			t.Errorf("obj %d: locked amount mismatch: have %v, want %v", i, obj.LockedAmount(), want)
		}
		if want := true; obj.Locked() != want {
			t.Errorf("obj %d: locked mismatch: have %v, want %v", i, obj.Locked(), want)
		}
	}
}

func TestLockAndUnlockUnd(t *testing.T) {

	stateObj, _ := New(common.Hash{}, NewDatabase(rawdb.NewMemoryDatabase()))
	// Add Balance, and Locked Amount (Balance - 1)
	for i := byte(0); i < 255; i++ {
		obj := stateObj.GetOrNewStateObject(common.BytesToAddress([]byte{i}))
		obj.AddBalance(big.NewInt(int64(i) + 1))
		stateObj.updateStateObject(obj)
	}
	stateObj.Finalise(false)

	for i := byte(0); i < 255; i++ {
		obj := stateObj.GetOrNewStateObject(common.BytesToAddress([]byte{i}))
		//t.Logf("balance %v, locked amount %v, available %v, locked %v", obj.Balance(), obj.LockedAmount(), obj.Available(), obj.Locked())
		if want := false; obj.Locked() != want {
			t.Errorf("locked mismatch: have %v, want %v", obj.Locked(), want)
		}

		// Add some more balance and locked UND
		obj.AddBalance(big.NewInt(int64(i) + 1))
		obj.AddLockedAmount(big.NewInt(int64(i) + 1))
		stateObj.updateStateObject(obj)
	}

	stateObj.Finalise(false)

	for i := byte(0); i < 255; i++ {
		obj := stateObj.GetOrNewStateObject(common.BytesToAddress([]byte{i}))
		//t.Logf("balance %v, locked amount %v, available %v, locked %v", obj.Balance(), obj.LockedAmount(), obj.Available(), obj.Locked())
		if want := true; obj.Locked() != want {
			t.Errorf("locked mismatch: have %v, want %v", obj.Locked(), want)
		}
		// Sub some balance and locked UND
		obj.SubBalance(big.NewInt(int64(i) + 1))
		obj.SubLockedAmount(big.NewInt(int64(i) + 1))
		stateObj.updateStateObject(obj)
	}

	stateObj.Finalise(false)

	for i := byte(0); i < 255; i++ {
		obj := stateObj.GetOrNewStateObject(common.BytesToAddress([]byte{i}))
		//t.Logf("balance %v, locked amount %v, available %v, locked %v", obj.Balance(), obj.LockedAmount(), obj.Available(), obj.Locked())
		if want := false; obj.Locked() != want {
			t.Errorf("locked mismatch: have %v, want %v", obj.Locked(), want)
		}
	}
}

func TestUndAvailable(t *testing.T) {
	// Create a random state test
	stateObj, _ := New(common.Hash{}, NewDatabase(rawdb.NewMemoryDatabase()))

	// Add Balance, and Locked Amount (Balance - 1)
	for i := byte(0); i < 255; i++ {
		obj := stateObj.GetOrNewStateObject(common.BytesToAddress([]byte{i}))
		obj.AddBalance(big.NewInt(int64(i) + 1))
		obj.AddLockedAmount(big.NewInt(int64(i)))
		stateObj.updateStateObject(obj)
	}
	stateObj.Finalise(false)

	// All accounts should have 1 Available
	for i := byte(0); i < 255; i++ {
		obj := stateObj.GetOrNewStateObject(common.BytesToAddress([]byte{i}))
		if want := big.NewInt(int64(1)); obj.Available().Cmp(want) != 0 {
			t.Errorf("obj %d: amount available mismatch: have %v, want %v", i, obj.Available(), want)
		}
	}
}

func TestCannotLockBelowZero(t *testing.T) {
	var amount int64 = 20
	var amountLocked int64 = 10
	addr := byte(1)
	stateObj, _ := New(common.Hash{}, NewDatabase(rawdb.NewMemoryDatabase()))
	obj := stateObj.GetOrNewStateObject(common.BytesToAddress([]byte{addr}))
	obj.AddBalance(big.NewInt(amount))
	obj.AddLockedAmount(big.NewInt(amountLocked))
	stateObj.updateStateObject(obj)
	stateObj.Finalise(false)

	for i := int64(0); i <= (amountLocked * 2); i++ {
		obj := stateObj.GetOrNewStateObject(common.BytesToAddress([]byte{addr}))
		obj.SubLockedAmount(big.NewInt(int64(1)))
		stateObj.updateStateObject(obj)
		//t.Logf("available %v, locked amount %v, balance %v, locked %v", obj.Available(), obj.LockedAmount(), obj.Balance(), obj.Locked())
	}
	stateObj.Finalise(false)

	obj = stateObj.GetOrNewStateObject(common.BytesToAddress([]byte{addr}))
	if want := big.NewInt(int64(0)); obj.LockedAmount().Cmp(want) != 0 {
		t.Errorf("locked amount mismatch: have %v, want %v", obj.LockedAmount(), want)
	}

	if want := false; obj.Locked() != want {
		t.Errorf("locked mismatch: have %v, want %v", obj.Locked(), want)
	}

	if want := big.NewInt(amount); obj.Available().Cmp(want) != 0 {
		t.Errorf("available mismatch: have %v, want %v", obj.Available(), want)
	}
}

func compareStateObjects(so0, so1 *stateObject, t *testing.T) {
	if so0.Address() != so1.Address() {
		t.Fatalf("Address mismatch: have %v, want %v", so0.address, so1.address)
	}
	if so0.Balance().Cmp(so1.Balance()) != 0 {
		t.Fatalf("Balance mismatch: have %v, want %v", so0.Balance(), so1.Balance())
	}
	if so0.Nonce() != so1.Nonce() {
		t.Fatalf("Nonce mismatch: have %v, want %v", so0.Nonce(), so1.Nonce())
	}
	if so0.LockedAmount().Cmp(so1.LockedAmount()) != 0 {
		t.Fatalf("LockedAmount mismatch: have %v, want %v", so0.LockedAmount(), so1.LockedAmount())
	}
	if so0.Locked() != so1.Locked() {
		t.Fatalf("Locked mismatch: have %v, want %v", so0.Locked(), so1.Locked())
	}
	if so0.data.Root != so1.data.Root {
		t.Errorf("Root mismatch: have %x, want %x", so0.data.Root[:], so1.data.Root[:])
	}
	if !bytes.Equal(so0.CodeHash(), so1.CodeHash()) {
		t.Fatalf("CodeHash mismatch: have %v, want %v", so0.CodeHash(), so1.CodeHash())
	}
	if !bytes.Equal(so0.code, so1.code) {
		t.Fatalf("Code mismatch: have %v, want %v", so0.code, so1.code)
	}

	if len(so1.dirtyStorage) != len(so0.dirtyStorage) {
		t.Errorf("Dirty storage size mismatch: have %d, want %d", len(so1.dirtyStorage), len(so0.dirtyStorage))
	}
	for k, v := range so1.dirtyStorage {
		if so0.dirtyStorage[k] != v {
			t.Errorf("Dirty storage key %x mismatch: have %v, want %v", k, so0.dirtyStorage[k], v)
		}
	}
	for k, v := range so0.dirtyStorage {
		if so1.dirtyStorage[k] != v {
			t.Errorf("Dirty storage key %x mismatch: have %v, want none.", k, v)
		}
	}
	if len(so1.originStorage) != len(so0.originStorage) {
		t.Errorf("Origin storage size mismatch: have %d, want %d", len(so1.originStorage), len(so0.originStorage))
	}
	for k, v := range so1.originStorage {
		if so0.originStorage[k] != v {
			t.Errorf("Origin storage key %x mismatch: have %v, want %v", k, so0.originStorage[k], v)
		}
	}
	for k, v := range so0.originStorage {
		if so1.originStorage[k] != v {
			t.Errorf("Origin storage key %x mismatch: have %v, want none.", k, v)
		}
	}
}
