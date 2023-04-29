DB=postgres://postgres:12345678@localhost:5432/ticket?sslmode=disable
N=2
V=2
TITLE=relation_table_ticket_and_track

start:
	go run cmd/main.go

start-consumer:
	go run cmd/event/consumer.go

test:
	go test -v ./...
	
migrate-create:
	migrate create -ext sql -dir database/migrations -seq "$(TITLE)"

migrate-up:
	migrate -database "${DB}" -path database/migrations up


migrate-down:
	migrate -database "${DB}" -path database/migrations down


migrate-to-version: 
	migrate -database "${DB}" -path database/migrations goto $(V)


drop-db: 
	migrate -database "${DB}" -path database/migrations drop


force-version: 
	migrate -database "${DB}" -path database/migrations force $(V)


migration-version: 
	migrate -database "${DB}" -path database/migrations version