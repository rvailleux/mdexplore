package ui

// CalculateMaxOffset returns the maximum valid scroll offset given content and viewport dimensions.
// Returns 0 if content fits within viewport (no scrolling needed).
func CalculateMaxOffset(contentTotalLines, viewportHeight int) int {
	if contentTotalLines <= viewportHeight {
		return 0
	}
	return contentTotalLines - viewportHeight
}

// ClampOffset ensures scroll offset stays within valid bounds [0, maxOffset].
func ClampOffset(offset, maxOffset int) int {
	if offset < 0 {
		return 0
	}
	if offset > maxOffset {
		return maxOffset
	}
	return offset
}

// CalculateViewportHeight returns the available height for content display.
// Accounts for title bar (2 lines), padding (2 lines), and help footer (1 line).
func CalculateViewportHeight(terminalHeight int) int {
	// Title (1) + newline (1) + content + padding (2) + help (1) + newline (1)
	const overhead = 6
	if terminalHeight <= overhead {
		return 1 // Minimum visible content
	}
	return terminalHeight - overhead
}

// CalculateScrollPercentage returns the scroll progress as a percentage (0-100).
// Returns 100 if content fits within viewport.
func CalculateScrollPercentage(scrollOffset, contentTotalLines, viewportHeight int) int {
	if contentTotalLines <= viewportHeight {
		return 100
	}
	maxOffset := contentTotalLines - viewportHeight
	if maxOffset == 0 {
		return 100
	}
	percentage := (scrollOffset * 100) / maxOffset
	if percentage < 0 {
		return 0
	}
	if percentage > 100 {
		return 100
	}
	return percentage
}

// GetVisibleRange returns the start and end line indices for the current viewport.
func GetVisibleRange(scrollOffset, contentTotalLines, viewportHeight int) (start, end int) {
	start = scrollOffset
	end = scrollOffset + viewportHeight
	if end > contentTotalLines {
		end = contentTotalLines
	}
	return start, end
}
