# Data Model: TTY Viewport Scrolling

**Feature**: TTY Viewport Scrolling
**Date**: 2026-02-26

---

## Viewport State

The viewport scrolling feature requires tracking scroll position and content dimensions within the TUI model.

### Model Fields (Additions to ui.Model)

| Field | Type | Description |
|-------|------|-------------|
| `ContentScrollOffset` | `int` | Number of lines scrolled from top (0 = at top) |
| `ContentTotalLines` | `int` | Total lines in rendered content |
| `ViewportHeight` | `int` | Available lines for content display (terminal height - headers/footers) |

### State Transitions

```
Initial State:
  ContentScrollOffset = 0
  ContentTotalLines = calculated from rendered markdown
  ViewportHeight = terminalHeight - headerHeight - footerHeight

Scroll Down:
  ContentScrollOffset = min(ContentScrollOffset + 1, max(0, ContentTotalLines - ViewportHeight))

Scroll Up:
  ContentScrollOffset = max(ContentScrollOffset - 1, 0)

Page Down:
  ContentScrollOffset = min(ContentScrollOffset + ViewportHeight, max(0, ContentTotalLines - ViewportHeight))

Page Up:
  ContentScrollOffset = max(ContentScrollOffset - ViewportHeight, 0)

Jump to Top:
  ContentScrollOffset = 0

Jump to Bottom:
  ContentScrollOffset = max(0, ContentTotalLines - ViewportHeight)

Terminal Resize:
  ViewportHeight = newTerminalHeight - headerHeight - footerHeight
  ContentScrollOffset = min(ContentScrollOffset, max(0, ContentTotalLines - ViewportHeight))
```

### Content Window Calculation

```
Visible Content Range:
  StartLine = ContentScrollOffset
  EndLine = min(ContentScrollOffset + ViewportHeight, ContentTotalLines)

Rendered Content for Display:
  visibleLines = allContentLines[StartLine:EndLine]
```

### Scroll Position Indicator

```
Scroll Percentage:
  if ContentTotalLines <= ViewportHeight:
    percentage = 100% (all content visible)
  else:
    percentage = (ContentScrollOffset / (ContentTotalLines - ViewportHeight)) * 100

Display Format:
  "[ 45% ]" or "Line 45/100"
```

---

## Entity Relationships

```
┌─────────────────────────────────────┐
│              ui.Model               │
├─────────────────────────────────────┤
│  ViewMode                           │
│  CurrentSection  ───────┐           │
│  ContentScrollOffset    │           │
│  ContentTotalLines      │           │
│  ViewportHeight         │           │
│  markdownRenderer       │           │
└──────────────────────│──┘           │
                       │              │
                       ▼              │
            ┌─────────────────────┐   │
            │   models.Section    │   │
            ├─────────────────────┤   │
            │  Content (raw MD)   │   │
            └─────────────────────┘   │
                       │              │
                       ▼              │
            ┌─────────────────────┐   │
            │  renderer.Render()  │◄──┘
            ├─────────────────────┤
            │ Rendered Content    │
            │ (formatted text)    │
            └─────────────────────┘
                       │
                       ▼
            ┌─────────────────────┐
            │  Viewport Window    │
            │  (visible slice)    │
            └─────────────────────┘
```

---

## Validation Rules

| Rule | Validation |
|------|------------|
| Scroll offset non-negative | `ContentScrollOffset >= 0` |
| Scroll offset bounded | `ContentScrollOffset <= max(0, ContentTotalLines - ViewportHeight)` |
| Viewport height positive | `ViewportHeight > 0` for scrolling to be meaningful |
| Total lines non-negative | `ContentTotalLines >= 0` |
