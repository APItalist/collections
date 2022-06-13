package collections

import (
    "sort"

    "github.com/apitalist/lang"
)

// NewSliceList creates a new slice-backed list, optionally filled with the specified elements. Slice lists are not
// concurrency-safe, so locks should be applied if concurrent list operations are needed. Alternatively, immutable
// lists can be used for concurrent access.
func NewSliceList[E lang.Ordered](elements ...E) SliceList[E] {
    data := make([]E, len(elements))
    copy(data, elements)
    result := make(sliceList[E], len(elements))
    copy(result, elements)
    return &result
}

type SliceList[E lang.Ordered] MutableList[E]

type sliceList[E lang.Ordered] []E

func (s *sliceList[E]) RemoveAt(index uint) error {
    if index >= uint(len(*s)) {
        return ErrIndexOutOfBounds
    }
    *s = append((*s)[:index], (*s)[index+1:]...)
    return nil
}

func (s *sliceList[E]) Iterator() MutableIterator[E] {
    return &sliceListIterator[E]{
        s,
        -1,
    }
}

func (s sliceList[E]) IsEmpty() bool {
    return len(s) == 0
}

func (s sliceList[E]) Size() uint {
    return uint(len(s))
}

func (s sliceList[E]) ToSlice() []E {
    return s
}

func (s sliceList[E]) Contains(e E) bool {
    _, err := s.IndexOf(e)
    return err == nil
}

func (s sliceList[E]) Get(index uint) (E, error) {
    var emptyResult E
    if index >= uint(len(s)) {
        return emptyResult, ErrIndexOutOfBounds
    }
    return s[index], nil
}

func (s sliceList[E]) IndexOf(e E) (uint, error) {
    for i, elem := range s {
        if elem == e {
            return uint(i), nil
        }
    }
    return 0, ErrElementNotFound
}

func (s sliceList[E]) LastIndexOf(e E) (uint, error) {
    for i := len(s) - 1; i >= 0; i-- {
        elem := s[i]
        if elem == e {
            return uint(i), nil
        }
    }
    return 0, ErrElementNotFound
}

func (s *sliceList[E]) SubList(from, to uint) (MutableList[E], error) {
    if to >= uint(len(*s)) {
        return nil, ErrIndexOutOfBounds
    }
    subSlice := (*s)[from:to]
    return &subSlice, nil
}

func (s *sliceList[E]) Add(e E) {
    *s = append(*s, e)
}

func (s *sliceList[E]) AddAll(c Collection[E, MutableIterator[E]]) {
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

func (s *sliceList[E]) Clear() {
    *s = nil
}

func (s *sliceList[E]) Remove(e E) {
    for i, entry := range *s {
        if entry == e {
            *s = append((*s)[:i], (*s)[i+1:]...)
        }
    }
}

func (s *sliceList[E]) RemoveAll(c Collection[E, MutableIterator[E]]) {
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

func (s *sliceList[E]) RemoveIf(p Predicate[E]) {
    tmpSlice := (*s)[:0]
    for _, e := range *s {
        if !p(e) {
            tmpSlice = append(tmpSlice, e)
        }
    }
    *s = tmpSlice
}

func (s *sliceList[E]) RetainAll(c Collection[E, MutableIterator[E]]) {
    s.RemoveIf(Predicate[E](c.Contains).Negate())
}

func (s *sliceList[E]) AddAt(index uint, element E) error {
    if index > uint(len(*s)) {
        return ErrIndexOutOfBounds
    }
    if index == uint(len(*s)) {
        *s = append(*s, element)
        return nil
    }
    *s = append(append((*s)[:index], element), (*s)[index:]...)
    return nil
}

func (s *sliceList[E]) Set(index uint, element E) error {
    if index >= uint(len(*s)) {
        return ErrIndexOutOfBounds
    }
    (*s)[index] = element
    return nil
}

func (s *sliceList[E]) Sort(f Comparator[E]) {
    sort.SliceStable(
        *s, func(i, j int) bool {
            return f((*s)[i], (*s)[j])
        },
    )
}

type sliceListIterator[E lang.Ordered] struct {
    backingSlice *sliceList[E]
    index        int
}

func (s *sliceListIterator[E]) ForEachRemaining(f Consumer[E]) error {
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

func (s sliceListIterator[E]) HasNext() bool {
    return s.index < len(*s.backingSlice)-1
}

func (s *sliceListIterator[E]) Next() (E, error) {
    var emptyResult E
    if s.index >= len(*s.backingSlice)-1 {
        return emptyResult, ErrIndexOutOfBounds
    }
    s.index++
    return (*s.backingSlice)[s.index], nil
}

func (s *sliceListIterator[E]) Remove() error {
    if s.index >= len(*s.backingSlice) {
        return ErrIndexOutOfBounds
    }
    *s.backingSlice = append((*s.backingSlice)[:s.index], (*s.backingSlice)[s.index+1:]...)
    return nil
}
