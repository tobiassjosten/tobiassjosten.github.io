package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// This file ports the content-authoring rules previously enforced by
// scripts/validate.go. The loaders and validators read the source under
// content/; they are exercised as subtests from content_test.go.

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
	"series":             true,
	"aliases":            true,
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
	AllProperties      map[string]string
	Title              string
	Authors            []string
	Date               string
	AmazonURL          string
	Image              string
	RelatedBooksRaw    string
	RelatedBooks       []string
	Rating             string
	CurrentlyReading   bool
	FeaturedOnHomepage bool
	Content            string
	HasContent         bool
}

// ArticleMetadata represents parsed article content
type ArticleMetadata struct {
	Path        string
	Slug        string
	Properties  map[string]string
	Title       string
	Date        string
	Draft       bool
	Content     string
	HasMoreMark bool
}

// ThinkerMetadata represents parsed thinker frontmatter
type ThinkerMetadata struct {
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
	Books    map[string]*BookMetadata
	Articles map[string]*ArticleMetadata
	Thinkers map[string]*ThinkerMetadata
	Authors  map[string]bool
	Errors   []ValidationError
	Warnings []ValidationWarning
}

// parseFrontMatter parses YAML frontmatter from content.
func parseFrontMatter(content string) (map[string]string, error) {
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid front matter format")
	}

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

		if inMultilineValue {
			if strings.HasPrefix(trimmed, "-") {
				item := strings.TrimSpace(strings.TrimPrefix(trimmed, "-"))
				item = strings.Trim(item, "\"")
				multilineValue = append(multilineValue, item)
			} else if strings.Contains(trimmed, ":") {
				fm[currentKey] = strings.Join(multilineValue, ",")
				inMultilineValue = false
				multilineValue = nil
			}
		}

		if !inMultilineValue && strings.Contains(trimmed, ":") {
			parts := strings.SplitN(trimmed, ":", 2)
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
				fm[key] = value
			} else if value == "" {
				currentKey = key
				inMultilineValue = true
				multilineValue = []string{}
			} else if value == "null" {
				fm[key] = value
			} else {
				fm[key] = strings.Trim(value, "\"")
			}
		}
	}

	if inMultilineValue {
		fm[currentKey] = strings.Join(multilineValue, ",")
	}

	return fm, nil
}

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

	if authorsRaw, exists := fm["authors"]; exists && authorsRaw != "" {
		for _, a := range strings.Split(authorsRaw, ",") {
			a = strings.TrimSpace(a)
			if a != "" {
				book.Authors = append(book.Authors, a)
			}
		}
	}

	if relatedBooksRaw, exists := fm["relatedBooks"]; exists && relatedBooksRaw != "" {
		book.RelatedBooksRaw = relatedBooksRaw
		if !strings.HasPrefix(relatedBooksRaw, "[") {
			for _, b := range strings.Split(relatedBooksRaw, ",") {
				b = strings.TrimSpace(b)
				if b != "" {
					book.RelatedBooks = append(book.RelatedBooks, b)
				}
			}
		}
	}

	if cr, exists := fm["currentlyReading"]; exists && cr == "true" {
		book.CurrentlyReading = true
	}
	if fh, exists := fm["featuredOnHomepage"]; exists && fh == "true" {
		book.FeaturedOnHomepage = true
	}

	return book, nil
}

func parseThinkerFile(path string) (*ThinkerMetadata, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	fm, err := parseFrontMatter(string(content))
	if err != nil {
		return nil, fmt.Errorf("parsing frontmatter: %w", err)
	}

	slug := filepath.Base(filepath.Dir(path))

	var books []string
	if booksRaw, exists := fm["books"]; exists && booksRaw != "" {
		for _, b := range strings.Split(booksRaw, ",") {
			b = strings.TrimSpace(b)
			if b != "" {
				books = append(books, b)
			}
		}
	}

	return &ThinkerMetadata{
		Path:  path,
		Slug:  slug,
		Title: fm["title"],
		Books: books,
	}, nil
}

func parseArticleFile(path string) (*ArticleMetadata, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	contentStr := string(content)

	parts := strings.SplitN(contentStr, "---", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid front matter format")
	}

	properties, _ := parseFrontMatter(contentStr)
	slug := strings.TrimSuffix(filepath.Base(path), ".md")

	bodyContent := parts[2]

	return &ArticleMetadata{
		Path:        path,
		Slug:        slug,
		Properties:  properties,
		Title:       properties["title"],
		Date:        properties["date"],
		Draft:       properties["draft"] == "true",
		Content:     strings.TrimSpace(bodyContent),
		HasMoreMark: strings.Contains(bodyContent, "<!--more-->"),
	}, nil
}

// loadValidationContext loads all books, articles, thinkers, and authors from
// the content/ source tree. The working directory must be the project root.
func loadValidationContext() (*ValidationContext, error) {
	ctx := &ValidationContext{
		Books:    make(map[string]*BookMetadata),
		Articles: make(map[string]*ArticleMetadata),
		Thinkers: make(map[string]*ThinkerMetadata),
		Authors:  make(map[string]bool),
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

	articlePaths, err := filepath.Glob("content/articles/*.md")
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

	thinkerPaths, err := filepath.Glob("content/thinkers/*/index.md")
	if err != nil {
		return nil, fmt.Errorf("finding thinker files: %w", err)
	}
	for _, path := range thinkerPaths {
		thinker, err := parseThinkerFile(path)
		if err != nil {
			ctx.Errors = append(ctx.Errors, ValidationError{
				BookSlug:  filepath.Base(filepath.Dir(path)),
				FilePath:  path,
				ErrorType: "PARSE_ERROR",
				Message:   fmt.Sprintf("Failed to parse: %v", err),
			})
			continue
		}
		ctx.Thinkers[thinker.Slug] = thinker
	}

	authorDirs, err := filepath.Glob("content/authors/*")
	if err != nil {
		return nil, fmt.Errorf("finding author directories: %w", err)
	}
	for _, dir := range authorDirs {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			ctx.Authors[filepath.Base(dir)] = true
		}
	}

	return ctx, nil
}

