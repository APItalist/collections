package slice_test

import (
	"errors"
	"fmt"
	"strings"

	"github.com/apitalist/collections"
	"github.com/apitalist/collections/slice"
)

func Example() {
	// The list variable will contain a typed slice:
	list := slice.New("a", "b", "c") //nolint:ineffassign

	// You can also convert an existing slice:
	existingSlice := []string{"a", "b", "c"}
	list = (*slice.Slice[string])(&existingSlice) //nolint:ineffassign

	// Instead of the code above, you can also use this simplified function:
	list = slice.NewFromSlice(existingSlice)

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

func ExampleNewFromSlice() {
	existingSlice := []string{"a", "b", "c"}
	// Here we convert an existing slice into a *Slice representation without copying the elements. Be careful about
	// modifying existingSlice afterwards!
	list := slice.NewFromSlice[string](existingSlice)

	// We can modify the underlying slice:
	existingSlice[0] = "d"
	// Or we can modify the abstraction:
	_ = list.Set(1, "f")

	fmt.Println(list)

	// Output: [d, f, c]
}

func ExampleSlice_Stream() {
	s := slice.New(1, 2, 3, 4, 5, 6)
	
	n, err := s.
		Stream().
		Filter(
			func(e int) bool {
				return e%2 == 0
			},
		).
		Filter(
			func(e int) bool {
				return e%3 == 0
			},
		).ToSlice()

	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	// Output: [6]
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
	list := slice.New("a", "b", "c", "b", "d")

	// Remove all b's from the list:
	list.Remove("b")

	fmt.Println(list)
	// Output: [a, c, d]
}

func ExampleSlice_RemoveAll() {
	list1 := slice.New("a", "b", "c", "b", "d")
	list2 := slice.New("b", "c")

	list1.RemoveAll(list2)

	fmt.Println(list1)

	// Output: [a, d]
}

func ExampleSlice_RemoveIf() {
	list := slice.New(1, 2, 3, 4, 5, 6, 7)

	list.RemoveIf(
		func(item int) bool {
			// Remove all even items
			return item%2 == 0
		},
	)

	fmt.Println(list)

	// Output: [1, 3, 5, 7]
}

func ExampleSlice_RetainAll() {
	list1 := slice.New(1, 2, 3, 4, 5, 6, 7)
	list2 := slice.New(2, 3, 4, 8)

	list1.RetainAll(list2)

	fmt.Println(list1)

	// Output: [2, 3, 4]
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

func ExampleSlice_MutableIterator() {
	list := slice.New[string]("a", "b", "c")

	iterator := list.MutableIterator()
	for iterator.HasNext() {
		item, err := iterator.Next()
		if err != nil {
			// This should never happen except when the list is concurrently changed.
			panic(err)
		}
		if item == "b" {
			if err := iterator.Remove(); err != nil {
				panic(err)
			}
		}
	}

	fmt.Println(list)

	// Output: [a, c]
}

func ExampleSlice_AddAll() {
	list1 := slice.New[string]("a", "b", "c")
	list2 := slice.New[string]("d")
	list2.AddAll(list1)

	fmt.Println(list2)

	// Output: [d, a, b, c]
}

func ExampleSlice_AddAt() {
	list := slice.New[string]("a", "b", "c")
	if err := list.AddAt(1, "d"); err != nil {
		// This should not happen.
		panic(err)
	}

	fmt.Println(list)

	// Output: [a, d, b, c]
}

func ExampleSlice_Clear() {
	list := slice.New[string]("a", "b", "c")
	list.Clear()

	fmt.Println(list)

	// Output: []
}

func ExampleSlice_Get() {
	list := slice.New[string]("a", "b", "c")
	item, err := list.Get(1)
	if err != nil {
		// This should not happen
		panic(err)
	}

	fmt.Println(item)

	// Output: b
}

func ExampleSlice_Set() {
	list := slice.New[string]("a", "b", "c")
	if err := list.Set(1, "d"); err != nil {
		panic(err)
	}

	fmt.Println(list)

	// Output: [a, d, c]
}

func ExampleSlice_Size() {
	list := slice.New("a", "b", "c")

	fmt.Println(list.Size())

	// Output: 3
}

func ExampleSlice_Sort() {
	type customData struct {
		data string
	}

	list := slice.New(customData{"c"}, customData{"b"}, customData{"a"})
	list.Sort(
		func(a, b customData) int {
			return strings.Compare(a.data, b.data)
		},
	)
	fmt.Println(list)

	// Output: [{a}, {b}, {c}]
}

func ExampleSlice_IndexOf() {
	list := slice.New[string]("a", "b", "c", "b", "d")
	index, err := list.IndexOf("b")
	if err != nil {
		// This should not happen
		panic(err)
	}

	fmt.Println(index)

	// Output: 1
}

func ExampleSlice_LastIndexOf() {
	list := slice.New[string]("a", "b", "c", "b", "d")
	index, err := list.LastIndexOf("b")
	if err != nil {
		// This should not happen
		panic(err)
	}

	fmt.Println(index)

	// Output: 3
}

func ExampleSlice_IsEmpty() {
	list1 := slice.New[string]("a", "b", "c")
	if list1.IsEmpty() {
		fmt.Println("List 1 is empty.")
	} else {
		fmt.Println("List 1 is not empty.")
	}
	list2 := slice.New[string]()
	if list2.IsEmpty() {
		fmt.Println("List 2 is empty.")
	} else {
		fmt.Println("List 2 is not empty.")
	}

	// Output: List 1 is not empty.
	// List 2 is empty.
}

func ExampleSlice_ToSlice() {
	list := slice.New("a", "b", "c")
	s := list.ToSlice()
	fmt.Println(s[0])
	// Output: a
}

func ExampleSlice_SubList() {
	list := slice.New(1, 2, 3, 4, 5, 6, 7)
	subList, err := list.SubList(1, 3)
	if err != nil {
		// This should not happen
		panic(err)
	}
	if err := subList.Set(0, 10); err != nil {
		panic(err)
	}

	fmt.Println(list)
	fmt.Println(subList)

	// Output: [1, 2, 3, 4, 5, 6, 7]
	// [10, 3]
}

func ExampleSlice_SubList_addingItems() {
	list := slice.New(1, 2, 3, 4, 5, 6, 7)
	subList, err := list.SubList(1, 3)
	if err != nil {
		// This should not happen
		panic(err)
	}

	// Adding an item to the sublist will overwrite the parent list.
	subList.Add(10)
	fmt.Println(list)
	fmt.Println(subList)

	// Output: [1, 2, 3, 4, 5, 6, 7]
	// [2, 3, 10]
}

func ExampleSlice_SubList_removingItems() {
	list := slice.New(1, 2, 3, 4, 5, 6, 7)
	subList, err := list.SubList(1, 3)
	if err != nil {
		// This should not happen
		panic(err)
	}

	// Adding an item to the sublist will overwrite the parent list.
	subList.Remove(2)
	fmt.Println(list)
	fmt.Println(subList)

	// Output: [1, 2, 3, 4, 5, 6, 7]
	// [3]
}

func ExampleIterator_hasNext() {
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

func ExampleIterator_forEachRemaining() {
	list := slice.New[string]("a", "b", "c")

	iterator := list.Iterator()
	_ = iterator.ForEachRemaining(
		func(item string) error {
			fmt.Println(item)
			return nil
		},
	)

	// Output: a
	// b
	// c
}

// ExampleSliceIterator_Next demonstrates the use of the Next() function. This function is best used in a for loop in
// conjunction with HasNext().
func ExampleIterator_Next() {
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

func ExampleIterator_Remove() {
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

func ExampleIterator_Set() {
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
