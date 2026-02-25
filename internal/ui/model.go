package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"mdexplore/internal/models"
)

// Model represents the state of the TUI application.
type Model struct {
	Filename string
	TOC      models.TableOfContents
	Selected int    // Currently selected heading index
	Error    error  // Error state (nil if no error)
	Quitting bool   // Whether the user has requested to quit
	Width    int    // Terminal width
	Height   int    // Terminal height
}

// InitialModel creates a new model with the given filename and TOC.
func InitialModel(filename string, toc models.TableOfContents) Model {
	return Model{
		Filename: filename,
		TOC:      toc,
		Selected: 0,
		Error:    nil,
		Quitting: false,
	}
}

// ErrorModel creates a model in error state.
func ErrorModel(err error) Model {
	return Model{
		Filename: "",
		TOC:      models.TableOfContents{},
		Selected: 0,
		Error:    err,
		Quitting: false,
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

// IsEmpty returns true if the TOC has no headings.
func (m Model) IsEmpty() bool {
	return m.TOC.IsEmpty()
}

// CanNavigateUp returns true if we can navigate up.
func (m Model) CanNavigateUp() bool {
	return m.Selected > 0
}

// CanNavigateDown returns true if we can navigate down.
func (m Model) CanNavigateDown() bool {
	return m.Selected < len(m.TOC.Headings)-1
}
