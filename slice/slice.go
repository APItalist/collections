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

// Slice is a slice-backed implementation from the MutableList interface. In order to guarantee proper
// operation it should always be used as a pointer.
type Slice[E comparable] []E

func (s *Slice[E]) RemoveAt(index uint) error {
    if index >= uint(len(*s)) {
        return collections.ErrIndexOutOfBounds
    }
    *s = append((*s)[:index], (*s)[index+1:]...)
    return nil
}

func (s *Slice[E]) Iterator() collections.MutableIterator[E] {
    return &SliceIterator[E]{
        s,
        -1,
    }
}

func (s *Slice[E]) IsEmpty() bool {
    return len(*s) == 0
}

func (s *Slice[E]) Size() uint {
    return uint(len(*s))
}

func (s *Slice[E]) ToSlice() []E {
    return *s
}

func (s Slice[E]) Contains(e E) bool {
    _, err := s.IndexOf(e)
    return err == nil
}

func (s *Slice[E]) Get(index uint) (E, error) {
    var emptyResult E
    if index >= uint(len(*s)) {
        return emptyResult, collections.ErrIndexOutOfBounds
    }
    return (*s)[index], nil
}

func (s *Slice[E]) IndexOf(e E) (uint, error) {
    for i, elem := range *s {
        if elem == e {
            return uint(i), nil
        }
    }
    return 0, collections.ErrElementNotFound
}

func (s *Slice[E]) LastIndexOf(e E) (uint, error) {
    for i := len(*s) - 1; i >= 0; i-- {
        elem := (*s)[i]
        if elem == e {
            return uint(i), nil
        }
    }
    return 0, collections.ErrElementNotFound
}

func (s *Slice[E]) SubList(from, to uint) (collections.MutableList[E], error) {
    if to >= uint(len(*s)) {
        return nil, collections.ErrIndexOutOfBounds
    }
    subSlice := (*s)[from:to]
    return &subSlice, nil
}

func (s *Slice[E]) Add(e E) {
    *s = append(*s, e)
}

func (s Slice[E]) AddAll(c collections.Collection[E, collections.MutableIterator[E]]) {
    iterator := c.Iterator()
    for iterator.HasNext() {
        e, err := iterator.Next()
        if err != nil {
            // todo better error handling here
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

func (s Slice[E]) RemoveAll(c collections.Collection[E, collections.MutableIterator[E]]) {
    iterator := c.Iterator()
    for iterator.HasNext() {
        e, err := iterator.Next()
        if err != nil {
            // Race condition, the underlying data set has been modified.
            continue
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

func (s *Slice[E]) RetainAll(c collections.Collection[E, collections.MutableIterator[E]]) {
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
    *s = append(append((*s)[:index], element), (*s)[index:]...)
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
            return f((*s)[i], (*s)[j])
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

// SliceIterator is an interator looping over a Slice. You can create it by calling Iterator() on a Slice.
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

// Next retrieves the next element. If no more elements are remaining, an ErrIndexOutOfBounds error is returned.
func (s *SliceIterator[E]) Next() (E, error) {
    var emptyResult E
    if s.index >= len(*s.backingSlice)-1 {
        return emptyResult, collections.ErrIndexOutOfBounds
    }
    s.index++
    return (*s.backingSlice)[s.index], nil
}

// Remove removes the current element from the underlying Slice.
func (s *SliceIterator[E]) Remove() error {
    if s.index >= len(*s.backingSlice) {
        return collections.ErrIndexOutOfBounds
    }
    *s.backingSlice = append((*s.backingSlice)[:s.index], (*s.backingSlice)[s.index+1:]...)
    return nil
}
