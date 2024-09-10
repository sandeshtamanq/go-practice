build:
	@go build -o bin/jwt cmd/main.go

run: build
	@./bin/jwt