# Implementation Plan: TUI Navigation Enhancement

**Branch**: `004-tui-navigation-enhancement` | **Date**: 2026-02-26 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/004-tui-navigation-enhancement/spec.md`

## Summary

Enhance the TUI navigation experience with four key improvements:
1. Right arrow opens content on leaf sections (like Enter)
2. Move nested numbering next to section titles for better visual hierarchy
3. Display full section content including subsections when viewing
4. Left arrow returns from content view to TOC navigation

This is a UI/UX refinement feature that builds on existing navigation patterns without changing core architecture.

## Technical Context

**Language/Version**: Go 1.23+
**Primary Dependencies**: Bubble Tea (TUI framework), Goldmark (markdown parsing), Lipgloss (styling)
**Storage**: N/A (in-memory document parsing)
**Testing**: Go's standard testing package + test scripts
**Target Platform**: Linux, macOS, Windows terminal environments
**Project Type**: CLI tool with interactive TUI
**Performance Goals**: Navigation response < 100ms, startup < 500ms
**Constraints**: Terminal width handling, UTF-8 support, 256-color terminals
**Scale/Scope**: Single document navigation, typical files < 10MB

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

| Principle | Status | Notes |
|-----------|--------|-------|
| **I. Test-Driven Development** | ✅ PASS | All navigation behaviors can be unit tested with Bubble Tea's test framework |
| **II. Terminal User Interface** | ✅ PASS | Feature enhances existing TUI navigation patterns |
| **III. Markdown Reader Focus** | ✅ PASS | Improvements directly support reading/navigation workflows |
| **IV. Agent-First Design** | ✅ PASS | Bidirectional arrow navigation optimized for agent workflows |
| **V. Simplicity** | ✅ PASS | Builds on existing patterns, no new dependencies needed |

**Verdict**: All constitution principles satisfied. Proceed with implementation.

## Project Structure

### Documentation (this feature)

```text
specs/004-tui-navigation-enhancement/
├── plan.md              # This file
├── research.md          # Phase 0 output (no research needed - straightforward UI changes)
├── data-model.md        # Phase 1 output (minimal model changes)
├── quickstart.md        # Phase 1 output (updated key bindings)
└── contracts/           # N/A - no external interfaces
```

### Source Code (repository root)

```text
cmd/mdexplore/
└── main.go              # No changes needed

internal/
├── models/
│   ├── section.go       # May need helper for descendant content extraction
│   └── navigation.go    # No changes needed
├── ui/
│   ├── model.go         # Add ReturnIndex to Model for cursor restoration
│   ├── update.go        # Modify key handlers (Right, Left, Enter)
│   └── view.go          # Modify renderSectionWithNumber, renderContent
└── parser/
    └── parser.go        # May need helper for full section content extraction

tests/
├── unit/
│   └── ui_test.go       # Add tests for new key handlers
└── integration/
    └── navigation_test.go # Test full navigation flows
```

**Structure Decision**: Minimal changes to existing structure. Primary work in `internal/ui/` package for view rendering and key handling. May add utility methods to models for content aggregation.

## Complexity Tracking

> **No constitution violations. No complexity justification needed.**

| Aspect | Assessment |
|--------|------------|
| Files Modified | 3-4 files in `internal/ui/` |
| New Dependencies | None |
| Architecture Changes | None |
| Breaking Changes | None (enhances existing behavior) |
| Test Coverage | All new behaviors require tests per TDD |

## Phase 0: Research

**Status**: SKIPPED - No research needed

**Rationale**: This feature consists of straightforward UI behavior modifications:
- Right arrow behavior: Simple conditional check (has children → expand, no children → view content)
- Numbering placement: Cosmetic change to render order in view layer
- Full section display: Recursive content aggregation from existing data structures
- Left arrow navigation: Add key binding to existing content view exit logic

All patterns are well-established in the codebase. No new technologies, libraries, or architectural patterns required.

## Phase 1: Design

### Data Model Changes

**Minimal changes required**:

1. **Model struct** (`internal/ui/model.go`):
   - Add `ReturnIndex int` to store cursor position when entering content view
   - Ensures cursor restoration when returning via left arrow

2. **Section helper** (`internal/models/section.go`):
   - Add `GetAllDescendants() []*Section` - returns flattened list of all nested sections
   - Add `GetFullContent() string` - aggregate RawContent from self + all descendants

### Interface Contracts

**No external interfaces** - this is an internal TUI enhancement.

### Key Bindings (for quickstart.md)

| Key | Action |
|-----|--------|
| `↑` / `↓` or `k` / `j` | Navigate up/down |
| `→` | Expand section OR open content (if leaf) |
| `←` | Collapse section OR return from content view |
| `Enter` | View full section content (including subsections) |
| `Esc` | Return from content view / Quit from TOC |
| `q` / `Ctrl+C` | Quit |

### Implementation Strategy

**Order of implementation (TDD)**:

1. **Numbering placement change** (simplest - pure UI)
   - Test: Verify render order
   - Implement: Move number rendering after tree prefix

2. **Right arrow on leaf** (key handler modification)
   - Test: Mock key press on leaf section
   - Implement: Conditional logic in Right key handler

3. **Left arrow navigation** (exit content view)
   - Test: Mock left arrow in content view mode
   - Implement: Add Left key handler for ViewContent mode

4. **Full section display** (content aggregation)
   - Test: Verify descendant content inclusion
   - Implement: Recursive content gathering

### Agent Context Update

Run `.specify/scripts/bash/update-agent-context.sh claude` after Phase 1 to update CLAUDE.md with any new technologies (none expected for this feature).

## Constitution Re-Check Post-Design

| Principle | Status | Verification |
|-----------|--------|--------------|
| **I. TDD** | ✅ PASS | Each behavior has testable acceptance criteria |
| **II. TUI** | ✅ PASS | All changes enhance terminal interface |
| **III. Markdown Focus** | ✅ PASS | Navigation improvements serve reading workflow |
| **IV. Agent-First** | ✅ PASS | Bidirectional arrows optimize agent navigation |
| **V. Simplicity** | ✅ PASS | No new dependencies, minimal code changes |

**Proceed to `/speckit.tasks` for task generation.**
