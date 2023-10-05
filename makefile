build:
	@go build -o bin/short

run: build
	@./bin/short
