package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableOfContents_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		headings []Heading
		expected bool
	}{
		{
			name:     "empty TOC",
			headings: []Heading{},
			expected: true,
		},
		{
			name: "non-empty TOC",
			headings: []Heading{
				{Level: 1, Text: "Heading 1", LineNumber: 1},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toc := TableOfContents{Headings: tt.headings}
			assert.Equal(t, tt.expected, toc.IsEmpty())
		})
	}
}

func TestTableOfContents_Count(t *testing.T) {
	toc := TableOfContents{
		Headings: []Heading{
			{Level: 1, Text: "H1", LineNumber: 1},
			{Level: 2, Text: "H2", LineNumber: 2},
			{Level: 3, Text: "H3", LineNumber: 3},
		},
	}
	assert.Equal(t, 3, toc.Count())
}

func TestTableOfContents_HeadingsByLevel(t *testing.T) {
	toc := TableOfContents{
		Headings: []Heading{
			{Level: 1, Text: "H1 First", LineNumber: 1},
			{Level: 2, Text: "H2 First", LineNumber: 2},
			{Level: 1, Text: "H1 Second", LineNumber: 3},
			{Level: 3, Text: "H3 First", LineNumber: 4},
			{Level: 2, Text: "H2 Second", LineNumber: 5},
		},
	}

	h1s := toc.HeadingsByLevel(1)
	assert.Equal(t, 2, len(h1s))
	assert.Equal(t, "H1 First", h1s[0].Text)
	assert.Equal(t, "H1 Second", h1s[1].Text)

	h2s := toc.HeadingsByLevel(2)
	assert.Equal(t, 2, len(h2s))

	h3s := toc.HeadingsByLevel(3)
	assert.Equal(t, 1, len(h3s))

	h4s := toc.HeadingsByLevel(4)
	assert.Equal(t, 0, len(h4s))
}
