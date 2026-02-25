package models

// Heading represents a single heading extracted from a markdown document.
type Heading struct {
	Level      int    // 1-6 (H1-H6)
	Text       string // Clean heading text (without markdown markers)
	LineNumber int    // Line number in source file (1-based)
}

// TableOfContents represents the complete structure of headings in a document.
type TableOfContents struct {
	Headings []Heading // Ordered list of headings as they appear in document
	Source   string    // Path to source file
}

// IsEmpty returns true if the TOC contains no headings.
func (t TableOfContents) IsEmpty() bool {
	return len(t.Headings) == 0
}

// Count returns the total number of headings.
func (t TableOfContents) Count() int {
	return len(t.Headings)
}

// HeadingsByLevel returns all headings at a specific level.
func (t TableOfContents) HeadingsByLevel(level int) []Heading {
	var result []Heading
	for _, h := range t.Headings {
		if h.Level == level {
			result = append(result, h)
		}
	}
	return result
}
