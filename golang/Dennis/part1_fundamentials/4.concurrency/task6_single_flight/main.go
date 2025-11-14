package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type call struct {
	err   error
	value any
	done  chan struct{}
}

type SingleFlight struct {
	mutex sync.Mutex
	calls map[string]*call
}

func NewSingleFlight() *SingleFlight {
	return &SingleFlight{
		calls: make(map[string]*call),
	}
}

func (s *SingleFlight) Do(
	ctx context.Context, 
	key string, 
	action func(context.Context) (any, error),
	) (any, error) {
    
	s.mutex.Lock()
	if call, found := s.calls[key]; found {
		s.mutex.Unlock()
		return s.wait(ctx, call)
	}

	call := &call{
		done: make(chan struct{}),
	}

	s.calls[key] = call
	s.mutex.Unlock()

	go func() {
		defer func() {
			if v := recover(); v != nil {
				call.err = errors.New("error from single flight")
			}

			close(call.done)

			s.mutex.Lock()
			delete(s.calls, key)
			s.mutex.Unlock()
		}()

		call.value, call.err = action(ctx)
	}()

	return s.wait(ctx, call)
}

func (s *SingleFlight) wait(
	ctx context.Context, 
	call *call) (any, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-call.done:
		return call.value, call.err
	}
}

func main() {
	fmt.Println("Go, go, go 2025")
}
