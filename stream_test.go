package collections_test

import (
	"fmt"

	"github.com/apitalist/collections/collect"
	"github.com/apitalist/collections/slice"
	"github.com/apitalist/collections/stream"
)

func Example_stream() {
	// Create a stream from elements:
	s := stream.Of(1, 2, 3, 4, 5, 6)

	// Filter them:
	s = s.Filter(func(i int) bool { return i%2 == 0 })

	// Map them to strings:
	s2 := stream.Map(
		s, func(input int) string {
			return fmt.Sprintf("n: %d", input)
		},
	)

	// Find the first element:
	fmt.Println(s2.FindFirst())

	// Output: n: 2
}

func ExampleStream() {
	// Create a stream from elements:
	s := stream.Of(1, 2, 3, 4, 5, 6)

	// Filter them:
	s = s.Filter(func(i int) bool { return i%2 == 0 })

	// Map them to strings:
	s2 := stream.Map(
		s, func(input int) string {
			return fmt.Sprintf("n: %d", input)
		},
	)

	// Find the first element:
	fmt.Println(s2.FindFirst())

	// Output: n: 2
}

func ExampleStream_lists() {
	// You can also create streams from lists:
	r := slice.
		New(1, 2, 3, 4).
		Stream().
		Filter(func(i int) bool { return i%2 == 0 }).
		ToSlice()

	fmt.Println(r)
	// Output: [2 4]
}

func ExampleStream_collect() {
	// You can collect the stream output:
	l, err := collect.ToList(
		slice.
			New(1, 2, 3, 4).
			Stream().
			Filter(func(i int) bool { return i%2 == 0 }),
	)
	if err != nil {
		panic(err)
	}

	i := l.IndexOf(2)

	fmt.Println(i)

	// Output: 0
}
