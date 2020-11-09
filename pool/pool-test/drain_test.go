package pool_test

import (
	"testing"

// 	. "github.com/pool-beta/pool-server/pool"
// 	. "github.com/pool-beta/pool-server/types"
)

func TestCreateDrain(t *testing.T) {
	// Init PoolFactory
	pf := initPoolFactory()
	if pf == nil {
		t.Errorf("Error in creating PoolFactory")
	}
}
