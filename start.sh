#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up | sed 's/[\]r//g'

echo "start the app"
exec "$@"
