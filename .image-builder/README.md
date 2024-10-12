# Image Builder Tool

This project includes a build tool for creating and managing image builds across different platforms. The build process is automated using a `Makefile` to handle the compilation of the tool, as well as specific image build tasks.

## Prerequisites

Make sure you have the following installed on your system:

- **Go**: The Go programming language (version 1.23.2 or higher recommended)
- **Docker**: Required for building and managing container images

## Makefile Targets

### Build for Host System

To compile the image builder tool for your current system, use:

```bash
make build
```

This will compile the Go source file `src/main.go` and place the resulting binary in the `bin/` directory as `bin/main`.

### Cross-Compilation

The tool can be cross-compiled for different operating systems, which is useful for building images on various platforms:

- **Windows**:

  ```sh
  make build-windows
  ```

  This will generate a `bin/main.exe` binary for Windows.

- **macOS**:

  ```sh
  make build-mac
  ```

  This will generate a `bin/main` binary for macOS.

- **Linux**:

  ```sh
  make build-linux
  ```

  This will generate a `bin/main` binary for Linux.

### Image Build Commands

The following commands are used to trigger specific image-building tasks:

- **Create Builder**:

  ```bash
  make builder
  ```

  This command runs the tool to create a new docker builder for building images.

- **Build Cloudflared**:

  ```bash
  make cloudflared
  ```

  This command uses the image builder tool to build the Cloudflared image.

### Cleaning Up

To remove the generated binaries and clean up the build directory, run:

```bash
make clean
```

This will remove all compiled binaries (`bin/main` and `bin/main.exe`).

## Directory Structure

- `src/`: Contains the Go source files.
- `bin/`: Output directory where the compiled binaries are placed.
