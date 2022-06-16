package collect_test

import (
	"fmt"

	"github.com/apitalist/collections/collect"
	"github.com/apitalist/collections/stream"
)

func ExampleToList() {
	// Create a stream:
	s := stream.Of(1, 2, 3, 4, 5).Filter(func(i int) bool { return i%2 == 0 })
	// Convert a stream to a list:
	l, err := collect.ToList(s)
	if err != nil {
		panic(err)
	}
	// Use the list:
	i, err := l.IndexOf(2)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	// Output: 0
}
