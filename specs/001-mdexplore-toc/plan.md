# Implementation Plan: Markdown TOC Explorer

**Branch**: `001-mdexplore-toc` | **Date**: 2025-02-25 | **Spec**: [spec.md](spec.md)
**Input**: Feature specification from `/specs/001-mdexplore-toc/spec.md`

## Summary

Implement a CLI command `mdexplore` that displays a hierarchical table of contents from markdown files. The tool uses a TUI-based interface (Bubble Tea) to render an interactive, styled view of document headings. Core functionality includes parsing ATX and Setext-style headings, displaying them in a tree structure, and handling file errors gracefully.

## Technical Context

**Language/Version**: Go 1.23+
**Primary Dependencies**: Bubble Tea (TUI), Goldmark (markdown parsing), Lipgloss (styling)
**Storage**: N/A - file-based, no persistent storage
**Testing**: Go standard testing + testify/assert
**Target Platform**: Linux, macOS, Windows (terminal environments)
**Project Type**: CLI tool
**Performance Goals**: Parse and display TOC for 1MB files in under 1 second
**Constraints**: UTF-8 encoding, requires modern terminal with color support
**Scale/Scope**: Single-file processing, handles files up to 10MB

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

| Principle | Status | Notes |
|-----------|--------|-------|
| I. Test-Driven Development | ✅ PASS | All code will follow Red-Green-Refactor; tests written first |
| II. Terminal User Interface | ✅ PASS | Using Bubble Tea TUI framework (not plain console) |
| III. Markdown Reader Focus | ✅ PASS | Core feature is reading/displaying markdown structure |
| IV. Agent-First Design | ✅ PASS | Fast startup, clear output, optimized for spec/documentation reading |
| V. Simplicity | ✅ PASS | Single purpose tool, minimal dependencies, clear abstractions |

**Gate Result**: ✅ PASSED - All constitution principles satisfied.

## Project Structure

### Documentation (this feature)

```text
specs/001-mdexplore-toc/
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
├── main.go              # CLI entry point

internal/
├── parser/
│   ├── parser.go        # Markdown parsing logic (Goldmark-based)
│   ├── parser_test.go   # Parser unit tests
│   └── heading.go       # Heading data structures
├── ui/
│   ├── model.go         # Bubble Tea model
│   ├── view.go          # Rendering/styling (Lipgloss)
│   └── update.go        # Event handling
├── toc/
│   ├── toc.go           # Table of contents logic
│   └── toc_test.go      # TOC unit tests
└── errors/
    └── errors.go        # Error types and handling

tests/
├── unit/                # Unit tests (mirror internal structure)
├── integration/         # Integration tests
└── fixtures/            # Test markdown files

go.mod
go.sum
```

**Structure Decision**: Go project with `cmd/` for CLI entry point, `internal/` for packages organized by concern (parsing, UI, TOC logic). Tests mirror source structure with additional fixtures for integration testing.

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

No violations - all constitution principles pass.
