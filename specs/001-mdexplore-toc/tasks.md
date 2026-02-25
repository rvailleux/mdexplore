# Tasks: Markdown TOC Explorer

**Input**: Design documents from `/specs/001-mdexplore-toc/`
**Prerequisites**: plan.md, spec.md, data-model.md, contracts/

**Tests**: Tests are included per TDD requirements from constitution

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- Go project: `cmd/`, `internal/`, `tests/` at repository root

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [X] T001 Create Go module and project structure: `go.mod`, `cmd/mdexplore/`, `internal/parser/`, `internal/ui/`, `internal/toc/`, `internal/errors/`, `tests/unit/`, `tests/integration/`, `tests/fixtures/`
- [X] T002 [P] Add dependencies to go.mod: `bubbletea`, `lipgloss`, `bubbles`, `goldmark`, `cobra`, `testify`
- [X] T003 [P] Configure linting with `golangci-lint` and formatting with `gofmt`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

### Data Models

- [X] T004 [P] Create Heading struct in `internal/parser/heading.go` with Level (int), Text (string), LineNumber (int)
- [X] T005 [P] Create TableOfContents struct in `internal/toc/toc.go` with Headings ([]Heading), Source (string)
- [X] T006 Create custom error types in `internal/errors/errors.go`: FileNotFoundError, PermissionDeniedError, InvalidFileError, ParseError

### Base Parser Infrastructure

- [X] T007 Create parser interface in `internal/parser/parser.go`: Parse(filepath string) (TableOfContents, error)
- [X] T008 Implement file validation in `internal/parser/parser.go`: check exists, is file (not dir), readable

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Display Table of Contents (Priority: P1) 🎯 MVP

**Goal**: Parse markdown file and display hierarchical TOC in interactive TUI

**Independent Test**: Run `mdexplore tests/fixtures/sample_with_headings.md --toc` and verify all headings appear in correct order with proper indentation

### Tests for User Story 1 ⚠️

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [X] T010 [P] [US1] Create test fixture `tests/fixtures/sample_with_headings.md` with H1, H2, H3 headings
- [X] T011 [P] [US1] Create test fixture `tests/fixtures/empty.md` with no headings
- [X] T012 [US1] Write unit tests in `internal/parser/parser_test.go` for ATX heading extraction (Red phase)
- [X] T013 [US1] Write unit tests in `internal/toc/toc_test.go` for TOC building from headings (Red phase)
- [X] T014 [US1] Write integration test in `tests/integration/toc_display_test.go` for end-to-end TOC display (Red phase)

### Implementation for User Story 1

- [X] T015 [US1] Implement ATX heading extraction in `internal/parser/parser.go` using Goldmark AST visitor (Green phase)
- [X] T016 [US1] Implement TOC building logic in `internal/toc/toc.go` with hierarchical tree structure (Green phase)
- [X] T017 [US1] Create Bubble Tea model in `internal/ui/model.go` with state: filename, toc, selected index, error
- [X] T018 [US1] Create view renderer in `internal/ui/view.go` using Lipgloss for styled TOC display with tree characters
- [X] T019 [US1] Create update handler in `internal/ui/update.go` for keyboard navigation (up/down, quit)
- [X] T020 [US1] Wire CLI entry point in `cmd/mdexplore/main.go` with `--toc` flag and TUI launch
- [X] T021 [US1] Handle "no headings" case: display friendly message when TOC is empty (Refactor phase)
- [X] T022 [US1] Refactor parser tests to pass with implementation (Green phase)
- [X] T023 [US1] Refactor TOC tests to pass with implementation (Green phase)

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently. Run `go test ./...` and verify all US1 tests pass.

---

## Phase 4: User Story 2 - Handle File Errors Gracefully (Priority: P2)

**Goal**: Display clear, styled error messages for file-related errors in the TUI

**Independent Test**: Run `mdexplore` with non-existent file, directory, and permission-denied scenarios - verify each shows appropriate styled error message

### Tests for User Story 2 ⚠️

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [X] T030 [P] [US2] Write unit tests in `internal/parser/parser_test.go` for FileNotFoundError handling (Red phase)
- [X] T031 [P] [US2] Write unit tests in `internal/parser/parser_test.go` for PermissionDeniedError handling (Red phase)
- [X] T032 [P] [US2] Write unit tests in `internal/parser/parser_test.go` for InvalidFileError (directory) handling (Red phase)
- [X] T033 [US2] Write integration tests in `tests/integration/error_handling_test.go` for error display in TUI (Red phase)

### Implementation for User Story 2

- [X] T034 [US2] Enhance file validation in `internal/parser/parser.go` to return specific error types (Green phase)
- [X] T035 [US2] Update Bubble Tea model in `internal/ui/model.go` to handle error states
- [X] T036 [US2] Create error view in `internal/ui/view.go` with styled error display using Lipgloss (emoji + message)
- [X] T037 [US2] Update main.go to set appropriate exit codes (2=file not found, 3=permission denied, 4=invalid file) per CLI contract
- [X] T038 [US2] Refactor parser tests to pass with error handling implementation (Green phase)

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently. Test: `mdexplore missing.md --toc` shows error, `mdexplore valid.md --toc` shows TOC.

---

## Phase 5: User Story 3 - Support Various Markdown Heading Formats (Priority: P3)

**Goal**: Recognize both ATX-style and Setext-style headings

**Independent Test**: Run `mdexplore` on files with ATX headings, Setext headings, and mixed formats - verify all headings extracted with correct levels

