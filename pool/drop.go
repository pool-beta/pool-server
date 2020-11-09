package pool

import (
	. "github.com/pool-beta/pool-server/types"
)

/*
	A Drop is a transaction within the pool/stream network

	It is initialized and finalized in two distinct steps

	Notes:
		- Should not need a lock since a tree of drops (or "flow") should not happen concurrently
			- Possbily when we start merging drops
*/

type Drop interface {
	// Getters
	Amount() USDollar

	// Accept the drop; follow thru
	Absorb()
	// Reject the drop; return the previous state
	Reject()
	// Add a drop to the list of droplets
	AddDroplet(Drop)
}

type drop struct {
	// Corresponding Pool
	pool Pool
	amount USDollar

	droplets []Drop
}

func NewDrop(pool Pool, amount USDollar) Drop {
	droplets := make([]Drop, 0)

	return &drop{
		pool: pool,
		amount: amount,
		droplets: droplets,
	}
}

func (d *drop) AddDroplet(drop Drop) {
	d.droplets = append(d.droplets, drop)
}

func (d *drop) Absorb() {
	// Absorb Droplets first
	for _, dlets := range d.droplets {
		dlets.Absorb()
	}

	// Do nothing (?); maybe log it
}

func (d *drop) Reject() {
	// Reject Droplets first
	for _, dlets := range d.droplets {
		dlets.Reject()
	}

	// Return to previous state

}

// Getters
func (d *drop) Amount() USDollar {
	return d.amount
}
