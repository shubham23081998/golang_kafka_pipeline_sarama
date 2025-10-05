#!/bin/bash
set -euo pipefail

# Always run from project root
cd "$(dirname "$0")/.." || exit 1

# Define image name and version
IMAGE_NAME="golang-kafka-pipeline"
IMAGE_TAG="0.0.1"

# Build Docker image
echo "Building Docker image: ${IMAGE_NAME}:${IMAGE_TAG} ..."
docker build -t "${IMAGE_NAME}:${IMAGE_TAG}" -f Dockerfile .

echo "✅ Successfully built ${IMAGE_NAME}:${IMAGE_TAG}"
