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


.PHONY: hassio-dev
hassoi-dev:
	docker run --rm -ti --name hassio-builder --privileged \
		-v $(shell pwd)/hassio-addon:/data \
		-v /var/run/docker.sock:/var/run/docker.sock:ro \
		homeassistant/amd64-builder -t /data --all --test \
		-i hassio -d thspinto/hassio-isecnet
