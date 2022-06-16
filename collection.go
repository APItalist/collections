package collections

// Collection is a generic type that can hold a collection of items. It does not enforce any specific ordering, or the
// ability to modify the collection, only a few simple elements.
//
// The first type parameter E determines the element type, while the second (I) determines the iterator type used
// for iterating over the elements of the collection.
type Collection[E comparable] interface {
	Iterable[E]

	// Contains returns true, if the collection contains the specified element.
	Contains(E) bool

	// IsEmpty returns true, if there are no elements in the current collection.
	IsEmpty() bool

	// Size returns the number of items in the current collection.
	Size() uint

	// ToSlice returns all elements in the current collection in a slice.
	ToSlice() []E
}

// MutableCollection is a collection variant that allows directly changing the elements of the current collection.
type MutableCollection[E comparable] interface {
	Collection[E]
	MutableIterable[E]

	// Add adds the specified element to the collection.
	Add(E)

	// AddAll adds all items from the passed collection to the current collection.
	AddAll(Collection[E])

	// Clear removes all items from the current collection.
	Clear()

	// Remove removes all instances of the specified element from the current list.
	Remove(E)

	// RemoveAll removes all elements that are in the passed collection from the current collection.
	RemoveAll(Collection[E])

	// RemoveIf calls the passed predicate function for each element in the collection. If the predicate returns
	// true, the element is removed from the current collection.
	RemoveIf(Predicate[E])

	// RetainAll removes all elements from the current collection that are not present in the passed collection.
	RetainAll(Collection[E])
}

// ImmutableCollection is a collection that has helper functions for creating changed copies of the collection.
type ImmutableCollection[E comparable, T any] interface {
	Collection[E]

	// WithAdded creates a new copy of the current collection with the specified element added.
	WithAdded(E) T

	// WithAddedAll creates a new copy of the current collection with all elements from the passed collection added.
	// The current collection is not changed.
	WithAddedAll(Collection[E]) T

	// WithCleared creates an empty copy of the current collection.
	WithCleared() T

	// WithRemoved creates a copy of the current collection, with all elements that match the passed element removed.
	// The current collection is not changed.
	WithRemoved(E) T

	// WithRemovedAll creates a copy of the current collection, with all elements matching the elements in the
	// passed collection removed. The current collection is not changed.
	WithRemovedAll(Collection[E]) T

	// WithRemovedIf calles the specified predicate for each element in the current collection and copies them into a
	// new collection if the predicate returns false. The current collection is not changed.
	WithRemovedIf(Predicate[E]) T

	// WithRetainedAll creates a copy of the elements of the current collection that are also present in the collection
	// passed in the parameter.
	WithRetainedAll(collection Collection[E]) T
}
