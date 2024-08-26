# Makefile

# Define your Go commands
build:
	@env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main cmd/web/main.go

run:
	@env APP_ENV=LOCAL go run cmd/web/main.go

run-worker:
	@go run cmd/worker/main.go

test:
	@go test ./... -v

migrate:
	@go run cmd/migrate/main.go up


migrate-down:
	@go run cmd/migrate/main.go down
