package collections_test

import (
	"fmt"

	"github.com/apitalist/collections"
)

// exampleIterator is an example implementation of an iterator. The struct should always be initialized with i = -1
type exampleIterator[T any] struct {
	data []T
	i    int
}

// ForEachRemaining will loop over all remaining elements of the iterator and execute the consumer function.
func (e *exampleIterator[T]) ForEachRemaining(c collections.Consumer[T]) error {
	for e.HasNext() {
		element, err := e.Next()
		if err != nil {
			panic(err)
		}
		if err := c(element); err != nil {
			return err
		}
	}
	return nil
}

// HasNext will return true if there is a next element in the iterator.
func (e exampleIterator[T]) HasNext() bool {
	return len(e.data)-1 > e.i
}

// Next will advance the internal pointer to the next element and return the next element.
func (e *exampleIterator[T]) Next() (T, error) {
	var defaultValue T
	if !e.HasNext() {
		return defaultValue, collections.ErrIndexOutOfBounds
	}
	e.i++
	return e.data[e.i], nil
}

func ExampleIterator() {
	// Example_slice 1: using HasNext() and Next() to loop over the iterator.
	var iterator1 collections.Iterator[string] = &exampleIterator[string]{
		data: []string{"a", "b", "c"},
		i:    -1,
	}

	for iterator1.HasNext() {
		element, err := iterator1.Next()
		if err != nil {
			// This should never happen because we used HasNext. The only way this can happen if the underlying data
			// structure has changed during iteration and the iterator does not implement concurrent access with locks.
			// The most appropriate way to handle this is to panic. Consider using the lang.Must2() function in this
			// case.
			panic(err)
		}
		fmt.Println(element)
	}

	// Example_slice 2: using ForEachRemaining() to loop over the iterator.
	var iterator2 collections.Iterator[string] = &exampleIterator[string]{
		data: []string{"d", "e", "f"},
		i:    -1,
	}

	printerFunc := func(e string) error {
		fmt.Println(e)
		return nil
	}

	err := iterator2.ForEachRemaining(printerFunc)
	if err != nil {
		// We never return an error in printerFunc, so ForEachRemaining will also not return an error. This must be
		// a race condition if the underlying data structure has changed during execution.
		panic(err)
	}

	// Output: a
	// b
	// c
	// d
	// e
	// f
}
