<!--
================================================================================
SYNC IMPACT REPORT - Constitution v1.0.1
================================================================================
Version Change: 1.0.0 → 1.0.1
Modified Principles:
  - Principle IV: Agent-First Design (UPDATED - clarified mdexplore is for coding agents)
Added Sections: None
Removed Sections: None
Templates Requiring Updates:
  - ✅ .specify/templates/plan-template.md - Constitution Check section reviewed
  - ✅ .specify/templates/spec-template.md - User scenarios align with principles
  - ✅ .specify/templates/tasks-template.md - TDD testing discipline reflected
Follow-up TODOs: None
================================================================================
-->

# AgentMark Constitution

## Core Principles

### I. Test-Driven Development (NON-NEGOTIABLE)

All code MUST be written using strict TDD discipline:

- **Red Phase**: Write a failing test that describes the desired behavior BEFORE any implementation code
- **Green Phase**: Write the minimum code necessary to make the test pass
- **Refactor Phase**: Clean up the implementation while keeping all tests green
- **No Exceptions**: No production code exists without a failing test preceding it
- **Test Coverage**: All user stories must have corresponding acceptance tests

**Rationale**: TDD ensures testable, well-designed code and prevents technical debt. For a tool used by coding agents, reliability is paramount—tests are the contract that guarantees behavior.

### II. Terminal User Interface

The application MUST provide a rich, interactive terminal user interface:

- **TUI Library**: Use a dedicated TUI library (e.g., Bubble Tea for Go, Rich/Textual for Python, Ink for Node.js) rather than plain console output
- **Interactive Experience**: Support keyboard navigation, scrolling, and interactive elements
- **Visual Polish**: Use colors, styling, and layout to create a pleasant reading experience
- **Responsive**: Handle terminal resizing gracefully
- **Accessibility**: Support common terminal capabilities and degradable features

**Rationale**: Coding agents spend significant time in terminal environments. A polished TUI makes the markdown reading experience efficient and enjoyable, distinguishing this tool from basic `cat` or `less` alternatives.

### III. Markdown Reader Focus

The project exists solely to render and navigate Markdown documents for coding agents:

- **Primary Purpose**: Parse and display Markdown with proper formatting
- **Code Highlighting**: Syntax highlighting for code blocks is essential
- **Navigation**: Support for jumping between sections, following links, and searching
- **Reader-First**: Optimize for reading, not editing (view mode over edit mode)
- **Agent Context**: Consider the workflows of AI coding agents (reading specs, documentation, READMEs)

**Rationale**: Scope clarity prevents feature creep. By focusing exclusively on reading (not editing) markdown, we deliver a superior experience for our target users—coding agents reviewing documentation and specifications.

### IV. Agent-First Design (mdexplore is for Coding Agents)

**mdexplore is primarily designed for coding agents.** All design decisions prioritize the coding agent workflow:

- **Spec Reading**: Optimize for reading technical specifications and implementation plans
- **Documentation Navigation**: Easy traversal of project documentation structures
- **Quick Context**: Fast startup and navigation to minimize context-switching overhead
- **Integration-Friendly**: Work well when invoked by scripts and automation
- **Clear Output**: Structured output suitable for parsing when needed

**Rationale**: Coding agents (both AI and human) have specific needs: quick access to information, clear formatting of technical content, and minimal friction. Design choices must serve this audience over general-purpose markdown viewers. The tool is built for agents who need to rapidly consume technical documentation.

### V. Simplicity

Keep the codebase and user interface simple:

- **YAGNI**: Do not add features until they are demonstrably needed
- **Minimal Dependencies**: Prefer standard library solutions; vet external dependencies carefully
- **Clear Abstractions**: Each module has a single, well-defined responsibility
- **Readable Code**: Prioritize clarity over cleverness
- **Incremental Growth**: Start small, expand based on usage patterns

**Rationale**: Simple systems are easier to maintain, test, and extend. For a tool that agents rely on, predictability and reliability trump feature richness.

## Technology Standards

### Language & Runtime

- **Primary Language**: Go 1.23+ (Bubble Tea TUI ecosystem)
- **Target Platforms**: Linux, macOS, Windows (terminal environments)
- **Minimum Requirements**: Modern terminal with UTF-8 and color support

### Dependencies

- **TUI Library**: Required—Bubble Tea for Go
- **Markdown Parser**: Required—Goldmark for CommonMark support
- **Syntax Highlighter**: Required—for code block rendering
- **Testing Framework**: Required—Go's standard testing package

### Project Structure

```
├── cmd/              # CLI entry points
├── internal/         # Internal packages
│   ├── ui/           # TUI components
│   ├── parser/       # Markdown parsing
│   ├── models/       # Data models
│   └── errors/       # Error types
├── tests/            # Test files
│   ├── unit/         # Unit tests (TDD: written first)
│   ├── integration/  # Integration tests
│   └── benchmark/    # Performance tests
└── docs/             # Documentation
```

## Development Workflow

### TDD Cycle Enforcement

1. Write failing test for new behavior
2. Run tests to confirm failure (Red)
3. Implement minimal code to pass (Green)
4. Refactor while tests pass
5. Commit with descriptive message
6. Repeat

### Code Quality Gates

- All tests MUST pass before merging
- New features MUST include tests (per TDD)
- TUI changes SHOULD be manually tested in common terminals

### Commit Convention

```
<type>: <description>

type:
  - feat: new feature
  - fix: bug fix
  - test: adding or updating tests
  - refactor: code change that neither fixes nor adds feature
  - docs: documentation only
  - chore: maintenance tasks
```

## Governance

### Authority

This constitution supersedes all other development practices. Any conflict between this document and other guidance MUST be resolved in favor of this constitution.

### Amendments

1. Proposed changes MUST be documented with rationale
2. Changes affecting principles require review of dependent templates
3. Version MUST be incremented per semantic versioning:
   - **MAJOR**: Breaking governance changes, principle removals/redefinitions
   - **MINOR**: New principles or expanded guidance
   - **PATCH**: Clarifications, wording improvements, typo fixes
4. Sync Impact Report MUST be updated with each amendment

### Compliance Review

- All implementation plans MUST pass Constitution Check before proceeding
- PR reviews MUST verify TDD compliance (tests exist for all new code)
- Regular reviews ensure templates remain aligned with principles

---

**Version**: 1.0.1 | **Ratified**: 2025-02-25 | **Last Amended**: 2026-02-25
