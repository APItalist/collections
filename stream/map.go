// Package stream provides a channel-based stream processor, which processes individual steps in separate goroutines.
// However, the individual step executions are not parallelized.
package stream

import (
	"errors"

	"github.com/apitalist/collections"
	"github.com/apitalist/lang"
)

// Map takes an input stream and a mapping function, then uses the mapping function to create an output stream.
// This is only required until Golang gets support for receiver generic types..
func Map[TInput, TOutput any](
	input collections.Stream[TInput],
	mapper func(TInput) TOutput,
) collections.Stream[TOutput] {
	output := make(chan TOutput)
	errorOutput := make(chan error)
	complete := make(chan struct{})
	s2 := &stream[TOutput]{
		input:      output,
		errorInput: errorOutput,
		complete:   complete,
	}
	iterator := input.Iterator()
	go func() {
		defer func() {
			close(output)
			close(errorOutput)
		}()
		for {
			var e TInput
			err := lang.Safe(
				func() {
					e = iterator.Next()
				},
			)
			if err != nil {
				if !errors.Is(err, collections.ErrIndexOutOfBounds) {
					select {
					case errorOutput <- err:
					case <-complete:
						return
					}
				}
				return
			}
			var item TOutput
			err = lang.Safe(
				func() {
					item = mapper(e)
				},
			)
			if err != nil {
				select {
				case errorOutput <- err:
				case <-complete:
				}
				return
			}
			select {
			case output <- item:
			case <-complete:
				return
			}
		}
	}()
	return s2
}
