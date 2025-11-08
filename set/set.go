package set

// exists is an empty struct used to represent the presence of an element in the set.
type exists struct{}

// Set is a generic set implementation using a map for storage.
type Set[T comparable] struct {
	data map[T]exists
}

// New creates a new Set and initializes it with the provided elements.
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
func (s *Set[T]) Insert(elements ...T) {
	for _, element := range elements {
		s.data[element] = exists{}
	}
}

// Remove deletes one or more elements from the set.
func (s *Set[T]) Remove(elements ...T) {
	for _, element := range elements {
		delete(s.data, element)
	}
}

// Contains checks if the set contains a specific element.
func (s Set[T]) Contains(element T) bool {
	_, exists := s.data[element]
	return exists
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s.data)
}

// Equals checks if two sets are equal.
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
func (s Set[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// ToSlice converts the set to a slice of its elements.
func (s Set[T]) ToSlice() []T {
	elements := make([]T, 0, len(s.data))
	for element := range s.data {
		elements = append(elements, element)
	}
	return elements
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
	s.data = make(map[T]exists)
}

// Copy creates a shallow copy of the set.
func (s *Set[T]) Copy() Set[T] {
	newSet := New[T]()
	for e := range s.data {
		newSet.Insert(e)
	}
	return newSet
}

// Merge adds all elements from another set into the current set.
func (s *Set[T]) Merge(other Set[T]) {
	for e := range other.data {
		s.Insert(e)
	}
}

// Union creates a new set that is the union of the current set and another set.
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := s.Copy()
	result.Merge(other)
	return result
}

// Intersection creates a new set that is the intersection of the current set and another set.
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
