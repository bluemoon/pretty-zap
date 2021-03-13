GOCMD=go
GOPRIVATE=github.com/bluemoon/*
GOBUILD=$(GOCMD) build
GOENV=$(GOCMD) env
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFILES=.
BINARY_NAME=pretty-zap-test
BUILD_DIR=bin
GOX_OSARCH ?= "darwin/amd64"

all: dev-dependencies mod-download tidy build

tidy: ## Cleans the Go module.
	@echo "==> Tidying module"
	@go mod tidy
.PHONY: tidy

mod-download: ## Downloads the Go module.
	@echo "==> Downloading Go module"
	@go env -w GOPRIVATE=$(GOPRIVATE)
	@go mod download -x
.PHONY: mod-download

dev-dependencies: ## Downloads the necessary dev dependencies.
	@echo "==> Downloading development dependencies"
	@go get github.com/ahmetb/govvv
.PHONY: dev-dependencies

build: test
	@echo "==> Building $(GOFILES)"
	$(GOBUILD) -o bin/$(BINARY_NAME) -v $(GOFILES)

lint: mod-download
	@golangci-lint --version
	@golangci-lint run -v ./...

test: mod-download
	@go test -race -v ./...

clean:
	rm -rfv $(BUILD_DIR)
