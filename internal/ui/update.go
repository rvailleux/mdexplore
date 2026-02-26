package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Update handles messages and updates the model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyPress(msg)

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}

	return m, nil
}

// handleKeyPress handles keyboard input.
func (m Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// If there's an error, any key quits
	if m.HasError() {
		m.Quitting = true
		return m, tea.Quit
	}

	// Handle content view mode separately
	if m.ViewMode == ViewContent {
		switch msg.String() {
		case "esc":
			// Return to TOC view
			m.ViewMode = ViewTOC
			m.CurrentSection = nil
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		}
		return m, nil
	}

	// Handle TOC view mode
	switch msg.String() {
	case "q", "ctrl+c":
		m.Quitting = true
		return m, tea.Quit

	case "up", "k":
		if m.CanNavigateUp() {
			m.Selected--
		}

	case "down", "j":
		if m.CanNavigateDown() {
			m.Selected++
		}

	case "right", "l":
		// Expand selected section
		if selected := m.GetSelectedSection(); selected != nil {
			if selected.HasChildren() && !m.IsExpanded(selected.ID) {
				m.Expand(selected.ID)
			}
		}

	case "left", "h":
		// Collapse selected section
		if selected := m.GetSelectedSection(); selected != nil {
			if m.IsExpanded(selected.ID) {
				m.Collapse(selected.ID)
			}
		}

	case "enter":
		// View content of selected section
		if selected := m.GetSelectedSection(); selected != nil {
			m.CurrentSection = selected
			m.ViewMode = ViewContent
		}

	case "esc":
		// Quit from TOC view
		m.Quitting = true
		return m, tea.Quit
	}

	return m, nil
}
