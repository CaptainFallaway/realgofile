FROM golang:1.24-alpine AS builder

WORKDIR /build

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin ./cmd/server/*.go

FROM alpine:latest

# Setup base enviroment

ENV PORT=3000
ENV DBPATH=/data/db.sqlite3
RUN mkdir -p /data
VOLUME ["/data"]

WORKDIR /app

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /build/bin ./bin
COPY ./database/migrations ./migrations

COPY ./build/docker/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

EXPOSE ${PORT}
ENTRYPOINT ["./entrypoint.sh"]
