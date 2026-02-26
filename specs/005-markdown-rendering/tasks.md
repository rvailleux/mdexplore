# Tasks: Rich Markdown Content Rendering

**Input**: Design documents from `/specs/005-markdown-rendering/`
**Prerequisites**: plan.md, spec.md, data-model.md, research.md, quickstart.md

**Tests**: Tests included per TDD discipline

**Organization**: Tasks are grouped by user story to enable independent implementation and testing.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

---

## Phase 1: Setup (Glamour Dependency)

**Purpose**: Add Glamour library dependency

- [X] T001 Add `github.com/charmbracelet/glamour` to go.mod in project root
- [X] T002 [P] Run `go mod tidy` to download dependencies
- [X] T003 [P] Verify build succeeds: `go build ./cmd/mdexplore`

**Checkpoint**: Glamour library available for use

---

## Phase 2: Foundational (Renderer Package)

**Purpose**: Core renderer infrastructure that supports all user stories

**⚠️ CRITICAL**: Must complete before user story implementation

- [X] T004 Create `internal/renderer/` directory
- [X] T005 Define `Renderer` interface in `internal/renderer/renderer.go`
- [X] T006 Implement `GlamourRenderer` struct in `internal/renderer/glamour.go`
- [X] T007 Add `NewGlamourRenderer(style string)` constructor
- [X] T008 Implement `Render(markdown string, width int)` method
- [X] T009 Add tests for renderer in `tests/unit/renderer_test.go`

**Checkpoint**: Renderer package ready - can render markdown to styled text

---

## Phase 3: User Story 1 - Formatted Headings (Priority: P1) 🎯 MVP

**Goal**: Render H1-H6 headings with distinct colors and bold styling

**Independent Test**: View markdown with headings and verify visual hierarchy

### Tests for User Story 1

- [X] T010 [P] [US1] Add test for H1-H6 rendering in `tests/unit/renderer_test.go`

### Implementation for User Story 1

- [X] T011 [US1] Verify Glamour renders headings with default "dark" theme
- [X] T012 [US1] Test heading hierarchy in terminal output

**Checkpoint**: Headings display with visual hierarchy

---

## Phase 4: User Story 2 - Styled Lists (Priority: P1)

**Goal**: Render ordered/unordered lists with bullets, numbers, and proper indentation

**Independent Test**: View markdown with lists and verify bullets and numbering

### Tests for User Story 2

- [X] T013 [P] [US2] Add test for unordered list rendering in `tests/unit/renderer_test.go`
- [X] T014 [P] [US2] Add test for ordered list rendering in `tests/unit/renderer_test.go`

### Implementation for User Story 2

- [X] T015 [US2] Verify Glamour renders unordered lists with bullets
- [X] T016 [US2] Verify Glamour renders ordered lists with numbers
- [X] T017 [US2] Test nested list indentation
- [X] T018 [US2] Test task list (- [ ], - [x]) rendering

**Checkpoint**: All list types render correctly

---

## Phase 5: User Story 3 - Formatted Code Blocks (Priority: P1)

**Goal**: Render code blocks with background styling and syntax highlighting

**Independent Test**: View markdown with code blocks and verify formatting

### Tests for User Story 3

- [X] T019 [P] [US3] Add test for fenced code block rendering in `tests/unit/renderer_test.go`
- [X] T020 [P] [US3] Add test for inline code rendering in `tests/unit/renderer_test.go`

### Implementation for User Story 3

- [X] T021 [US3] Verify Glamour renders code blocks with background
- [X] T022 [US3] Test syntax highlighting for common languages (Go, Python, JavaScript)
- [X] T023 [US3] Verify inline code has distinct styling

**Checkpoint**: Code blocks have visual separation and syntax highlighting

---

## Phase 6: User Story 4 - Styled Text Formatting (Priority: P1)

**Goal**: Render bold, italic, and strikethrough text with appropriate styling

**Independent Test**: View markdown with **bold**, *italic*, ~~strikethrough~~ text

### Tests for User Story 4

- [X] T024 [P] [US4] Add test for bold text rendering in `tests/unit/renderer_test.go`
- [X] T025 [P] [US4] Add test for italic text rendering in `tests/unit/renderer_test.go`
- [X] T026 [P] [US4] Add test for strikethrough rendering in `tests/unit/renderer_test.go`

### Implementation for User Story 4

