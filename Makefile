build:
	go build -v -o srv-trust
lint:
	golangci-lint run