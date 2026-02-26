package models

import (
	"testing"
)

// T003: Unit test for Section model
func TestSectionCreation(t *testing.T) {
	section := &Section{
		ID:        "L10",
		Level:     2,
		Title:     "Getting Started",
		StartLine: 10,
		EndLine:   50,
		RawContent: "## Getting Started\n\nThis is content.",
		Children:  nil,
		Parent:    nil,
	}

	if section.ID != "L10" {
		t.Errorf("Expected ID 'L10', got '%s'", section.ID)
	}
	if section.Level != 2 {
		t.Errorf("Expected Level 2, got %d", section.Level)
	}
	if section.Title != "Getting Started" {
		t.Errorf("Expected Title 'Getting Started', got '%s'", section.Title)
	}
	if section.StartLine != 10 {
		t.Errorf("Expected StartLine 10, got %d", section.StartLine)
	}
	if section.EndLine != 50 {
		t.Errorf("Expected EndLine 50, got %d", section.EndLine)
	}
}

func TestSectionHasChildren(t *testing.T) {
	// Section without children
	sectionNoChildren := &Section{
		ID:       "L10",
		Children: nil,
	}
	if sectionNoChildren.HasChildren() {
		t.Error("Expected HasChildren() to return false for nil Children")
	}

	// Section with empty children slice
	sectionEmptyChildren := &Section{
		ID:       "L10",
		Children: []*Section{},
	}
	if sectionEmptyChildren.HasChildren() {
		t.Error("Expected HasChildren() to return false for empty Children")
	}

	// Section with children
	sectionWithChildren := &Section{
		ID: "L10",
		Children: []*Section{
			{ID: "L20"},
		},
	}
	if !sectionWithChildren.HasChildren() {
		t.Error("Expected HasChildren() to return true for section with Children")
	}
}

func TestSectionContainsLine(t *testing.T) {
	section := &Section{
		StartLine: 10,
		EndLine:   50,
	}

	// Line before section
	if section.ContainsLine(5) {
		t.Error("Expected ContainsLine(5) to be false for section starting at 10")
	}

	// Line at start
	if !section.ContainsLine(10) {
		t.Error("Expected ContainsLine(10) to be true")
	}

	// Line in middle
	if !section.ContainsLine(30) {
		t.Error("Expected ContainsLine(30) to be true")
	}

	// Line at end
	if !section.ContainsLine(50) {
		t.Error("Expected ContainsLine(50) to be true")
	}

	// Line after section
	if section.ContainsLine(51) {
		t.Error("Expected ContainsLine(51) to be false for section ending at 50")
	}
}

func TestSectionGetDepth(t *testing.T) {
	// H1 section (depth 0)
	h1 := &Section{Level: 1}
	if h1.GetDepth() != 0 {
		t.Errorf("Expected GetDepth() = 0 for H1, got %d", h1.GetDepth())
	}

	// H2 section (depth 1)
	h2 := &Section{Level: 2}
	if h2.GetDepth() != 1 {
		t.Errorf("Expected GetDepth() = 1 for H2, got %d", h2.GetDepth())
	}

	// H3 section (depth 2)
	h3 := &Section{Level: 3}
	if h3.GetDepth() != 2 {
		t.Errorf("Expected GetDepth() = 2 for H3, got %d", h3.GetDepth())
	}
}

// T004: Unit test for SectionTree
func TestSectionTreeCreation(t *testing.T) {
	root := &Section{
		ID:       "root",
		Level:    0,
		Title:    "Root",
		Children: []*Section{},
	}

	tree := &SectionTree{
		Root:     root,
		Source:   "/path/to/file.md",
		Sections: []*Section{},
		ByID:     make(map[string]*Section),
	}

	if tree.Root != root {
		t.Error("Expected tree.Root to be set")
	}
	if tree.Source != "/path/to/file.md" {
		t.Errorf("Expected Source '/path/to/file.md', got '%s'", tree.Source)
	}
}

func TestSectionTreeGetH1Sections(t *testing.T) {
	// Create H1 sections
	h1_1 := &Section{ID: "L10", Level: 1, Title: "Introduction"}
	h1_2 := &Section{ID: "L50", Level: 1, Title: "Getting Started"}

	// Create tree with root containing H1 sections
	tree := &SectionTree{
		Root: &Section{
			ID:       "root",
			Children: []*Section{h1_1, h1_2},
		},
		Sections: []*Section{h1_1, h1_2},
		ByID: map[string]*Section{
			"L10": h1_1,
			"L50": h1_2,
		},
	}

	h1Sections := tree.GetH1Sections()
	if len(h1Sections) != 2 {
		t.Errorf("Expected 2 H1 sections, got %d", len(h1Sections))
	}
	if h1Sections[0].ID != "L10" {
		t.Errorf("Expected first H1 ID 'L10', got '%s'", h1Sections[0].ID)
	}
	if h1Sections[1].ID != "L50" {
		t.Errorf("Expected second H1 ID 'L50', got '%s'", h1Sections[1].ID)
	}
}

func TestSectionTreeFindByID(t *testing.T) {
	section := &Section{ID: "L10", Level: 1, Title: "Introduction"}
	tree := &SectionTree{
		ByID: map[string]*Section{
			"L10": section,
		},
	}

	// Find existing
	found, ok := tree.FindByID("L10")
	if !ok {
		t.Error("Expected to find section with ID 'L10'")
	}
	if found != section {
		t.Error("Expected found section to be the same as original")
	}

	// Find non-existing
	_, ok = tree.FindByID("L99")
	if ok {
		t.Error("Expected not to find section with ID 'L99'")
	}
}

// T005: Unit test for GetAllDescendants
func TestSectionGetAllDescendants(t *testing.T) {
	// Create hierarchy:
	// H1 (L10)
	//   H2 (L20)
	//     H3 (L30)
	//   H2 (L40)

	h3 := &Section{ID: "L30", Level: 3}
	h2_1 := &Section{ID: "L20", Level: 2, Children: []*Section{h3}}
	h3.Parent = h2_1
	h2_2 := &Section{ID: "L40", Level: 2}
	h1 := &Section{ID: "L10", Level: 1, Children: []*Section{h2_1, h2_2}}
	h2_1.Parent = h1
	h2_2.Parent = h1

	descendants := h1.GetAllDescendants()
	if len(descendants) != 3 {
		t.Errorf("Expected 3 descendants from H1, got %d", len(descendants))
	}

	// Check H2 only has H3 as descendant
	h2Descendants := h2_1.GetAllDescendants()
	if len(h2Descendants) != 1 {
		t.Errorf("Expected 1 descendant from H2, got %d", len(h2Descendants))
	}

	// Check H3 has no descendants
	h3Descendants := h3.GetAllDescendants()
	if len(h3Descendants) != 0 {
		t.Errorf("Expected 0 descendants from H3, got %d", len(h3Descendants))
	}
}
