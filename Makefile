GOCMD=go
GOBUILD=CGO_ENABLED=0 GOARCH=amd64 $(GOCMD) build
GOCLEAN=$(GOCMD) clean -testcache
GOTEST=$(GOCMD) test ./...

BINARY_NAME=bin/gol

all: build test
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST)
clean:
	$(GOCLEAN)
run: build
	./$(BINARY_NAME)

.PHONY: all test clean
