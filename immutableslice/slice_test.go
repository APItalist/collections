package immutableslice_test

import (
	"fmt"
	"strings"

	"github.com/apitalist/collections/immutableslice"
)

func Example() {
	// Create a new immutable slice:
	s := immutableslice.New("a", "b", "c")

	// Add some items. Don't forget that you MUST store the result as the original is not modified:
	s = s.WithAdded("d").WithAdded("e")

	// This won't work:
	s.WithAdded("f")

	// We can also create a copy with an item removed:
	s = s.WithRemoved("a")

	fmt.Println(s)

	// Output: [b, c, d, e]
}

func ExampleImmutableSlice_withAdded() {
	// Create a new immutable slice:
	s := immutableslice.New("a", "b", "c")

	// Add some items. Don't forget that you MUST store the result as the original is not modified:
	s = s.WithAdded("d").WithAdded("e")

	fmt.Println(s)
	// Output: [a, b, c, d, e]
}

func ExampleImmutableSlice_withAddedAll() {
	s1 := immutableslice.New("a", "b", "c")
	s2 := immutableslice.New("d", "e")

	s1 = s1.WithAddedAll(s2)

	fmt.Println(s1)
	// Output: [a, b, c, d, e]
}

func ExampleImmutableSlice_withAddedAt() {
	// Create a new immutable slice:
	s := immutableslice.New("a", "b", "c")

	// Add an item at a specific position. Remember, you need to save the return value.
	s, err := s.WithAddedAt(1, "d")
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
	// Output: [a, d, b, c]
}

func ExampleImmutableSlice_withCleared() {
	s := immutableslice.New("a", "b", "c")

	s = s.WithCleared()

	fmt.Println(s)
	// Output: []
}

func ExampleImmutableSlice_contains() {
	s := immutableslice.New("a", "b", "c")

	if s.Contains("b") {
		fmt.Println("The list contains 'b'.")
	} else {
		fmt.Println("The list does not contain 'b'.")
	}

	// Output: The list contains 'b'.
}

func ExampleImmutableSlice_indexOf() {
	s := immutableslice.New("a", "b", "c")

	i, err := s.IndexOf("b")
	if err != nil {
		panic(err)
	}

	fmt.Printf("'b' is in position %d\n", i)

	// Output: 'b' is in position 1
}

func ExampleImmutableSlice_isEmpty() {
	s := immutableslice.New("a", "b", "c")

	if s.IsEmpty() {
		fmt.Println("The list is empty.")
	} else {
		fmt.Println("The list is not empty.")
	}

	// Output: The list is not empty.
}

func ExampleImmutableSlice_iterator() {
	s := immutableslice.New("a", "b", "c")

	iterator := s.Iterator()

	for iterator.HasNext() {
		e, err := iterator.Next()
		if err != nil {
			panic(err)
		}
		fmt.Println(e)
	}

	// Output: a
	// b
	// c
}

func ExampleImmutableSlice_lastIndexOf() {
	s := immutableslice.New("a", "b", "c", "b", "a")

	i, err := s.LastIndexOf("b")
	if err != nil {
		panic(err)
	}

	fmt.Printf("'b' is last in position %d\n", i)

	// Output: 'b' is last in position 3
}

func ExampleImmutableSlice_withSet() {
	// Create a new immutable slice:
	s := immutableslice.New("a", "b", "c")

	// Set the item at a specific position. Remember, you need to save the return value.
	s, err := s.WithSet(1, "d")
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
	// Output: [a, d, c]
}

type customData struct {
	data string
}

func ExampleImmutableSlice_withSorted() {
	// Create a new immutable slice:
	s := immutableslice.New(customData{"c"}, customData{"b"}, customData{"a"})

	// Create a sorted copy of the slice
	s = s.WithSorted(
		func(a, b customData) int {
			return strings.Compare(a.data, b.data)
		},
	)

	fmt.Println(s)
	// Output: [{a}, {b}, {c}]
}

func ExampleImmutableSlice_withRemoved() {
	// Create a new immutable slice:
	s := immutableslice.New("a", "b", "c", "b", "e")

	// Remove all items from the list that patch the provided parameter:
	// Don't forget that you MUST store the result as the original is not modified:
	s = s.WithRemoved("b")

	fmt.Println(s)
	// Output: [a, c, e]
}

func ExampleImmutableSlice_withRemovedAll() {
	s1 := immutableslice.New("a", "b", "c", "b", "e")
	s2 := immutableslice.New("b", "c")

	s1 = s1.WithRemovedAll(s2)

	fmt.Println(s1)
	// Output: [a, e]
}

func ExampleImmutableSlice_withRemovedAt() {
	s1 := immutableslice.New("a", "b", "c", "b", "e")

	s1, err := s1.WithRemovedAt(2)
	if err != nil {
		panic(err)
	}

	fmt.Println(s1)
	// Output: [a, b, b, e]
}

func ExampleImmutableSlice_withRemovedIf() {
	s1 := immutableslice.New("a", "b", "c", "b", "e")

	s1 = s1.WithRemovedIf(
		func(e string) bool {
			return e == "b"
		},
	)

	fmt.Println(s1)
	// Output: [a, c, e]
}

func ExampleImmutableSlice_withRetainedAll() {
	s1 := immutableslice.New("a", "b", "c", "b", "e")
	s2 := immutableslice.New("b", "c")

	s1 = s1.WithRetainedAll(s2)

	fmt.Println(s1)
	// Output: [b, c, b]
}

func ExampleImmutableSlice_subList() {
	s1 := immutableslice.New("a", "b", "c", "b", "e")

	s2, err := s1.SubList(2, 4)
	if err != nil {
		panic(err)
	}

	fmt.Println(s2)
	// Output: [c, b]
}

func ExampleImmutableSlice_get() {
	// Create a new immutable slice:
	s := immutableslice.New("a", "b", "c", "b", "e")

	// Get an item:
	item, err := s.Get(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(item)

	// Output: b
}

func ExampleNew() {
	// Create a new immutable slice using the New function:
	s := immutableslice.New("a", "b", "c") //nolint:ineffassign

	// If you want to create an empty slice, you may want to pass the specific type:
	s = immutableslice.New[string]()

	// Then you can use the slice. Remember, any change to the slice results in a new slice being created. You MUST
	// store the returned new slice:
	s = s.WithAdded("d").WithAdded("e")

	// You can quickly print the slice:
	fmt.Println(s)

	// Output: [d, e]
}
