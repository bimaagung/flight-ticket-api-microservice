DB=mysql://root@tcp(localhost:3306)/airplane?multiStatements=true

start:
	go run cmd/main.go

migrate-create:
	migrate create -ext sql -dir database/migrations -seq

migrate-up:
	migrate -database "${DB}" -path database/migrations up

migrate-force:
	migrate -database "${DB}" -path database/migrations force

migrate-down:
	migrate -database "${DB}" -path database/migrations down

test:
	go test -v ./...