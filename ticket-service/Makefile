DB=postgres://postgres:12345678@localhost:5432/ticket?sslmode=disable
N=3
V=2
TITLE=rename_column_production_table_airplane

start:
	go run cmd/main.go

test:
	go test -v ./...
	
migrate-create:
	migrate create -ext sql -dir database/migrations -seq "$(TITLE)"

migrate-up:
	migrate -database "${DB}" -path database/migrations up $(N)


migrate-down:
	migrate -database "${DB}" -path database/migrations down $(N)


migrate-to-version: 
	migrate -database "${DB}" -path database/migrations goto $(V)


drop-db: 
	migrate -database "${DB}" -path database/migrations drop


force-version: 
	migrate -database "${DB}" -path database/migrations force $(V)


migration-version: 
	migrate -database "${DB}" -path database/migrations version