# Implementation Plan: Enhanced Section Navigation

**Branch**: `002-section-navigation` | **Date**: 2026-02-25 | **Spec**: [spec.md](spec.md)
**Input**: Feature specification from `/specs/002-section-navigation/spec.md`

## Summary

This feature enhances mdexplore with three key capabilities: (1) displaying line number ranges for each section in the TOC, (2) adding a `--level` flag to limit TOC depth, and (3) implementing hierarchical navigation in the TUI where users can expand sections with left arrow, view content with Enter, and return with Escape.

## Technical Context

**Language/Version**: Go 1.23+ (existing codebase)
**Primary Dependencies**: Bubble Tea (TUI), Goldmark (markdown parsing), Lipgloss (styling), Cobra (CLI)
**Storage**: File-based (reads markdown files from disk)
**Testing**: Go's standard testing package with testify for assertions
**Target Platform**: Linux, macOS, Windows terminal environments
**Project Type**: CLI tool with TUI
**Performance Goals**: Parse files up to 10MB in under 500ms, TUI response <16ms
**Constraints**: Must maintain backward compatibility with existing TOC display, adhere to TDD discipline
**Scale/Scope**: Single-user CLI tool, processes one file at a time

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

### Principle I: Test-Driven Development ✅
- All new code requires tests written first (Red-Green-Refactor)
- Parser changes need unit tests in `internal/parser/parser_test.go`
- UI changes need tests verifying key handling and state transitions
- Integration tests needed for end-to-end navigation flow

### Principle II: Terminal User Interface ✅
- Bubble Tea already in use, continue using for hierarchical navigation
- Keyboard navigation (↑/↓/←/→/Enter/Escape/q) must be implemented
- Visual styling for selected items, expanded sections, and content view
- Handle terminal resizing gracefully

### Principle III: Markdown Reader Focus ✅
- Scope remains reading/navigation, not editing
- Content display must preserve markdown formatting
- Line numbers enhance reading/navigation experience
- Hierarchical navigation supports spec reading workflows

### Principle IV: Agent-First Design ✅
- Line numbers enable quick reference for agents
- Hierarchical navigation reduces cognitive load
- Keyboard-driven interface for efficient navigation
- Fast context switching between TOC and content view

### Principle V: Simplicity ✅
- Extend existing models rather than replacing
- Reuse existing parser infrastructure
- Build on existing TUI model structure
- No external dependencies beyond existing stack

**Constitution Check Result**: ✅ All principles satisfied. Proceed with implementation.

## Project Structure

### Documentation (this feature)

```text
specs/002-section-navigation/
├── plan.md              # This file
├── research.md          # Phase 0 output (N/A - no research needed)
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 output (CLI contract)
└── tasks.md             # Phase 2 output (/speckit.tasks command)
```

### Source Code (repository root)

```text
cmd/mdexplore/
└── main.go              # CLI entry point (add --level/-L flag handling)

internal/
├── models/
│   ├── heading.go       # Extend Heading with StartLine/EndLine, add Section struct
│   └── section.go       # NEW: Section tree structure with hierarchical relationships
├── parser/
│   └── parser.go        # Modify Parse() to calculate end lines and build section tree
├── ui/
│   ├── model.go         # Extend Model with viewMode, expandedSections, levelFilter
│   ├── update.go        # Add handlers for ←/→/Enter/Escape keys
│   ├── view.go          # Add hierarchical rendering and content view
│   └── content_view.go  # NEW: Section content rendering with markdown formatting
└── renderer/
    └── renderer.go      # NEW: Markdown content renderer for section display

tests/
├── unit/                # Unit tests for parser, models, UI state
├── integration/         # End-to-end navigation tests
└── fixtures/            # Test markdown files with various structures
```

**Structure Decision**: Extend existing structure. The existing `internal/` layout is appropriate. Add `renderer/` package for content display, extend `models/` with section tree support.

## Complexity Tracking

> No Constitution violations. All changes align with existing architecture and principles.

