postgres:
	docker run --name postgres14 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mokopass -d postgres:14-alpine

postgresrun:
	docker start postgres14

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:mokopass@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:mokopass@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:mokopass@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:mokopass@localhost:5432/simple_bank?sslmode=disable" -verbose down 1


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

.PHONY: postgres postgresrun createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc server mock testcoverhtml dockerbuildimage dockerrunbuildcontainer dockerstart