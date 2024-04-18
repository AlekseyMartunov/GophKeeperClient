FROM golang:1.21-alpine AS builder

WORKDIR /usr/local/src

COPY ["app/go.mod", "app/go.sum", "./"]
RUN go mod download

COPY ["app", "./"]
RUN go build -o ./bin/app cmd/client/main.go

FROM alpine:latest AS runner
COPY  --from=builder usr/local/src/bin/app /
CMD ["/app"]
