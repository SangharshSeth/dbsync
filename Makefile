# Name of the executable
APP_NAME := dbsync

# Source directories
SRC_DIR := ./

# Go flags
GO_FLAGS := -v

# Default target - build the project
all: build

# Build the Go project
build:
	@echo "Building the project..."
	@go build $(GO_FLAGS) -o $(APP_NAME) $(SRC_DIR)

# Run the Go project
run: build
	@echo "Running the project..."
	@./$(APP_NAME)

# Test the Go project
test:
	@echo "Running tests..."
	@go test $(GO_FLAGS) $(SRC_DIR)

# Clean build files
clean:
	@echo "Cleaning build files..."
	@rm -f $(APP_NAME)

# Clean and rebuild the project
rebuild: clean build

# Format the Go code
fmt:
	@echo "Formatting code..."
	@go fmt $(SRC_DIR)

# Display help
help:
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all      - Default target, builds the project"
	@echo "  build    - Build the Go project"
	@echo "  run      - Run the Go project"
	@echo "  test     - Test the Go project"
	@echo "  clean    - Clean build files"
	@echo "  rebuild  - Clean and then build the project"
	@echo "  fmt      - Format the Go code"
	@echo "  help     - Display this help message"

.PHONY: all build run test clean rebuild fmt help