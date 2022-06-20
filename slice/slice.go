// Package slice offers a go slice-backed list implementation that is mutable (can be changed in-place). The List API
// gives a rich set of features and supports code completion at minimal overhead. However, the Slice implementation is
// not concurrency-safe, parallel modifications of the Slice should be avoided by using locks.
package slice

import (
	"fmt"
	"sort"
	"strings"

	"github.com/apitalist/collections"
	"github.com/apitalist/collections/stream"
)

// New creates a new slice-backed list, optionally filled with the specified elements. Slice lists are not
// concurrency-safe, so locks should be applied if concurrent list operations are needed. Alternatively, immutable
// lists can be used for concurrent access.
func New[E comparable](elements ...E) *Slice[E] {
	data := make([]E, len(elements))
	copy(data, elements)
	result := make(Slice[E], len(elements))
	copy(result, elements)
	return &result
}

// NewFromSlice converts an already existing go slice into a Slice pointer.
func NewFromSlice[E comparable](existingSlice []E) *Slice[E] {
	return (*Slice[E])(&existingSlice)
}

// Slice is a slice-backed implementation from the MutableList interface.
//
// In order to guarantee this type works correctly, you should always use it as a pointer.
//
// Correct:
//
//     var mySlice *slice.Slice
//
// Incorrect:
//
//     var mySlice slice.Slice
//
// You can can create this Slice using the New() function:
//
//     myData := slice.New[string]()
//
// You can also pass data to the New() function, which eliminates the need for a type specification:
//
//     myData := slice.New("a", "b", "c")
//
// You can also convert an existing slice:
//
//     mySlice := []string{"a", "b", "c"}
//     myData := slice.NewFromSlice(mySlice)
//
// Or, you can alternatively convert it manually:
//
//     myData := []string{"a", "b", "c"}
//     mySlice := *slice.Slice(&myData)
type Slice[E comparable] []E

// Stream creates a processing stream from the current slice.
func (s *Slice[E]) Stream() collections.Stream[E] {
	return stream.FromCollection[E](s)
}

// RemoveAt removes the element at the specified index. If the specified index does not exist a
// collections.ErrIndexOutOfBounds is thrown in a panic.
func (s *Slice[E]) RemoveAt(index uint) collections.MutableList[E] {
	if index >= uint(len(*s)) {
		panic(collections.ErrIndexOutOfBounds)
	}
	*s = append((*s)[:index], (*s)[index+1:]...)
	return s
}

// Iterator returns an iterator to loop over the elements. Multiple concurrent iterators for the Slice may exist, but
// concurrent modification must be locked externally.
func (s *Slice[E]) Iterator() collections.Iterator[E] {
	return &Iterator[E]{
		s,
		-1,
	}
}

// MutableIterator returns a mutable iterator to loop over the elements. It also offers the ability to remove the
// current element from the list. Multiple concurrent iterators for the Slice may exist, but concurrent modification
// must be locked externally.
func (s *Slice[E]) MutableIterator() collections.MutableIterator[E] {
	return &Iterator[E]{
		s,
		-1,
	}
}

// IsEmpty returns true if the current Slice is empty.
func (s Slice[E]) IsEmpty() bool {
	return len(s) == 0
}

// Size returns the number of elements in the slice.
func (s Slice[E]) Size() uint {
	return uint(len(s))
}

// ToSlice returns the underlying raw slice. Modifications to this underlying slice will translate to the current slice.
func (s Slice[E]) ToSlice() []E {
	return s
}

// Contains will return true if the specified element is contained within the slice.
func (s Slice[E]) Contains(e E) bool {
	for _, elem := range s {
		if elem == e {
			return true
		}
	}
	return false
}

// Get will return the element at the specified index. If the index is larger than the number of elements, a
// collections.ErrIndexOutOfBounds is thrown in a panic.
func (s Slice[E]) Get(index uint) E {
	if index >= uint(len(s)) {
		panic(collections.ErrIndexOutOfBounds)
	}
	return s[index]
}

// IndexOf returns the index of the first element that matches the specified element. If no element is found,
// a collections.ErrElementNotFound is thrown in a panic.
func (s Slice[E]) IndexOf(e E) uint {
	for i, elem := range s {
		if elem == e {
			return uint(i)
		}
	}
	panic(collections.ErrElementNotFound)
}

// LastIndexOf returns the index of the last element that matches the specified element. If no element is found,
// a collections.ErrElementNotFound is thrown in a panic.
func (s Slice[E]) LastIndexOf(e E) uint {
	for i := len(s) - 1; i >= 0; i-- {
		elem := s[i]
		if elem == e {
			return uint(i)
		}
	}
	panic(collections.ErrElementNotFound)
}

