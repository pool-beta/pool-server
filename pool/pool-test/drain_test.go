package pool_test

import (
	"testing"

	. "github.com/pool-beta/pool-server/pool"
	. "github.com/pool-beta/pool-server/types"
)

func TestCreateDrain(t *testing.T) {
	var err error

	// Init PoolFactory
	pf := initPoolFactory()
	if pf == nil {
		t.Errorf("Error in creating PoolFactory")
	}

	user1 := NewUserID()

	// Create Pool
	pool1, err := pf.CreatePool("pool1", user1)
	pool1.Fund(USDollar(100))

	// Create Debit Drain
	debit1 := NewDrain("debit1", user1)

	// Connect Pool & Drain
	stream, err := NewStream(user1, pool1, debit1)
	if err != nil {
		t.Errorf("Error in NewStream -- pool1: %v; pool2: %v; user: %v", pool1, debit1, user1)
	}
	debit1.AddPusher(stream)
	pool1.AddPuller(stream)

	// Drop
	err = debit1.Push(nil)
	if err == nil {
		t.Errorf("Should be able to push a drain")
	}
}
