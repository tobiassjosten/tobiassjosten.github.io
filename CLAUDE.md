# Claude instructions

## Tech stack

- Hugo generates the website
- Plain HTML/CSS/JS for frontend
- GitHub Pages hosts the site
- Cloudflare as CDN and reverse proxy
- Go for utility scripts

## Development

Don't bother backing up files. Everything is tracked by Git and can easily be restored.

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

## Content

There are different types of content in the site:
- Pages — 
- Articles — 
- Tech Notes — 
- Book reviews —

Images related to specific articles or tech notes are stored in
