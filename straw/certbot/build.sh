#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Create Buildx builder if it doesn't exist
if ! docker buildx inspect mybuilder > /dev/null 2>&1; then
  docker buildx create --name multiplatform --driver docker-container --use
fi

# Build and push the image
docker buildx build \
  --platform linux/amd64,linux/arm/v6,linux/arm64 \
  --build-arg GOLANG_IMAGE_VERSION=$GOLANG_IMAGE_VERSION \
  --build-arg CERTBOT_IMAGE_VERSION=$CERTBOT_IMAGE_VERSION \
  --tag williamveith/certbot:$CERTBOT_IMAGE_VERSION \
  --file certbot/Dockerfile \
  --push \
  .