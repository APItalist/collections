package collections

// List is an structure where elements are stored in-order and elements may repeat. This interface offers functions
// to access a specific item of that list. The underlying implementation, for example slice from the slice package,
// determines the execution speed of these operations.
//
//
// This interface has two subinterfaces: MutableList, which adds the operations to change the list, and ImmutableList,
// which adds methods to create modified copies of this list.
type List[E comparable, T any] interface {
	Collection[E]

	Get(index uint) (E, error)
	IndexOf(E) (uint, error)
	LastIndexOf(E) (uint, error)
	SubList(from, to uint) (T, error)
}

// MutableList is a List type that can be directly modified. However, depending on the implementation, the list may not
// be safe to concurrently modify. If concurrency is desired, please select an implementation that is safe for
// concurrent access or consider switching to ImmutableList, which is concurrency safe by design.
type MutableList[E comparable] interface {
	List[E, MutableList[E]]
	MutableCollection[E]

	AddAt(index uint, element E) error
	Set(index uint, element E) error
	Sort(Comparator[E])

	// RemoveAt removes the element at the specified index. If the specified index does not exist a
	// collections.ErrIndexOutOfBounds is returned.
	RemoveAt(index uint) error
}

// ImmutableList is a List type that cannot be modified, but has helper functions to create a copy of the current
// List with the modified values in it. This is most useful when lists need to be modified and concurrently accessed.
type ImmutableList[E comparable] interface {
	List[E, ImmutableList[E]]
	ImmutableCollection[E, ImmutableList[E]]

	WithAddedAt(index uint, element E) (ImmutableList[E], error)
	WithSet(index uint, element E) (ImmutableList[E], error)
	WithSorted(Comparator[E]) ImmutableList[E]
	WithRemovedAt(index uint) (ImmutableList[E], error)
}
