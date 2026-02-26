# Quickstart: Enhanced Section Navigation

**Feature**: Enhanced Section Navigation
**Date**: 2026-02-25

---

## Installation

```bash
# Build from source
go build -o mdexplore ./cmd/mdexplore

# Or install to $GOPATH/bin
go install ./cmd/mdexplore
```

---

## Basic Usage

### View Table of Contents

```bash
# Display TOC with line numbers
./mdexplore README.md

# Output shows:
# L1-45   ● Introduction
# L47-89  ├── Getting Started
# L91-120 └── Architecture
```

### Limit Depth

```bash
# Show only H1 and H2 headings
./mdexplore --level 2 README.md

# Or use shorthand
./mdexplore -L 2 README.md

# Show only top-level sections
./mdexplore -L 1 docs/specification.md
```

---

## Interactive Navigation

### Launch TUI

```bash
./mdexplore README.md
```

Initial view shows only top-level (H1) sections:

```
📄 README.md - Table of Contents

  L1-45   ● Introduction
  L47-120 ● Getting Started
  L122-200 ● Architecture

[↑/↓] Navigate  [←] Expand  [Enter] View  [q] Quit
```

### Expand Sections

1. Use `↑`/`↓` to select a section
2. Press `←` to expand and see subsections

```
📄 README.md - Table of Contents

  L1-45   ● Introduction
  L47-120 ● Getting Started [selected, expanded]
  L49-80  │   ├── Prerequisites
  L82-110 │   ├── Installation
  L112-118│   └── Configuration
  L122-200 ● Architecture
```

### View Section Content

1. Navigate to any section
2. Press `Enter` to view its content

```
┌─ L49-80: Prerequisites ───────────────────────────┐
│                                                    │
│  ## Prerequisites                                  │
│                                                    │
│  - Go 1.23 or later                               │
│  - Git                                            │
│  - Make (optional)                                │
│                                                    │
└─ Press Esc to return to navigation ───────────────┘
```

3. Press `Esc` to return to TOC navigation

### Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `↑` / `↓` or `k` / `j` | Navigate up/down |
| `←` | Expand section |
| `→` | Collapse section |
| `Enter` | View section content |
| `Esc` | Return / Quit |
| `q` | Quit |

---

## Scripting Usage

### Extract TOC to File

```bash
# Save TOC to file
./mdexplore README.md > toc.txt

# Process with other tools
./mdexplore -L 2 README.md | grep "Getting Started"
```

### Check File Structure

```bash
# Count top-level sections
./mdexplore -L 1 README.md | wc -l

# Find specific section line numbers
./mdexplore README.md | grep "Architecture"
```

---

## Troubleshooting

### File Not Found

```bash
$ ./mdexplore missing.md

❌ Error

File not found: missing.md

Press any key to exit
```

**Solution**: Check the file path and ensure the file exists.

### Permission Denied

```bash
$ ./mdexplore /root/secret.md

🔒 Error

Permission denied: /root/secret.md
```

**Solution**: Check file permissions or run with appropriate access.

### No Headings Found

```bash
$ ./mdexplore empty.md

📄 empty.md - Table of Contents

  No headings found in this file.

[↑/↓] Navigate  [q] Quit
```

**Solution**: Ensure the markdown file contains headings (`# Heading` syntax).

---

## Tips

1. **Start broad, then narrow**: Use `-L 1` to see document structure, then expand sections of interest
2. **Quick reference**: Line numbers help you jump to sections in your editor
3. **Keyboard navigation**: Learn the vim-style shortcuts (`j`/`k`) for faster navigation
4. **Content preview**: Press Enter to quickly preview a section without leaving the tool
