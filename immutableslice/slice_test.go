package immutableslice_test

import (
	"fmt"

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

func ExampleImmutableSlice_withRemoved() {
	// Create a new immutable slice:
	s := immutableslice.New("a", "b", "c", "b", "e")

	// Remove all items from the list that patch the provided parameter:
	// Don't forget that you MUST store the result as the original is not modified:
	s = s.WithRemoved("b")

	fmt.Println(s)
	// Output: [a, c, e]
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
