package collections_test

import (
    "fmt"

    "github.com/apitalist/collections"
)

func ExamplePredicate() {
    element := 5
    isLargerThanThree := func(e int) bool {
        return e > 3
    }
    p := collections.Predicate[int](isLargerThanThree).Negate()
    if !p(element) {
        fmt.Printf("%d is larger than 3", element)
    } else {
        fmt.Printf("%d is smaller than 3", element)
    }

    // Output: 5 is larger than 3
}
