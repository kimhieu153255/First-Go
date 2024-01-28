postgres:
	docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root testGo

dropdb:
	docker exec -it postgres16 dropdb testGo

migrateup:
	migrate -path internal/config/db/migration -database "postgresql://root:secret@localhost:5432/testGo?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/config/db/migration -database "postgresql://root:secret@localhost:5432/testGo?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./tests/direct/... -coverpkg ./internal/config/db/sqlc/...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination internal/config/db/mock/store.go github.com/kimhieu153255/first-go/internal/config/db/sqlc Store

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server mock