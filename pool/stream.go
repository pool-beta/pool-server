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
	GetConfig() *streamConfig

	Pull(Drop) error
	Push(Drop) error
}

type stream struct {
	owner UserID
	// Source
	pullPool Pool
	// Destination
	pushPool Pool
	// Push/Pull Config
	config *streamConfig
}

func NewStream(owner UserID, pullPool Pool, pushPool Pool) (Stream, error) {
	if pushPool == nil || pullPool == nil {
		return nil, fmt.Errorf("Invalid pushPool or pullPool")
	}



	return &stream{
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

func (s *stream) GetConfig() *streamConfig {
	return s.config
}
 
// -------------------------------------------------------------------------------------------------

type StreamConfig interface {
	// Pull Config
	GetAllowOverdraft() bool
	SetAllowOverdraft(bool)

	GetAllowFlexibleOverdraft() bool
	SetAllowFlexibleOverdraft(bool)

	GetPercentageOverdraft() Percent
	SetPercentageOverdraft(Percent)

	GetMaxOverdraft() USDollar
	SetMaxOverdraft(USDollar)

	GetMinOverdraft() USDollar
	SetMinOverdraft(USDollar)
}

type streamConfig struct {
	// Pull Config
	allowOverdraft bool
	allowFlexibleOverdraft bool
	percentageOverdraft Percent
	maxOverdraft USDollar
	minOverdraft USDollar
}

func newStreamConfig() StreamConfig {
	return &streamConfig{
		allowOverdraft: false,
		allowFlexibleOverdraft: false,
		percentageOverdraft: NewPercent(Number(0), Number(1)),
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

func (c *streamConfig) GetPercentageOverdraft() Percent {
	return c.percentageOverdraft
}

func (c *streamConfig) SetPercentageOverdraft(percent Percent) {
	c.percentageOverdraft = percent
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