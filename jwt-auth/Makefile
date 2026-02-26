# Application Name
APP_NAME := jwt-auth

# Main entry point
CMD_PATH := ./cmd/api

# Binary output directory
BIN_DIR := bin

# Config file (default)
CONFIG := configs/local.yaml

# -----------------------------
# Run application (development)
# -----------------------------
run:
	@echo "Starting $(APP_NAME)..."
	@go run $(CMD_PATH)/main.go -config=$(CONFIG)

# -----------------------------
# Build binary
# -----------------------------
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)

# -----------------------------
# Run built binary
# -----------------------------
run-build: build
	@echo "Running built binary..."
	@./$(BIN_DIR)/$(APP_NAME) -config=$(CONFIG)

# -----------------------------
# Run tests
# -----------------------------
test:
	@echo "Running tests..."
	@go test -v ./...

# -----------------------------
# Format code
# -----------------------------
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# -----------------------------
# Vet code (static analysis)
# -----------------------------
vet:
	@echo "Running go vet..."
	@go vet ./...

# -----------------------------
# Clean build artifacts
# -----------------------------
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BIN_DIR)

# -----------------------------
# Build for Linux (production)
# -----------------------------
build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BIN_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)