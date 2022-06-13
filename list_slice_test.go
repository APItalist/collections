package collections_test

import (
	"fmt"

	"github.com/apitalist/collections"
)

func ExampleSliceList_Add() {
	// Create a new list
	list := collections.NewSliceList("a", "b", "c", "d")

	// Add an element to the list
	list.Add("e")

	// Iterate over the list. We ignore the returning error since our output function never fails.
	_ = list.Iterator().ForEachRemaining(
		func(e string) error {
			fmt.Print(e)
			return nil
		},
	)

	// Output: abcde
}
