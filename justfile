# !/usr/bin/env just --justfile
#
name := "ego-set"

default:
	@just --list --unsorted --justfile {{justfile()}} | grep -v default

# Install grpc plugins and other go tools
tools:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/environment
	log_info "Install go tools"
	go get -tool golang.org/x/tools/cmd/goimports
	go get -tool golang.org/x/tools/cmd/stringer
	go install tool

# QA all code
qa:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/environment
	log_info "Check go.mod and lint code"
	go mod tidy
	go mod verify
	log_info "Format code"
	go tool goimports -w .
	gofmt -l -s -w $(find . -type f -name '*.go'| grep -v "/vendor/\|/.git/")
	log_info "Linting"
	golangci-lint run -v
	log_info "Vulnerability checking"
	go run golang.org/x/vuln/cmd/govulncheck@latest --show verbose ./...

# unittest all code
unittest:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/environment
	log_info "Run unittests"
	go test -v ./...

