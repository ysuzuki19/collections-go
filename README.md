# collections-go

A modern Go collections library with generic data structures.

## Installation

```bash
go get github.com/ysuzuki19/collections-go
```

## Available Collections

- **[Set](./set/README.md)** - A generic set implementation with common set operations

## Quick Start

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

## Documentation

For detailed documentation on each collection type and available methods, please refer to the individual package READMEs:

- **Set**: See [set/README.md](./set/README.md) for complete API reference and usage examples

## Requirements

- Go 1.18 or higher (for generics support)
