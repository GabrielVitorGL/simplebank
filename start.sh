#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
echo "ORIGINAL>>>"
echo $DB_SOURCE
echo "MODIFICADO>>>"
echo $DB_SOURCE | sed 's/[\]r//g'
/app/migrate -path /app/migration -database "postgresql://root:pass412@postgres:5432/simple_bank?sslmode=disable" -verbose up

echo "start the app"
exec "$@"
