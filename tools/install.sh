#!/usr/bin/env bash

echo "Make sure you have installed 'protoc' protobuf compiler: https://grpc.io/docs/protoc-installation"

go mod tidy && go mod vendor && go mod verify

go install \
  github.com/evilmartians/lefthook \
  github.com/golangci/golangci-lint/cmd/golangci-lint \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
  google.golang.org/protobuf/cmd/protoc-gen-go
