include .env
ifneq (,$(wildcard ./.env.local))
        include .env.local
endif
export

.PHONY: serve
serve:
	@echo "running server"
	@hugo --cleanDestinationDir server

.PHONY: deploy
deploy:
	@echo "deploying to cloudflare pages"
	@BRANCH=$$(git branch --show-current) ./scripts/deploy-cloudflare.sh

.PHONY: validate
validate:
	@echo "validating site"
	@go test ./tests/...

.PHONY: highlights
highlights:
	@go run scripts/download-highlights.go $(filter-out highlights,$(MAKECMDGOALS))

.PHONY: new
new:
	@go run scripts/create-content.go $(filter-out new,$(MAKECMDGOALS))

%:
	@:
