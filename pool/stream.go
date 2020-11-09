package pool

import (
	"fmt"

	. "github.com/pool-beta/pool-server/types"
)

/* 
	Stream implements a connection between pools

	Note:
		- Drops should NEVER be started at a stream; instead, they should be started on the pool they correspond to
*/

type Stream interface {
	StreamConfig
	Owner() UserID
	Pull(Drop) error
	Push(Drop) error
}

type stream struct {
	*streamConfig
	
	owner UserID
	// Source
	pullPool Pool
	// Destination
	pushPool Pool
	// Push/Pull Config
}

func NewStream(owner UserID, pullPool Pool, pushPool Pool) (Stream, error) {
	if pushPool == nil || pullPool == nil {
		return nil, fmt.Errorf("Invalid pushPool or pullPool")
	}

	config := newStreamConfig()


	return &stream{
		streamConfig: config,
		owner: owner,
		pullPool: pullPool,
		pushPool: pushPool,
	}, nil
}

func (s *stream) Pull(drop Drop) error {
	return s.pullPool.Pull(drop)
}

func (s *stream) Push(drop Drop) error {
	return nil
}

func (s *stream) Owner() UserID {
	return s.owner
}

// -------------------------------------------------------------------------------------------------

type StreamConfig interface {
	// Pull Config
	GetAllowOverdraft() bool
	SetAllowOverdraft(bool)

	GetAllowFlexibleOverdraft() bool
	SetAllowFlexibleOverdraft(bool)

	GetPercentOverdraft() Percent
	SetPercentOverdraft(Percent)

	GetMaxOverdraft() USDollar
	SetMaxOverdraft(USDollar)

	GetMinOverdraft() USDollar
	SetMinOverdraft(USDollar)
}

type streamConfig struct {
	// Pull Config
	allowOverdraft bool
	allowFlexibleOverdraft bool
	percentOverdraft Percent
	maxOverdraft USDollar
	minOverdraft USDollar // TODO: Not Fully Supported
}

func newStreamConfig() *streamConfig {
	return &streamConfig{
		allowOverdraft: false,
		allowFlexibleOverdraft: false,
		percentOverdraft: NewPercent(0, 1),
		maxOverdraft: USDollar(0),
		minOverdraft: USDollar(0),		
	}
}

func (c *streamConfig) GetAllowOverdraft() bool {
	return c.allowOverdraft
}

func (c *streamConfig) SetAllowOverdraft(value bool) {
	c.allowOverdraft = value
}

func (c *streamConfig) GetAllowFlexibleOverdraft() bool {
	return c.allowFlexibleOverdraft
}

func (c *streamConfig) SetAllowFlexibleOverdraft(value bool) {
	c.allowFlexibleOverdraft = value
}

func (c *streamConfig) GetPercentOverdraft() Percent {
	return c.percentOverdraft
}

func (c *streamConfig) SetPercentOverdraft(percent Percent) {
	c.percentOverdraft = percent
}

func (c *streamConfig) GetMaxOverdraft() USDollar {
	return c.maxOverdraft
}

func (c *streamConfig) SetMaxOverdraft(max USDollar) {
	c.maxOverdraft = max
}

func (c *streamConfig) GetMinOverdraft() USDollar {
	return c.minOverdraft
}

func (c *streamConfig) SetMinOverdraft(min USDollar) {
	c.minOverdraft = min
}