package frame

import "fmt"

type Formatter func(Frame) string

// defaultFormatter provides the default string representation of a Frame.
// Use for .String() method.
func defaultFormatter(f Frame) string {
	return fmt.Sprintf("%s()\n\t%s:%d", f.Function, f.File, f.Line)
}
