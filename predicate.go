package collections

// NewPredicate creates a predicate from the specified function.
func NewPredicate[E any](f func(E) bool) Predicate[E] {
    return f
}

// Predicate is a utility to make a boolean decision about an element, for example for a filter. Converting a simple
// function to a predicate offers the ability to add boolean operations, such as creating and-relations with other
// predicates.
type Predicate[E any] func(E) bool

// Negate creates a new predicate with the negated effect.
func (p Predicate[E]) Negate() Predicate[E] {
    return func(e E) bool {
        return !p(e)
    }
}

// And creates a new predicate with the current predicate and the passed predicate combined in an AND boolean relation.
func (p Predicate[E]) And(p2 Predicate[E]) Predicate[E] {
    return func(e E) bool {
        return p(e) && p2(e)
    }
}

// Or creates a new predicate with the current predicate and the passed predicate combined in an OR boolean relation.
func (p Predicate[E]) Or(p2 Predicate[E]) Predicate[E] {
    return func(e E) bool {
        return p(e) || p2(e)
    }
}
