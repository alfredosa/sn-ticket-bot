build:
	@go build -o bin/sn-bot

test:
	@cd $(CURDIR) && go test -v ./...

run: build test 
	@echo "Binary built, running:"
	@./bin/sn-bot
