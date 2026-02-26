# Tasks: TUI Navigation Enhancement

**Input**: Design documents from `/specs/004-tui-navigation-enhancement/`
**Prerequisites**: plan.md, spec.md, data-model.md, quickstart.md

**Tests**: Tests included per TDD discipline (Constitution Principle I)

**Organization**: Tasks are grouped by user story to enable independent implementation and testing.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3, US4)
- Include exact file paths in descriptions

---

## Phase 1: Foundational (Model Changes)

**Purpose**: Core model changes that support all user stories

**⚠️ CRITICAL**: Must complete before user story implementation

- [ ] T001 [P] Add `ReturnIndex int` field to Model in `internal/ui/model.go`
- [ ] T002 Add `IsLeaf() bool` method to Section in `internal/models/section.go`
- [ ] T003 Add `GetFullContent() string` method to Section in `internal/models/section.go`

**Checkpoint**: Model changes ready - all helper methods available for user stories

---

## Phase 2: User Story 2 - Nested Numbering Next to Title (Priority: P1)

**Goal**: Move section numbers next to titles (e.g., "├── 1.1. Title")

**Independent Test**: Run `mdexplore README.md` and verify numbers appear after tree prefix

### Tests for User Story 2

> **NOTE: Write these tests FIRST, ensure they FAIL before implementation**

- [ ] T004 [P] [US2] Add test for `renderSectionWithNumber` number placement in `tests/unit/view_test.go`

### Implementation for User Story 2

- [ ] T005 [US2] Modify `renderSectionWithNumber` in `internal/ui/view.go` to place number after tree prefix
- [ ] T006 [US2] Update rendering order: indent → tree prefix → number → title

**Checkpoint**: Numbering displays correctly next to titles

---

## Phase 3: User Story 1 - Right Arrow on Leaf Opens Content (Priority: P1)

**Goal**: Right arrow opens content view on leaf sections (like Enter)

**Independent Test**: Navigate to leaf section, press right arrow, verify content opens

### Tests for User Story 1

- [ ] T007 [P] [US1] Add test for right arrow on leaf section in `tests/unit/update_test.go`
- [ ] T008 [P] [US1] Add test for right arrow on parent section (preserves expand) in `tests/unit/update_test.go`

### Implementation for User Story 1

- [ ] T009 [US1] Modify Right key handler in `internal/ui/update.go` to check `IsLeaf()`
- [ ] T010 [US1] On leaf: store `ReturnIndex`, set `ViewMode = ViewContent`, set `CurrentSection`
- [ ] T011 [US1] On parent: preserve existing expand behavior

**Checkpoint**: Right arrow opens content on leaf, expands on parent

---

## Phase 4: User Story 4 - Left Arrow Returns to Navigation (Priority: P1)

**Goal**: Left arrow closes content view and returns to TOC

**Independent Test**: Open content view, press left arrow, verify return to TOC with cursor restored

### Tests for User Story 4

- [ ] T012 [P] [US4] Add test for left arrow in content view in `tests/unit/update_test.go`
- [ ] T013 [P] [US4] Add test for cursor restoration when returning in `tests/unit/update_test.go`
- [ ] T014 [P] [US4] Add test for left arrow in TOC mode (preserves collapse) in `tests/unit/update_test.go`

### Implementation for User Story 4

- [ ] T015 [US4] Add Left key handler for `ViewContent` mode in `internal/ui/update.go`
- [ ] T016 [US4] On left in content view: restore `Selected = ReturnIndex`, set `ViewMode = ViewTOC`
- [ ] T017 [US4] Preserve existing collapse behavior for left arrow in TOC mode

**Checkpoint**: Left arrow returns from content view, cursor restored

---

## Phase 5: User Story 3 - Full Section Display Including Subsections (Priority: P1)

**Goal**: Enter shows full section content including all subsections

**Independent Test**: Press Enter on section with children, verify all descendant content displayed

### Tests for User Story 3

- [ ] T018 [P] [US3] Add test for `GetFullContent()` including descendants in `tests/unit/section_test.go`
- [ ] T019 [P] [US3] Add test for full content view in `tests/unit/update_test.go`

### Implementation for User Story 3

- [ ] T020 [US3] Implement `GetFullContent()` in `internal/models/section.go` to aggregate RawContent recursively
- [ ] T021 [US3] Modify Enter key handler in `internal/ui/update.go` to use `GetFullContent()`
- [ ] T022 [US3] Update `renderContent` in `internal/ui/view.go` to display aggregated content

**Checkpoint**: Enter displays full section with all subsections

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Final improvements and documentation

- [ ] T023 [P] Update README.md with new key bindings
- [ ] T024 Run `golangci-lint run` and fix any issues
- [ ] T025 Run full test suite: `go test ./...`
- [ ] T026 Manual test in terminal to verify navigation feels intuitive

---

## Dependencies & Execution Order

### Phase Dependencies

- **Foundational (Phase 1)**: No dependencies - can start immediately
- **User Stories (Phase 2-5)**: All depend on Foundational phase completion
  - US2 (Numbering) → US1 (Right arrow) → US4 (Left arrow) → US3 (Full display)
  - This order per implementation plan (simplest to most complex)
- **Polish (Phase 6)**: Depends on all user stories complete

### User Story Dependencies

All user stories can technically be done in parallel after Foundational, but recommended order:

1. **US2 (Numbering)**: Pure UI change, simplest
2. **US1 (Right arrow)**: Builds on existing Enter behavior
3. **US4 (Left arrow)**: Complements US1 bidirectional navigation
4. **US3 (Full display)**: Most complex, requires content aggregation

### Within Each User Story

- Tests MUST be written and FAIL before implementation
- Implementation follows test
- Verify story works independently before next

### Parallel Opportunities

- All Foundational tasks marked [P] can run in parallel
- All tests for a user story marked [P] can run in parallel
- US2 and US1+US4 could be worked in parallel (different concerns: UI vs navigation)

---

## Implementation Strategy

### Recommended Order (TDD)

1. **Complete Phase 1: Foundational**
   - Add model fields and helper methods
   - These support all user stories

2. **Complete Phase 2: US2 (Numbering placement)**
   - Simplest change, establishes pattern
   - Pure view layer modification

3. **Complete Phase 3: US1 (Right arrow)**
   - Builds on existing navigation
   - Uses IsLeaf() from Foundational

4. **Complete Phase 4: US4 (Left arrow)**
   - Complements US1
   - Uses ReturnIndex from Foundational

5. **Complete Phase 5: US3 (Full display)**
   - Most complex
   - Uses GetFullContent() from Foundational

6. **Complete Phase 6: Polish**
   - Documentation, lint, final tests

### Checkpoint Strategy

Stop and validate at each user story:
- After US2: Visual check of numbering placement
- After US1: Test right arrow on leaf and parent
- After US4: Test bidirectional navigation (right in, left out)
- After US3: Test full content display with nested sections

---

## Notes

- [P] tasks = different files, no dependencies
- All tasks follow TDD: test first, ensure failure, then implement
- Run `go test ./...` after each user story
- Run `golangci-lint run` before final commit
- Verify with manual terminal testing
