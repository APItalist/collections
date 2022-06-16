// Package slice offers a go slice-backed list implementation that is mutable (can be changed in-place). The List API
// gives a rich set of features and supports code completion at minimal overhead. However, the Slice implementation is
// not concurrency-safe, parallel modifications of the Slice should be avoided by using locks.
package slice

import (
	"fmt"
	"sort"
	"strings"

	"github.com/apitalist/collections"
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
// You can can create this Slice using the New() function, or manually by converting an underlying slice:
//
//     myData := []string{"a", "b", "c"}
//     mySlice := *slice.Slice(&myData)
type Slice[E comparable] []E

// RemoveAt removes the element at the specified index. If the specified index does not exist a
// collections.ErrIndexOutOfBounds is returned.
func (s *Slice[E]) RemoveAt(index uint) error {
	if index >= uint(len(*s)) {
		return collections.ErrIndexOutOfBounds
	}
	*s = append((*s)[:index], (*s)[index+1:]...)
	return nil
}

// Iterator returns an iterator to loop over the elements. Multiple concurrent iterators for the Slice may exist, but
// concurrent modification must be locked externally.
func (s *Slice[E]) Iterator() collections.Iterator[E] {
	return &SliceIterator[E]{
		s,
		-1,
	}
}

// MutableIterator returns a mutable iterator to loop over the elements. It also offers the ability to remove the
// current element from the list. Multiple concurrent iterators for the Slice may exist, but concurrent modification
// must be locked externally.
func (s *Slice[E]) MutableIterator() collections.MutableIterator[E] {
	return &SliceIterator[E]{
		s,
		-1,
	}
}

// IsEmpty returns true if the current Slice is empty.
func (s *Slice[E]) IsEmpty() bool {
	return len(*s) == 0
}

// Size returns the number of elements in the slice.
func (s *Slice[E]) Size() uint {
	return uint(len(*s))
}

// ToSlice returns the underlying raw slice. Modifications to this underlying slice will translate to the current slice.
func (s *Slice[E]) ToSlice() []E {
	return *s
}

// Contains will return true if the specified element is contained within the slice.
func (s Slice[E]) Contains(e E) bool {
	_, err := s.IndexOf(e)
	return err == nil
}

// Get will return the element at the specified index. If the index is larger than the number of elements, a
// collections.ErrIndexOutOfBounds is returned.
func (s *Slice[E]) Get(index uint) (E, error) {
	var emptyResult E
	if index >= uint(len(*s)) {
		return emptyResult, collections.ErrIndexOutOfBounds
	}
	return (*s)[index], nil
}

// IndexOf returns the index of the first element that matches the specified element. If no element is found,
// a collections.ErrElementNotFound is returned.
func (s *Slice[E]) IndexOf(e E) (uint, error) {
	for i, elem := range *s {
		if elem == e {
			return uint(i), nil
		}
	}
	return 0, collections.ErrElementNotFound
}

// LastIndexOf returns the index of the last element that matches the specified element. If no element is found,
// a collections.ErrElementNotFound is returned.
func (s *Slice[E]) LastIndexOf(e E) (uint, error) {
	for i := len(*s) - 1; i >= 0; i-- {
		elem := (*s)[i]
		if elem == e {
			return uint(i), nil
		}
	}
	return 0, collections.ErrElementNotFound
}

// SubList will return a part of the current Slice. If the specified bounds are invalid, a
// collections.ErrIndexOutOfBounds is returned.
func (s *Slice[E]) SubList(from, to uint) (collections.MutableList[E], error) {
	if from > to {
		return nil, collections.ErrIndexOutOfBounds
	}
	if to >= uint(len(*s)) {
		return nil, collections.ErrIndexOutOfBounds
	}
	subSlice := (*s)[from:to]
	newSlice := make([]E, len(subSlice))
	copy(newSlice, subSlice)
	return NewFromSlice(newSlice), nil
}

// Add adds a new element to the slice.
func (s *Slice[E]) Add(e E) {
	*s = append(*s, e)
}

// AddAll adds all elements from the passed collection to the current slice. Depending on the
func (s *Slice[E]) AddAll(c collections.Collection[E]) {
	iterator := c.Iterator()
	for iterator.HasNext() {
		e, err := iterator.Next()
		if err != nil {
			panic(err)
		}
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
		e, err := iterator.Next()
		if err != nil {
			panic(err)
		}
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

func (s *Slice[E]) AddAt(index uint, element E) error {
	if index > uint(len(*s)) {
		return collections.ErrIndexOutOfBounds
	}
	if index == uint(len(*s)) {
		*s = append(*s, element)
		return nil
	}
	*s = append((*s)[:index+1], (*s)[index:]...)
	(*s)[index] = element
	return nil
}

func (s *Slice[E]) Set(index uint, element E) error {
	if index >= uint(len(*s)) {
		return collections.ErrIndexOutOfBounds
	}
	(*s)[index] = element
	return nil
}

func (s *Slice[E]) Sort(f collections.Comparator[E]) {
	sort.SliceStable(
		*s, func(i, j int) bool {
			return f((*s)[i], (*s)[j]) < 0
		},
	)
}

func (s Slice[E]) String() string {
	result := make([]string, len(s))
	for i, e := range s {
		result[i] = fmt.Sprintf("%v", e)
	}
	return "[" + strings.Join(result, ", ") + "]"
}

// SliceIterator is an interator looping over a Slice. You can create it by calling Iterator() or MutableIterator() on a
// Slice.
type SliceIterator[E comparable] struct {
	backingSlice *Slice[E]
	index        int
}

// ForEachRemaining executes the specified consumer function on each remaining elements until no more elements remain
// in the iterator or an error occurs.
func (s *SliceIterator[E]) ForEachRemaining(f collections.Consumer[E]) error {
	for s.HasNext() {
		element, err := s.Next()
		if err != nil {
			// No more elements remaining
			return nil
		}
		if err := f(element); err != nil {
			return err
		}
	}
	return nil
}

// HasNext returns true if the iterator has more elements remaining.
func (s SliceIterator[E]) HasNext() bool {
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
func (s *SliceIterator[E]) Next() (E, error) {
	var emptyResult E
	if s.index >= len(*s.backingSlice)-1 {
		return emptyResult, collections.ErrIndexOutOfBounds
	}
	s.index++
	return (*s.backingSlice)[s.index], nil
}

// Set sets the current element in the underlying Slice. If the iterator currently doesn't point to a valid element,
// for example Next() hasn't been called yet, a collections.ErrIndexOutOfBounds is returned.
func (s *SliceIterator[E]) Set(e E) error {
	if s.index >= len(*s.backingSlice) {
		return collections.ErrIndexOutOfBounds
	}
	(*s.backingSlice)[s.index] = e
	return nil
}

// Remove removes the current element from the underlying Slice.
func (s *SliceIterator[E]) Remove() error {
	if s.index >= len(*s.backingSlice) {
		return collections.ErrIndexOutOfBounds
	}
	*s.backingSlice = append((*s.backingSlice)[:s.index], (*s.backingSlice)[s.index+1:]...)
	return nil
}
