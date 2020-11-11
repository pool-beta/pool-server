package pool_test

import (
	"testing"

	. "github.com/pool-beta/pool-server/pool"
	. "github.com/pool-beta/pool-server/pool/types"
	. "github.com/pool-beta/pool-server/user/types"
)

func TestSinglePoolMethods(t *testing.T) {
	// Init PoolFactory
	pf := initPoolFactory()
	if pf == nil {
		t.Errorf("Error in creating PoolFactory")
	}

	var ok bool
	var err error
	// Create Users
	user1 := NewUserID()
	user2 := NewUserID()
	
	// Create Pool
	pool, err := pf.CreatePool("pool1", user1, POOL)
	if err != nil {
		t.Errorf("Error in CreatePool for user %v", user1)
	}
	ok = pool.AdminCheck(user1, "owner")
	if !ok {
		t.Errorf("Creater should be owner of pool -- user: %v", user1)
	}

	// Invalid Check
	ok = pool.AdminCheck(user2, "owner")
	if ok {
		t.Errorf("Should not be owner -- user: %v", user2)
	}

	// Add/Remove Owner
	err = pool.AddOwner(user2)
	if err != nil {
		t.Errorf("Error in adding owner -- user: %v", user2)
	}
	ok = pool.AdminCheck(user2, "owner")
	if !ok {
		t.Errorf("Should be owner of pool -- user: %v", user2)
	}
}

func initPoolFactory() PoolFactory {
	poolFactory, err := NewPoolFactory()
	if err != nil {
		return nil
	}
	return poolFactory
}
