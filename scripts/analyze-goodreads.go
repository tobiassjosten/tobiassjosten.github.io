package scripts

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func parseFrontMatter(content string) (map[string]string, error) {
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid front matter format")
	}

	fm := make(map[string]string)
	lines := strings.Split(strings.TrimSpace(parts[1]), "\n")

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || !strings.Contains(trimmed, ":") {
			continue
		}

		parts := strings.SplitN(trimmed, ":", 2)
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		fm[key] = strings.Trim(value, "\"")
	}

	return fm, nil
}

func loadExistingBookTitles() (map[string]bool, error) {
	existingTitles := make(map[string]bool)

	bookPaths, err := filepath.Glob("content/books/*/index.md")
	if err != nil {
		return nil, fmt.Errorf("finding book files: %w", err)
	}

	for _, path := range bookPaths {
		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		fm, err := parseFrontMatter(string(content))
		if err != nil {
			continue
		}

		if title, exists := fm["title"]; exists {
			normalizedTitle := strings.ToLower(strings.TrimSpace(title))
			existingTitles[normalizedTitle] = true
		}
	}

	return existingTitles, nil
}

func readAnalysisCSV(filename string) ([]GoodreadsBook, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("reading CSV: %w", err)
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("CSV file is empty or has no data rows")
	}

	header := records[0]
	columnIndices := make(map[string]int)
	for i, col := range header {
		columnIndices[col] = i
	}

	requiredColumns := []string{"Book Id", "Title", "Author", "Exclusive Shelf"}
	for _, col := range requiredColumns {
		if _, exists := columnIndices[col]; !exists {
			return nil, fmt.Errorf("missing required column: %s", col)
		}
	}

	var books []GoodreadsBook
	for _, record := range records[1:] {
		if len(record) <= columnIndices["Exclusive Shelf"] {
			continue
		}

		exclusiveShelf := record[columnIndices["Exclusive Shelf"]]
		if exclusiveShelf == "to-read" {
			continue
		}

		book := GoodreadsBook{
			ID:          record[columnIndices["Book Id"]],
			Title:       record[columnIndices["Title"]],
			Author:      record[columnIndices["Author"]],
			Rating:      record[columnIndices["My Rating"]],
			DateRead:    record[columnIndices["Date Read"]],
			Bookshelves: record[columnIndices["Bookshelves"]],
		}

		book.Bookshelves = strings.ReplaceAll(book.Bookshelves, "/", "-")

		books = append(books, book)
	}

	return books, nil
}

func formatRating(rating string) string {
	if rating == "" || rating == "0" {
		return "Not rated"
	}
	return rating
}

func bookExists(title string, existingBooks map[string]bool) bool {
	normalizedTitle := strings.ToLower(strings.TrimSpace(title))

	if existingBooks[normalizedTitle] {
		return true
	}

	mainTitle := strings.Split(normalizedTitle, ":")[0]
	mainTitle = strings.TrimSpace(mainTitle)

	for existingTitle := range existingBooks {
		if existingTitle == mainTitle {
			return true
		}

		if strings.HasPrefix(normalizedTitle, existingTitle+":") {
			return true
		}
	}

	return false
}

// AnalyzeGoodreads lists books from a Goodreads export that are not yet present
// under content/books/. The CSV path defaults to goodreads_library_export.csv
// and can be overridden by the first argument.
func AnalyzeGoodreads(args []string) error {
	csvFile := "goodreads_library_export.csv"
	if len(args) > 0 {
		csvFile = args[0]
	}

	existingBooks, err := loadExistingBookTitles()
	if err != nil {
		return fmt.Errorf("loading existing books: %w", err)
	}

	goodreadsBooks, err := readAnalysisCSV(csvFile)
	if err != nil {
		return fmt.Errorf("reading Goodreads CSV: %w", err)
	}

	fmt.Printf("=== Books from Goodreads Not Yet in content/books/ ===\n\n")
	fmt.Printf("Loaded %d existing books\n", len(existingBooks))
	fmt.Printf("Found %d books in Goodreads (excluding to-read)\n\n", len(goodreadsBooks))

	missingCount := 0
	filteredCount := 0
	for _, book := range goodreadsBooks {
		if bookExists(book.Title, existingBooks) {
			continue
		}

		if book.Bookshelves == "" || book.DateRead == "" {
			filteredCount++
			continue
		}

		fmt.Printf(
			`https://www.goodreads.com/book/show/%s, Title: "%s", Author: "%s", Rating: "%s", Date Read: "%s", Bookshelves: "%s"`+"\n",
			book.ID, book.Title, book.Author, formatRating(book.Rating), book.DateRead, book.Bookshelves,
		)
		missingCount++
	}

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Total books in Goodreads (excluding to-read): %d\n", len(goodreadsBooks))
	fmt.Printf("Books already in content/books/: %d\n", len(goodreadsBooks)-missingCount-filteredCount)
	fmt.Printf("Books filtered (empty Bookshelves or Date Read): %d\n", filteredCount)
	fmt.Printf("Books missing from content/books/: %d\n", missingCount)

	return nil
}
