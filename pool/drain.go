package pool

import (
	"fmt"
)

type Drain interface {
	// Extends Pool
	Pool
}

type drain struct {

}

func NewDrain() Drain {
	return nil
}

func (d *drain) Pull(drop Drop) error {
	return fmt.Errorf("Can't pull from a drain -- drop: %v", drop)
}