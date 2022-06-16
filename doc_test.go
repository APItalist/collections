package collections_test

import (
	"fmt"

	"github.com/apitalist/collections/immutableslice"
	"github.com/apitalist/collections/slice"
)

func Example_slice() {
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

func Example_immutableSlice() {
	// The immutableslice subpackage offers a go slice-backed list implementation that is immutable
	list := immutableslice.New("a", "b", "c")

	// Items can be added to lists. The returned list must be saved as the original list is not changed.
	list = list.WithAdded("d")

	// This won't do anything:
	list.WithAdded("e")

	// Or items can be removed. The returned list must be saved as the original list is not changed.
	list = list.WithRemoved("a")

	// This won't do anything:
	list.WithRemoved("b")

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