package benchmark

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"mdexplore/internal/parser"
)

func BenchmarkParser_SmallFile(b *testing.B) {
	content := generateMarkdown(100, 5) // 100 headings, 5 paragraphs
	tmpDir := b.TempDir()
	tmpFile := filepath.Join(tmpDir, "small.md")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		b.Fatal(err)
	}

	p := parser.NewGoldmarkParser()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := p.Parse(tmpFile)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParser_MediumFile(b *testing.B) {
	content := generateMarkdown(500, 20) // 500 headings, 20 paragraphs
	tmpDir := b.TempDir()
	tmpFile := filepath.Join(tmpDir, "medium.md")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		b.Fatal(err)
	}

	p := parser.NewGoldmarkParser()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := p.Parse(tmpFile)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParser_LargeFile(b *testing.B) {
	content := generateMarkdown(1000, 50) // 1000 headings, 50 paragraphs
	tmpDir := b.TempDir()
	tmpFile := filepath.Join(tmpDir, "large.md")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		b.Fatal(err)
	}

	p := parser.NewGoldmarkParser()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := p.Parse(tmpFile)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParser_OneMBFile(b *testing.B) {
	// Create approximately 1MB file
	content := generateMarkdown(5000, 100) // Lots of content
	tmpDir := b.TempDir()
	tmpFile := filepath.Join(tmpDir, "onemb.md")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		b.Fatal(err)
	}

	// Verify size
	info, err := os.Stat(tmpFile)
	if err != nil {
		b.Fatal(err)
	}
	b.Logf("File size: %d bytes (%.2f MB)", info.Size(), float64(info.Size())/(1024*1024))

	p := parser.NewGoldmarkParser()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := p.Parse(tmpFile)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// generateMarkdown creates markdown content with specified number of headings and paragraphs
func generateMarkdown(numHeadings, numParagraphs int) string {
	var sb strings.Builder

	for i := 0; i < numHeadings; i++ {
		level := (i % 6) + 1 // Levels 1-6
		sb.WriteString(fmt.Sprintf("%s Heading %d\n\n", strings.Repeat("#", level), i+1))

		// Add some paragraphs after each heading
		for j := 0; j < numParagraphs/numHeadings+1; j++ {
			sb.WriteString("Lorem ipsum dolor sit amet, consectetur adipiscing elit. ")
			sb.WriteString("Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. ")
			sb.WriteString("Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.\n\n")
		}
	}

	return sb.String()
}
