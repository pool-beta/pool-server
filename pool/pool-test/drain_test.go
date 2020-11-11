package pool_test

import (
	"fmt"
	"testing"

	. "github.com/pool-beta/pool-server/pool"
	. "github.com/pool-beta/pool-server/types"
)

func TestSimpleDrain(t *testing.T) {
	var err error

	initialAmount, _ := NewUSDollar(100, 0)
	pullAmount, _ := NewUSDollar(25, 0)
	expectedAmount, _ := NewUSDollar(75, 0)
	
	// Init PoolFactory
	pf := initPoolFactory()
	if pf == nil {
		t.Errorf("Error in creating PoolFactory")
	}

	user1 := NewUserID()

	// Create Pool
	pool1, err := pf.CreatePool("pool1", user1, POOL)
	pool1.Fund(initialAmount)

	// Create Debit Drain
	debit1, err := pf.CreatePool("debit1", user1, DRAIN)

	// Connect Pool & Drain
	stream1, err := NewStream(user1, pool1, debit1)
	stream1.SetAllowOverdraft(true)
	stream1.SetMaxOverdraft(USDollar(2500))
	stream1.SetPercentOverdraft(NewPercent(1, 1))

	if err != nil {
		t.Errorf("Error in NewStream -- pool1: %v; pool2: %v; user: %v", pool1, debit1, user1)
	}

	// Pull Drop
	drop1 := NewDrop(debit1, pullAmount)
	err = debit1.Pull(drop1)
	if err != nil {
		fmt.Println(err.Error())
		t.Errorf("Error in drain pull -- drain: %v; drop: %v", debit1, drop1)
	}

	// Check pool reserve
	reserve := pool1.GetReserve()
	if reserve != expectedAmount {
		t.Errorf("Does not match -- expected: %v; actual: %v", expectedAmount.String(), reserve.String())
	}
}

func TestDrainInvalidMethods(t *testing.T) {
	var err error

	// Init PoolFactory
	pf := initPoolFactory()
	if pf == nil {
		t.Errorf("Error in creating PoolFactory")
	}

	user1 := NewUserID()
	drain1, err := pf.CreatePool("drain1", user1, DRAIN)

	err = drain1.Push(nil)
	if err == nil {
		t.Errorf("Drain should be not able to push")
	}
}


