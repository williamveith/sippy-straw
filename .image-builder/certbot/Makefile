# Define the source file and directory
BUILD_FILE_NAME=entrypoint.go
SRC_DIR=src
BUILD_PATH=$(SRC_DIR)/$(BUILD_FILE_NAME)

# Define output binary name and output directory
BINARY_NAME=entrypoint
BIN_DIR=bin
BIN_PATH=$(BIN_DIR)/$(BINARY_NAME)

# Build for host system
build:
	go build -o $(BIN_PATH) $(BUILD_PATH)

# Cross-compile for Windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BIN_PATH).exe $(BUILD_PATH)

# Cross-compile for macOS
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_PATH) $(BUILD_PATH)

# Cross-compile for Linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_PATH) $(BUILD_PATH)

builder:
	$(BIN_PATH) create-builder

certbot:
	go run main.go

# Clean up build artifacts
clean:
	rm -f $(BIN_PATH) $(BIN_PATH).exe