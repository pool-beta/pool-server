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

func TestPercent(t *testing.T) {
	var ok bool

	testNum := Number(38)
	testDen := Number(203)

	p := NewPercent(testNum, testDen)

	fmt.Printf("test percent: %v\n", p.String())

	toOne, ok := p.ToOne()
	if !ok {
		t.Errorf("Unexpected false on ToOne -- percent: %v", p.String())
	}
	fmt.Printf("toOne percent: %v\n", toOne.String())

	total := toOne.Add(p)
	fmt.Printf("total percent: %v\n", total.String())

	ok = total.IsOne()
	if !ok {
		t.Errorf("Expected to be one -- actual percent: %v", total.String())
	}
}