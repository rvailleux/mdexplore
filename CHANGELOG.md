# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2026-02-26

### Added

- Line numbers display in TOC format `L[start]-[end]` for each section
- `--level` / `-L` flag to limit maximum heading depth displayed
- Hierarchical navigation TUI with expand/collapse functionality
- Right arrow (→) to expand sections, Left arrow (←) to collapse
- Enter key to view section content in a separate view
- Escape key to return from content view to TOC
- Section content extraction and display
- `internal/renderer` package for markdown content rendering
- Comprehensive test coverage for Section and SectionTree models

### Changed

- Updated UI model to use SectionTree instead of flat TableOfContents
- Improved help text with navigation key documentation
- Enhanced visual styling with line number indicators

## [0.1.0] - 2026-02-25

### Added

- Initial implementation of mdexplore
- Interactive TUI using Bubble Tea
- Support for ATX-style headings (`# Heading`)
- Support for Setext-style headings (`Heading\n===`)
- Code block awareness (excludes headings in code blocks)
- YAML frontmatter support
- Error handling with styled error messages
- Comprehensive test suite
- CLI with --help, --version, and --toc flags

[Unreleased]: https://github.com/rvailleux/mdexplore/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/rvailleux/mdexplore/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/rvailleux/mdexplore/releases/tag/v0.1.0
