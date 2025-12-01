package frame

import (
	"runtime"
)

// Capture captures the current stack frames, skipping the specified number of frames.
//
// Example:
//
//	frames := frame.Capture(0)
func Capture(skip int) Frames {
	var frames Frames

	pcs := make([]uintptr, 32)
	n := runtime.Callers(skip+1, pcs)
	pcs = pcs[:n]

	callersFrames := runtime.CallersFrames(pcs)
	for {
		callersFrame, more := callersFrames.Next()
		frames.Push(Frame{
			Function: callersFrame.Function,
			File:     callersFrame.File,
			Line:     callersFrame.Line,
		})
		if !more {
			break
		}
	}
	return frames
}

//go:generate go run github.com/ysuzuki19/robustruct/cmd/gen/testdocgen -file=$GOFILE
