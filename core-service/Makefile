build:
	go build -o core-service .

docs:
	swag init -g internal/adapter/http.go

docs-docker:
	docker run --rm -v $(pwd):/code ghcr.io/swaggo/swag:latest -- init -g internal/adapter/http.go

migrate: 
	go run internal/postgres/migrate.go