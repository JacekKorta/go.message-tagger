# syntax=docker/dockerfile:1

FROM golang:1.19-alpine3.15 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /message-tagger-bin

FROM alpine:latest

COPY --from=build /message-tagger-bin /message-tagger-bin

ENTRYPOINT [ "/message-tagger-bin" ]