// SubList will return a part of the current Slice. If the specified bounds are invalid, a
// collections.ErrIndexOutOfBounds is thrown in a panic.
func (s Slice[E]) SubList(from, to uint) collections.MutableList[E] {
	if from > to {
		panic(collections.ErrIndexOutOfBounds)
	}
	if to >= uint(len(s)) {
		panic(collections.ErrIndexOutOfBounds)
	}
	subSlice := s[from:to]
	newSlice := make([]E, len(subSlice))
	copy(newSlice, subSlice)
	return NewFromSlice(newSlice)
}

// Add adds a new element to the slice.
func (s *Slice[E]) Add(e E) {
	*s = append(*s, e)
}

// AddAll adds all elements from the passed collection to the current slice.
func (s *Slice[E]) AddAll(c collections.Collection[E]) {
	iterator := c.Iterator()
	for iterator.HasNext() {
		e := iterator.Next()
		s.Add(e)
	}
}

func (s *Slice[E]) Clear() {
	*s = nil
}

func (s *Slice[E]) Remove(e E) {
	for i, entry := range *s {
		if entry == e {
			*s = append((*s)[:i], (*s)[i+1:]...)
		}
	}
}

func (s *Slice[E]) RemoveAll(c collections.Collection[E]) {
	iterator := c.Iterator()
	for iterator.HasNext() {
		e := iterator.Next()
		s.Remove(e)
	}
}

func (s *Slice[E]) RemoveIf(p collections.Predicate[E]) {
	tmpSlice := (*s)[:0]
	for _, e := range *s {
		if !p(e) {
			tmpSlice = append(tmpSlice, e)
		}
	}
	*s = tmpSlice
}

func (s *Slice[E]) RetainAll(c collections.Collection[E]) {
	s.RemoveIf(collections.Predicate[E](c.Contains).Negate())
}

func (s *Slice[E]) AddAt(index uint, element E) collections.MutableList[E] {
	if index > uint(len(*s)) {
		panic(collections.ErrIndexOutOfBounds)
	}
	if index == uint(len(*s)) {
		*s = append(*s, element)
		return s
	}
	*s = append((*s)[:index+1], (*s)[index:]...)
	(*s)[index] = element
	return s
}

// Set sets the element at index to the specified value. If the specified index is not found, an ErrIndexOutOfBounds
// is thrown in a panic.
func (s *Slice[E]) Set(index uint, element E) collections.MutableList[E] {
	if index >= uint(len(*s)) {
		panic(collections.ErrIndexOutOfBounds)
	}
	(*s)[index] = element
	return s
}

// Sort sorts the slice according to the comparator passed as the argument.
func (s *Slice[E]) Sort(f collections.Comparator[E]) collections.MutableList[E] {
	sort.SliceStable(
		*s, func(i, j int) bool {
			return f((*s)[i], (*s)[j]) < 0
		},
	)
	return s
}

// String creates a printable string with brackets and comma-delimiters from the current slice.
func (s Slice[E]) String() string {
	result := make([]string, len(s))
	for i, e := range s {
		result[i] = fmt.Sprintf("%v", e)
	}
	return "[" + strings.Join(result, ", ") + "]"
}

// Iterator is an interator looping over a Slice. You can create it by calling Iterator() or MutableIterator() on a
// Slice.
type Iterator[E comparable] struct {
	backingSlice *Slice[E]
	index        int
}

// ForEachRemaining executes the specified consumer function on each remaining elements until no more elements remain
// in the iterator or an error occurs.
func (s *Iterator[E]) ForEachRemaining(f collections.Consumer[E]) {
	for s.HasNext() {
		f(s.Next())
	}
}

// HasNext returns true if the iterator has more elements remaining.
func (s Iterator[E]) HasNext() bool {
	return s.index < len(*s.backingSlice)-1
}

// Next retrieves the next element. If no more elements are remaining, an ErrIndexOutOfBounds error is returned. This
// function is best used in a for loop in conjunction with HasNext():
//
//     for iterator.HasNext() {
//         element, err := iterator.Next()
//         if err != nil {
//             // Possible race condition, this should not happen:
//             panic(err)
//         }
//     }
func (s *Iterator[E]) Next() E {
	if s.index >= len(*s.backingSlice)-1 {
		panic(collections.ErrIndexOutOfBounds)
	}
	s.index++
	return (*s.backingSlice)[s.index]
}

// Remove removes the current element from the underlying Slice.
func (s *Iterator[E]) Remove() {
	if s.index >= len(*s.backingSlice) {
		panic(collections.ErrIndexOutOfBounds)
	}
	*s.backingSlice = append((*s.backingSlice)[:s.index], (*s.backingSlice)[s.index+1:]...)
}
