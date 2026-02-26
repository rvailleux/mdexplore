# CLI Contract: mdexplore

**Feature**: Enhanced Section Navigation
**Date**: 2026-02-25
**Status**: Draft

---

## Command Interface

### Usage

```bash
mdexplore <file> [flags]
mdexplore --version
mdexplore --help
```

### Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `file` | Yes* | Path to markdown file to explore |

*Required unless using `--version` or `--help`.

### Flags

| Flag | Shorthand | Type | Default | Description |
|------|-----------|------|---------|-------------|
| `--level` | `-L` | `int` | `0` | Maximum heading level to display (0 = no limit) |
| `--toc` | - | `bool` | `true` | Display table of contents |
| `--help` | `-h` | `bool` | `false` | Display help information |
| `--version` | `-v` | `bool` | `false` | Display version information |

### Exit Codes

| Code | Meaning | Trigger |
|------|---------|---------|
| `0` | Success | Normal execution completed |
| `1` | General error | Unknown or unexpected error |
| `2` | File not found | Specified file does not exist |
| `3` | Permission denied | Cannot read specified file |
| `4` | Invalid file | Path is directory or invalid format |
| `5` | Parse error | Failed to parse markdown |

### Output Modes

#### Non-TTY Mode (piped/redirected)

When stdout is not a terminal, outputs plain text TOC:

```
L1-45 Introduction
L47-89 Getting Started
L91-120 Architecture
L93-110   Overview
```

- Respects `--level` flag
- No interactive features available
- Suitable for scripting

#### TTY Mode (interactive terminal)

Launches Bubble Tea TUI with:
- Interactive navigation
- Hierarchical section expansion
- Content viewing on Enter

---

## Keyboard Interface (TUI Mode)

### Navigation Keys

| Key | Action |
|-----|--------|
| `↑` or `k` | Move selection up |
| `↓` or `j` | Move selection down |
| `←` | Expand selected section (reveals children) |
| `→` | Collapse selected section (hides children) |
| `Enter` | View content of selected section |
| `Esc` | Return to navigation from content view / Quit if in TOC |
| `q` | Quit application |
| `Ctrl+C` | Force quit |

### Content View Keys

| Key | Action |
|-----|--------|
| `Esc` | Return to TOC navigation |
| `q` | Quit application |
| `Ctrl+C` | Force quit |

---

## Examples

### Basic Usage

```bash
# Display TOC for README.md
mdexplore README.md

# Show only H1 and H2 headings
mdexplore --level 2 README.md

# Short flag form
mdexplore -L 2 docs/guide.md

# Show only top-level sections
mdexplore -L 1 specification.md
```

### Scripting Usage

```bash
# Get filtered TOC for processing
mdexplore -L 2 README.md > toc.txt

# Check if file has headings
mdexplore README.md > /dev/null && echo "Has headings"

# Error handling
if ! mdexplore README.md 2>/dev/null; then
    echo "Failed to parse"
fi
```

---

## Error Output

Errors are written to stderr with descriptive messages:

```
❌ Error

File not found: /path/to/missing.md

Press any key to exit
```

---

## Version Output

```
mdexplore version dev
```

---

## Help Output

```
mdexplore displays a hierarchical table of contents from markdown files.

Usage:
  mdexplore <file> [flags]

Flags:
  -L, --level int    Maximum heading level to display (0 = no limit)
  -h, --help         Display help information
  -v, --version      Display version information

Examples:
  mdexplore README.md --toc
  mdexplore -L 2 docs/guide.md

Navigation (TUI mode):
  ↑/↓ or k/j    Navigate
  ←             Expand section
  →             Collapse section
  Enter         View section content
  Esc           Return / Quit
  q             Quit
```
