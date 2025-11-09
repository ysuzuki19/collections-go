package main

import (
	"github.com/ysuzuki19/collections-go/set"
)

func main() {
	s := set.New[int]()

	if !s.IsEmpty() {
		panic("Set should be empty initially")
	}

	s.Insert(1)
	if s.IsEmpty() {
		panic("Set should not be empty after insertion")
	}

	s.Insert(2)
	s.Insert(3)

	if !s.Contains(2) {
		panic("Set should contain 2")
	}

	s.Remove(2)
	if s.Contains(2) {
		panic("Set should not contain 2 after removal")
	}
}
