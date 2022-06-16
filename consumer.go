package collections

// Consumer is a function that can receive a typed parameter, and potentially return an error.
type Consumer[E any] func(E) error
