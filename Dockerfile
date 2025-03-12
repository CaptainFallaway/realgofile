FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
# RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin ./cmd/api/*.go

FROM alpine:latest AS prod

WORKDIR /app

# COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /build/bin ./bin
COPY ./database/migrations /app/migrations

COPY ./build/docker/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

ENV DBSTRING=/app/data/db.sqlite3
ENV PORT=3000

# RUN mkdir -p /app/data
VOLUME ["/app/data"]

EXPOSE ${PORT}
ENTRYPOINT ["/app/entrypoint.sh"]