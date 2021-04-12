test: lint unit-test

lint:
	golangci-lint run

unit-test:
	go test ./...

mock-server:
	$(MAKE) -C mockServer run
