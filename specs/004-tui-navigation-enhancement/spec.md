# Feature Specification: TUI Navigation Enhancement

**Feature Branch**: `004-tui-navigation-enhancement`
**Created**: 2026-02-26
**Status**: Draft
**Input**: User description: "when pressing right arrow on a 'leaf' section, do the same as 'enter' (open the content). Place the nested numbering next to each section's title. When typing 'enter' to open a section, display the full section, including the sub sections titles and contents. when pressing the left arrow, close the content and go back to navigation"

---

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Right Arrow on Leaf Opens Content (Priority: P1)

As a coding agent, I want to press the right arrow key on a leaf section (section without children) to open its content, just like pressing Enter, so that I can navigate more intuitively.

**Why this priority**: This provides a consistent navigation experience where the right arrow always "drills down" into content - expanding for parent sections, opening content for leaf sections. This matches the mental model of right = deeper/into.

**Independent Test**: Can be tested by launching the TUI, navigating to any leaf section (one without children), pressing the right arrow, and verifying the content view opens.

**Acceptance Scenarios**:

1. **Given** a leaf section is selected (no children), **When** the right arrow key is pressed, **Then** the content view opens showing the section content.
2. **Given** a parent section is selected (has children), **When** the right arrow key is pressed, **Then** the section expands to show children (existing behavior preserved).
3. **Given** content view is open from a leaf section via right arrow, **When** the user navigates back, **Then** the cursor returns to the previously selected section.

---

### User Story 2 - Nested Numbering Next to Title (Priority: P1)

As a coding agent, I want to see the section number (e.g., "1.1.") displayed directly next to the section title rather than at the beginning of the line, so that the visual hierarchy is clearer and the title is easier to read.

**Why this priority**: Moving the numbering next to the title creates a cleaner visual layout where the tree structure (indentation and tree characters) is separated from the content identification (number + title), making it easier to scan the document structure.

**Independent Test**: Can be tested by running mdexplore on any markdown file and verifying that section numbers appear immediately before the title text, after the tree prefix characters.

**Acceptance Scenarios**:

1. **Given** the TOC is displayed, **When** viewing any section, **Then** the section number appears immediately before the title (e.g., "├── 1.1. Title" instead of "1.1. ├── Title").
2. **Given** a selected section, **When** viewing the highlighted line, **Then** the numbering style is consistent with the selection highlighting.
3. **Given** nested sections at various depths, **When** displayed, **Then** numbering is consistently positioned next to titles at all levels.

---

### User Story 3 - Full Section Display Including Subsections (Priority: P1)

As a coding agent, when I press Enter to view a section, I want to see the complete content including all subsection titles and their contents, so that I can read the entire logical unit without navigating back and forth.

**Why this priority**: Sections often form logical units where the parent provides context and subsections provide details. Viewing them together provides better reading flow and context preservation, especially for specifications and documentation.

**Independent Test**: Can be tested by opening a section with subsections via Enter and verifying that all descendant content is displayed in the content view.

**Acceptance Scenarios**:

1. **Given** a section with subsections exists, **When** Enter is pressed to view it, **Then** the content view displays the parent section heading and content, followed by all subsection headings and their contents recursively.
2. **Given** a leaf section (no subsections), **When** Enter is pressed, **Then** only that section's content is displayed (existing behavior).
3. **Given** a section with multiple levels of nesting, **When** viewed, **Then** all descendant content is included in the display, properly formatted with appropriate heading levels.
4. **Given** the full section content is displayed, **When** viewing the content, **Then** subsection headings are visually distinguished (e.g., with appropriate markdown heading markers or styling).

---

### User Story 4 - Left Arrow Returns to Navigation (Priority: P1)

As a coding agent, I want to press the left arrow key to close the content view and return to the TOC navigation, so that I can quickly get back to browsing sections.

**Why this priority**: The left arrow provides a natural "back" navigation that complements the right arrow/Enter "forward" navigation. It's more ergonomic than reaching for the Esc key and creates a bidirectional navigation model.

**Independent Test**: Can be tested by opening a section's content view, pressing the left arrow, and verifying the TOC is displayed with the cursor on the previously viewed section.

**Acceptance Scenarios**:

1. **Given** the content view is open, **When** the left arrow key is pressed, **Then** the content view closes and the TOC navigation is displayed.
2. **Given** the content view was opened via Enter or right arrow, **When** returning via left arrow, **Then** the cursor is positioned on the section that was being viewed.
3. **Given** the TOC is displayed (not in content view), **When** the left arrow is pressed on a collapsed section, **Then** the existing collapse behavior is preserved (no change).
4. **Given** the TOC is displayed, **When** the left arrow is pressed on an expanded section, **Then** the section collapses (existing behavior preserved).

---

### Edge Cases

- What happens when a section has no content (empty section)?
- How does the display handle very deeply nested subsections (hierarchical depth > 6)?
- What happens when the full section content exceeds the terminal height?
- How should the cursor position be restored when returning from content view if the section is now off-screen due to expansion changes?
- What happens when viewing a section that was filtered out by --level but is visible via parent?

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST open content view when right arrow is pressed on a leaf section (section without children)
- **FR-002**: System MUST preserve existing expand/collapse behavior when right arrow is pressed on parent sections
- **FR-003**: System MUST display section numbers immediately before the section title, after tree prefix characters
- **FR-004**: System MUST display the full hierarchical content when Enter is pressed, including the selected section and all its descendants
- **FR-005**: System MUST return to TOC navigation when left arrow is pressed in content view
- **FR-006**: System MUST position the cursor on the previously viewed section when returning from content view via left arrow
- **FR-007**: System MUST preserve existing collapse behavior for left arrow in TOC navigation mode
- **FR-008**: System MUST include subsection headings in the full section display with appropriate visual hierarchy

### Key Entities

- **ContentViewState**: Represents the current content display:
  - `Mode`: enum indicating current view (TOC vs Content)
  - `SelectedSection`: reference to the section being viewed
  - `Content`: the rendered full content including subsections
  - `ReturnIndex`: position in TOC to restore when returning

- **NavigationContext**: Tracks navigation state for bidirectional movement:
  - `SourceSection`: the section from which content view was entered
  - `ViewHistory`: stack of viewed sections for potential future enhancement

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users can open leaf section content using right arrow with 100% reliability
- **SC-002**: Section numbering is positioned next to titles consistently across all sections
- **SC-003**: Full section display includes 100% of descendant content when Enter is pressed
- **SC-004**: Left arrow returns to navigation in under 100ms with cursor on correct section
- **SC-005**: Navigation feels intuitive with arrow keys (right = in/down, left = out/up) for 90%+ of users
