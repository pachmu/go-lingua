SHELL := /bin/bash
PROJECT_NAME := $(shell basename $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST)))))
CURRENT_DIR := $(shell pwd)
DOCKER_IMAGE_NAME := $(PROJECT_NAME):latest

prepare_image:
	docker build -t $(DOCKER_IMAGE_NAME) -f Dockerfile .

clean:
	rm -rf ./build

build: prepare_image
	docker run --privileged --rm -t \
		--name $(PROJECT_NAME)-build-$(TIMESTAMP) \
		-v $(CURRENT_DIR):/$(PROJECT_NAME) \
		$(DOCKER_IMAGE_NAME) \
		bash -c "GOPATH=/ GOOS=linux GOARCH=amd64 go build -o /$(PROJECT_NAME)/build/$(PROJECT_NAME) /src/go-lingua"