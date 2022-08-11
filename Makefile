BIN_DIR := $(GOPATH)/bin
BUILD_DIR := build
# man-pages is also defined in goreleaser.yml
MAN_DIR := man-pages
DOCS_DIR := docs
BINARY := icm

.PHONY: all
all: test lint build

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run --enable revive --enable gofumpt --enable errorlint --enable godot --enable errname

.PHONY: build
build:
	export CGO_ENABLED=0; go build -o $(BUILD_DIR)/$(BINARY)


# Individual commands

.PHONY: build-docs
build-docs: build
	$(shell $(BUILD_DIR)/$(BINARY) doc man $(DOCS_DIR)/$(MAN_DIR)/man1)
	$(shell $(BUILD_DIR)/$(BINARY) doc markdown $(DOCS_DIR))

.PHONY: install
install: build
	cp $(BUILD_DIR)/$(BINARY) $(BIN_DIR)/$(BINARY)

.PHONY: clean
clean:
	go clean -x -testcache
	rm -rf $(BUILD_DIR)

.PHONY: format
format:
	gofumpt -l -w .

.PHONY: update-owners
update-owners: build
	echo '' > $(HOME)/.icm/data/owner.csv && ./$(BUILD_DIR)/$(BINARY) update && cp $(HOME)/.icm/data/owner.csv data/file/owner.csv