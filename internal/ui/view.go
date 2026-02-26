package ui

import (
	"fmt"
	"os"
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

	numberStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF9F1C")).
		Bold(true)

	lineNumberStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666"))

	selectedLineNumberStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Background(lipgloss.Color("#2D2D2D"))

	selectedNumberStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFB347")).
		Background(lipgloss.Color("#2D2D2D")).
		Bold(true)

	errorStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF5F56"))

	helpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888"))

	emptyStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Italic(true)

	contentStyle = lipgloss.NewStyle().
		Padding(1, 2)
)

// View renders the TUI based on the current model state.
func (m Model) View() string {
	if m.HasError() {
		return m.renderError()
	}

	if m.IsEmpty() {
		return m.renderEmpty()
	}

	if m.ViewMode == ViewContent {
		return m.renderContent()
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

	// Get visible sections
	visibleSections := m.GetVisibleSections()

	// Get numbered sections for display
	numberedSections := m.Tree.AssignNumbers()
	sectionToNumber := make(map[string]*models.NumberedSection)
	for _, ns := range numberedSections {
		sectionToNumber[ns.Section.ID] = ns
	}

	// TOC items
	for i, section := range visibleSections {
		ns := sectionToNumber[section.ID]
		line := renderSectionWithNumber(section, ns, i, m.Selected, m.IsExpanded(section.ID))
		if i == m.Selected {
			b.WriteString(selectedStyle.Render(line))
		} else {
			b.WriteString(normalStyle.Render(line))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(helpStyle.Render("[↑/↓] Navigate  [→] Expand  [←] Collapse  [Enter] View  [q] Quit"))
	b.WriteString("\n")

	return b.String()
}

// renderContent renders the content view for the current section.
func (m Model) renderContent() string {
	if m.CurrentSection == nil {
		return m.renderTOC()
	}

	var b strings.Builder

	// Title bar
	b.WriteString("\n")
	title := titleStyle.Render(fmt.Sprintf("📄 %s - L%d-%d: %s",
		m.Filename,
		m.CurrentSection.StartLine,
		m.CurrentSection.EndLine,
		m.CurrentSection.Title))
	b.WriteString(title)
	b.WriteString("\n\n")

	// Content - use RawContent if available, otherwise extract from file
	content := m.CurrentSection.RawContent
	if content == "" {
		// Try to extract from file
		if extracted, err := extractSectionFromFile(m.Filename, m.CurrentSection.StartLine, m.CurrentSection.EndLine); err == nil {
			content = extracted
		}
	}
	b.WriteString(contentStyle.Render(content))

	b.WriteString("\n\n")
	b.WriteString(helpStyle.Render("[Esc] Return to navigation  [q] Quit"))
	b.WriteString("\n")

	return b.String()
}

// extractSectionFromFile reads content from a file between given line numbers.
func extractSectionFromFile(filepath string, startLine, endLine int) (string, error) {
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

// renderSectionWithNumber renders a single section with section number, line numbers and tree characters.
func renderSectionWithNumber(section *models.Section, ns *models.NumberedSection, index, selected int, isExpanded bool) string {
	var b strings.Builder

	// Section number (e.g., "1.1.")
	if ns != nil {
		b.WriteString(numberStyle.Render(fmt.Sprintf("%-6s ", ns.DisplayNumber)))
	}

	// Line numbers in format "L[start]-[end]"
	lineNum := fmt.Sprintf("L%d-%d", section.StartLine, section.EndLine)
	b.WriteString(lineNumberStyle.Render(fmt.Sprintf("%-10s ", lineNum)))

	// Indentation based on depth
	indent := strings.Repeat("  ", section.GetDepth())
	b.WriteString(indent)

	// Tree prefix based on level and expansion state
	prefix := getTreePrefix(section.Level, section.HasChildren(), isExpanded)
	b.WriteString(treeStyle.Render(prefix))

	// Section title
	b.WriteString(" ")
	b.WriteString(section.Title)

	return b.String()
}

// getTreePrefix returns the appropriate tree character for the section level.
func getTreePrefix(level int, hasChildren bool, isExpanded bool) string {
	if !hasChildren {
		switch level {
		case 1:
			return "●"
		case 2, 3, 4, 5, 6:
			return "└──"
		default:
			return "●"
		}
	}

	// Has children - show expansion indicator
	if isExpanded {
		switch level {
		case 1:
			return "▼"
		case 2:
			return "├──"
		case 3:
			return "└──"
		default:
			return "└──"
		}
	} else {
		switch level {
		case 1:
			return "▶"
		case 2:
			return "├──"
		case 3:
			return "└──"
		default:
			return "└──"
		}
	}
}
