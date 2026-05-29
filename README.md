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

Validate content:

```bash
go run scripts/validate.go
```

See `CLAUDE.md` for the content model and conventions.
