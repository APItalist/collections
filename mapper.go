package collections

type Mapper[T comparable, K comparable] func(T) (K, error)
