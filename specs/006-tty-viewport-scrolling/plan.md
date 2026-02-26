# Implementation Plan: TTY Viewport Scrolling

**Branch**: `006-tty-viewport-scrolling` | **Date**: 2026-02-26 | **Spec**: [spec.md](spec.md)
**Input**: Feature specification from `/specs/006-tty-viewport-scrolling/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/plan-template.md` for the execution workflow.

## Summary

Add vertical scrolling capability to the TTY content view using Bubble Tea's viewport component. When rendered markdown content exceeds terminal height, users can navigate using Up/Down arrows for line-by-line scrolling and Page Up/Page Down for faster navigation. A scroll position indicator will provide visual feedback on document location.

## Technical Context

**Language/Version**: Go 1.23+
**Primary Dependencies**: Bubble Tea (TUI framework), Bubbles (viewport component), Lipgloss (styling), Goldmark (markdown rendering)
**Storage**: N/A (in-memory content rendering)
**Testing**: Go standard testing package
**Target Platform**: Linux, macOS, Windows terminal environments
**Project Type**: CLI TUI application
**Performance Goals**: <50ms scroll response time for documents under 1,000 lines
**Constraints**: Support terminals as small as 24 lines height, 80 columns width
**Scale/Scope**: Handle documents up to 10,000 lines

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

| Principle | Status | Notes |
|-----------|--------|-------|
| **I. TDD** | ✅ PASS | All viewport scrolling logic will be developed via TDD - tests for scroll offset calculation, boundary checking, and viewport rendering will be written first |
| **II. Terminal User Interface** | ✅ PASS | Feature enhances TUI using Bubble Tea viewport component, maintains keyboard navigation, handles terminal resize |
| **III. Markdown Reader Focus** | ✅ PASS | Scrolling improves the core reading experience without adding editing capabilities |
| **IV. Agent-First Design** | ✅ PASS | Quick navigation aids agents reviewing long technical specifications and documentation |
| **V. Simplicity** | ✅ PASS | Uses standard Bubble Tea viewport component; no custom scroll logic beyond key mappings |

**Gate Decision**: ✅ ALL CHECKS PASS - Proceed to Phase 0

## Project Structure

### Documentation (this feature)

```text
specs/006-tty-viewport-scrolling/
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command)
├── data-model.md        # Phase 1 output (/speckit.plan command)
├── quickstart.md        # Phase 1 output (/speckit.plan command)
├── contracts/           # Phase 1 output (/speckit.plan command)
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)
```

### Source Code (repository root)

```text
cmd/mdexplore/
└── main.go              # Entry point (no changes)

internal/
├── ui/
│   ├── model.go         # Add ContentScrollOffset field, ViewportHeight tracking
│   ├── view.go          # Modify renderContent() to use viewport window
│   ├── update.go        # Add scroll key handlers (up/down/pgup/pgdown)
│   └── viewport.go      # NEW: Viewport helper for content windowing
├── models/
│   └── section.go       # No changes
├── parser/
│   └── parser.go        # No changes
└── renderer/
    └── renderer.go      # No changes

tests/
├── unit/
│   └── viewport_test.go # NEW: Unit tests for viewport calculations
├── integration/
│   └── scrolling_test.go # NEW: Integration tests for scroll behavior
└── benchmark/
    └── scroll_bench_test.go # NEW: Performance benchmarks
```

**Structure Decision**: This is a focused enhancement to existing UI components. The main changes are in `internal/ui/` with the addition of viewport state tracking and scroll key handling. A new `viewport.go` helper will encapsulate viewport windowing logic.

## Constitution Check (Post-Phase 1 Re-evaluation)

After design and research phase, re-validating against constitution:

| Principle | Status | Notes |
|-----------|--------|-------|
| **I. TDD** | ✅ PASS | Unit tests for viewport calculations, integration tests for scroll behavior, and benchmarks defined in project structure |
| **II. Terminal User Interface** | ✅ PASS | Uses Bubble Tea viewport component, handles resize, keyboard navigation |
| **III. Markdown Reader Focus** | ✅ PASS | Pure reading enhancement - no editing features added |
| **IV. Agent-First Design** | ✅ PASS | Efficient navigation for long technical documents; vim-style keybindings |
| **V. Simplicity** | ✅ PASS | Minimal changes to existing code; leverages standard library patterns |

**Post-Design Gate Decision**: ✅ ALL CHECKS PASS - Ready for task generation

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

No violations. All constitution principles are satisfied with the planned approach.

## Generated Artifacts

| Artifact | Path | Status |
|----------|------|--------|
| Research | `research.md` | ✅ Complete |
| Data Model | `data-model.md` | ✅ Complete |
| Quickstart | `quickstart.md` | ✅ Complete |
| Contracts | `contracts/ui-keys.md` | ✅ Complete |
| Agent Context | `CLAUDE.md` | ✅ Updated |

## Next Steps

Run `/speckit.tasks` to generate the task breakdown for implementation.
