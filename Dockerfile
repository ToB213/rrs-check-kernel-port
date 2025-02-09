FROM golang:latest

ADD ./main.go /go/src

WORKDIR /go/src