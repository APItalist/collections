package mapset_test

import (
    "fmt"
    "sort"
    "strings"

    "github.com/apitalist/collections"
    "github.com/apitalist/collections/mapset"
    "github.com/apitalist/lang/try"
    "github.com/apitalist/lang/try/catch"
)

func Example() {
    // The set variable will contain a typed slice:
    set := mapset.New("a", "b", "c")

    // We can add new items to it:
    set.Add("d")

    // We can also remove items from it:
    set.Remove("a")

    // Let's loop over the items. Note that the MapSet doesn't guarantee that you will iterate over it in order.
    var result []string
    set.Iterator().ForEachRemaining(
        func(e string) {
            result = append(result, e)
        },
    )
    // Sets are not sorted by default, so we must sort the result for the output
    sort.SliceStable(result, func(i, j int) bool {
        return strings.Compare(result[i], result[j]) < 0
    })

    // We can also print slices directly:
    fmt.Printf("Here's the result directly: %v\n", result)

    // Output: Here's the result directly: [b c d]
}

func ExampleNew() {
    // Create an empty set by specifying the type:
    set1 := mapset.New[string]()
    set1.Add("a")
    fmt.Println(set1)

    // Create a set by specifying some elements:
    set2 := mapset.New("b")
    fmt.Println(set2)

    // Create a set and explicitly assign it to a MutableSet interface type:
    var set3 collections.MutableSet[string] = mapset.New[string]()
    set3.Add("c")
    fmt.Println(set3)

    // Output: [a]
    // [b]
    // [c]
}

func ExampleMapSet_stream() {
    s := mapset.New(1, 2, 3, 4, 5, 6)

    n := s.
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
    fmt.Println(n)

    // Output: [6]
}

func ExampleMapSet_add() {
    // Create a new set
    var set collections.MutableSet[string] = mapset.New("a", "b", "c", "d")

    // Add an element to the set
    set.Add("e")

    s := set.ToSlice()
    // Sets are not sorted by default, so we must sort the result for the output
    sort.SliceStable(s, func(i, j int) bool {
        return strings.Compare(s[i], s[j]) < 0
    })

    fmt.Println(s)

    // Output: [a b c d e]
}

func ExampleMapSet_remove() {
    set := mapset.New("a", "b", "c", "b", "d")

    // Remove all b's from the set:
    set.Remove("b")

    s := set.ToSlice()
    sort.SliceStable(s, func(i, j int) bool {
        return strings.Compare(s[i], s[j]) < 0
    })

    fmt.Println(s)
    // Output: [a c d]
}

func ExampleMapSet_removeAll() {
    set1 := mapset.New("a", "b", "c", "b", "d")
    set2 := mapset.New("b", "c")

    set1.RemoveAll(set2)

    s := set1.ToSlice()
    // Sets are not sorted by default, so we must sort the result for the output
    sort.SliceStable(s, func(i, j int) bool {
        return strings.Compare(s[i], s[j]) < 0
    })

    fmt.Println(s)

    // Output: [a d]
}

func ExampleMapSet_removeIf() {
    set := mapset.New(1, 2, 3, 4, 5, 6, 7)

    set.RemoveIf(
        func(item int) bool {
            // Remove all even items
            return item%2 == 0
        },
    )

    s := set.ToSlice()
    // Sets are not sorted by default, so we must sort the result for the output
    sort.SliceStable(s, func(i, j int) bool {
        return s[i] < s[j]
    })

    fmt.Println(s)

    // Output: [1 3 5 7]
}

func ExampleMapSet_retainAll() {
    set1 := mapset.New(1, 2, 3, 4, 5, 6, 7)
    set2 := mapset.New(2, 3, 4, 8)

    set1.RetainAll(set2)

    s := set1.ToSlice()
    // Sets are not sorted by default, so we must sort the result for the output
    sort.SliceStable(s, func(i, j int) bool {
        return s[i] < s[j]
    })

    fmt.Println(s)

    // Output: [2 3 4]
}

func ExampleMapSet_contains() {
    var set collections.MutableSet[string] = mapset.New("a", "b", "c", "d")

    if set.Contains("c") {
        fmt.Println("The set contains 'c'.")
    } else {
        fmt.Println("The set does not contain 'c'.")
    }

    // Output: The set contains 'c'.
}

