package collections_test

import (
	"fmt"

	"github.com/apitalist/collections/slice"
)

func Example() {
	// The slice subpackage offers a go slice-backed list implementation
	list := slice.New("a", "b", "c")

	// Items can be added to lists:
	list.Add("d")

	// Or items can be removed:
	list.Remove("a")

	// We can also print all items in a list:
	iterator := list.Iterator()
	for iterator.HasNext() {
		element, err := iterator.Next()
		if err != nil {
			// This should never happen:
			panic(err)
		}
		fmt.Println(element)
	}

	// Output: b
	// c
	// d
}
