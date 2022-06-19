// Package collect provides collectors from streams to other formats.
package collect

import (
	"github.com/apitalist/collections"
	"github.com/apitalist/collections/slice"
)

func ToList[T comparable](s collections.Stream[T]) collections.MutableList[T] {
	l := slice.New[T]()
	iterator := s.Iterator()
	for iterator.HasNext() {
		l.Add(iterator.Next())
	}
	return l
}
