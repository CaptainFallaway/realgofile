#!/bin/sh
set -e

mkdir -p /app/data
goose sqlite3 ${DBPATH} -dir=migrations up

exec /app/bin