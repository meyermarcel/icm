PACKAGES := $(shell go list ./... | grep -v /vendor)
BIN_DIR := $(GOPATH)/bin
BUILD_DIR := build
BINARY := iso6346

.PHONY: all
all: clean test lint build

.PHONY: clean
clean:
	go clean -x
	rm -rf release
	rm -rf $(BUILD_DIR)

.PHONY: fmt
fmt:
	goimports -w $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: test
test:
	go test $(PACKAGES)

.PHONY: lint
lint: 
	gometalinter --disable-all --enable=vet --enable=golint --enable=gotype --enable=goimports --vendor ./...

.PHONY: build
build:
	go build -o $(BUILD_DIR)/$(BINARY)

VERSION ?= vlatest
PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64

.PHONY: release
release: windows linux darwin

.PHONY: dep
dep:
	dep ensure

.PHONY: install
install:
	go build -o $(BIN_DIR)/$(BINARY)
