package models

import (
	"fmt"
	"strconv"
	"strings"
)

// Section represents a hierarchical document section with heading, line range, and content.
type Section struct {
	ID         string     // Unique identifier (format: "L{StartLine}")
	Level      int        // 1-6 (H1-H6)
	Title      string     // Clean heading text (no markdown markers)
	StartLine  int        // Line number where heading starts (1-based)
	EndLine    int        // Line number where section ends (inclusive)
	RawContent string     // Raw markdown content between StartLine and EndLine
	Children   []*Section // Nested subsections (immediate children only)
	Parent     *Section   // Parent section (nil for H1 sections)
}

// HasChildren returns true if section has nested subsections.
func (s *Section) HasChildren() bool {
	return len(s.Children) > 0
}

// GetAllDescendants returns flattened list of all nested sections recursively.
func (s *Section) GetAllDescendants() []*Section {
	var result []*Section
	for _, child := range s.Children {
		result = append(result, child)
		result = append(result, child.GetAllDescendants()...)
	}
	return result
}

// GetVisibleDescendants returns immediate children (for expansion display).
func (s *Section) GetVisibleDescendants() []*Section {
	return s.Children
}

// ContainsLine returns true if the given line number is within this section.
func (s *Section) ContainsLine(line int) bool {
	return line >= s.StartLine && line <= s.EndLine
}

// GetDepth returns the nesting depth (0 for H1, 1 for H2, etc.).
func (s *Section) GetDepth() int {
	return s.Level - 1
}

// SectionTree contains the complete hierarchical structure of a document.
type SectionTree struct {
	Root     *Section          // Virtual root node containing all H1 sections as children
	Source   string            // File path
	Sections []*Section        // Flattened list of all sections for indexed access
	ByID     map[string]*Section // Lookup map from ID to section
}

// GetH1Sections returns all top-level sections.
func (t *SectionTree) GetH1Sections() []*Section {
	if t.Root == nil {
		return nil
	}
	return t.Root.Children
}

// FindByID looks up a section by its unique ID.
func (t *SectionTree) FindByID(id string) (*Section, bool) {
	section, ok := t.ByID[id]
	return section, ok
}

// GetFlattenedVisible returns sections in display order respecting expansion state.
// If expanded is nil or empty, returns only H1 sections.
func (t *SectionTree) GetFlattenedVisible(expanded map[string]bool) []*Section {
	var result []*Section
	h1Sections := t.GetH1Sections()

	for _, h1 := range h1Sections {
		result = append(result, h1)
		if expanded != nil && expanded[h1.ID] {
			result = t.appendVisibleChildren(result, h1, expanded)
		}
	}

	return result
}

// appendVisibleChildren recursively appends visible children to result.
func (t *SectionTree) appendVisibleChildren(result []*Section, parent *Section, expanded map[string]bool) []*Section {
	for _, child := range parent.Children {
		result = append(result, child)
		if expanded[child.ID] {
			result = t.appendVisibleChildren(result, child, expanded)
		}
	}
	return result
}

// FilterByMaxLevel returns a new tree with only sections up to given level.
// If maxLevel is 0, returns the original tree (no filtering).
func (t *SectionTree) FilterByMaxLevel(maxLevel int) *SectionTree {
	if maxLevel == 0 {
		return t
	}

	filtered := &SectionTree{
		Source:   t.Source,
		Sections: make([]*Section, 0),
		ByID:     make(map[string]*Section),
		Root: &Section{
			ID:       "root",
			Level:    0,
			Title:    "Root",
			Children: []*Section{},
		},
	}

	// Filter and add sections
	for _, section := range t.Sections {
		if section.Level <= maxLevel {
			filtered.Sections = append(filtered.Sections, section)
			filtered.ByID[section.ID] = section
		}
	}

	// Rebuild tree structure
	for _, section := range filtered.Sections {
		if section.Level == 1 {
			filtered.Root.Children = append(filtered.Root.Children, section)
		} else {
			// Find parent in filtered set
			if section.Parent != nil {
				if _, ok := filtered.ByID[section.Parent.ID]; ok {
					section.Parent.Children = append(section.Parent.Children, section)
				}
			}
		}
	}

	return filtered
}

// NavigationState tracks the current TUI navigation context and view state.
type NavigationState struct {
	SelectedIndex    int            // Current cursor position in visible list
	ExpandedSections map[string]bool // Set of section IDs currently expanded
	ViewMode         ViewMode       // Current display mode (TOC or Content)
	CurrentSection   *Section       // Section being viewed (when in Content mode)
	LevelFilter      int            // 0 = no filter, >0 = max heading level to show
}

// ViewMode represents the current display mode.
type ViewMode int

const (
	ViewTOC ViewMode = iota      // Table of contents navigation view
	ViewContent                  // Section content display view
)

// IsExpanded returns true if given section is expanded.
func (n *NavigationState) IsExpanded(sectionID string) bool {
	if n.ExpandedSections == nil {
		return false
	}
	return n.ExpandedSections[sectionID]
}

// ToggleExpanded flips the expansion state of a section.
func (n *NavigationState) ToggleExpanded(sectionID string) {
	if n.ExpandedSections == nil {
		n.ExpandedSections = make(map[string]bool)
	}
	n.ExpandedSections[sectionID] = !n.ExpandedSections[sectionID]
}

