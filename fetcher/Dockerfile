# Build Stage
FROM golang:1.21.4-alpine3.17 AS build

WORKDIR /app

RUN apk add --no-cache libc-dev gcc

COPY go.mod go.sum ./

RUN go mod download

COPY fetcher/ ./fetcher/
COPY .env .

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o ethereum-fetcher ./fetcher/cmd

EXPOSE 8080

CMD ["./ethereum-fetcher"]