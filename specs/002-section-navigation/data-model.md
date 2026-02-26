# Data Model: Enhanced Section Navigation

**Feature**: Enhanced Section Navigation
**Date**: 2026-02-25
**Status**: Draft

---

## Entity: Section

Represents a hierarchical document section with heading, line range, and content.

### Fields

| Field | Type | Description |
|-------|------|-------------|
| ID | `string` | Unique identifier (format: "L{StartLine}") |
| Level | `int` | Heading level 1-6 (H1-H6) |
| Title | `string` | Clean heading text (no markdown markers) |
| StartLine | `int` | Line number where heading starts (1-based) |
| EndLine | `int` | Line number where section ends (inclusive) |
| RawContent | `string` | Raw markdown content between StartLine and EndLine |
| Children | `[]*Section` | Nested subsections (immediate children only) |
| Parent | `*Section` | Parent section (nil for H1 sections) |

### Validation Rules

- `Level` must be 1-6
- `StartLine` must be >= 1
- `EndLine` must be >= `StartLine`
- `Title` must be non-empty
- `Children` must all have `Level = Parent.Level + 1` (or greater for skipped levels)
- Parent relationship must be reciprocal (child.Parent == parent)

### Methods

```go
// HasChildren returns true if section has nested subsections
func (s *Section) HasChildren() bool

// GetAllDescendants returns flattened list of all nested sections recursively
func (s *Section) GetAllDescendants() []*Section

// GetVisibleDescendants returns immediate children (for expansion display)
func (s *Section) GetVisibleDescendants() []*Section

// ContainsLine returns true if the given line number is within this section
func (s *Section) ContainsLine(line int) bool

// GetDepth returns the nesting depth (0 for H1, 1 for H2, etc.)
func (s *Section) GetDepth() int
```

---

## Entity: SectionTree

Container for the complete hierarchical structure of a document.

### Fields

| Field | Type | Description |
|-------|------|-------------|
| Root | `*Section` | Virtual root node containing all H1 sections as children |
| Source | `string` | Absolute path to source file |
| Sections | `[]*Section` | Flattened list of all sections for indexed access |
| ByID | `map[string]*Section` | Lookup map from ID to section |

### Methods

```go
// GetH1Sections returns all top-level sections
func (t *SectionTree) GetH1Sections() []*Section

// FindByID looks up a section by its unique ID
func (t *SectionTree) FindByID(id string) (*Section, bool)

// GetFlattenedVisible returns sections in display order respecting expansion state
func (t *SectionTree) GetFlattenedVisible(expanded map[string]bool) []*Section

// FilterByMaxLevel returns a new tree with only sections up to given level
func (t *SectionTree) FilterByMaxLevel(maxLevel int) *SectionTree
```

---

## Entity: NavigationState

Tracks the current TUI navigation context and view state.

### Fields

| Field | Type | Description |
|-------|------|-------------|
| SelectedIndex | `int` | Current cursor position in visible list |
| ExpandedSections | `map[string]bool` | Set of section IDs currently expanded |
| ViewMode | `ViewMode` | Current display mode (TOC or Content) |
| CurrentSection | `*Section` | Section being viewed (when in Content mode) |
| LevelFilter | `int` | 0 = no filter, >0 = max heading level to show |

### Types

```go
type ViewMode int

const (
    ViewTOC ViewMode = iota      // Table of contents navigation view
    ViewContent                   // Section content display view
)
```

### Methods

```go
// IsExpanded returns true if given section is expanded
func (n *NavigationState) IsExpanded(sectionID string) bool

// ToggleExpanded flips the expansion state of a section
func (n *NavigationState) ToggleExpanded(sectionID string)

// Expand marks a section as expanded
func (n *NavigationState) Expand(sectionID string)

// Collapse marks a section as collapsed
func (n *NavigationState) Collapse(sectionID string)

// CanNavigateUp returns true if selection can move up
func (n *NavigationState) CanNavigateUp(visibleCount int) bool

// CanNavigateDown returns true if selection can move down
func (n *NavigationState) CanNavigateDown(visibleCount int) bool
```

---

## Entity: DisplaySection

View-model for rendering a section in the TUI (combines Section + display state).

### Fields

| Field | Type | Description |
|-------|------|-------------|
| Section | `*Section` | Reference to underlying section |
| DisplayIndex | `int` | Position in the currently visible list |
| Depth | `int` | Visual indentation level (0 for H1, etc.) |
| IsSelected | `bool` | Whether this item has keyboard focus |
| IsExpanded | `bool` | Whether this section's children are visible |
| CanExpand | `bool` | Whether this section has children to expand |

### Display Format

```
TOC View:
  L1-45   ● Introduction              (H1, not selected)
  L47-89  ├── Getting Started         (H2, selected, not expanded)
  L91-120 ├── Architecture            (H2, expanded)
  L93-110 │   └── Overview            (H3, visible because parent expanded)

Content View:
  ┌─ L47-89: Getting Started ─────────────────────┐
  │                                                │
  │  [Section content rendered with formatting]   │
  │                                                │
  └─ Press Esc to return to navigation ───────────┘
```

---

## Relationships

```
SectionTree
    └── Root (virtual)
        └── Sections []*Section (H1)
            └── Each has Children []*Section (H2)
                └── Each has Children []*Section (H3)
                    └── ... (recursively nested)

NavigationState
    └── references SectionTree
    └── tracks ExpandedSections (section IDs)
    └── references CurrentSection (when viewing content)

DisplaySection
    └── references Section
    └── references NavigationState for display properties
```

---

## State Transitions

### Navigation Flow

```
[TOC View, H1 only visible]
    ↓ (select section, press ←)
[TOC View, section expanded, children visible]
    ↓ (navigate to child, press ←)
[TOC View, subsection expanded, grandchildren visible]
    ↓ (press Enter on any section)
[Content View, section content displayed]
    ↓ (press Esc)
[TOC View, previous state restored]
```

### Level Filter Application

```
[Parse Document]
    ↓
[Build Full SectionTree (all levels)]
    ↓
[Apply LevelFilter if > 0]
    ↓
[Display Filtered Sections]
```

---

## File Organization

```
internal/models/
├── heading.go       # Existing Heading struct (preserved for compatibility)
├── section.go       # NEW: Section, SectionTree, NavigationState
└── section_test.go  # NEW: Unit tests for section operations
```
