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
    isSmallerThanSix := func(e int) bool {
        return e < 6
    }
    p := collections.Predicate[int](isLargerThanThree).And(isSmallerThanSix)
    if p(element) {
        fmt.Printf("%d is larger than 3 and smaller than 6", element)
    } else {
        fmt.Printf("%d is not larger than 3 and smaller than 6", element)
    }

    // Output: 5 is larger than 3 and smaller than 6
}
