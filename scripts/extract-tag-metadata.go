package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type TagFrontmatter struct {
	Title    string `yaml:"title"`
	Category string `yaml:"category"`
	Summary  string `yaml:"summary"`
}

type TagData struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

func main() {
	// Find all tag directories
	dirs, err := filepath.Glob("*/index.md")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	tags := make(map[string]TagData)

	for _, path := range dirs {
		// Skip content/ and other Hugo directories
		if strings.HasPrefix(path, "content/") ||
			strings.HasPrefix(path, "layouts/") ||
			strings.HasPrefix(path, "static/") ||
			strings.HasPrefix(path, "data/") ||
			strings.HasPrefix(path, "archetypes/") ||
			strings.HasPrefix(path, "scripts/") ||
			strings.HasPrefix(path, "books/") {
			continue
		}

		// Read the index.md file
		content, err := ioutil.ReadFile(path)
		if err != nil {
			continue
		}

		// Parse frontmatter
		parts := strings.SplitN(string(content), "---", 3)
		if len(parts) < 3 {
			continue
		}

		var fm TagFrontmatter
		err = yaml.Unmarshal([]byte(parts[1]), &fm)
		if err != nil || fm.Category == "" {
			continue
		}

		// Extract description from content
		description := strings.TrimSpace(parts[2])
		if fm.Summary != "" {
			description = fm.Summary
		}

		// Store tag data
		tags[fm.Category] = TagData{
			Title:       fm.Title,
			Description: description,
		}

		fmt.Printf("Extracted: %s - %s\n", fm.Category, fm.Title)
	}

	// Write to data/tags.yaml
	os.MkdirAll("data", 0755)
	output, err := yaml.Marshal(tags)
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}

	err = ioutil.WriteFile("data/tags.yaml", output, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("\n=== Extraction Complete ===\n")
	fmt.Printf("Extracted %d tag descriptions to data/tags.yaml\n", len(tags))
}
