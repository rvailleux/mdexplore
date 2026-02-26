# UI Key Bindings Contract

**Feature**: TTY Viewport Scrolling
**Date**: 2026-02-26

---

## Content View Key Bindings

When in content view mode (`ViewMode == ViewContent`), the following keys are supported:

| Key | Action | Behavior |
|-----|--------|----------|
| `↑` | Scroll Up | Move viewport up by 1 line (if not at top) |
| `↓` | Scroll Down | Move viewport down by 1 line (if not at bottom) |
| `k` | Scroll Up | Vim-style: move viewport up by 1 line |
| `j` | Scroll Down | Vim-style: move viewport down by 1 line |
| `PgUp` | Page Up | Move viewport up by one screen height |
| `PgDown` | Page Down | Move viewport down by one screen height |
| `Home` | Jump to Top | Set scroll offset to 0 (top of content) |
| `End` | Jump to Bottom | Set scroll offset to end of content |
| `g` | Jump to Top | Vim-style: go to beginning of content |
| `G` | Jump to Bottom | Vim-style: go to end of content |
| `Esc` | Return to TOC | Exit content view, return to table of contents |
| `Left` | Return to TOC | Alternative: exit content view |
| `h` | Return to TOC | Vim-style: exit content view |
| `q` | Quit | Exit application |
| `Ctrl+C` | Quit | Exit application |

---

## Scroll Behavior Contract

### Scroll Boundaries
- Scrolling up when at top has no effect (scroll offset remains 0)
- Scrolling down when at bottom has no effect (scroll offset remains at maximum)

### Scroll Position Indicator
- Displayed in help/footer line as percentage: `[ 45% ]`
- 100% indicates bottom of content
- 0% indicates top of content
- Hidden when all content fits within viewport

### Terminal Resize
- On resize, viewport height is recalculated
- Scroll offset is adjusted if viewport would show empty lines
- Content remains anchored to valid scroll position

---

## State Transitions

```
┌─────────────────────────────────────────────────────────────┐
│                      Content View Mode                       │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  ┌─────────────┐        ┌─────────────┐                    │
│  │   At Top    │◄──────►│   Scrolled  │                    │
│  │  (offset=0) │  ↑/↓   │  (offset>0) │                    │
│  └─────────────┘        └──────┬──────┘                    │
│         ▲                      │                            │
│         │                      ▼                            │
│         │               ┌─────────────┐                    │
│         │               │  At Bottom  │                    │
│         └───────────────┤(offset=max) │                    │
│               Home/g    └─────────────┘    End/G           │
│                                                             │
└─────────────────────────────────────────────────────────────┘
                              │
                              │ Esc/Left/h
                              ▼
                    ┌─────────────────┐
                    │   TOC View Mode │
                    └─────────────────┘
```
