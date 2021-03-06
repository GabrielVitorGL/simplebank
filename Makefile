postgres:
	docker run --name postgres14 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pass412 -d postgres:14-alpine

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

migrateup1:
	migrate -path db/migration -database "postgresql://root:pass412@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:pass412@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:pass412@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	powershell -command "docker run --rm -v $${pwd}:/src -w /src kjconroy/sqlc generate"

test:
	go clean -testcache && go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

simplebank:
	docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:pass412@postgres14:5432/simple_bank?sslmode=disable" simplebank:latest
	
.PHONY: postgres stopps removeps createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock simplebank
