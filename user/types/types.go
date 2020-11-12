package types

import (
	"math/rand"
)

type UserID = uint64

func NewUserID() UserID {
	r := rand.Uint64()
	return r
}

type UserName = string