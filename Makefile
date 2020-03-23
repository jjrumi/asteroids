.ONESHELL:

SHELL := /bin/bash
.SHELLFLAGS := -ec

install-vendor:
	go mod vendor
	$(GOPATH)/bin/modvendor -copy="**/*.c **/*.h **/*.proto **/*.m"