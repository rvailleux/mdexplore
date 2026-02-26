package benchmark

import (
	"fmt"
	"testing"

	"mdexplore/internal/ui"
)

// BenchmarkScrollCalculations benchmarks viewport calculation functions
func BenchmarkCalculateMaxOffset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui.CalculateMaxOffset(10000, 24)
	}
}

func BenchmarkClampOffset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui.ClampOffset(500, 976)
	}
}

func BenchmarkCalculateScrollPercentage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui.CalculateScrollPercentage(500, 10000, 24)
	}
}

func BenchmarkGetVisibleRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui.GetVisibleRange(500, 10000, 24)
	}
}

// BenchmarkFullScrollOperation simulates a complete scroll operation
func BenchmarkFullScrollOperation(b *testing.B) {
	contentLines := 10000
	viewportHeight := 24
	scrollOffset := 500

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Simulate scroll down operation
		maxOffset := ui.CalculateMaxOffset(contentLines, viewportHeight)
		newOffset := scrollOffset + 1
		scrollOffset = ui.ClampOffset(newOffset, maxOffset)
		_, _ = ui.GetVisibleRange(scrollOffset, contentLines, viewportHeight)
		_ = ui.CalculateScrollPercentage(scrollOffset, contentLines, viewportHeight)
	}
}

// BenchmarkScrollPerformanceLargeDocument tests performance with large documents
func BenchmarkScrollPerformanceLargeDocument(b *testing.B) {
	documentSizes := []int{100, 1000, 10000}

	for _, size := range documentSizes {
		b.Run(fmt.Sprintf("%d_lines", size), func(b *testing.B) {
			contentLines := size
			viewportHeight := 24
			scrollOffset := 0

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// Simulate scrolling through the document
				maxOffset := ui.CalculateMaxOffset(contentLines, viewportHeight)
				scrollOffset = (scrollOffset + 10) % (maxOffset + 1)
				_, _ = ui.GetVisibleRange(scrollOffset, contentLines, viewportHeight)
				_ = ui.CalculateScrollPercentage(scrollOffset, contentLines, viewportHeight)
			}
		})
	}
}

// BenchmarkCalculateViewportHeight benchmarks viewport height calculation
func BenchmarkCalculateViewportHeight(b *testing.B) {
	terminalHeights := []int{24, 50, 100}

	for _, height := range terminalHeights {
		b.Run(fmt.Sprintf("height_%d", height), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ui.CalculateViewportHeight(height)
			}
		})
	}
}
