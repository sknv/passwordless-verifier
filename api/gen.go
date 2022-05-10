package api

//nolint:lll // allowed to be long
//go:generate protoc -I=./openapi -I=../third_party/include --go_out=./openapi --go_opt=paths=source_relative --openapiv2_out=./openapi --openapiv2_opt=logtostderr=true openapi.proto
