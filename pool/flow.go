package pool

import (
	"fmt"

	. "github.com/pool-beta/pool-server/types"
	. "github.com/pool-beta/pool-server/pool/types"
)

/* 
	Flow is the interface to start/accept/decline a drop system.
	Drop should not be used outside of package pool

	Both Push and Pull have one pool that gains and another that loses.
	PushDrop adds money while PullDrop takes aways when accepted.
*/

type Flow interface {
	PullDrop() Drop
	PushDrop() Drop

	Valid()
	Invalid()

	Accept() error
	Decline() error
}

type flow struct {
	push Drop
	pull Drop
	status FlowStatus
}

func NewFlow(pool Pool, amount USDollar, flowType FlowType) (Flow, error) {
	// Create Drops
	var push, pull Drop
	switch (flowType) {
	case PULL:
		push = newDrop(pool, amount) 
		pull = newDrop(pool, amount)
	case PUSH:
		pull = newDrop(pool, amount)
		push = newDrop(pool, amount)
	default:
		return nil, fmt.Errorf("Invalid FlowType -- flowType: %v", flowType)
	}

	flow := &flow{
		push: push,
		pull: pull,
		status: NONE,
	}

	var err error
	switch (flowType) {
	case PULL:
		err = pool.Pull(flow)
	case PUSH:
		err = pool.Push(flow)
	default:
		return nil, fmt.Errorf("Invalid FlowType -- flowType: %v", flowType)
	}

	if err != nil {
		flow.status = INVALID
	} else {
		flow.status = VALID
	}
	
	return flow , nil
}

func (f *flow) Valid() {
	f.status = VALID
}

func (f *flow) Invalid() {
	f.status = INVALID
}

func (f *flow) Accept() error {
	if f.pull == nil && f.push == nil {
		return fmt.Errorf("No push or pull drop")
	}

	if f.pull != nil {
		f.pull.Discard()
	}
	if f.push != nil {
		f.push.Absorb()
	}
	return nil
}

func (f *flow) Decline() error {
	if f.pull == nil && f.push == nil {
		return fmt.Errorf("No push or pull drop")
	}

	if f.pull != nil {
		f.pull.Absorb()
	}
	if f.push != nil {
		f.push.Discard()
	}

	return nil
}

func (f *flow) PullDrop() Drop {
	return f.pull
}

func (f *flow) PushDrop() Drop {
	return f.push
}