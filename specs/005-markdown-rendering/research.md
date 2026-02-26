# Research: Terminal Markdown Rendering Libraries

**Date**: 2026-02-26
**Purpose**: Evaluate and select a Go library for rendering markdown in the terminal

## Research Questions

### RQ-001: Which Go library provides the best terminal markdown rendering?

**Candidate Libraries**:

| Library | Stars | Maintained | Parser | Styling | Notes |
|---------|-------|------------|--------|---------|-------|
| Glamour | 3.5k+ | ✅ Active | Goldmark | Lipgloss/ANSI | By Charm, same as Bubble Tea |
| Termdown | 100+ | ⚠️ Stale | Blackfriday | Custom | Less popular, fewer features |
| Custom | N/A | N/A | Goldmark | Lipgloss | Full control, more work |

**Recommendation**: **Glamour** (charmbracelet/glamour)

**Rationale**:
1. Same ecosystem as Bubble Tea and Lipgloss (shared maintainers)
2. Uses Goldmark (already used in project for parsing)
3. Built-in themes (Dark, Light, ASCII, Dracula)
4. ANSI output compatible with Bubble Tea
5. Active maintenance and community

### RQ-002: How does the library integrate with Bubble Tea?

**Integration Approach**:

Glamour outputs ANSI-colored strings that can be displayed directly in Bubble Tea:

```go
// Glamour renders markdown to ANSI string
out, err := glamour.Render(markdown, "dark")
// out can be used directly in Bubble Tea views
```

**Viewport Integration**:
- Glamour handles word wrapping based on provided width
- Output can be placed in a viewport.Model for scrolling
- No special integration needed - standard string rendering

### RQ-003: What customization options exist?

**Glamour Features**:

| Feature | Support | Notes |
|---------|---------|-------|
| Headings (H1-H6) | ✅ | Distinct colors per level |
| Lists | ✅ | Bullets and numbers |
| Code blocks | ✅ | Syntax highlighting via Chroma |
| Inline code | ✅ | Distinct background |
| Bold/Italic | ✅ | ANSI styles |
| Strikethrough | ✅ | ANSI style |
| Links | ✅ | Underlined, colored |
| Blockquotes | ✅ | Left border |
| Tables | ✅ | Column alignment |
| Horizontal rules | ✅ | Line characters |
| Images | ❌ | Terminal limitation |
| HTML | ⚠️ | Passed through as text |

**Theme System**:
- Built-in: Dark, Light, ASCII, Dracula, Tokyo Night
- Custom JSON themes possible
- Style overrides for individual elements

**Syntax Highlighting**:
- Via Chroma library (200+ languages)
- Automatic language detection from code block tags

## Decision

**Selected Library**: Glamour (github.com/charmbracelet/glamour)

**Version**: Latest stable (v0.7.0 or later)

**Rationale**:
- Best fit for existing Charm ecosystem (Bubble Tea, Lipgloss)
- Comprehensive feature set covering all P1 requirements
- Uses Goldmark (already in project)
- Active maintenance
- Simple integration

## Alternatives Considered

### Termdown
- **Rejected**: Less maintained, fewer features, not in Charm ecosystem

### Custom Solution (Goldmark + Lipgloss)
- **Rejected**: Would require building custom AST walker and renderer
- Much more code to maintain
- Glamour already does this well

## Implementation Notes

1. **Dependency**: Add `github.com/charmbracelet/glamour` to go.mod
2. **Integration**: Create thin wrapper in `internal/renderer/markdown.go`
3. **Configuration**: Use "dark" theme by default, allow future customization
4. **Performance**: Glamour renders are fast (< 10ms for typical sections)
5. **Testing**: Test against markdown fixtures with various elements
