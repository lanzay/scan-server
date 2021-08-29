go-run:
	go run -v ./cmd/scan-server
go-build:
	go get -d -v ./...
	go build -v -o ./bin/ ./cmd/scan-server

up:
	docker-compose -p scan-server-DEV -f docker-compose.yml -f docker-compose.dev.yml up --build -d

up-prod:
	docker-compose -p scan-server-PROD -f docker-compose.yml -f docker-compose.prod.yml up --build -d