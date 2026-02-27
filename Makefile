# Makefile

.PHONY: build run test migrate seed clean

build:
	go build -o bin/go-backend-project ./cmd/main.go

run: build
	./bin/go-backend-project

test:
	go test ./... -v

migrate:
	./scripts/migrate.sh

seed:
	./scripts/seed.sh

clean:
	go clean
	rm -rf bin/