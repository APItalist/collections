package collections

// Stream is an interface that describes a process for streaming data, one by one. The functions in this interface can
// be used to add more processing elements to a stream.
//
// Each stream must be terminated by a terminal function, such as AllMatch, AnyMatch, ToSlice, etc in order to ensure
// no resources are left dangling. Individual stream items should also not be used more than once, which can lead to
// unpredictable behavior.
type Stream[T any] interface {
	// AllMatch returns true if the predicate returns true for all items in the stream.
	//
	// This is a terminal element in the stream.
	AllMatch(Predicate[T]) (bool, error)

	// AnyMatch returns true if the predicate returns true for any item in the stream.
	//
	// This is a terminal element in the stream.
	AnyMatch(Predicate[T]) (bool, error)

	// Filter creates a stream with the items where the predicate returned true.
	Filter(Predicate[T]) Stream[T]

	// ToSlice gathers all items in the stream into a slice.
	//
	// This is a terminal element in the stream.
	ToSlice() ([]T, error)

	// FindFirst returns the first item in the stream. If no item is found an ErrElementNotFound error is returned.
	//
	// This is a terminal element in the stream.
	FindFirst() (T, error)

	// FindAny returns any element from the stream. If no item is found an ErrElementNotFound error is returned.
	//
	// This is a terminal element in the stream.
	FindAny() (T, error)

	// Count returns the number of items in the stream. It returns an error if an upstream element passed an error.
	//
	// This is a terminal element in the stream.
	Count() (uint, error)

	// Map applies a mapper function to all stream elements. If you require a type conversion, please use the Map()
	// function without a receiver.
	Map(func(T) (T, error)) Stream[T]

	// Iterator returns an iterator that loops over the stream. Please note that Close() must be called on the iterator
	// to properly close the stream.
	Iterator() IteratorCloser[T]
}
