package pool

import (
	"fmt"

	. "github.com/pool-beta/pool-server/types"
	. "github.com/pool-beta/pool-server/user/types"
	. "github.com/pool-beta/pool-server/pool/types"
)

/* 
	Stream implements a connection between pools

	Note:
		- Drops should NEVER be started at a stream; instead, they should be started on the pool they correspond to
*/

type Stream interface {
	StreamConfig
	StreamID() StreamID
	Owner() UserID
	Pull(amount USDollar) (Drop, error)
	Push(amount USDollar) (Drop, error)
}

type stream struct {
	*streamConfig
	streamID StreamID
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
	streamID := NewStreamID()

	stream := &stream{
		streamConfig: config,
		streamID: streamID,
		owner: owner,
		pullPool: pullPool,
		pushPool: pushPool,
	}

	// TODO: Add the stream to pool
	pullPool.AddPuller(stream)
	pushPool.AddPusher(stream)

	return stream, nil
}

func (s *stream) Pull(amount USDollar) (Drop, error) {
	drop := newDrop(s.pullPool, amount)
	return drop, s.pullPool.PullDrop(drop, true)
}

func (s *stream) Push(amount USDollar) (Drop, error) {
	_ = newDrop(s.pushPool, amount)
	return nil, nil
}

func (s *stream) StreamID() StreamID {
	return s.streamID
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

	// Push Config
	GetAllowPush() bool
	SetAllowPush(bool)

	GetPercentPush() Percent
	SetPercentPush(Percent)

	// GetMaxPush() USDollar
	// SetMaxPush(USDollar)

}

type streamConfig struct {
	// Pull Config
	allowOverdraft bool
	allowFlexibleOverdraft bool
	percentOverdraft Percent
	maxOverdraft USDollar
	minOverdraft USDollar // TODO: Not Fully Supported

	// Push Config
	allowPush bool
	percentPush Percent
}

func newStreamConfig() *streamConfig {
	return &streamConfig{
		allowOverdraft: false,
		allowFlexibleOverdraft: false,
		percentOverdraft: NewPercent(0, 1),
		maxOverdraft: USDollar(0),
		minOverdraft: USDollar(0),

		allowPush: false,
		percentPush: NewPercent(0, 1),
	}
}

// Pull Configs

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

// Push Configs

func (c *streamConfig) GetAllowPush() bool {
	return c.allowPush
}

func (c *streamConfig) SetAllowPush(value bool) {
	c.allowPush = value
}

func (c *streamConfig) GetPercentPush() Percent {
	return c.percentPush
}

func (c *streamConfig) SetPercentPush(percent Percent) {
	c.percentPush = percent
}