export GOOSE_DRIVER := "sqlite3"
export GOOSE_DBSTRING := "./data/db.sqlite3"
export GOOSE_MIGRATION_DIR := "./database/migrations"

export DBSTRING := GOOSE_DBSTRING

export PORT := "3000"

DOCKERFILE := "./build/docker/Dockerfile"

default: 
  just run server

dev APP="server":
  air --build.cmd "CGO_ENABLED=0 go build -o ./tmp/main ./cmd/{{APP}}/*"

run APP *ARGS:
  @mkdir -p tmp
  @CGO_ENABLED=0 go build -o ./tmp/main ./cmd/{{APP}}/* 
  ./tmp/main {{ARGS}}

docker-build:
  docker build --load -t realgofile:latest -f {{DOCKERFILE}} . 

sqlc-gen:
  go tool sqlc generate

sqlc-test:
  go tool sqlc compile
  
migrate-up:
  go tool goose up

migrate-down:
  go tool goose down  

sqlite:
  sqlite3 $GOOSE_DBSTRING

new-migraition MIGRATION_NAME:
  go tool goose create {{MIGRATION_NAME}} sql