func ExampleMapSet_iterator() {
    set := mapset.New[string]("a", "b", "c")

    var s []string
    iterator := set.Iterator()
    for iterator.HasNext() {
        s = append(s, iterator.Next())
    }
    // Sets are not sorted by default, so we must sort the result for the output
    sort.SliceStable(s, func(i, j int) bool {
        return strings.Compare(s[i], s[j]) < 0
    })
    fmt.Println(s)

    // Output: [a b c]
}

func ExampleMapSet_mutableIterator() {
    set := mapset.New[string]("a", "b", "c")

    iterator := set.MutableIterator()
    for iterator.HasNext() {
        item := iterator.Next()
        if item == "b" {
            iterator.Remove()
        }
    }

    s := set.ToSlice()
    // Sets are not sorted by default, so we must sort the result for the output
    sort.SliceStable(s, func(i, j int) bool {
        return strings.Compare(s[i], s[j]) < 0
    })
    fmt.Println(s)

    // Output: [a c]
}

func ExampleMapSet_addAll() {
    set1 := mapset.New[string]("a", "b", "c")
    set2 := mapset.New[string]("d")
    set2.AddAll(set1)

    s := set2.ToSlice()
    // Sets are not sorted by default, so we must sort the result for the output
    sort.SliceStable(s, func(i, j int) bool {
        return strings.Compare(s[i], s[j]) < 0
    })
    fmt.Println(s)

    // Output: [a b c d]
}

func ExampleMapSet_clear() {
    set := mapset.New[string]("a", "b", "c")
    set.Clear()

    fmt.Println(set)

    // Output: []
}

func ExampleMapSet_size() {
    set := mapset.New("a", "b", "c")

    fmt.Println(set.Size())

    // Output: 3
}

func ExampleMapSet_isEmpty() {
    set1 := mapset.New[string]("a", "b", "c")
    if set1.IsEmpty() {
        fmt.Println("set 1 is empty.")
    } else {
        fmt.Println("set 1 is not empty.")
    }

    set2 := mapset.New[string]()
    if set2.IsEmpty() {
        fmt.Println("set 2 is empty.")
    } else {
        fmt.Println("set 2 is not empty.")
    }

    // Output: set 1 is not empty.
    // set 2 is empty.
}

func ExampleMapSet_toSlice() {
    set := mapset.New("a", "b", "c")
    s := set.ToSlice()
    // Sets are not sorted by default, so we must sort the result for the output
    sort.SliceStable(s, func(i, j int) bool {
        return strings.Compare(s[i], s[j]) < 0
    })
    fmt.Println(s[0])
    // Output: a
}

func ExampleMapSet_iteratorHasNext() {
    set := mapset.New[string]("a", "b", "c")

    var result []string
    iterator := set.Iterator()
    for iterator.HasNext() {
        item := iterator.Next()
        result = append(result, item)
    }

    // Sets are not sorted by default, so we must sort the result for the output
    sort.SliceStable(result, func(i, j int) bool {
        return strings.Compare(result[i], result[j]) < 0
    })
    fmt.Println(result)
    // Output: [a b c]
}

func ExampleMapSet_iteratorForEachRemaining() {
    set := mapset.New[string]("a", "b", "c")

    var s []string
    iterator := set.Iterator()
    iterator.ForEachRemaining(
        func(item string) {
            s = append(s, item)
        },
    )

    // Sets are not sorted by default, so we must sort the result for the output
    sort.SliceStable(s, func(i, j int) bool {
        return strings.Compare(s[i], s[j]) < 0
    })
    fmt.Println(s)
    // Output: [a b c]
}

func ExampleIterator_next() {
    set := mapset.New[string]("a", "b", "c")

    iterator := set.Iterator()

    var s []string
    // Get the first element:
    s = append(s, iterator.Next())

    // Get the second element:
    s = append(s, iterator.Next())

    // Get the third element:
    s = append(s, iterator.Next())

    // Sets are not sorted by default, so we must sort the result for the output
    sort.SliceStable(s, func(i, j int) bool {
        return strings.Compare(s[i], s[j]) < 0
    })
    fmt.Println(s)

    // This will result in an error since the fourth element doesn't exist.
    try.Catch(
        func() {
            _ = iterator.Next()
        },
        catch.ErrorByValue(
            collections.ErrIndexOutOfBounds, func(_ error) {
                fmt.Println("set finished!")
            },
        ),
    )

    // Output: [a b c]
    // set finished!
}
