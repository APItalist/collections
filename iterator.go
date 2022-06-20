package collections

import "io"

// Iterable is an interface designating an element as being able to be looped over. The T type parameter designates
// the type of the elements in the iterator.
type Iterable[T any] interface {
	// Iterator returns an Iterator implementation that loops over a list of elements. Depending on the underlying
	// implementation, the iterator may or may not loop over the elements in a guaranteed order.
	Iterator() Iterator[T]
}

// MutableIterable is a typed interface that returns an iterator that can be updated.
type MutableIterable[T any] interface {
	// MutableIterator returns a mutable iterator to loop over the elements. It also offers the ability to remove the
	// current element from the list. Multiple concurrent iterators for the Slice may exist, but concurrent modification
	// must be locked externally.
	MutableIterator() MutableIterator[T]
}

// Iterator designates the methods needed for iterating over a specific set of items.
type Iterator[T any] interface {
	// ForEachRemaining executes the specified function for each remaining item, or until a panic occurs.
	ForEachRemaining(Consumer[T])

	// HasNext returns true if there is a next element in the iterator. The iterator pointer is not advanced.
	HasNext() bool

	// Next returns the next element in the iterator and advances the internal pointer by one. If no more elements are
	// remaining, an ErrIndexOutOfBounds error is thrown in a panic. This function is best used in a for loop in
	// conjunction with HasNext():
	//
	//     for iterator.HasNext() {
	//         element := iterator.Next()
	//         // Use element here
	//     }
	Next() T
}

// MutableIterator extends on the Iterator interface by adding methods to modify the underlying data structure.
type MutableIterator[T any] interface {
	Iterator[T]
	
	// Remove removes the current element. Throws an ErrIndexOutOfBounds in a panic if the current iterator does not
	// point to a valid element (e.g. before calling Next()).
	Remove()
}

// IteratorCloser is an iterator that must be closed for an orderly shutdown.
type IteratorCloser[T any] interface {
	Iterator[T]

	io.Closer
}
