#!/bin/sh

set -e

golangci-lint run ./...
echo 'lint OK'
go test -race -short ./...
echo 'go test OK'