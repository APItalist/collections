package stream_test

import (
	"fmt"

	"github.com/apitalist/collections/stream"
)

func ExampleOf() {
	result := stream.Of(1, 2, 3, 4, 5).ToSlice()
	for _, r := range result {
		fmt.Printf("%d\n", r)
	}
	// Output: 1
	// 2
	// 3
	// 4
	// 5
}