// Expand marks a section as expanded.
func (n *NavigationState) Expand(sectionID string) {
	if n.ExpandedSections == nil {
		n.ExpandedSections = make(map[string]bool)
	}
	n.ExpandedSections[sectionID] = true
}

// Collapse marks a section as collapsed.
func (n *NavigationState) Collapse(sectionID string) {
	if n.ExpandedSections == nil {
		n.ExpandedSections = make(map[string]bool)
	}
	n.ExpandedSections[sectionID] = false
}

// CanNavigateUp returns true if selection can move up.
func (n *NavigationState) CanNavigateUp(visibleCount int) bool {
	return n.SelectedIndex > 0
}

// CanNavigateDown returns true if selection can move down.
func (n *NavigationState) CanNavigateDown(visibleCount int) bool {
	return n.SelectedIndex < visibleCount-1
}

// DisplaySection is a view-model for rendering a section in the TUI.
type DisplaySection struct {
	Section      *Section // Reference to underlying section
	DisplayIndex int      // Position in the currently visible list
	Depth        int      // Visual indentation level (0 for H1, etc.)
	IsSelected   bool     // Whether this item has keyboard focus
	IsExpanded   bool     // Whether this section's children are visible
	CanExpand    bool     // Whether this section has children to expand
}

// SectionNumber represents a hierarchical section number (e.g., "1.1.2").
type SectionNumber struct {
	Parts []int // The number components (e.g., [1, 1] for "1.1")
}

// String returns the formatted section number (e.g., "1.1.").
func (sn SectionNumber) String() string {
	if len(sn.Parts) == 0 {
		return ""
	}
	result := ""
	for _, part := range sn.Parts {
		result += fmt.Sprintf("%d.", part)
	}
	return result
}

// Parent returns the parent section number.
func (sn SectionNumber) Parent() SectionNumber {
	if len(sn.Parts) <= 1 {
		return SectionNumber{}
	}
	return SectionNumber{Parts: sn.Parts[:len(sn.Parts)-1]}
}

// IsAncestorOf returns true if this number is an ancestor of the given number.
func (sn SectionNumber) IsAncestorOf(other SectionNumber) bool {
	if len(sn.Parts) >= len(other.Parts) {
		return false
	}
	for i := range sn.Parts {
		if sn.Parts[i] != other.Parts[i] {
			return false
		}
	}
	return true
}

// Equals returns true if the section numbers are equal.
func (sn SectionNumber) Equals(other SectionNumber) bool {
	if len(sn.Parts) != len(other.Parts) {
		return false
	}
	for i := range sn.Parts {
		if sn.Parts[i] != other.Parts[i] {
			return false
		}
	}
	return true
}

// Depth returns the depth level (number of parts).
func (sn SectionNumber) Depth() int {
	return len(sn.Parts)
}

// ParseSectionNumber parses a section number string (e.g., "1.1").
func ParseSectionNumber(s string) (SectionNumber, error) {
	var parts []int
	s = strings.TrimSuffix(s, ".")
	if s == "" {
		return SectionNumber{}, fmt.Errorf("empty section number")
	}
	for _, part := range strings.Split(s, ".") {
		num, err := strconv.Atoi(part)
		if err != nil || num < 1 {
			return SectionNumber{}, fmt.Errorf("invalid section number part: %s", part)
		}
		parts = append(parts, num)
	}
	return SectionNumber{Parts: parts}, nil
}

// NumberedSection wraps a Section with its hierarchical number.
type NumberedSection struct {
	Section       *Section
	Number        SectionNumber
	DisplayNumber string
}

// AssignNumbers assigns hierarchical numbers to all sections in the tree.
func (t *SectionTree) AssignNumbers() []*NumberedSection {
	var result []*NumberedSection
	counters := make(map[int]int) // Track counters at each depth

	for _, section := range t.Sections {
		depth := section.GetDepth()
		counters[depth]++
		// Reset deeper counters
		for d := depth + 1; d <= 6; d++ {
			counters[d] = 0
		}

		// Build the number parts
		parts := make([]int, depth+1)
		for d := 0; d <= depth; d++ {
			parts[d] = counters[d]
		}

		num := SectionNumber{Parts: parts}
		result = append(result, &NumberedSection{
			Section:       section,
			Number:        num,
			DisplayNumber: num.String(),
		})
	}

	return result
}

// FindByNumber finds a section by its hierarchical number.
func (t *SectionTree) FindByNumber(target SectionNumber) (*Section, bool) {
	numbered := t.AssignNumbers()
	for _, ns := range numbered {
		if ns.Number.Equals(target) {
			return ns.Section, true
		}
	}
	return nil, false
}

// GetSectionPath returns the path of sections from root to the given section.
func (t *SectionTree) GetSectionPath(section *Section) []*Section {
	var path []*Section
	current := section
	for current != nil && current.Parent != nil && current.Parent.Level > 0 {
		path = append([]*Section{current}, path...)
		current = current.Parent
	}
	if current != nil && current.Level > 0 {
		path = append([]*Section{current}, path...)
	}
	return path
}
