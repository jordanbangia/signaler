package signaler

import (
	"context"
	"sync"
)

type Signaler interface {
	Subscribe() <-chan struct{}
	Trigger()
}

type signaler struct {
	lock sync.Mutex

	ctx    context.Context
	cancel context.CancelFunc
}

func (s *signaler) Subscribe() <-chan struct{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.ctx == nil {
		s.ctx, s.cancel = context.WithCancel(context.Background())
	}
	return s.ctx.Done()
}

func (s *signaler) Trigger() {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.cancel == nil {
		return
	}

	s.cancel()

	s.ctx = nil
	s.cancel = nil
}
