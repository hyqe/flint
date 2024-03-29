#!/bin/bash
set -e

tag=$(cat VERSION)
image="flint:$tag"
container="flint"

echo "building $image"
docker build --no-cache -t $image .

echo "removing containter $container"
docker rm -f $container



mkdir -p ~/.flint

echo "running $image as $container"
docker run \
    -d \
    -p 2000:2000 \
    -v ~/.flint:/app/cache \
    --restart unless-stopped \
    --name $container \
    $image \
        --storage cache \
        --verbose