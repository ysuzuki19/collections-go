package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/ysuzuki19/collections-go/set"
)

func Equals[T comparable](s set.Set[T], slice []T) bool {
	if s.Len() != len(slice) {
		return false
	}
	for _, v := range slice {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}

type Suite struct {
	suite.Suite
	require *require.Assertions
}

func (s *Suite) SetupTest() {
	s.require = require.New(s.T())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestNew() {
	s.require.True(Equals(set.New[int](), []int{}))
	s.require.True(Equals(set.New(1, 2, 3), []int{1, 2, 3}))

	s.require.True(Equals(set.New[string](), []string{}))
	s.require.True(Equals(set.New("a", "b", "c"), []string{"a", "b", "c"}))

	v := struct{ X int }{X: 10}
	s.require.True(Equals(set.New(v), []struct{ X int }{v}))
}

func (s *Suite) TestInsert() {
	st := set.New[int]()
	// single insert
	st.Insert(1)
	s.require.True(Equals(st, []int{1}))
	// multiple insert
	st.Insert(2, 3)
	s.require.True(Equals(st, []int{1, 2, 3}))
	// insert existing and new
	st.Insert(2, 4)
	s.require.True(Equals(st, []int{1, 2, 3, 4}))
}

func (s *Suite) TestRemove() {
	st := set.New(1, 2, 3, 4)
	// single remove
	st.Remove(2)
	s.require.True(Equals(st, []int{1, 3, 4}))
	// multiple remove
	st.Remove(1, 3)
	s.require.True(Equals(st, []int{4}))
	// remove non-existing
	st.Remove(5)
	s.require.True(Equals(st, []int{4}))
	// remove existing and non-existing
	st.Remove(10, 4)
	s.require.True(Equals(st, []int{}))
}

func (s *Suite) TestContains() {
	st := set.New(1, 2, 3)
	s.require.True(st.Contains(1), "expected set to contain 1")
	s.require.True(st.Contains(2), "expected set to contain 2")
	s.require.False(st.Contains(4), "expected set not to contain 4")
	s.require.False(st.Contains(0), "expected set not to contain 0")
}

func (s *Suite) TestLen() {
	st := set.New[int]()
	s.require.Equal(0, st.Len())
	st.Insert(1)
	s.require.Equal(1, st.Len())
	st.Insert(2, 3, 4)
	s.require.Equal(4, st.Len())
	st.Remove(2)
	s.require.Equal(3, st.Len())
}

func (s *Suite) TestIsEmpty() {
	st := set.New[int]()
	s.require.True(st.IsEmpty(), "expected set to be empty")
	st.Insert(1)
	s.require.False(st.IsEmpty(), "expected set not to be empty")
	st.Remove(1)
	s.require.True(st.IsEmpty(), "expected set to be empty after removing all elements")
}

func (s *Suite) TestToSlice() {
	st := set.New(1, 2, 3)
	slice := st.ToSlice()
	s.require.Len(slice, 3, "expected slice length 3")
	// check all elements are present (order not guaranteed)
	s.require.True(Equals(st, []int{1, 2, 3}))
}

func (s *Suite) TestClear() {
	st := set.New(1, 2, 3, 4)
	s.require.True(Equals(st, []int{1, 2, 3, 4}))
	st.Clear()
	s.require.True(Equals(st, []int{}))
}

func (s *Suite) TestCopy() {
	s1 := set.New(1, 2, 3)
	s2 := s1.Copy()
	s.require.True(Equals(s2, []int{1, 2, 3}))
	// modify original and check copy is unchanged
	s1.Insert(4)
	s.require.True(Equals(s1, []int{1, 2, 3, 4}))
	s.require.True(Equals(s2, []int{1, 2, 3}))
	// modify copy and check original is unchanged
	s2.Insert(5)
	s.require.True(Equals(s1, []int{1, 2, 3, 4}))
	s.require.True(Equals(s2, []int{1, 2, 3, 5}))
}

func (s *Suite) TestMerge() {
	s1 := set.New(1, 2, 3)
	s2 := set.New(3, 4, 5)
	s1.Merge(s2)
	s.require.True(Equals(s1, []int{1, 2, 3, 4, 5}))
	s.require.True(Equals(s2, []int{3, 4, 5})) // s2 should be unchanged
}

func (s *Suite) TestUnion() {
	s1 := set.New(1, 2, 3)
	s2 := set.New(3, 4, 5)
	s3 := s1.Union(s2)
	s.require.True(Equals(s3, []int{1, 2, 3, 4, 5}))
	s.require.True(Equals(s1, []int{1, 2, 3})) // s1 should be unchanged
	s.require.True(Equals(s2, []int{3, 4, 5})) // s2 should be unchanged
}

func (s *Suite) TestIntersection() {
	s1 := set.New(1, 2, 3, 4)
	s2 := set.New(3, 4, 5, 6)
	s3 := s1.Intersection(s2)
	s.require.True(Equals(s3, []int{3, 4}))
	s.require.True(Equals(s1, []int{1, 2, 3, 4})) // s1 should be unchanged
	s.require.True(Equals(s2, []int{3, 4, 5, 6})) // s2 should be unchanged
	// test empty intersection
	s4 := set.New(1, 2)
	s5 := set.New(3, 4)
	s6 := s4.Intersection(s5)
	s.require.True(Equals(s6, []int{}))
}

func (s *Suite) TestDifference() {
	s1 := set.New(1, 2, 3, 4)
	s2 := set.New(3, 4, 5, 6)
	s.require.True(Equals(s1.Difference(s2), []int{1, 2}))
	s.require.True(Equals(s2.Difference(s1), []int{5, 6}))
	s.require.True(Equals(s1, []int{1, 2, 3, 4})) // s1 should be unchanged
	s.require.True(Equals(s2, []int{3, 4, 5, 6})) // s2 should be unchanged
	// test difference with no common elements
	s4 := set.New(1, 2)
	s5 := set.New(3, 4)
	s.require.True(Equals(s4.Difference(s5), []int{1, 2}))
	s.require.True(Equals(s5.Difference(s4), []int{3, 4}))
}

func (s *Suite) TestSymmetricDifference() {
	s1 := set.New(1, 2, 3, 4)
	s2 := set.New(3, 4, 5, 6)
	s3 := s1.SymmetricDifference(s2)
	s.require.True(Equals(s3, []int{1, 2, 5, 6}))
	s.require.True(Equals(s1, []int{1, 2, 3, 4})) // s1 should be unchanged
	s.require.True(Equals(s2, []int{3, 4, 5, 6})) // s2 should be unchanged
	// test symmetric difference with no common elements
	s4 := set.New(1, 2)
	s5 := set.New(3, 4)
	s6 := s4.SymmetricDifference(s5)
	s.require.True(Equals(s6, []int{1, 2, 3, 4}))
}

func (s *Suite) TestEquals() {
	s1 := set.New(1, 2, 3)
	s2 := set.New(1, 2, 3)
	s3 := set.New(1, 2)
	s4 := set.New(1, 2, 3, 4)
	s5 := set.New(3, 2, 1) // different order, same elements
	s6 := set.New(1, 2, 4) // same length, different elements

	s.require.True(s1.Equals(s2), "sets with same elements should be equal")
	s.require.True(s1.Equals(s5), "sets with same elements in different order should be equal")
	s.require.False(s1.Equals(s3), "sets with different length should not be equal")
	s.require.False(s1.Equals(s4), "sets with different elements should not be equal")
	s.require.False(s1.Equals(s6), "sets with same length, different elements should not be equal")
}
