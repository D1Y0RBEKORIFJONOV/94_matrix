SWAGGER_CMD = swag
GO_RUN_CMD = go run

swagger-init:
	$(SWAGGER_CMD) init -g internal/http/handler/book.go -o internal/app/docs

run:
	$(GO_RUN_CMD) cmd/app/main.go

all: swagger-init run

migrate-file:
	migrate create -ext sql -dir migrations/ -seq book

DB_URL := "postgres://postgres:+_+diyor2005+_+@postgres_container:5432/book?sslmode=disable"


migrate-up:
	migrate -path migrations -database $(DB_URL) -verbose up

migrate-down:
	migrate -path migrations -database $(DB_URL) -verbose down

migrate-force:
	migrate -path migrations -database $(DB_URL) -verbose force 1
