DB_URL=postgresql://root:mokopass@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgres14 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mokopass -d postgres:14-alpine

postgresrun:
	docker start postgres14

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

db_docs:
	dbdocs build docs/db.dbml

db_schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go simple-bank/db/sqlc Store

testcoverhtml:
	go test ./... -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

dockerbuildimage:
	docker build -t simplebank:latest .

dockerrunbuildcontainer:
	docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE=postgresql://root:mokopass@postgres14:5432/simple_bank?sslmode=disable simplebank:latest

dockerstart:
	docker start simplebank

.PHONY: postgres postgresrun createdb dropdb migrateup migrateup1 migratedown migratedown1 new_migration db_docs db_schema sqlc server mock testcoverhtml dockerbuildimage dockerrunbuildcontainer dockerstart