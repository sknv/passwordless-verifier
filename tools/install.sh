#!/usr/bin/env bash

go mod tidy && go mod vendor && go mod verify

go install \
  github.com/evilmartians/lefthook \
  github.com/golangci/golangci-lint/cmd/golangci-lint \
  github.com/deepmap/oapi-codegen/cmd/oapi-codegen \
  github.com/pressly/goose/v3/cmd/goose \
  github.com/matryer/moq
