package types

import (
	"fmt"
	"testing"
)

func TestNewPoolID(t *testing.T) {
	pid := NewPoolID()
	fmt.Printf("Pool ID: %v\n", pid)
}