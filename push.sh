#!/bin/bash
image_latest="hyqe/flint:latest"
image_version="hyqe/flint:$(cat VERSION)"
platforms="linux/amd64,linux/arm64,linux/arm/v7"

docker login
docker buildx build --platform $platforms -t $image_latest -t $image_version --push .