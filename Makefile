postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it postgres dropdb --username=postgres  simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate
test: 
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/hoangduc02011998/simple_bank/db/sqlc Store

.PHONY: postgres, createdb, dropdb, sqlc, test, server, mock
