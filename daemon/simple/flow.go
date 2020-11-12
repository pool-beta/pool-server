package simple

import (
	"math/rand"
	. "github.com/pool-beta/pool-server/pool/types"
)

/*
	flow implements simple Flow and it is the interface to manipulate drops.
	
	Since a drop on a pool may create a network of drops, flow allows for simple
	interface for working with them.

	**Do not interface directly with drops**
*/

type flow struct {
	id FlowID
}

func NewFlowID() FlowID {
	r := rand.Uint64()
	return r
}

func newFlow() {

}