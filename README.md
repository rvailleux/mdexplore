# mdexplore - Markdown TOC Explorer

A CLI tool that displays a hierarchical table of contents from markdown files using an interactive TUI.

## Features

- **Interactive TUI**: Navigate through your markdown document structure with an elegant Bubble Tea interface
- **Line Numbers**: Each section shows its line range (`L[start]-[end]`) for quick reference
- **Hierarchical Navigation**: Expand/collapse sections to focus on specific parts of your document
- **Content View**: View section content directly in the terminal without leaving the tool
- **Depth Limiting**: Use `--level` flag to limit TOC depth (e.g., show only H1 and H2)
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
mdexplore README.md

# Limit depth to H1 and H2 headings only
mdexplore -L 2 README.md

# Show only top-level sections
mdexplore --level 1 README.md

# Non-TTY mode (pipe to other tools)
mdexplore README.md > toc.txt

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
| `→` | Expand selected section (reveals subsections) |
| `←` | Collapse selected section |
| `Enter` | View section content |
| `Esc` | Return from content view / Quit from TOC |
| `q` / `Ctrl+C` | Quit |

## Example Output

### Table of Contents View
```
📄 README.md - Table of Contents

  L1-4      ▶ Introduction
  L5-8      ├── Installation
  L9-14     │   └── Requirements
  L15-18    │   └── Setup
  L19-20    ├── Usage
  L21-24    │   └── Basic Commands
  L25-28    │   └── Advanced Options
  L29-31    └── License

[↑/↓] Navigate  [→] Expand  [←] Collapse  [Enter] View  [q] Quit
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

- Go 1.23+
- Terminal with UTF-8 and color support

## License

MIT
