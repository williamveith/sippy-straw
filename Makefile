# Define the source file and directory
MODULE_STRAW=straw
MODULE_APP=appserver

SRC_DIR=Buildfiles
BUILD_FILE_NAME=main.go
BUILD_PATH=${MODULE_STRAW}/$(SRC_DIR)/$(BUILD_FILE_NAME)

# Define output binary name and output directory
BINARY_NAME=main
BIN_DIR=bin
BUILDER_ROOT_DIR=Buildfiles
BIN_PATH=$(BIN_DIR)/$(BINARY_NAME)
MAIN_PATH=$(BUILDER_ROOT_DIR)/$(BIN_PATH)

# Set the working directory to Buildfiles where go.mod is located
GO_MOD_DIR=$(MODULE_STRAW)/$(SRC_DIR)

# Load straw/.env
ifneq (,$(wildcard $(MODULE_STRAW)/.env))
    include $(MODULE_STRAW)/.env
    export
endif

# Build for host system
compile:
	cd $(GO_MOD_DIR) && go build -o $(BIN_PATH) $(BUILD_FILE_NAME)

# Cross-compile for Windows
compile-windows:
	cd $(GO_MOD_DIR) && GOOS=windows GOARCH=amd64 go build -o $(BIN_PATH).exe $(BUILD_FILE_NAME)

# Cross-compile for macOS
compile-mac:
	cd $(GO_MOD_DIR) && GOOS=darwin GOARCH=amd64 go build -o $(BIN_PATH) $(BUILD_FILE_NAME)

# Cross-compile for Linux
compile-linux:
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

start:
	docker-compose -f $(MODULE_APP)/docker-compose.yml up --build -d
	docker-compose -f $(MODULE_STRAW)/docker-compose.yml up --build -d
	make open

stop:
	docker-compose -f $(MODULE_APP)/docker-compose.yml down
	docker-compose -f $(MODULE_STRAW)/docker-compose.yml down

clean:
	cd $(GO_MOD_DIR) && rm -f $(BIN_PATH) $(BIN_PATH).exe

open:
	open "https://$(HOST_NAME)"
