.PHONY: build-bin
build-bin: lint
	go build -o berglas-aws

.PHONY: lint
lint:
	golangci-lint run
