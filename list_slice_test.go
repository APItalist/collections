package collections_test

import (
    "fmt"

    "github.com/apitalist/collections"
)

func ExampleNewSlice() {
    // Create an empty list by specifying the type:
    list1 := collections.NewSlice[string]()
    list1.Add("a")
    fmt.Println(list1)

    // Create a list by specifying some elements:
    list2 := collections.NewSlice("b")
    fmt.Println(list2)

    // Create a list and explicitly assign it to a MutableList interface type:
    var list3 collections.MutableList[string] = collections.NewSlice[string]()
    list3.Add("c")
    fmt.Println(list3)

    // Output: [a]
    // [b]
    // [c]
}

func ExampleSlice_Add() {
    // Create a new list
    var list collections.MutableList[string] = collections.NewSlice("a", "b", "c", "d")

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
    var list collections.MutableList[string] = collections.NewSlice("a", "b", "c", "d")

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
    var list collections.MutableList[string] = collections.NewSlice("a", "b", "c", "d")

    if list.Contains("c") {
        fmt.Println("The list contains 'c'.")
    } else {
        fmt.Println("The list does not contain 'c'.")
    }

    // Output: The list contains 'c'.
}

func ExampleSlice_IsEmpty() {
    var list collections.MutableList[string] = collections.NewSlice[string]()

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
    list := collections.NewSlice[string]("a", "b", "c")

    // Slice has a helper to print out nicely as a string:
    fmt.Println(list)

    // Output: [a, b, c]
}

func ExampleSlice_Iterator() {
    list := collections.NewSlice[string]("a", "b", "c")

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

func ExampleSliceIterator() {
    list := collections.NewSlice[string]("a", "b", "c")

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

func ExampleSliceIterator_Remove() {
    list := collections.NewSlice[string]("a", "b", "c")

    iterator := list.Iterator()
    for iterator.HasNext() {
        item, err := iterator.Next()
        if err != nil {
            // This should never happen except when the list is concurrently changed.
            panic(err)
        }
        if item == "b" {
            // Remove can return an error if the list has been changed in a concurrent goroutine, which is not the case
            // here.
            _ = iterator.Remove()
        }
    }
    fmt.Println(list)

    // Output: [a, c]
}
