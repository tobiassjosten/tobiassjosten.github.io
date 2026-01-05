package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var allowedProperties = map[string]bool{
	"title":              true,
	"authors":            true,
	"date":               true,
	"amazonURL":          true,
	"image":              true,
	"relatedBooks":       true,
	"rating":             true,
	"currentlyReading":   true,
	"featuredOnHomepage": true,
}

var allowedArticleProperties = map[string]bool{
	"title": true,
	"date":  true,
}

var requiredArticleProperties = []string{
	"title",
	"date",
}

// BookMetadata represents parsed book frontmatter
type BookMetadata struct {
	Path               string
	Slug               string
	AllProperties      map[string]string // raw properties from frontmatter
	Title              string
	Authors            []string
	Date               string
	AmazonURL          string
	Image              string
	RelatedBooksRaw    string // raw value to detect bracket syntax
	RelatedBooks       []string
	Rating             string
	CurrentlyReading   bool
	FeaturedOnHomepage bool
	Content            string // content body after frontmatter
	HasContent         bool   // whether book has non-empty content
}

// ArticleMetadata represents parsed article content
type ArticleMetadata struct {
	Path        string
	Slug        string
	Properties  map[string]string // raw properties from frontmatter
	Title       string
	Date        string
	Draft       bool
	Content     string
	HasMoreMark bool
}

// PhilosopherMetadata represents parsed philosopher frontmatter
type PhilosopherMetadata struct {
	Path  string
	Slug  string
	Title string
	Books []string
}

// ValidationError represents a validation error
type ValidationError struct {
	BookSlug  string
	FilePath  string
	ErrorType string
	Message   string
}

// ValidationWarning represents a validation warning
type ValidationWarning struct {
	BookSlug string
	FilePath string
	Message  string
}

// ValidationContext holds all data needed for validation
type ValidationContext struct {
	Books        map[string]*BookMetadata        // slug -> metadata
	Articles     map[string]*ArticleMetadata     // slug -> metadata
	Philosophers map[string]*PhilosopherMetadata // slug -> metadata
	Authors      map[string]bool                 // author slug -> exists
	Errors       []ValidationError
	Warnings     []ValidationWarning
}

// parseFrontMatter parses YAML frontmatter from book content
// Adapted from refine_book_frontmatter.go
func parseFrontMatter(content string) (map[string]string, error) {
	// Split front matter from content
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid front matter format")
	}

	// Parse front matter into key-value pairs
	fm := make(map[string]string)
	lines := strings.Split(strings.TrimSpace(parts[1]), "\n")

	var currentKey string
	var inMultilineValue bool
	var multilineValue []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		// Check if this is a continuation of a multiline value
		if inMultilineValue {
			if strings.HasPrefix(trimmed, "-") {
				// Part of array - remove prefix, trim whitespace and quotes
				item := strings.TrimSpace(strings.TrimPrefix(trimmed, "-"))
				item = strings.Trim(item, "\"")
				multilineValue = append(multilineValue, item)
			} else if strings.Contains(trimmed, ":") {
				// New key, save previous multiline value
				fm[currentKey] = strings.Join(multilineValue, ",")
				inMultilineValue = false
				multilineValue = nil
				// Process new key below
			}
		}

		if !inMultilineValue && strings.Contains(trimmed, ":") {
			parts := strings.SplitN(trimmed, ":", 2)
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			// Handle arrays
			if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
				// Inline array like [a, b, c] - store as-is to detect bracket syntax
				fm[key] = value
			} else if value == "" {
				// Start of multiline value (like authors:, relatedBooks:)
				currentKey = key
				inMultilineValue = true
				multilineValue = []string{}
			} else if value == "null" {
				// Null value
				fm[key] = value
			} else {
				// Regular key-value
				fm[key] = strings.Trim(value, "\"")
			}
		}
	}

	// Save last multiline value if any
	if inMultilineValue {
		fm[currentKey] = strings.Join(multilineValue, ",")
	}

	return fm, nil
}

