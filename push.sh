#!/bin/bash

version=$(cat VERSION)

docker buildx create --name builder --use

docker buildx inspect --bootstrap

docker buildx build \
    --platform linux/amd64,linux/arm64,linux/arm/v7 \
    -t hyqe/flint:$version username/demo:latest \
    .

docker login

docker push hyqe/flint:$version
docker push hyqe/flint:latest