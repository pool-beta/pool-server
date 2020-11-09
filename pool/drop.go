package pool

import (
	. "github.com/pool-beta/pool-server/types"
)

/*
	A Drop is a transaction within the pool/stream network

	It is initialized and finalized in two distinct steps 
*/

type Drop interface {
	Finalize()
}

type drop struct {
	// Corresponding Pool
	pool Pool
	amount USDollar
}

func NewDrop(pool Pool, amount USDollar) Drop {
	return &drop{
		pool: pool,
		amount: amount,
	}
}

func (d *drop) Finalize() {
	
}
