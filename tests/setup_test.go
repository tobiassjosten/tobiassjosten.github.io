package tests

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// internalHost is the baseURL host used for the test build. Any absolute link
// pointing at this host is treated as internal and resolved against siteDir;
// links to any other host are external and skipped.
const internalHost = "test.local"

// siteDir holds the freshly built site for the duration of the test run.
var siteDir string

// ctx holds the parsed content/ source tree, loaded once.
var ctx *ValidationContext

// redirectMap holds the static/_redirects rules, keyed by source path with any
// trailing slash stripped, mapping to the raw target.
var redirectMap map[string]string

// TestMain moves to the project root, builds the site into a temp directory,
// and loads the content source so every test shares the same artifacts.
func TestMain(m *testing.M) {
	if err := os.Chdir(".."); err != nil {
		fmt.Fprintf(os.Stderr, "chdir to project root: %v\n", err)
		os.Exit(1)
	}
	if _, err := os.Stat("config.yaml"); err != nil {
		fmt.Fprintf(os.Stderr, "not at project root (config.yaml not found): %v\n", err)
		os.Exit(1)
	}

	dir, err := os.MkdirTemp("", "site-build-*")
	if err != nil {
		fmt.Fprintf(os.Stderr, "creating temp build dir: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)
	siteDir = dir

	cmd := exec.Command("hugo",
		"--environment", "production",
		"--destination", dir,
		"--cleanDestinationDir",
		"--gc",
		"--minify",
		"--baseURL", "http://"+internalHost+"/",
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Fprintf(os.Stderr, "hugo build failed: %v\n%s\n", err, out)
		os.Exit(1)
	}

	ctx, err = loadValidationContext()
	if err != nil {
		fmt.Fprintf(os.Stderr, "loading content: %v\n", err)
		os.Exit(1)
	}

	if err := loadRedirects(); err != nil {
		fmt.Fprintf(os.Stderr, "loading redirects: %v\n", err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

// htmlPages returns every built .html file in the site, keyed by its absolute
// path on disk.
func htmlPages(t *testing.T) []string {
	t.Helper()
	var pages []string
	err := filepath.WalkDir(siteDir, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(p, ".html") {
			pages = append(pages, p)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking build dir: %v", err)
	}
	return pages
}

// parsePage loads a built HTML file into a goquery document.
func parsePage(t *testing.T, page string) *goquery.Document {
	t.Helper()
	f, err := os.Open(page)
	if err != nil {
		t.Fatalf("opening %s: %v", page, err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		t.Fatalf("parsing %s: %v", page, err)
	}
	return doc
}

// pageURLPath returns the site-root-relative URL directory for a built page,
// used to resolve relative links. e.g. <siteDir>/books/x/index.html -> /books/x/
func pageURLPath(page string) string {
	rel, err := filepath.Rel(siteDir, page)
	if err != nil {
		return "/"
	}
	return "/" + filepath.ToSlash(filepath.Dir(rel)) + "/"
}

// pageLinks extracts every href/src/srcset reference from a parsed page.
func pageLinks(doc *goquery.Document) []string {
	var links []string
	add := func(v string) {
		v = strings.TrimSpace(v)
		if v != "" {
			links = append(links, v)
		}
	}

	doc.Find("a[href], link[href], area[href]").Each(func(_ int, s *goquery.Selection) {
		if v, ok := s.Attr("href"); ok {
			add(v)
		}
	})
	doc.Find("img[src], script[src], iframe[src], source[src], audio[src], video[src]").Each(func(_ int, s *goquery.Selection) {
		if v, ok := s.Attr("src"); ok {
			add(v)
		}
	})
	doc.Find("img[srcset], source[srcset]").Each(func(_ int, s *goquery.Selection) {
		if v, ok := s.Attr("srcset"); ok {
			for _, cand := range strings.Split(v, ",") {
				fields := strings.Fields(cand)
				if len(fields) > 0 {
					add(fields[0])
				}
			}
		}
	})

	return links
}

// resolveInternal classifies a link found on the page at pageDir. It returns
// the site-root-relative path to check, and whether the link is internal and
// worth checking at all.
func resolveInternal(link, pageDir string) (target string, check bool) {
	u, err := url.Parse(link)
	if err != nil {
		return "", false
	}

	switch u.Scheme {
	case "mailto", "tel", "javascript", "data":
		return "", false
	}

	// Absolute URL: only internal if it points at our build host.
	if u.Host != "" {
		if u.Host != internalHost {
			return "", false
		}
	}

	p := u.Path
	if p == "" {
		// Pure fragment or query against the current page.
		return "", false
	}

	if !strings.HasPrefix(p, "/") {
		p = path.Join(pageDir, p)
	}
	return path.Clean(p), true
}

// targetExists reports whether a site-root-relative path resolves to a file in
// the build output. Directory paths resolve via their index.html.
func targetExists(target string) bool {
	full := filepath.Join(siteDir, filepath.FromSlash(target))
	info, err := os.Stat(full)
	if err != nil {
		return false
	}
	if info.IsDir() {
		_, err := os.Stat(filepath.Join(full, "index.html"))
		return err == nil
	}
	return true
}

// normPath strips a trailing slash so links and redirect sources compare
// regardless of how they were written ("/foo" and "/foo/" are the same).
func normPath(p string) string {
	if p == "/" {
		return p
	}
	return strings.TrimSuffix(p, "/")
}

// loadRedirects parses static/_redirects into redirectMap.
func loadRedirects() error {
	redirectMap = make(map[string]string)
	data, err := os.ReadFile("static/_redirects")
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		redirectMap[normPath(fields[0])] = fields[1]
	}
	return nil
}

// hasRedirect reports whether a site-root-relative path matches a redirect rule.
func hasRedirect(target string) bool {
	_, ok := redirectMap[normPath(target)]
	return ok
}

// resolveRedirect follows a redirect target through any chain of further rules
// and reports the final destination and whether it is reachable (an external
// URL, or an internal path that resolves to a file).
func resolveRedirect(to string) (final string, ok bool) {
	seen := make(map[string]bool)
	cur := to
	for i := 0; i < 16; i++ {
		if strings.HasPrefix(cur, "http://") || strings.HasPrefix(cur, "https://") {
			return cur, true
		}
		if targetExists(cur) {
			return cur, true
		}
		n := normPath(cur)
		if seen[n] {
			return cur, false // redirect cycle
		}
		seen[n] = true
		if next, found := redirectMap[n]; found {
			cur = next
			continue
		}
		return cur, false
	}
	return cur, false
}
