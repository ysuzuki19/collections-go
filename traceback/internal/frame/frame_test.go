package frame_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/ysuzuki19/collections-go/traceback/internal/frame"
)

type FrameSuite struct {
	suite.Suite
	*require.Assertions
}

func (s *FrameSuite) SetupTest() {
	s.Assertions = require.New(s.T())
}

func TestFrameSuite(t *testing.T) {
	suite.Run(t, new(FrameSuite))
}

func (s *FrameSuite) TestString() {
	// testdoc begin Frame.String
	f := frame.Frame{
		Function: "main.doSomething",
		File:     "/path/to/file.go",
		Line:     42,
	}
	s.Equal("main.doSomething()\n\t/path/to/file.go:42", f.String())
	// testdoc end
}

func (s *FrameSuite) TestFormat() {
	// testdoc begin Frame.Format
	f := frame.Frame{
		Function: "main.doSomething",
		File:     "/path/to/file.go",
		Line:     42,
	}
	formatter := func(args frame.Frame) string {
		return args.Function + " at " + args.File + ":" + fmt.Sprint(args.Line)
	}
	s.Equal("main.doSomething at /path/to/file.go:42", f.Format(formatter))
	// testdoc end
}
