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

	switch msg.String() {
	case "q", "ctrl+c", "esc":
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
	}

	return m, nil
}
