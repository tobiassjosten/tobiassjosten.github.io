package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
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
	{"newsletter", "Create newsletter (auto-dated/numbered)", false},
	{"interview", "Create interview (auto-dated/numbered)", false},
	{"book", "Create book review (requires slug)", true},
	{"author", "Create author (requires slug)", true},
	{"philosopher", "Create philosopher (requires slug)", true},
	{"project", "Create project (requires slug)", true},
	{"business", "Create business (requires slug)", true},
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
	fmt.Println("  make new newsletter")
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
	case "newsletter":
		return createNewsletter()
	case "interview":
		return createInterview()
	case "book":
		return createBook(slug)
	case "author":
		return createAuthor(slug)
	case "philosopher":
		return createPhilosopher(slug)
	case "project":
		return createProject(slug)
	case "business":
		return createBusiness(slug)
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

func createNewsletter() error {
	date := time.Now().Format("2006-01-02")

	nextNum, err := findNextNumber("content/newsletters", `newsletter-(\d+)`)
	if err != nil {
		return fmt.Errorf("finding next newsletter number: %w", err)
	}

	filePath := fmt.Sprintf("content/newsletters/%s-newsletter-%d.md", date, nextNum)

	if fileExists(filePath) {
		return fmt.Errorf("newsletter already exists: %s", filePath)
	}

	if err := runHugoNew(filePath); err != nil {
		return err
	}

	if err := replaceIssueNumber(filePath, nextNum); err != nil {
		return fmt.Errorf("updating issue number: %w", err)
	}

	fmt.Printf("✓ Created newsletter #%d: %s\n", nextNum, filePath)
	fmt.Println("  Update title and add <!--more--> separator")
	return nil
}

func createInterview() error {
	date := time.Now().Format("2006-01-02")

	nextNum, err := findNextNumber("content/interviews", `interview-(\d+)`)
	if err != nil {
		return fmt.Errorf("finding next interview number: %w", err)
	}

	filePath := fmt.Sprintf("content/interviews/%s-interview-%d.md", date, nextNum)

	if fileExists(filePath) {
		return fmt.Errorf("interview already exists: %s", filePath)
	}

	if err := runHugoNew(filePath); err != nil {
		return err
	}

	fmt.Printf("✓ Created interview #%d: %s\n", nextNum, filePath)
	fmt.Println("  Fill in interviewee details and add <!--more--> separator")
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

func createPhilosopher(slug string) error {
	dirPath := fmt.Sprintf("content/philosophers/%s", slug)
	filePath := fmt.Sprintf("%s/index.md", dirPath)

	if fileExists(dirPath) {
		return fmt.Errorf("philosopher already exists: %s", dirPath)
	}

	if err := runHugoNew(filePath); err != nil {
		return err
	}

	fmt.Printf("✓ Created philosopher: %s\n", filePath)
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

func createBusiness(slug string) error {
	dirPath := fmt.Sprintf("content/businesses/%s", slug)
	filePath := fmt.Sprintf("%s/index.md", dirPath)

	if fileExists(dirPath) {
		return fmt.Errorf("business already exists: %s", dirPath)
	}

	if err := runHugoNew(filePath); err != nil {
		return err
	}

	fmt.Printf("✓ Created business: %s\n", filePath)
	fmt.Println("  Fill in description and project_url")
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

func findNextNumber(dir, pattern string) (int, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.md"))
	if err != nil {
		return 0, err
	}

	re := regexp.MustCompile(pattern)
	var numbers []int

	for _, file := range files {
		matches := re.FindStringSubmatch(filepath.Base(file))
		if len(matches) > 1 {
			if num, err := strconv.Atoi(matches[1]); err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	if len(numbers) == 0 {
		return 1, nil
	}

	sort.Ints(numbers)
	return numbers[len(numbers)-1] + 1, nil
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

func replaceIssueNumber(filePath string, issueNum int) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	updated := strings.Replace(string(content), "issue: 1", fmt.Sprintf("issue: %d", issueNum), 1)

	return os.WriteFile(filePath, []byte(updated), 0644)
}
