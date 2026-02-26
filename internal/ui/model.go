package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"mdexplore/internal/models"
	"mdexplore/internal/renderer"
)

// ViewMode represents the current display mode.
type ViewMode int

const (
	ViewTOC ViewMode = iota      // Table of contents navigation view
	ViewContent                  // Section content display view
)

// glamourRenderer wraps the renderer package function for use in the UI.
type glamourRenderer struct{}

// Render implements the renderer.Renderer interface.
func (r *glamourRenderer) Render(markdown string, width int) (string, error) {
	return renderer.RenderMarkdown(markdown, width)
}

// Model represents the state of the TUI application.
type Model struct {
	Filename         string
	Tree             *models.SectionTree
	Selected         int              // Currently selected index in visible sections
	Error            error            // Error state (nil if no error)
	Quitting         bool             // Whether the user has requested to quit
	Width            int              // Terminal width
	Height           int              // Terminal height
	ViewMode         ViewMode         // Current view mode
	ExpandedSections map[string]bool  // Set of expanded section IDs
	CurrentSection   *models.Section  // Section being viewed (in content mode)
	ReturnIndex      int              // Position to restore when returning from content view
	markdownRenderer renderer.Renderer // Markdown renderer for content display
}

// InitialModel creates a new model with the given filename and section tree.
func InitialModel(filename string, tree *models.SectionTree) Model {
	return Model{
		Filename:         filename,
		Tree:             tree,
		Selected:         0,
		Error:            nil,
		Quitting:         false,
		ViewMode:         ViewTOC,
		ExpandedSections: make(map[string]bool),
		markdownRenderer: &glamourRenderer{},
	}
}

// InitialModelWithSelection creates a new model with a pre-selected section.
func InitialModelWithSelection(filename string, tree *models.SectionTree, targetSection *models.Section) Model {
	m := Model{
		Filename:         filename,
		Tree:             tree,
		Selected:         0,
		Error:            nil,
		Quitting:         false,
		ViewMode:         ViewContent, // Start in content view
		ExpandedSections: make(map[string]bool),
		CurrentSection:   targetSection,
		markdownRenderer: &glamourRenderer{},
	}

	// Expand all parent sections to make target visible
	if targetSection != nil && tree != nil {
		path := tree.GetSectionPath(targetSection)
		for _, section := range path {
			m.Expand(section.ID)
		}

		// Find the index of the target section in visible sections
		visible := m.GetVisibleSections()
		for i, section := range visible {
			if section.ID == targetSection.ID {
				m.Selected = i
				break
			}
		}
	}

	return m
}

// ErrorModel creates a model in error state.
func ErrorModel(err error) Model {
	return Model{
		Filename:         "",
		Tree:             nil,
		Selected:         0,
		Error:            err,
		Quitting:         false,
		ViewMode:         ViewTOC,
		ExpandedSections: make(map[string]bool),
	}
}

// Init initializes the TUI.
func (m Model) Init() tea.Cmd {
	return nil
}

// HasError returns true if the model is in an error state.
func (m Model) HasError() bool {
	return m.Error != nil
}

// IsEmpty returns true if the tree has no sections.
func (m Model) IsEmpty() bool {
	return m.Tree == nil || len(m.Tree.GetH1Sections()) == 0
}

// GetVisibleSections returns the list of currently visible sections based on expansion state.
func (m Model) GetVisibleSections() []*models.Section {
	if m.Tree == nil {
		return nil
	}
	return m.Tree.GetFlattenedVisible(m.ExpandedSections)
}

// CanNavigateUp returns true if we can navigate up.
func (m Model) CanNavigateUp() bool {
	return m.Selected > 0
}

// CanNavigateDown returns true if we can navigate down.
func (m Model) CanNavigateDown() bool {
	visible := m.GetVisibleSections()
	return m.Selected < len(visible)-1
}

// IsExpanded returns true if the given section ID is expanded.
func (m Model) IsExpanded(sectionID string) bool {
	return m.ExpandedSections[sectionID]
}

// ToggleExpanded toggles the expansion state of a section.
func (m Model) ToggleExpanded(sectionID string) {
	m.ExpandedSections[sectionID] = !m.ExpandedSections[sectionID]
}

// Expand marks a section as expanded.
func (m Model) Expand(sectionID string) {
	m.ExpandedSections[sectionID] = true
}

// Collapse marks a section as collapsed.
func (m Model) Collapse(sectionID string) {
	m.ExpandedSections[sectionID] = false
}

// GetSelectedSection returns the currently selected section.
func (m Model) GetSelectedSection() *models.Section {
	visible := m.GetVisibleSections()
	if m.Selected < 0 || m.Selected >= len(visible) {
		return nil
	}
	return visible[m.Selected]
}
