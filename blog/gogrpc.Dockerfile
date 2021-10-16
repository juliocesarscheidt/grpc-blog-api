FROM golang:1.14
LABEL maintainer="Julio Cesar <julio@blackdevs.com.br>"

ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /go/

RUN apt-get update -yqq && \
    apt-get install -yqq \
    build-essential \
    apt-transport-https \
    curl netcat \
    protobuf-compiler \
    golang-goprotobuf-dev

RUN curl --silent \
    -L https://github.com/ktr0731/evans/releases/download/0.9.3/evans_linux_amd64.tar.gz \
    -o evans_linux_amd64.tar.gz && \
    tar -xzvf evans_linux_amd64.tar.gz && \
    rm -rf evans_linux_amd64.tar.gz && \
    chmod +x evans && \
    mv evans /usr/local/bin/

RUN go get -u \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/*
