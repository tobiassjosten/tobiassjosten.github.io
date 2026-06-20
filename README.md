# tobiassjosten.net

Source for my personal site at [tobiassjosten.net](https://tobiassjosten.net). Built with [Hugo](https://gohugo.io), deployed to Cloudflare Pages.

## Local development

```bash
hugo --environment production \
     --destination public \
     --cleanDestinationDir \
     --gc \
     --minify \
     --panicOnWarning \
     --printPathWarnings
```

Validate the whole site (builds it and asserts internal links resolve, covers
render, and content frontmatter is valid):

```bash
go test ./tests/...
```

See `CLAUDE.md` for the content model and conventions.
