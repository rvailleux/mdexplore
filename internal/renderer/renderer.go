package renderer

import (
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

// RenderMarkdown renders markdown content as styled terminal output.
func RenderMarkdown(content string, width int) (string, error) {
	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return "", err
	}

	return renderer.Render(content)
}

// ExtractSectionContent reads a file and extracts content between start and end lines.
func ExtractSectionContent(filepath string, startLine, endLine int) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(content), "\n")

	// Adjust for 1-based line numbers
	start := startLine - 1
	if start < 0 {
		start = 0
	}
	if start > len(lines) {
		start = len(lines)
	}

	end := endLine
	if end > len(lines) {
		end = len(lines)
	}

	sectionLines := lines[start:end]
	return strings.Join(sectionLines, "\n"), nil
}

// StyleContent applies basic styling to raw content when Glamour fails.
func StyleContent(content string) string {
	return lipgloss.NewStyle().
		Padding(1, 2).
		Render(content)
}
