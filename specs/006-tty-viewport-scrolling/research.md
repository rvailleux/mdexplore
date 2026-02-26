# Research: TTY Viewport Scrolling

**Feature**: TTY Viewport Scrolling
**Date**: 2026-02-26
**Purpose**: Research and document technical decisions for implementing viewport scrolling in the TUI

---

## Decision: Viewport Implementation Strategy

### Decision
Use Bubble Tea's `bubbles/viewport` component for scrollable content viewing, integrated with custom scroll offset tracking in the model.

### Rationale
- **Native Integration**: `bubbles/viewport` is designed specifically for Bubble Tea applications
- **Built-in Features**: Handles terminal resize, key events, and viewport windowing
- **Keyboard Support**: Already implements Page Up, Page Down, Up, Down, Home, End key handling
- **Performance**: Efficiently renders only visible content lines
- **Testability**: Component-based design allows isolated testing

### Alternatives Considered

| Alternative | Pros | Cons | Decision |
|-------------|------|------|----------|
| Custom viewport logic | Full control | More code to maintain, potential bugs | Rejected - unnecessary complexity |
| `lipgloss` window only | Lightweight | No built-in scroll state management | Rejected - need scroll tracking |
| Raw terminal manipulation | Maximum performance | Platform-specific code, brittle | Rejected - too low-level |

---

## Decision: Scroll Position Indicator Style

### Decision
Display a simple percentage-based indicator in the help/footer line, e.g., `[ 45% ]` or position text like `Line 45/100`.

### Rationale
- **Minimal UI Clutter**: Single line at bottom, doesn't consume content space
- **Standard Pattern**: Similar to `less`, `vim`, and other pager tools
- **Easy to Implement**: Calculated from scroll offset and total content lines

### Alternatives Considered

| Alternative | Pros | Cons | Decision |
|-------------|------|------|----------|
| Scrollbar on right | Visual familiarity | Takes horizontal space, complex in terminal | Rejected - too complex for TUI |
| Progress bar | Visual progress | Requires additional rendering logic | Rejected - percentage is sufficient |
| No indicator | Simplest | Users lose context in long docs | Rejected - poor UX |

---

## Decision: Content Height Calculation

### Decision
Calculate content height from rendered string line count after markdown-to-text conversion.

### Rationale
- **Accurate Sizing**: Rendered content (with Glamour) determines actual terminal lines
- **Wrapped Line Awareness**: Markdown rendering may wrap long lines, affecting total height
- **Dynamic Adjustment**: Recalculate on terminal resize events

### Implementation Notes
```go
// Split rendered content by newlines to get actual display lines
lines := strings.Split(renderedContent, "\n")
totalHeight := len(lines)
```

---

## Decision: Scroll Key Mapping

### Decision
Map standard navigation keys:
- `↑` / `k`: Scroll up one line
- `↓` / `j`: Scroll down one line
- `Page Up`: Scroll up one screen (viewport height)
- `Page Down`: Scroll down one screen
- `Home` / `g`: Jump to top
- `End` / `G`: Jump to bottom

### Rationale
- **Vim Compatibility**: `j`/`k` navigation familiar to developers
- **Standard Pager Keys**: Page Up/Down match `less`, `more`
- **Home/End**: Universal document navigation

---

## Decision: Viewport State in Model

### Decision
Add the following fields to the Model struct:
```go
type Model struct {
    // ... existing fields ...
    ContentScrollOffset int    // Current scroll position (lines from top)
    ContentTotalLines   int    // Total rendered lines
    ViewportHeight      int    // Available height for content (terminal - headers/footers)
}
```

### Rationale
- **Separation of Concerns**: Scroll state separate from TOC navigation
- **Persistent State**: Maintains position when switching views
- **Testability**: Pure functions for viewport calculations

---

## Decision: Terminal Resize Handling

### Decision
On `tea.WindowSizeMsg`:
1. Update `Width` and `Height`
2. Recalculate `ViewportHeight` (accounting for title bar and help text)
3. Adjust `ContentScrollOffset` if viewport now shows beyond content end

### Rationale
- **Responsive**: Content adapts to terminal changes
- **Boundary Safe**: Prevents scroll position from pointing to non-existent content

---

## Summary

All technical decisions align with the Bubble Tea ecosystem and follow terminal UI conventions. The implementation will:

1. Use `bubbles/viewport` for core scroll functionality
2. Track scroll state in the Model struct
3. Display percentage indicator in help line
4. Support vim-style and standard pager keybindings
5. Handle terminal resize gracefully
