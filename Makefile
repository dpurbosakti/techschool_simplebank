DB_URL=postgresql://root:mokopass@localhost:5432/simple_bank?sslmode=disable

## postgres: run a PostgreSQL container with specific configurations
postgres:
	docker run --name postgres14 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mokopass -d postgres:14-alpine

## postgresrun: start docker 
postgresrun:
	docker start postgres14

## createdb: create database
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

## dropdb: drop database
dropdb:
	docker exec -it postgres14 dropdb simple_bank

## migrateup: migrate all up schema sql
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

## migrateup1: migrate up schema sql by 1
migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

## migratedown: migrate all down schema sql
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

## migratedown1: migrate down schema sql by 1
migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

## new_migration: init sql migration file with name as parameter
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

## db_docs: create documentation for database dbml
db_docs:
	dbdocs build docs/db.dbml

## db_schema: create schema sql using db.dbml
db_schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml

## sqlc: generate repository code from query
sqlc:
	sqlc generate

## test: run test
test:
	go test -v -cover ./...

## server: run server
server:
	go run main.go

## mock: generate mock data
mock:
	mockgen -package mockdb -destination db/mock/store.go simple-bank/db/sqlc Store

## testcoverhtml: run test and create coverprofile
testcoverhtml:
	go test ./... -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

## dockerbuildimage: build image
dockerbuildimage:
	docker build -t simplebank:latest .

## dockerrunbuildcontainer: run an image container with specific configurations
dockerrunbuildcontainer:
	docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE=postgresql://root:mokopass@postgres14:5432/simple_bank?sslmode=disable simplebank:latest

## dockerstart: start docker container
dockerstart:
	docker start simplebank

## proto: generate go code from proto
proto:
	rm -f pb/*.go
	rm -f docs/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
    proto/*.proto

# after running this command, before u call the RPC, try tp check the package and service, then use package <package_name> and service <service_name> 
## evans: connect to evans
evans:
	evans --host localhost --port 9090 -r repl

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: postgres postgresrun createdb dropdb migrateup migrateup1 migratedown migratedown1 new_migration db_docs db_schema sqlc server mock testcoverhtml dockerbuildimage dockerrunbuildcontainer dockerstart proto evans help