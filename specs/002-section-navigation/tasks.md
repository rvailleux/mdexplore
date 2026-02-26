# Tasks: Enhanced Section Navigation

**Input**: Design documents from `/specs/002-section-navigation/`
**Prerequisites**: plan.md, spec.md, data-model.md, contracts/cli.md, quickstart.md

**Tests**: Test tasks included per TDD Constitution requirement (Principle I).

**Organization**: Tasks are grouped by user story to enable independent implementation and testing.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (US1, US2, US3)
- Include exact file paths in descriptions

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Add Glamour dependency for markdown content rendering

- [X] T001 Add Glamour dependency to go.mod for markdown content rendering
- [X] T002 [P] Create test fixtures directory `tests/fixtures/` with sample markdown files

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core Section and SectionTree models that ALL user stories depend on

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

### Tests (TDD - Write First)

- [X] T003 [P] Create unit test for Section model in `internal/models/section_test.go` - Test Section struct creation and basic methods
- [X] T004 [P] Create unit test for SectionTree in `internal/models/section_test.go` - Test tree building from flat headings
- [X] T005 Create unit test for Section.GetAllDescendants() method in `internal/models/section_test.go`

### Implementation

- [X] T006 [P] Create Section struct with ID, Level, Title, StartLine, EndLine, RawContent, Children, Parent fields in `internal/models/section.go`
- [X] T007 [P] Create SectionTree struct with Root, Source, Sections, ByID fields in `internal/models/section.go`
- [X] T008 Implement Section.HasChildren(), GetAllDescendants(), ContainsLine() methods in `internal/models/section.go`
- [X] T009 Implement SectionTree.GetH1Sections(), FindByID(), GetFlattenedVisible() methods in `internal/models/section.go`

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Display Line Numbers in TOC (Priority: P1) 🎯 MVP

**Goal**: Display start and end line numbers for each section in TOC format "L[start]-[end] [Heading Text]"

**Independent Test**: Run `mdexplore README.md` and verify each TOC entry shows line range (e.g., "L45-78 Introduction")

### Tests (TDD - Write First)

- [X] T010 [P] [US1] Create parser test for EndLine calculation in `internal/parser/parser_test.go` - Test that last section ends at EOF
- [X] T011 [P] [US1] Create parser test for single-line section in `internal/parser/parser_test.go` - Test StartLine == EndLine for headings with no content
- [X] T012 [US1] Create integration test for line number display in `tests/integration/line_numbers_test.go`

### Implementation

- [X] T013 [US1] Extend GoldmarkParser.Parse() to calculate EndLine for each heading in `internal/parser/parser.go`
- [X] T014 [US1] Modify parser to build SectionTree with line ranges in `internal/parser/parser.go`
- [X] T015 [US1] Update renderHeading() to display "L[start]-[end]" format in `internal/ui/view.go`
- [X] T016 [US1] Update non-TTY printTOC() to show line ranges in `cmd/mdexplore/main.go`

**Status**: ✅ COMPLETE - Line numbers display correctly in both TTY and non-TTY modes

---

## Phase 4: User Story 2 - Limit TOC Depth with --level Flag (Priority: P1)

**Goal**: Add `--level` / `-L` flag to limit maximum heading depth displayed in TOC

**Independent Test**: Run `mdexplore -L 2 README.md` and verify only H1 and H2 headings are displayed

### Tests (TDD - Write First)

- [X] T017 [P] [US2] Create CLI test for --level flag parsing in `cmd/mdexplore/main_test.go`
- [X] T018 [P] [US2] Create parser test for level filtering in `internal/parser/parser_test.go`
- [X] T019 [US2] Create integration test for level-limited TOC output in `tests/integration/level_filter_test.go`

### Implementation

- [X] T020 [US2] Add --level/-L flag to Cobra root command in `cmd/mdexplore/main.go`
- [X] T021 [US2] Pass level filter through to parser in `cmd/mdexplore/main.go`
- [X] T022 [US2] Implement SectionTree.FilterByMaxLevel() method in `internal/models/section.go`
- [X] T023 [US2] Apply level filter in parser when building SectionTree in `internal/parser/parser.go`

