MAKEFLAGS += --no-print-directory
SHELL := zsh
.DEFAULT_GOAL := help

.PHONY: build
build: 
	CGO_ENABLED=1 goreleaser build --single-target

.PHONY: release
releaser: 
	goreleaser release

.PHONY: help
help: ## List of the aivailable commands
	@echo "MEDELSE COMMANDS"
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z1-9_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""
