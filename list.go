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

	// Get returns the item located at the specified index, starting at 0 for the first item. If the list contains the
	// item, the item will be returned. If the list does not contain the item, an ErrIndexOutOfBounds error will be
	// returned.
	Get(index uint) (E, error)

	// IndexOf returns the index of the first item in the list matching the specified element, starting at index 0. If
	// no item is found, an ErrElementNotFound is returned.
	IndexOf(E) (uint, error)

	// LastIndexOf returns the index of the last item in the list matching the specified element, starting at index 0.
	// If no item is found, an ErrElementNotFound is returned.
	LastIndexOf(E) (uint, error)

	// SubList creates a list from a part from the current list, starting at the from parameter (inclusive) up until
	// the to parameter (exclusive). If the specified bounds are invalid (from is larger than to, or to is larger than
	// the list length), an ErrIndexOutOfBounds is returned.
	SubList(from, to uint) (T, error)
}

// MutableList is a List type that can be directly modified. However, depending on the implementation, the list may not
// be safe to concurrently modify. If concurrency is desired, please select an implementation that is safe for
// concurrent access or consider switching to ImmutableList, which is concurrency safe by design.
type MutableList[E comparable] interface {
	List[E, MutableList[E]]
	MutableCollection[E]

	// AddAt inserts an item into the current list at the specified index, shifting list elements coming after the
	// specified index backwards. If the specified index is too large, an ErrIndexOutOfBounds is returned.
	AddAt(index uint, element E) error

	// Set replaces the item at the indicated index with the supplied parameter. If the index is larger than the list
	// end, an ErrIndexOutOfBounds is returned.
	Set(index uint, element E) error

	// Sort sorts the current list with the help of the passed comparator function. The comparator function will be
	// called several times, each time with two values. If the comparator returns a value smaller than zero, the two
	// items will be swapped in the list.
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

	// WithAddedAt creates a new copy of the current list, with the specified element added at the index, shifting
	// other elements back. The current list remains unchanged in the process. If the specified index is after the end
	// of the list, an ErrIndexOutOfBounds is returned.
	WithAddedAt(index uint, element E) (ImmutableList[E], error)

	// WithSet creates a new copy of the current list, with the element at the specified index replaced with the
	// specified value. The current list remains unchanged in the process. If the specified index is after the end
	// of the list, an ErrIndexOutOfBounds is returned.
	WithSet(index uint, value E) (ImmutableList[E], error)

	// WithSorted returns a new list with the elements from the current list sorted according to the passed comparator
	// function. The current list remains unchanged in the process. The comparator function will be called several
	// times, each time with two values. If the comparator returns a value smaller than zero, the two items will be
	// swapped in the list.
	WithSorted(Comparator[E]) ImmutableList[E]

	// WithRemovedAt returns a new list with the item at the specified index removed. The current list remains unchanged
	// in the process. If the specified index is after the end of the index, an ErrIndexOutOfBounds is returned.
	WithRemovedAt(index uint) (ImmutableList[E], error)
}
