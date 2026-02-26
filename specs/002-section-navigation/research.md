# Research: Enhanced Section Navigation

**Feature**: Enhanced Section Navigation
**Date**: 2026-02-25
**Status**: N/A - Using Established Patterns

---

## Summary

No external research required. This feature builds entirely on existing technologies and patterns already established in the codebase.

---

## Technology Decisions

### Parser: Goldmark (Existing)

**Decision**: Continue using Goldmark for markdown parsing.

**Rationale**:
- Already integrated in codebase
- Provides line position data needed for line number extraction
- Supports CommonMark and common extensions
- No additional dependencies needed

### TUI Library: Bubble Tea (Existing)

**Decision**: Continue using Bubble Tea for terminal UI.

**Rationale**:
- Already integrated in codebase
- Supports hierarchical navigation patterns via component state
- Message-based architecture supports view mode switching
- Existing key handling can be extended for new navigation keys

### Content Rendering: Glamour (To Add)

**Decision**: Add Glamour for markdown content rendering in the TUI.

**Rationale**:
- Industry standard for terminal markdown rendering
- Produces beautiful, formatted output
- Integrates well with Bubble Tea
- Supports syntax highlighting for code blocks
- Trade-off: Additional dependency, but justified for quality content display

**Alternatives Considered**:
- Render raw markdown: Too hard to read
- Strip markdown: Loses formatting
- Custom renderer: Too complex, reinvents wheel

---

## Design Patterns

### Hierarchical Navigation

**Pattern**: Tree view with expansion state

**Implementation**:
- Each section tracks its own expanded/collapsed state
- Visible list is computed by flattening tree based on expansion state
- Selection index operates on flattened visible list
- This pattern is common in file explorers and tree widgets

### View Mode Switching

**Pattern**: State machine with two modes (TOC vs Content)

**Implementation**:
- `ViewMode` enum with `ViewTOC` and `ViewContent` values
- Different render functions for each mode
- Escape key transitions from Content back to TOC
- State restoration preserves navigation context

---

## Conclusion

All technologies and patterns are established and well-understood. Proceeding directly to implementation.
