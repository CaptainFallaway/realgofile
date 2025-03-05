export GOOSE_DRIVER := "sqlite3"
export GOOSE_DBSTRING := "./data/db.sqlite3"
export GOOSE_MIGRATION_DIR := "./database/migrations"

export PORT := "3000"

DOCKERFILE := "./build/docker/Dockerfile"

dev APP="server":
  air --build.cmd "go build -o ./tmp/main ./cmd/{{APP}}/*"

docker-build:
  docker build --load -t solaris-server:latest -f {{DOCKERFILE}} . 

sqlc-gen:
  go tool sqlc generate

sqlc-test:
  go tool sqlc compile
  
migrate-up:
  go tool goose up

migrate-down:
  go tool goose down  

sqlite:
  sqlite3 ./data/db.sqlite3

new-migraition MIGRATION_NAME:
  go tool goose create {{MIGRATION_NAME}} sql