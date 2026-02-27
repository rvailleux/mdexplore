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
		// Recalculate viewport height and clamp scroll offset if needed
		if m.ViewMode == ViewContent {
			newViewportHeight := CalculateViewportHeight(m.Height)
			maxOffset := CalculateMaxOffset(m.ContentTotalLines, newViewportHeight)
			m.ContentScrollOffset = ClampOffset(m.ContentScrollOffset, maxOffset)
		}
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
		return m.handleContentViewKeys(msg)
	}

	// Handle TOC view mode
	return m.handleTOCViewKeys(msg)
}

// handleContentViewKeys handles keys when in content view mode.
func (m Model) handleContentViewKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	keyStr := msg.String()

	// Handle navigation and quit keys by string
	switch keyStr {
	case "esc", "left", "h":
		// Return to TOC view
		m.ViewMode = ViewTOC
		m.CurrentSection = nil
		m.Selected = m.ReturnIndex
		return m, nil
	case "q", "ctrl+c":
		m.Quitting = true
		return m, tea.Quit
	}

	// Handle scroll keys by key type (for arrow keys, page keys, home/end)
	switch msg.Type {
	case tea.KeyUp:
		if m.ContentScrollOffset > 0 {
			m.ContentScrollOffset--
		}
	case tea.KeyDown:
		maxOffset := CalculateMaxOffset(m.ContentTotalLines, m.ViewportHeight)
		if m.ContentScrollOffset < maxOffset {
			m.ContentScrollOffset++
		}
	case tea.KeyPgUp:
		m.ContentScrollOffset = max(0, m.ContentScrollOffset-m.ViewportHeight)
	case tea.KeyPgDown:
		maxOffset := CalculateMaxOffset(m.ContentTotalLines, m.ViewportHeight)
		m.ContentScrollOffset = min(maxOffset, m.ContentScrollOffset+m.ViewportHeight)
	case tea.KeyHome:
		m.ContentScrollOffset = 0
	case tea.KeyEnd:
		m.ContentScrollOffset = CalculateMaxOffset(m.ContentTotalLines, m.ViewportHeight)
	}

	// Handle vim-style navigation with rune keys
	if msg.Type == tea.KeyRunes && len(msg.Runes) == 1 {
		switch msg.Runes[0] {
		case 'k':
			if m.ContentScrollOffset > 0 {
				m.ContentScrollOffset--
			}
		case 'j':
			maxOffset := CalculateMaxOffset(m.ContentTotalLines, m.ViewportHeight)
			if m.ContentScrollOffset < maxOffset {
				m.ContentScrollOffset++
			}
		case 'g':
			m.ContentScrollOffset = 0
		case 'G':
			m.ContentScrollOffset = CalculateMaxOffset(m.ContentTotalLines, m.ViewportHeight)
		}
	}

	return m, nil
}

// handleTOCViewKeys handles keys when in TOC view mode.
func (m Model) handleTOCViewKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
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
		// Expand selected section or open content if leaf
		if selected := m.GetSelectedSection(); selected != nil {
			if selected.HasChildren() {
				// Parent section: expand/collapse
				if !m.IsExpanded(selected.ID) {
					m.Expand(selected.ID)
				}
			} else {
				// Leaf section: open content view
				m.ReturnIndex = m.Selected
				m.CurrentSection = selected
				m.ViewMode = ViewContent
				m.ContentScrollOffset = 0 // Reset scroll position
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
			m.ReturnIndex = m.Selected
			m.CurrentSection = selected
			m.ViewMode = ViewContent
			m.ContentScrollOffset = 0 // Reset scroll position
		}

	case "esc":
		// Quit from TOC view
		m.Quitting = true
		return m, tea.Quit
	}

	return m, nil
}
