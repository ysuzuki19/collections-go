package frame_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/ysuzuki19/collections-go/traceback/internal/frame"
)

type FramesSuite struct {
	suite.Suite
	*require.Assertions
}

func (s *FramesSuite) SetupTest() {
	s.Assertions = require.New(s.T())
}

func TestFramesSuite(t *testing.T) {
	suite.Run(t, new(FramesSuite))
}

func (s *FramesSuite) TestString() {
	{
		// testdoc begin Frames.String
		frames := frame.Frames{}
		s.Equal("", frames.String())
		// testdoc end
		frames.Push(frame.Frame{File: "file1.go", Line: 10, Function: "func1"})
		s.Equal("func1()\n\tfile1.go:10\n", frames.String())
		frames.Push(frame.Frame{File: "file2.go", Line: 20, Function: "func2"})
		s.Equal("func1()\n\tfile1.go:10\nfunc2()\n\tfile2.go:20\n", frames.String())
	}
}

func (s *FramesSuite) TestPush() {
	{
		// testdoc begin Frames.Push
		frames := frame.Frames{}
		frames.Push(frame.Frame{File: "file1.go", Line: 10, Function: "func1"})
		s.Equal(1, frames.Len())
		// testdoc end
	}
}

func (s *FramesSuite) TestLen() {
	{
		// testdoc begin Frames.Len
		frames := frame.Frames{}
		s.Equal(0, frames.Len())
		frames.Push(frame.Frame{File: "file1.go", Line: 10, Function: "func1"})
		s.Equal(1, frames.Len())
		// testdoc end
		frames.Push(frame.Frame{File: "file2.go", Line: 20, Function: "func2"})
		s.Equal(2, frames.Len())
	}
}

func (s *FramesSuite) TestFormat() {
	// testdoc begin Frames.Format
	frames := frame.Frames{}
	frames.Push(frame.Frame{File: "file1.go", Line: 10, Function: "func1"})
	frames.Push(frame.Frame{File: "file2.go", Line: 20, Function: "func2"})
	formatter := func(args frame.Frame) string {
		return args.Function + " at " + args.File + ":" + fmt.Sprint(args.Line)
	}
	s.Equal("func1 at file1.go:10\nfunc2 at file2.go:20\n", frames.Format(formatter))
	// testdoc end
}
