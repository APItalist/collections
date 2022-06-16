// Package immutableslice offers an immutable (unchangable) list implementation, where the slice is copied every time
// an element is added. This ensures easy use for concurrent access.
package immutableslice

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/apitalist/collections"
	"github.com/apitalist/collections/stream"
)

// New creates a new immutable slice, optionally with the passed elements already added to the slice. You can use it
// in two ways. If you want to create an empty slice, specify the type:
//
//     s := immutableslice.New[string]()
//
// If you add initial types, the type will be inferred. You can skip explicitly specifying it:
//
//     s := immutableslice.New("a", "b", "c")
func New[E comparable](elements ...E) ImmutableSlice[E] {
	data := make([]E, len(elements))
	copy(data, elements)
	return &slice[E]{
		data: data,
	}
}

// ImmutableSlice is a go slice-backed immutable list. Immutability ensures that the implementation is safe to use in a
// concurrent-access environment.
//
// You can create a new ImmutableSlice using the New() function:
//
//     s := immutableslice.New[string]()
//
// You can also pass items to initialize the slice. This also saves you from needing to pass the type specification:
//
//     s := immutableslice.New("a", "b", "c")
//
// You can then use the slice. However, any modification will make a copy, which you need to store.
//
// Correct:
//
//     s = s.WithAdded("d")
//
// Incorrect:
//
//     s.WithAdded("d")
type ImmutableSlice[E comparable] interface {
	collections.ImmutableList[E]
}

type slice[E comparable] struct {
	data []E
}

// Stream creates a processing stream from the current slice.
func (s slice[E]) Stream() collections.Stream[E] {
	return stream.FromCollection[E](s)
}

func (s slice[E]) Iterator() collections.Iterator[E] {
	return &sliceIterator[E]{
		slice: s,
		index: -1,
		lock:  &sync.Mutex{},
	}
}

func (s slice[E]) Contains(e E) bool {
	for _, entry := range s.data {
		if e == entry {
			return true
		}
	}
	return false
}

func (s slice[E]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s slice[E]) Size() uint {
	return uint(len(s.data))
}

func (s slice[E]) ToSlice() []E {
	result := make([]E, len(s.data))
	copy(result, s.data)
	return result
}

func (s slice[E]) Get(index uint) (E, error) {
	var defaultReturn E
	if index >= uint(len(s.data)) {
		return defaultReturn, collections.ErrIndexOutOfBounds
	}
	return s.data[index], nil
}

func (s slice[E]) IndexOf(e E) (uint, error) {
	for i, entry := range s.data {
		if e == entry {
			return uint(i), nil
		}
	}
	return 0, collections.ErrElementNotFound
}

func (s slice[E]) LastIndexOf(e E) (uint, error) {
	for i := len(s.data) - 1; i >= 0; i-- {
		elem := (s.data)[i]
		if elem == e {
			return uint(i), nil
		}
	}
	return 0, collections.ErrElementNotFound
}

func (s slice[E]) SubList(from, to uint) (collections.ImmutableList[E], error) {
	if from > to || to >= uint(len(s.data)) {
		return nil, collections.ErrIndexOutOfBounds
	}
	return slice[E]{
		s.data[from:to],
	}, nil
}

func (s slice[E]) WithAdded(e E) collections.ImmutableList[E] {
	newSlice := make([]E, len(s.data)+1)
	copy(newSlice[:len(s.data)], s.data)
	newSlice[len(s.data)] = e
	return &slice[E]{
		newSlice,
	}
}

func (s slice[E]) WithAddedAll(c collections.Collection[E]) collections.ImmutableList[E] {
	newSlice := make([]E, len(s.data)+int(c.Size()))
	copy(newSlice[:len(s.data)], s.data)
	iterator := c.Iterator()
	i := len(s.data)
	for iterator.HasNext() {
		e, err := iterator.Next()
		if err != nil {
			panic(err)
		}
		newSlice[i] = e
		i++
	}
	return &slice[E]{
		newSlice,
	}
}

