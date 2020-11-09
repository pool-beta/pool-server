package types

import (
	"fmt"
	"testing"
)

func TestUSDollar(t *testing.T) {
	ben, _ := NewUSDollar(100, 0)
	lincoln, _ := NewUSDollar(0, 5)

	fmt.Printf("ben: %v\n", ben.String())
	fmt.Printf("lincoln: %v\n", lincoln.String())

	// Test Invalid
	var err error
	_, err = NewUSDollar(-100, 0)
	if err == nil {
		t.Errorf("Should not allow negative dollars")
	}

	_, err = NewUSDollar(0, -1)
	if err == nil {
		t.Errorf("Should not allow negative dollars")
	}

	_, err = NewUSDollar(0, 100)
	if err == nil {
		t.Errorf("Should not allow more than 99 cents")
	}
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

func TestMultiplyPercent(t *testing.T) {
	initialAmount, _ := NewUSDollar(1000, 50)

	percents := []Percent{
		NewPercent(1, 2),
		NewPercent(1, 4), 
	}

	e1, _ := NewUSDollar(500, 25)
	e2, _ := NewUSDollar(250, 12)

	expectedValues := []USDollar{
		e1,
		e2, // Round Down
	}

	for i, p := range percents {
		value := initialAmount.MultiplyPercent(p)
		if value != expectedValues[i] {
			t.Errorf("Does not match -- initial amount: %v; percent: %v; expected: %v; actual: %v", initialAmount.String(), p.String(), expectedValues[i].String(), value.String())
		}
	}
}