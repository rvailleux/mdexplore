# Research: Markdown TOC Explorer

**Feature**: mdexplore TOC Display
**Date**: 2025-02-25

## Technology Decisions

### Language Selection: Go

**Decision**: Use Go (Golang) for implementation

**Rationale**:
- Excellent TUI library ecosystem (Bubble Tea, Lipgloss, Bubbles)
- Fast startup times critical for CLI tools used by agents
- Single binary deployment - no runtime dependencies
- Strong standard library for file I/O and testing
- Native support for building cross-platform CLI tools

**Alternatives Considered**:
- Python with Rich/Textual: Larger distribution size, slower startup
- Node.js with Ink: Larger binary size via bundling, slower startup
- Rust with Ratatui: Excellent performance but steeper learning curve, longer compile times

### TUI Library: Bubble Tea

**Decision**: Use Charm's Bubble Tea framework with Lipgloss for styling

**Rationale**:
- Elm-inspired architecture - excellent for stateful TUI applications
- Built-in support for common components (viewport, list, spinner)
- Lipgloss provides declarative styling with CSS-like syntax
- Active community and good documentation
- Follows Go idioms and conventions

**Components to Use**:
- `bubbletea` - Core TUI framework
- `lipgloss` - Styling and layout
- `bubbles` - Pre-built components (list for TOC display)

### Markdown Parsing: Goldmark

**Decision**: Use yuin/goldmark - extensible CommonMark parser

**Rationale**:
- CommonMark compliant
- Extensible architecture for custom extensions
- Good performance characteristics
- Widely used in Go ecosystem (Hugo, etc.)
- Supports AST traversal for heading extraction

### Testing: Go Standard Library

**Decision**: Use Go's built-in testing package with testify/assert

**Rationale**:
- Follows TDD principle - tests are first-class citizens
- Native integration with Go toolchain
- testify/assert for readable assertions
- No external test runners needed

## Architecture Decisions

### Application Pattern

The application will follow a layered architecture:

1. **CLI Layer**: Command-line argument parsing using `spf13/cobra` or standard `flag` package
2. **Parser Layer**: Goldmark-based markdown parsing with custom AST visitor
3. **Model Layer**: Bubble Tea model for state management and TUI rendering
4. **View Layer**: Lipgloss-styled components for TOC display

### Heading Extraction Strategy

- Use Goldmark's AST visitor pattern to traverse the document
- Extract all Heading nodes (ATX and Setext styles supported by CommonMark)
- Track line numbers for potential future navigation features
- Filter out headings inside code blocks (FencedCodeBlock, CodeBlock contexts)

### Error Handling Strategy

- All errors bubble up to the Bubble Tea model
- Error state rendered as a styled error message in the TUI
- Non-interactive mode (future): errors written to stderr with exit codes

## Performance Considerations

- Target: Parse and display TOC for 1MB files in under 1 second
- Goldmark streaming parser - processes file in chunks
- No full document storage - headings extracted during parse
- Lazy rendering - only visible TOC items rendered

## Security Considerations

- File path validation to prevent directory traversal attacks
- Size limits to prevent memory exhaustion (configurable, default 10MB)
- No execution of code within markdown files

## Dependencies Summary

| Package | Purpose | Version |
|---------|---------|---------|
| github.com/charmbracelet/bubbletea | TUI framework | v1.1.0 |
| github.com/charmbracelet/lipgloss | Styling | v0.13.0 |
| github.com/charmbracelet/bubbles | UI components | v0.20.0 |
| github.com/yuin/goldmark | Markdown parser | v1.7.4 |
| github.com/spf13/cobra | CLI framework | v1.8.1 |
| github.com/stretchr/testify | Test assertions | v1.9.0 |

## Open Questions Resolved

1. **Q**: Should we support Setext-style headings?
   **A**: Yes - Goldmark supports them natively as part of CommonMark

2. **Q**: How to handle headings in code blocks?
   **A**: Exclude them - track AST context during traversal

3. **Q**: What output format for the TOC?
   **A**: Interactive TUI list with hierarchical indentation, styled with Lipgloss
