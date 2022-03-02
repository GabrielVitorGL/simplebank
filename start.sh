#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
echo "source"
DB_SOURCE_CORRIGIDO=`echo "$DB_SOURCE" | sed 's/[\]r//g'`
echo "corrigiu"
/app/migrate -path /app/migration -database "$DB_SOURCE_CORRIGIDO" -verbose up
echo "rodou"

echo "start the app"
exec "$@"