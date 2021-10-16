# Go GRPC API

## Commands

```bash

################ blog GRPC ################
cd blog

export GO111MODULE=on
go mod init github.com/juliocesarscheidt/blog
go mod tidy

go mod download

# generate proto GRPC
protoc \
  --go_out=blogpb \
  --go_opt=paths=source_relative \
  --go-grpc_out=blogpb \
  --go-grpc_opt=paths=source_relative \
  --proto_path=./protofiles \
  ./protofiles/*.proto

# server
go run main.go

```
