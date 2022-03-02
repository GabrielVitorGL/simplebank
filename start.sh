#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
DB_SOURCE_CORRIGIDO=`echo "$DB_SOURCE" | sed 's/[\]r//g'`
/app/migrate -path /app/migration -database "$DB_SOURCE_CORRIGIDO" -verbose up

echo "start the app"
exec "$@"