func (s slice[E]) WithCleared() collections.ImmutableList[E] {
	return &slice[E]{}
}

func (s slice[E]) WithRemoved(e E) collections.ImmutableList[E] {
	var newSlice []E
	for _, element := range s.data {
		if element != e {
			newSlice = append(newSlice, element)
		}
	}
	return &slice[E]{
		data: newSlice,
	}
}

func (s slice[E]) WithRemovedAll(c collections.Collection[E]) collections.ImmutableList[E] {
	var newSlice []E
	for _, element := range s.data {
		if !c.Contains(element) {
			newSlice = append(newSlice, element)
		}
	}
	return &slice[E]{
		data: newSlice,
	}
}

func (s slice[E]) WithRemovedIf(p collections.Predicate[E]) collections.ImmutableList[E] {
	var newSlice []E
	for _, element := range s.data {
		if !p(element) {
			newSlice = append(newSlice, element)
		}
	}
	return &slice[E]{
		data: newSlice,
	}
}

func (s slice[E]) WithRetainedAll(c collections.Collection[E]) collections.ImmutableList[E] {
	var newSlice []E
	for _, element := range s.data {
		if c.Contains(element) {
			newSlice = append(newSlice, element)
		}
	}
	return &slice[E]{
		data: newSlice,
	}
}

func (s slice[E]) WithAddedAt(index uint, element E) (collections.ImmutableList[E], error) {
	if index > uint(len(s.data)) {
		return nil, collections.ErrIndexOutOfBounds
	}
	newSlice := make([]E, len(s.data)+1)
	copy(newSlice[:index], s.data[:index])
	newSlice[index] = element
	copy(newSlice[index+1:], s.data[index:])
	return &slice[E]{
		data: newSlice,
	}, nil
}

func (s slice[E]) WithSet(index uint, element E) (collections.ImmutableList[E], error) {
	if index >= uint(len(s.data)) {
		return nil, collections.ErrIndexOutOfBounds
	}
	newSlice := make([]E, len(s.data))
	copy(newSlice, s.data)
	newSlice[index] = element
	return &slice[E]{
		data: newSlice,
	}, nil
}

func (s slice[E]) WithSorted(c collections.Comparator[E]) collections.ImmutableList[E] {
	newSlice := make([]E, len(s.data))
	copy(newSlice, s.data)
	sort.SliceStable(
		newSlice, func(i, j int) bool {
			return c(newSlice[i], newSlice[j]) < 0
		},
	)
	return &slice[E]{
		data: newSlice,
	}
}

func (s slice[E]) WithRemovedAt(index uint) (collections.ImmutableList[E], error) {
	if index >= uint(len(s.data)) {
		return nil, collections.ErrIndexOutOfBounds
	}
	newSlice := make([]E, len(s.data)-1)
	copy(newSlice[:index], s.data[:index])
	copy(newSlice[index:], s.data[index+1:])
	return &slice[E]{
		data: newSlice,
	}, nil
}

func (s slice[E]) String() string {
	result := make([]string, len(s.data))
	for i, e := range s.data {
		result[i] = fmt.Sprintf("%v", e)
	}
	return "[" + strings.Join(result, ", ") + "]"
}

type sliceIterator[E comparable] struct {
	slice slice[E]
	index int
	lock  *sync.Mutex
}

func (s *sliceIterator[E]) ForEachRemaining(c collections.Consumer[E]) error {
	s.lock.Lock()
	for s.index < len(s.slice.data)-1 {
		element := s.slice.data[s.index]
		s.index++
		s.lock.Unlock()
		err := c(element)
		if err != nil {
			return err
		}
		s.lock.Lock()
	}
	s.lock.Unlock()
	return nil
}

func (s sliceIterator[E]) HasNext() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.index < len(s.slice.data)-1
}

func (s *sliceIterator[E]) Next() (E, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	var defaultReturn E
	if s.index >= len(s.slice.data)-1 {
		return defaultReturn, collections.ErrIndexOutOfBounds
	}
	s.index++
	return s.slice.data[s.index], nil
}
