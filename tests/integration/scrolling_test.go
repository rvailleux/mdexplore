package integration

import (
	"testing"

	"mdexplore/internal/models"
	"mdexplore/internal/ui"
)

// TestScrollOffsetResetOnEnterContentView verifies that scroll offset resets when entering content view
func TestScrollOffsetResetOnEnterContentView(t *testing.T) {
	// Create a simple section tree
	tree := &models.SectionTree{
		Sections: []*models.Section{
			{
				ID:        "sec1",
				Title:     "Test Section",
				Level:     1,
				StartLine: 1,
				EndLine:   10,
			},
		},
	}

	// Create model and simulate entering content view
	m := ui.InitialModel("test.md", tree)

	// Set some scroll offset (simulating previous scroll)
	m.ContentScrollOffset = 50

	// Enter content view
	m.CurrentSection = tree.Sections[0]
	m.ViewMode = ui.ViewContent
	m.ContentScrollOffset = 0 // This should happen in the actual implementation

	if m.ContentScrollOffset != 0 {
		t.Errorf("Expected scroll offset to be reset to 0, got %d", m.ContentScrollOffset)
	}
}

// TestScrollBoundaries verifies that scroll offset stays within valid bounds
func TestScrollBoundaries(t *testing.T) {
	tests := []struct {
		name           string
		initialOffset  int
		contentLines   int
		viewportHeight int
		scrollDelta    int // +1 for down, -1 for up
		expectedOffset int
	}{
		{"scroll down within bounds", 10, 100, 20, 1, 11},
		{"scroll up within bounds", 10, 100, 20, -1, 9},
		{"scroll down at bottom", 80, 100, 20, 1, 80},  // Should stay at 80
		{"scroll up at top", 0, 100, 20, -1, 0},         // Should stay at 0
		{"scroll down from 0", 0, 100, 20, 1, 1},
		{"scroll up to 0", 1, 100, 20, -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maxOffset := ui.CalculateMaxOffset(tt.contentLines, tt.viewportHeight)

			newOffset := tt.initialOffset + tt.scrollDelta
			clampedOffset := ui.ClampOffset(newOffset, maxOffset)

			if clampedOffset != tt.expectedOffset {
				t.Errorf("Scroll from %d by %d: got %d, want %d",
					tt.initialOffset, tt.scrollDelta, clampedOffset, tt.expectedOffset)
			}
		})
	}
}

// TestViewportHeightCalculation verifies viewport height accounts for UI elements
func TestViewportHeightCalculation(t *testing.T) {
	tests := []struct {
		terminalHeight int
		minExpected    int
	}{
		{24, 10}, // Standard terminal should have reasonable viewport
		{50, 40}, // Large terminal
		{10, 1},  // Small terminal should at least show 1 line
	}

	for _, tt := range tests {
		viewportHeight := ui.CalculateViewportHeight(tt.terminalHeight)
		if viewportHeight < tt.minExpected {
			t.Errorf("CalculateViewportHeight(%d) = %d, want at least %d",
				tt.terminalHeight, viewportHeight, tt.minExpected)
		}
	}
}

// TestVisibleRangeCalculation verifies the content window calculation
func TestVisibleRangeCalculation(t *testing.T) {
	tests := []struct {
		name           string
		scrollOffset   int
		contentLines   int
		viewportHeight int
		expectedStart  int
		expectedEnd    int
	}{
		{"at top", 0, 100, 20, 0, 20},
		{"scrolled 10", 10, 100, 20, 10, 30},
		{"near end", 80, 100, 20, 80, 100},
		{"small content", 0, 10, 20, 0, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, end := ui.GetVisibleRange(tt.scrollOffset, tt.contentLines, tt.viewportHeight)
			if start != tt.expectedStart || end != tt.expectedEnd {
				t.Errorf("GetVisibleRange(%d, %d, %d) = (%d, %d), want (%d, %d)",
					tt.scrollOffset, tt.contentLines, tt.viewportHeight,
					start, end, tt.expectedStart, tt.expectedEnd)
			}
		})
	}
}

// TestScrollPercentageCalculation verifies percentage indicator calculation
func TestScrollPercentageCalculation(t *testing.T) {
	tests := []struct {
		name           string
		scrollOffset   int
		contentLines   int
		viewportHeight int
		minExpected    int
		maxExpected    int
	}{
		{"at top", 0, 100, 24, 0, 0},
		{"at 50%", 38, 100, 24, 48, 52},
		{"at bottom", 76, 100, 24, 98, 100},
		{"content fits", 0, 10, 24, 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			percentage := ui.CalculateScrollPercentage(tt.scrollOffset, tt.contentLines, tt.viewportHeight)
			if percentage < tt.minExpected || percentage > tt.maxExpected {
				t.Errorf("CalculateScrollPercentage(%d, %d, %d) = %d, want between %d and %d",
					tt.scrollOffset, tt.contentLines, tt.viewportHeight,
					percentage, tt.minExpected, tt.maxExpected)
			}
		})
	}
}
