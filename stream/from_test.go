package stream_test

import (
	"fmt"

	"github.com/apitalist/collections/slice"
	"github.com/apitalist/collections/stream"
)

func ExampleFromCollection() {
	s := slice.New(1, 2, 3, 4, 5, 6)

	s2 := stream.FromCollection[int](s).Filter(
		func(e int) bool {
			return e%2 == 0
		},
	).ToSlice()

	fmt.Println(s2)

	// Output: [2 4 6]
}
