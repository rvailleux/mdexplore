package integration

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"mdexplore/internal/models"
	"mdexplore/internal/ui"
)

// TestScrollKeyHandling verifies that scroll keys update the model correctly
func TestScrollKeyHandling(t *testing.T) {
	// Create a section tree
	tree := &models.SectionTree{
		Sections: []*models.Section{
			{
				ID:        "sec1",
				Title:     "Test Section",
				Level:     1,
				StartLine: 1,
				EndLine:   100,
			},
		},
	}

	// Create model in content view mode
	m := ui.InitialModel("test.md", tree)
	m.ViewMode = ui.ViewContent
	m.CurrentSection = tree.Sections[0]
	m.ContentScrollOffset = 10
	m.ContentTotalLines = 100
	m.ViewportHeight = 24
	m.Height = 30

	// Test scroll down
	keyMsg := tea.KeyMsg{Type: tea.KeyDown}
	newModel, _ := m.Update(keyMsg)
	updatedModel := newModel.(ui.Model)

	if updatedModel.ContentScrollOffset != 11 {
		t.Errorf("Scroll down: expected offset 11, got %d", updatedModel.ContentScrollOffset)
	}

	// Test scroll up
	m.ContentScrollOffset = 10
	keyMsg = tea.KeyMsg{Type: tea.KeyUp}
	newModel, _ = m.Update(keyMsg)
	updatedModel = newModel.(ui.Model)

	if updatedModel.ContentScrollOffset != 9 {
		t.Errorf("Scroll up: expected offset 9, got %d", updatedModel.ContentScrollOffset)
	}

	// Test at top boundary
	m.ContentScrollOffset = 0
	keyMsg = tea.KeyMsg{Type: tea.KeyUp}
	newModel, _ = m.Update(keyMsg)
	updatedModel = newModel.(ui.Model)

	if updatedModel.ContentScrollOffset != 0 {
		t.Errorf("Scroll up at top: expected offset 0, got %d", updatedModel.ContentScrollOffset)
	}

	// Test at bottom boundary
	m.ContentScrollOffset = 76 // 100 - 24 = 76 max
	keyMsg = tea.KeyMsg{Type: tea.KeyDown}
	newModel, _ = m.Update(keyMsg)
	updatedModel = newModel.(ui.Model)

	if updatedModel.ContentScrollOffset != 76 {
		t.Errorf("Scroll down at bottom: expected offset 76, got %d", updatedModel.ContentScrollOffset)
	}
}

// TestScrollKeysWithString verifies key string matching
func TestScrollKeysWithString(t *testing.T) {
	tests := []struct {
		name        string
		keyType     tea.KeyType
		keyString   string
		initialOffset int
		expectedOffset int
	}{
		{"down arrow", tea.KeyDown, "down", 10, 11},
		{"up arrow", tea.KeyUp, "up", 10, 9},
		{"j key", tea.KeyRunes, "j", 10, 11},
		{"k key", tea.KeyRunes, "k", 10, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &models.SectionTree{
				Sections: []*models.Section{
					{ID: "sec1", Title: "Test", Level: 1, StartLine: 1, EndLine: 100},
				},
			}

			m := ui.InitialModel("test.md", tree)
			m.ViewMode = ui.ViewContent
			m.CurrentSection = tree.Sections[0]
			m.ContentScrollOffset = tt.initialOffset
			m.ContentTotalLines = 100
			m.ViewportHeight = 24

			var keyMsg tea.KeyMsg
			if tt.keyType == tea.KeyRunes {
				keyMsg = tea.KeyMsg{Type: tt.keyType, Runes: []rune(tt.keyString)}
			} else {
				keyMsg = tea.KeyMsg{Type: tt.keyType}
			}

			newModel, _ := m.Update(keyMsg)
			updatedModel := newModel.(ui.Model)

			if updatedModel.ContentScrollOffset != tt.expectedOffset {
				t.Errorf("Expected offset %d, got %d", tt.expectedOffset, updatedModel.ContentScrollOffset)
			}
		})
	}
}
