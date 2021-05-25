#!/bin/bash
image_latest="hyqe/flint:latest"
image_version="hyqe/flint:$(cat VERSION)"

docker login
docker build -t $image_latest -t $image_version .
docker push $image_latest