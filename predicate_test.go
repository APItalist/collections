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
	isEven := func(e int) bool {
		return e%2 == 0
	}
	p := collections.NewPredicate(isLargerThanThree).And(isSmallerThanSix).Or(isEven).Negate()
	if p(element) {
		fmt.Printf("%d is not larger than 3 and smaller than 6 or even", element)
	} else {
		fmt.Printf("%d is larger than 3 and smaller than 6 or even", element)
	}

	// Output: 5 is larger than 3 and smaller than 6 or even
}

func ExamplePredicate_And() {
	element := 5
	isLargerThanThree := func(e int) bool {
		return e > 3
	}
	isSmallerThanSix := func(e int) bool {
		return e < 6
	}
	p := collections.NewPredicate(isLargerThanThree).And(isSmallerThanSix)
	if p(element) {
		fmt.Printf("%d is larger than 3 and smaller than 6", element)
	} else {
		fmt.Printf("%d is not larger than 3 and smaller than 6", element)
	}

	// Output: 5 is larger than 3 and smaller than 6
}

func ExamplePredicate_Or() {
	element := 5
	isLargerThanFour := func(e int) bool {
		return e > 4
	}
	isSmallerThanOne := func(e int) bool {
		return e < 1
	}
	p := collections.NewPredicate(isLargerThanFour).Or(isSmallerThanOne)
	if p(element) {
		fmt.Printf("%d is larger than 4 or smaller than 1", element)
	} else {
		fmt.Printf("%d is not larger than 4 or smaller than 1", element)
	}

	// Output: 5 is larger than 4 or smaller than 1
}

func ExamplePredicate_Negate() {
	element := 5
	isLargerThanFour := func(e int) bool {
		return e > 4
	}
	p := collections.NewPredicate(isLargerThanFour).Negate()
	if p(element) {
		fmt.Printf("%d is not larger than 4", element)
	} else {
		fmt.Printf("%d is larger than 4", element)
	}

	// Output: 5 is larger than 4
}
