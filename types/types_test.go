package types

import (
	"fmt"
	"testing"
)

func TestUSDollar(t *testing.T) {
	ben := NewUSDollar(100, 0)
	lincoln := NewUSDollar(0, 5)

	fmt.Printf("ben: %v\n", ben.String())
	fmt.Printf("lincoln: %v\n", lincoln.String())
}