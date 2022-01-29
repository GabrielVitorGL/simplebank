postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pass412 -d postgres:14-alpine

stopps:
	docker stop postgres14

removeps:
	docker rm postgres14

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:pass412@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:pass412@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	powershell -command "docker run --rm -v $${pwd}:/src -w /src kjconroy/sqlc generate"

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres stopps removeps createdb dropdb migrateup migratedown sqlc test server
