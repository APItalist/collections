package collections

import (
    "sync"

    "github.com/apitalist/lang"
)

type sliceIterator[E lang.Ordered] struct {
    backingSlice *sliceListMutable[E]
    index        uint
    lock         *sync.RWMutex
}

func (s sliceIterator[E]) ForEachRemaining(f Consumer[E]) error {
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

func (s sliceIterator[E]) HasNext() bool {
    s.lock.RLock()
    defer s.lock.RUnlock()
    return s.index < uint(len(s.backingSlice.data))-1
}

func (s sliceIterator[E]) Next() (E, error) {
    var emptyResult E
    s.lock.RLock()
    defer s.lock.RUnlock()
    if s.index >= uint(len(s.backingSlice.data))-1 {
        return emptyResult, ErrIndexOutOfBounds
    }
    s.index++
    return s.backingSlice.data[s.index], nil
}

func (s sliceIterator[E]) Remove() error {
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.index >= uint(len(s.backingSlice.data)) {
        return ErrIndexOutOfBounds
    }
    s.backingSlice.data = append(s.backingSlice.data[:s.index], s.backingSlice.data[s.index+1:]...)
    return nil
}
