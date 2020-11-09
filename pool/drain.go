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

func (d *drain) Pull(drop Drop) error {
	return fmt.Errorf("Can't pull from a drain -- drop: %v", drop)
}