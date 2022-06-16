package stream

import (
	"sync"

	"github.com/apitalist/collections"
)

type stream[T any] struct {
	// input is a channel where items can be received from upstream.
	input <-chan T
	// complete is a channel to signal upstream that processing is complete and no more items should be sent.
	complete chan struct{}
	// errorInput is a channel where upstream processors can send errors downstream.
	errorInput <-chan error
}

func (s stream[T]) AllMatch(p collections.Predicate[T]) (bool, error) {
	for {
		select {
		case item, ok := <-s.input:
			if !ok {
				return true, nil
			}
			if !p(item) {
				close(s.complete)
				return false, nil
			}
		case err, ok := <-s.errorInput:
			if !ok {
				return true, nil
			}
			return false, err
		}
	}
}

func (s stream[T]) AnyMatch(p collections.Predicate[T]) (bool, error) {
	for {
		select {
		case item, ok := <-s.input:
			if !ok {
				return false, nil
			}
			if p(item) {
				close(s.complete)
				return true, nil
			}
		case err := <-s.errorInput:
			return false, err
		}
	}
}

func (s *stream[T]) Filter(p collections.Predicate[T]) collections.Stream[T] {
	output := make(chan T)
	errorOutput := make(chan error)
	s2 := &stream[T]{
		input:      output,
		errorInput: errorOutput,
		complete:   s.complete,
	}
	go func() {
		defer func() {
			close(output)
			close(errorOutput)
		}()
		for {
			var e T
			var ok bool
			var err error
			select {
			case e, ok = <-s.input:
				if !ok {
					return
				}
			case err, ok = <-s.errorInput:
				if !ok {
					return
				}
			case <-s.complete:
				return
			}
			if err != nil {
				select {
				case errorOutput <- err:
					return
				case <-s.complete:
					return
				}
			}
			if p(e) {
				select {
				case output <- e:
				case <-s.complete:
					return
				}
			}
		}
	}()
	return s2
}

func (s stream[T]) ToSlice() ([]T, error) {
	defer close(s.complete)
	var result []T
	for {
		select {
		case item, ok := <-s.input:
			if !ok {
				return result, nil
			}
			result = append(result, item)
		case err := <-s.errorInput:
			return result, err
		}
	}
}

func (s stream[T]) FindFirst() (T, error) {
	var defaultValue T
	defer close(s.complete)
	for {
		select {
		case item, ok := <-s.input:
			if !ok {
				return defaultValue, nil
			}
			return item, nil
		case err := <-s.errorInput:
			return defaultValue, err
		}
	}
}

func (s stream[T]) FindAny() (T, error) {
	return s.FindFirst()
}

func (s stream[T]) Count() (uint, error) {
	defer close(s.complete)
	count := uint(0)
	for {
		select {
		case _, ok := <-s.input:
			if !ok {
				return count, nil
			}
			count++
		case err, ok := <-s.errorInput:
			if !ok {
				return count, nil
			}
			return count, err
		}
	}
}

func (s *stream[T]) Map(f func(T) (T, error)) collections.Stream[T] {
	output := make(chan T)
	errorOutput := make(chan error)
	s2 := &stream[T]{
		input:      output,
		errorInput: errorOutput,
		complete:   s.complete,
	}
	go func() {
		defer func() {
			close(output)
			close(errorOutput)
		}()
		for {
			var e T
			var ok bool
			var err error
			select {
			case e, ok = <-s.input:
				if !ok {
					return
				}
			case err, ok = <-s.errorInput:
				if !ok {
					return
				}
			case <-s.complete:
				return
			}
			if err != nil {
				select {
				case errorOutput <- err:
					return
				case <-s.complete:
					return
				}
			}
			newItem, err := f(e)
			if err != nil {
				select {
				case errorOutput <- err:
					return
				case <-s.complete:
					return
				}
			}
			select {
			case output <- newItem:
			case <-s.complete:
				return
			}
		}
	}()
	return s2
}

func (s *stream[T]) Iterator() collections.IteratorCloser[T] {
	return &iterator[T]{
		s.input,
		s.errorInput,
		s.complete,
		&sync.Mutex{},
		nil,
		nil,
		false,
	}
}

type iterator[T any] struct {
	input      <-chan T
	errorInput <-chan error
	complete   chan struct{}
	lock       *sync.Mutex
	lastItem   *T
	lastError  error
	finished   bool
}

func (i *iterator[T]) ForEachRemaining(c collections.Consumer[T]) error {
	for i.HasNext() {
		item, err := i.Next()
		if err != nil {
			return err
		}
		if err := c(item); err != nil {
			return err
		}
	}
	return nil
}

func (i *iterator[T]) Close() error {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.finish()
	return nil
}

func (i *iterator[T]) HasNext() bool {
	i.lock.Lock()
	defer i.lock.Unlock()
	if i.lastItem != nil {
		return true
	}
	select {
	case item, ok := <-i.input:
		if !ok {
			i.finish()
			return false
		}
		i.lastItem = &item
		return true
	case err, ok := <-i.errorInput:
		if !ok {
			i.finish()
			return false
		}
		i.lastError = err
		return false
	case <-i.complete:
		return false
	}
}

func (i *iterator[T]) finish() {
	if !i.finished {
		i.finished = true
		close(i.complete)
	}
}

func (i *iterator[T]) Next() (T, error) {
	var defaultValue T
	i.lock.Lock()
	defer i.lock.Unlock()
	if i.lastError != nil {
		return defaultValue, i.lastError
	}
	if i.lastItem != nil {
		item := *i.lastItem
		i.lastItem = nil
		return item, nil
	}
	select {
	case item, ok := <-i.input:
		if !ok {
			i.finish()
			return defaultValue, collections.ErrIndexOutOfBounds
		}
		return item, nil
	case err, ok := <-i.errorInput:
		i.finish()
		if !ok {
			return defaultValue, collections.ErrIndexOutOfBounds
		}
		i.lastError = err
		return defaultValue, err
	case <-i.complete:
		i.finish()
		return defaultValue, collections.ErrIndexOutOfBounds
	}
}
