// Package mapset offers a Set implementation (a collection that can only contain each item once) that is backed by a
// go map.
package mapset

import (
    "fmt"

    "github.com/apitalist/collections"
    "github.com/apitalist/collections/stream"
)

// New creates a new MapSet, a set implementation that stores data in a map.
func New[V comparable](elements ...V) MapSet[V] {
    s := make(mapSet[V], len(elements))
    for _, e := range elements {
        s[e] = struct{}{}
    }
    return &s
}

// MapSet is an interface describing a map-based set implementation.
type MapSet[V comparable] interface {
    collections.MutableSet[V]
}

type mapSet[V comparable] map[V]struct{}

func (m *mapSet[V]) MutableIterator() collections.MutableIterator[V] {
    return &iterator[V]{
        set:  m,
        data: m.ToSlice(),
        i:    -1,
    }
}

func (m mapSet[V]) Add(e V) {
    m[e] = struct{}{}
}

func (m mapSet[V]) AddAll(c collections.Collection[V]) {
    c.Iterator().ForEachRemaining(func(e V) {
        m.Add(e)
    })
}

func (m *mapSet[V]) Clear() {
    newM := make(map[V]struct{}, 0)
    *m = newM
}

func (m mapSet[V]) Remove(e V) {
    _, ok := m[e]
    if !ok {
        panic(collections.ErrElementNotFound)
    }
    delete(m, e)
}

func (m mapSet[V]) RemoveAll(c collections.Collection[V]) {
    i := c.Iterator()
    for i.HasNext() {
        e := i.Next()
        delete(m, e)
    }
}

func (m mapSet[V]) RemoveIf(p collections.Predicate[V]) {
    for e := range m {
        if p(e) {
            delete(m, e)
        }
    }
}

func (m mapSet[V]) RetainAll(c collections.Collection[V]) {
    m.RemoveIf(collections.Predicate[V](c.Contains).Negate())
}

func (m mapSet[V]) Iterator() collections.Iterator[V] {
    return &iterator[V]{
        data: m.ToSlice(),
        i:    -1,
    }
}

func (m mapSet[V]) Contains(e V) bool {
    _, ok := m[e]
    return ok
}

func (m mapSet[V]) IsEmpty() bool {
    return len(m) == 0
}

func (m mapSet[V]) Size() uint {
    return uint(len(m))
}

func (m mapSet[V]) ToSlice() []V {
    result := make([]V, len(m))
    i := 0
    for e := range m {
        result[i] = e
        i++
    }
    return result
}

func (m mapSet[V]) String() string {
    i := 0
    result := "["
    for e := range m {
        if i > 0 {
            result += ", "
        }
        i++
        result += fmt.Sprintf("%v", e)
    }
    result += "]"
    return result
}

func (m mapSet[V]) Stream() collections.Stream[V] {
    return stream.FromCollection[V](m)
}

type iterator[V comparable] struct {
    set  *mapSet[V]
    data []V
    i    int
}

func (i *iterator[V]) Remove() {
    if i.set == nil {
        panic(fmt.Errorf("iterator is not mutable"))
    }
    if i.i >= len(i.data) {
        panic(collections.ErrIndexOutOfBounds)
    }
    delete(*i.set, i.data[i.i])
}

func (i *iterator[V]) ForEachRemaining(c collections.Consumer[V]) {
    for i.HasNext() {
        c(i.Next())
    }
}

func (i *iterator[V]) HasNext() bool {
    return i.i < len(i.data)-1
}

func (i *iterator[V]) Next() V {
    if i.i >= len(i.data)-1 {
        panic(collections.ErrIndexOutOfBounds)
    }
    i.i++
    return i.data[i.i]
}
