package collections

// Collection is a generic type that can hold a collection of items. It does not enforce any specific ordering, or the
// ability to modify the collection, only a few simple elements.
//
// The first type parameter E determines the element type, while the second (I) determines the iterator type used
// for iterating over the elements of the collection.
type Collection[E comparable] interface {
	Iterable[E]

	Contains(E) bool
	IsEmpty() bool
	Size() uint
	ToSlice() []E
}

type MutableCollection[E comparable] interface {
	Collection[E]
	MutableIterable[E]

	Add(E)
	AddAll(Collection[E])
	Clear()
	Remove(E)
	RemoveAll(Collection[E])
	RemoveIf(Predicate[E])
	RetainAll(Collection[E])
}

type ImmutableCollection[E comparable, T any] interface {
	Collection[E]

	WithAdded(E) T
	WithAddedAll(Collection[E]) T
	WithCleared() T
	WithRemoved(E) T
	WithRemovedAll(Collection[E]) T
	WithRemovedIf(Predicate[E]) T
	WithRetainedAll(collection Collection[E]) T
}
