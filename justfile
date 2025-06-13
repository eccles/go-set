#!/usr/bin/env just --justfile
#
name := "go-set"

default:
	@just --list --unsorted --justfile {{justfile()}} | grep -v default

# Install grpc plugins and other go tools
tools:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/log
	log_info "Install go tools"
	which go
	go get -modfile=tools/go.mod -tool golang.org/x/tools/cmd/goimports
	go get -modfile=tools/go.mod -tool golang.org/x/pkgsite/cmd/pkgsite
	go get -modfile=tools/go.mod -tool golang.org/x/perf/cmd/benchstat
	go get -modfile=tools/go.mod -tool honnef.co/go/tools/cmd/staticcheck
	go install -modfile=tools/go.mod tool

# QA all code
qa:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/log
	log_info "Check go.mod and lint code"
	which go
	go mod tidy
	go mod verify
	(cd tools && go mod tidy && go mod verify)
	log_info "Format code"
	go tool -modfile=tools/go.mod goimports -w .
	gofmt -l -s -w $(find . -type f -name '*.go'| grep -v "/vendor/\|/.git/")
	log_info "Vetting"
	go vet ./...
	log_info "Linting"
	go tool -modfile=tools/go.mod staticcheck ./...
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
	go test -bench=. -benchmem ./... | tee benchmark-new.txt
	go tool -modfile=tools/go.mod benchstat benchmark.txt benchmark-new.txt

# generate documentation server
doc:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/log
	log_info "Run documentation server at localhost:8080"
	which go
	go tool -modfile=tools/go.mod pkgsite

# publish package to proxy
publish:
        #!/usr/bin/env bash
        set -euo pipefail
        source ./scripts/source/log
        log_info "Publish"
        which go
        VERSION=$(git tag -l | sort -r -V | head -1)
        GOPROXY=proxy.golang.org go list -m github.com/eccles/go-set@${VERSION}
