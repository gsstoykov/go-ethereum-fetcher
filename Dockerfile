# Build Stage
FROM golang:1.21.4-alpine3.17 AS build

WORKDIR /app

RUN apk add --no-cache libc-dev gcc

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o api-go ./cmd

EXPOSE 8080

CMD ["./api-go"]