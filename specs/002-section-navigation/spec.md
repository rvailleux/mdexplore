# Feature Specification: Enhanced Section Navigation

**Feature Branch**: `002-section-navigation`
**Created**: 2026-02-25
**Status**: Draft
**Input**: User description: "add line start and end number in front of each section. Add a --level X (or -L X) option to limit title to deepth X (default no limit). Add a navigation tui that enables the user to navigate in the sections : at first, only top level sections title is listed, then on selection of a section then typing the left arrow, it displays the sublevel titles included in the section. Typing enter on a selection cat the whole section in a readable format, esc close the section display and makes the navigation to continue."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Display Line Numbers in TOC (Priority: P1)

As a coding agent, I want to see the line number range for each section in the table of contents so that I can quickly locate sections in large markdown files.

**Why this priority**: Line numbers are essential for agents to quickly reference specific sections when working with files. This is the foundational feature that enables precise navigation.

**Independent Test**: Can be tested by running mdexplore on any markdown file and verifying that each TOC entry displays its start and end line numbers.

**Acceptance Scenarios**:

1. **Given** a markdown file with multiple sections, **When** the TOC is displayed, **Then** each heading shows its start and end line numbers (e.g., "L45-78 Introduction")
2. **Given** the last section in a file, **When** displayed in TOC, **Then** it shows start line to end of file line number
3. **Given** a single-line section, **When** displayed in TOC, **Then** it shows the same line number for start and end (e.g., "L23-23 Heading")

---

### User Story 2 - Limit TOC Depth with --level Flag (Priority: P1)

As a coding agent, I want to limit the TOC display to a specific depth level so that I can focus on high-level structure without being overwhelmed by nested subsections.

**Why this priority**: Large documentation files often have deeply nested structures. Limiting depth helps agents quickly grasp the overall document organization.

**Independent Test**: Can be tested by running `mdexplore file.md --level 2` and verifying only headings up to level 2 are displayed.

**Acceptance Scenarios**:

1. **Given** a markdown file with headings up to level 4, **When** `--level 2` is specified, **Then** only headings level 1 and 2 are displayed
2. **Given** a level limit of 1, **When** TOC is generated, **Then** only top-level (H1) headings are shown
3. **Given** no level flag, **When** TOC is generated, **Then** all heading levels are displayed (default: no limit)
4. **Given** `-L 3` shorthand flag, **When** TOC is generated, **Then** headings up to level 3 are shown

---

### User Story 3 - Hierarchical Section Navigation TUI (Priority: P1)

As a coding agent, I want to navigate through document sections hierarchically so that I can efficiently explore large documents without cognitive overload.

**Why this priority**: This is the core navigation experience. Starting with only top-level sections allows agents to understand document structure at a glance before drilling into details.

**Independent Test**: Can be tested by launching the TUI and verifying the navigation flow: top-level list → expand with left arrow → view content with enter → return with esc.

**Acceptance Scenarios**:

1. **Given** the TUI is launched, **When** navigation starts, **Then** only top-level (H1) sections are listed initially
2. **Given** a top-level section is selected, **When** the left arrow key is pressed, **Then** the section expands to show its nested subsections (H2, H3, etc.)
3. **Given** a section with subsections is expanded, **When** the user navigates to a subsection and presses left arrow, **Then** that subsection further expands to show its children
4. **Given** any section is selected, **When** Enter is pressed, **Then** the full content of that section is displayed in a readable format
5. **Given** section content is being displayed, **When** Escape is pressed, **Then** the content view closes and returns to the navigation view with previous selection state preserved

---

### Edge Cases

- What happens when a section has no content between it and the next section?
- How does the system handle circular navigation at the top and bottom of the list?
- What happens when Enter is pressed on a section that contains only subsections and no direct content?
- How are empty sections (headings with no content) displayed?
- What happens when the --level flag is combined with the navigation TUI?
- How does the system handle malformed markdown with skipped heading levels (e.g., H1 followed by H3)?

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST display the start and end line number for each section in the TOC format "L[start]-[end] [Heading Text]"
- **FR-002**: System MUST accept a `--level` (or `-L`) flag to limit the maximum heading depth displayed in TOC
- **FR-003**: System MUST default to showing all heading levels when no `--level` flag is provided
- **FR-004**: System MUST display only top-level (H1) sections when the TUI navigation first launches
- **FR-005**: System MUST expand a selected section to show its subsections when the left arrow key (←) is pressed
- **FR-006**: System MUST display the full content of a selected section when Enter is pressed
- **FR-007**: System MUST return to the navigation view when Escape is pressed from the content view
- **FR-008**: System MUST preserve navigation state (current position, expanded sections) when returning from content view
- **FR-009**: System MUST support collapsing an expanded section (optional enhancement via right arrow or second left arrow press)
- **FR-010**: System MUST render section content in a readable format with proper formatting preserved

### Key Entities

- **Section**: Represents a document section with a heading, containing:
  - `level`: The heading level (1-6)
  - `title`: The heading text
  - `startLine`: The line number where the section starts
  - `endLine`: The line number where the section ends
  - `content`: The markdown content between start and end lines
  - `children`: Nested subsections within this section
  - `parent`: Reference to parent section (if any)

- **NavigationState**: Tracks the current TUI navigation context:
  - `selectedIndex`: Currently highlighted item index
  - `expandedSections`: Set of section IDs that are expanded
  - `viewMode`: Current display mode ("toc" or "content")
  - `currentContent`: Content being displayed when in content view

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users can identify the exact line range of any section within 2 seconds of viewing the TOC
- **SC-002**: Users can limit TOC depth to reduce displayed items by 50% or more on deeply nested documents
- **SC-003**: Users can navigate to any section content and return to navigation in under 5 seconds
- **SC-004**: The hierarchical navigation reduces cognitive load by showing at most 7±2 items at the initial view
- **SC-005**: 100% of markdown heading levels (H1-H6) are correctly parsed and assigned accurate line numbers
