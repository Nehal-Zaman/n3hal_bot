build:
	@go build -o bin/n3hal_bot -ldflags="-s -w"

run:build
	@./bin/n3hal_bot