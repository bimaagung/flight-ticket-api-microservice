TRACK_BINARY=trackApp
AIRPLANE_BINARY=airplaneApp
TICKET_BINARY=ticketApp

up:
	@echo Starting Docker Images....
	docker-compose up -d
	@echo Docker Images Started!

up-build: build-track build-airplane build-ticket
	@echo Stopping docker images (if running...)
	docker-compose down
	@echo Building (when required) and starting docker images...
	docker-compose up --build -d
	@echo Docker images build and started!

down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!

build-track:
	@echo Building track binary
	chdir ..\track-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${TRACK_BINARY} ./cmd
	@echo Done

build-airplane:
	@echo Building airplane binary
	chdir ..\airplane-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${AIRPLANE_BINARY} ./cmd
	@echo Done

build-ticket:
	@echo Building ticket binary
	chdir ..\ticket-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${TICKET_BINARY} ./cmd
	@echo Done