#!/bin/bash

image="flint"
container="flint"

echo "building image..."
docker build . -t $image

echo "removing old container..."
docker rm -f $container

echo "running image..."
docker run \
    -p 2000:2000 \
    --name $container \
    $image \
        -v