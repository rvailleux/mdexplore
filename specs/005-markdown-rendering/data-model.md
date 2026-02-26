# Data Model: Rich Markdown Content Rendering

## Entity Changes

### Renderer Interface (internal/renderer/markdown.go)

**New Interface**:
```go
// Renderer converts markdown to styled terminal output
type Renderer interface {
    // Render converts markdown text to terminal-ready styled text
    Render(markdown string, width int) (string, error)

    // RenderWithStyle renders with a specific style theme
    RenderWithStyle(markdown string, width int, style string) (string, error)
}
```

### GlamourRenderer (internal/renderer/glamour.go)

**New Struct**:
```go
// GlamourRenderer implements Renderer using the Glamour library
type GlamourRenderer struct {
    defaultStyle string  // "dark", "light", "ascii", etc.
}

// NewGlamourRenderer creates a new renderer with default style
func NewGlamourRenderer(style string) *GlamourRenderer

// Render implements the Renderer interface
func (r *GlamourRenderer) Render(markdown string, width int) (string, error)
```

### UI Model Integration (internal/ui/model.go)

**No changes required** - rendering is a view-layer concern.

The Model's `CurrentSection` already contains the `RawContent` which is passed to the renderer in the view layer.

### View Changes (internal/ui/view.go)

**Integration point**:
```go
// In renderContent(), replace:
//   content := m.CurrentSection.GetFullContent()
//   b.WriteString(contentStyle.Render(content))

// With:
//   content := m.CurrentSection.GetFullContent()
//   rendered, err := m.renderer.Render(content, m.Width)
//   b.WriteString(rendered)
```

## State Flow

```
[Section.RawContent]
        ↓
[GlamourRenderer.Render()]
        ↓
[ANSI Styled String]
        ↓
[Bubble Tea View Display]
```

## Validation Rules

- Renderer must handle empty content gracefully
- Width parameter must be positive (minimum 40 for readability)
- Invalid markdown should render as plain text (never error on display)
- ANSI codes from Glamour must be compatible with terminal

## Dependencies

- `github.com/charmbracelet/glamour` - Markdown to terminal rendering
- Already have: `github.com/yuin/goldmark` (used internally by Glamour)
