#!/bin/bash
GOOS=darwin GOARCH=amd64  go build -o bin/flint_darwin_amd64 -ldflags "-s -w" main.go
GOOS=darwin GOARCH=arm64  go build -o bin/flint_darwin_arm64 -ldflags "-s -w" main.go
GOOS=linux GOARCH=amd64  go build -o bin/flint_linux_amd64 -ldflags "-s -w" main.go
GOOS=linux GOARCH=arm64  go build -o bin/flint_linux_arm64 -ldflags "-s -w" main.go