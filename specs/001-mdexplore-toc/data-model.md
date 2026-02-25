# Data Model: Markdown TOC Explorer

**Feature**: mdexplore TOC Display
**Date**: 2025-02-25

## Entities

### Heading

Represents a single heading extracted from a markdown document.

```go
type Heading struct {
    Level      int    // 1-6 (H1-H6)
    Text       string // Clean heading text (without markdown markers)
    LineNumber int    // Line number in source file (1-based)
}
```

**Validation Rules**:
- Level must be between 1 and 6 inclusive
- Text must not be empty after trimming
- LineNumber must be positive

### TableOfContents

Represents the complete structure of headings in a document.

```go
type TableOfContents struct {
    Headings []Heading // Ordered list of headings as they appear in document
    Source   string    // Path to source file
}
```

**Validation Rules**:
- Headings slice may be empty (valid for documents with no headings)
- Source must be a valid file path
- Headings must be in ascending LineNumber order

### ParseResult

Represents the result of parsing a markdown file.

```go
type ParseResult struct {
    TOC     TableOfContents
    Error   error  // nil on success
    Size    int64  // File size in bytes
}
```

## Relationships

```
TableOfContents 1--* Heading
ParseResult 1--1 TableOfContents
```

## State Transitions

### Parsing Flow

```
[File Path] → [Parse Markdown] → [Extract Headings] → [Build TOC]
     ↓               ↓                    ↓               ↓
  validate      Goldmark AST         filter code      sort/order
  exists        traversal            block headings   headings
```

### UI State Transitions

```
[Initial] → [Loading] → [Display TOC] → [Quit]
                ↓              ↓
             [Error] ← [File Not Found]
             [Error] ← [Permission Denied]
             [Error] ← [Parse Error]
```

## Heading Extraction Logic

1. **ATX-Style**: Lines starting with 1-6 `#` characters
   - Example: `## Section Title` → Level 2, Text "Section Title"

2. **Setext-Style**: Underlined text
   - Text followed by `===` → Level 1
   - Text followed by `---` → Level 2

3. **Exclusions**:
   - Headings inside fenced code blocks (```)
   - Headings inside indented code blocks
   - Headings in frontmatter (YAML between `---` markers at start)

## Tree Structure Building

The TOC is displayed hierarchically based on heading levels:

```
H1 (Level 1)
├── H2 (Level 2)
│   ├── H3 (Level 3)
│   └── H3 (Level 3)
└── H2 (Level 2)
    └── H3 (Level 3)
```

**Indentation Rules**:
- Each level adds 2 spaces of indentation
- Visual tree characters (`├──`, `└──`) enhance readability
