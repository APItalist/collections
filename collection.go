package collections

type Collection[E comparable, I Iterator[E]] interface {
    Iterable[E, I]

    Contains(E) bool
    IsEmpty() bool
    Size() uint
    ToSlice() []E
}

type MutableCollection[E comparable, I MutableIterator[E]] interface {
    Collection[E, I]

    Add(E)
    AddAll(Collection[E, I])
    Clear()
    Remove(E)
    RemoveAll(Collection[E, I])
    RemoveIf(Predicate[E])
    RetainAll(Collection[E, I])
}

type ImmutableCollection[E comparable, T any, I Iterator[E]] interface {
    Collection[E, I]

    WithAdded(E) T
    WithAddedAll(Collection[E, I]) T
    WithCleared() T
    WithRemoved(E) T
    WithRemovedAll(Collection[E, I]) T
    WithRemovedIf(Predicate[E]) T
    WithRetainedAll(collection Collection[E, I]) T
}
