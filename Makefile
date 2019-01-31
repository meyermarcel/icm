PACKAGES := $(shell go list ./... | grep -v /vendor)
BIN_DIR := $(GOPATH)/bin
BUILD_DIR := build
# man-pages is also defined in goreleaser.yml
MAN_DIR := man-pages
DOCS_DIR := docs
BINARY := icm

.PHONY: all
all: dep clean test lint build man markdown

.PHONY: dep
dep:
	dep ensure -v

.PHONY: clean
clean: dep
	go clean -x -testcache
	rm -rf $(BUILD_DIR)

.PHONY: test
test: clean
	go test $(PACKAGES)

.PHONY: lint
lint: test
	gometalinter --disable-all --enable=vet --enable=golint --enable=gotype --enable=goimports --vendor ./...

.PHONY: build
build: lint
	export CGO_ENABLED=0; go build -tags osusergo -o $(BUILD_DIR)/$(BINARY)

.PHONY: man
man: build
	$(shell $(BUILD_DIR)/$(BINARY) misc man $(MAN_DIR)/man1)

.PHONY: markdown
markdown: build
	$(shell $(BUILD_DIR)/$(BINARY) misc markdown $(DOCS_DIR))

.PHONY: install
install:
	go build -o $(BIN_DIR)/$(BINARY)

.PHONY: fmt
fmt:
	goimports -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")