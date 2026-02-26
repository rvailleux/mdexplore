package unit

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"mdexplore/internal/renderer"
)

func TestRenderMarkdown_Headings(t *testing.T) {
	tests := []struct {
		name     string
		markdown string
		want     string
	}{
		{
			name:     "H1 heading",
			markdown: "# Heading 1",
			want:     "Heading 1",
		},
		{
			name:     "H2 heading",
			markdown: "## Heading 2",
			want:     "Heading 2",
		},
		{
			name:     "H3 heading",
			markdown: "### Heading 3",
			want:     "Heading 3",
		},
		{
			name:     "H4 heading",
			markdown: "#### Heading 4",
			want:     "Heading 4",
		},
		{
			name:     "H5 heading",
			markdown: "##### Heading 5",
			want:     "Heading 5",
		},
		{
			name:     "H6 heading",
			markdown: "###### Heading 6",
			want:     "Heading 6",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := renderer.RenderMarkdown(tt.markdown, 80)
			require.NoError(t, err)
			assert.Contains(t, got, tt.want)
		})
	}
}

func TestRenderMarkdown_UnorderedList(t *testing.T) {
	markdown := `- Item 1
- Item 2
- Item 3`

	got, err := renderer.RenderMarkdown(markdown, 80)
	require.NoError(t, err)

	// Glamour renders bullets as various characters
	assert.Contains(t, got, "Item 1")
	assert.Contains(t, got, "Item 2")
	assert.Contains(t, got, "Item 3")
}

func TestRenderMarkdown_OrderedList(t *testing.T) {
	markdown := `1. First item
2. Second item
3. Third item`

	got, err := renderer.RenderMarkdown(markdown, 80)
	require.NoError(t, err)

	assert.Contains(t, got, "First item")
	assert.Contains(t, got, "Second item")
	assert.Contains(t, got, "Third item")
}

func TestRenderMarkdown_TaskList(t *testing.T) {
	markdown := `- [ ] Unchecked task
- [x] Checked task`

	got, err := renderer.RenderMarkdown(markdown, 80)
	require.NoError(t, err)

	assert.Contains(t, got, "Unchecked task")
	assert.Contains(t, got, "Checked task")
}

func TestRenderMarkdown_CodeBlock(t *testing.T) {
	markdown := "```go\nfunc main() {\n    fmt.Println(\"Hello\")\n}\n```"

	got, err := renderer.RenderMarkdown(markdown, 80)
	require.NoError(t, err)

	assert.Contains(t, got, "func main()")
	assert.Contains(t, got, "Hello")
}

func TestRenderMarkdown_InlineCode(t *testing.T) {
	markdown := "Use the `printf` function to print output."

	got, err := renderer.RenderMarkdown(markdown, 80)
	require.NoError(t, err)

	assert.Contains(t, got, "printf")
}

func TestRenderMarkdown_BoldText(t *testing.T) {
	markdown := "This is **bold** text."

	got, err := renderer.RenderMarkdown(markdown, 80)
	require.NoError(t, err)

	assert.Contains(t, got, "bold")
}

func TestRenderMarkdown_ItalicText(t *testing.T) {
	markdown := "This is *italic* text."

	got, err := renderer.RenderMarkdown(markdown, 80)
	require.NoError(t, err)

	assert.Contains(t, got, "italic")
}

func TestRenderMarkdown_Strikethrough(t *testing.T) {
	markdown := "This is ~~strikethrough~~ text."

	got, err := renderer.RenderMarkdown(markdown, 80)
	require.NoError(t, err)

	assert.Contains(t, got, "strikethrough")
}

func TestRenderMarkdown_Link(t *testing.T) {
	markdown := "[Link text](https://example.com)"

	got, err := renderer.RenderMarkdown(markdown, 80)
	require.NoError(t, err)

	assert.Contains(t, got, "Link text")
}

func TestRenderMarkdown_Blockquote(t *testing.T) {
	markdown := "> This is a quoted text."

	got, err := renderer.RenderMarkdown(markdown, 80)
	require.NoError(t, err)

	assert.Contains(t, got, "quoted text")
}

func TestRenderMarkdown_HorizontalRule(t *testing.T) {
	markdown := "Text before\n\n---\n\nText after"

	got, err := renderer.RenderMarkdown(markdown, 80)
	require.NoError(t, err)

	assert.Contains(t, got, "Text before")
	assert.Contains(t, got, "Text after")
}

func TestRenderMarkdown_Table(t *testing.T) {
	markdown := `| Column 1 | Column 2 |
|----------|----------|
| Data 1   | Data 2   |`

	got, err := renderer.RenderMarkdown(markdown, 80)
	require.NoError(t, err)

	assert.Contains(t, got, "Column 1")
	assert.Contains(t, got, "Column 2")
	assert.Contains(t, got, "Data 1")
	assert.Contains(t, got, "Data 2")
}

func TestRenderMarkdown_EmptyContent(t *testing.T) {
	got, err := renderer.RenderMarkdown("", 80)
	require.NoError(t, err)
	// Empty content should not cause an error
	// Glamour returns "\n\n" for empty input, which is acceptable
	assert.True(t, got == "" || strings.TrimSpace(got) == "", "Expected empty or whitespace-only output")
}

func TestRenderMarkdown_WidthHandling(t *testing.T) {
	// Long paragraph that should wrap
	markdown := "This is a very long paragraph that should be wrapped when rendered with a narrow width. " +
		"It contains many words that need to be properly wrapped to fit within the specified width."

	// Test with narrow width
	got, err := renderer.RenderMarkdown(markdown, 40)
	require.NoError(t, err)

	// The content should be rendered
	assert.Contains(t, got, "paragraph")
}

func TestExtractSectionContent(t *testing.T) {
	// Create a temporary file for testing
	content := "Line 1\nLine 2\nLine 3\nLine 4\nLine 5"

	tmpFile := t.TempDir() + "/test.txt"
	err := renderer.WriteTestFile(tmpFile, content)
	require.NoError(t, err)

	got, err := renderer.ExtractSectionContent(tmpFile, 2, 4)
	require.NoError(t, err)

	// ExtractSectionContent extracts from startLine to endLine (inclusive start, exclusive end)
	// Lines 2-4 means lines 2 and 3 (0-indexed: 1 and 2)
	assert.Contains(t, got, "Line 2")
	assert.Contains(t, got, "Line 3")
}

func TestStyleContent(t *testing.T) {
	content := "Test content"
	got := renderer.StyleContent(content)

	// Styled content should contain the original text
	assert.Contains(t, got, "Test content")
}
