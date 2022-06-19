package collections

// Set is a Collection type that guarantee that one value can only be contained once.
type Set[V comparable] interface {
	Collection[V]
}

type MutableSet[V comparable] interface {
	Set[V]
	MutableCollection[V]
}

type ImmutableSet[V comparable] interface {
	Set[V]
	ImmutableCollection[V, ImmutableSet[V]]
}
