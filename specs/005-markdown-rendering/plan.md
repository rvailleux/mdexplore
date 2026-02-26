# Implementation Plan: Rich Markdown Content Rendering

**Branch**: `005-markdown-rendering` | **Date**: 2026-02-26 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/005-markdown-rendering/spec.md`

## Summary

Implement rich markdown content rendering in the terminal with styled headings, lists, code blocks, text formatting, links, blockquotes, and tables. The feature will leverage an existing terminal markdown rendering library integrated with the current Bubble Tea TUI framework.

## Technical Context

**Language/Version**: Go 1.23+
**Primary Dependencies**: NEEDS CLARIFICATION - Research markdown-to-terminal rendering library
**Storage**: N/A (in-memory rendering)
**Testing**: Go's standard testing package
**Target Platform**: Linux, macOS, Windows terminal environments
**Project Type**: CLI tool with interactive TUI
**Performance Goals**: Render sections under 100ms for content up to 1000 lines
**Constraints**: Terminal width handling (80-200 columns), UTF-8 support, 256-color terminals
**Scale/Scope**: Single document rendering, typical sections < 1000 lines

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

| Principle | Status | Notes |
|-----------|--------|-------|
| **I. Test-Driven Development** | ✅ PASS | All rendering behaviors can be unit tested |
| **II. Terminal User Interface** | ✅ PASS | Feature enhances TUI with rich markdown display |
| **III. Markdown Reader Focus** | ✅ PASS | Core feature directly supports markdown reading |
| **IV. Agent-First Design** | ✅ PASS | Rich formatting improves spec/documentation readability |
| **V. Simplicity** | ⚠️ NEEDS RESEARCH | Library choice impacts complexity - evaluate options |

**Verdict**: Proceed to Phase 0 research to resolve markdown rendering library selection.

## Project Structure

### Documentation (this feature)

```text
specs/005-markdown-rendering/
├── plan.md              # This file
├── research.md          # Phase 0 output - library comparison
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
└── tasks.md             # Phase 2 output (via /speckit.tasks)
```

### Source Code (repository root)

```text
internal/
├── ui/
│   ├── model.go         # No changes (rendering is view concern)
│   ├── update.go        # No changes
│   ├── view.go          # Integrate markdown renderer in content view
│   └── renderer.go      # NEW: Markdown to terminal rendering adapter
├── renderer/
│   └── markdown.go      # NEW: Markdown rendering using selected library
└── models/
    └── section.go       # No changes needed
```

**Structure Decision**: Add a new `renderer` package to encapsulate markdown-to-terminal rendering logic. Keep UI layer thin by delegating rendering to this package.

## Phase 0: Research

### Research Questions

**RQ-001**: Which Go library provides the best terminal markdown rendering?
**RQ-002**: How does the library integrate with Bubble Tea's viewport/component model?
**RQ-003**: What customization options exist for colors, bullets, and styling?

### Library Candidates

1. **Glamour** (github.com/charmbracelet/glamour)
   - By Charm (same org as Bubble Tea/Lipgloss)
   - Uses Goldmark parser + Lipgloss styling
   - Themes support (Dark, Light, ASCII)
   - ANSI output

2. **Termdown** (github.com/williamchanning/termdown)
   - Alternative markdown viewer
   - Less widely used

3. **Custom solution with Goldmark + Lipgloss**
   - Build custom renderer using existing parser
   - More control but more code

### Research Tasks

- Compare Glamour vs alternatives for features, performance, maintenance
- Test Glamour integration with Bubble Tea viewport
- Verify syntax highlighting support for common languages
- Test table rendering quality
- Evaluate theme customization options

## Phase 1: Design

### Data Model

**Minimal changes** - rendering is a view-layer concern:

```go
// Renderer interface for markdown-to-terminal rendering
type Renderer interface {
    Render(markdown string, width int) (string, error)
}

// GlamourRenderer implements Renderer using Glamour library
type GlamourRenderer struct {
    style glamour.Style
}
```

### Interface Contracts

No external interfaces - internal rendering abstraction only.

### Integration Points

1. **View layer** (`internal/ui/view.go`):
   - Replace `contentStyle.Render(content)` with markdown renderer
   - Pass terminal width for proper wrapping

2. **Renderer package** (`internal/renderer/markdown.go`):
   - Wrap Glamour or alternative library
   - Provide clean interface for UI layer

### Implementation Strategy

**Order of implementation (TDD)**:

1. **Renderer package** (core capability)
   - Test: Renderer converts markdown to styled text
   - Implement: Glamour wrapper with error handling

2. **Heading styles** (P1)
   - Test: H1-H6 render with distinct colors
   - Implement: Glamour theme configuration

3. **Lists** (P1)
   - Test: Ordered/unordered lists render correctly
   - Implement: Verify Glamour list rendering

4. **Code blocks** (P1)
   - Test: Code blocks have background/syntax highlighting
   - Implement: Glamour syntax highlighting config

5. **Text formatting** (P1)
   - Test: Bold, italic, strikethrough render correctly
   - Implement: Style configuration

6. **Links, blockquotes, tables** (P2-P3)
   - Test: Each element renders appropriately
   - Implement: Style adjustments as needed

### Agent Context Update

Run `.specify/scripts/bash/update-agent-context.sh claude` after Phase 1 to document new dependency (Glamour or alternative).

## Constitution Re-Check Post-Design

| Principle | Status | Verification |
|-----------|--------|--------------|
| **I. TDD** | ✅ PASS | Each rendering feature testable |
| **II. TUI** | ✅ PASS | Integrates with Bubble Tea |
| **III. Markdown Focus** | ✅ PASS | Core markdown rendering feature |
| **IV. Agent-First** | ✅ PASS | Improves documentation readability |
| **V. Simplicity** | ✅ PASS | Use existing library (Glamour), minimal custom code |

**Proceed to `/speckit.tasks` for task generation.**