### Tests for User Story 3 ⚠️

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [X] T040 [P] [US3] Create test fixture `tests/fixtures/setext_headings.md` with underlined H1 (===) and H2 (---)
- [X] T041 [P] [US3] Create test fixture `tests/fixtures/mixed_headings.md` with both ATX and Setext styles
- [X] T042 [US3] Write unit tests in `internal/parser/parser_test.go` for Setext H1 detection (===) (Red phase)
- [X] T043 [US3] Write unit tests in `internal/parser/parser_test.go` for Setext H2 detection (---) (Red phase)
- [X] T044 [US3] Write integration test in `tests/integration/mixed_formats_test.go` for mixed heading formats (Red phase)

### Implementation for User Story 3

- [X] T045 [US3] Extend Goldmark parser configuration in `internal/parser/parser.go` to enable Setext heading extension (Green phase)
- [X] T046 [US3] Verify AST visitor correctly extracts Setext headings with proper level assignment (Green phase)
- [X] T047 [US3] Refactor parser tests to pass with Setext support (Green phase)
- [X] T048 [US3] Test edge case: headings inside code blocks are excluded (may require AST context tracking)

**Checkpoint**: All user stories should now be independently functional. Test with: ATX-only file, Setext-only file, mixed file, and verify correct TOC in all cases.

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [X] T050 [P] Add code block exclusion logic in `internal/parser/parser.go` to filter headings inside fenced/indented code blocks
- [X] T051 [P] Add frontmatter (YAML) handling to skip `---` delimited headers at file start
- [X] T052 Add performance benchmark in `tests/benchmark/parser_bench_test.go` to verify <1s for 1MB files
- [X] T053 Add help text (`--help`) in `cmd/mdexplore/main.go` showing usage and examples per CLI contract
- [X] T054 Add version flag (`--version`) in `cmd/mdexplore/main.go`
- [X] T055 Update `README.md` with installation and usage instructions
- [X] T056 Run full test suite: `go test -v ./...` and verify 100% pass rate
- [X] T057 Validate quickstart.md examples work correctly

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3-5)**: All depend on Foundational phase completion
  - User stories can then proceed in parallel (if staffed)
  - Or sequentially in priority order (P1 → P2 → P3)
- **Polish (Phase 6)**: Depends on all desired user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2) - No dependencies on other stories
- **User Story 2 (P2)**: Can start after Foundational (Phase 2) - Extends US1 with error handling
- **User Story 3 (P3)**: Can start after Foundational (Phase 2) - Extends parser from US1

### Within Each User Story

- Tests MUST be written and FAIL before implementation (TDD)
- Models before services
- Core parsing/UI before integration
- Story complete before moving to next priority

### Parallel Opportunities

- All Setup tasks marked [P] can run in parallel
- All Foundational tasks marked [P] can run in parallel (within Phase 2)
- All test fixtures can be created in parallel
- Once Foundational phase completes, all user stories can start in parallel (if team capacity allows)
- All tests for a user story marked [P] can run in parallel
- Parser tests and TOC tests can be written in parallel

---

## Parallel Example: User Story 1

```bash
# Launch all test fixtures together:
Task: "Create test fixture tests/fixtures/sample_with_headings.md"
Task: "Create test fixture tests/fixtures/empty.md"

# Launch all tests for User Story 1 together (write first - Red phase):
Task: "Write unit tests in internal/parser/parser_test.go"
Task: "Write unit tests in internal/toc/toc_test.go"
Task: "Write integration test in tests/integration/toc_display_test.go"

# Launch models in parallel (if not already done in Foundational):
Task: "Create Heading struct in internal/parser/heading.go"
Task: "Create TableOfContents struct in internal/toc/toc.go"

# Then implement to make tests pass (Green phase):
Task: "Implement ATX heading extraction in internal/parser/parser.go"
Task: "Implement TOC building logic in internal/toc/toc.go"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL - blocks all stories)
3. Complete Phase 3: User Story 1 (includes TDD cycle: Red-Green-Refactor)
4. **STOP and VALIDATE**: Test User Story 1 independently
   - Run: `go test ./...`
   - Manual test: `go run cmd/mdexplore/main.go tests/fixtures/sample_with_headings.md --toc`
5. Deploy/demo if ready

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 → Test independently → Deploy/Demo (MVP!)
3. Add User Story 2 → Test independently → Deploy/Demo
4. Add User Story 3 → Test independently → Deploy/Demo
5. Complete Phase 6: Polish → Final release
6. Each story adds value without breaking previous stories

### Parallel Team Strategy

With multiple developers:

1. Team completes Setup + Foundational together
2. Once Foundational is done:
   - Developer A: User Story 1 (core TOC functionality)
   - Developer B: User Story 2 (error handling)
   - Developer C: User Story 3 (Setext heading support)
3. Stories complete and integrate independently
4. Team converges on Phase 6: Polish

---

## Task Summary

| Phase | Tasks | Description |
|-------|-------|-------------|
| Phase 1: Setup | 3 | Project structure, dependencies, linting |
| Phase 2: Foundational | 5 | Data models, error types, parser interface |
| Phase 3: US1 (P1) | 14 | TOC display with TDD (tests + implementation) |
| Phase 4: US2 (P2) | 9 | Error handling with TDD |
| Phase 5: US3 (P3) | 9 | Setext headings with TDD |
| Phase 6: Polish | 8 | Edge cases, docs, validation |
| **Total** | **48** | |

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- Each user story should be independently completable and testable
- **CRITICAL**: Verify tests fail (Red) before implementing (Green)
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
- Avoid: vague tasks, same file conflicts, cross-story dependencies that break independence
- All tasks follow TDD per constitution Principle I
