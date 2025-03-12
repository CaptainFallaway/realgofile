#!/bin/sh
set -e

mkdir -p /data
goose sqlite3 $DBSTRING -dir=migrations up

exec ./bin