- [X] T027 [US4] Verify Glamour renders **bold** text
- [X] T028 [US4] Verify Glamour renders *italic* text
- [X] T029 [US4] Verify Glamour renders ~~strikethrough~~ text

**Checkpoint**: Text formatting styles render correctly

---

## Phase 7: User Story 5 - Links and References (Priority: P2)

**Goal**: Render hyperlinks in distinguishable color with underlining

**Independent Test**: View markdown with [links](url) and verify styling

### Tests for User Story 5

- [X] T030 [P] [US5] Add test for link rendering in `tests/unit/renderer_test.go`

### Implementation for User Story 5

- [X] T031 [US5] Verify Glamour renders links with underlining
- [X] T032 [US5] Verify bare URLs are styled as links

**Checkpoint**: Links are visually distinct

---

## Phase 8: User Story 6 - Blockquotes and Horizontal Rules (Priority: P2)

**Goal**: Render blockquotes and horizontal rules with visual distinction

**Independent Test**: View markdown with > blockquotes and --- rules

### Tests for User Story 6

- [X] T033 [P] [US6] Add test for blockquote rendering in `tests/unit/renderer_test.go`
- [X] T034 [P] [US6] Add test for horizontal rule rendering in `tests/unit/renderer_test.go`

### Implementation for User Story 6

- [X] T035 [US6] Verify Glamour renders blockquotes with left border
- [X] T036 [US6] Verify Glamour renders horizontal rules as line separators

**Checkpoint**: Blockquotes and rules visually distinct

---

## Phase 9: User Story 7 - Tables (Priority: P3)

**Goal**: Render tables with proper column alignment and borders

**Independent Test**: View markdown with tables and verify alignment

### Tests for User Story 7

- [X] T037 [P] [US7] Add test for table rendering in `tests/unit/renderer_test.go`

### Implementation for User Story 7

- [X] T038 [US7] Verify Glamour renders tables with aligned columns
- [X] T039 [US7] Verify table headers are visually distinguished

**Checkpoint**: Tables render with proper formatting

---

## Phase 10: Integration (UI Layer)

**Purpose**: Integrate renderer with TUI content view

- [X] T040 Add `renderer` field to Model in `internal/ui/model.go`
- [X] T041 Initialize renderer in `InitialModel()` function
- [X] T042 Modify `renderContent()` in `internal/ui/view.go` to use renderer
- [X] T043 Pass terminal width to renderer for proper wrapping
- [X] T044 Handle renderer errors gracefully (fallback to plain text)

**Checkpoint**: TUI displays rendered markdown in content view

---

## Phase 11: Polish & Cross-Cutting Concerns

**Purpose**: Final improvements and documentation

- [X] T045 [P] Update README.md with markdown rendering features
- [X] T046 Run `go test ./...` and verify all tests pass
- [X] T047 Run `golangci-lint run` and fix any issues
- [X] T048 Manual test in terminal with various markdown files
- [X] T049 Verify performance under 100ms for typical sections

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - start immediately
- **Foundational (Phase 2)**: Depends on Setup - BLOCKS all user stories
- **User Stories (Phases 3-9)**: Can proceed in parallel after Foundational
  - US1, US2, US3, US4 (P1) can be parallel
  - US5, US6 (P2) can be parallel after P1
  - US7 (P3) can be after P2
- **Integration (Phase 10)**: Depends on Foundational and at least US1
- **Polish (Phase 11)**: Depends on all user stories and integration

### Within Each User Story

- Tests MUST be written and FAIL before implementation
- Implementation follows test
- Verify story works independently

### Parallel Opportunities

- All P1 user stories (US1-US4) can run in parallel after Foundational
- All tests for a user story marked [P] can run in parallel

---

## Implementation Strategy

### Recommended Order

1. **Setup**: Add Glamour dependency
2. **Foundational**: Create renderer package with tests
3. **US1-US4 (P1)**: Core rendering features - can be parallel
4. **Integration**: Wire renderer into TUI
5. **US5-US7 (P2-P3)**: Additional features
6. **Polish**: Documentation, final tests

### Checkpoint Strategy

Stop and validate at each phase:
- After Foundational: Renderer works in isolation
- After US1-US4: Core markdown renders correctly
- After Integration: TUI shows rendered content
- After Polish: Feature complete and documented

---

## Notes

- [P] tasks = different files, no dependencies
- Glamour handles most rendering - tasks are verification/integration
- Run `go test ./...` after each user story
- Manual terminal testing essential for visual verification
