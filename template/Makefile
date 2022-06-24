{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
include .env

SHELL                := /bin/bash
ROOTDIR              ?= $(CURDIR)
GOPATH               := $(shell go env GOPATH)
GO                   := $(shell command -v go 2> /dev/null)
GOFMT 				 := $(shell command -v gofumpt 2> /dev/null)
PACKAGE              := {{$pkg}}
BIN                  := $(GOPATH)/bin
BASE                 := $(GOPATH)/src/$(PACKAGE)
DOCKER_COMPOSE       := /usr/local/bin/docker-compose
LOCAL_USER_ID        := $(shell id -u)
ZENKIT_BUILD_VERSION := 1.12.0
BUILD_IMG            := zenoss/zenkit-build:$(ZENKIT_BUILD_VERSION)
COVERAGE_DIR         := coverage
DOCKER_PARAMS        := --rm -v $(ROOTDIR):/go/src/$(PACKAGE):rw \
							-v /var/run/docker.sock:/var/run/docker.sock \
							-e LOCAL_USER_ID=$(LOCAL_USER_ID) \
							-w /go/src/$(PACKAGE)
DOCKER_CMD           := docker run -t $(DOCKER_PARAMS) $(BUILD_IMG)

DOCKER_COMPOSE_BASE  := $(DOCKER_COMPOSE)
ifdef PROJECT_NAME
DOCKER_COMPOSE_BASE  += -p $(PROJECT_NAME)
endif

M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: default
default: init-containerized

.PHONY: docker-compose
docker-compose: $(DOCKER_COMPOSE)

$(DOCKER_COMPOSE): DOCKER_COMPOSE_VERSION := 1.14.0
$(DOCKER_COMPOSE):
	@if [ ! -w $(@D) ]; then echo 'No docker-compose found. Please run "sudo make docker-compose" to install it.'; exit 2; else true; fi
	@curl -L https://github.com/docker/compose/releases/download/$(DOCKER_COMPOSE_VERSION)/docker-compose-`uname -s`-`uname -m` > $@
	@chmod +x $@

$(BIN):
	@mkdir -p $@
$(BIN)/%: $(BIN) | $(BASE)
	@[ -x $@ ] || echo "$(M)" Installing $* && go install $(REPOSITORY)

GINKGO = $(BIN)/ginkgo
$(BIN)/ginkgo: REPOSITORY=github.com/onsi/ginkgo/v2/ginkgo@latest

GOSTATICCHECK = $(BIN)/staticcheck
$(BIN)/staticcheck: REPOSITORY=honnef.co/go/tools/cmd/staticcheck@latest

REVIVE = $(BIN)/revive
$(BIN)/revive: REPOSITORY=github.com/mgechev/revive@latest

GOCOV = $(BIN)/gocov
$(BIN)/gocov: REPOSITORY=github.com/axw/gocov/gocov@latest

GOCOVXML = $(BIN)/gocov-xml
$(BIN)/gocov-xml: REPOSITORY=github.com/AlekSi/gocov-xml@latest

.PHONY: vendor
vendor:
	@echo "$(M) retrieving dependencies…"
	cd $(ROOTDIR) && $(GO) mod vendor
	@echo "$(M) dependencies have been vendored"

.PHONY: fmt
fmt:
	@echo "$(M) running gofumpt…"
	$(GOFMT) -w -l $(ROOTDIR)
	@echo "$(M) formatted with gofumpt"

.PHONY: lint
lint: 
	@echo "$(M) running revive linter…"
	$(REVIVE) \
		-set_exit_status \
		-config $(ROOTDIR)/.revive.toml \
		-formatter=friendly \
		-exclude $(ROOTDIR)/vendor/... \
		$(ROOTDIR)/...
	@echo "$(M) linted with revive linter"

.PHONY: build
build: export COMMIT_SHA ?= $(shell git rev-parse HEAD)
build: export GIT_BRANCH ?= $(shell git symbolic-ref HEAD | sed -e "s/^refs\/heads\///")
build: export PULL_REQUEST = ${ghprbPullLink}
build: | $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE_BASE) build {{Name}}

.PHONY: run
run: $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE_BASE) up --build

.PHONY: down
down: $(DOCKER_COMPOSE)
	$(DOCKER_COMPOSE_BASE) down

.PHONY: check test tests
check test tests: unit-test-containerized

.PHONY: unit-test
unit-test: COVERAGE_PROFILE := coverprofile.out
unit-test: COVERAGE_HTML    := $(COVERAGE_DIR)/index.html
unit-test: COVERAGE_XML     := $(COVERAGE_DIR)/coverage.xml
unit-test: fmt
	@mkdir -p $(COVERAGE_DIR)
	@$(GINKGO) \
		run \
		-r \
		--mod vendor \
		--cover \
		--covermode=count \
		--skip-package vendor \
		--junit-report=junit.xml
	@$(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	@$(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

.PHONY: unit-test-containerized
unit-test-containerized: export ZENKIT_BUILD_VERSION ?= $(ZENKIT_BUILD_VERSION)
unit-test-containerized: $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE_BASE) -f docker-compose.test.yml run $(DOCKER_PARAMS) test make unit-test
	@$(DOCKER_COMPOSE_BASE) -f docker-compose.test.yml down

.PHONY: clean
clean:: down
	rm -f junit.xml
	rm -f coverprofile.out
	rm -rf $(COVERAGE_DIR)

.PHONY: mrclean
mrclean: clean
	rm -rf vendor

.git: vendor
	@git init
	@git add .; git commit -m "Initial commit"

.PHONY: init
init: .git

.PHONY: init-containerized
init-containerized:
	@$(DOCKER_CMD) make vendor
