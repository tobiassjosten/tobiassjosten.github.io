package tests

import (
	"fmt"
	"path"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// TestBookCovers asserts that every rendered book page displays a cover image
// drawn from the book's own bundle that resolves in the build output.
func TestBookCovers(t *testing.T) {
	for slug := range ctx.Books {
		slug := slug
		t.Run(slug, func(t *testing.T) {
			page := path.Join(siteDir, "books", slug, "index.html")
			if !targetExists("/books/" + slug + "/") {
				t.Fatalf("book page not rendered at /books/%s/", slug)
			}

			doc := parsePage(t, page)
			prefix := fmt.Sprintf("/books/%s/", slug)

			found := false
			doc.Find("img[src]").EachWithBreak(func(_ int, s *goquery.Selection) bool {
				src, _ := s.Attr("src")
				target, check := resolveInternal(src, prefix)
				if check && strings.HasPrefix(target, prefix) && target != prefix && targetExists(target) {
					found = true
					return false
				}
				return true
			})

			if !found {
				t.Errorf("book %q has no resolvable cover image from its own bundle", slug)
			}
		})
	}
}
