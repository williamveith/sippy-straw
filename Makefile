# Define modules and directories
MODULE_STRAW = straw
MODULE_APP = appserver
SRC_DIR = Buildfiles
BIN_DIR = bin
GO_MOD_DIR = $(MODULE_STRAW)/$(SRC_DIR)

# Define file paths
BUILD_FILE = main.go
BUILD_PATH = $(GO_MOD_DIR)/$(BUILD_FILE)
BIN_PATH = $(GO_MOD_DIR)/$(BIN_DIR)/main

# Load straw/.env
ifneq (,$(wildcard $(CURDIR)/$(MODULE_STRAW)/.env))
    include $(CURDIR)/$(MODULE_STRAW)/.env
    export
endif

# Build for host system
compile:
	make clean
	GOFLAGS="-mod=mod" go build -o $(BIN_PATH) $(BUILD_PATH)

# Cross-compile for Windows
compile-windows:
	make clean
	GOOS=windows GOARCH=amd64 GOFLAGS="-mod=mod" go build -o $(BIN_PATH).exe $(BUILD_PATH)

# Cross-compile for macOS
compile-mac:
	make clean
	GOOS=darwin GOARCH=amd64 GOFLAGS="-mod=mod" go build -o $(BIN_PATH) $(BUILD_PATH)

# Cross-compile for Linux
compile-linux:
	make clean
	GOOS=linux GOARCH=amd64 GOFLAGS="-mod=mod" go build -o $(BIN_PATH) $(BUILD_PATH)

# Builder and service targets
builder:
	$(BIN_PATH) create-builder

builder-clean:
	$(BIN_PATH) clean-builder

nginx:
	$(BIN_PATH) build-nginx
	make builder-clean
	
cloudflared:
	$(BIN_PATH) build-cloudflared
	make builder-clean

certbot:
	$(BIN_PATH) build-certbot
	make builder-clean

# Image signing and verification
sign:
	$(BIN_PATH) sign-image $(IMAGE_TAG)

verify:
	$(BIN_PATH) verify-image $(IMAGE_TAG)

# Docker compose commands
start:
	docker-compose -f $(MODULE_APP)/docker-compose.yml up --build -d
	docker-compose -f $(MODULE_STRAW)/docker-compose.yml up --build -d
	make open

stop:
	docker-compose -f $(MODULE_APP)/docker-compose.yml down
	docker-compose -f $(MODULE_STRAW)/docker-compose.yml down

# Clean up binaries
clean:
	rm -f $(BIN_PATH) $(BIN_PATH).exe

# Open the site
open:
	open "https://$(DOMAIN_NAME)"

all:
	make compile
	make builder
	$(BIN_PATH) build-nginx
	$(BIN_PATH) build-cloudflared
	$(BIN_PATH) build-certbot
	make builder-clean
