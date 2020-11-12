package pool_test

import (
	
)

import (
	"fmt"
	"testing"

	. "github.com/pool-beta/pool-server/pool"
	. "github.com/pool-beta/pool-server/types"
	. "github.com/pool-beta/pool-server/pool/types"
	. "github.com/pool-beta/pool-server/user/types"
)

func TestSimpleTank(t *testing.T) {
	var err error
	pullAmount, _ := NewUSDollar(25, 0)
	numOfPulls := 100

	// Init PoolFactory
	pf := initPoolFactory()
	if pf == nil {
		t.Errorf("Error in creating PoolFactory")
	}

	user1 := NewUserID()

	// Create Pool
	pool1, err := pf.CreatePool("pool1", user1, POOL)

	// Create Tank
	tank1, err := pf.CreatePool("tank1", user1, TANK)

	// Connect Pool & Tank
	stream1, err := NewStream(user1, tank1, pool1)
	stream1.SetAllowOverdraft(true)
	stream1.SetMaxOverdraft(USDollar(2500))
	stream1.SetPercentOverdraft(NewPercent(1, 1))

	if err != nil {
		t.Errorf("Error in NewStream -- pool1: %v; pool2: %v; user: %v", pool1, tank1, user1)
	}

	// Pull on tank1
	for i := 0; i < numOfPulls; i++ {
		flow1, err := NewFlow(pool1, pullAmount, PULL)
		if err != nil {
			fmt.Println(err.Error())
			t.Errorf("Error in pool pull -- pool: %v; flow: %v", pool1, flow1)
		}

		// Check pool reserve
		reserve := pool1.GetReserve()
		expected := pullAmount * USDollar(i)
		if reserve != expected {
			t.Errorf("Does not match -- expected: %v; actual: %v", expected.String(), reserve.String())
		}

		flow1.Accept()

		// Check pool reserve
		reserve = pool1.GetReserve()
		expected = pullAmount * USDollar(i + 1)
		if reserve != expected {
			t.Errorf("Does not match -- expected: %v; actual: %v", expected.String(), reserve.String())
		}
	}
}

