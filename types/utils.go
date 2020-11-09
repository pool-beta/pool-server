package types

import (
	"math/rand"
)

func NewPoolID() PoolID {
	r := rand.Uint64()
	return r
}