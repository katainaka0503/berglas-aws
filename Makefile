GOCMD=go
GOLANGCI_LINT=golangci-lint
DOCKER=docker
DOCKER_REPOSITORY=katainaka0503/berglas-aws
GET_TAG=git rev-parse HEAD
BINARY_NAME=berglas-aws
BINARY_NAME_LINUX=$(BINARY_NAME)_linux

.PHONNY: clean
clean:
	$(GOCMD) clean

.PHONY: build
build: build-bin-linux
	$(DOCKER) build . -t $(DOCKER_REPOSITORY):$(shell $(GET_TAG))

.PHONY: build-bin
build-bin: lint
	$(GOCMD) build -o $(BINARY_NAME)

.PHONY: build-bin-linux
build-bin-linux: lint
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -o $(BINARY_NAME_LINUX)

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run
