package api

//nolint:lll // allowed to be long
//go:generate protoc -I=./openapi/v1 -I=../third_party/include --go_out=./openapi/v1 --go_opt=paths=source_relative --openapiv2_out=./openapi/v1 --openapiv2_opt=logtostderr=true openapi_v1.proto
