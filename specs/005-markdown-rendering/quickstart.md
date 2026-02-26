# Quickstart: Rich Markdown Content Rendering

## Features

The content view now renders markdown with rich formatting:

### Headings
```
# H1 Heading        → Bold, prominent color
## H2 Heading       → Bold, slightly smaller
### H3 Heading      → Bold, smaller still
#### H4-H6          → Progressively smaller
```

### Lists

**Unordered Lists:**
- Item 1     → ● Bullet
- Item 2     ● Sub-item → ○ Open bullet
  - Nested   └── 1.1. Title

**Ordered Lists:**
1. First item
2. Second item
   1. Nested item

**Task Lists:**
☐ Unchecked task
☑ Checked task

### Code

**Inline code:** `code` appears with monospace font and distinct background

**Code blocks:**
```go
func main() {
    fmt.Println("Syntax highlighted!")
}
```

### Text Formatting

- **Bold text** → Bright/bold styling
- *Italic text* → Italic or distinct color
- ~~Strikethrough~~ → Crossed out or dimmed
- ***Bold italic*** → Combined styling

### Links
[Link text](url) → Underlined, colored text

### Blockquotes
> Quoted text appears with left border or prefix

### Tables
| Column 1 | Column 2 |
|----------|----------|
| Data 1   | Data 2   |

## Using in TUI

Navigate to any section and press **Enter** to view the formatted content.

The markdown is automatically rendered with appropriate styling based on terminal capabilities.
