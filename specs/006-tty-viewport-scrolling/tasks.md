# Tasks: TTY Viewport Scrolling

**Feature**: TTY Viewport Scrolling
**Branch**: `006-tty-viewport-scrolling`
**Generated**: 2026-02-26

---

## Task Overview

| Phase | Description | Task Count |
|-------|-------------|------------|
| Phase 1 | Foundational Setup | 3 |
| Phase 2 | User Story 1: Arrow Key Scrolling | 6 |
| Phase 3 | User Story 2: Scroll Position Indicator | 3 |
| Phase 4 | User Story 3: Page-Up/Page-Down Navigation | 3 |
| Phase 5 | Polish & Cross-Cutting Concerns | 3 |
| **Total** | | **18** |

---

## Dependencies

```
Phase 1 (Foundational)
    │
    ├──► Phase 2 (US1: Arrow Key Scrolling) ──┐
    │                                           │
    ├──► Phase 3 (US2: Position Indicator) ────┼──► Phase 5 (Polish)
    │                                           │
    └──► Phase 4 (US3: Page Navigation) ───────┘
```

**Story Dependencies**: Each user story builds on the foundational viewport state from Phase 1 but can be developed in parallel once Phase 1 is complete.

**Independent Test Criteria**:
- **US1**: Can scroll long content with arrow keys, boundaries respected
- **US2**: Percentage indicator updates correctly as you scroll
- **US3**: Page keys jump by screen height

---

## Phase 1: Foundational Setup

**Goal**: Add viewport state fields to the model and establish scroll calculation helpers
**Blocking**: Required for all user stories

- [x] T001 Add viewport state fields to ui.Model in `internal/ui/model.go` (ContentScrollOffset, ContentTotalLines, ViewportHeight)
- [x] T002 [P] Create viewport helper functions in `internal/ui/viewport.go` (CalculateMaxOffset, ClampOffset)
- [x] T003 Initialize viewport state in InitialModel functions in `internal/ui/model.go`

---

## Phase 2: User Story 1 - Scroll Long Content with Arrow Keys

**Goal**: Users can scroll content exceeding terminal height using Up/Down arrows
**Priority**: P1
**Independent Test**: Can fully test by opening a markdown file with content longer than terminal height and verifying arrow keys allow viewing all content from top to bottom

- [x] T004 [P] [US1] Write unit tests for viewport calculations in `tests/unit/viewport_test.go` (scroll offset, boundary checks)
- [x] T005 [US1] Add arrow key handlers (up/down/j/k) in `internal/ui/update.go` for ViewContent mode
- [x] T006 [P] [US1] Modify renderContent() in `internal/ui/view.go` to display viewport window based on ContentScrollOffset
- [x] T007 [US1] Calculate ViewportHeight accounting for headers/footers in `internal/ui/view.go`
- [x] T008 [P] [US1] Write integration tests for arrow key scrolling in `tests/integration/scrolling_test.go`
- [x] T009 [US1] Reset scroll offset when entering content view in `internal/ui/update.go` handleKeyPress enter/right

**Acceptance Criteria**:
- Given terminal height 24 lines and content 50 lines, pressing Down reveals line 25
- Given scrolled down 10 lines, pressing Up reveals previous content
- Given content shorter than terminal, no scrolling occurs
- Given at top, pressing Up stays at top
- Given at bottom, pressing Down stays at bottom

---

## Phase 3: User Story 2 - Visual Scroll Position Indicator

**Goal**: Users see a visual indicator of their current scroll position
**Priority**: P2
**Independent Test**: Can test by scrolling through a long document and verifying position indicator updates correctly

- [x] T010 [P] [US2] Create scroll percentage calculation function in `internal/ui/viewport.go`
- [x] T011 [US2] Add scroll position indicator to help footer in `internal/ui/view.go` renderContent()
- [x] T012 [P] [US2] Write unit tests for percentage calculation in `tests/unit/viewport_test.go`

**Acceptance Criteria**:
- At top: indicator shows 0% or beginning state
- At middle: indicator shows ~50%
- At end: indicator shows 100%
- Hidden when all content fits within viewport

---

## Phase 4: User Story 3 - Page-Up and Page-Down Navigation

**Goal**: Users can navigate quickly using Page Up/Page Down keys
**Priority**: P3
**Independent Test**: Can test by pressing Page Up/Page Down in a long document and verifying viewport jumps by approximately one screen height

- [x] T013 [P] [US3] Add Page Up key handler in `internal/ui/update.go` handleKeyPress()
- [x] T014 [P] [US3] Add Page Down key handler in `internal/ui/update.go` handleKeyPress()
- [x] T015 [US3] Add Home/End key handlers (g/G) for jump to top/bottom in `internal/ui/update.go`

**Acceptance Criteria**:
- Page Down from lines 1-24 shows lines 25-48 (or to end)
- Page Up from lines 50-73 shows lines 26-49
- Less than one page remaining jumps to final lines

---

## Phase 5: Polish & Cross-Cutting Concerns

**Goal**: Handle edge cases, terminal resize, and performance
**Priority**: Polish

- [x] T016 Handle terminal resize events in `internal/ui/update.go` WindowSizeMsg to adjust ViewportHeight and clamp scroll offset
- [x] T017 [P] Write benchmark tests for scroll performance in `tests/benchmark/scroll_bench_test.go` (<50ms response time)
- [x] T018 Handle fast repeated key presses - ensure responsive scrolling without visual glitches in `internal/ui/update.go`

---

## Implementation Strategy

### MVP Scope (Phase 1 + Phase 2)

The Minimum Viable Product includes:
1. Viewport state fields in model
2. Arrow key scrolling (up/down)
3. Boundary enforcement (can't scroll past top/bottom)
4. Viewport window rendering

This delivers the core value: users can read all content in long documents.

### Incremental Delivery

1. **Phase 1** → Phase 2 can begin immediately after
2. **Phase 2** → Parallel work on Phase 3 and Phase 4 can start
3. **Phase 3 & 4** → Can be developed independently after Phase 1
4. **Phase 5** → Final polish after all functional stories complete

### Parallel Execution Opportunities

Tasks marked with **[P]** can be executed in parallel:
- T002, T004, T008: Viewport helper, unit tests, integration tests
- T006: Window rendering can be done concurrently with key handlers
- T010, T012: Position indicator calculations and tests
- T013, T014, T017: Page navigation and benchmarks

---

## File Summary

| File Path | Tasks | Description |
|-----------|-------|-------------|
| `internal/ui/model.go` | T001, T003 | Add viewport state fields |
| `internal/ui/viewport.go` | T002, T010 | New: Viewport helper functions |
| `internal/ui/update.go` | T005, T009, T013-T016, T018 | Key handlers and resize logic |
| `internal/ui/view.go` | T007, T011 | Viewport height calc and indicator display |
| `tests/unit/viewport_test.go` | T004, T012 | Unit tests for viewport logic |
| `tests/integration/scrolling_test.go` | T008 | Integration tests for scrolling |
| `tests/benchmark/scroll_bench_test.go` | T017 | Performance benchmarks |
