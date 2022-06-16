package collections

// Comparator is a function that compares two values. If the first parameter is larger than the second parameter, a
// negative number is returned. If the second parameter is larger, a positive number is returned. If they are equal,
// zero is returned.
type Comparator[E any] func(a, b E) int
