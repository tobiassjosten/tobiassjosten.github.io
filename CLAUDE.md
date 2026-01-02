# Claude instructions

## Tech stack

- Hugo generates the website
- Plain HTML/CSS/JS for frontend
- GitHub Pages hosts the site
- Cloudflare as CDN and reverse proxy
- Go for utility scripts

## Development

Don't bother backing up files. Everything is tracked by Git and can easily be restored.

Only create new utility scripts when explicitly told to do so. Instead rely on shell commands and Posix tools for ad-hoc tasks.

Default to not commenting code. Instead use clear and descriptive names for variables, functions, files, and directories.

For CSS and styling, rely on Bootstrap 5, using its built-in extension points like variables. Use custom CSS and code only when necessary.

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

There are different types of content in the site:
- Pages — 
- Articles — 
- Tech Notes — 
- Book reviews —

Images related to specific articles or tech notes are stored in
