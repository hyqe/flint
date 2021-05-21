#!/bin/bash

image="flint"
container="flint"

echo "building image..."
docker build . -t $image

echo "removing old container..."
docker rm -f $container


echo "running image..."
docker run \
    -p 1389:1389 \
    --name $container \
    $image