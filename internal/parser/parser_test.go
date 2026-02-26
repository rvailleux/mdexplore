package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mdexplore/internal/errors"
	"mdexplore/internal/models"
)

func TestGoldmarkParser_Parse(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []models.Heading
		wantErr  bool
	}{
		{
			name:    "ATX headings",
			content: "# H1\n## H2\n### H3\n",
			expected: []models.Heading{
				{Level: 1, Text: "H1", LineNumber: 1},
				{Level: 2, Text: "H2", LineNumber: 2},
				{Level: 3, Text: "H3", LineNumber: 3},
			},
		},
		{
			name:    "no headings",
			content: "Just plain text.\nNo headings here.\n",
			expected: []models.Heading{},
		},
		{
			name:    "headings with text",
			content: "## Installation\nFollow these steps.\n## Usage\nBasic usage info.\n",
			expected: []models.Heading{
				{Level: 2, Text: "Installation", LineNumber: 1},
				{Level: 2, Text: "Usage", LineNumber: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temp file
			tmpDir := t.TempDir()
			tmpFile := filepath.Join(tmpDir, "test.md")
			err := os.WriteFile(tmpFile, []byte(tt.content), 0644)
			require.NoError(t, err)

			parser := NewGoldmarkParser()
			result, err := parser.Parse(tmpFile)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, len(tt.expected), len(result.Headings))
			for i, exp := range tt.expected {
				assert.Equal(t, exp.Level, result.Headings[i].Level, "level mismatch at index %d", i)
				assert.Equal(t, exp.Text, result.Headings[i].Text, "text mismatch at index %d", i)
			}
		})
	}
}

func TestGoldmarkParser_Parse_FileNotFound(t *testing.T) {
	parser := NewGoldmarkParser()
	_, err := parser.Parse("/nonexistent/path/file.md")
	assert.Error(t, err)
	assert.IsType(t, errors.FileNotFoundError{}, err)
}

func TestGoldmarkParser_Parse_Directory(t *testing.T) {
	tmpDir := t.TempDir()
	parser := NewGoldmarkParser()
	_, err := parser.Parse(tmpDir)
	assert.Error(t, err)
	assert.IsType(t, errors.InvalidFileError{}, err)
}

func TestGoldmarkParser_ExtractHeadingsFromSampleFile(t *testing.T) {
	parser := NewGoldmarkParser()
	result, err := parser.Parse("../../tests/fixtures/sample_with_headings.md")
	require.NoError(t, err)
	assert.Equal(t, 8, len(result.Headings))
	assert.Equal(t, "Introduction", result.Headings[0].Text)
	assert.Equal(t, 1, result.Headings[0].Level)
	assert.Equal(t, "Installation", result.Headings[1].Text)
	assert.Equal(t, 2, result.Headings[1].Level)
}

func TestGoldmarkParser_Parse_EmptyFile(t *testing.T) {
	parser := NewGoldmarkParser()
	result, err := parser.Parse("../../tests/fixtures/empty.md")
	require.NoError(t, err)
	assert.True(t, len(result.Headings) == 0)
}

// User Story 2: Error handling tests

func TestGoldmarkParser_Parse_FileNotFoundError(t *testing.T) {
	parser := NewGoldmarkParser()
	_, err := parser.Parse("/nonexistent/directory/file.md")
	require.Error(t, err)
	assert.IsType(t, errors.FileNotFoundError{}, err)
	assert.Contains(t, err.Error(), "file not found")
}

// User Story 3: Setext heading tests

func TestGoldmarkParser_Parse_SetextH1(t *testing.T) {
	content := `Introduction
============

This is content.`

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "setext_h1.md")
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	require.NoError(t, err)

	parser := NewGoldmarkParser()
	result, err := parser.Parse(tmpFile)
	require.NoError(t, err)
	assert.Equal(t, 1, len(result.Headings))
	assert.Equal(t, 1, result.Headings[0].Level)
	assert.Equal(t, "Introduction", result.Headings[0].Text)
}

