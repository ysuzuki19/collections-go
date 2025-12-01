package frame

// Frame represents a single stack frame.
type Frame struct {
	Function string
	File     string
	Line     int
}

// Format returns a formatted string representation of the frame using the given layout.
//
// Example:
//
//	f := frame.Frame{
//		Function: "main.doSomething",
//		File:     "/path/to/file.go",
//		Line:     42,
//	}
//	formatter := func(args frame.Frame) string {
//		return args.Function + " at " + args.File + ":" + fmt.Sprint(args.Line)
//	}
//	s.Equal("main.doSomething at /path/to/file.go:42", f.Format(formatter))
func (f Frame) Format(formatter Formatter) string {
	return formatter(f)
}

// String returns a formatted string representation of the frame.
// Format follows Go's standard stack trace style:
//
// Example:
//
//	f := frame.Frame{
//		Function: "main.doSomething",
//		File:     "/path/to/file.go",
//		Line:     42,
//	}
//	s.Equal("main.doSomething()\n\t/path/to/file.go:42", f.String())
func (f Frame) String() string {
	return f.Format(defaultFormatter)
}

//go:generate go run github.com/ysuzuki19/robustruct/cmd/gen/testdocgen -file=$GOFILE
