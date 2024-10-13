package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read necessary environment variables
	GOLANG_IMAGE_VERSION := os.Getenv("GOLANG_IMAGE_VERSION")
	if GOLANG_IMAGE_VERSION == "" {
		GOLANG_IMAGE_VERSION = "latest"
	}

	// Read necessary environment variables
	BASE_IMAGE_VERSION := os.Getenv("BASE_IMAGE_VERSION")
	if BASE_IMAGE_VERSION == "" {
		BASE_IMAGE_VERSION = "latest"
	}

	DOMAIN_NAME := os.Getenv("DOMAIN_NAME")
	if DOMAIN_NAME == "" {
		log.Fatal("DOMAIN_NAME is not set in .env file")
	}

	EMAIL_ADDRESS := os.Getenv("EMAIL_ADDRESS")
	if EMAIL_ADDRESS == "" {
		log.Fatal("EMAIL_ADDRESS is not set in .env file")
	}

	// Define the paths
	dockerfilePath := filepath.Join("src", "Dockerfile")
	entrypointPath := filepath.Join("src", "entrypoint.go")

	// Check if Dockerfile and entrypoint.go exist
	if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
		log.Fatalf("Dockerfile not found at %s", dockerfilePath)
	}
	if _, err := os.Stat(entrypointPath); os.IsNotExist(err) {
		log.Fatalf("entrypoint.go not found at %s", entrypointPath)
	}

	// Build the Docker image
	cmd := exec.Command("docker", "buildx", "build",
		"--file", dockerfilePath,
		"--tag", fmt.Sprintf("williamveith/certbot:%s", BASE_IMAGE_VERSION),
		"--builder", "multiplatform",
		"--load",
		"--build-arg", fmt.Sprintf("GOLANG_IMAGE_VERSION=%s", GOLANG_IMAGE_VERSION),
		"--build-arg", fmt.Sprintf("BASE_IMAGE_VERSION=%s", BASE_IMAGE_VERSION),
		"--build-arg", fmt.Sprintf("DOMAIN_NAME=%s", DOMAIN_NAME),
		"--build-arg", fmt.Sprintf("EMAIL_ADDRESS=%s", EMAIL_ADDRESS),
		".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error building Docker image: %v", err)
	}

	fmt.Println("Docker image built successfully!")
}
