run: build 
	@./bin/rest-book

build: 
	@go build -o bin/rest-book

test: 
	@go test -v ./...
