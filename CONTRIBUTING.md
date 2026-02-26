# Contributing to mdexplore

Thank you for your interest in contributing to mdexplore! We welcome contributions from the community and are pleased to have you join us.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [How to Contribute](#how-to-contribute)
  - [Reporting Bugs](#reporting-bugs)
  - [Suggesting Features](#suggesting-features)
  - [Pull Requests](#pull-requests)
- [Development Setup](#development-setup)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Commit Message Guidelines](#commit-message-guidelines)

## Code of Conduct

This project and everyone participating in it is governed by our [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally
3. Create a new branch for your contribution
4. Make your changes
5. Test your changes
6. Submit a pull request

## How to Contribute

### Reporting Bugs

Before creating a bug report, please:

1. Check the [existing issues](https://github.com/rvailleux/mdexplore/issues) to see if the bug has already been reported
2. Update to the latest version to verify the bug still exists

When reporting a bug, please include:

- **Clear title and description**
- **Steps to reproduce** the issue
- **Expected behavior** vs **actual behavior**
- **Environment details**: OS, Go version, terminal emulator
- **Sample markdown file** that triggers the issue (if applicable)
- **Screenshots** if the issue is UI-related

### Suggesting Features

Feature requests are welcome! Please provide:

- **Clear use case**: What problem does this solve?
- **Detailed description**: How should it work?
- **Mockups or examples** (if applicable)

### Pull Requests

1. **Create a new branch** from `main`:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following our coding standards

3. **Test your changes**:
   ```bash
   go test ./...
   ```

4. **Commit your changes** with a clear commit message

5. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```

6. **Open a Pull Request** against the `main` branch

## Development Setup

### Prerequisites

- Go 1.23 or later
- Git
- Make (optional)

### Building

```bash
# Clone the repository
git clone https://github.com/rvailleux/mdexplore.git
cd mdexplore

# Build the binary
go build -o mdexplore ./cmd/mdexplore

# Install locally
go install ./cmd/mdexplore
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run benchmarks
go test ./tests/benchmark -bench=.
```

## Coding Standards

### Go Code Style

- Follow standard Go conventions (gofmt, goimports)
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions focused and small
- Handle errors explicitly

### Example:

```go
// ParseSectionTree reads a markdown file and builds a hierarchical section tree.
func (p *GoldmarkParser) ParseSectionTree(filepath string) (*models.SectionTree, error) {
    // Implementation
}
```

### Project Structure

```
cmd/mdexplore/       # CLI entry point
internal/
  models/            # Data models
  parser/            # Markdown parsing
  ui/                # TUI components
  renderer/          # Content rendering
  errors/            # Error types
tests/
  integration/       # Integration tests
  benchmark/         # Performance tests
  fixtures/          # Test data
```

## Testing

We require tests for all new functionality:

- **Unit tests**: Test individual functions and methods
- **Integration tests**: Test end-to-end workflows
- **Benchmarks**: For performance-critical code

### Test-Driven Development

This project follows TDD principles:

1. Write a failing test
2. Write minimal code to pass the test
3. Refactor while keeping tests green

### Example Test:

```go
func TestSectionContainsLine(t *testing.T) {
    section := &Section{StartLine: 10, EndLine: 50}

    if !section.ContainsLine(30) {
        t.Error("Expected ContainsLine(30) to be true")
    }
}
```

## Commit Message Guidelines

We follow [Conventional Commits](https://www.conventionalcommits.org/):

### Format

```
<type>: <description>

[optional body]

[optional footer]
```

### Types

- **feat**: New feature
- **fix**: Bug fix
- **docs**: Documentation changes
- **test**: Adding or updating tests
- **refactor**: Code refactoring
- **perf**: Performance improvements
- **chore**: Maintenance tasks

### Examples

```
feat: add line number display to TOC

fix: handle empty sections correctly
docs: update README with new navigation keys
test: add parser tests for Setext headings
refactor: simplify SectionTree building logic
```

## Questions?

Feel free to:
- Open an issue for questions
- Start a discussion on GitHub
- Reach out to maintainers

Thank you for contributing! 🎉
