package parser

import (
	"bytes"
	"fmt"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"

	"mdexplore/internal/errors"
	"mdexplore/internal/models"
)

// Parser defines the interface for parsing markdown files.
type Parser interface {
	Parse(filepath string) (models.TableOfContents, error)
}

// GoldmarkParser implements Parser using the Goldmark library.
type GoldmarkParser struct {
	maxFileSize int64
}

// NewGoldmarkParser creates a new parser with default settings.
func NewGoldmarkParser() *GoldmarkParser {
	return &GoldmarkParser{
		maxFileSize: 10 * 1024 * 1024, // 10MB default
	}
}

// SetMaxFileSize sets the maximum file size allowed (in bytes).
func (p *GoldmarkParser) SetMaxFileSize(size int64) {
	p.maxFileSize = size
}

// validateFile checks if the file exists, is readable, and is a file (not directory).
func (p *GoldmarkParser) validateFile(filepath string) error {
	info, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.FileNotFoundError{Path: filepath}
		}
		if os.IsPermission(err) {
			return errors.PermissionDeniedError{Path: filepath}
		}
		return fmt.Errorf("failed to stat file: %w", err)
	}

	if info.IsDir() {
		return errors.InvalidFileError{
			Path:   filepath,
			Reason: "expected a file, got directory",
		}
	}

	// Check file size
	if info.Size() > p.maxFileSize {
		return errors.FileTooLargeError{
			Path:    filepath,
			Size:    info.Size(),
			MaxSize: p.maxFileSize,
		}
	}

	// Try to open for reading to verify permissions
	file, err := os.Open(filepath)
	if err != nil {
		if os.IsPermission(err) {
			return errors.PermissionDeniedError{Path: filepath}
		}
		return fmt.Errorf("failed to open file: %w", err)
	}
	file.Close()

	return nil
}

// stripFrontmatter removes YAML frontmatter from content if present.
// Frontmatter is delimited by --- at the start of the file.
func stripFrontmatter(content []byte) []byte {
	// Check if content starts with ---
	if !bytes.HasPrefix(content, []byte("---")) {
		return content
	}

	// Find the closing ---
	lines := bytes.Split(content, []byte("\n"))
	if len(lines) < 2 {
		return content
	}

	// Find closing delimiter
	for i := 1; i < len(lines); i++ {
		if bytes.HasPrefix(bytes.TrimSpace(lines[i]), []byte("---")) {
			// Return content after frontmatter
			var result []byte
			for j := i + 1; j < len(lines); j++ {
				if len(result) > 0 {
					result = append(result, '\n')
				}
				result = append(result, lines[j]...)
			}
			return result
		}
	}

	return content
}

// Parse reads a markdown file and extracts its table of contents.
func (p *GoldmarkParser) Parse(filepath string) (models.TableOfContents, error) {
	// Validate file first
	if err := p.validateFile(filepath); err != nil {
		return models.TableOfContents{}, err
	}

	// Read file content
	content, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsPermission(err) {
			return models.TableOfContents{}, errors.PermissionDeniedError{Path: filepath}
		}
		return models.TableOfContents{}, fmt.Errorf("failed to read file: %w", err)
	}

	// Strip YAML frontmatter
	content = stripFrontmatter(content)

	// Create Goldmark parser with default extensions
	gm := goldmark.New(
		goldmark.WithExtensions(
		// Enable Setext headings (level 1 and 2 underlined headings)
		),
	)

	// Parse the markdown
	doc := gm.Parser().Parse(text.NewReader(content))

	// Extract headings using AST visitor
	var headings []models.Heading
	var inCodeBlock bool

	_ = ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			switch n.Kind() {
			case ast.KindFencedCodeBlock, ast.KindCodeBlock:
				inCodeBlock = true
			case ast.KindHeading:
				if !inCodeBlock {
					h := n.(*ast.Heading)
					lineNumber := 1
					if segment := h.Lines(); segment.Len() > 0 {
						lineNumber = segment.At(0).Start
					}

					text := extractText(content, h)
					headings = append(headings, models.Heading{
						Level:      h.Level,
						Text:       text,
						LineNumber: lineNumber,
					})
				}
			}
		} else {
			switch n.Kind() {
			case ast.KindFencedCodeBlock, ast.KindCodeBlock:
				inCodeBlock = false
			}
		}
		return ast.WalkContinue, nil
	})

	return models.TableOfContents{
		Headings: headings,
		Source:   filepath,
	}, nil
}

