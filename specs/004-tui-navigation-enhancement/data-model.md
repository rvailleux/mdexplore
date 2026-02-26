# Data Model: TUI Navigation Enhancement

## Entity Changes

### Model (internal/ui/model.go)

**New Field**:
```go
ReturnIndex int  // Position in TOC to restore when returning from content view
```

**Purpose**: Store the cursor position before entering content view so it can be restored when user presses left arrow to return.

---

### Section (internal/models/section.go)

**New Methods**:

```go
// GetFullContent returns the combined RawContent of this section and all descendants
func (s *Section) GetFullContent() string

// IsLeaf returns true if section has no children
func (s *Section) IsLeaf() bool
```

**Rationale**:
- `GetFullContent()`: Aggregates content recursively for full section display (FR-004)
- `IsLeaf()`: Helper for determining right arrow behavior (FR-001)

---

## State Transitions

### Content View Entry

```
[TOC View] --Enter or Right(on leaf)--> [Content View]
   |                                         |
   | Store ReturnIndex = Selected            |
   v                                         v
[Model] ----------------------------> [Model]
 Selected = target index               ViewMode = ViewContent
                                       CurrentSection = target
                                       ReturnIndex = previous Selected
```

### Content View Exit

```
[Content View] --Left or Esc--> [TOC View]
      |                              |
      | Restore Selected = ReturnIndex|
      v                              v
[Model] ----------------------> [Model]
 ViewMode = ViewTOC              Selected = ReturnIndex
```

## Validation Rules

- `ReturnIndex` must be within bounds of visible sections when restoring
- `GetFullContent()` must maintain original markdown formatting
- `IsLeaf()` must be consistent with `HasChildren()` (inverse relationship)
