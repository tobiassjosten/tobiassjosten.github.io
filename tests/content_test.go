package tests

import (
	"sort"
	"testing"
)

// TestContent enforces the content-authoring invariants on the source tree
// under content/ (ported from the former scripts/validate.go). Each rule is a
// subtest; errors are reported via t.Errorf and warnings via t.Logf.
func TestContent(t *testing.T) {
	// Parse errors gathered while loading the source are reported up front;
	// the per-rule subtests below reset the error slice before running.
	loadErrors := append([]ValidationError(nil), ctx.Errors...)
	t.Run("Parsing", func(t *testing.T) {
		for _, e := range loadErrors {
			t.Errorf("%s: %s", e.FilePath, e.Message)
		}
	})

	rules := []struct {
		name string
		fn   func(*ValidationContext)
	}{
		{"AllowedProperties", validateAllowedProperties},
		{"Authors", validateAuthors},
		{"UnusedAuthors", validateUnusedAuthors},
		{"UniqueTitles", validateUniqueTitles},
		{"UniqueAmazonURLs", validateUniqueAmazonURLs},
		{"ImageFiles", validateImageFiles},
		{"ReadingStatus", validateReadingStatus},
		{"BooleanFlags", validateBooleanFlags},
		{"RelatedBooks", validateRelatedBooks},
		{"ArticleMoreMarker", validateArticleMoreMarker},
		{"ArticleProperties", validateArticleProperties},
		{"ThinkerBooks", validateThinkerBooks},
	}

	for _, rule := range rules {
		t.Run(rule.name, func(t *testing.T) {
			ctx.Errors = nil
			ctx.Warnings = nil
			rule.fn(ctx)

			errs := append([]ValidationError(nil), ctx.Errors...)
			sort.Slice(errs, func(i, j int) bool { return errs[i].BookSlug < errs[j].BookSlug })
			for _, e := range errs {
				t.Errorf("[%s] %s: %s", e.BookSlug, e.FilePath, e.Message)
			}

			warns := append([]ValidationWarning(nil), ctx.Warnings...)
			sort.Slice(warns, func(i, j int) bool { return warns[i].BookSlug < warns[j].BookSlug })
			for _, w := range warns {
				t.Logf("warning [%s] %s: %s", w.BookSlug, w.FilePath, w.Message)
			}
		})
	}
}
