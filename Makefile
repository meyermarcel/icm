PACKAGES := $(shell go list ./... | grep -v /vendor)
BIN_DIR := $(GOPATH)/bin
BINARY := iso6346

.PHONY: all
all: clean lint test install

.PHONY: clean
clean:
	go clean -x
	rm -rf release

GOMETALINTER := $(BIN_DIR)/gometalinter

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

.PHONY: lint
lint: $(GOMETALINTER)
	gometalinter --disable-all --enable=golint --enable=gotype --enable=goimports --enable=errcheck  --vendor ./...

.PHONY: test
test:
	go test $(PACKAGES)

.PHONY: build
build:
	go build -o $(BINARY)

.PHONY: install
install:
	go build -o $(BIN_DIR)/$(BINARY)

VERSION ?= vlatest
PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64

.PHONY: release
release: windows linux darwin
