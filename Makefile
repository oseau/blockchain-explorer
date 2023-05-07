dev: ## start dev server
	@$(MAKE) log.info MSG="================ DEV ================"
	@docker compose -f docker-compose.yml -f docker-compose.dev.yml up --build backend

.PHONY: login
login: ## login dev docker
	@$(MAKE) log.info MSG="================ DEV ================"
	@docker compose -f docker-compose.yml -f docker-compose.dev.yml exec backend bash

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

# https://github.com/zephinzer/godev/blob/62012ce006df8a3131ee93a74bcec2066405fb60/Makefile#L220
log.info:
	-@printf -- "\033[0;32m> [INFO] ${MSG}\033[0m\n"
