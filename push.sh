#!/bin/bash

version=$(cat VERSION)

docker build . -t hyqe/flint:$version hyqe/flint:latest

docker login
docker push hyqe/flint:$version
docker push hyqe/flint:latest