package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func createCustomBuilder() {
	builderName := os.Getenv("BUILDER_NAME")
	image := os.Getenv("BUILDER_IMAGE")

	exists := checkBuilderExists(builderName)

	if exists {
		removeBuilder(builderName)
	}

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

func checkBuilderExists(builderName string) bool {
	cmd := exec.Command("docker", "buildx", "inspect", builderName)
	output, err := cmd.CombinedOutput()
	if err != nil && strings.Contains(string(output), "not found") {
		return false
	}
	return true
}

func removeBuilder(builderName string) {
	fmt.Printf("Removing existing builder: %s\n", builderName)
	cmd := exec.Command("docker", "buildx", "rm", builderName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "no builder") {
			fmt.Println("No existing builder found, skipping removal.")
		} else {
			fmt.Println("Failed to remove builder:", err)
			fmt.Println(string(output))
		}
	} else {
		fmt.Println("Builder removed successfully!")
	}
}

func buildNginxImage() {
	imageTag := "williamveith/nginx:mainline-alpine3.20-slim"

	cmd := exec.Command(
		"docker", "buildx", "imagetools", "create",
		"--tag", imageTag,
		"nginx:mainline-alpine3.20-slim",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to build and push the image:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("Image pulled, re-tagged, and pushed successfully!")
	fmt.Println(string(output))

	signImage(imageTag)
}

func buildCloudflareImage() {
	builderName := os.Getenv("BUILDER_NAME")
	alpineVersion := os.Getenv("CLOUDFLARED_ALPINE_VERSION")
	cloudflaredVersion := os.Getenv("CLOUDFLARED_VERSION")
	IMAGE_TAG := fmt.Sprintf("williamveith/cloudflared:%s", cloudflaredVersion)

	projectRoot := findProjectRoot()

	// Construct the paths for Dockerfile and build context
	dockerfilePath := filepath.Join(projectRoot, "straw", "Dockerfiles", "Dockerfile.cloudflare")
	buildContext := filepath.Join(projectRoot, "straw", "cloudflared")

	cmd := exec.Command("docker", "buildx", "build",
		"--platform", "linux/arm64,linux/amd64,linux/386,linux/arm/v7,linux/arm/v6",
		"--tag", IMAGE_TAG,
		"--builder", builderName,
		"--push",
		"--file", dockerfilePath,
		"--build-arg", fmt.Sprintf("CLOUDFLARED_ALPINE_VERSION=%s", alpineVersion),
		"--build-arg", fmt.Sprintf("CLOUDFLARED_VERSION=%s", cloudflaredVersion),
		buildContext)

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

	projectRoot := findProjectRoot()

	dockerfilePath := filepath.Join(projectRoot, "straw", "Dockerfiles", "Dockerfile.certbot")
	buildContext := filepath.Join(projectRoot, "straw", "certbot")

	cmd := exec.Command("docker", "buildx", "build",
		"--platform", "linux/amd64,linux/arm/v6,linux/arm64",
		"--build-arg", fmt.Sprintf("CERTBOT_GOLANG_VERSION=%s", certbotGolangVersion),
		"--build-arg", fmt.Sprintf("CERTBOT_IMAGE_VERSION=%s", certbotImageVersion),
		"--tag", IMAGE_TAG,
		"--file", dockerfilePath,
		"--push",
		buildContext)

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

	projectRoot := findProjectRoot()
	keyPath := filepath.Join(projectRoot, "straw", "Buildfiles")

	os.Setenv("COSIGN_PASSWORD", "")
	os.Setenv("COSIGN_EXPERIMENTAL", "1")

	cmd := exec.Command("cosign", "sign",
		"--key", fmt.Sprintf("%s/cosign.key", keyPath),
		IMAGE_TAG)

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
	projectRoot := findProjectRoot()
	keyPath := filepath.Join(projectRoot, "straw", "Buildfiles")

	cmd := exec.Command("cosign", "verify",
		"--key", fmt.Sprintf("%s/cosign.pub", keyPath),
		IMAGE_TAG)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Could not verify image:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println("Image was verified successfully!")
	fmt.Println(string(output))
}

func loadEnv() {
	projectRoot := findProjectRoot()

	envPath := filepath.Join(projectRoot, "straw", ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Error loading .env file: %v", err)
	} else {
		fmt.Println(".env file loaded successfully from:", envPath)
	}
}

func findProjectRoot(startDir ...string) string {
	var dir string
	if len(startDir) == 0 || startDir[0] == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		dir = wd
	} else {
		dir = startDir[0]
	}

	currentDir := dir
	for {
		envPath := filepath.Join(currentDir, "straw", ".env")
		if _, err := os.Stat(envPath); err == nil {
			return currentDir
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			log.Fatal(".env file not found")
		}
		currentDir = parentDir
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: create-builder, build-cloudflared, or build-certbot")
		return
	}

	loadEnv()

	switch os.Args[1] {
	case "create-builder":
		fmt.Println("Creating builder...")
		createCustomBuilder()
	case "build-nginx":
		fmt.Println("Building nginx...")
		buildNginxImage()
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
	}
}
