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
build: lint
	$(DOCKER) build . -t $(DOCKER_REPOSITORY):$(shell $(GET_TAG))

.PHONY: build-bin
build-bin: lint
	$(GOCMD) build -o $(BINARY_NAME)

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run