| Decision | Rationale |
|----------|-----------|
| Extend Heading vs new Section struct | Keep Heading for flat TOC, add Section for tree navigation to maintain backward compatibility |
| Add viewMode to Model | Clean separation between TOC navigation and content viewing states |
| New renderer package | Separation of concerns - UI handles interaction, renderer handles markdown-to-display conversion |

## Phase 0: Research

No external research needed. The implementation uses existing technologies:
- Goldmark parser already integrated
- Bubble Tea patterns well-established
- Line number calculation from existing parser

**research.md**: Mark as N/A - using established patterns from existing codebase.

## Phase 1: Design

### Data Model

Key entities from spec:

**Section** (extends current Heading concept):
```go
type Section struct {
    ID         string     // Unique identifier for navigation
    Level      int        // 1-6 (H1-H6)
    Title      string     // Heading text
    StartLine  int        // Line where heading starts (1-based)
    EndLine    int        // Line where section ends (next heading or EOF)
    Content    string     // Raw markdown content of this section
    Children   []*Section // Nested subsections
    Parent     *Section   // Parent section (nil for H1)
}

type SectionTree struct {
    Root     *Section // Virtual root containing all H1 sections
    Source   string   // File path
    Sections []*Section // Flat list for indexed access
}
```

**Enhanced Model** (extends ui.Model):
```go
type ViewMode int
const (
    ViewTOC ViewMode = iota
    ViewContent
)

type NavigationState struct {
    SelectedIndex    int          // Current cursor position
    ExpandedSections map[string]bool // Section IDs currently expanded
    ViewMode         ViewMode     // Current display mode
    CurrentSection   *Section     // Section being viewed (when in content mode)
    LevelFilter      int          // 0 = no filter, >0 = max level to show
}
```

### CLI Contract

**Interface**: `mdexplore <file> [flags]`

**Flags**:
- `--level N` / `-L N`: Limit TOC to headings up to level N (default: 0 = no limit)
- `--toc`: Display table of contents (existing, default true)
- `--help` / `-h`: Display help
- `--version` / `-v`: Display version

**Exit Codes** (existing, preserved):
- 0: Success
- 1: General error
- 2: File not found
- 3: Permission denied
- 4: Invalid file
- 5: Parse error

**Usage Examples**:
```bash
# Show all headings
mdexplore README.md

# Show only H1 and H2
mdexplore README.md --level 2

# Show only top-level sections
mdexplore -L 1 docs/guide.md
```

**Keyboard Navigation** (TUI mode):
- `↑`/`k`: Navigate up
- `↓`/`j`: Navigate down
- `←`: Expand/collapse selected section
- `→`: Collapse selected section (optional)
- `Enter`: View section content
- `Esc`: Return from content view / quit from TOC
- `q`: Quit

### Quickstart

```bash
# Build the tool
go build -o mdexplore ./cmd/mdexplore

# View TOC with line numbers
./mdexplore README.md

# Limit depth to H1 and H2
./mdexplore --level 2 README.md

# Navigate in TUI
# - Use ↑/↓ to select a section
# - Press ← to expand and see subsections
# - Press Enter to view section content
# - Press Esc to return to navigation
# - Press q to quit
```

## Phase 2: Tasks

To be generated by `/speckit.tasks` command.

## Design Decisions

### Line Number Calculation
- Use existing Goldmark parser's line position data
- Calculate EndLine as: next heading's StartLine - 1, or EOF if last section
- Store line numbers in Section struct, display as "L[start]-[end]" in TOC

### Hierarchical Navigation
- Start with only H1 sections visible
- Left arrow expands selected section, revealing its immediate children
- Each section tracks its own expanded/collapsed state
- Selection index operates on the flattened visible list

### Content View
- When Enter pressed on section: switch to ViewContent mode
- Render section.Content with markdown formatting preserved
- Use glamour or similar for terminal markdown rendering (if not already present)
- Escape returns to ViewTOC mode with selection preserved

### Level Flag
- Applied at parser level: filter headings before building TOC/SectionTree
- Affects both TUI and non-TTY output modes
- 0 = no limit (default), 1-6 = max level to include
