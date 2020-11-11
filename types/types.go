package types

import (
	"fmt"
)

type PoolID = uint64
type StreamID = uint64
type UserID = uint64

/* PoolType */
type PoolType uint8
const (
	UNKNOWN PoolType = iota
	POOL
	TANK
	DRAIN
	// GROUP 
)

/* USDollar */
type USDollar uint64

func NewUSDollar(dollar int, cent int) (USDollar, error) {
	if dollar < 0 || cent < 0 {
		return USDollar(0), fmt.Errorf("invalid values for dollar or cent -- dollar: %v; cent: %v", dollar, cent)
	}
	if cent > 99 {
		return USDollar(0), fmt.Errorf("cent must be less than 99 -- cent: %v", cent)
	}

	d := USDollar(dollar) * USDollar(100)
	c := USDollar(cent)
	return d + c, nil
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

// Returns USDollar that p% of us 
func (us *USDollar) MultiplyPercent(p Percent) USDollar {
	usFloat := float64(*us)
	product := usFloat * p.Float() / float64(100)

	return USDollar(product)
}

/* Percent */
type Number uint8

type Percent interface {
	// Getter
	Numerator() Number
	Denominator() Number
	// Math
	Flip() Percent
	Add(Percent) Percent
	Multiply(Percent) Percent
	Divide(Percent) Percent
	ToOne() (Percent, bool)
	IsOne() bool
	// Out
	Float() float64
	String() string
}

type percent struct {
	numerator Number
	denominator Number
}

// TODO: Return error for divide zero (?)
func NewPercent(numerator Number, denominator Number) Percent {
	divisor := gcd(numerator, denominator)
	numerator = numerator / divisor
	denominator = denominator / divisor

	return &percent{
		numerator: numerator,
		denominator: denominator,
	}
}

func (p *percent) Flip() Percent {
	return NewPercent(p.denominator, p.numerator)
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

func (p *percent) Multiply(other Percent) Percent {
	x := p.numerator * other.Numerator()
	y := p.denominator * other.Denominator()

	divisor := gcd(x, y)
	numerator := x / divisor
	denominator := y / divisor

	return NewPercent(numerator, denominator)
}

func (p *percent) Divide(other Percent) Percent {
	return p.Multiply(other.Flip())
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

func (p *percent) Float() float64 {
	numFloat := float64(p.numerator)
	denFloat := float64(p.denominator)
	percentFloat := numFloat / denFloat * float64(100)

	return percentFloat
}

func (p *percent) String() string {
	return fmt.Sprintf("%v%%", p.Float())
}

func (p *percent) Numerator() Number {
	return p.numerator
}

func (p *percent) Denominator() Number {
	return p.denominator
}

// Returns the map of Percents normalized 
func NormalizePercents(percents map[uint64]Percent) map[uint64]Percent {
	// Add them up
	total := NewPercent(0, 1)
	for _, p := range percents {
		total = total.Add(p)
	}

	if total.IsOne() {
		return percents
	} else {
		normalized := make(map[uint64]Percent)
		for k, p := range percents {
			normalized[k] = p.Divide(total)
		}
		return normalized
	}
}

func gcd(a Number, b Number) Number {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}