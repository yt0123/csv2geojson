FROM golang:1.9-alpine

RUN apk add --no-cache make git
RUN go get -u github.com/golang/dep/cmd/dep