func TestGoldmarkParser_Parse_SetextH2(t *testing.T) {
	content := `Section
-------

This is content.`

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "setext_h2.md")
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	require.NoError(t, err)

	parser := NewGoldmarkParser()
	result, err := parser.Parse(tmpFile)
	require.NoError(t, err)
	assert.Equal(t, 1, len(result.Headings))
	assert.Equal(t, 2, result.Headings[0].Level)
	assert.Equal(t, "Section", result.Headings[0].Text)
}

func TestGoldmarkParser_Parse_SetextHeadingsFromSampleFile(t *testing.T) {
	parser := NewGoldmarkParser()
	result, err := parser.Parse("../../tests/fixtures/setext_headings.md")
	require.NoError(t, err)
	// Should have: Introduction (H1), Section One (H2), ATX Subsection (H3), Another H1 (H1), Final Section (H2)
	assert.Equal(t, 5, len(result.Headings))
	assert.Equal(t, 1, result.Headings[0].Level)
	assert.Equal(t, "Introduction", result.Headings[0].Text)
	assert.Equal(t, 2, result.Headings[1].Level)
	assert.Equal(t, "Section One", result.Headings[1].Text)
}

func TestGoldmarkParser_Parse_MixedHeadingsFromSampleFile(t *testing.T) {
	parser := NewGoldmarkParser()
	result, err := parser.Parse("../../tests/fixtures/mixed_headings.md")
	require.NoError(t, err)
	// Should have: ATX H1, Setext H1, ATX H2, Setext H2, ATX H3, Another Setext H1, Final Setext H2 = 7 headings
	assert.Equal(t, 7, len(result.Headings))
	// Check ATX headings are preserved
	assert.Equal(t, 1, result.Headings[0].Level)
	assert.Equal(t, "ATX Heading 1", result.Headings[0].Text)
	// Check Setext headings are detected
	assert.Equal(t, 1, result.Headings[1].Level)
	assert.Equal(t, "Setext H1", result.Headings[1].Text)
}

func TestGoldmarkParser_Parse_HeadingsInCodeBlocksExcluded(t *testing.T) {
	parser := NewGoldmarkParser()
	result, err := parser.Parse("../../tests/fixtures/headings_in_code_blocks.md")
	require.NoError(t, err)
	// Should have only 4 valid headings, not the ones in code blocks
	assert.Equal(t, 4, len(result.Headings))
	assert.Equal(t, "Valid Heading", result.Headings[0].Text)
	assert.Equal(t, "Another Valid Heading", result.Headings[1].Text)
	assert.Equal(t, "Third Valid Heading", result.Headings[2].Text)
	assert.Equal(t, "Final Valid Heading", result.Headings[3].Text)
}

func TestGoldmarkParser_Parse_FrontmatterSkipped(t *testing.T) {
	parser := NewGoldmarkParser()
	result, err := parser.Parse("../../tests/fixtures/with_frontmatter.md")
	require.NoError(t, err)
	// Should have 3 headings: Introduction (H1), Section One (H2), Section Two (H2)
	// The frontmatter should be skipped (not treated as content)
	assert.Equal(t, 3, len(result.Headings))
	assert.Equal(t, "Introduction", result.Headings[0].Text)
	assert.Equal(t, 1, result.Headings[0].Level)
	assert.Equal(t, "Section One", result.Headings[1].Text)
	assert.Equal(t, 2, result.Headings[1].Level)
	assert.Equal(t, "Section Two", result.Headings[2].Text)
	assert.Equal(t, 2, result.Headings[2].Level)
}

func TestGoldmarkParser_Parse_PermissionDeniedError(t *testing.T) {
	// Skip this test if running as root (root can read any file)
	if os.Getuid() == 0 {
		t.Skip("Skipping permission test when running as root")
	}

	// Create a file with no read permissions
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "noperm.md")
	err := os.WriteFile(tmpFile, []byte("# Test"), 0200) // Write-only permission
	require.NoError(t, err)

	parser := NewGoldmarkParser()
	_, err = parser.Parse(tmpFile)
	require.Error(t, err)
	assert.IsType(t, errors.PermissionDeniedError{}, err)
}

