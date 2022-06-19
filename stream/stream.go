package stream

import (
	"sync"

	"github.com/apitalist/collections"
	"github.com/apitalist/lang"
)

type stream[T any] struct {
	// input is a channel where items can be received from upstream.
	input <-chan T
	// complete is a channel to signal upstream that processing is complete and no more items should be sent.
	complete chan struct{}
	// errorInput is a channel where upstream processors can send errors downstream.
	errorInput <-chan error
}

func (s stream[T]) AllMatch(p collections.Predicate[T]) bool {
	for {
		select {
		case item, ok := <-s.input:
			if !ok {
				return true
			}
			if !p(item) {
				close(s.complete)
				return false
			}
		case err, ok := <-s.errorInput:
			if !ok {
				return true
			}
			panic(err)
		}
	}
}

func (s stream[T]) AnyMatch(p collections.Predicate[T]) bool {
	for {
		select {
		case item, ok := <-s.input:
			if !ok {
				return false
			}
			if p(item) {
				close(s.complete)
				return true
			}
		case err, ok := <-s.errorInput:
			if !ok {
				return true
			}
			panic(err)
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

func (s stream[T]) ToSlice() []T {
	defer close(s.complete)
	var result []T
	for {
		select {
		case item, ok := <-s.input:
			if !ok {
				return result
			}
			result = append(result, item)
		case err, ok := <-s.errorInput:
			if !ok {
				return result
			}
			panic(err)
		}
	}
}

func (s stream[T]) FindFirst() T {
	defer close(s.complete)
	for {
		select {
		case item, ok := <-s.input:
			if !ok {
				panic(collections.ErrElementNotFound)
			}
			return item
		case err, ok := <-s.errorInput:
			// TODO possible race condition here:
			if !ok {
				panic(collections.ErrElementNotFound)
			}
			panic(err)
		}
	}
}

func (s stream[T]) FindAny() T {
	return s.FindFirst()
}

func (s stream[T]) Count() uint {
	defer close(s.complete)
	count := uint(0)
	for {
		select {
		case _, ok := <-s.input:
			if !ok {
				return count
			}
			count++
		case err, ok := <-s.errorInput:
			if !ok {
				return count
			}
			panic(err)
		}
	}
}

func (s *stream[T]) Map(f func(T) T) collections.Stream[T] {
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
			var newItem T
			err = lang.Safe(
				func() {
					newItem = f(e)
				},
			)
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

func (i *iterator[T]) ForEachRemaining(c collections.Consumer[T]) {
	for i.HasNext() {
		c(i.Next())
	}
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

func (i *iterator[T]) Next() T {
	i.lock.Lock()
	defer i.lock.Unlock()
	if i.lastError != nil {
		panic(i.lastError)
	}
	if i.lastItem != nil {
		item := *i.lastItem
		i.lastItem = nil
		return item
	}
	select {
	case item, ok := <-i.input:
		if !ok {
			i.finish()
			panic(collections.ErrIndexOutOfBounds)
		}
		return item
	case err, ok := <-i.errorInput:
		i.finish()
		if !ok {
			panic(collections.ErrIndexOutOfBounds)
		}
		i.lastError = err
		panic(err)
	case <-i.complete:
		i.finish()
		panic(collections.ErrIndexOutOfBounds)
	}
}
