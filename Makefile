PACKAGES := $(shell go list ./... | grep -v /vendor)
BIN_DIR := $(GOPATH)/bin
BUILD_DIR := build
BINARY := icm

.PHONY: all
all: dep clean test lint build

.PHONY: clean
clean:
	go clean -x
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
	export CGO_ENABLED=0; go build -tags osusergo -o $(BUILD_DIR)/$(BINARY)

.PHONY: dep
dep:
	dep ensure -v

.PHONY: install
install:
	go build -o $(BIN_DIR)/$(BINARY)
