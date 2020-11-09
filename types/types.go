package types

import (
	"fmt"
)

type PoolID = uint64
type UserID = uint64
type USDollar uint64

func NewUSDollar(dollar int, cent int) USDollar {
	d := USDollar(dollar) * USDollar(100)
	c := USDollar(cent)
	return d + c
}

func (us *USDollar) String() string {
	d := *us / USDollar(100)
	c := *us % USDollar(100)

	lead := ""
	if c < 10 {
		lead = "0"
	}

	return fmt.Sprintf("$%v.%v%v", d, lead, c)
}