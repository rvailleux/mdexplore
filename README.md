# mdexplore - Markdown TOC Explorer

A CLI tool that displays a hierarchical table of contents from markdown files using an interactive TUI.

## Features

- **Interactive TUI**: Navigate through your markdown document structure with an elegant Bubble Tea interface
- **Multiple Heading Formats**: Supports both ATX (`# Heading`) and Setext (`Heading\n===`) style headings
- **Code Block Awareness**: Automatically excludes headings inside fenced and indented code blocks
- **Frontmatter Support**: Skips YAML frontmatter at the start of files
- **Error Handling**: Clear, styled error messages for file-related issues
- **Fast Performance**: Parses 1MB files in under 100ms

## Installation

```bash
go build -o mdexplore ./cmd/mdexplore
go install ./cmd/mdexplore
```

## Usage

```bash
# Display table of contents for a markdown file
mdexplore README.md --toc

# Display table of contents (default behavior)
mdexplore README.md

# Show help
mdexplore --help

# Show version
mdexplore --version
```

## Key Bindings

When the TUI is open:

- `↑` / `↓` or `k` / `j` - Navigate up/down
- `q` / `Ctrl+C` / `Esc` - Quit

## Example Output

```
📄 README.md - Table of Contents

  ● Introduction
  ├── Installation
  │   ├── Requirements
  │   └── Setup
  ├── Usage
  │   ├── Basic Commands
  │   └── Advanced Options
  └── License

[↑/↓] Navigate  [q] Quit
```

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

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run benchmarks
go test ./tests/benchmark -bench=.

# Build
go build -o mdexplore ./cmd/mdexplore
```

## Requirements

- Go 1.22+
- Terminal with UTF-8 and color support

## License

MIT
