# Binary name
BINARY_NAME=server

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## build: Build the binary
build:
	@echo "Building..."
	go build -o $(GOBIN)/$(BINARY_NAME) ./cmd

## run: Build and run the binary
run: build
	@echo "Running..."
	$(GOBIN)/$(BINARY_NAME)

## clean: Clean build files
clean:
	@echo "Cleaning..."
	go clean
	rm -rf $(GOBIN)

## test: Run tests
test:
	@echo "Testing..."
	go test ./... -v

## fmt: Format code
fmt:
	@echo "Formatting..."
	go fmt ./...

## vet: Run go vet
vet:
	@echo "Vetting..."
	go vet ./...

## help: Display this help screen
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
