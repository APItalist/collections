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
func (e *exampleIterator[T]) ForEachRemaining(c collections.Consumer[T]) {
	for e.HasNext() {
		c(e.Next())
	}
}

// HasNext will return true if there is a next element in the iterator.
func (e exampleIterator[T]) HasNext() bool {
	return len(e.data)-1 > e.i
}

// Next will advance the internal pointer to the next element and return the next element.
func (e *exampleIterator[T]) Next() T {
	if !e.HasNext() {
		panic(collections.ErrIndexOutOfBounds)
	}
	e.i++
	return e.data[e.i]
}

func ExampleIterator() {
	// Example_slice 1: using HasNext() and Next() to loop over the iterator.
	var iterator1 collections.Iterator[string] = &exampleIterator[string]{
		data: []string{"a", "b", "c"},
		i:    -1,
	}

	for iterator1.HasNext() {
		element := iterator1.Next()
		fmt.Println(element)
	}

	// Example_slice 2: using ForEachRemaining() to loop over the iterator.
	var iterator2 collections.Iterator[string] = &exampleIterator[string]{
		data: []string{"d", "e", "f"},
		i:    -1,
	}

	printerFunc := func(e string) {
		fmt.Println(e)
	}

	iterator2.ForEachRemaining(printerFunc)

	// Output: a
	// b
	// c
	// d
	// e
	// f
}
