package collections

import "github.com/ysuzuki19/collections-go/set"

func NewSet[T comparable](elements ...T) set.Set[T] {
	return set.New(elements...)
}
