FROM juliocesarmidia/gogrpc:1.14
LABEL maintainer="Julio Cesar <julio@blackdevs.com.br>"

WORKDIR /go/src/app

COPY go.mod go.sum /go/src/app/
RUN go mod download

COPY . /go/src/app/

CMD [ "go", "run", "main.go" ]
