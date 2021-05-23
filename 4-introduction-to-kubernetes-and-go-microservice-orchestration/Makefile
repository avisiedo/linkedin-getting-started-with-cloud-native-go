ifneq (,$(shell command -v podman 2>/dev/null))
DOCKER := podman
else
	ifneq (,$(shell command -v docker 2>/dev/null))
DOCKER := docker
	else
DOCKER :=
	endif
endif

BIN := microservice
PORT ?= 8080
IMG_BASE ?= quay.io/avisied0
IMG ?= $(IMG_BASE)/go-cloud-native:latest

.PHONY: default
default: help

.PHONY: help
help:
	@cat HELP.txt

.PHONY: all
all: build test

.PHONY: build
build:
	go build -mod vendor -o $(BIN)

.PHONY: test
test:
	go test -mod vendor

.PHONY: clean
clean:
	-rm $(BIN)

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: run
run: build test
	PORT=$(PORT) ./microservice

.PHONY: container-build
container-build:
	$(DOCKER) build -t $(IMG) -f Dockerfile .

.PHONY: container-push
container-push:
	$(DOCKER) push $(IMG)

.PHONY: container-run
container-run:
	$(DOCKER) run -it -e PORT=$(PORT) -p $(PORT):$(PORT) $(IMG)

.PHONY: k8s-create
k8s-create:
	kustomize build config/ | kubectl create -f -

.PHONY: k8s-delete
k8s-delete:
	kustomize build config/ | kubectl delete -f -
