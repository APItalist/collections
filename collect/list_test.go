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
	l := collect.ToList(s)
	// Use the list:
	i := l.IndexOf(2)
	fmt.Println(i)

	// Output: 0
}
