# Feature Specification: Nested Section Numbering and Select Option

**Feature Branch**: `003-nested-numbering`
**Created**: 2026-02-26
**Status**: Draft
**Input**: User description: "add nested numbering (1., 1.1, 1.1.1) display for sections. Add --select CLI option to pre-select a section at opening (e.g., --select 1.1). When ESC is pressed in content view, close the display and return to navigation. In TTY mode with --select, print only the selected section and its subsections."

---

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Nested Numbering Display (Priority: P1)

As a coding agent, I want to see hierarchical section numbering (1., 1.1, 1.1.1) in the TOC so that I can easily reference and navigate to specific sections using a clear numbering scheme.

**Why this priority**: Numbering provides a universal reference system that is easier to communicate than line numbers or titles alone. It helps users understand the document structure at a glance.

**Independent Test**: Can be tested by running mdexplore on any markdown file and verifying that each TOC entry displays its hierarchical number (e.g., "1. Introduction", "1.1. Getting Started").

**Acceptance Scenarios**:

1. **Given** a markdown file with nested sections, **When** the TOC is displayed, **Then** each H1 heading shows as "1. Title", "2. Title", etc.
2. **Given** a section with H2 children, **When** displayed in TOC, **Then** H2 headings show as "1.1. Title", "1.2. Title", etc.
3. **Given** deeply nested sections (H3, H4), **When** displayed in TOC, **Then** headings show appropriate depth (e.g., "1.1.1. Title", "1.1.1.1. Title").
4. **Given** skipped heading levels (H1 → H3), **When** displayed in TOC, **Then** numbering continues logically (e.g., "1.1.1. Title" even if H2 is missing).
5. **Given** the non-TTY output mode, **When** TOC is printed, **Then** numbering is included in the output.

---

### User Story 2 - Pre-select Section with --select (Priority: P1)

As a coding agent, I want to use a `--select` option to jump directly to a specific section so that I can quickly access the content I need without manual navigation.

**Why this priority**: This feature significantly speeds up workflows where users know exactly which section they want to view, eliminating the need for manual navigation through the TOC.

**Independent Test**: Can be tested by running `mdexplore file.md --select 1.1` and verifying that the specified section is pre-selected and its content is displayed.

**Acceptance Scenarios**:

1. **Given** a valid section number like "1.1", **When** `--select 1.1` is specified, **Then** the TUI opens with section 1.1 pre-selected.
2. **Given** an invalid section number, **When** `--select 99.99` is specified, **Then** an appropriate error message is shown.
3. **Given** a parent section number, **When** `--select 1` is specified, **Then** the TUI opens with section 1 selected and its children visible.
4. **Given** the `--select` option combined with `--level`, **When** both flags are used, **Then** the level filter is applied first, then the selection is made from visible sections.

---

### User Story 3 - ESC Returns to Navigation (Priority: P1)

As a coding agent, I want to press ESC to close the content view and return to TOC navigation so that I can continue exploring other sections.

**Why this priority**: This provides a smooth navigation flow - view content, press ESC to go back, select another section. This is standard behavior in many TUI applications.

**Independent Test**: Can be tested by launching the TUI, viewing a section's content, pressing ESC, and verifying the TOC is displayed again.

**Acceptance Scenarios**:

1. **Given** a section's content is being displayed, **When** ESC is pressed, **Then** the content view closes and the TOC is shown.
2. **Given** the content view was opened via --select, **When** ESC is pressed, **Then** navigation returns to TOC at the previously selected section.
3. **Given** the TOC is being displayed, **When** ESC is pressed, **Then** the application quits (existing behavior preserved).

---

### User Story 4 - TTY Mode with --select (Priority: P1)

As a coding agent, I want to use `--select` in non-TTY mode to print only the selected section and its subsections so that I can extract specific parts of documentation programmatically.

**Why this priority**: This enables scripting workflows where users can extract specific sections for processing, documentation generation, or integration with other tools.

**Independent Test**: Can be tested by running `mdexplore file.md --select 1.1 > output.txt` and verifying only that section and its subsections are in the output.

**Acceptance Scenarios**:

1. **Given** a valid section number and non-TTY mode, **When** `--select 1.1` is used with output redirection, **Then** only section 1.1 and its subsections are printed.
2. **Given** a parent section number, **When** `--select 1` is used, **Then** section 1 and all its descendants are printed.
3. **Given** the --select option, **When** used with line numbers, **Then** the output includes both numbering and line ranges.

---

### Edge Cases

- What happens when section numbers exceed single digits (10., 10.1, etc.)?
- How are sections numbered when using --level filtering (does numbering reflect filtered or full structure)?
- What happens when --select points to a section that was filtered out by --level?
- How should we handle malformed section numbers in --select (e.g., "1.1.1.1.1" when only 3 levels exist)?

---

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST display hierarchical section numbering in format "N.", "N.N.", "N.N.N." etc.
- **FR-002**: Numbering MUST be displayed alongside line numbers in format "1. L5-8 Title"
- **FR-003**: System MUST accept a `--select` flag with a section number (e.g., "1.1") to pre-select that section
- **FR-004**: When `--select` is used, the TUI MUST open with the specified section selected and its content displayed
- **FR-005**: When ESC is pressed in content view, System MUST return to TOC navigation
- **FR-006**: When `--select` is used in non-TTY mode, System MUST print only the selected section and its subsections
- **FR-007**: System MUST validate section numbers and provide clear error messages for invalid selections
- **FR-008**: Numbering MUST be consistent between TTY and non-TTY output modes

### Key Entities

- **SectionNumber**: Represents a hierarchical section identifier:
  - `parts`: []int - The number components (e.g., [1, 1] for "1.1")
  - `depth`: int - The depth level (number of parts)
  - Methods: `String()`, `Parent()`, `IsAncestorOf()`, `Equals()`

- **NumberedSection**: Extends Section with display numbering:
  - `Section`: *Section - Reference to underlying section
  - `Number`: SectionNumber - The hierarchical number
  - `DisplayNumber`: string - Formatted string (e.g., "1.1.")

- **Selection**: Represents a section selection:
  - `TargetNumber`: SectionNumber - The requested section
  - `TargetSection`: *Section - The resolved section (nil if not found)
  - `IsValid`: bool - Whether the selection is valid

---

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users can identify any section using hierarchical numbering within 1 second
- **SC-002**: Users can navigate to a specific section using --select in under 3 seconds
- **SC-003**: 100% accuracy in section numbering across all heading levels (H1-H6)
- **SC-004**: --select works correctly for 100% of valid section numbers in the document
- **SC-005**: Error messages for invalid --select values are clear and actionable
