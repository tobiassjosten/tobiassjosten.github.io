# Claude instructions

## Tech stack

- Hugo generates the website
- Plain HTML/CSS/JS for frontend
- GitHub Pages hosts the site
- Go for utility scripts

## Development

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
