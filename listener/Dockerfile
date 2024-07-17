# Build Stage
FROM golang:1.21.4-alpine3.17 AS build

WORKDIR /app

RUN apk add --no-cache libc-dev gcc

COPY go.mod go.sum ./

RUN go mod download

COPY listener/ ./listener/
COPY .env .
COPY contract/ ./contract/

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o contract-listener ./listener/cmd

CMD ["./contract-listener"]
