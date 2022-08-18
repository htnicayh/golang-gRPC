postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=hyacinth -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root banking

dropdb: 
	docker exec -it postgres dropdb banking

migrateup:
	migrate -path database/migration -database "postgresql://root:hyacinth@localhost:5432/banking?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://root:hyacinth@localhost:5432/banking?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown