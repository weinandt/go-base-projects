# GRPC Base

## Pre-reqs
- Protoc installation
- Protoc plugins (make sure you path includes the go/bin):
    - `go install google.golang.org/protobuf/cmd/protoc-gen-go`
    - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`

## Generating Go Code from proto
`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user/user.proto`

## Running
1. `go run server.go`
2. `go run client.go`