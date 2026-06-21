package scripts

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ImageRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type ImageResponse struct {
	Data []ImageData `json:"data"`
}

type ImageData struct {
	B64JSON string `json:"b64_json"`
}

type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// IllustrateThinkers generates a portrait for every thinker missing its image,
// using the OpenAI image API (requires OPENAI_API_KEY).
func IllustrateThinkers(args []string) error {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("OPENAI_API_KEY not set")
	}

	thinkerPaths, err := filepath.Glob("content/thinkers/*/index.md")
	if err != nil {
		return fmt.Errorf("finding thinkers: %w", err)
	}

	fmt.Printf("Found %d thinkers\n\n", len(thinkerPaths))

	generated := 0
	skipped := 0
	failed := 0

	for i, path := range thinkerPaths {
		dir := filepath.Dir(path)

		title, image, err := parseFrontmatter(path)
		if err != nil {
			fmt.Printf("[%d/%d] ✗ %s: %v\n", i+1, len(thinkerPaths), dir, err)
			failed++
			continue
		}

		imagePath := filepath.Join(dir, image)
		if _, err := os.Stat(imagePath); err == nil {
			skipped++
			continue
		}

		fmt.Printf("[%d/%d] Generating portrait for %s\n", i+1, len(thinkerPaths), title)

		if err := generatePortrait(apiKey, title, imagePath); err != nil {
			fmt.Printf("  ✗ Failed: %v\n", err)
			failed++
			continue
		}

		fmt.Printf("  ✓ Saved %s\n", imagePath)
		generated++

		time.Sleep(2 * time.Second)
	}

	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Generated: %d\n", generated)
	fmt.Printf("Skipped (already exists): %d\n", skipped)
	fmt.Printf("Failed: %d\n", failed)

	return nil
}

func parseFrontmatter(path string) (title, image string, err error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", "", fmt.Errorf("reading file: %w", err)
	}

	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return "", "", fmt.Errorf("invalid frontmatter format")
	}

	for _, line := range strings.Split(parts[1], "\n") {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "title:") {
			title = strings.TrimSpace(strings.TrimPrefix(line, "title:"))
			title = strings.Trim(title, "\"")
		}

		if strings.HasPrefix(line, "image:") {
			image = strings.TrimSpace(strings.TrimPrefix(line, "image:"))
			image = strings.Trim(image, "\"")
		}
	}

	if title == "" {
		return "", "", fmt.Errorf("missing title")
	}
	if image == "" {
		return "", "", fmt.Errorf("missing image")
	}

	return title, image, nil
}

func generatePortrait(apiKey, name, outputPath string) error {
	prompt := fmt.Sprintf(
		`A minimalist black-and-white line art illustration of %s, shown as a bust portrait from the chest up, wearing a simple ancient Greek robe (himation) draped over one shoulder.

		Style: clean vector-style line art, bold black ink outlines, smooth flowing strokes, minimal detail, no shading, no grayscale, no color, only black lines on a white background.

		The drawing should resemble a philosopher, slightly stylized but recognizable, similar to modern educational or encyclopedia illustrations.

		Composition: centered portrait, simple lines defining facial features, beard, and robe folds. Background plain white. High contrast, crisp edges, scalable SVG-style illustration.
		`,
		name,
	)

	reqBody := ImageRequest{
		Model:  "gpt-image-1",
		Prompt: prompt,
		N:      1,
		Size:   "1024x1024",
	}

	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("marshaling request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", strings.NewReader(string(reqJSON)))
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 120 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode != 200 {
		var errResp ErrorResponse
		if json.Unmarshal(body, &errResp) == nil && errResp.Error.Message != "" {
			return fmt.Errorf("API error (%d): %s", resp.StatusCode, errResp.Error.Message)
		}
		return fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var imgResp ImageResponse
	if err := json.Unmarshal(body, &imgResp); err != nil {
		return fmt.Errorf("parsing response: %w", err)
	}

	if len(imgResp.Data) == 0 {
		return fmt.Errorf("no image data in response")
	}

	imageBytes, err := base64.StdEncoding.DecodeString(imgResp.Data[0].B64JSON)
	if err != nil {
		return fmt.Errorf("decoding base64: %w", err)
	}

	if err := os.WriteFile(outputPath, imageBytes, 0644); err != nil {
		return fmt.Errorf("writing image: %w", err)
	}

	return nil
}
