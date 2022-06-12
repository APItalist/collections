package collections

import (
    "github.com/apitalist/lang"
)

type List[E lang.Ordered, T any, I Iterator[E]] interface {
    Collection[E, I]

    Get(index uint) (E, error)
    IndexOf(E) (uint, error)
    LastIndexOf(E) (uint, error)
    SubList(from, to uint) T
}

type MutableList[E lang.Ordered] interface {
    List[E, MutableList[E], MutableIterator[E]]
    MutableCollection[E, MutableIterator[E]]

    AddAt(index uint, element E) error
    Set(index uint, element E) error
    Sort(Comparator[E])
    RemoveAt(index uint) error
}

type ImmutableList[E lang.Ordered] interface {
    List[E, ImmutableList[E], Iterator[E]]
    ImmutableCollection[E, ImmutableList[E], Iterator[E]]

    WithAddedAt(index uint, element E) ImmutableList[E]
    WithSet(index uint, element E) (ImmutableList[E], error)
    WithSorted(Comparator[E]) ImmutableList[E]
    WithRemovedAt(index uint) (ImmutableList[E], error)
}
