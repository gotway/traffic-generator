PROJECT := traffic
VERSION := $(git describe --abbrev=0 --tags)
LD_FLAGS := -X main.version=$(VERSION) -s -w
SOURCE_FILES ?= ./internal/... ./pkg/... ./cmd/...
UNAME := $(uname -s)

export CGO_ENABLED := 0
export GO111MODULE := on
export GOBIN := $(shell pwd)/bin

.PHONY: all
all: help

.PHONY: help
help:	### Show targets documentation
ifeq ($(UNAME), Linux)
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
else
	@awk -F ':.*###' '$$0 ~ FS {printf "%15s%s\n", $$1 ":", $$2}' \
		$(MAKEFILE_LIST) | grep -v '@awk' | sort
endif

GOLANGCI_LINT := $(GOBIN)/golangci-lint
GOLANGCI_LINT_VERSION := v1.46.2
golangci-lint:
	$(call go-install,github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION))

GORELEASER := $(GOBIN)/goreleaser
GORELEASER_VERSION := v1.9.2
goreleaser:
	$(call go-install,github.com/goreleaser/goreleaser@$(GORELEASER_VERSION))

.PHONY: clean
clean: ### Clean build files
	@rm -rf ./bin
	@go clean

.PHONY: build
build: clean ### Build binary
	@go build -tags netgo -a -v -ldflags "${LD_FLAGS}" -o ./bin/traffic ./cmd/traffic/*.go
	@chmod +x ./bin/*

.PHONY: run
run: ### Quick run
	@CGO_ENABLED=1 go run -race cmd/traffic/*.go

.PHONY: deps
deps: ### Optimize dependencies
	@go mod tidy

.PHONY: install
install: ### Install binary in your system
	@go install -v cmd/traffic/*.go

.PHONY: lint
lint: golangci-lint ### Lint
	$(GOLANGCI_LINT) run

.PHONY: release
release: goreleaser ### Dry-run release
	$(GORELEASER) release --snapshot --rm-dist

.PHONY: test-clean
test-clean: ### Clean test cache
	@go clean -testcache ./...

.PHONY: test
test: lint ### Run tests
	@go test -v  -coverprofile=cover.out -timeout 10s ./...

.PHONY: cover
cover: test ### Run tests and generate coverage
	@go tool cover -html=cover.out -o=cover.html

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-install
@[ -f $(1) ] || { \
go install $(1) ; \
}
endef
