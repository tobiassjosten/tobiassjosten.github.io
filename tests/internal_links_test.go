package tests

import (
	"path/filepath"
	"sort"
	"testing"
)

// TestInternalLinks asserts that every internal href/src/srcset reference in
// the built site resolves to a real file in the build output. External links
// (other hosts, mailto, tel, ...) are skipped.
func TestInternalLinks(t *testing.T) {
	type broken struct {
		page string
		link string
	}
	var failures []broken
	seen := make(map[string]bool) // page|target -> already checked

	for _, page := range htmlPages(t) {
		pageDir := pageURLPath(page)
		doc := parsePage(t, page)

		for _, link := range pageLinks(doc) {
			target, check := resolveInternal(link, pageDir)
			if !check {
				continue
			}
			key := target
			if seen[key] {
				continue
			}
			seen[key] = true
			if !targetExists(target) && !hasRedirect(target) {
				rel, _ := filepath.Rel(siteDir, page)
				failures = append(failures, broken{page: rel, link: link})
			}
		}
	}

	sort.Slice(failures, func(i, j int) bool {
		if failures[i].page == failures[j].page {
			return failures[i].link < failures[j].link
		}
		return failures[i].page < failures[j].page
	})

	for _, f := range failures {
		t.Errorf("broken internal link %q on page %s", f.link, f.page)
	}
}
