# Define the source file and directory
BUILD_FILE_NAME=main.go
SRC_DIR=Buildfiles
BUILD_PATH=$(SRC_DIR)/$(BUILD_FILE_NAME)

# Define output binary name and output directory
BINARY_NAME=main
BIN_DIR=bin
BUILDER_ROOT_DIR=Buildfiles
BIN_PATH=$(BIN_DIR)/$(BINARY_NAME)
MAIN_PATH=${BUILDER_ROOT_DIR}/${BIN_PATH}

# Set the working directory to Buildfiles where go.mod is located
GO_MOD_DIR=$(SRC_DIR)

# Build for host system
build:
	cd $(GO_MOD_DIR) && go build -o $(BIN_PATH) $(BUILD_FILE_NAME)

# Cross-compile for Windows
build-windows:
	cd $(GO_MOD_DIR) && GOOS=windows GOARCH=amd64 go build -o $(BIN_PATH).exe $(BUILD_FILE_NAME)

# Cross-compile for macOS
build-mac:
	cd $(GO_MOD_DIR) && GOOS=darwin GOARCH=amd64 go build -o $(BIN_PATH) $(BUILD_FILE_NAME)

# Cross-compile for Linux
build-linux:
	cd $(GO_MOD_DIR) && GOOS=linux GOARCH=amd64 go build -o $(BIN_PATH) $(BUILD_FILE_NAME)

builder:
	$(MAIN_PATH) create-builder

cloudflared:
	$(MAIN_PATH) build-cloudflared

certbot:
	$(MAIN_PATH) build-certbot

sign:
	${MAIN_PATH} sign-image $(IMAGE_TAG)

verify:
	${MAIN_PATH} verify-image $(IMAGE_TAG)

# Clean up build artifacts
clean:
	cd $(GO_MOD_DIR) && rm -f $(BIN_PATH) $(BIN_PATH).exe