// parseBookFile parses a single book file
func parseBookFile(path string) (*BookMetadata, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	contentStr := string(content)

	fm, err := parseFrontMatter(contentStr)
	if err != nil {
		return nil, fmt.Errorf("parsing frontmatter: %w", err)
	}

	// Extract content body after frontmatter
	parts := strings.SplitN(contentStr, "---", 3)
	bodyContent := ""
	hasContent := false
	if len(parts) >= 3 {
		bodyContent = strings.TrimSpace(parts[2])
		hasContent = len(bodyContent) > 0
	}

	slug := filepath.Base(filepath.Dir(path))

	book := &BookMetadata{
		Path:          path,
		Slug:          slug,
		AllProperties: fm,
		Title:         fm["title"],
		Date:          fm["date"],
		AmazonURL:     fm["amazonURL"],
		Image:         fm["image"],
		Rating:        fm["rating"],
		Content:       bodyContent,
		HasContent:    hasContent,
	}

	// Parse authors array
	if authorsRaw, exists := fm["authors"]; exists && authorsRaw != "" {
		// Multiline format - split by comma
		for _, a := range strings.Split(authorsRaw, ",") {
			a = strings.TrimSpace(a)
			if a != "" {
				book.Authors = append(book.Authors, a)
			}
		}
	}

	// Parse relatedBooks
	if relatedBooksRaw, exists := fm["relatedBooks"]; exists && relatedBooksRaw != "" {
		book.RelatedBooksRaw = relatedBooksRaw
		// Check for bracket syntax
		if !strings.HasPrefix(relatedBooksRaw, "[") {
			// Multiline format - split by comma
			for _, b := range strings.Split(relatedBooksRaw, ",") {
				b = strings.TrimSpace(b)
				if b != "" {
					book.RelatedBooks = append(book.RelatedBooks, b)
				}
			}
		}
		// If bracket syntax, RelatedBooks remains empty (will be caught by validation)
	}

	// Parse boolean flags
	if cr, exists := fm["currentlyReading"]; exists && cr == "true" {
		book.CurrentlyReading = true
	}
	if fh, exists := fm["featuredOnHomepage"]; exists && fh == "true" {
		book.FeaturedOnHomepage = true
	}

	return book, nil
}

// parsePhilosopherFile parses a single philosopher file
func parsePhilosopherFile(path string) (*PhilosopherMetadata, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	contentStr := string(content)

	fm, err := parseFrontMatter(contentStr)
	if err != nil {
		return nil, fmt.Errorf("parsing frontmatter: %w", err)
	}

	slug := filepath.Base(filepath.Dir(path))
	title := fm["title"]

	// Parse books array
	var books []string
	if booksRaw, exists := fm["books"]; exists && booksRaw != "" {
		for _, b := range strings.Split(booksRaw, ",") {
			b = strings.TrimSpace(b)
			if b != "" {
				books = append(books, b)
			}
		}
	}

	philosopher := &PhilosopherMetadata{
		Path:  path,
		Slug:  slug,
		Title: title,
		Books: books,
	}

	return philosopher, nil
}

// parseArticleFile parses a single article file
func parseArticleFile(path string) (*ArticleMetadata, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	contentStr := string(content)

	// Split front matter from content
	parts := strings.SplitN(contentStr, "---", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid front matter format")
	}

	// Parse all properties from frontmatter
	properties, _ := parseFrontMatter(contentStr)

	// Get slug from filename
	slug := strings.TrimSuffix(filepath.Base(path), ".md")

	// Check for <!--more--> in content body
	bodyContent := parts[2]
	hasMoreMark := strings.Contains(bodyContent, "<!--more-->")

	article := &ArticleMetadata{
		Path:        path,
		Slug:        slug,
		Properties:  properties,
		Title:       properties["title"],
		Date:        properties["date"],
		Draft:       properties["draft"] == "true",
		Content:     strings.TrimSpace(bodyContent),
		HasMoreMark: hasMoreMark,
	}

	return article, nil
}

