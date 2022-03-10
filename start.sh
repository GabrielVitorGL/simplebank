#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
/app/migrate -path /app/migration -database "postgresql://root:pass412@localhost:5432/simple_bank?sslmode=disable" -verbose up

echo "start the app"
exec "$@"
