package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"mdexplore/internal/errors"
	"mdexplore/internal/models"
	"mdexplore/internal/parser"
	"mdexplore/internal/ui"
)

var (
	version       = "dev"
	showToc       bool
	showHelp      bool
	showVer       bool
	maxLevel      int
	selectSection string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "mdexplore <file> [flags]",
		Short: "Display table of contents from markdown files",
		Long: `mdexplore displays a hierarchical table of contents from markdown files.

It supports both ATX-style (# Heading) and Setext-style (underlined) headings.
The tool launches an interactive TUI for navigating the TOC.

Navigation Keys (TUI mode):
  ↑/↓ or k/j    Navigate up/down
  →             Expand section
  ←             Collapse section
  Enter         View section content
  Esc           Return / Quit
  q             Quit

Options:
  -L, --level N    Maximum heading level to display (0 = no limit)
      --select N   Pre-select section by number (e.g., 1.1)`,
        Example: `  mdexplore README.md
  mdexplore README.md --level 2
  mdexplore docs/guide.md -L 1`,
		Args: func(cmd *cobra.Command, args []string) error {
			// Accept 1 positional arg (the file), flags are handled separately
			if len(args) > 1 {
				return fmt.Errorf("accepts at most 1 arg, received %d", len(args))
			}
			return nil
		},
		Version: version,
		RunE:    run,
	}

	rootCmd.Flags().BoolVar(&showToc, "toc", true, "Display table of contents")
	rootCmd.Flags().BoolVarP(&showHelp, "help", "h", false, "Display help information")
	rootCmd.Flags().BoolVarP(&showVer, "version", "v", false, "Display version information")
	rootCmd.Flags().IntVarP(&maxLevel, "level", "L", 0, "Maximum heading level to display (0 = no limit)")
	rootCmd.Flags().StringVar(&selectSection, "select", "", "Pre-select section by number (e.g., 1.1)")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	// Check if file argument provided
	if len(args) == 0 {
		cmd.Usage()
		return nil
	}

	filepath := args[0]

	// Parse the markdown file into a section tree
	p := parser.NewGoldmarkParser()
	tree, err := p.ParseSectionTree(filepath)

	if err != nil {
		// Print error directly for non-TTY environments
		fmt.Fprintln(os.Stderr, err)
		setExitCode(err)
		return nil
	}

	// Apply level filter if specified
	if maxLevel > 0 {
		tree = tree.FilterByMaxLevel(maxLevel)
	}

	// Handle --select option
	var targetSection *models.Section
	if selectSection != "" {
		sectionNum, err := models.ParseSectionNumber(selectSection)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: invalid section number '%s': %v\n", selectSection, err)
			os.Exit(1)
		}

		var found bool
		targetSection, found = tree.FindByNumber(sectionNum)
		if !found {
			fmt.Fprintf(os.Stderr, "Error: section '%s' not found\n", selectSection)
			os.Exit(1)
		}
	}

	// Check if we have a TTY
	if !isTerminal() {
		// Non-TTY mode: print selected section or full TOC
		if targetSection != nil {
			printSelectedSection(tree, targetSection)
		} else {
			printSectionTree(tree)
		}
		return nil
	}

	// Create the TUI model (with optional pre-selection)
	var model ui.Model
	if targetSection != nil {
		model = ui.InitialModelWithSelection(filepath, tree, targetSection)
	} else {
		model = ui.InitialModel(filepath, tree)
	}

	// Run the TUI
	program := tea.NewProgram(model)
	if _, err := program.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running TUI: %v\n", err)
		return err
	}

	// Set appropriate exit codes based on error type
	if model.HasError() {
		setExitCode(model.Error)
	}

	return nil
}

// isTerminal checks if stdout is a terminal
func isTerminal() bool {
	fileInfo, _ := os.Stdout.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}

// printSectionTree prints the section tree to stdout in non-TTY mode
func printSectionTree(tree *models.SectionTree) {
	if tree == nil || len(tree.GetH1Sections()) == 0 {
		fmt.Println("No headings found.")
		return
	}

	numberedSections := tree.AssignNumbers()
	for _, ns := range numberedSections {
		indent := strings.Repeat("  ", ns.Section.GetDepth())
		fmt.Printf("%-6s L%d-%d %s%s\n", ns.DisplayNumber, ns.Section.StartLine, ns.Section.EndLine, indent, ns.Section.Title)
	}
}

// printSelectedSection prints only the selected section and its subsections
func printSelectedSection(tree *models.SectionTree, target *models.Section) {
	if tree == nil || target == nil {
		fmt.Println("No headings found.")
		return
	}

	numberedSections := tree.AssignNumbers()

	// Find the target's section number to identify its descendants
	var targetNum models.SectionNumber
	for _, ns := range numberedSections {
		if ns.Section.ID == target.ID {
			targetNum = ns.Number
			break
		}
	}

	for _, ns := range numberedSections {
		// Print if this is the target or a descendant of the target
		if ns.Section.ID == target.ID || targetNum.IsAncestorOf(ns.Number) {
			indent := strings.Repeat("  ", ns.Section.GetDepth()-target.GetDepth())
			fmt.Printf("%-6s L%d-%d %s%s\n", ns.DisplayNumber, ns.Section.StartLine, ns.Section.EndLine, indent, ns.Section.Title)
		}
	}
}

// printTOC prints the table of contents to stdout in non-TTY mode (legacy)
func printTOC(toc models.TableOfContents) {
	if toc.IsEmpty() {
		fmt.Println("No headings found.")
		return
	}

	for _, heading := range toc.Headings {
		indent := strings.Repeat("  ", heading.Level-1)
		fmt.Printf("%s%s\n", indent, heading.Text)
	}
}

func setExitCode(err error) {
	switch err.(type) {
	case errors.FileNotFoundError:
		os.Exit(2)
	case errors.PermissionDeniedError:
		os.Exit(3)
	case errors.InvalidFileError:
		os.Exit(4)
	case errors.ParseError:
		os.Exit(5)
	default:
		os.Exit(1)
	}
}
