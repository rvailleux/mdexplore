package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"mdexplore/internal/errors"
	"mdexplore/internal/models"
)

// Styles for the TUI
var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4"))

	selectedStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00D26A")).
		Background(lipgloss.Color("#2D2D2D"))

	normalStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DDDDDD"))

	treeStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888"))

	errorStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF5F56"))

	helpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888"))

	emptyStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Italic(true)
)

// View renders the TUI based on the current model state.
func (m Model) View() string {
	if m.HasError() {
		return m.renderError()
	}

	if m.IsEmpty() {
		return m.renderEmpty()
	}

	return m.renderTOC()
}

// renderError renders an error message.
func (m Model) renderError() string {
	var errorIcon string
	var message string

	switch m.Error.(type) {
	case errors.FileNotFoundError:
		errorIcon = "❌"
		message = m.Error.Error()
	case errors.PermissionDeniedError:
		errorIcon = "🔒"
		message = m.Error.Error()
	case errors.InvalidFileError:
		errorIcon = "📁"
		message = m.Error.Error()
	case errors.FileTooLargeError:
		errorIcon = "📄"
		message = m.Error.Error()
	default:
		errorIcon = "⚠️"
		message = m.Error.Error()
	}

	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(errorStyle.Render(fmt.Sprintf("%s Error", errorIcon)))
	b.WriteString("\n\n")
	b.WriteString(normalStyle.Render(message))
	b.WriteString("\n\n")
	b.WriteString(helpStyle.Render("Press any key to exit"))
	b.WriteString("\n")

	return b.String()
}

// renderEmpty renders a message when the file has no headings.
func (m Model) renderEmpty() string {
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(titleStyle.Render(fmt.Sprintf("📄 %s - Table of Contents", m.Filename)))
	b.WriteString("\n\n")
	b.WriteString(emptyStyle.Render("  No headings found in this file."))
	b.WriteString("\n\n")
	b.WriteString(helpStyle.Render("[↑/↓] Navigate  [q] Quit"))
	b.WriteString("\n")

	return b.String()
}

// renderTOC renders the table of contents.
func (m Model) renderTOC() string {
	var b strings.Builder

	// Title
	b.WriteString("\n")
	title := titleStyle.Render(fmt.Sprintf("📄 %s - Table of Contents", m.Filename))
	b.WriteString(title)
	b.WriteString("\n\n")

	// TOC items
	for i, heading := range m.TOC.Headings {
		line := renderHeading(heading, i, m.Selected)
		if i == m.Selected {
			b.WriteString(selectedStyle.Render(line))
		} else {
			b.WriteString(normalStyle.Render(line))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(helpStyle.Render("[↑/↓] Navigate  [q] Quit"))
	b.WriteString("\n")

	return b.String()
}

// renderHeading renders a single heading with tree characters.
func renderHeading(heading models.Heading, index, selected int) string {
	indent := strings.Repeat("  ", heading.Level-1)
	prefix := getTreePrefix(heading.Level)
	return fmt.Sprintf("  %s%s %s", indent, prefix, heading.Text)
}

// getTreePrefix returns the appropriate tree character for the heading level.
func getTreePrefix(level int) string {
	switch level {
	case 1:
		return "●"
	case 2:
		return "├──"
	case 3:
		return "└──"
	case 4, 5, 6:
		return "└──"
	default:
		return "●"
	}
}
