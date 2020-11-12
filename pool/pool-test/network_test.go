package pool_test

import (
	"fmt"
	"testing"

	. "github.com/pool-beta/pool-server/pool"
	. "github.com/pool-beta/pool-server/types"
	. "github.com/pool-beta/pool-server/pool/types"
	. "github.com/pool-beta/pool-server/user/types"
)

func TestSimpleNetwork(t *testing.T) {
	pf := initPoolFactory()

	// Users
	user1 := NewUserID()

	// Create Pool
	pool1, err := pf.CreatePool("pool1", user1, POOL)
	if err != nil {
		t.Errorf("Error in create pool")
	}
	pool1.Fund(USDollar(1000))

	// Create Debit Drain
	debit1, err := pf.CreatePool("debit1", user1, DRAIN)
	if err != nil {
		t.Errorf("Error in create pool")
	}

	// Create Tanks
	tank1, err := pf.CreatePool("tank1", user1, TANK)
	if err != nil {
		t.Errorf("Error in create pool")
	}

	// Connect Pool & Drain
	stream1, err := NewStream(user1, pool1, debit1)
	stream1.SetAllowOverdraft(true)
	stream1.SetMaxOverdraft(USDollar(2500))
	stream1.SetPercentOverdraft(NewPercent(1, 1))

	// Connect Pool & Tank
	stream2, err := NewStream(user1, tank1, pool1)
	stream2.SetAllowOverdraft(true)
	stream2.SetMaxOverdraft(USDollar(2500))
	stream2.SetPercentOverdraft(NewPercent(1, 1))

	// Pull on debit
	flow1, err := NewFlow(debit1, USDollar(1200), PULL)
	if err != nil {
		fmt.Println(err.Error())
		t.Errorf("Error in drain pull -- drain: %v; drop: %v", debit1, flow1)
	}
	flow1.Accept()

	var reserve, expectedAmount USDollar
	expectedAmount = USDollar(0)
	// Check pool reserve
	reserve = pool1.GetReserve()
	if reserve != expectedAmount {
		t.Errorf("Does not match pool -- expected: %v; actual: %v", expectedAmount.String(), reserve.String())
	}

	reserve = debit1.GetReserve()
	if reserve != expectedAmount {
		t.Errorf("Does not match debit -- expected: %v; actual: %v", expectedAmount.String(), reserve.String())
	}

	reserve = tank1.GetReserve()
	if reserve != expectedAmount {
		t.Errorf("Does not match tank -- expected: %v; actual: %v", expectedAmount.String(), reserve.String())
	}

}