# https://tech.davis-hansson.com/p/make/
SHELL := bash
.ONESHELL:

BIN_DIR := $(GOPATH)/bin
BUILD_DIR := build
# man-pages is also defined in goreleaser.yml
MAN_DIR := man-pages
DOCS_DIR := docs
BINARY := icm

.PHONY: all
all: test lint build markdown

.PHONY: init-csv
init-csv:
	[ -f data/file/owner.csv ] || echo 'AAA;my company;my city;my country' > data/file/owner.csv

.PHONY: test
test: init-csv
	go test ./...

.PHONY: lint
lint: init-csv
	golangci-lint run --enable revive --enable gofumpt --enable errorlint --enable godot --enable errname

.PHONY: build
build: init-csv
	export CGO_ENABLED=0; go build -o $(BUILD_DIR)/$(BINARY)

.PHONY: markdown
markdown: build
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
