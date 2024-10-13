package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func createCustomBuilder() {
	builderName := os.Getenv("BUILDER_NAME")
	image := os.Getenv("BUILDER_IMAGE")

	cmd := exec.Command("docker", "buildx", "create",
		"--name", builderName,
		"--driver", "docker-container",
		"--driver-opt", fmt.Sprintf("image=%s", image))

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Build failed with error:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("Build completed successfully!")
	fmt.Println(string(output))
}

func buildCloudflareImage() {
	builderName := os.Getenv("BUILDER_NAME")
	alpineVersion := os.Getenv("ALPINE_VERSION")
	cloudflaredVersion := os.Getenv("CLOUDFLARED_VERSION")

	cmd := exec.Command("docker", "buildx", "build",
		"--platform", "linux/arm64,linux/amd64,linux/386,linux/arm/v7,linux/arm/v6",
		"--tag", fmt.Sprintf("williamveith/cloudflared:%s", cloudflaredVersion),
		"--builder", builderName,
		"--push",
		"--file", "cloudflared/src/Dockerfile.cloudflare",
		"--build-arg", fmt.Sprintf("ALPINE_VERSION=%s", alpineVersion),
		"--build-arg", fmt.Sprintf("CLOUDFLARED_VERSION=%s", cloudflaredVersion),
		".")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Build failed with error:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("Build completed successfully!")
	fmt.Println(string(output))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: create-builder or build-cloudflared")
		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	switch os.Args[1] {
	case "create-builder":
		fmt.Println("Creating builder...")
		createCustomBuilder()
	case "build-cloudflared":
		fmt.Println("Building cloudflared...")
		buildCloudflareImage()
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
