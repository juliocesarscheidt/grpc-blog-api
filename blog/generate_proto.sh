#!/bin/bash

protoc \
  --go_out=blogpb \
  --go_opt=paths=source_relative \
  --go-grpc_out=blogpb \
  --go-grpc_opt=paths=source_relative \
  --proto_path=./protofiles \
  ./protofiles/*.proto