// loadValidationContext loads all books, articles, philosophers, and authors
func loadValidationContext() (*ValidationContext, error) {
	ctx := &ValidationContext{
		Books:        make(map[string]*BookMetadata),
		Articles:     make(map[string]*ArticleMetadata),
		Philosophers: make(map[string]*PhilosopherMetadata),
		Authors:      make(map[string]bool),
	}

	bookPaths, err := filepath.Glob("content/books/*/index.md")
	if err != nil {
		return nil, fmt.Errorf("finding book files: %w", err)
	}

	for _, path := range bookPaths {
		book, err := parseBookFile(path)
		if err != nil {
			ctx.Errors = append(ctx.Errors, ValidationError{
				BookSlug:  filepath.Base(filepath.Dir(path)),
				FilePath:  path,
				ErrorType: "PARSE_ERROR",
				Message:   fmt.Sprintf("Failed to parse: %v", err),
			})
			continue
		}
		ctx.Books[book.Slug] = book
	}

	articlePaths, err := filepath.Glob("content/articles/**.md")
	if err != nil {
		return nil, fmt.Errorf("finding article files: %w", err)
	}

	for _, path := range articlePaths {
		if filepath.Base(path) == "_index.md" {
			continue
		}

		article, err := parseArticleFile(path)
		if err != nil {
			ctx.Errors = append(ctx.Errors, ValidationError{
				BookSlug:  strings.TrimSuffix(filepath.Base(path), ".md"),
				FilePath:  path,
				ErrorType: "PARSE_ERROR",
				Message:   fmt.Sprintf("Failed to parse: %v", err),
			})
			continue
		}
		ctx.Articles[article.Slug] = article
	}

	philosopherPaths, err := filepath.Glob("content/philosophers/*/index.md")
	if err != nil {
		return nil, fmt.Errorf("finding philosopher files: %w", err)
	}

	for _, path := range philosopherPaths {
		philosopher, err := parsePhilosopherFile(path)
		if err != nil {
			// Add parsing error but continue
			slug := filepath.Base(filepath.Dir(path))
			ctx.Errors = append(ctx.Errors, ValidationError{
				BookSlug:  slug,
				FilePath:  path,
				ErrorType: "PARSE_ERROR",
				Message:   fmt.Sprintf("Failed to parse: %v", err),
			})
			continue
		}
		ctx.Philosophers[philosopher.Slug] = philosopher
	}

	// Load all authors
	authorDirs, err := filepath.Glob("content/authors/*")
	if err != nil {
		return nil, fmt.Errorf("finding author directories: %w", err)
	}

	for _, dir := range authorDirs {
		// Check if it's a directory
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			slug := filepath.Base(dir)
			ctx.Authors[slug] = true
		}
	}

	return ctx, nil
}

// addError adds a validation error
func (ctx *ValidationContext) addError(book *BookMetadata, errorType, message string) {
	ctx.Errors = append(ctx.Errors, ValidationError{
		BookSlug:  book.Slug,
		FilePath:  book.Path,
		ErrorType: errorType,
		Message:   message,
	})
}

// addWarning adds a validation warning
func (ctx *ValidationContext) addWarning(book *BookMetadata, message string) {
	ctx.Warnings = append(ctx.Warnings, ValidationWarning{
		BookSlug: book.Slug,
		FilePath: book.Path,
		Message:  message,
	})
}

// validateAllowedProperties checks that only allowed properties exist
func validateAllowedProperties(ctx *ValidationContext) {
	for _, book := range ctx.Books {
		for key := range book.AllProperties {
			if !allowedProperties[key] {
				ctx.addError(book, "INVALID_PROPERTY",
					fmt.Sprintf("Property '%s' is not allowed", key))
			}
		}
	}
}

// validateAuthors checks that authors exist
func validateAuthors(ctx *ValidationContext) {
	for _, book := range ctx.Books {
		if len(book.Authors) == 0 {
			ctx.addError(book, "MISSING_AUTHOR", "Missing required 'author' or 'authors' field")
			continue
		}

		for _, authorSlug := range book.Authors {
			if !ctx.Authors[authorSlug] {
				ctx.addError(book, "INVALID_AUTHOR",
					fmt.Sprintf("Author '%s' does not exist in content/authors/", authorSlug))
			}
		}
	}
}

// validateUniqueTitles checks that all titles are unique
func validateUniqueTitles(ctx *ValidationContext) {
	titleMap := make(map[string][]string) // title -> []slugs

	for _, book := range ctx.Books {
		if book.Title == "" {
			ctx.addError(book, "MISSING_TITLE", "Missing required 'title' field")
			continue
		}
		titleMap[book.Title] = append(titleMap[book.Title], book.Slug)
	}

	for title, slugs := range titleMap {
		if len(slugs) > 1 {
			for _, slug := range slugs {
				book := ctx.Books[slug]
				ctx.addError(book, "DUPLICATE_TITLE",
					fmt.Sprintf("Title '%s' is used by multiple books: %s",
						title, strings.Join(slugs, ", ")))
			}
		}
	}
}

