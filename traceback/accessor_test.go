package traceback_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/ysuzuki19/collections-go/traceback"
)

type AccessorSuite struct {
	suite.Suite
	*require.Assertions
}

func (s *AccessorSuite) SetupTest() {
	s.Assertions = require.New(s.T())
}

func TestAccessorSuite(t *testing.T) {
	suite.Run(t, new(AccessorSuite))
}

func (s *AccessorSuite) TestFramesOf_TraceError() {
	someFunc := func() error {
		return traceback.New("something went wrong")
	}
	// testdoc begin FramesOf
	err := someFunc()
	frames := traceback.FramesOf(err)

	// can use stack frames here
	_ = frames.String()

	// custom formatting
	_ = frames.Format(func(args traceback.FormatterArgs) string {
		return fmt.Sprintf("%s:%d (%s)", args.File, args.Line, args.Function)
	})
	// testdoc end
	s.Greater(frames.Len(), 0)
}

func (s *AccessorSuite) TestFramesOf_WrappedTraceError() {
	err := traceback.New("original error")
	wrapped := fmt.Errorf("wrapped: %w", err)
	frames := traceback.FramesOf(wrapped)
	s.Greater(frames.Len(), 0)
}

func (s *AccessorSuite) TestFramesOf_NonTraceError() {
	err := fmt.Errorf("regular error")
	frames := traceback.FramesOf(err)
	s.Equal(0, frames.Len())
}

func (s *AccessorSuite) TestFramesOf_Nil() {
	frames := traceback.FramesOf(nil)
	s.Equal(0, frames.Len())
}
