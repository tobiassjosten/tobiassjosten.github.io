package site

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gosimple/slug"
)

type GoodreadsBook struct {
	ID       string
	Title    string
	Author   string
	Rating   string
	DateRead string
}

func isDigitsOnly(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}

func stripSubtitle(title string) string {
	if idx := strings.Index(title, ":"); idx != -1 {
		if !isDigitsOnly(strings.TrimSpace(title[:idx])) {
			return strings.TrimSpace(title[:idx])
		}
	}

	if idx := strings.Index(title, "–"); idx != -1 {
		return strings.TrimSpace(title[:idx])
	}

	if idx := strings.Index(title, " ("); idx != -1 {
		return strings.TrimSpace(title[:idx])
	}

	return title
}

func loadExistingBooks() (map[string]bool, error) {
	existingTitles := make(map[string]bool)

	bookPaths, err := filepath.Glob("content/books/*/index.md")
	if err != nil {
		return nil, err
	}

	for _, path := range bookPaths {
		slug := filepath.Base(filepath.Dir(path))
		existingTitles[slug] = true
	}

	return existingTitles, nil
}

func loadExistingAuthors() (map[string]bool, error) {
	existingAuthors := make(map[string]bool)

	authorPaths, err := filepath.Glob("content/authors/*/index.md")
	if err != nil {
		return nil, err
	}

	for _, path := range authorPaths {
		slug := filepath.Base(filepath.Dir(path))
		existingAuthors[slug] = true
	}

	return existingAuthors, nil
}

func readGoodreadsCSV(filename string) ([]GoodreadsBook, error) {
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
		return nil, fmt.Errorf("CSV file is empty")
	}

	header := records[0]
	columnIndices := make(map[string]int)
	for i, col := range header {
		columnIndices[col] = i
	}

	var books []GoodreadsBook
	for i, record := range records[1:] {
		if len(record) <= columnIndices["Exclusive Shelf"] {
			return nil, fmt.Errorf("malformed CSV record #%d: %v", i, record)
		}

		if record[columnIndices["Exclusive Shelf"]] == "to-read" {
			continue
		}

		if record[columnIndices["My Rating"]] == "0" && record[columnIndices["Date Read"]] != "" {
			fmt.Printf(
				"Skipping unrated book: https://www.goodreads.com/book/show/%s\n",
				record[columnIndices["Book Id"]],
			)
			continue
		}

		book := GoodreadsBook{
			ID:       record[columnIndices["Book Id"]],
			Title:    stripSubtitle(record[columnIndices["Title"]]),
			Author:   record[columnIndices["Author"]],
			Rating:   record[columnIndices["My Rating"]],
			DateRead: record[columnIndices["Date Read"]],
		}

		books = append(books, book)
	}

	return books, nil
}

func createAuthor(authorName string) error {
	authorSlug := slug.Make(authorName)

	authorDir := filepath.Join("content/authors", authorSlug)

	if err := os.MkdirAll(authorDir, 0755); err != nil {
		return fmt.Errorf("creating author directory: %w", err)
	}

	indexPath := filepath.Join(authorDir, "index.md")
	content := fmt.Sprintf("---\ntitle: \"%s\"\n---\n", authorName)

	if err := os.WriteFile(indexPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("writing author index: %w", err)
	}

	return nil
}

