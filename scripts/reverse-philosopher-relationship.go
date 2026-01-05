package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type BookInfo struct {
	Path         string
	Slug         string
	Philosophers []string
}

type PhilosopherInfo struct {
	Path  string
	Slug  string
	Books []string
}

func main() {
	fmt.Println("=== Reverse Book-Philosopher Relationship ===")
	fmt.Println()

	// Step 1: Find all books with philosophers property
	fmt.Println("Step 1: Finding books with philosophers property...")
	books, err := findBooksWithPhilosophers()
	if err != nil {
		fmt.Printf("Error finding books: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d books with philosophers property\n", len(books))
	fmt.Println()

	// Step 2: Build reverse mapping
	fmt.Println("Step 2: Building philosopher -> books mapping...")
	philosopherBooks := buildPhilosopherMapping(books)

	fmt.Printf("Found %d philosophers with books\n", len(philosopherBooks))
	fmt.Println()

	// Step 3: Display mapping for review
	fmt.Println("Step 3: Philosopher -> Books mapping:")
	fmt.Println()
	displayMapping(philosopherBooks)

	// Step 4: Update philosopher files
	fmt.Println("\nStep 4: Updating philosopher files...")
	err = updatePhilosopherFiles(philosopherBooks)
	if err != nil {
		fmt.Printf("Error updating philosophers: %v\n", err)
		os.Exit(1)
	}

	// Step 5: Remove philosophers property from books
	fmt.Println("\nStep 5: Removing philosophers property from books...")
	err = removePhilosophersFromBooks(books)
	if err != nil {
		fmt.Printf("Error updating books: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n=== Migration Complete ===")
	fmt.Printf("Updated %d philosopher files\n", len(philosopherBooks))
	fmt.Printf("Updated %d book files\n", len(books))
}

func findBooksWithPhilosophers() ([]BookInfo, error) {
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
		if !strings.Contains(contentStr, "philosophers:") {
			continue
		}

		// Extract philosophers list
		philosophers := extractPhilosophers(contentStr)
		if len(philosophers) == 0 {
			continue
		}

		slug := filepath.Base(filepath.Dir(path))
		books = append(books, BookInfo{
			Path:         path,
			Slug:         slug,
			Philosophers: philosophers,
		})
	}

	return books, nil
}

func extractPhilosophers(content string) []string {
	var philosophers []string

	lines := strings.Split(content, "\n")
	inPhilosophers := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "philosophers:") {
			inPhilosophers = true
			continue
		}

		if inPhilosophers {
			if strings.HasPrefix(trimmed, "-") {
				// Extract philosopher slug
				philosopher := strings.TrimSpace(strings.TrimPrefix(trimmed, "-"))
				philosopher = strings.Trim(philosopher, "\"")
				if philosopher != "" {
					philosophers = append(philosophers, philosopher)
				}
			} else if strings.Contains(trimmed, ":") && !strings.HasPrefix(trimmed, "#") {
				// New property, stop
				break
			}
		}
	}

	return philosophers
}

func buildPhilosopherMapping(books []BookInfo) map[string][]string {
	mapping := make(map[string][]string)

	for _, book := range books {
		for _, philosopher := range book.Philosophers {
			mapping[philosopher] = append(mapping[philosopher], book.Slug)
		}
	}

	// Sort book lists for consistency
	for philosopher := range mapping {
		sort.Strings(mapping[philosopher])
	}

	return mapping
}

func displayMapping(mapping map[string][]string) {
	philosophers := make([]string, 0, len(mapping))
	for philosopher := range mapping {
		philosophers = append(philosophers, philosopher)
	}
	sort.Strings(philosophers)

	for _, philosopher := range philosophers {
		books := mapping[philosopher]
		fmt.Printf("  %s:\n", philosopher)
		for _, book := range books {
			fmt.Printf("    - %s\n", book)
		}
		fmt.Println()
	}
}

func updatePhilosopherFiles(mapping map[string][]string) error {
	for philosopher, books := range mapping {
		path := filepath.Join("content/philosophers", philosopher, "index.md")

		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("  ⚠️  Warning: Philosopher file not found: %s\n", path)
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

		fmt.Printf("  ✅ Updated %s (%d books)\n", philosopher, len(books))
	}

	return nil
}

func removePhilosophersFromBooks(books []BookInfo) error {
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

		// Remove philosophers property
		newFrontmatter := removePhilosophersProperty(frontmatter)

		// Reconstruct file
		newContent := "---\n" + newFrontmatter + "---" + body

		err = os.WriteFile(book.Path, []byte(newContent), 0644)
		if err != nil {
			return fmt.Errorf("writing %s: %w", book.Path, err)
		}

		fmt.Printf("  ✅ Removed philosophers from %s\n", book.Slug)
	}

	return nil
}

func removePhilosophersProperty(frontmatter string) string {
	lines := strings.Split(frontmatter, "\n")
	var newLines []string
	inPhilosophers := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "philosophers:") {
			inPhilosophers = true
			continue
		}

		if inPhilosophers {
			if strings.HasPrefix(trimmed, "-") {
				// Skip philosopher list items
				continue
			} else if strings.Contains(trimmed, ":") || trimmed == "" {
				// End of philosophers list
				inPhilosophers = false
			}
		}

		if !inPhilosophers {
			newLines = append(newLines, line)
		}
	}

	return strings.Join(newLines, "\n")
}
