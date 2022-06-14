package slice_test

import (
	"errors"
	"fmt"

	"github.com/apitalist/collections"
	"github.com/apitalist/collections/slice"
)

func Example() {
	// The list variable will contain a typed slice:
	list := slice.New("a", "b", "c")

	// We can add new items to it:
	list.Add("d")

	// We can also remove items from it:
	list.Remove("a")

	// Or we can remove by position:
	if err := list.RemoveAt(0); err != nil {
		// Position 0 is guaranteed to be filled, so this should never happen.
		panic(err)
	}

	// Let's loop over the items:
	if err := list.Iterator().ForEachRemaining(
		func(e string) error {
			fmt.Printf("Printing an element: %s\n", e)
			return nil
		},
	); err != nil {
		// This should also never happen
		panic(err)
	}

	// We can also print slices directly:
	fmt.Printf("Here's the slice directly: %v\n", list)

	// Output: Printing an element: c
	// Printing an element: d
	// Here's the slice directly: [c, d]
}

func ExampleNew() {
	// Create an empty list by specifying the type:
	list1 := slice.New[string]()
	list1.Add("a")
	fmt.Println(list1)

	// Create a list by specifying some elements:
	list2 := slice.New("b")
	fmt.Println(list2)

	// Create a list and explicitly assign it to a MutableList interface type:
	var list3 collections.MutableList[string] = slice.New[string]()
	list3.Add("c")
	fmt.Println(list3)

	// Output: [a]
	// [b]
	// [c]
}

func ExampleSlice_Add() {
	// Create a new list
	var list collections.MutableList[string] = slice.New("a", "b", "c", "d")

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

func ExampleSlice_Remove() {
	// Create a new list
	var list collections.MutableList[string] = slice.New("a", "b", "c", "d")

	// Add an element to the list
	list.Remove("c")

	// Iterate over the list. We ignore the returning error since our output function never fails.
	_ = list.Iterator().ForEachRemaining(
		func(e string) error {
			fmt.Print(e)
			return nil
		},
	)

	// Output: abd
}

func ExampleSlice_Contains() {
	var list collections.MutableList[string] = slice.New("a", "b", "c", "d")

	if list.Contains("c") {
		fmt.Println("The list contains 'c'.")
	} else {
		fmt.Println("The list does not contain 'c'.")
	}

	// Output: The list contains 'c'.
}

func ExampleSlice_IsEmpty() {
	var list collections.MutableList[string] = slice.New[string]()

	if list.IsEmpty() {
		fmt.Println("The list is empty.")
	}
	list.Add("a")
	if !list.IsEmpty() {
		fmt.Println("The list is not empty.")
	}

	// Output: The list is empty.
	// The list is not empty.
}

func ExampleSlice_String() {
	list := slice.New[string]("a", "b", "c")

	// Slice has a helper to print out nicely as a string:
	fmt.Println(list)

	// Output: [a, b, c]
}

func ExampleSlice_Iterator() {
	list := slice.New[string]("a", "b", "c")

	iterator := list.Iterator()
	for iterator.HasNext() {
		item, err := iterator.Next()
		if err != nil {
			// This should never happen except when the list is concurrently changed.
			panic(err)
		}
		fmt.Println(item)
	}

	// Output: a
	// b
	// c
}

func ExampleSlice_AddAll() {
	list1 := slice.New[string]("a", "b", "c")
	list2 := slice.New[string]("d")
	list2.AddAll(list1)

	fmt.Println(list2)

	// Output: [d, a, b, c]
}

func ExampleSliceIterator() {
	list := slice.New[string]("a", "b", "c")

	iterator := list.Iterator()
	for iterator.HasNext() {
		item, err := iterator.Next()
		if err != nil {
			// This should never happen except when the list is concurrently changed.
			panic(err)
		}
		fmt.Println(item)
	}

	// Output: a
	// b
	// c
}

// ExampleSliceIterator_Next demonstrates the use of the Next() function. This function is best used in a for loop in
// conjunction with HasNext().
func ExampleSliceIterator_Next() {
	list := slice.New[string]("a", "b", "c")

	iterator := list.Iterator()

	// Get the first element:
	e1, err := iterator.Next()
	if err != nil {
		// This should never happen
		panic(err)
	}
	fmt.Println(e1)

	// Get the second element:
	e2, err := iterator.Next()
	if err != nil {
		// This should never happen
		panic(err)
	}
	fmt.Println(e2)

	// Get the third element:
	e3, err := iterator.Next()
	if err != nil {
		// This should never happen
		panic(err)
	}
	fmt.Println(e3)

	// This will result in an error since the fourth element doesn't exist.
	_, err = iterator.Next()
	if errors.Is(err, collections.ErrIndexOutOfBounds) {
		fmt.Println("List finished!")
	} else {
		panic(err)
	}

	// Output: a
	// b
	// c
	// List finished!
}

func ExampleSliceIterator_Remove() {
	list := slice.New[string]("a", "b", "c")

	iterator := list.MutableIterator()
	for iterator.HasNext() {
		item, err := iterator.Next()
		if err != nil {
			// This should never happen except when the list is concurrently changed.
			panic(err)
		}
		if item == "b" {
			err = iterator.Remove()
			if err != nil {
				// Remove can return an error if the list has been changed in a concurrent goroutine, which is not the case
				// here, so this should never happen.
				panic(err)
			}
		}
	}
	fmt.Println(list)

	// Output: [a, c]
}

func ExampleSliceIterator_Set() {
	list := slice.New[string]("a", "b", "c")

	iterator := list.MutableIterator()
	for iterator.HasNext() {
		item, err := iterator.Next()
		if err != nil {
			// This should never happen except when the list is concurrently changed.
			panic(err)
		}
		if item == "b" {
			err = iterator.Set("d")
			if err != nil {
				// Set can return an error if the list has been changed in a concurrent goroutine, which is not the case
				// here, so this should never happen.
				panic(err)
			}
		}
	}
	fmt.Println(list)

	// Output: [a, d, c]
}
