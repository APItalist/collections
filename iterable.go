package collections

type Iterable[T any, I Iterator[T]] interface {
    Iterator() I
}
