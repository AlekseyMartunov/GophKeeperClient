FROM golang:latest AS builder

WORKDIR /usr/local/src

COPY ["app/go.mod", "app/go.sum", "./"]
RUN go mod download

COPY ["app", "./"]

RUN go build -o app cmd/client/main.go
ENV MIGRATION_PATH=/migrations

