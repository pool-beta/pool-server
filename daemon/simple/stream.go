package simple

import (
	"fmt"

	ppool "github.com/pool-beta/pool-server/pool"
	. "github.com/pool-beta/pool-server/types"
)

type stream struct {
	stream ppool.Stream
}

func newStream() (Stream, error) {
	return nil, nil
}


func (s *stream) EnableOverDraft() error {
	if s.stream == nil {
		return fmt.Errorf("stream cannot be nil")
	}

	s.stream.SetAllowOverdraft(true)
	return nil
}

func (s *stream) DisableOverDraft() error {
	if s.stream == nil {
		return fmt.Errorf("stream cannot be nil")
	}

	s.stream.SetAllowOverdraft(false)
	return nil
}

func (s *stream) EnableFlexibleOverdraft() error {
	if s.stream == nil {
		return fmt.Errorf("stream cannot be nil")
	}

	s.stream.SetAllowFlexibleOverdraft(true)
	return nil
}

func (s *stream) DisableFlexibleOverdraft() error {
	if s.stream == nil {
		return fmt.Errorf("stream cannot be nil")
	}

	s.stream.SetAllowFlexibleOverdraft(false)
	return nil
}

func (s *stream) SetPercentOverdraft(percent Percent) error {
	if s.stream == nil {
		return fmt.Errorf("stream cannot be nil")
	}

	s.stream.SetPercentOverdraft(percent)
	return nil
}

func (s *stream) SetMaxOverdraft(amount USDollar) error {
	if s.stream == nil {
		return fmt.Errorf("stream cannot be nil")
	}

	s.stream.SetMaxOverdraft(amount)
	return nil
}
