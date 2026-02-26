# =====================================================
# Application Configuration
# =====================================================

APP_NAME := jwt-auth
CMD_PATH := ./cmd/api
BIN_DIR  := bin
CONFIG   := configs/local.yaml

# =====================================================
# Declare phony targets
# =====================================================

.PHONY: run build run-build test fmt vet clean build-linux

# =====================================================
# Run application (development)
# =====================================================

run:
	@echo "Starting $(APP_NAME)..."
	@exec go run $(CMD_PATH)/main.go -config=$(CONFIG)

# =====================================================
# Build binary
# =====================================================

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)

# =====================================================
# Run built binary
# =====================================================

run-build: build
	@echo "Running built binary..."
	@exec ./$(BIN_DIR)/$(APP_NAME) -config=$(CONFIG)

# =====================================================
# Run tests
# =====================================================

test:
	@echo "Running tests..."
	@go test -v ./...

# =====================================================
# Format code
# =====================================================

fmt:
	@echo "Formatting code..."
	@go fmt ./...

# =====================================================
# Static analysis
# =====================================================

vet:
	@echo "Running go vet..."
	@go vet ./...

# =====================================================
# Clean build artifacts
# =====================================================

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BIN_DIR)

# =====================================================
# Build for Linux (production release)
# =====================================================

build-linux:
	@echo "Building Linux binary..."
	@mkdir -p $(BIN_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)