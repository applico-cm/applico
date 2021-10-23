DOCKER := docker
GO := go

PROJECT ?= applico
# TODO use the correct ECR repository 
IMAGE_REPO = docker.io
IMAGE_TAG ?= dev
GO_VERSION ?= 1.17
OUTPUT_DIR ?= ./bin
IMAGE_NAME_API := $(IMAGE_REPO)/applico-api:$(IMAGE_TAG)


default: test

images:
	@echo $(IMAGE_NAME_API)

test:
	$(GO) test -v -tags=unit ./...

build-binary-local:
	GOSUMDB=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -ldflags="-w -s" -o $(OUTPUT_DIR)/applico-api ./cmd/applico-api

build-images-local:
	$(DOCKER) build -f ./cmd/applico-api/Dockerfile -t $(IMAGE_NAME_API) .

build: build-images-local

push-images:
	$(DOCKER) push $(IMAGE_NAME_API)

delete-images:
	$(DOCKER) rmi -f $(IMAGE_NAME_API)
	$(DOCKER) image prune --filter label=build=$(IMAGE_NAME_API) --force

test-ci:
	$(DOCKER) run --rm -v `pwd`:/applico -w /applico golang:$(GO_VERSION) make test

.PHONY: default
