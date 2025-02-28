FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin ./cmd/server/*.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /build/bin ./bin
COPY ./database/migrations /migrations

COPY ./build/docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENV DBPATH=$GOOSE_DBSTRING
ENV PORT=3000

# Redundancy check because it's fun
RUN mkdir -p ./data
VOLUME ["./data"]

EXPOSE ${PORT}
ENTRYPOINT ["/entrypoint.sh"]
