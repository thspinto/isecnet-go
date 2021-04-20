test: lint unit-test
.PHONY: lint
lint:
	golangci-lint run

.PHONY: unit-test
unit-test:
	go test ./...

.PHONY: mock-alarm-central
mock-server:
	$(MAKE) -C internal/mockAlarmCentral run

.PHONY: gen
gen:
	go generate ./...
