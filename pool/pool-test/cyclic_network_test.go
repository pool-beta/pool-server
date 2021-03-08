package pool_test

import (
	"fmt"
	"testing"
	"strconv"

	. "github.com/pool-beta/pool-server/pool"
	. "github.com/pool-beta/pool-server/types"
	. "github.com/pool-beta/pool-server/pool/types"
	. "github.com/pool-beta/pool-server/user/types"
)

// Testing 3 pool system; tank-pool1 pool3-drain
func TestCyclicNetwork(t *testing.T) {
	// Testing Parameters
	numOfPools := 3

	pf := initPoolFactory()

	user1 := NewUserID()

	// Create Debit Drain
	drain1, err := pf.CreatePool("debit1", user1, DRAIN)
	if err != nil {
		t.Errorf("Error in create pool")
	}

	// Create Tanks
	tank1, err := pf.CreatePool("tank1", user1, TANK)
	if err != nil {
		t.Errorf("Error in create pool")
	}

	var poolList []Pool
	for i := 0; i < numOfPools; i++ {
		// Create Pool
		poolName := "pool" + strconv.Itoa(i+1)
		pool1, err := pf.CreatePool(poolName, user1, POOL)
		if err != nil {
			t.Errorf("Error in create pool")
		}
		pool1.Fund(USDollar(1000))
		poolList = append(poolList, pool1)
	}

	// Connect with Streams

	// Connects Pools
	for i := 0; i < numOfPools; i++ {
		first_index := i
		second_index := (i + 1) % numOfPools

		stream1, err := NewStream(user1, poolList[first_index], poolList[second_index])
		if err != nil {
			fmt.Println(err.Error())
			t.Errorf("Error in stream create -- pool: %v; stream: %v", poolList[i], stream1)
		}
		stream1.SetAllowOverdraft(true)
		stream1.SetMaxOverdraft(USDollar(3000))
		stream1.SetPercentOverdraft(NewPercent(1, 1))
	}

	// Connect Tank
	stream1, err := NewStream(user1, tank1, poolList[0])
	stream1.SetAllowOverdraft(true)
	stream1.SetMaxOverdraft(USDollar(3000))
	stream1.SetPercentOverdraft(NewPercent(1, 1))

	// Connect Drain
	stream1, err = NewStream(user1, poolList[len(poolList)-1], drain1)
	stream1.SetAllowOverdraft(true)
	stream1.SetMaxOverdraft(USDollar(3000))
	stream1.SetPercentOverdraft(NewPercent(1, 1))

	// Make Transactions (Flow)
	// Pull on debit
	flow1, err := NewFlow(drain1, USDollar(3000), PULL)
	if err != nil {
		fmt.Println(err.Error())
		t.Errorf("Error in drain pull -- drain: %v; drop: %v", drain1, flow1)
	}
	flow1.Accept()

	// Check actual vs expected (Errorf if not correct)
	var reserve, expectedAmount USDollar
	expectedAmount = USDollar(0)
	// Check all pool reserves
	for _, pool1 := range poolList {
		reserve = pool1.GetReserve()
		if reserve != expectedAmount {
			t.Errorf("Does not match pool -- expected: %v; actual: %v", expectedAmount.String(), reserve.String())
		}
	}
}

//Testing 3 pool system; tank-pool1 pool2-drain
func TestCyclicNetwork2(t *testing.T) {
	// Testing Parameters
	numOfPools := 3

	pf := initPoolFactory()

	user1 := NewUserID()

	// Create Debit Drain
	drain1, err := pf.CreatePool("debit1", user1, DRAIN)
	if err != nil {
		t.Errorf("Error in create pool")
	}

	// Create Tanks
	tank1, err := pf.CreatePool("tank1", user1, TANK)
	if err != nil {
		t.Errorf("Error in create pool")
	}

	var poolList []Pool
	for i := 0; i < numOfPools; i++ {
		// Create Pool
		poolName := "pool" + strconv.Itoa(i+1)
		pool1, err := pf.CreatePool(poolName, user1, POOL)
		if err != nil {
			t.Errorf("Error in create pool")
		}
		pool1.Fund(USDollar(1000))
		poolList = append(poolList, pool1)
	}

	// Connect with Streams

	// Connects Pools
	for i := 0; i < numOfPools; i++ {
		first_index := i
		second_index := (i + 1) % numOfPools

		stream1, err := NewStream(user1, poolList[first_index], poolList[second_index])
		if err != nil {
			fmt.Println(err.Error())
			t.Errorf("Error in stream create -- pool: %v; stream: %v", poolList[i], stream1)
		}
		stream1.SetAllowOverdraft(true)
		stream1.SetMaxOverdraft(USDollar(3000))
		stream1.SetPercentOverdraft(NewPercent(1, 1))
	}

	// Connect Tank
	stream1, err := NewStream(user1, tank1, poolList[0])
	stream1.SetAllowOverdraft(true)
	stream1.SetMaxOverdraft(USDollar(3000))
	stream1.SetPercentOverdraft(NewPercent(1, 1))

	// Connect Drain
	stream1, err = NewStream(user1, poolList[len(poolList)-2], drain1)
	stream1.SetAllowOverdraft(true)
	stream1.SetMaxOverdraft(USDollar(3000))
	stream1.SetPercentOverdraft(NewPercent(1, 1))

	// Make Transactions (Flow)
	// Pull on debit
	flow1, err := NewFlow(drain1, USDollar(3000), PULL)
	if err != nil {
		fmt.Println(err.Error())
		t.Errorf("Error in drain pull -- drain: %v; drop: %v", drain1, flow1)
	}
	flow1.Accept()

	// Check actual vs expected (Errorf if not correct)
	var reserve, expectedAmount USDollar
	expectedAmount = USDollar(0)
	// Check all pool reserves
	for _, pool1 := range poolList {
		reserve = pool1.GetReserve()
		if reserve != expectedAmount {
			t.Errorf("Does not match pool -- expected: %v; actual: %v", expectedAmount.String(), reserve.String())
		}
	}
}