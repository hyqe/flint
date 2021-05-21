#!/bin/bash

tag=$(cat VERSION)
image="flint:$tag"
container="flint"

echo "building $image"
docker build . -t $image

echo "removing $container"
docker rm -f $container

echo "running $image as $container"
docker run \
    -p 2000:2000 \
    --name $container \
    $image \
        --verbose