func downloadCover(book GoodreadsBook) error {
	goodreadsURL := fmt.Sprintf("https://www.goodreads.com/book/show/%s", book.ID)

	resp, err := http.Get(goodreadsURL)
	if err != nil {
		return fmt.Errorf("failed to fetch Goodreads page: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Goodreads page returned status %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("Error creating document: %v", err)
	}

	imageURL, exists := doc.Find(".BookCover__image").First().Find("img").First().Attr("src")
	if !exists {
		return errors.New("Image URL not found on Goodreads page")
	}

	resp, err = http.Get(imageURL)
	if err != nil {
		return fmt.Errorf("downloading cover: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("cover not available, %s: %d", imageURL, resp.StatusCode)
	}

	bookSlug := slug.Make(book.Title)

	imagePath := filepath.Join("content/books", bookSlug, bookSlug+".jpg")

	imageFile, err := os.Create(imagePath)
	if err != nil {
		return fmt.Errorf("creating image file: %w", err)
	}
	defer imageFile.Close()

	if _, err := io.Copy(imageFile, resp.Body); err != nil {
		return fmt.Errorf("saving image: %w", err)
	}

	time.Sleep(2 * time.Second)

	return nil
}

func createBook(book GoodreadsBook) error {
	bookSlug := slug.Make(book.Title)
	authorSlug := slug.Make(book.Author)

	bookDir := filepath.Join("content/books", bookSlug)

	if err := os.MkdirAll(bookDir, 0755); err != nil {
		return fmt.Errorf("creating book directory: %w", err)
	}

	dateFormatted := strings.ReplaceAll(book.DateRead, "/", "-")
	mainTitle := stripSubtitle(book.Title)

	var frontmatter strings.Builder
	frontmatter.WriteString("---\n")
	frontmatter.WriteString(fmt.Sprintf("title: \"%s\"\n", mainTitle))
	frontmatter.WriteString(fmt.Sprintf("author: \"%s\"\n", authorSlug))
	frontmatter.WriteString(fmt.Sprintf("date: \"%s\"\n", dateFormatted))
	frontmatter.WriteString("amazonURL: \"\"\n")
	frontmatter.WriteString(fmt.Sprintf("image: \"%s.jpg\"\n", bookSlug))

	if book.Rating != "" && book.Rating != "0" {
		frontmatter.WriteString(fmt.Sprintf("rating: %s\n", book.Rating))
	}

	if book.DateRead == "" {
		frontmatter.WriteString("currentlyReading: true\n")
	}

	frontmatter.WriteString("---\n")

	indexPath := filepath.Join(bookDir, "index.md")
	if err := os.WriteFile(indexPath, []byte(frontmatter.String()), 0644); err != nil {
		return fmt.Errorf("writing book index: %w", err)
	}

	return nil
}

func main() {
	csvFile := "goodreads_library_export.csv"

	fmt.Println("=== Goodreads Import Tool ===\n")

	existingBooks, err := loadExistingBooks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading existing books: %v\n", err)
		os.Exit(1)
	}

	existingAuthors, err := loadExistingAuthors()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading existing authors: %v\n", err)
		os.Exit(1)
	}

	books, err := readGoodreadsCSV(csvFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading Goodreads CSV: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d books to potentially import\n", len(books))
	fmt.Printf("Existing books: %d\n", len(existingBooks))
	fmt.Printf("Existing authors: %d\n\n", len(existingAuthors))

	imported := 0
	skipped := 0
	authorsCreated := 0
	coversFailed := 0

	for i, book := range books {
		bookSlug := slug.Make(book.Title)
		authorSlug := slug.Make(book.Author)

		if existingBooks[bookSlug] {
			skipped++
			continue
		}

		fmt.Printf("[%d/%d] Importing: %s\n", i+1, len(books), book.Title)
		fmt.Printf("  Slug: %s\n", bookSlug)
		fmt.Printf("  Author: %s (%s)\n", book.Author, authorSlug)

		if !existingAuthors[authorSlug] {
			if err := createAuthor(book.Author); err != nil {
				fmt.Printf("  ✗ Failed to create author: %v\n", err)
				continue
			}

			existingAuthors[authorSlug] = true
			authorsCreated++

			fmt.Printf("  ✓ Created author: %s\n", authorSlug)
		}

		if err := createBook(book); err != nil {
			fmt.Printf("  ✗ Failed to create book: %v\n", err)
			continue
		}
		fmt.Printf("  ✓ Created book metadata\n")

		if err := downloadCover(book); err != nil {
			fmt.Printf("  ⚠ Cover download failed: %v\n", err)
			coversFailed++
		} else {
			fmt.Printf("  ✓ Downloaded cover\n")
		}

		imported++
		existingBooks[bookSlug] = true

		fmt.Println()
	}

	fmt.Printf("\n=== Import Summary ===\n")
	fmt.Printf("Books imported: %d\n", imported)
	fmt.Printf("Books skipped (already exist): %d\n", skipped)
	fmt.Printf("Authors created: %d\n", authorsCreated)
	fmt.Printf("Covers failed: %d\n", coversFailed)
}
