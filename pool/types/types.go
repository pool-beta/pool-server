package types

import (
	"math/rand"
)

/* ID */
type PoolID = uint64
type StreamID = uint64
type DropID = uint64
type FlowID = uint64

func NewPoolID() PoolID {
	r := rand.Uint64()
	return r
}

func NewStreamID() StreamID {
	r := rand.Uint64()
	return r
}

func NewDropID() DropID {
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

/* FlowType */
type FlowType uint8
const (
	NOFLOW FlowType = iota
	PULL
	PUSH
)

/* FlowStatus */
type FlowStatus uint8
const (
	NONE FlowStatus = iota
	INVALID
	VALID
	ACCEPTED
	DECLINED
)