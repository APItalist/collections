package collections

import "fmt"

// ErrIndexOutOfBounds indicates that the specified list index is too large for the amount of list elements, or an
// iterator ran out of elements.
var ErrIndexOutOfBounds = fmt.Errorf("index out of bounds")

// ErrElementNotFound indicates that an IndexOf or similar operation did not find the specified element.
var ErrElementNotFound = fmt.Errorf("element not found")
