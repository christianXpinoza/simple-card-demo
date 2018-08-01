VERSION = 0.0.1
BUILD = $(shell git rev-parse --short=6 HEAD)
LDFLAGS = -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"
DOCKER_REPO = chespinoza/simple-card-demo
BUILD_DIR = build

default: help

.PHONY: build 
build: ## Build binaries for current host
	mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/simple-card-demo .

.PHONY: clean 
clean: ## Clean files generated on build
	rm -rf build/*

.PHONY: docker 
docker: ## Build binaries into docker container
	docker build --force-rm -t $(DOCKER_REPO):$(VERSION) --build-arg build=$(BUILD) --build-arg version=$(VERSION) .

.PHONY: help
help: ## Help
	@echo "Please use 'make <target>' where <target> is ..."
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
