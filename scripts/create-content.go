package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type contentType struct {
	name        string
	description string
	requiresSlug bool
}

var contentTypes = []contentType{
	{"article", "Create article (requires slug)", true},
	{"book", "Create book review (requires slug)", true},
	{"author", "Create author (requires slug)", true},
	{"thinker", "Create thinker (requires slug)", true},
	{"project", "Create project (requires slug)", true},
	{"presentation", "Create presentation (requires slug)", true},
}

func main() {
	args := os.Args[1:]

	var selectedType string
	var slug string

	if len(args) == 0 {
		selectedType = promptContentType()
	} else {
		selectedType = args[0]
		if !isValidContentType(selectedType) {
			printUsage()
			os.Exit(1)
		}
	}

	ct := getContentType(selectedType)
	if ct.requiresSlug {
		if len(args) >= 2 {
			slug = args[1]
		} else {
			slug = promptSlug(selectedType)
		}

		if err := validateSlug(slug); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			fmt.Fprintf(os.Stderr, "Usage: make new %s <slug>\n", selectedType)
			os.Exit(1)
		}
	}

	if err := createContent(selectedType, slug); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: make new <type> [slug]")
	fmt.Println()
	fmt.Println("Content types:")
	for _, ct := range contentTypes {
		fmt.Printf("  %-13s - %s\n", ct.name, ct.description)
	}
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  make new article my-article-slug")
	fmt.Println("  make new book clean-code")
}

func promptContentType() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select content type:")
	for i, ct := range contentTypes {
		fmt.Printf("  %d. %s - %s\n", i+1, ct.name, ct.description)
	}
	fmt.Print("\nEnter number or name: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if num, err := strconv.Atoi(input); err == nil && num >= 1 && num <= len(contentTypes) {
		return contentTypes[num-1].name
	}

	if isValidContentType(input) {
		return input
	}

	fmt.Fprintf(os.Stderr, "Invalid content type: %s\n", input)
	os.Exit(1)
	return ""
}

func promptSlug(contentType string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter slug for %s: ", contentType)

	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func isValidContentType(name string) bool {
	for _, ct := range contentTypes {
		if ct.name == name {
			return true
		}
	}
	return false
}

func getContentType(name string) contentType {
	for _, ct := range contentTypes {
		if ct.name == name {
			return ct
		}
	}
	return contentType{}
}

func validateSlug(slug string) error {
	if slug == "" {
		return fmt.Errorf("slug is required")
	}

	slugPattern := regexp.MustCompile(`^[a-z0-9-]+$`)
	if !slugPattern.MatchString(slug) {
		return fmt.Errorf("invalid slug '%s' - use lowercase letters, numbers, and hyphens only", slug)
	}

	return nil
}

func createContent(contentType, slug string) error {
	switch contentType {
	case "article":
		return createArticle(slug)
	case "book":
		return createBook(slug)
	case "author":
		return createAuthor(slug)
	case "thinker":
		return createThinker(slug)
	case "project":
		return createProject(slug)
	case "presentation":
		return createPresentation(slug)
	default:
		return fmt.Errorf("unknown content type: %s", contentType)
	}
}

func createArticle(slug string) error {
	filePath := fmt.Sprintf("content/articles/%s.md", slug)

	if fileExists(filePath) {
		return fmt.Errorf("article already exists: %s", filePath)
	}

	if err := runHugoNew(filePath); err != nil {
		return err
	}

	fmt.Printf("✓ Created article: %s\n", filePath)
	fmt.Println("  Remember to add <!--more--> separator")
	return nil
}

func createBook(slug string) error {
	dirPath := fmt.Sprintf("content/books/%s", slug)
	filePath := fmt.Sprintf("%s/index.md", dirPath)

	if fileExists(dirPath) {
		return fmt.Errorf("book already exists: %s", dirPath)
	}

	if err := runHugoNew(filePath); err != nil {
		return err
	}

	fmt.Printf("✓ Created book: %s\n", filePath)
	fmt.Println("  Next steps:")
	fmt.Printf("    1. Add cover image to %s/\n", dirPath)
	fmt.Println("    2. Update image field in frontmatter")
	fmt.Println("    3. Fill in authors, Amazon URL, and rating")
	return nil
}

func createAuthor(slug string) error {
	dirPath := fmt.Sprintf("content/authors/%s", slug)
	filePath := fmt.Sprintf("%s/index.md", dirPath)

	if fileExists(dirPath) {
		return fmt.Errorf("author already exists: %s", dirPath)
	}

	if err := runHugoNew(filePath); err != nil {
		return err
	}

	fmt.Printf("✓ Created author: %s\n", filePath)
	return nil
}

func createThinker(slug string) error {
	dirPath := fmt.Sprintf("content/thinkers/%s", slug)
	filePath := fmt.Sprintf("%s/index.md", dirPath)

	if fileExists(dirPath) {
		return fmt.Errorf("thinker already exists: %s", dirPath)
	}

	if err := runHugoNew(filePath); err != nil {
		return err
	}

	fmt.Printf("✓ Created thinker: %s\n", filePath)
	fmt.Println("  Next steps:")
	fmt.Printf("    1. Add portrait image (PNG) to %s/\n", dirPath)
	fmt.Println("    2. Update image field in frontmatter")
	fmt.Println("    3. Fill in birth_date and death_date")
	return nil
}

func createProject(slug string) error {
	dirPath := fmt.Sprintf("content/projects/%s", slug)
	filePath := fmt.Sprintf("%s/index.md", dirPath)

	if fileExists(dirPath) {
		return fmt.Errorf("project already exists: %s", dirPath)
	}

	if err := runHugoNew(filePath); err != nil {
		return err
	}

	fmt.Printf("✓ Created project: %s\n", filePath)
	fmt.Println("  Fill in description, project_url, and group")
	return nil
}

func createPresentation(slug string) error {
	year := time.Now().Format("2006")
	dirPath := fmt.Sprintf("content/presentations/%s-%s", year, slug)
	filePath := fmt.Sprintf("%s/index.md", dirPath)

	if fileExists(dirPath) {
		return fmt.Errorf("presentation already exists: %s", dirPath)
	}

	if err := runHugoNew(filePath); err != nil {
		return err
	}

	fmt.Printf("✓ Created presentation: %s\n", filePath)
	fmt.Println("  Fill in event, URLs, and optionally add image")
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func runHugoNew(path string) error {
	cmd := exec.Command("hugo", "new", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

