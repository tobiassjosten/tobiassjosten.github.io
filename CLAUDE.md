# Claude instructions

## Tech stack

- Hugo generates the website
- Plain HTML/CSS/JS for frontend
- Cloudflare Pages hosts the site
- Go for utility scripts

## Development

Don't bother backing up files. Everything is tracked by Git and can easily be restored.

Only create new utility scripts when explicitly told to do so. Instead rely on shell commands and Posix tools for ad-hoc tasks.

Default to not commenting code. Instead use clear and descriptive names for variables, functions, files, and directories.

For CSS and styling, rely on Bootstrap 5, using its built-in extension points like variables. Use custom CSS and code only when necessary.

Don't run `hugo server`, but instead use `hugo build`.

Run the following commands to verify that the website builds correctly:

```bash
hugo --environment production \
     --destination public \
     --cleanDestinationDir \
     --gc \
     --minify \
     --panicOnWarning \
     --printPathWarnings
```

```bash
hugo mod tidy
hugo mod verify
```

Run the following commands to verify that the content is valid:

```bash
go run scripts/validate.go
```

## Content

Content lives under `content/` and is grouped by type. Each type has its own section template under `layouts/section/` and, where relevant, a single-page template under `layouts/<type>/single.html`.

- **Pages** — standalone pages with their own URL: `connect/`, `now/`, `uses/`, `projects/`, `services/`, `resume/`, `speaking/`. Each is an `_index.md` directly under `content/`.
- **Articles** — long-form posts under `content/articles/`. Flat layout, filenames are slugs (no date prefix), date comes from frontmatter.
- **Book reviews** — under `content/books/<slug>/index.md`. Bundle directories so the cover image (and optional `highlights.md` and `notes.md`) sit alongside the review. Frontmatter references one or more authors by slug. `currentlyReading: true` marks an unfinished read; `featuredOnHomepage: true` opts a book into the homepage rotation.
- **Authors** — `content/authors/<slug>/index.md`, referenced by books. Title is the author's full name.
- **Thinkers** — `content/thinkers/<slug>/index.md`. Each thinker lists the books of theirs I've read (by slug) in `books:`. A thinker with body content is considered "Completed"; without, "Pending". Sort key on the section page is `birth_date` (a signed year, negative for BCE).
- **Interviews** — `content/interviews/`, filenames `YYYY-MM-DD-<slug>.md`.
- **Presentations** — `content/presentations/<slug>/index.md`, bundle directory.

Images related to a specific piece of content live in the same bundle directory (e.g. `content/books/<slug>/<slug>.jpg`). Hugo resolves them via `.Resources.GetMatch`.

`scripts/validate.go` enforces required frontmatter (e.g. `amazonURL` on book reviews with content) and cross-reference integrity (e.g. a thinker's `books:` entries must resolve to real book slugs). Run it before committing content changes.
