#!/bin/bash

: '
Releases can be found here: https://github.com/cloudflare/cloudflared/releases

URLs for cloudflared binaries:

x86_64:
https://github.com/cloudflare/cloudflared/releases/download/${CLOUDFLARED_VERSION}/cloudflared-linux-amd64

aarch64:
https://github.com/cloudflare/cloudflared/releases/download/${CLOUDFLARED_VERSION}/cloudflared-linux-arm64

armv7l:
https://github.com/cloudflare/cloudflared/releases/download/${CLOUDFLARED_VERSION}/cloudflared-linux-armhf

i386:
https://github.com/cloudflare/cloudflared/releases/download/${CLOUDFLARED_VERSION}/cloudflared-linux-386
'

CLOUDFLARED_VERSION="2024.8.2"
IMAGE_NAME="overtime0022/cloudflared"

# Build the Docker image
docker build --build-arg CLOUDFLARED_VERSION=${CLOUDFLARED_VERSION} -t ${IMAGE_NAME}:${CLOUDFLARED_VERSION} .

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Docker image built successfully."
else
  echo "Docker image build failed."
  exit 1
fi

# Push the Docker image to the registry
docker push ${IMAGE_NAME}:${CLOUDFLARED_VERSION}

# Check if the push was successful
if [ $? -eq 0 ]; then
  echo "Docker image pushed successfully."
else
  echo "Docker image push failed."
  exit 1
fi
