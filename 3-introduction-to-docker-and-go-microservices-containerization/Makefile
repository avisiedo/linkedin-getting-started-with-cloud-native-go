BIN := microservice
PORT ?= 8080

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
