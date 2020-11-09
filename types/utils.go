package types

import (
	"math/rand"
)

func NewPoolID() PoolID {
	r := rand.Uint64()
	return r
}

func NewUserID() UserID {
	r := rand.Uint64()
	return r
}