func (ctx *ValidationContext) addError(book *BookMetadata, errorType, message string) {
	ctx.Errors = append(ctx.Errors, ValidationError{
		BookSlug:  book.Slug,
		FilePath:  book.Path,
		ErrorType: errorType,
		Message:   message,
	})
}

func (ctx *ValidationContext) addWarning(book *BookMetadata, message string) {
	ctx.Warnings = append(ctx.Warnings, ValidationWarning{
		BookSlug: book.Slug,
		FilePath: book.Path,
		Message:  message,
	})
}

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

func validateUniqueTitles(ctx *ValidationContext) {
	titleMap := make(map[string][]string)
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
				ctx.addError(ctx.Books[slug], "DUPLICATE_TITLE",
					fmt.Sprintf("Title '%s' is used by multiple books: %s",
						title, strings.Join(slugs, ", ")))
			}
		}
	}
}

func validateUniqueAmazonURLs(ctx *ValidationContext) {
	urlMap := make(map[string][]string)
	for _, book := range ctx.Books {
		if book.AmazonURL == "" {
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
				ctx.addError(ctx.Books[slug], "DUPLICATE_AMAZON_URL",
					fmt.Sprintf("Amazon URL '%s' is used by multiple books: %s",
						url, strings.Join(slugs, ", ")))
			}
		}
	}
}

func validateImageFiles(ctx *ValidationContext) {
	for _, book := range ctx.Books {
		if book.Image == "" {
			ctx.addError(book, "MISSING_IMAGE", "Missing required 'image' field")
			continue
		}
		imagePath := filepath.Join(filepath.Dir(book.Path), book.Image)
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

func validateReadingStatus(ctx *ValidationContext) {
	datePattern := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	for _, book := range ctx.Books {
		if book.CurrentlyReading {
			continue
		}
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

func validateBooleanFlags(ctx *ValidationContext) {
	for _, book := range ctx.Books {
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

func validateRelatedBooks(ctx *ValidationContext) {
	for _, book := range ctx.Books {
		if book.RelatedBooksRaw != "" && strings.HasPrefix(book.RelatedBooksRaw, "[") {
			ctx.addError(book, "INVALID_RELATED_BOOKS_FORMAT",
				"relatedBooks uses bracket syntax []. Use multiline YAML format instead (one per line with '  - ')")
			continue
		}
		for _, relatedSlug := range book.RelatedBooks {
			if _, exists := ctx.Books[relatedSlug]; !exists {
				ctx.addError(book, "INVALID_RELATED_BOOK",
					fmt.Sprintf("Related book '%s' does not exist", relatedSlug))
			}
		}
		for _, relatedSlug := range book.RelatedBooks {
			relatedBook, exists := ctx.Books[relatedSlug]
			if !exists {
				continue
			}
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

func validateUnusedAuthors(ctx *ValidationContext) {
	referencedAuthors := make(map[string]bool)
	for _, book := range ctx.Books {
		for _, authorSlug := range book.Authors {
			referencedAuthors[authorSlug] = true
		}
	}
	for authorSlug := range ctx.Authors {
		if !referencedAuthors[authorSlug] {
			ctx.Errors = append(ctx.Errors, ValidationError{
				BookSlug:  authorSlug,
				FilePath:  filepath.Join("content/authors", authorSlug),
				ErrorType: "UNUSED_AUTHOR",
				Message:   fmt.Sprintf("Author '%s' has no books referencing them", authorSlug),
			})
		}
	}
}

func validateArticleMoreMarker(ctx *ValidationContext) {
	for _, article := range ctx.Articles {
		if !article.HasMoreMark {
			ctx.Errors = append(ctx.Errors, ValidationError{
				BookSlug:  article.Slug,
				FilePath:  article.Path,
				ErrorType: "MISSING_MORE_MARKER",
				Message:   "Article is missing <!--more--> separator for summary generation",
			})
		}
	}
}

func validateArticleProperties(ctx *ValidationContext) {
	datePattern := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	for _, article := range ctx.Articles {
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
		if article.Date != "" && !datePattern.MatchString(article.Date) {
			ctx.Errors = append(ctx.Errors, ValidationError{
				BookSlug:  article.Slug,
				FilePath:  article.Path,
				ErrorType: "INVALID_DATE_FORMAT",
				Message:   fmt.Sprintf("Article has invalid date format: %s (expected YYYY-MM-DD)", article.Date),
			})
		}
	}
}

func validateThinkerBooks(ctx *ValidationContext) {
	for _, thinker := range ctx.Thinkers {
		for _, bookSlug := range thinker.Books {
			if _, exists := ctx.Books[bookSlug]; !exists {
				ctx.Errors = append(ctx.Errors, ValidationError{
					BookSlug:  thinker.Slug,
					FilePath:  thinker.Path,
					ErrorType: "INVALID_BOOK_REFERENCE",
					Message:   fmt.Sprintf("Thinker references non-existent book: '%s'", bookSlug),
				})
			}
		}
	}
}