// validateUniqueAmazonURLs checks that all Amazon URLs are unique
// and that books with content have an Amazon URL
func validateUniqueAmazonURLs(ctx *ValidationContext) {
	urlMap := make(map[string][]string) // url -> []slugs

	for _, book := range ctx.Books {
		if book.AmazonURL == "" {
			// Only require amazonURL if book has actual content
			if book.HasContent {
				ctx.addError(book, "MISSING_AMAZON_URL",
					"Missing required 'amazonURL' field for book with content")
			}
			continue
		}

		urlMap[book.AmazonURL] = append(urlMap[book.AmazonURL], book.Slug)
	}

	for url, slugs := range urlMap {
		if len(slugs) > 1 {
			for _, slug := range slugs {
				book := ctx.Books[slug]
				ctx.addError(book, "DUPLICATE_AMAZON_URL",
					fmt.Sprintf("Amazon URL '%s' is used by multiple books: %s",
						url, strings.Join(slugs, ", ")))
			}
		}
	}
}

// validateImageFiles checks that image files exist
func validateImageFiles(ctx *ValidationContext) {
	for _, book := range ctx.Books {
		if book.Image == "" {
			ctx.addError(book, "MISSING_IMAGE", "Missing required 'image' field")
			continue
		}

		bookDir := filepath.Dir(book.Path)
		imagePath := filepath.Join(bookDir, book.Image)

		info, err := os.Stat(imagePath)
		if os.IsNotExist(err) {
			ctx.addError(book, "MISSING_IMAGE_FILE",
				fmt.Sprintf("Image file '%s' not found in book directory", book.Image))
			continue
		}

		if info.Size() < 100 {
			ctx.addError(book, "IMAGE_FILE_TOO_SMALL",
				fmt.Sprintf("Image file '%s' is suspiciously small (%d bytes)", book.Image, info.Size()))
		}
	}
}

// validateReadingStatus checks rating/date vs currentlyReading logic
func validateReadingStatus(ctx *ValidationContext) {
	datePattern := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

	for _, book := range ctx.Books {
		if book.CurrentlyReading {
			// If currentlyReading, date and rating can be null/empty
			continue
		}

		// Not currently reading - must have both rating and valid date
		hasValidDate := book.Date != "" && book.Date != "null" && datePattern.MatchString(book.Date)
		hasValidRating := false

		if book.Rating != "" && book.Rating != "null" {
			if rating, err := strconv.Atoi(book.Rating); err == nil && rating >= 1 && rating <= 5 {
				hasValidRating = true
			} else {
				ctx.addError(book, "INVALID_RATING",
					fmt.Sprintf("Rating must be an integer between 1-5, got '%s'", book.Rating))
			}
		}

		if !hasValidDate {
			ctx.addError(book, "INVALID_READING_STATUS",
				"Book is not currentlyReading but missing valid date (YYYY-MM-DD format)")
		}
		if !hasValidRating {
			ctx.addError(book, "INVALID_READING_STATUS",
				"Book is not currentlyReading but missing rating (1-5)")
		}
	}
}

// validateBooleanFlags checks that boolean flags are only true
func validateBooleanFlags(ctx *ValidationContext) {
	for _, book := range ctx.Books {
		// Check for explicit false values (should have been removed)
		if val, exists := book.AllProperties["currentlyReading"]; exists && val == "false" {
			ctx.addError(book, "INVALID_BOOLEAN",
				"Property 'currentlyReading' should not be explicitly set to false (omit instead)")
		}
		if val, exists := book.AllProperties["featuredOnHomepage"]; exists && val == "false" {
			ctx.addError(book, "INVALID_BOOLEAN",
				"Property 'featuredOnHomepage' should not be explicitly set to false (omit instead)")
		}
	}
}

