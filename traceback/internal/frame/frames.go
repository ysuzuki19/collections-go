package frame

import (
	"strings"
)

// Frames represents a collection of stack frames.
type Frames struct {
	frames []Frame
}

// String returns a formatted string representation of the frames.
//
// Example:
//
//	frames := frame.Frames{}
//	s.Equal("", frames.String())
func (fs Frames) String() string {
	var sb strings.Builder
	for _, f := range fs.frames {
		sb.WriteString(f.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

// Format returns a formatted string representation of the frames using the given layout.
//
// Example:
//
//	frames := frame.Frames{}
//	frames.Push(frame.Frame{File: "file1.go", Line: 10, Function: "func1"})
//	frames.Push(frame.Frame{File: "file2.go", Line: 20, Function: "func2"})
//	formatter := func(args frame.Frame) string {
//		return args.Function + " at " + args.File + ":" + fmt.Sprint(args.Line)
//	}
//	s.Equal("func1 at file1.go:10\nfunc2 at file2.go:20\n", frames.Format(formatter))
func (fs Frames) Format(formatter Formatter) string {
	var sb strings.Builder
	for _, f := range fs.frames {
		sb.WriteString(f.Format(formatter))
		sb.WriteString("\n")
	}
	return sb.String()
}

// Push adds a new frame to the collection.
//
// Example:
//
//	frames := frame.Frames{}
//	frames.Push(frame.Frame{File: "file1.go", Line: 10, Function: "func1"})
//	s.Equal(1, frames.Len())
func (fs *Frames) Push(other Frame) {
	fs.frames = append(fs.frames, other)
}

// Len returns the number of frames in the collection.
//
// Example:
//
//	frames := frame.Frames{}
//	s.Equal(0, frames.Len())
//	frames.Push(frame.Frame{File: "file1.go", Line: 10, Function: "func1"})
//	s.Equal(1, frames.Len())
func (fs Frames) Len() int {
	return len(fs.frames)
}

//go:generate go run github.com/ysuzuki19/robustruct/cmd/gen/testdocgen -file=$GOFILE
