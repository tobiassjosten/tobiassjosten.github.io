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
