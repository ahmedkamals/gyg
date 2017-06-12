NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)

# Binary output name.
BINARY=./bin/$(shell basename `pwd`)

# Version and build values.
VERSION=1.0.0
BUILD=`git rev-parse HEAD`

# Setup the -ldflags option for go build here.
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Default target.
.DEFAULT_GOAL: ${BINARY}

.PHONY: clean install

${BINARY}:
	@echo "$(OK_COLOR)==> Building$(NO_COLOR)"
	go build -v ${LDFLAGS} -o ${BINARY} .

all: format vet test bench clean deps build
default: deps install build

deps:
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
	@go get -t -d -v ./...
	@echo $(DEPS) | xargs -n1 go get -d

updatedeps:
	@echo "$(OK_COLOR)==> Updating all dependencies$(NO_COLOR)"
	@go get -t -d -v -u ./...
	@echo $(DEPS) | xargs -n1 go get -d -u

build: clean \
    vet \
	test \
	${BINARY}

# Installs the project dependencies
install:
	@echo "$(OK_COLOR)==> Installing$(NO_COLOR)"
	go install ${LDFLAGS} .

# Cleaning the project, by deleting binaries.
clean:
	@echo "$(OK_COLOR)==> Cleaning$(NO_COLOR)"
	if [ -f ${BINARY} ] ; then rm -rf bin ; fi
	go clean

format:
	@echo "$(OK_COLOR)==> Formatting$(NO_COLOR)"
	go fmt ./...

test:
	@echo "$(OK_COLOR)==> Testing$(NO_COLOR)"
	go test -v ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''

# Runs benchmarks
bench:
	@echo "$(OK_COLOR)==> Benchmarking$(NO_COLOR)"
	go test ./... -bench=.

# https://golang.org/cmd/vet/
vet:
	go vet ./...

# Documentation.
doc:
	godoc -http=:6060

# External dependencies
# =====================

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	@echo "$(OK_COLOR)==> Linting$(NO_COLOR)"
	golint ./...

dev:
	DEBUG=* go get && go install && gin -p 8911 -i

godep:
	godep save ./...