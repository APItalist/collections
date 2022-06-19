package stream_test

import (
	"fmt"

	"github.com/apitalist/collections/stream"
)

func Example_allMatch() {
	allEven := stream.
		Of(2, 4, 6).
		AllMatch(
			func(e int) bool {
				return e%2 == 0
			},
		)
	if allEven {
		fmt.Println("All numbers are even.")
	} else {
		fmt.Printf("Not all numbers are even.")
	}

	// Output: All numbers are even.
}

func Example_anyMatch() {
	allEven := stream.
		Of(1, 2, 3, 4, 5, 6).
		AnyMatch(
			func(e int) bool {
				return e%2 == 0
			},
		)
	if allEven {
		fmt.Println("There are even numbers in the stream.")
	} else {
		fmt.Printf("There are no even numbers in the stream.")
	}

	// Output: There are even numbers in the stream.
}

func Example_filter() {
	r := stream.
		Of(1, 2, 3, 4, 5, 6).
		Filter(
			func(e int) bool {
				return e%2 == 0
			},
		).
		ToSlice()
	fmt.Println(r)

	// Output: [2 4 6]
}

func Example_toSlice() {
	r := stream.
		Of(1, 2, 3, 4, 5, 6).
		Filter(
			func(e int) bool {
				return e%2 == 0
			},
		).
		ToSlice()
	fmt.Println(r)

	// Output: [2 4 6]
}

func Example_findFirst() {
	r := stream.
		Of(1, 2, 3, 4, 5, 6).
		Filter(
			func(e int) bool {
				return e%2 == 0
			},
		).
		FindFirst()
	fmt.Println(r)

	// Output: 2
}

func Example_findAny() {
	r := stream.
		Of(1, 2, 3, 4, 5, 6).
		Filter(
			func(e int) bool {
				return e%2 == 0
			},
		).
		FindAny()
	fmt.Println(r)
}

func Example_count() {
	r := stream.
		Of(1, 2, 3, 4, 5, 6).
		Filter(
			func(e int) bool {
				return e%2 == 0
			},
		).
		Count()
	fmt.Println(r)

	// Output: 3
}

func Example_map() {
	s := stream.
		Of(1, 2, 3, 4, 5, 6).
		// The Map stream function is only usable for same-type elements due to Golang restrictions:
		Map(
			func(e int) int {
				return e * 2
			},
		)

	// Use this map function to convert types.
	r := stream.Map(
		s, func(input int) string {
			return fmt.Sprintf("%d", input)
		},
	).ToSlice()
	fmt.Println(r)

	// Output: [2 4 6 8 10 12]
}

func ExampleMap() {
	s := stream.Map(
		stream.
			Of(1, 2, 3, 4, 5, 6, 7, 8),
		func(input int) string {
			return fmt.Sprintf("%d", input)
		},
	).ToSlice()
	fmt.Println(s)
	// Output: [1 2 3 4 5 6 7 8]
}
