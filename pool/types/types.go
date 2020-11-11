package types

import (
	"math/rand"
)

/* ID */
type PoolID = uint64
type StreamID = uint64

func NewPoolID() PoolID {
	r := rand.Uint64()
	return r
}

func NewStreamID() StreamID {
	r := rand.Uint64()
	return r
}

/* PoolType */
type PoolType uint8
const (
	UNKNOWN PoolType = iota
	POOL
	TANK
	DRAIN
	// GROUP 
)
