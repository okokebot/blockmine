FROM golang:latest

RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app

ENV GO111MODULE=on

EXPOSE 8080

CMD go run cmd/blockmine/main.go