**Status**: ✅ COMPLETE - Level filtering works correctly

---

## Phase 5: User Story 3 - Hierarchical Section Navigation TUI (Priority: P1)

**Goal**: Hierarchical navigation with expand/collapse (left arrow), content view (Enter), and return (Escape)

**Independent Test**: Launch TUI, verify: only H1 shown initially → left arrow expands → Enter shows content → Escape returns

### Tests (TDD - Write First)

- [X] T024 [P] [US3] Create UI model test for NavigationState in `internal/ui/model_test.go` - Test expansion state tracking
- [X] T025 [P] [US3] Create test for GetFlattenedVisible() with expansion in `internal/models/section_test.go`
- [ ] T026 [US3] Create integration test for keyboard navigation flow in `tests/integration/navigation_test.go`

### Implementation - Navigation State & Model

- [X] T027 [P] [US3] Create ViewMode type (ViewTOC, ViewContent) in `internal/ui/model.go`
- [X] T028 [P] [US3] Extend Model with NavigationState (ExpandedSections, ViewMode, CurrentSection) in `internal/ui/model.go`
- [X] T029 [US3] Implement NavigationState methods (IsExpanded, ToggleExpanded, Expand, Collapse) in `internal/ui/model.go`

### Implementation - Key Handlers

- [X] T030 [US3] Add left arrow handler to expand selected section in `internal/ui/update.go`
- [X] T031 [US3] Add right arrow handler to collapse selected section in `internal/ui/update.go`
- [X] T032 [US3] Add Enter handler to switch to content view in `internal/ui/update.go`
- [X] T033 [US3] Add Escape handler to return from content view to TOC in `internal/ui/update.go`

### Implementation - View Rendering

- [X] T034 [US3] Create content_view.go with content view rendering in `internal/ui/content_view.go`
- [X] T035 [US3] Implement GetFlattenedVisible() for hierarchical display in `internal/models/section.go`
- [X] T036 [US3] Update View() to handle ViewContent mode in `internal/ui/view.go`
- [X] T037 [US3] Update renderTOC() to show only H1 initially and expanded children in `internal/ui/view.go`
- [X] T038 [US3] Add visual indicators for expandable sections in `internal/ui/view.go`

### Implementation - Content Rendering

- [ ] T039 [US3] Create renderer package with markdown content rendering in `internal/renderer/renderer.go`
- [ ] T040 [US3] Implement section content extraction by line range in `internal/renderer/renderer.go`
- [ ] T041 [US3] Integrate Glamour for formatted content display in `internal/renderer/renderer.go`

**Status**: ⚠️ PARTIAL - Core navigation implemented, Glamour content rendering pending

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [ ] T042 [P] Update help text with new navigation keys in `cmd/mdexplore/main.go`
- [ ] T043 Add visual styling for expanded/collapsed states in `internal/ui/view.go`
- [ ] T044 Handle edge case: empty sections (no content between headings) in `internal/parser/parser.go`
- [ ] T045 Handle edge case: skipped heading levels (H1 -> H3) in `internal/parser/parser.go`
- [ ] T046 [P] Create test fixture with deeply nested structure in `tests/fixtures/nested.md`
- [ ] T047 [P] Create test fixture with skipped levels in `tests/fixtures/skipped_levels.md`
- [ ] T048 Run quickstart.md validation - verify all examples work

---

## Implementation Summary

### ✅ Completed (US1 & US2)
- Line numbers display in "L[start]-[end]" format
- `--level` / `-L` flag for depth limiting
- Section/SectionTree hierarchical models
- Basic TUI navigation with expand/collapse
- Content view mode (Enter to view, Escape to return)

### ⚠️ Partial (US3)
- Navigation structure implemented
- Content view mode works but shows raw markdown (Glamour integration pending)

### ⏳ Pending
- Glamour-based markdown rendering for content view
- Polish items (help text, edge cases, test fixtures)
