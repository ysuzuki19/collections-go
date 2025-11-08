package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/ysuzuki19/collections-go/set"
)

type Suite struct {
	suite.Suite
}

func (s *Suite) SetupTest() {
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestNew() {
	require := require.New(s.T())

	// testdoc begin New
	s1 := set.New[int]()
	s2 := set.New(1, 2, 3)
	s3 := set.New[string]()
	s4 := set.New("a", "b", "c")

	type Pos struct{ X, Y int }
	v := []Pos{{1, 2}, {3, 4}}
	s5 := set.New(v...)
	// testdoc end

	require.True(s1.Equals(set.New[int]()))
	require.True(s2.Equals(set.New(1, 2, 3)))
	require.True(s3.Equals(set.New[string]()))
	require.True(s4.Equals(set.New("a", "b", "c")))
	require.True(s5.Equals(set.New(v...)))
}

func (s *Suite) TestInsert() {
	require := require.New(s.T())
	// testdoc begin Set.Insert
	st := set.New[int]()
	st.Insert(1)
	require.True(st.Equals(set.New(1)))
	st.Insert(2, 3)
	require.True(st.Equals(set.New(1, 2, 3)))
	st.Insert(2, 4)
	require.True(st.Equals(set.New(1, 2, 3, 4)))
	// testdoc end
}

func (s *Suite) TestRemove() {
	require := require.New(s.T())
	// testdoc begin Set.Remove
	st := set.New(1, 2, 3, 4)
	st.Remove(2)
	require.True(st.Equals(set.New(1, 3, 4)))
	st.Remove(1, 3)
	require.True(st.Equals(set.New(4)))
	st.Remove(5)
	require.True(st.Equals(set.New(4)))
	st.Remove(10, 4)
	require.True(st.IsEmpty())
	// testdoc end
}

func (s *Suite) TestContains() {
	require := require.New(s.T())
	// testdoc begin Set.Contains
	st := set.New(1, 2, 3)
	require.True(st.Contains(1))
	require.True(st.Contains(2))
	require.False(st.Contains(4))
	require.False(st.Contains(0))
	// testdoc end
}

func (s *Suite) TestLen() {
	require := require.New(s.T())
	// testdoc begin Set.Len
	st := set.New[int]()
	require.Equal(0, st.Len())
	st.Insert(1)
	require.Equal(1, st.Len())
	// testdoc end
	st.Insert(2, 3, 4)
	require.Equal(4, st.Len())
	st.Remove(2)
	require.Equal(3, st.Len())
}

func (s *Suite) TestIsEmpty() {
	require := require.New(s.T())
	// testdoc begin Set.IsEmpty
	st := set.New[int]()
	require.True(st.IsEmpty())
	st.Insert(1)
	require.False(st.IsEmpty())
	// testdoc end
	st.Remove(1)
	require.True(st.IsEmpty())
}

func (s *Suite) TestToSlice() {
	require := require.New(s.T())
	// testdoc begin Set.ToSlice
	v := []int{1, 2, 3}
	st := set.New(v...)
	slice := st.ToSlice()
	require.ElementsMatch(v, slice)
	// testdoc end
}

func (s *Suite) TestClear() {
	require := require.New(s.T())
	// testdoc begin Set.Clear
	st := set.New(1)
	require.Equal(1, st.Len())
	st.Clear()
	require.Equal(0, st.Len())
	// testdoc end
}

func (s *Suite) TestCopy() {
	require := require.New(s.T())
	// testdoc begin Set.Copy
	s1 := set.New(1, 2, 3)
	s2 := s1.Copy()
	require.True(s2.Equals(set.New(1, 2, 3)))
	// testdoc end
	// modify original and check copy is unchanged
	s1.Insert(4)
	require.True(s1.Equals(set.New(1, 2, 3, 4)))
	require.True(s2.Equals(set.New(1, 2, 3)))
	// modify copy and check original is unchanged
	s2.Insert(5)
	require.True(s1.Equals(set.New(1, 2, 3, 4)))
	require.True(s2.Equals(set.New(1, 2, 3, 5)))
}

func (s *Suite) TestMerge() {
	require := require.New(s.T())
	// testdoc begin Set.Merge
	s1 := set.New(1, 2, 3)
	s2 := set.New(3, 4, 5)
	s1.Merge(s2)
	require.True(s1.Equals(set.New(1, 2, 3, 4, 5)))
	// testdoc end
	require.True(s2.Equals(set.New(3, 4, 5))) // s2 should be unchanged
}

func (s *Suite) TestUnion() {
	require := require.New(s.T())
	// testdoc begin Set.Union
	s1 := set.New(1, 2, 3)
	s2 := set.New(3, 4, 5)
	s3 := s1.Union(s2)
	require.True(s3.Equals(set.New(1, 2, 3, 4, 5)))
	require.True(s1.Equals(set.New(1, 2, 3))) // s1 should be unchanged
	require.True(s2.Equals(set.New(3, 4, 5))) // s2 should be unchanged
	// testdoc end
}

func (s *Suite) TestIntersection() {
	require := require.New(s.T())
	// testdoc begin Set.Intersection
	s1 := set.New(1, 2, 3, 4)
	s2 := set.New(3, 4, 5, 6)
	s3 := s1.Intersection(s2)
	require.True(s3.Equals(set.New(3, 4)))
	require.True(s1.Equals(set.New(1, 2, 3, 4))) // s1 should be unchanged
	require.True(s2.Equals(set.New(3, 4, 5, 6))) // s2 should be unchanged
	// testdoc end
	// test empty intersection
	s4 := set.New(1, 2)
	s5 := set.New(3, 4)
	s6 := s4.Intersection(s5)
	require.True(s6.Equals(set.New[int]()))
}

func (s *Suite) TestDifference() {
	require := require.New(s.T())
	// testdoc begin Set.Difference
	s1 := set.New(1, 2, 3, 4)
	s2 := set.New(3, 4, 5, 6)

	require.True(s1.Difference(s2).Equals(set.New(1, 2)))
	require.True(s2.Difference(s1).Equals(set.New(5, 6)))
	require.True(s1.Equals(set.New(1, 2, 3, 4))) // s1 should be unchanged
	require.True(s2.Equals(set.New(3, 4, 5, 6))) // s2 should be unchanged
	// testdoc end

	// test difference with no common elements
	s4 := set.New(1, 2)
	s5 := set.New(3, 4)
	require.True(s4.Difference(s5).Equals(set.New(1, 2)))
	require.True(s5.Difference(s4).Equals(set.New(3, 4)))
}

func (s *Suite) TestSymmetricDifference() {
	require := require.New(s.T())
	// testdoc begin Set.SymmetricDifference
	s1 := set.New(1, 2, 3, 4)
	s2 := set.New(3, 4, 5, 6)
	s3 := s1.SymmetricDifference(s2)
	require.True(s3.Equals(set.New(1, 2, 5, 6)))
	require.True(s1.Equals(set.New(1, 2, 3, 4))) // s1 should be unchanged
	require.True(s2.Equals(set.New(3, 4, 5, 6))) // s2 should be unchanged
	// testdoc end

	// test symmetric difference with no common elements
	s4 := set.New(1, 2)
	s5 := set.New(3, 4)
	s6 := s4.SymmetricDifference(s5)
	require.True(s6.Equals(set.New(1, 2, 3, 4)))
}

func (s *Suite) TestEquals() {
	require := require.New(s.T())
	// testdoc begin Set.Equals
	s1 := set.New(1, 2, 3)

	s2 := set.New(1, 2, 3)
	require.True(s1.Equals(s2))

	s3 := set.New(1, 2)
	require.False(s1.Equals(s3))
	// testdoc end

	s4 := set.New(1, 2, 3, 4)
	require.False(s1.Equals(s4))

	s5 := set.New(3, 2, 1) // different order, same elements
	require.True(s1.Equals(s5))

	s6 := set.New(1, 2, 4) // same length, different elements
	require.False(s1.Equals(s6))
}
