package collections

import "github.com/apitalist/lang"

type Collection[E lang.Ordered, I Iterator[E]] interface {
    Iterable[E, I]

    Contains(E) bool
    IsEmpty() bool
    Size() uint
    ToSlice() []E
}

type MutableCollection[E lang.Ordered, I MutableIterator[E]] interface {
    Collection[E, I]

    Add(E)
    AddAll(Collection[E, I])
    Clear()
    Remove(E)
    RemoveAll(Collection[E, I])
    RemoveIf(Predicate[E])
    RetainAll(Collection[E, I])
}

type ImmutableCollection[E lang.Ordered, T any, I Iterator[E]] interface {
    Collection[E, I]

    WithAdded(E) T
    WithAddedAll(Collection[E, I]) T
    WithCleared() T
    WithRemoved(E) T
    WithRemovedAll(Collection[E, I]) T
    WithRemovedIf(Predicate[E]) T
    WithRetainedAll(collection Collection[E, I]) T
}
