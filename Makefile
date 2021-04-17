test: lint
.PHONY: lint
lint:
	golangci-lint run

.PHONY: unit-test
unit-test:
	go test ./...

.PHONY: mock-server
mock-server:
	$(MAKE) -C mockServer run

.PHONY: proto
proto:
	protoc --go-grpc_out=internal/genproto/zones -I api/protobuf api/protobuf/zones.proto
	protoc --go_out=internal/genproto/zones -I api/protobuf api/protobuf/zones.proto
