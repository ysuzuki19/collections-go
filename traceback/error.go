package traceback

import (
	"errors"
	"fmt"

	"github.com/ysuzuki19/collections-go/traceback/internal/frame"
)

// Error is an error type that captures a stack trace.
type Error struct {
	cause  error
	frames frame.Frames
}

var _ error = (*Error)(nil)

// New creates a new Error with the given message.
//
// Example:
//
//	err := traceback.New("something went wrong")
//	fmt.Println(err.Error())
func New(message string) *Error {
	return &Error{
		cause:  errors.New(message),
		frames: frame.Capture(2),
	}
}

// Errorf creates a new Error with a formatted message.
//
// Example:
//
//	err := traceback.Errorf("failed to process item %d", 42)
//	fmt.Println(err.Error())
func Errorf(format string, args ...any) *Error {
	return &Error{
		cause:  fmt.Errorf(format, args...),
		frames: frame.Capture(2),
	}
}

// From wraps an existing error with a Error, adding stack trace information.
// Returns nil if err is nil.
//
// Example:
//
//	someOperation := func() (any, error) { return nil, fmt.Errorf("some error") }
//	result, err := someOperation()
//	if err != nil {
//		err = traceback.From(err)
//	}
func From(err error) *Error {
	if err == nil {
		return nil
	}
	return &Error{
		cause:  err,
		frames: frame.Capture(2),
	}
}

// Wrap wraps an existing error with a Error and additional context message.
// Returns nil if err is nil.
//
// Example:
//
//	someOperation := func() (any, error) { return nil, fmt.Errorf("some error") }
//	result, err := someOperation()
//	if err != nil {
//		err = traceback.Wrap(err, "failed to complete operation")
//	}
func Wrap(err error, message string) *Error {
	if err == nil {
		return nil
	}
	return &Error{
		cause:  fmt.Errorf("%s: %w", message, err),
		frames: frame.Capture(2),
	}
}

// Wrapf wraps an existing error with a formatted context message.
// Returns nil if err is nil.
//
// Example:
//
//	fetchUser := func(userID int) (any, error) {
//		return nil, fmt.Errorf("user not found userID=%d", userID)
//	}
//	userID := 123
//	result, err := fetchUser(userID)
//	if err != nil {
//		err = traceback.Wrapf(err, "failed to fetch user %d", userID)
//	}
func Wrapf(err error, format string, args ...any) *Error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(format, args...)
	return &Error{
		cause:  fmt.Errorf("%s: %w", msg, err),
		frames: frame.Capture(2),
	}
}

// Error returns the error message.
func (e *Error) Error() string {
	if e.cause != nil {
		return e.cause.Error()
	}
	return ""
}

// Frames returns the captured stack frames.
func (e *Error) Frames() frame.Frames {
	return e.frames
}

// String returns a formatted stack trace string.
//
// Example:
//
//	err := traceback.New("something went wrong")
//	fmt.Println(err.String())
func (e *Error) String() string {
	return e.frames.String()
}

//go:generate go run github.com/ysuzuki19/robustruct/cmd/gen/testdocgen -file=$GOFILE
