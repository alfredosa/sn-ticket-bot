build:
	@go build -o bin/sn-bot

test:
	@go test ./...

run: build test 
	@echo "Binary built, running:"
	@./bin/sn-bot