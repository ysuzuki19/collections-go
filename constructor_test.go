package collections_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ysuzuki19/collections-go"
	"github.com/ysuzuki19/collections-go/set"
)

func TestSet(t *testing.T) {
	s := collections.NewSet(1, 2, 3)
	expected := set.New(1, 2, 3)
	require.Equal(t, expected, s)
}
