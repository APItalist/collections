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
		s, func(input int) (string, error) {
			return fmt.Sprintf("n: %d", input), nil
		},
	)

	// Find the first element:
	firstElement, err := s2.FindFirst()
	if err != nil {
		panic(err)
	}

	fmt.Println(firstElement)

	// Output: n: 2
}

func ExampleStream() {
	// Create a stream from elements:
	s := stream.Of(1, 2, 3, 4, 5, 6)

	// Filter them:
	s = s.Filter(func(i int) bool { return i%2 == 0 })

	// Map them to strings:
	s2 := stream.Map(
		s, func(input int) (string, error) {
			return fmt.Sprintf("n: %d", input), nil
		},
	)

	// Find the first element:
	firstElement, err := s2.FindFirst()
	if err != nil {
		panic(err)
	}

	fmt.Println(firstElement)

	// Output: n: 2
}

func ExampleStream_lists() {
	// You can also create streams from lists:
	r, err := slice.
		New(1, 2, 3, 4).
		Stream().
		Filter(func(i int) bool { return i%2 == 0 }).
		ToSlice()
	if err != nil {
		panic(err)
	}

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

	i, err := l.IndexOf(2)
	if err != nil {
		panic(err)
	}

	fmt.Println(i)

	// Output: 0
}
