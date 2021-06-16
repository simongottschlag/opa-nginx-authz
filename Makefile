SHELL := /bin/bash

.PHONY: lint
.SILENT: lint
lint:
	golangci-lint run

.PHONY: cover
.SILENT: cover
cover:
	mkdir -p tmp/
	go test -timeout 1m -coverpkg=./... -coverprofile=tmp/coverage.out ./...
	go tool cover -html=tmp/coverage.out	
