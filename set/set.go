package set

// exists is an empty struct used to represent the presence of an element in the set.
type exists struct{}

// Set is a generic set implementation using a map for storage.
type Set[T comparable] struct {
	data map[T]exists
}

// New creates a new Set and initializes it with the provided elements.
//
// Example:
//
//	s1 := set.New[int]()
//	s2 := set.New(1, 2, 3)
//	s3 := set.New[string]()
//	s4 := set.New("a", "b", "c")
//
//	type Pos struct{ X, Y int }
//	v := []Pos{{1, 2}, {3, 4}}
//	s5 := set.New(v...)
func New[T comparable](elements ...T) Set[T] {
	s := Set[T]{
		data: make(map[T]exists),
	}
	for _, e := range elements {
		s.Insert(e)
	}
	return s
}

// Insert adds one or more elements to the set.
//
// Example:
//
//	st := set.New[int]()
//	st.Insert(1)
//	require.True(st.Equals(set.New(1)))
//	st.Insert(2, 3)
//	require.True(st.Equals(set.New(1, 2, 3)))
//	st.Insert(2, 4)
//	require.True(st.Equals(set.New(1, 2, 3, 4)))
func (s *Set[T]) Insert(elements ...T) {
	for _, element := range elements {
		s.data[element] = exists{}
	}
}

// Remove deletes one or more elements from the set.
//
// Example:
//
//	st := set.New(1, 2, 3, 4)
//	st.Remove(2)
//	require.True(st.Equals(set.New(1, 3, 4)))
//	st.Remove(1, 3)
//	require.True(st.Equals(set.New(4)))
//	st.Remove(5)
//	require.True(st.Equals(set.New(4)))
//	st.Remove(10, 4)
//	require.True(st.IsEmpty())
func (s *Set[T]) Remove(elements ...T) {
	for _, element := range elements {
		delete(s.data, element)
	}
}

// Contains checks if the set contains a specific element.
//
// Example:
//
//	st := set.New(1, 2, 3)
//	require.True(st.Contains(1))
//	require.True(st.Contains(2))
//	require.False(st.Contains(4))
//	require.False(st.Contains(0))
func (s Set[T]) Contains(element T) bool {
	_, exists := s.data[element]
	return exists
}

// Len returns the number of elements in the set.
//
// Example:
//
//	st := set.New[int]()
//	require.Equal(0, st.Len())
//	st.Insert(1)
//	require.Equal(1, st.Len())
func (s Set[T]) Len() int {
	return len(s.data)
}

// Equals checks if two sets are equal.
//
// Example:
//
//	s1 := set.New(1, 2, 3)
//
//	s2 := set.New(1, 2, 3)
//	require.True(s1.Equals(s2))
//
//	s3 := set.New(1, 2)
//	require.False(s1.Equals(s3))
func (s Set[T]) Equals(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for element := range s.data {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// IsEmpty checks if the set is empty.
//
// Example:
//
//	st := set.New[int]()
//	require.True(st.IsEmpty())
//	st.Insert(1)
//	require.False(st.IsEmpty())
func (s Set[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// ToSlice converts the set to a slice of its elements.
//
// Example:
//
//	v := []int{1, 2, 3}
//	st := set.New(v...)
//	slice := st.ToSlice()
//	require.ElementsMatch(v, slice)
func (s Set[T]) ToSlice() []T {
	elements := make([]T, 0, len(s.data))
	for element := range s.data {
		elements = append(elements, element)
	}
	return elements
}

// Clear removes all elements from the set.
//
// Example:
//
//	st := set.New(1)
//	require.Equal(1, st.Len())
//	st.Clear()
//	require.Equal(0, st.Len())
func (s *Set[T]) Clear() {
	s.data = make(map[T]exists)
}

// Copy creates a shallow copy of the set.
//
// Example:
//
//	s1 := set.New(1, 2, 3)
//	s2 := s1.Copy()
//	require.True(s2.Equals(set.New(1, 2, 3)))
func (s *Set[T]) Copy() Set[T] {
	newSet := New[T]()
	for e := range s.data {
		newSet.Insert(e)
	}
	return newSet
}

// Merge adds all elements from another set into the current set.
//
// Example:
//
//	s1 := set.New(1, 2, 3)
//	s2 := set.New(3, 4, 5)
//	s1.Merge(s2)
//	require.True(s1.Equals(set.New(1, 2, 3, 4, 5)))
func (s *Set[T]) Merge(other Set[T]) {
	for e := range other.data {
		s.Insert(e)
	}
}

// Union creates a new set that is the union of the current set and another set.
//
// Example:
//
//	s1 := set.New(1, 2, 3)
//	s2 := set.New(3, 4, 5)
//	s3 := s1.Union(s2)
//	require.True(s3.Equals(set.New(1, 2, 3, 4, 5)))
//	require.True(s1.Equals(set.New(1, 2, 3))) // s1 should be unchanged
//	require.True(s2.Equals(set.New(3, 4, 5))) // s2 should be unchanged
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := s.Copy()
	result.Merge(other)
	return result
}

// Intersection creates a new set that is the intersection of the current set and another set.
//
// Example:
//
//	s1 := set.New(1, 2, 3, 4)
//	s2 := set.New(3, 4, 5, 6)
//	s3 := s1.Intersection(s2)
//	require.True(s3.Equals(set.New(3, 4)))
//	require.True(s1.Equals(set.New(1, 2, 3, 4))) // s1 should be unchanged
//	require.True(s2.Equals(set.New(3, 4, 5, 6))) // s2 should be unchanged
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := New[T]()
	for e := range s.data {
		if other.Contains(e) {
			result.Insert(e)
		}
	}
	return result
}

// Difference creates a new set that contains elements in the current set but not in the other set.
//
// Example:
//
//	s1 := set.New(1, 2, 3, 4)
//	s2 := set.New(3, 4, 5, 6)
//
//	require.True(s1.Difference(s2).Equals(set.New(1, 2)))
//	require.True(s2.Difference(s1).Equals(set.New(5, 6)))
//	require.True(s1.Equals(set.New(1, 2, 3, 4))) // s1 should be unchanged
//	require.True(s2.Equals(set.New(3, 4, 5, 6))) // s2 should be unchanged
func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := New[T]()
	for e := range s.data {
		if !other.Contains(e) {
			result.Insert(e)
		}
	}
	return result
}

// SymmetricDifference creates a new set with elements in either set but not in both.
//
// Example:
//
//	s1 := set.New(1, 2, 3, 4)
//	s2 := set.New(3, 4, 5, 6)
//	s3 := s1.SymmetricDifference(s2)
//	require.True(s3.Equals(set.New(1, 2, 5, 6)))
//	require.True(s1.Equals(set.New(1, 2, 3, 4))) // s1 should be unchanged
//	require.True(s2.Equals(set.New(3, 4, 5, 6))) // s2 should be unchanged
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	result := New[T]()
	// Add elements from s that are not in other
	for e := range s.data {
		if !other.Contains(e) {
			result.Insert(e)
		}
	}
	// Add elements from other that are not in s
	for e := range other.data {
		if !s.Contains(e) {
			result.Insert(e)
		}
	}
	return result
}

//go:generate go run github.com/ysuzuki19/robustruct/cmd/gen/testdocgen -file=$GOFILE
