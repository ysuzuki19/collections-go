package traceback

import (
	"errors"

	"github.com/ysuzuki19/collections-go/traceback/internal/frame"
)

// FramesOf extracts the stack frames from the given error.
// If the error is not a traceback.Error, it returns an empty slice.
//
// Example:
//
//	err := someFunc()
//	frames := traceback.FramesOf(err)
//
//	// can use stack frames here
//	_ = frames.String()
//
//	// custom formatting
//	_ = frames.Format(func(args traceback.FormatterArgs) string {
//		return fmt.Sprintf("%s:%d (%s)", args.File, args.Line, args.Function)
//	})
func FramesOf(err error) frame.Frames {
	var te *Error
	if errors.As(err, &te) {
		return te.frames
	}
	return frame.Frames{} // fallback to empty frames
}

//go:generate go run github.com/ysuzuki19/robustruct/cmd/gen/testdocgen -file=$GOFILE
