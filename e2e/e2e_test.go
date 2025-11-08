package e2e_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ysuzuki19/collections-go"
)

func TestSet(t *testing.T) {
	s := collections.NewSet(1, 2, 3)
	v := s.ToSlice()
	expected := []int{1, 2, 3}
	require.Equal(t, expected, v)
}
