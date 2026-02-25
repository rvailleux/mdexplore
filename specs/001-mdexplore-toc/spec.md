# Feature Specification: Markdown TOC Explorer

**Feature Branch**: `001-mdexplore-toc`
**Created**: 2025-02-25
**Status**: Draft
**Input**: User description: "create a cli command 'mdexplore mdfile.md --toc' that will display the titles and subtitles of the file."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Display Table of Contents (Priority: P1)

As a coding agent, I want to see a structured table of contents from a markdown file so I can quickly understand the document structure and navigate to relevant sections.

**Why this priority**: This is the core functionality requested. Without it, the feature delivers no value. The TOC display enables rapid document comprehension without reading the entire file.

**Independent Test**: Can be fully tested by running `mdexplore sample.md --toc` on a markdown file with known headings and verifying all titles/subtitles appear in the correct order and hierarchy.

**Acceptance Scenarios**:

1. **Given** a markdown file with headings (H1, H2, H3), **When** the user runs `mdexplore file.md --toc`, **Then** all headings are displayed in a hierarchical tree structure showing the document outline
2. **Given** a markdown file with no headings, **When** the user runs `mdexplore file.md --toc`, **Then** a clear message indicates no headings were found
3. **Given** a markdown file with nested headings (H1 containing H2 containing H3), **When** the TOC is displayed, **Then** the indentation visually represents the heading hierarchy levels

---

### User Story 2 - Handle File Errors Gracefully (Priority: P2)

As a coding agent, I want clear error messages when the markdown file cannot be read so I can quickly identify and fix the issue.

**Why this priority**: Error handling is essential for a good user experience, but secondary to the core functionality. Clear errors save debugging time.

**Independent Test**: Can be fully tested by running `mdexplore` with various invalid inputs (non-existent file, directory instead of file, permission-denied file) and verifying helpful error messages are shown.

**Acceptance Scenarios**:

1. **Given** a non-existent file path, **When** the user runs `mdexplore missing.md --toc`, **Then** an error message clearly states the file was not found
2. **Given** a path that is a directory (not a file), **When** the user runs `mdexplore ./dir --toc`, **Then** an error message explains that a file is required, not a directory
3. **Given** a file without read permissions, **When** the user runs `mdexplore restricted.md --toc`, **Then** an error message indicates permission was denied

---

### User Story 3 - Support Various Markdown Heading Formats (Priority: P3)

As a coding agent, I want the tool to recognize both ATX-style (`# Heading`) and Setext-style (underlined) headings so it works with any markdown file format.

**Why this priority**: While most modern markdown uses ATX-style headings, supporting Setext-style ensures broader compatibility with legacy or differently-formatted documents.

**Independent Test**: Can be fully tested by running `mdexplore` on files with ATX-style headings, Setext-style headings, and mixed formats, verifying all are correctly extracted.

**Acceptance Scenarios**:

1. **Given** a file with ATX-style headings (`# H1`, `## H2`, etc.), **When** the TOC is generated, **Then** all headings are extracted with correct levels
2. **Given** a file with Setext-style headings (underlined with `===` or `---`), **When** the TOC is generated, **Then** these headings are recognized as H1 and H2 respectively

---

### Edge Cases

- What happens when headings skip levels (e.g., H1 directly to H3)?
- How does the system handle special characters or emoji in heading text?
- What happens when the same heading text appears multiple times?
- How does the system handle very long heading text (e.g., > 200 characters)?
- What happens with headings inside code blocks (should they be excluded)?
- How does the system handle markdown files with frontmatter (YAML header)?

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST provide a CLI command named `mdexplore`
- **FR-002**: The command MUST accept a file path as the first positional argument
- **FR-003**: The command MUST support a `--toc` flag to trigger table of contents display
- **FR-004**: When `--toc` is specified, the system MUST parse the markdown file and extract all headings (H1-H6)
- **FR-005**: The system MUST display headings in a hierarchical tree structure that reflects heading levels
- **FR-006**: The system MUST support ATX-style headings (`# Heading`, `## Heading`, etc.)
- **FR-007**: The system MUST handle cases where the specified file does not exist with a clear error message
- **FR-008**: The system MUST handle cases where the specified path is not a file (e.g., directory) with a clear error message
- **FR-009**: The system MUST handle permission errors when reading files with a clear error message
- **FR-010**: The system MUST complete TOC generation for files up to 1MB in under 1 second

### Key Entities

- **MarkdownDocument**: Represents the input markdown file being analyzed
  - Attributes: filePath, content, headings
- **Heading**: Represents a single heading extracted from the document
  - Attributes: level (1-6), text, lineNumber
- **TableOfContents**: Represents the structured output of all headings
  - Attributes: headings (ordered list), hierarchical representation

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users can view the complete document structure of a 1000-line markdown file in under 1 second
- **SC-002**: 100% of valid markdown headings (H1-H6) in a file are accurately extracted and displayed
- **SC-003**: 100% of error scenarios (missing file, permission denied, invalid path) produce clear, actionable error messages
- **SC-004**: The heading hierarchy is visually accurate (proper indentation reflecting heading levels) in 100% of cases
- **SC-005**: Users can identify the main topics and subtopics of a document at a glance without scrolling through the full content

## Assumptions

- Markdown files are encoded in UTF-8
- The tool is invoked from a terminal/environment with standard output capabilities
- Headings inside code blocks should be excluded from TOC (as they are code, not document structure)
- Frontmatter (YAML headers) should be recognized and not treated as headings
