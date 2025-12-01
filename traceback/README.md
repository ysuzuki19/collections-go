# traceback

A lightweight Go package for capturing stack traces with errors.

## Usage

### Creating errors with stack traces

```go
// Create a new error
err := traceback.New("something went wrong")

// Create with formatted message
err := traceback.Errorf("failed to process item %d", 42)
```

### Wrapping existing errors

```go
result, err := someOperation()
if err != nil {
    // Add stack trace to existing error
    err = traceback.From(err)

    // Or wrap with additional context
    err = traceback.Wrap(err, "failed to complete operation")

    // Or wrap with formatted context
    err = traceback.Wrapf(err, "failed to fetch user %d", userID)
}
```

### Accessing stack frames

```go
// Extract frames from error interface (works with errors.As)
func doSomething() error {
    return traceback.New("something went wrong")
}

err := doSomething()  // type is error, not *traceback.Error
frames := traceback.FramesOf(err)
fmt.Println(frames.String())
```

### Custom formatting

```go
frames := traceback.FramesOf(err)

output := frames.Format(func(f traceback.FormatterArgs) string {
    return fmt.Sprintf("%s:%d (%s)", f.File, f.Line, f.Function)
})
```

## API

| Function                      | Description                               |
| ----------------------------- | ----------------------------------------- |
| `New(message)`                | Create a new error with stack trace       |
| `Errorf(format, args...)`     | Create a new error with formatted message |
| `From(err)`                   | Wrap an existing error with stack trace   |
| `Wrap(err, message)`          | Wrap with additional context message      |
| `Wrapf(err, format, args...)` | Wrap with formatted context message       |
| `FramesOf(err)`               | Extract stack frames from any error       |

## Features

- Lightweight with no external dependencies
- Compatible with Go 1.13+ error wrapping (`errors.As`, `errors.Is`)
- Customizable stack trace formatting
- Nil-safe: `From`, `Wrap`, and `Wrapf` return nil when given nil error
