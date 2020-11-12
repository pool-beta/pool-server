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
	// Absorb the amount of drop
	Absorb()
	// Discard the amount of drop
	Discard()
	// Adds the amount withheld from pool
	AddWithheld(USDollar)
	// Add a drop to the list of droplets
	AddDroplet(Drop)
}

type drop struct {
	// Corresponding Pool
	pool Pool
	amount USDollar

	withheld USDollar
	droplets []Drop
}

func newDrop(pool Pool, amount USDollar) Drop {
	droplets := make([]Drop, 0)

	return &drop{
		pool: pool,
		amount: amount,
		droplets: droplets,
	}
}

func (d *drop) AddWithheld(amount USDollar) {
	d.withheld = amount
}

func (d *drop) AddDroplet(drop Drop) {
	d.droplets = append(d.droplets, drop)
}

// Add the drop amount into the pool
func (d *drop) Absorb() {
	// Absorb Droplets first
	for _, dlets := range d.droplets {
		dlets.Absorb()
	}

	d.pool.Fund(d.withheld)
}

// Discard the drop
func (d *drop) Discard() {
	// Reject Droplets first
	for _, dlets := range d.droplets {
		dlets.Discard()
	}

	// Discard so do nothing; maybe log
}

// Getters
func (d *drop) Amount() USDollar {
	return d.amount
}
