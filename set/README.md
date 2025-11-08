# Set

A generic, type-safe Set implementation for Go using generics.

## Features

- **Generic**: Works with any comparable type (int, string, custom structs, etc.)
- **Type-safe**: Compile-time type checking
- **Rich API**: Common set operations including union, intersection, difference, and more

## Usage

### Creating a Set

```go
import "github.com/ysuzuki19/collections-go/set"

// Empty set (necessary to specify type)
s1 := set.New[int]()

// Set with initial elements
s2 := set.New(1, 2, 3)

// String set
s3 := set.New("apple", "banana", "cherry")
```

### Basic Operations

```go
s := set.New[int]()

// Insert elements
s.Insert(1)
s.Insert(2, 3, 4)

// Check if element exists
if s.Contains(2) {
    fmt.Println("Set contains 2")
}

// Remove elements
s.Remove(3)
s.Remove(1, 4)

// Get size
fmt.Println("Size:", s.Len())

// Check if empty
if s.IsEmpty() {
    fmt.Println("Set is empty")
}

// Clear all elements
s.Clear()
```

### Set Operations

```go
s1 := set.New(1, 2, 3, 4)
s2 := set.New(3, 4, 5, 6)

// Union: elements in either set
union := s1.Union(s2)  // {1, 2, 3, 4, 5, 6}

// Intersection: elements in both sets
intersection := s1.Intersection(s2)  // {3, 4}

// Difference: elements in s1 but not in s2
diff := s1.Difference(s2)  // {1, 2}

// Symmetric Difference: elements in either set but not in both
symDiff := s1.SymmetricDifference(s2)  // {1, 2, 5, 6}
```

### Copying and Merging

```go
s1 := set.New(1, 2, 3)

// Create a shallow copy
s2 := s1.Copy()

// Merge another set into current set (modifies s1)
s3 := set.New(4, 5)
s1.Merge(s3)  // s1 is now {1, 2, 3, 4, 5}
```

### Comparison and Conversion

```go
s1 := set.New(1, 2, 3)
s2 := set.New(3, 2, 1)

// Check equality (order doesn't matter)
if s1.Equals(s2) {
    fmt.Println("Sets are equal")
}

// Convert to slice
slice := s1.ToSlice()  // []int{1, 2, 3} (order not guaranteed)
```

## Performance

- **Insert**: O(1) average case
- **Remove**: O(1) average case
- **Contains**: O(1) average case
- **Union**: O(n + m) where n and m are the sizes of the sets
- **Intersection**: O(min(n, m))
- **Difference**: O(n) where n is the size of the current set
- **SymmetricDifference**: O(n + m)
