package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"mdexplore/internal/parser"
)

func TestMixedFormats_SetextOnly(t *testing.T) {
	p := parser.NewGoldmarkParser()
	result, err := p.Parse("../fixtures/setext_headings.md")
	require.NoError(t, err)

	// Verify all Setext headings are extracted with correct levels
	expected := []struct {
		Level int
		Text  string
	}{
		{1, "Introduction"},
		{2, "Section One"},
		{3, "ATX Subsection"}, // This is ATX inside Setext file
		{1, "Another H1"},
		{2, "Final Section"},
	}

	assert.Equal(t, len(expected), result.Count())
	for i, exp := range expected {
		assert.Equal(t, exp.Level, result.Headings[i].Level, "level at index %d", i)
		assert.Equal(t, exp.Text, result.Headings[i].Text, "text at index %d", i)
	}
}

func TestMixedFormats_MixedATXAndSetext(t *testing.T) {
	p := parser.NewGoldmarkParser()
	result, err := p.Parse("../fixtures/mixed_headings.md")
	require.NoError(t, err)

	// Verify both ATX and Setext headings are extracted
	expected := []struct {
		Level int
		Text  string
	}{
		{1, "ATX Heading 1"},
		{1, "Setext H1"},
		{2, "ATX Heading 2"},
		{2, "Setext H2"},
		{3, "ATX Heading 3"},
		{1, "Another Setext H1"},
		{2, "Final Setext H2"},
	}

	assert.Equal(t, len(expected), result.Count())
	for i, exp := range expected {
		assert.Equal(t, exp.Level, result.Headings[i].Level, "level at index %d", i)
		assert.Equal(t, exp.Text, result.Headings[i].Text, "text at index %d", i)
	}
}

func TestMixedFormats_OrderPreserved(t *testing.T) {
	p := parser.NewGoldmarkParser()
	result, err := p.Parse("../fixtures/mixed_headings.md")
	require.NoError(t, err)

	// Verify headings appear in document order
	assert.Equal(t, "ATX Heading 1", result.Headings[0].Text)
	assert.Equal(t, "Setext H1", result.Headings[1].Text)
	assert.Equal(t, "ATX Heading 2", result.Headings[2].Text)
	assert.Equal(t, "Setext H2", result.Headings[3].Text)
}
