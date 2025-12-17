FROM golang:1.24.10-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
    GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o orga ./cmd/api

FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from-builder /app/orga .

EXPOSE 8080

CMD ["./orga"]
