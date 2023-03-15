DB=mysql://root@tcp(localhost:3306)/airplane?multiStatements=true&tls=false
N=1
V=1
TITLE=

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