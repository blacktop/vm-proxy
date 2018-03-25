REPO=blacktop
NAME=vm-proxy
VERSION=$(shell cat VERSION)
MESSAGE?="New release ${VERSION}"

# TODO remove \|/templates/\|/api
SOURCE_FILES?=$$(go list ./... | grep -v /vendor/)
TEST_PATTERN?=.
TEST_OPTIONS?=

GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GIT_DESCRIBE=$(git describe --tags)


setup: ## Install all the build and lint dependencies
	@echo "===> Installing deps"
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/golang/dep/...
	go get -u github.com/pierrre/gotestcover
	go get -u golang.org/x/tools/cmd/cover
	dep ensure
	gometalinter --install

.PHONY: run.server
run.server: ## Run vm-proxy server
	@echo "===> Running vm-proxy server..."
	go run server/*.go -V

.PHONY: test
test: ## Run all the tests
	gotestcover $(TEST_OPTIONS) -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=30s

.PHONY: test.release
test.release: ## Run all the tests
	 goreleaser release --rm-dist --skip-publish --skip-validate

cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

lint: ## Run all the linters
	gometalinter --vendor --disable-all \
		--enable=deadcode \
		--enable=ineffassign \
		--enable=gosimple \
		--enable=staticcheck \
		--enable=gofmt \
		--enable=goimports \
		--enable=dupl \
		--enable=misspell \
		--enable=errcheck \
		--enable=vet \
		--enable=vetshadow \
		--deadline=10m \
		./...
		markdownfmt -w README.md

.PHONY: release
release: ## Create a new release from the VERSION
	@echo "===> Creating Release"
	@hack/make/release ${VERSION}
	@goreleaser --rm-dist

ci: lint test ## Run all the tests and code checks

build: ## Build a beta version of vm-proxy
	@echo "===> Building Binaries"
	CGO_ENABLED=0 go build -buildmode=pie -i -o $(NAME) -ldflags="-w -s" server/*.go

.PHONY: build.vbox
build.vbox: ## Build a beta version of vbox client
	@echo "===> Building vbox client"
	docker build -t $(REPO)/vbox:dev clients/vbox/

.PHONY: ssh.vbox
ssh.vbox: ## Build a beta version of vbox client
	@echo "===> sshing into vbox client"
	docker run -it --rm --entrypoint=sh $(REPO)/vbox:dev

clean: ## Clean up artifacts
	@rm -rf $(HOME)/.vmproxy || true
	@rm -rf dist/ || true
	@rm vm-proxy || true

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help