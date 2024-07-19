# https://tech.davis-hansson.com/p/make/
SHELL := bash
.ONESHELL:

BIN_DIR := $(GOPATH)/bin
BUILD_DIR := build
# man-pages is also defined in goreleaser.yml
MAN_DIR := man-pages
DOCS_DIR := docs
MARKDOWN_FILES := $(DOCS_DIR)/*.md
BINARY := icm

.PHONY: all
all: test lint build markdown

.PHONY: dummy-csv
dummy-csv:
	@echo 'AAA;my company;my city;my country' > data/file/owner.csv

.PHONY: test
test: dummy-csv
	go test ./...

.PHONY: lint
lint: dummy-csv
# See .golangci.yml
	golangci-lint run

.PHONY: build
build: dummy-csv
	export CGO_ENABLED=0; go build -o $(BUILD_DIR)/$(BINARY)

.PHONY: markdown
markdown: build
	rm $(MARKDOWN_FILES)
	./$(BUILD_DIR)/$(BINARY) doc markdown $(DOCS_DIR)

# Individual commands

.PHONY: format
format:
	gofumpt -l -w .

.PHONY: download-owners
download-owners: build
	./$(BUILD_DIR)/$(BINARY) download-owners -o data/file/owner.csv

.PHONY: man-page
man-page: build
	./$(BUILD_DIR)/$(BINARY) doc man $(DOCS_DIR)/$(MAN_DIR)/man1

.PHONY: install
install: build
	cp $(BUILD_DIR)/$(BINARY) $(BIN_DIR)/$(BINARY)

.PHONY: clean
clean:
	go clean -x -testcache
	rm -rf $(BUILD_DIR)
