PACKAGE  = influx
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)

GO        = go
GOBIN   = $(GOPATH)/bin
GODOC   = godoc
GOFMT   = gofmt
GOLINT  = gometalinter

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[0;35m▶\033[0m")

.PHONY: all
all: dep check

# Tests
.PHONY: dep
dep:
	$Q go get -u github.com/golang/dep/cmd/dep
	dep ensure -vendor-only

# Check
.PHONY: check
check: lint test

# Tests
.PHONY: test
test:
	$(info $(M) running go test…) @
	$Q $(GO) test -cover -race -v ./...

# Tools
.PHONY: lint
lint:
	$(info $(M) running $(GOLINT)…) @
	$Q GO_VENDOR=1 $(GOLINT) "--vendor" \
					"--disable=gotype" \
					"--disable=vetshadow" \
					"--disable=gocyclo" \
					"--fast" \
					"--json" \
					"./..."

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret


.PHONY: version
version:
	@echo $(VERSION)
