package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

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
	alpineVersion := os.Getenv("CLOUDFLARED_ALPINE_VERSION")
	cloudflaredVersion := os.Getenv("CLOUDFLARED_VERSION")
	IMAGE_TAG := fmt.Sprintf("williamveith/cloudflared:%s", cloudflaredVersion)

	cmd := exec.Command("docker", "buildx", "build",
		"--platform", "linux/arm64,linux/amd64,linux/386,linux/arm/v7,linux/arm/v6",
		"--tag", IMAGE_TAG,
		"--builder", builderName,
		"--push",
		"--file", "cloudflared/src/Dockerfile.cloudflare",
		"--build-arg", fmt.Sprintf("CLOUDFLARED_ALPINE_VERSION=%s", alpineVersion),
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
	signImage(IMAGE_TAG)
}

func buildCertbotImage() {
	certbotGolangVersion := os.Getenv("CERTBOT_GOLANG_VERSION")
	certbotImageVersion := os.Getenv("CERTBOT_IMAGE_VERSION")
	IMAGE_TAG := fmt.Sprintf("williamveith/certbot:%s", certbotImageVersion)

	cmd := exec.Command("docker", "buildx", "build",
		"--platform", "linux/amd64,linux/arm/v6,linux/arm64",
		"--build-arg", fmt.Sprintf("CERTBOT_GOLANG_VERSION=%s", certbotGolangVersion),
		"--build-arg", fmt.Sprintf("CERTBOT_IMAGE_VERSION=%s", certbotImageVersion),
		"--tag", IMAGE_TAG,
		"--file", "Dockerfiles/Dockerfile.certbot",
		"--push",
		".")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Build failed with error:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("Build completed successfully!")
	fmt.Println(string(output))
	signImage(IMAGE_TAG)
}

func signImage(IMAGE_TAG string) {
	fmt.Println("Signing Image...")
	KEY_PATH := os.Getenv("KEY_PATH")
	if KEY_PATH == "" {
		KEY_PATH = "Buildfiles"
	}

	// Set COSIGN_PASSWORD to an empty string to bypass the password prompt
	os.Setenv("COSIGN_PASSWORD", "")
	// Enable experimental features for Tlog signing
	os.Setenv("COSIGN_EXPERIMENTAL", "1")

	// Create command to sign the image
	cmd := exec.Command("cosign", "sign",
		"--key", fmt.Sprintf("%s/cosign.key", KEY_PATH),
		IMAGE_TAG)

	// Simulate typing 'y' for Tlog upload confirmation
	cmd.Stdin = strings.NewReader("y\n")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Could not sign image:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("Image Signed Successfully!")
	fmt.Println(string(output))
}

func verifyImage(IMAGE_TAG string) {
	KEY_PATH := "Buildfiles"
	cmd := exec.Command("cosign", "verify",
		"--key", fmt.Sprintf("%s/cosign.pub", KEY_PATH),
		IMAGE_TAG)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Could not verify image:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("Build completed successfully!")
	fmt.Println(string(output))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: create-builder, build-cloudflared, or build-certbot")
		return
	}

	err := godotenv.Load(".env")
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
	case "build-certbot":
		fmt.Println("Building certbot...")
		buildCertbotImage()
	case "verify-image":
		if len(os.Args) < 3 {
			fmt.Println("Please provide an image tag to verify.")
			return
		}
		verifyImage(os.Args[2])
	case "sign-image":
		if len(os.Args) < 3 {
			fmt.Println("Please provide an image tag to sign.")
			return
		}
		signImage(os.Args[2])
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
