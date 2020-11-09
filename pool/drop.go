package pool

import (

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
}

func NewDrop() Drop {
	return nil
}
