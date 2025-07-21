# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go library implementing a generic Set data structure using Go's generics and iterators. The implementation is based on Python's set operations and uses an underlying map[T]struct{} as the storage mechanism.

**Key Architecture Points:**
- Single-file library (`set.go`) with comprehensive test coverage (`set_test.go`)
- Generic type `Set[T comparable]` where T must be comparable (supports ==, != operations)
- Python-inspired API with Union, Intersection, Difference, and SymmetricDifference operations
- Iterator support using Go's `iter.Seq[T]` for integration with modern Go patterns
- Not goroutine-safe (similar to Go's map type)

## Development Commands

This project uses `just` as a make replacement. Common commands:

### Core Development Workflow
```bash
# Install required Go tools
just tools

# Run full quality assurance (mod tidy, vet, format, lint, vulnerability check)
just qa

# Run tests with coverage report (generates coverage.html)
just unittest

# Run benchmarks and compare with previous results
just benchmark
```

### Additional Commands
```bash
# Start documentation server at localhost:8080
just doc

# Publish package to Go proxy
just publish
```

## Testing and Quality

- Tests use example-based testing with `Example` functions for documentation
- 100% test coverage is maintained (see coverage.html)
- golangci-lint with custom configuration (.golangci.yml) enforces code quality
- Benchmark tests compare performance against previous runs
- Vulnerability scanning with govulncheck

## Code Style

- Go 1.24+ required
- Line length limit: 100 characters
- Function length limit: 80 lines / 30 statements
- UK English spelling in comments
- Apache 2.0 license headers required