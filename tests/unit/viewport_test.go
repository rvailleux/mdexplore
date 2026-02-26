package unit

import (
	"testing"

	"mdexplore/internal/ui"
)

func TestCalculateMaxOffset(t *testing.T) {
	tests := []struct {
		name            string
		contentLines    int
		viewportHeight  int
		expectedMax     int
	}{
		{"content fits viewport", 10, 20, 0},
		{"content equals viewport", 20, 20, 0},
		{"content larger than viewport", 50, 20, 30},
		{"empty content", 0, 20, 0},
		{"single line content", 1, 20, 0},
		{"large content", 1000, 24, 976},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ui.CalculateMaxOffset(tt.contentLines, tt.viewportHeight)
			if result != tt.expectedMax {
				t.Errorf("CalculateMaxOffset(%d, %d) = %d; want %d",
					tt.contentLines, tt.viewportHeight, result, tt.expectedMax)
			}
		})
	}
}

func TestClampOffset(t *testing.T) {
	tests := []struct {
		name       string
		offset     int
		maxOffset  int
		expected   int
	}{
		{"within bounds", 10, 30, 10},
		{"at lower bound", 0, 30, 0},
		{"at upper bound", 30, 30, 30},
		{"below bounds", -5, 30, 0},
		{"above bounds", 35, 30, 30},
		{"negative max", 5, 0, 0},
		{"negative both", -5, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ui.ClampOffset(tt.offset, tt.maxOffset)
			if result != tt.expected {
				t.Errorf("ClampOffset(%d, %d) = %d; want %d",
					tt.offset, tt.maxOffset, result, tt.expected)
			}
		})
	}
}

func TestCalculateViewportHeight(t *testing.T) {
	tests := []struct {
		name           string
		terminalHeight int
		expected       int
	}{
		{"standard terminal", 24, 18},
		{"large terminal", 50, 44},
		{"small terminal", 10, 4},
		{"minimum terminal", 6, 1},
		{"very small terminal", 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ui.CalculateViewportHeight(tt.terminalHeight)
			if result != tt.expected {
				t.Errorf("CalculateViewportHeight(%d) = %d; want %d",
					tt.terminalHeight, result, tt.expected)
			}
		})
	}
}

func TestCalculateScrollPercentage(t *testing.T) {
	tests := []struct {
		name           string
		scrollOffset   int
		contentLines   int
		viewportHeight int
		expected       int
	}{
		{"at top", 0, 100, 24, 0},
		{"at middle", 38, 100, 24, 50},
		{"at bottom", 76, 100, 24, 100},
		{"content fits", 0, 20, 24, 100},
		{"single page", 0, 10, 10, 100},
		{"near top", 7, 100, 24, 9},
		{"near bottom", 70, 100, 24, 92},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ui.CalculateScrollPercentage(tt.scrollOffset, tt.contentLines, tt.viewportHeight)
			// Allow small variance due to integer division
			variance := 2
			if result < tt.expected-variance || result > tt.expected+variance {
				t.Errorf("CalculateScrollPercentage(%d, %d, %d) = %d; want ~%d",
					tt.scrollOffset, tt.contentLines, tt.viewportHeight, result, tt.expected)
			}
		})
	}
}

func TestGetVisibleRange(t *testing.T) {
	tests := []struct {
		name           string
		scrollOffset   int
		contentLines   int
		viewportHeight int
		expectedStart  int
		expectedEnd    int
	}{
		{"at top", 0, 100, 20, 0, 20},
		{"scrolled down", 10, 100, 20, 10, 30},
		{"near end", 80, 100, 20, 80, 100},
		{"at end", 80, 100, 20, 80, 100},
		{"content smaller than viewport", 0, 10, 20, 0, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, end := ui.GetVisibleRange(tt.scrollOffset, tt.contentLines, tt.viewportHeight)
			if start != tt.expectedStart || end != tt.expectedEnd {
				t.Errorf("GetVisibleRange(%d, %d, %d) = (%d, %d); want (%d, %d)",
					tt.scrollOffset, tt.contentLines, tt.viewportHeight,
					start, end, tt.expectedStart, tt.expectedEnd)
			}
		})
	}
}

func TestScrollUp(t *testing.T) {
	// Simulate scrolling up: offset decreases by 1, but not below 0
	offset := 10
	maxOffset := 50

	// Scroll up
	offset = ui.ClampOffset(offset-1, maxOffset)
	if offset != 9 {
		t.Errorf("After scroll up, offset = %d; want 9", offset)
	}

	// Scroll up at top
	offset = 0
	offset = ui.ClampOffset(offset-1, maxOffset)
	if offset != 0 {
		t.Errorf("Scroll up at top should stay at 0, got %d", offset)
	}
}

func TestScrollDown(t *testing.T) {
	// Simulate scrolling down: offset increases by 1, but not above max
	offset := 10
	maxOffset := 50

	// Scroll down
	offset = ui.ClampOffset(offset+1, maxOffset)
	if offset != 11 {
		t.Errorf("After scroll down, offset = %d; want 11", offset)
	}

	// Scroll down at bottom
	offset = maxOffset
	offset = ui.ClampOffset(offset+1, maxOffset)
	if offset != maxOffset {
		t.Errorf("Scroll down at bottom should stay at %d, got %d", maxOffset, offset)
	}
}
