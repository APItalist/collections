// Package collect provides collectors from streams to other formats.
package collect

import (
	"github.com/apitalist/collections"
	"github.com/apitalist/collections/slice"
)

func ToList[T comparable](s collections.Stream[T]) (collections.MutableList[T], error) {
	l := slice.New[T]()
	iterator := s.Iterator()
	for iterator.HasNext() {
		e, err := iterator.Next()
		if err != nil {
			return nil, err
		}
		l.Add(e)
	}
	return l, nil
}
