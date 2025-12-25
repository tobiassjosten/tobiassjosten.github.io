package main

import (
	"fmt"
	"os"
	"sort"

	"gopkg.in/yaml.v3"
)

type TagData struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

func main() {
	// Read tags from data/tags.yaml
	data, err := os.ReadFile("data/tags.yaml")
	if err != nil {
		fmt.Printf("Error reading tags.yaml: %v\n", err)
		return
	}

	var tags map[string]TagData
	err = yaml.Unmarshal(data, &tags)
	if err != nil {
		fmt.Printf("Error parsing tags.yaml: %v\n", err)
		return
	}

	// Get sorted tag list
	var tagNames []string
	for tag := range tags {
		tagNames = append(tagNames, tag)
	}
	sort.Strings(tagNames)

	// Generate _redirects file
	file, err := os.Create("static/_redirects")
	if err != nil {
		fmt.Printf("Error creating _redirects: %v\n", err)
		return
	}
	defer file.Close()

	// Write header
	file.WriteString("# Tag page redirects - old Jekyll URLs to new Hugo tag pages\n")
	file.WriteString("# Format: /old-path /new-path 301\n\n")

	// Write tag index page redirects
	file.WriteString("# Tag index pages\n")
	for _, tag := range tagNames {
		file.WriteString(fmt.Sprintf("/%s/ /tags/%s/ 301\n", tag, tag))
	}

	// Write tag feed redirects
	file.WriteString("\n# Tag feed redirects - old .atom feeds to new .xml feeds\n")
	for _, tag := range tagNames {
		file.WriteString(fmt.Sprintf("/%s/feed.atom /tags/%s/feed.xml 301\n", tag, tag))
	}

	fmt.Printf("Generated _redirects file with %d tags (%d redirects)\n", len(tagNames), len(tagNames)*2)
}
