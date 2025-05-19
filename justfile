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
	go get -modfile=tools/go.mod -tool golang.org/x/tools/cmd/goimports
	go get -modfile=tools/go.mod -tool golang.org/x/tools/cmd/stringer
	go get -modfile=tools/go.mod -tool golang.org/x/pkgsite/cmd/pkgsite
	go get -modfile=tools/go.mod -tool golang.org/x/perf/cmd/benchstat
	go install -modfile=tools/go.mod tool

# QA all code
qa:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/environment
	log_info "Check go.mod and lint code"
	go mod tidy
	go mod verify
	(cd tools && go mod tidy && go mod verify)
	log_info "Format code"
	go tool -modfile=tools/go.mod goimports -w .
	gofmt -l -s -w $(find . -type f -name '*.go'| grep -v "/vendor/\|/.git/")
	log_info "Vetting"
	go vet ./...
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
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# benchmark all code
benchmark:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/environment
	log_info "Run benchmarks"
	go test -bench=. -benchmem ./... | tee benchmark-new.txt
	go tool -modfile=tools/go.mod benchstat benchmark-new.txt benchmark.txt

# generate documentation server
doc:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/environment
	log_info "Run documentation server at localhost:8080"
	go tool -modfile=tools/go.mod pkgsite

