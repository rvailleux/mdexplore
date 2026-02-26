# mdexplore - Markdown TOC Explorer

[![CI](https://github.com/rvailleux/mdexplore/actions/workflows/ci.yml/badge.svg)](https://github.com/rvailleux/mdexplore/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/rvailleux/mdexplore)](https://goreportcard.com/report/github.com/rvailleux/mdexplore)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/go-1.23+-blue.svg)](https://golang.org)

A CLI tool that displays a hierarchical table of contents from markdown files using an interactive TUI.

<p align="center">
  <img src="docs/screenshot.png" alt="mdexplore screenshot" width="600">
</p>

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Key Bindings](#key-bindings)
- [Examples](#examples)
- [Command-Line Options](#command-line-options)
- [Exit Codes](#exit-codes)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Interactive TUI**: Navigate through your markdown document structure with an elegant Bubble Tea interface
- **Line Numbers**: Each section shows its line range (`L[start]-[end]`) for quick reference
- **Hierarchical Numbering**: Sections are numbered (1., 1.1, 1.1.1) for easy reference
- **Pre-select Sections**: Use `--select` flag to jump directly to a specific section
- **Hierarchical Navigation**: Expand/collapse sections to focus on specific parts of your document
- **Content View**: View section content directly in the terminal without leaving the tool
- **Depth Limiting**: Use `--level` flag to limit TOC depth (e.g., show only H1 and H2)
- **Multiple Heading Formats**: Supports both ATX (`# Heading`) and Setext (`Heading\n===`) style headings
- **Code Block Awareness**: Automatically excludes headings inside fenced and indented code blocks
- **Frontmatter Support**: Skips YAML frontmatter at the start of files
- **Error Handling**: Clear, styled error messages for file-related issues
- **Fast Performance**: Parses 1MB files in under 100ms

## Installation

### From Source

```bash
go install github.com/rvailleux/mdexplore/cmd/mdexplore@latest
```

### Build Locally

```bash
git clone https://github.com/rvailleux/mdexplore.git
cd mdexplore
go build -o mdexplore ./cmd/mdexplore
```

### Pre-built Binaries

Download pre-built binaries from the [releases page](https://github.com/rvailleux/mdexplore/releases).

## Usage

```bash
# Display table of contents for a markdown file
mdexplore README.md

# Limit depth to H1 and H2 headings only
mdexplore -L 2 README.md

# Show only top-level sections
mdexplore --level 1 README.md

# Pre-select a section by number (jumps directly to content)
mdexplore README.md --select 1.1

# Non-TTY mode (pipe to other tools)
mdexplore README.md > toc.txt

# Print only specific section and subsections (with --select in non-TTY mode)
mdexplore README.md --select 1.3

# Show help
mdexplore --help

# Show version
mdexplore --version
```

## Key Bindings

When the TUI is open:

| Key | Action |
|-----|--------|
| `↑` / `↓` or `k` / `j` | Navigate up/down |
| `→` | Expand section **OR** open content (if leaf section) |
| `←` | Collapse section **OR** return from content view |
| `Enter` | View full section content (including subsections) |
| `Esc` | Return from content view / Quit from TOC |
| `q` / `Ctrl+C` | Quit |

## Examples

### Table of Contents View
```
📄 README.md - Table of Contents

  L1-4      ▶ 1. Introduction
  L5-8      ├── 1.1. Installation
  L9-14     │   └── 1.1.1. Requirements
  L15-18    │   └── 1.1.2. Setup
  L19-20    ├── 1.2. Usage
  L21-24    │   └── 1.2.1. Basic Commands
  L25-28    │   └── 1.2.2. Advanced Options
  L29-31    └── 1.3. License

[↑/↓] Navigate  [→] Expand/Open  [←] Collapse/Back  [Enter] Full Content  [q] Quit
```

### Content View
```
📄 README.md - L5-8: Installation

  ## Installation

  To install mdexplore, run:
  go install ./cmd/mdexplore

[Esc] Return to navigation  [q] Quit
```

## Command-Line Options

| Flag | Shorthand | Description |
|------|-----------|-------------|
| `--level N` | `-L N` | Maximum heading level to display (0 = no limit) |
| `--select N` | | Pre-select section by number (e.g., 1.1) |
| `--toc` | | Display table of contents (default: true) |
| `--help` | `-h` | Display help information |
| `--version` | `-v` | Display version information |

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | File not found |
| 3 | Permission denied |
| 4 | Invalid file (directory instead of file) |
| 5 | Parse error |

## Development

### Prerequisites

- Go 1.23 or later
- Git

### Building

```bash
# Clone the repository
git clone https://github.com/rvailleux/mdexplore.git
cd mdexplore

# Build
go build -o mdexplore ./cmd/mdexplore

# Install locally
go install ./cmd/mdexplore
```

### Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run benchmarks
go test ./tests/benchmark -bench=.
```

### Project Structure

```
cmd/mdexplore/       # CLI entry point
internal/
  models/            # Data models (Section, SectionTree)
  parser/            # Markdown parsing (Goldmark)
  ui/                # TUI components (Bubble Tea)
  renderer/          # Content rendering
  errors/            # Error types
tests/
  integration/       # Integration tests
  benchmark/         # Performance tests
  fixtures/          # Test data
```

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

- [Code of Conduct](CODE_OF_CONDUCT.md)
- [Security Policy](SECURITY.md)
- [Changelog](CHANGELOG.md)

### Quick Start for Contributors

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (`go test ./...`)
5. Commit your changes (`git commit -m 'feat: add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) - The TUI framework
- Markdown parsing powered by [Goldmark](https://github.com/yuin/goldmark)
- Styled with [Lipgloss](https://github.com/charmbracelet/lipgloss)
