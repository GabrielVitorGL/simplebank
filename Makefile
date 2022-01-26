postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=jogos3dg -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:jogos3dg@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:jogos3dg@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	powershell -command "docker run --rm -v $${pwd}:/src -w /src kjconroy/sqlc generate"

.PHONY: postgres createdb dropdb migrateup migratedown sqlc