// extractText extracts the text content from a heading node.
func extractText(content []byte, h *ast.Heading) string {
	var result []byte
	for child := h.FirstChild(); child != nil; child = child.NextSibling() {
		if text, ok := child.(*ast.Text); ok {
			result = append(result, text.Segment.Value(content)...)
		}
	}
	return string(result)
}

// byteOffsetToLine converts a byte offset to a 1-based line number.
func byteOffsetToLine(content []byte, offset int) int {
	line := 1
	for i := 0; i < offset && i < len(content); i++ {
		if content[i] == '\n' {
			line++
		}
	}
	return line
}

// ParseSectionTree reads a markdown file and builds a hierarchical section tree with line ranges.
func (p *GoldmarkParser) ParseSectionTree(filepath string) (*models.SectionTree, error) {
	// Validate file first
	if err := p.validateFile(filepath); err != nil {
		return nil, err
	}

	// Read original file content (before stripping frontmatter)
	originalContent, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsPermission(err) {
			return nil, errors.PermissionDeniedError{Path: filepath}
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Count total lines in original file
	lines := bytes.Split(originalContent, []byte("\n"))
	totalLines := len(lines)
	if len(originalContent) > 0 && len(lines[len(lines)-1]) == 0 {
		totalLines--
	}
	if totalLines < 1 {
		totalLines = 1
	}

	// Calculate frontmatter offset
	frontmatterOffset := 0
	if bytes.HasPrefix(originalContent, []byte("---")) {
		lines := bytes.Split(originalContent, []byte("\n"))
		for i := 1; i < len(lines); i++ {
			if bytes.HasPrefix(bytes.TrimSpace(lines[i]), []byte("---")) {
				// Count lines in frontmatter including the closing ---
				frontmatterOffset = i + 1
				break
			}
		}
	}

	// Strip YAML frontmatter for parsing
	content := stripFrontmatter(originalContent)

	// Create Goldmark parser
	gm := goldmark.New()

	// Parse the markdown
	doc := gm.Parser().Parse(text.NewReader(content))

	// Extract sections with line numbers
	var sections []*models.Section
	var inCodeBlock bool

	_ = ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			switch n.Kind() {
			case ast.KindFencedCodeBlock, ast.KindCodeBlock:
				inCodeBlock = true
			case ast.KindHeading:
				if !inCodeBlock {
					h := n.(*ast.Heading)
					lineNumber := 1
					if segment := h.Lines(); segment.Len() > 0 {
						// Convert byte offset to line number
						offset := segment.At(0).Start
						lineNumber = byteOffsetToLine(content, offset) + frontmatterOffset
					}

					text := extractText(content, h)
					section := &models.Section{
						ID:        fmt.Sprintf("L%d", lineNumber),
						Level:     h.Level,
						Title:     text,
						StartLine: lineNumber,
						Children:  []*models.Section{},
					}
					sections = append(sections, section)
				}
			}
		} else {
			switch n.Kind() {
			case ast.KindFencedCodeBlock, ast.KindCodeBlock:
				inCodeBlock = false
			}
		}
		return ast.WalkContinue, nil
	})

	// Calculate end lines for each section
	for i, section := range sections {
		if i < len(sections)-1 {
			// End at the line before the next heading
			section.EndLine = sections[i+1].StartLine - 1
		} else {
			// Last section ends at end of file
			section.EndLine = totalLines
		}
	}

	// Build hierarchical tree structure
	root := &models.Section{
		ID:       "root",
		Level:    0,
		Title:    "Root",
		Children: []*models.Section{},
	}

	byID := make(map[string]*models.Section)
	for _, section := range sections {
		byID[section.ID] = section
	}

	// Build parent-child relationships
	var lastParent [7]*models.Section // Track last section at each level (0-6)
	lastParent[0] = root

	for _, section := range sections {
		// Find the appropriate parent
		parentLevel := section.Level - 1
		for parentLevel > 0 && lastParent[parentLevel] == nil {
			parentLevel--
		}

		if parent := lastParent[parentLevel]; parent != nil {
			section.Parent = parent
			parent.Children = append(parent.Children, section)
		}

		// Update this level and reset deeper levels
		lastParent[section.Level] = section
		for i := section.Level + 1; i < len(lastParent); i++ {
			lastParent[i] = nil
		}
	}

	return &models.SectionTree{
		Root:     root,
		Source:   filepath,
		Sections: sections,
		ByID:     byID,
	}, nil
}
