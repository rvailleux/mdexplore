# Quickstart: TTY Viewport Scrolling

**Feature**: TTY Viewport Scrolling
**Date**: 2026-02-26

---

## Development Setup

### Prerequisites
- Go 1.23+
- Terminal with UTF-8 and color support

### Running the Application

```bash
# Build the application
go build -o mdexplore ./cmd/mdexplore

# Run with a markdown file
./mdexplore README.md

# Navigate to a section and view content
# When content exceeds terminal height, use scroll keys:
#   ↑/↓     - Scroll line by line
#   PgUp    - Scroll up one screen
#   PgDown  - Scroll down one screen
#   Home/g  - Jump to top
#   End/G   - Jump to bottom
#   Esc     - Return to TOC
#   q       - Quit
```

### Testing Scroll Behavior

```bash
# Create a test file with long content
cat > /tmp/long-doc.md << 'EOF'
# Section 1

Line 1
Line 2
... (repeat to create 100+ lines)

# Section 2
More content...
EOF

# Run and test scrolling
./mdexplore /tmp/long-doc.md
```

---

## Key Implementation Points

### 1. Model State

Add to `internal/ui/model.go`:
```go
type Model struct {
    // ... existing fields ...
    ContentScrollOffset int  // Lines scrolled from top
    ContentTotalLines   int  // Total rendered lines
    ViewportHeight      int  // Available content height
}
```

### 2. Viewport Windowing

In `internal/ui/view.go`, modify `renderContent()`:
```go
// After rendering markdown, extract visible window
lines := strings.Split(renderedContent, "\n")
m.ContentTotalLines = len(lines)

// Calculate visible range
start := m.ContentScrollOffset
end := min(start+m.ViewportHeight, m.ContentTotalLines)
visibleLines := lines[start:end]

// Join and display
contentToShow := strings.Join(visibleLines, "\n")
```

### 3. Key Handlers

In `internal/ui/update.go`, add to `handleKeyPress()`:
```go
case "up", "k":
    if m.ContentScrollOffset > 0 {
        m.ContentScrollOffset--
    }

case "down", "j":
    maxOffset := max(0, m.ContentTotalLines-m.ViewportHeight)
    if m.ContentScrollOffset < maxOffset {
        m.ContentScrollOffset++
    }

case "pgup":
    m.ContentScrollOffset = max(0, m.ContentScrollOffset-m.ViewportHeight)

case "pgdown":
    maxOffset := max(0, m.ContentTotalLines-m.ViewportHeight)
    m.ContentScrollOffset = min(maxOffset, m.ContentScrollOffset+m.ViewportHeight)
```

### 4. Scroll Position Display

In `renderContent()`, add to help line:
```go
scrollPercent := 0
if m.ContentTotalLines > m.ViewportHeight {
    scrollPercent = (m.ContentScrollOffset * 100) /
                    (m.ContentTotalLines - m.ViewportHeight)
}
helpText := fmt.Sprintf("[%3d%%] [↑/↓] Scroll [PgUp/PgDown] Page [Esc] Back [q] Quit", scrollPercent)
```

---

## Running Tests

```bash
# Run all tests
go test ./...

# Run viewport-specific tests
go test ./internal/ui/... -v

# Run with coverage
go test -cover ./internal/ui/...

# Run benchmarks
go test -bench=. ./tests/benchmark/...
```
