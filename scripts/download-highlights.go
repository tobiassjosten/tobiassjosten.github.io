package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type HighlightsResponse struct {
	Count    int         `json:"count"`
	Next     *string     `json:"next"`
	Previous *string     `json:"previous"`
	Results  []Highlight `json:"results"`
}

type Highlight struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Note string `json:"note"`
}

func main() {
	bookID, err := getBookID()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := validateBookID(bookID); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	token := getAccessToken()

	highlights, err := fetchHighlights(bookID, token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching highlights: %v\n", err)
		os.Exit(1)
	}

	if len(highlights) == 0 {
		fmt.Fprintf(os.Stderr, "No highlights found for book %s\n", bookID)
		os.Exit(0)
	}

	outputHighlights(highlights)
}

func getBookID() (string, error) {
	if args := os.Args[1:]; len(args) > 0 {
		return args[0], nil
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter book ID: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func validateBookID(id string) error {
	if id == "" {
		return fmt.Errorf("book ID cannot be empty")
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid book ID '%s' - must be a number", id)
	}

	return nil
}

func getAccessToken() string {
	if token := os.Getenv("READWISE_ACCESS_TOKEN"); token != "" {
		return token
	}

	fmt.Println("READWISE_ACCESS_TOKEN environment variable not set")
	fmt.Println("Get your access token from: https://readwise.io/access_token")
	fmt.Print("Enter access token: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func buildAPIURL(bookID string, page int) string {
	return fmt.Sprintf(
		"https://readwise.io/api/v2/highlights/?book_id=%s&page=%d&page_size=1000",
		bookID,
		page,
	)
}

func makeAPIRequest(url, token string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Authorization", "Token "+token)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}

	return resp, nil
}

func fetchPage(bookID, token string, page int) (*HighlightsResponse, error) {
	url := buildAPIURL(bookID, page)

	resp, err := makeAPIRequest(url, token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		return nil, fmt.Errorf("rate limit exceeded (429). Please wait before retrying")
	}

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	var highlightsResp HighlightsResponse
	if err := json.Unmarshal(body, &highlightsResp); err != nil {
		return nil, fmt.Errorf("parsing JSON response: %w", err)
	}

	return &highlightsResp, nil
}

func fetchHighlights(bookID, token string) ([]Highlight, error) {
	var allHighlights []Highlight
	page := 1

	for {
		response, err := fetchPage(bookID, token, page)
		if err != nil {
			return nil, err
		}

		allHighlights = append(allHighlights, response.Results...)

		if response.Next == nil || *response.Next == "" {
			break
		}

		page++

		time.Sleep(1 * time.Second)
	}

	slices.Reverse(allHighlights)

	return allHighlights, nil
}

func outputHighlights(highlights []Highlight) {
	for i, h := range highlights {
		if i > 0 {
			fmt.Println()
		}

		fmt.Println(h.Text)
		if h.Note != "" {
			fmt.Println("-", h.Note)
		}
	}
}
