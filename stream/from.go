package stream

import "github.com/apitalist/collections"

func FromCollection[E comparable](c collections.Collection[E]) collections.Stream[E] {
	input := make(chan E)
	complete := make(chan struct{})
	errorInput := make(chan error)
	s := &stream[E]{
		input:      input,
		complete:   complete,
		errorInput: errorInput,
	}
	iterator := c.Iterator()
	go func() {
		defer func() {
			close(input)
			close(errorInput)
		}()
		for iterator.HasNext() {
			e, err := iterator.Next()
			if err != nil {
				select {
				case errorInput <- err:
				case <-complete:
					return
				}
			}
			select {
			case input <- e:
			case <-complete:
				break
			}
		}
	}()
	return s
}
