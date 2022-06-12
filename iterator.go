package collections

type Iterator[T any] interface {
    // ForEachRemaining executes the specified function for each remaining item, or until an error occurs. If an
    // error occurs, the error is returned and no further elements are processed.
    ForEachRemaining(Consumer[T]) error
    // HasNext returns true if there is a next element in the iterator. The iterator pointer is not advanced.
    HasNext() bool
    // Next returns the next element in the iterator, or an error if no more elements exist in the iterator.
    Next() (T, error)
}

type MutableIterator[T any] interface {
    Iterator[T]

    // Remove removes the current element. Throws an ErrIndexOutOfBounds if the current iterator does not point to a
    // valid element.
    Remove() error
}
