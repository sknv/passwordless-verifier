//go:build tools

package tools

import (
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	_ "github.com/evilmartians/lefthook/cmd/lefthook"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/matryer/moq"
	_ "github.com/pressly/goose/v3/cmd/goose"
)
