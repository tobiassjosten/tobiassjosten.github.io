package tests

import (
	"sort"
	"testing"
)

// TestRedirects asserts that every rule in static/_redirects ultimately
// resolves: an external URL, or an internal path that exists in the build
// (following redirect chains). This keeps the redirect map from rotting as
// content moves.
func TestRedirects(t *testing.T) {
	type dead struct {
		from  string
		final string
	}
	var failures []dead

	for from := range redirectMap {
		final, ok := resolveRedirect(redirectMap[from])
		if !ok {
			failures = append(failures, dead{from: from, final: final})
		}
	}

	sort.Slice(failures, func(i, j int) bool { return failures[i].from < failures[j].from })
	for _, f := range failures {
		t.Errorf("redirect %q leads to unreachable target %q", f.from, f.final)
	}
}
