# CLI Contract: mdexplore

**Feature**: Markdown TOC Explorer
**Date**: 2025-02-25

## Command Interface

### Usage

```bash
mdexplore <file> [--toc]
```

### Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `file` | Yes | Path to markdown file to analyze |

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--toc` | false | Display table of contents (default behavior when no other flags) |
| `--help` | - | Display help information |
| `--version` | - | Display version information |

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | File not found |
| 3 | Permission denied |
| 4 | Invalid file (directory instead of file) |
| 5 | Parse error |

## Output Formats

### Interactive TUI Mode (default with --toc)

When `--toc` is specified, the application launches an interactive TUI:

```
┌─────────────────────────────────────────────────┐
│  📄 README.md - Table of Contents               │
├─────────────────────────────────────────────────┤
│                                                 │
│  ├── Introduction                               │
│  ├── Installation                               │
│  │   ├── Requirements                          │
│  │   └── Setup                                 │
│  ├── Usage                                      │
│  │   ├── Basic Commands                        │
│  │   └── Advanced Options                      │
│  └── License                                    │
│                                                 │
│  [↑/↓] Navigate  [q] Quit  [enter] Copy path   │
└─────────────────────────────────────────────────┘
```

**Key Bindings**:
- `↑` / `↓` or `k` / `j` - Navigate up/down
- `q` / `Ctrl+C` / `Esc` - Quit
- `Enter` - Copy heading text to clipboard (optional feature)

### Error Output

Errors are displayed in the TUI with styling:

```
┌─────────────────────────────────────────────────┐
│  ❌ Error                                       │
├─────────────────────────────────────────────────┤
│                                                 │
│  File not found: /path/to/missing.md            │
│                                                 │
│  Press any key to exit                          │
└─────────────────────────────────────────────────┘
```

## Examples

### Display TOC

```bash
$ mdexplore README.md --toc
# Launches TUI with hierarchical TOC
```

### Error: File Not Found

```bash
$ mdexplore missing.md --toc
# TUI shows error: "File not found: missing.md"
# Exit code: 2
```

### Error: Directory

```bash
$ mdexplore ./docs --toc
# TUI shows error: "Expected a file, got directory: ./docs"
# Exit code: 4
```

## Constraints

- Maximum file size: 10MB (configurable, fails gracefully with clear message)
- Supported encoding: UTF-8
- Minimum terminal width: 40 columns
- Minimum terminal height: 10 rows
