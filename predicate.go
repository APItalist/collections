package collections

type Predicate[E any] func(E) bool

func (p Predicate[E]) Negate() Predicate[E] {
    return func(e E) bool {
        return !p(e)
    }
}

func (p Predicate[E]) And(p2 Predicate[E]) Predicate[E] {
    return func(e E) bool {
        return p(e) && p2(e)
    }
}

func (p Predicate[E]) Or(p2 Predicate[E]) Predicate[E] {
    return func(e E) bool {
        return p(e) || p2(e)
    }
}
