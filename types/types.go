package types

import (
	"fmt"
)

type PoolID = uint64
type UserID = uint64

/* USDollar */
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


/* Percent */
type Number uint8
type Percent interface {
	Numerator() Number
	Denominator() Number
	Add(Percent) Percent
	ToOne() (Percent, bool)
	IsOne() bool
	String() string
}

type percent struct {
	numerator Number
	denominator Number
}

func NewPercent(numerator Number, denominator Number) Percent {
	return &percent{
		numerator: numerator,
		denominator: denominator,
	}
}

func (p *percent) Add(other Percent) Percent {
	x1 := p.numerator * other.Denominator()
	x2 := other.Numerator() * p.denominator
	x := x1 + x2

	y := p.denominator * other.Denominator()

	divisor := gcd(x, y)

	numerator := x / divisor
	denominator := y / divisor

	return &percent{
		numerator: numerator,
		denominator: denominator,
	}
}

func (p *percent) ToOne() (Percent, bool) {
	if p.numerator > p.denominator {
		return nil, false
	}

	x := p.denominator - p.numerator
	divisor := gcd(x, p.denominator)

	numerator := x / divisor
	denominator := p.denominator / divisor

	return &percent{
		numerator: numerator,
		denominator: denominator,
	}, true
}

func (p *percent) IsOne() bool {
	return p.numerator == p.denominator
}

func (p *percent) String() string {
	numFloat := float64(p.numerator)
	denFloat := float64(p.denominator)
	percentFloat := numFloat / denFloat * float64(100)

	return fmt.Sprintf("%v%%", percentFloat)
}


func (p *percent) Numerator() Number {
	return p.numerator
}

func (p *percent) Denominator() Number {
	return p.denominator
}

func gcd(a Number, b Number) Number {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
} 