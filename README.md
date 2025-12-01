# collections-go

[![Go Reference](https://pkg.go.dev/badge/github.com/ysuzuki19/collections-go.svg)](https://pkg.go.dev/github.com/ysuzuki19/collections-go)

A modern Go collections library with generic data structures.

## Installation

```bash
go get github.com/ysuzuki19/collections-go
```

## Available Collections

- Set - A generic set implementation with common set operations
  - [README](./set/README.md)
  - [![Docs](https://pkg.go.dev/badge/github.com/ysuzuki19/collections-go/set.svg)](https://pkg.go.dev/github.com/ysuzuki19/collections-go/set)

## Available Utilities

- Traceback - A lightweight Go package for capturing stack traces with errors
  - [README](./traceback/README.md)
  - [![Docs](https://pkg.go.dev/badge/github.com/ysuzuki19/collections-go/traceback.svg)](https://pkg.go.dev/github.com/ysuzuki19/collections-go/traceback)

## Quick Start

### Set

You can create collections using the top-level constructors:

```go
import "github.com/ysuzuki19/collections-go"

// Create a set
s := collections.NewSet()
s := collections.NewSet(1, 2, 3, 4)
```

Or import specific packages directly:

```go
import "github.com/ysuzuki19/collections-go/set"

// Create a set
s := set.New()
s := set.New(1, 2, 3, 4)
```
