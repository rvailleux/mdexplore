package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mdexplore/internal/models"
	"mdexplore/internal/parser"
)

func TestTOCDisplay_FromSampleFile(t *testing.T) {
	p := parser.NewGoldmarkParser()
	result, err := p.Parse("../fixtures/sample_with_headings.md")
	require.NoError(t, err)

	// Verify TOC structure
	assert.Equal(t, "../fixtures/sample_with_headings.md", result.Source)
	assert.Equal(t, 8, result.Count())

	// Verify hierarchical structure
	expected := []struct {
		Level int
		Text  string
	}{
		{1, "Introduction"},
		{2, "Installation"},
		{3, "Requirements"},
		{3, "Setup"},
		{2, "Usage"},
		{3, "Basic Commands"},
		{3, "Advanced Options"},
		{2, "License"},
	}

	for i, exp := range expected {
		if i < len(result.Headings) {
			assert.Equal(t, exp.Level, result.Headings[i].Level, "level at index %d", i)
			assert.Equal(t, exp.Text, result.Headings[i].Text, "text at index %d", i)
		}
	}
}

func TestTOCDisplay_EmptyFile(t *testing.T) {
	p := parser.NewGoldmarkParser()
	result, err := p.Parse("../fixtures/empty.md")
	require.NoError(t, err)

	assert.True(t, result.IsEmpty())
	assert.Equal(t, 0, result.Count())
}

func TestTOCDisplay_BuildHierarchy(t *testing.T) {
	// Test building hierarchical representation
	tocData := models.TableOfContents{
		Headings: []models.Heading{
			{Level: 1, Text: "Root", LineNumber: 1},
			{Level: 2, Text: "Child 1", LineNumber: 2},
			{Level: 3, Text: "Grandchild 1", LineNumber: 3},
			{Level: 3, Text: "Grandchild 2", LineNumber: 4},
			{Level: 2, Text: "Child 2", LineNumber: 5},
		},
		Source: "test.md",
	}

	assert.Equal(t, 5, tocData.Count())
	assert.Equal(t, 1, len(tocData.HeadingsByLevel(1)))
	assert.Equal(t, 2, len(tocData.HeadingsByLevel(2)))
	assert.Equal(t, 2, len(tocData.HeadingsByLevel(3)))
}