// validateRelatedBooks checks relatedBooks format and references
func validateRelatedBooks(ctx *ValidationContext) {
	for _, book := range ctx.Books {
		// Check for bracket syntax
		if book.RelatedBooksRaw != "" && strings.HasPrefix(book.RelatedBooksRaw, "[") {
			ctx.addError(book, "INVALID_RELATED_BOOKS_FORMAT",
				"relatedBooks uses bracket syntax []. Use multiline YAML format instead (one per line with '  - ')")
			continue
		}

		// Check that referenced books exist
		for _, relatedSlug := range book.RelatedBooks {
			if _, exists := ctx.Books[relatedSlug]; !exists {
				ctx.addError(book, "INVALID_RELATED_BOOK",
					fmt.Sprintf("Related book '%s' does not exist", relatedSlug))
			}
		}

		// Warn if related books don't refer back (bidirectional check)
		for _, relatedSlug := range book.RelatedBooks {
			relatedBook, exists := ctx.Books[relatedSlug]
			if !exists {
				continue // Already reported as error above
			}

			// Check if relatedBook refers back to this book
			hasBackReference := false
			for _, backSlug := range relatedBook.RelatedBooks {
				if backSlug == book.Slug {
					hasBackReference = true
					break
				}
			}

			if !hasBackReference {
				ctx.addWarning(book,
					fmt.Sprintf("Related book '%s' does not reference back to '%s'",
						relatedSlug, book.Slug))
			}
		}
	}
}

// validateUnusedAuthors checks for authors with no books referencing them
func validateUnusedAuthors(ctx *ValidationContext) {
	referencedAuthors := make(map[string]bool)

	for _, book := range ctx.Books {
		for _, authorSlug := range book.Authors {
			referencedAuthors[authorSlug] = true
		}
	}

	for authorSlug := range ctx.Authors {
		if !referencedAuthors[authorSlug] {
			authorPath := filepath.Join("content/authors", authorSlug)
			ctx.Errors = append(ctx.Errors, ValidationError{
				BookSlug:  authorSlug,
				FilePath:  authorPath,
				ErrorType: "UNUSED_AUTHOR",
				Message:   fmt.Sprintf("Author '%s' has no books referencing them", authorSlug),
			})
		}
	}
}

// validateArticleMoreMarker checks that articles have <!--more--> marker
func validateArticleMoreMarker(ctx *ValidationContext) {
	for _, article := range ctx.Articles {
		if !article.HasMoreMark {
			ctx.Errors = append(ctx.Errors, ValidationError{
				BookSlug:  article.Slug,
				FilePath:  article.Path,
				ErrorType: "MISSING_MORE_MARKER",
				Message:   fmt.Sprintf("Article is missing <!--more--> separator for summary generation"),
			})
		}
	}
}

// validateArticleProperties checks that articles have valid frontmatter properties
func validateArticleProperties(ctx *ValidationContext) {
	for _, article := range ctx.Articles {
		// Check for invalid properties
		for prop := range article.Properties {
			if !allowedArticleProperties[prop] {
				ctx.Errors = append(ctx.Errors, ValidationError{
					BookSlug:  article.Slug,
					FilePath:  article.Path,
					ErrorType: "INVALID_PROPERTY",
					Message:   fmt.Sprintf("Article has invalid property: %s", prop),
				})
			}
		}

		// Check for required properties
		for _, prop := range requiredArticleProperties {
			value := article.Properties[prop]
			if value == "" || value == "null" {
				ctx.Errors = append(ctx.Errors, ValidationError{
					BookSlug:  article.Slug,
					FilePath:  article.Path,
					ErrorType: "MISSING_PROPERTY",
					Message:   fmt.Sprintf("Article missing required property: %s", prop),
				})
			}
		}

		// Validate date format (if present)
		if article.Date != "" {
			datePattern := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
			if !datePattern.MatchString(article.Date) {
				ctx.Errors = append(ctx.Errors, ValidationError{
					BookSlug:  article.Slug,
					FilePath:  article.Path,
					ErrorType: "INVALID_DATE_FORMAT",
					Message:   fmt.Sprintf("Article has invalid date format: %s (expected YYYY-MM-DD)", article.Date),
				})
			}
		}
	}
}

