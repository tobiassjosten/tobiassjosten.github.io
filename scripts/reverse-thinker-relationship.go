package scripts

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type BookInfo struct {
	Path     string
	Slug     string
	Thinkers []string
}

type ThinkerInfo struct {
	Path  string
	Slug  string
	Books []string
}

// ReverseThinkerRelationship is a one-off migration: it reads the thinkers
// listed on each book, writes the reverse (books per thinker) into the thinker
// files, and removes the thinkers property from the books.
func ReverseThinkerRelationship(args []string) error {
	fmt.Println("=== Reverse Book-Thinker Relationship ===")
	fmt.Println()

	// Step 1: Find all books with thinkers property
	fmt.Println("Step 1: Finding books with thinkers property...")
	books, err := findBooksWithThinkers()
	if err != nil {
		return fmt.Errorf("finding books: %w", err)
	}

	fmt.Printf("Found %d books with thinkers property\n", len(books))
	fmt.Println()

	// Step 2: Build reverse mapping
	fmt.Println("Step 2: Building thinker -> books mapping...")
	thinkerBooks := buildThinkerMapping(books)

	fmt.Printf("Found %d thinkers with books\n", len(thinkerBooks))
	fmt.Println()

	// Step 3: Display mapping for review
	fmt.Println("Step 3: Thinker -> Books mapping:")
	fmt.Println()
	displayMapping(thinkerBooks)

	// Step 4: Update thinker files
	fmt.Println("\nStep 4: Updating thinker files...")
	if err := updateThinkerFiles(thinkerBooks); err != nil {
		return fmt.Errorf("updating thinkers: %w", err)
	}

	// Step 5: Remove thinkers property from books
	fmt.Println("\nStep 5: Removing thinkers property from books...")
	if err := removeThinkersFromBooks(books); err != nil {
		return fmt.Errorf("updating books: %w", err)
	}

	fmt.Println("\n=== Migration Complete ===")
	fmt.Printf("Updated %d thinker files\n", len(thinkerBooks))
	fmt.Printf("Updated %d book files\n", len(books))

	return nil
}

func findBooksWithThinkers() ([]BookInfo, error) {
	var books []BookInfo

	bookPaths, err := filepath.Glob("content/books/*/index.md")
	if err != nil {
		return nil, err
	}

	for _, path := range bookPaths {
		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		contentStr := string(content)
		if !strings.Contains(contentStr, "thinkers:") {
			continue
		}

		// Extract thinkers list
		thinkers := extractThinkers(contentStr)
		if len(thinkers) == 0 {
			continue
		}

		slug := filepath.Base(filepath.Dir(path))
		books = append(books, BookInfo{
			Path:     path,
			Slug:     slug,
			Thinkers: thinkers,
		})
	}

	return books, nil
}

func extractThinkers(content string) []string {
	var thinkers []string

	lines := strings.Split(content, "\n")
	inThinkers := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "thinkers:") {
			inThinkers = true
			continue
		}

		if inThinkers {
			if strings.HasPrefix(trimmed, "-") {
				// Extract thinker slug
				thinker := strings.TrimSpace(strings.TrimPrefix(trimmed, "-"))
				thinker = strings.Trim(thinker, "\"")
				if thinker != "" {
					thinkers = append(thinkers, thinker)
				}
			} else if strings.Contains(trimmed, ":") && !strings.HasPrefix(trimmed, "#") {
				// New property, stop
				break
			}
		}
	}

	return thinkers
}

func buildThinkerMapping(books []BookInfo) map[string][]string {
	mapping := make(map[string][]string)

	for _, book := range books {
		for _, thinker := range book.Thinkers {
			mapping[thinker] = append(mapping[thinker], book.Slug)
		}
	}

	// Sort book lists for consistency
	for thinker := range mapping {
		sort.Strings(mapping[thinker])
	}

	return mapping
}

func displayMapping(mapping map[string][]string) {
	thinkers := make([]string, 0, len(mapping))
	for thinker := range mapping {
		thinkers = append(thinkers, thinker)
	}
	sort.Strings(thinkers)

	for _, thinker := range thinkers {
		books := mapping[thinker]
		fmt.Printf("  %s:\n", thinker)
		for _, book := range books {
			fmt.Printf("    - %s\n", book)
		}
		fmt.Println()
	}
}

func updateThinkerFiles(mapping map[string][]string) error {
	for thinker, books := range mapping {
		path := filepath.Join("content/thinkers", thinker, "index.md")

		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("  ⚠️  Warning: Thinker file not found: %s\n", path)
			continue
		}

		contentStr := string(content)

		// Split frontmatter and body
		parts := strings.SplitN(contentStr, "---", 3)
		if len(parts) < 3 {
			fmt.Printf("  ⚠️  Warning: Invalid frontmatter in %s\n", path)
			continue
		}

		frontmatter := parts[1]
		body := parts[2]

		// Add books property to frontmatter
		booksYAML := "books:\n"
		for _, book := range books {
			booksYAML += fmt.Sprintf("  - %s\n", book)
		}

		// Insert books property before closing ---
		newFrontmatter := strings.TrimSpace(frontmatter) + "\n" + booksYAML

		// Reconstruct file
		newContent := "---\n" + newFrontmatter + "---" + body

		err = os.WriteFile(path, []byte(newContent), 0644)
		if err != nil {
			return fmt.Errorf("writing %s: %w", path, err)
		}

		fmt.Printf("  ✅ Updated %s (%d books)\n", thinker, len(books))
	}

	return nil
}

func removeThinkersFromBooks(books []BookInfo) error {
	for _, book := range books {
		content, err := os.ReadFile(book.Path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", book.Path, err)
		}

		contentStr := string(content)

		// Split frontmatter and body
		parts := strings.SplitN(contentStr, "---", 3)
		if len(parts) < 3 {
			fmt.Printf("  ⚠️  Warning: Invalid frontmatter in %s\n", book.Path)
			continue
		}

		frontmatter := parts[1]
		body := parts[2]

		// Remove thinkers property
		newFrontmatter := removeThinkersProperty(frontmatter)

		// Reconstruct file
		newContent := "---\n" + newFrontmatter + "---" + body

		err = os.WriteFile(book.Path, []byte(newContent), 0644)
		if err != nil {
			return fmt.Errorf("writing %s: %w", book.Path, err)
		}

		fmt.Printf("  ✅ Removed thinkers from %s\n", book.Slug)
	}

	return nil
}

func removeThinkersProperty(frontmatter string) string {
	lines := strings.Split(frontmatter, "\n")
	var newLines []string
	inThinkers := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "thinkers:") {
			inThinkers = true
			continue
		}

		if inThinkers {
			if strings.HasPrefix(trimmed, "-") {
				// Skip thinker list items
				continue
			} else if strings.Contains(trimmed, ":") || trimmed == "" {
				// End of thinkers list
				inThinkers = false
			}
		}

		if !inThinkers {
			newLines = append(newLines, line)
		}
	}

	return strings.Join(newLines, "\n")
}
