# Quickstart: TUI Navigation Enhancement

## Updated Key Bindings

| Key | Action |
|-----|--------|
| `↑` / `↓` or `k` / `j` | Navigate up/down in TOC |
| `→` | Expand section **OR** open content (if leaf section) |
| `←` | Collapse section **OR** return from content view |
| `Enter` | View full section content (including all subsections) |
| `Esc` | Return from content view / Quit from TOC |
| `q` / `Ctrl+C` | Quit |

## Navigation Patterns

### Drill Down (Right Arrow)
- On **parent section** (has children): expands to show children
- On **leaf section** (no children): opens content view

### Return Up (Left Arrow)
- In **content view**: returns to TOC at previous position
- In **TOC**: collapses expanded section

### View Full Content (Enter)
- Opens content view showing section + all descendants
- Use for reading complete logical units

## Visual Layout

```
📄 README.md - Table of Contents

  L1-13     ▶ 1. Introduction
  L14-26    ├── 1.1. Table of Contents
  L27-39    ├── 1.2. Features
  L40-41    └── 1.3. Installation
  L42-47        ├── 1.3.1. From Source
  L48-55        ├── 1.3.2. Build Locally
  L56-59        └── 1.3.3. Pre-built Binaries

[↑/↓] Navigate  [→] Expand/Open  [←] Collapse/Back  [Enter] Full Content  [q] Quit
```

Note: Section numbers (1., 1.1., etc.) now appear next to titles for better readability.