// validatePhilosopherBooks checks that philosopher book references are valid
func validatePhilosopherBooks(ctx *ValidationContext) {
	for _, philosopher := range ctx.Philosophers {
		for _, bookSlug := range philosopher.Books {
			if _, exists := ctx.Books[bookSlug]; !exists {
				ctx.Errors = append(ctx.Errors, ValidationError{
					BookSlug:  philosopher.Slug,
					FilePath:  philosopher.Path,
					ErrorType: "INVALID_BOOK_REFERENCE",
					Message:   fmt.Sprintf("Philosopher references non-existent book: '%s'", bookSlug),
				})
			}
		}
	}
}

// validateBookMeta runs all validations
func validateBookMeta(ctx *ValidationContext) {
	validateAllowedProperties(ctx)
	validateAuthors(ctx)
	validateUnusedAuthors(ctx)
	validateUniqueTitles(ctx)
	validateUniqueAmazonURLs(ctx)
	validateImageFiles(ctx)
	validateReadingStatus(ctx)
	validateBooleanFlags(ctx)
	validateRelatedBooks(ctx)
}

// printResults prints validation results
func printResults(ctx *ValidationContext) {
	fmt.Println("=== Content Validation ===")
	fmt.Println()
	fmt.Printf("Validating %d books, %d articles, and %d philosophers...\n", len(ctx.Books), len(ctx.Articles), len(ctx.Philosophers))
	fmt.Println()

	if len(ctx.Errors) > 0 {
		fmt.Println("ERRORS FOUND:")
		fmt.Println()

		// Sort errors by book slug for consistent output
		sort.Slice(ctx.Errors, func(i, j int) bool {
			if ctx.Errors[i].BookSlug == ctx.Errors[j].BookSlug {
				return ctx.Errors[i].ErrorType < ctx.Errors[j].ErrorType
			}
			return ctx.Errors[i].BookSlug < ctx.Errors[j].BookSlug
		})

		for _, err := range ctx.Errors {
			fmt.Printf("[%s] %s\n", err.ErrorType, err.BookSlug)
			fmt.Printf("  %s\n", err.FilePath)
			fmt.Printf("  → %s\n", err.Message)
			fmt.Println()
		}
	}

	if len(ctx.Warnings) > 0 {
		fmt.Println("WARNINGS:")
		fmt.Println()

		// Sort warnings by book slug
		sort.Slice(ctx.Warnings, func(i, j int) bool {
			return ctx.Warnings[i].BookSlug < ctx.Warnings[j].BookSlug
		})

		for _, warn := range ctx.Warnings {
			fmt.Printf("[NON_RECIPROCAL_RELATION] %s\n", warn.BookSlug)
			fmt.Printf("  %s\n", warn.FilePath)
			fmt.Printf("  → %s\n", warn.Message)
			fmt.Println()
		}
	}

	// Print summary
	fmt.Println("=== Summary ===")
	fmt.Printf("Total books: %d\n", len(ctx.Books))

	// Count unique books with errors
	booksWithErrors := make(map[string]bool)
	for _, err := range ctx.Errors {
		booksWithErrors[err.BookSlug] = true
	}
	fmt.Printf("Books with errors: %d\n", len(booksWithErrors))

	// Count unique books with warnings
	booksWithWarnings := make(map[string]bool)
	for _, warn := range ctx.Warnings {
		booksWithWarnings[warn.BookSlug] = true
	}
	fmt.Printf("Books with warnings: %d\n", len(booksWithWarnings))

	fmt.Printf("Total errors: %d\n", len(ctx.Errors))
	fmt.Printf("Total warnings: %d\n", len(ctx.Warnings))
	fmt.Println()

	if len(ctx.Errors) > 0 {
		fmt.Println("✗ VALIDATION FAILED")
	} else {
		fmt.Println("✓ VALIDATION PASSED")
	}
}

func main() {
	// Load validation context
	ctx, err := loadValidationContext()
	if err != nil {
		fmt.Printf("Error loading validation context: %v\n", err)
		os.Exit(1)
	}

	// Run validations
	validateBookMeta(ctx)
	validateArticleMoreMarker(ctx)
	validateArticleProperties(ctx)
	validatePhilosopherBooks(ctx)

	// Print results
	printResults(ctx)

	// Exit with appropriate code
	if len(ctx.Errors) > 0 {
		os.Exit(1)
	}
}
