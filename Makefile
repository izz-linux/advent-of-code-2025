.PHONY: all clean run build test fmt vet help update-readme init-modules

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=gofmt
GOVET=$(GOCMD) vet
GORUN=$(GOCMD) run

# Find all day directories
DAYS := $(wildcard day*)

# Default target
all: fmt vet run-all

# Help target
help:
	@echo "Advent of Code 2025 - Available Make Targets:"
	@echo ""
	@echo "  make all          - Format, vet, and run all solutions"
	@echo "  make run-all      - Run all day solutions"
	@echo "  make run-day01    - Run specific day (e.g., day01, day02, etc.)"
	@echo "  make build        - Build all day binaries"
	@echo "  make build-day01  - Build specific day binary"
	@echo "  make test         - Run all tests"
	@echo "  make fmt          - Format all Go code"
	@echo "  make vet          - Run go vet on all packages"
	@echo "  make clean        - Remove all built binaries"
	@echo "  make update-readme - Update README with solution links"
	@echo "  make init-modules  - Initialize go.mod in all day directories"
	@echo ""

# Initialize go.mod in all day directories
init-modules:
	@echo "Initializing Go modules in day directories..."
	@for dir in $(DAYS); do \
		if [ -f $$dir/main.go ] && [ ! -f $$dir/go.mod ]; then \
			echo "Initializing module in $$dir..."; \
			cd $$dir && $(GOCMD) mod init github.com/izz-linux/advent-of-code-2025/$$dir; \
			cd ..; \
		elif [ -f $$dir/go.mod ]; then \
			echo "$$dir already has go.mod"; \
		fi \
	done
	@echo "Module initialization complete!"

# Run all solutions
run-all:
	@echo "Running all Advent of Code solutions..."
	@for dir in $(DAYS); do \
		if [ -f $$dir/main.go ]; then \
			echo ""; \
			echo "=== Running $$dir ==="; \
			cd $$dir && $(GORUN) main.go || exit 1; \
			cd ..; \
		fi \
	done
	@echo ""
	@echo "All solutions completed!"

# Run specific day (e.g., make run-day01)
run-day%:
	@if [ -f day$*/main.go ]; then \
		echo "Running day$*..."; \
		cd day$* && $(GORUN) main.go; \
	else \
		echo "Error: day$*/main.go not found"; \
		exit 1; \
	fi

# Build all binaries
build:
	@echo "Building all solutions..."
	@for dir in $(DAYS); do \
		if [ -f $$dir/main.go ]; then \
			echo "Building $$dir..."; \
			cd $$dir && $(GOBUILD) -o ../bin/$$dir main.go || exit 1; \
			cd ..; \
		fi \
	done
	@echo "Build complete! Binaries in ./bin/"

# Build specific day (e.g., make build-day01)
build-day%:
	@if [ -f day$*/main.go ]; then \
		echo "Building day$*..."; \
		mkdir -p bin; \
		cd day$* && $(GOBUILD) -o ../bin/day$* main.go; \
		echo "Binary created: bin/day$*"; \
	else \
		echo "Error: day$*/main.go not found"; \
		exit 1; \
	fi

# Run tests
test:
	@echo "Running tests..."
	@for dir in $(DAYS); do \
		if [ -f $$dir/main_test.go ]; then \
			echo "Testing $$dir..."; \
			cd $$dir && $(GOTEST) -v || exit 1; \
			cd ..; \
		fi \
	done

# Format all Go code
fmt:
	@echo "Formatting code..."
	@for dir in $(DAYS); do \
		if [ -f $$dir/main.go ]; then \
			$(GOFMT) -w $$dir/*.go; \
		fi \
	done
	@$(GOFMT) -w *.go 2>/dev/null || true
	@echo "Formatting complete!"

# Run go vet
vet:
	@echo "Running go vet..."
	@for dir in $(DAYS); do \
		if [ -f $$dir/main.go ]; then \
			echo "Vetting $$dir..."; \
			cd $$dir && $(GOVET) . 2>/dev/null || echo "  (skipped - no module)"; \
			cd ..; \
		fi \
	done
	@echo "Vet complete!"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@for dir in $(DAYS); do \
		if [ -d $$dir ]; then \
			cd $$dir && $(GOCLEAN) || true; \
			cd ..; \
		fi \
	done
	@echo "Clean complete!"

# Update README with latest solutions
update-readme:
	@echo "Updating README..."
	@$(GORUN) update-readme.go
	@echo "README updated!"

# Create new day directory (e.g., make new-day DAY=04)
new-day:
	@if [ -z "$(DAY)" ]; then \
		echo "Error: Please specify DAY (e.g., make new-day DAY=04)"; \
		exit 1; \
	fi; \
	DAY_DIR=day$(DAY); \
	if [ -d $$DAY_DIR ]; then \
		echo "Error: $$DAY_DIR already exists"; \
		exit 1; \
	fi; \
	mkdir -p $$DAY_DIR; \
	echo 'package main\n\nimport (\n\t"fmt"\n)\n\nfunc main() {\n\tfmt.Println("Day $(DAY) - Part 1")\n\t// Your solution here\n}' > $$DAY_DIR/main.go; \
	touch $$DAY_DIR/input.txt; \
	echo "Created $$DAY_DIR with template files"; \
	$(MAKE) update-readme
