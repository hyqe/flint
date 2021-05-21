#!/bin/bash

version=$(cat VERSION)

docker build . -t hyqe/flint:latest hyqe/flint:$version

docker login
docker push hyqe/flint:latest
docker push hyqe/flint:$version