func TestGoldmarkParser_Parse_InvalidFileError(t *testing.T) {
	// Try to parse a directory
	tmpDir := t.TempDir()
	parser := NewGoldmarkParser()
	_, err := parser.Parse(tmpDir)
	require.Error(t, err)
	assert.IsType(t, errors.InvalidFileError{}, err)
	assert.Contains(t, err.Error(), "expected a file, got directory")
}

// User Story 1: Line number tests

func TestGoldmarkParser_ParseSectionTree_EndLineCalculation(t *testing.T) {
	content := `# Introduction
This is the intro.
More content here.

# Getting Started
Steps to start.
More steps.

# Conclusion
Final thoughts.`

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "sections.md")
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	require.NoError(t, err)

	parser := NewGoldmarkParser()
	tree, err := parser.ParseSectionTree(tmpFile)
	require.NoError(t, err)

	// Should have 3 H1 sections
	require.Equal(t, 3, len(tree.GetH1Sections()))

	// First section: Introduction (L1 - L4, line before "Getting Started")
	intro := tree.GetH1Sections()[0]
	assert.Equal(t, "Introduction", intro.Title)
	assert.Equal(t, 1, intro.StartLine)
	assert.Equal(t, 4, intro.EndLine)

	// Second section: Getting Started (L5 - L7)
	gettingStarted := tree.GetH1Sections()[1]
	assert.Equal(t, "Getting Started", gettingStarted.Title)
	assert.Equal(t, 5, gettingStarted.StartLine)
	assert.Equal(t, 8, gettingStarted.EndLine)

	// Third section: Conclusion (L9 - end of file)
	conclusion := tree.GetH1Sections()[2]
	assert.Equal(t, "Conclusion", conclusion.Title)
	assert.Equal(t, 9, conclusion.StartLine)
	assert.Equal(t, 10, conclusion.EndLine)
}

func TestGoldmarkParser_ParseSectionTree_SingleLineSection(t *testing.T) {
	content := `# Heading Only
# Next Heading`

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "single_line.md")
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	require.NoError(t, err)

	parser := NewGoldmarkParser()
	tree, err := parser.ParseSectionTree(tmpFile)
	require.NoError(t, err)

	// First heading should have StartLine == EndLine
	first := tree.GetH1Sections()[0]
	assert.Equal(t, 1, first.StartLine)
	assert.Equal(t, 1, first.EndLine)

	// Second heading
	second := tree.GetH1Sections()[1]
	assert.Equal(t, 2, second.StartLine)
	assert.Equal(t, 2, second.EndLine)
}

func TestGoldmarkParser_ParseSectionTree_HierarchicalSections(t *testing.T) {
	content := `# Main Section
Main content.

## Subsection A
Sub A content.
More content.

## Subsection B
Sub B content.

### Deep Subsection
Deep content.`

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "hierarchical.md")
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	require.NoError(t, err)

	parser := NewGoldmarkParser()
	tree, err := parser.ParseSectionTree(tmpFile)
	require.NoError(t, err)

	// Should have 1 H1
	h1Sections := tree.GetH1Sections()
	require.Equal(t, 1, len(h1Sections))

	mainSection := h1Sections[0]
	assert.Equal(t, "Main Section", mainSection.Title)
	assert.Equal(t, 2, len(mainSection.Children)) // Two H2 children

	// Check subsection A
	subA := mainSection.Children[0]
	assert.Equal(t, "Subsection A", subA.Title)
	assert.Equal(t, 2, subA.Level)
	assert.Equal(t, 4, subA.StartLine)
	assert.Equal(t, 7, subA.EndLine)

	// Check subsection B
	subB := mainSection.Children[1]
	assert.Equal(t, "Subsection B", subB.Title)
	assert.Equal(t, 2, subB.Level)
	assert.Equal(t, 8, subB.StartLine)
	assert.Equal(t, 10, subB.EndLine) // Ends at line 10 (line before Deep Subsection at 11)
	assert.Equal(t, 1, len(subB.Children)) // One H3 child

	// Check deep subsection
	deep := subB.Children[0]
	assert.Equal(t, "Deep Subsection", deep.Title)
	assert.Equal(t, 3, deep.Level)
	assert.Equal(t, 11, deep.StartLine) // Starts at line 11
	assert.Equal(t, 12, deep.EndLine)
}
