.ONESHELL:

SHELL := /bin/bash
.SHELLFLAGS := -ec

BIN=$(CURDIR)/bin

install-vendor:
	go get -u github.com/goware/modvendor
	go mod vendor
	$(GOPATH)/bin/modvendor -copy="**/*.c **/*.h **/*.proto **/*.m"

build: install-vendor
	go build \
		-mod=vendor \
		-o $(BIN)/asteroids \
		./cmd

run: build
	$(BIN)/asteroids