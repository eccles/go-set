#!/usr/bin/env just --justfile
#
name := "go-set"

default:
	@just --list --unsorted --justfile {{justfile()}} | grep -v default

# QA all code
qa:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/log
	log_info "Check go.mod and lint code"
	which go
	go mod tidy
	go mod verify
	log_info "Vetting"
	go vet ./...
	log_info "Formatting"
	golangci-lint fmt ./...
	log_info "Linting"
	golangci-lint run ./...
	log_info "Vulnerability checking"
	go run golang.org/x/vuln/cmd/govulncheck@latest --show verbose ./...

# unittest all code
unittest:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/log
	log_info "Run unittests"
	which go
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# benchmark all code
benchmark:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/log
	log_info "Run benchmarks"
	which go
	rm -f benchmark-new.txt
	go test -run='^$' -count=10 -bench=. -benchmem ./... | tee benchmark-new.txt
	go run golang.org/x/perf/cmd/benchstat@latest benchmark.txt benchmark-new.txt

# generate documentation server
doc:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/log
	log_info "Run documentation server at localhost:8080"
	which go
	go run golang.org/x/pkgsite/cmd/pkgsite@latest

# publish package to proxy
publish:
        #!/usr/bin/env bash
        set -euo pipefail
        source ./scripts/source/log
        log_info "Publish"
        which go
        VERSION=$(git tag -l | sort -r -V | head -1)
        GOPROXY=proxy.golang.org go list -m github.com/eccles/go-set@${VERSION}
