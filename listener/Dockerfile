# Build Stage
FROM golang:1.21.4-alpine3.17 AS build

WORKDIR /app

RUN apk add --no-cache libc-dev gcc

COPY go.mod go.sum ./
RUN go mod download

COPY listener/ ./listener/
COPY contract/ ./contract/

RUN go build \
    -ldflags="-s -w -linkmode external -extldflags -static" \
    -tags netgo \
    -o contract-listener ./listener/cmd

# Final Stage
FROM scratch

COPY --from=build /app/contract-listener /contract-listener
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

CMD ["/contract-listener"]
