.DEFAULT_GOAL := build

# Format the code
fmt:
	go fmt ./...
.PHONY: fmt

# Lint the code
lint: fmt
	golint ./...
.PHONY: lint

# Vet the code
vet: fmt
	go vet ./...
.PHONY: vet

# Build the project
build:
	go build .
.PHONY: build

# Clean the build artifacts
clean:
	go clean
.PHONY: clean

