package pool_test

import (
	"fmt"
	"testing"

	. "github.com/pool-beta/pool-server/pool/types"
)

func TestNewPoolID(t *testing.T) {
	pid := NewPoolID()
	fmt.Printf("Pool ID: %v\n", pid)
}
