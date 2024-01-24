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

.PHONY: createdb dropdb postgres migrateup migratedown