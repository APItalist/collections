package collections

import (
    "sort"
    "sync"

    "github.com/apitalist/lang"
)

// NewSliceListMutable creates a mutable, slice-backed list.
func NewSliceListMutable[E lang.Ordered](elements ...E) MutableList[E] {
    data := make([]E, len(elements))
    copy(data, elements)
    return &sliceListMutable[E]{
        data,
        &sync.Mutex{},
    }
}

type sliceListMutable[E lang.Ordered] struct {
    data []E
    lock *sync.Mutex
}

func (s *sliceListMutable[E]) RemoveAt(index uint) error {
    s.lock.Lock()
    defer s.lock.Unlock()

    if index >= uint(len(s.data)) {
        return ErrIndexOutOfBounds
    }
    s.data = append(s.data[:index], s.data[index+1:]...)
    return nil
}

func (s *sliceListMutable[E]) Iterator() MutableIterator[E] {
    return &sliceIterator[E]{
        s,
        0,
        s.lock,
    }
}

func (s sliceListMutable[E]) IsEmpty() bool {
    s.lock.Lock()
    defer s.lock.Unlock()
    return len(s.data) == 0
}

func (s sliceListMutable[E]) Size() uint {
    s.lock.Lock()
    defer s.lock.Unlock()
    return uint(len(s.data))
}

func (s sliceListMutable[E]) ToSlice() []E {
    s.lock.Lock()
    defer s.lock.Unlock()
    result := make([]E, len(s.data))
    copy(result, s.data)
    return result
}

func (s sliceListMutable[E]) Contains(e E) bool {
    _, err := s.IndexOf(e)
    return err == nil
}

func (s sliceListMutable[E]) Get(index uint) (E, error) {
    var emptyResult E
    s.lock.Lock()
    defer s.lock.Unlock()
    if index >= uint(len(s.data)) {
        return emptyResult, ErrIndexOutOfBounds
    }
    return s.data[index], nil
}

func (s sliceListMutable[E]) IndexOf(e E) (uint, error) {
    for i, elem := range s.data {
        if elem == e {
            return uint(i), nil
        }
    }
    return 0, ErrElementNotFound
}

func (s sliceListMutable[E]) LastIndexOf(e E) (uint, error) {
    // TODO implement me
    panic("implement me")
}

func (s sliceListMutable[E]) SubList(from, to uint) MutableList[E] {
    // TODO implement me
    panic("implement me")
}

func (s sliceListMutable[E]) Add(e E) {
    s.lock.Lock()
    defer s.lock.Unlock()
    s.data = append(s.data, e)
}

func (s sliceListMutable[E]) AddAll(c Collection[E, MutableIterator[E]]) {
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

func (s sliceListMutable[E]) Clear() {
    s.lock.Lock()
    defer s.lock.Unlock()
    s.data = nil
}

func (s sliceListMutable[E]) Remove(e E) {
    s.lock.Lock()
    defer s.lock.Unlock()
    for i, entry := range s.data {
        if entry == e {
            s.data = append(s.data[:i], s.data[i+1:]...)
        }
    }
}

func (s sliceListMutable[E]) RemoveAll(c Collection[E, MutableIterator[E]]) {
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

func (s sliceListMutable[E]) RemoveIf(p Predicate[E]) {
    for i, element := range s.data {
        if p(element) {
            s.data = append(s.data[:i], s.data[i+1:]...)
        }
    }
}

func (s sliceListMutable[E]) RetainAll(c Collection[E, MutableIterator[E]]) {
    // TODO implement me
    panic("implement me")
}

func (s sliceListMutable[E]) AddAt(index uint, element E) error {
    s.lock.Lock()
    defer s.lock.Unlock()
    if index > uint(len(s.data)) {
        return ErrIndexOutOfBounds
    }
    if index == uint(len(s.data)) {
        s.data = append(s.data, element)
        return nil
    }
    s.data = append(append(s.data[:index], element), s.data[index:]...)
    return nil
}

func (s sliceListMutable[E]) Set(index uint, element E) error {
    s.lock.Lock()
    defer s.lock.Unlock()
    if index >= uint(len(s.data)) {
        return ErrIndexOutOfBounds
    }
    s.data[index] = element
    return nil
}

func (s sliceListMutable[E]) Sort(f Comparator[E]) {
    s.lock.Lock()
    defer s.lock.Unlock()
    sort.SliceStable(
        s.data, func(i, j int) bool {
            return f(s.data[i], s.data[j])
        },
    )
}
