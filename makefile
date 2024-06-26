sqlc:
	sqlc generate

postgresql:
	docker run --name postgres15 --network bank-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:15

migratedb:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/golomt?sslmode=disable" -verbose up

downdb:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/golomt?sslmode=disable" -verbose down

test:
		go test -v -cover ./test

run:
	go run main.go

mock:
	mockgen -destination db/mock/store.go simplebank/db/sqlc Store	

.PHONY: test mock run sqlc downdb server migratedb