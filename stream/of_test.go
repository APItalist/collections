package stream_test

import (
	"fmt"

	"github.com/apitalist/collections/stream"
)

func ExampleOf() {
	result, err := stream.Of(1, 2, 3, 4, 5).ToSlice()
	if err != nil {
		panic(err)
	}
	for _, r := range result {
		fmt.Printf("%d\n", r)
	}
	// Output: 1
	// 2
	// 3
	// 4
	// 5
}
