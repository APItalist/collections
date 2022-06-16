package stream_test

import (
	"fmt"

	"github.com/apitalist/collections/stream"
)

func Example_allMatch() {
	allEven, err := stream.
		Of(2, 4, 6).
		AllMatch(
			func(e int) bool {
				return e%2 == 0
			},
		)
	if err != nil {
		panic(err)
	}
	if allEven {
		fmt.Println("All numbers are even.")
	} else {
		fmt.Printf("Not all numbers are even.")
	}

	// Output: All numbers are even.
}

func Example_anyMatch() {
	allEven, err := stream.
		Of(1, 2, 3, 4, 5, 6).
		AnyMatch(
			func(e int) bool {
				return e%2 == 0
			},
		)
	if err != nil {
		panic(err)
	}
	if allEven {
		fmt.Println("There are even numbers in the stream.")
	} else {
		fmt.Printf("There are no even numbers in the stream.")
	}

	// Output: There are even numbers in the stream.
}

func Example_filter() {
	r, err := stream.
		Of(1, 2, 3, 4, 5, 6).
		Filter(
			func(e int) bool {
				return e%2 == 0
			},
		).
		ToSlice()
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	// Output: [2 4 6]
}

func Example_toSlice() {
	r, err := stream.
		Of(1, 2, 3, 4, 5, 6).
		Filter(
			func(e int) bool {
				return e%2 == 0
			},
		).
		ToSlice()
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	// Output: [2 4 6]
}

func Example_findFirst() {
	r, err := stream.
		Of(1, 2, 3, 4, 5, 6).
		Filter(
			func(e int) bool {
				return e%2 == 0
			},
		).
		FindFirst()
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	// Output: 2
}

func Example_findAny() {
	r, err := stream.
		Of(1, 2, 3, 4, 5, 6).
		Filter(
			func(e int) bool {
				return e%2 == 0
			},
		).
		FindAny()
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}

func Example_count() {
	r, err := stream.
		Of(1, 2, 3, 4, 5, 6).
		Filter(
			func(e int) bool {
				return e%2 == 0
			},
		).
		Count()
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	// Output: 3
}

func Example_map() {
	s := stream.
		Of(1, 2, 3, 4, 5, 6).
		// The Map stream function is only usable for same-type elements due to Golang restrictions:
		Map(
			func(e int) (int, error) {
				return e * 2, nil
			},
		)

	// Use this map function to convert types.
	r, err := stream.Map(
		s, func(input int) (string, error) {
			return fmt.Sprintf("%d", input), nil
		},
	).ToSlice()

	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	// Output: [2 4 6 8 10 12]
}

func ExampleMap() {
	s, err := stream.Map(
		stream.
			Of(1, 2, 3, 4, 5, 6, 7, 8),
		func(input int) (string, error) {
			return fmt.Sprintf("%d", input), nil
		},
	).ToSlice()
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	// Output: [1 2 3 4 5 6 7 8]
}
