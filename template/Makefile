{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
include .env

SHELL                := /bin/bash
ROOTDIR              ?= $(CURDIR)
GOPATH               := $(shell go env GOPATH)
GO                   := $(shell command -v go 2> /dev/null)
GOFUMPT				 := $(shell command -v gofumpt 2> /dev/null)
GOLANGCI_LINT 		 := $(shell command -v golangci-lint 2> /dev/null)
REVIVE				 := $(shell command -v revive 2> /dev/null)
GINKGO				 := $(shell command -v ginkgo 2> /dev/null)
GOCOV				 := $(shell command -v gocov 2> /dev/null)
GOCOVXML			 := $(shell command -v gocov-xml 2> /dev/null)
DOCKER			     := $(shell command -v docker 2> /dev/null)
LOCAL_USER_ID        := $(shell id -u)
ZENKIT_BUILD_VERSION := 1.18.2
BUILD_IMG            := zenoss/zenkit-build:$(ZENKIT_BUILD_VERSION)
COVERAGE_DIR         := coverage
DOCKER_PARAMS        := \
	--rm \
	--volume $(ROOTDIR):/workspace:rw \
	--volume /var/run/docker.sock:/var/run/docker.sock \
	--env LOCAL_USER_ID=$(LOCAL_USER_ID) \
	--workdir /workspace
DOCKER_CMD           := docker run -t $(DOCKER_PARAMS) $(BUILD_IMG)

DOCKER_COMPOSE_BASE  := $(DOCKER) compose --progress=quiet
ifdef PROJECT_NAME
DOCKER_COMPOSE_BASE  += -p $(PROJECT_NAME)
endif

M = $(shell printf "\033[34;1m▶\033[0m")
RED = $(shell printf "\033[31;1m▶\033[0m")

.PHONY: default
default: init-containerized

.PHONY: docker-compose
docker-compose: $(DOCKER)

.PHONY: vendor
vendor:
	@echo "$(M) retrieving dependencies…"
	cd $(ROOTDIR) && $(GO) mod vendor
	@echo "$(M) dependencies have been vendored"

.PHONY: go-check
go-check:
ifeq ($(strip $(GOLANGCI_LINT)),)
	@echo "Warning: golangci-lint is not available on this system, please install it"
else
	@echo "Running golangci-lint..."
	-@$(GOLANGCI_LINT) run --presets "unused,performance,bugs" --tests=false -D "protogetter,contextcheck,musttag" ./...
endif

.PHONY: vuln-sarif
vuln-sarif:
	$(DOCKER_COMPOSE_BASE) -f docker-compose.test.yml run $(DOCKER_PARAMS) test make vuln-sarif-host

.PHONY: vuln-sarif-host
vuln-sarif-host: 
	$(shell GOMEMLIMIT=512MiB govulncheck -test=true -tags=integration -format=sarif ./... | jq ".runs[].results[].locations[].physicalLocation.region += {\"endLine\": $(shell wc -l < go.mod)}" > vuln.sarif)

.PHONY: fmt
fmt:
	@echo "$(M) gofumpt: running"
	@if [[ "$(shell $(GOFUMPT) -l -w .)" ]]; then \
		echo "$(RED) gofumpt: Please commit formatted files and if this runs on Jenkins rerun the job"; \
	else \
		echo "$(M) gofumpt: files look good"; \
	fi


.PHONY: lint
lint: 
	@echo "$(M) running revive linter…"
	$(REVIVE) \
		-config $(ROOTDIR)/.revive.toml \
		-formatter=friendly \
		-exclude $(ROOTDIR)/vendor/... \
		$(ROOTDIR)/...
	@echo "$(M) linted with revive linter"

.PHONY: build
build: export COMMIT_SHA ?= $(shell git rev-parse HEAD)
build: export GIT_BRANCH ?= $(shell git symbolic-ref HEAD | sed -e "s/^refs\/heads\///")
build: export PULL_REQUEST = ${ghprbPullLink}
build: | $(DOCKER)
	@$(DOCKER_COMPOSE_BASE) build {{Name}}

.PHONY: run
run: $(DOCKER)
	@$(DOCKER_COMPOSE_BASE) up --build

.PHONY: down
down: $(DOCKER)
	$(DOCKER_COMPOSE_BASE) -f docker-compose.test.yml down

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
		--race \
		--skip-package vendor \
		--junit-report=junit.xml
	@$(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	@$(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

.PHONY: unit-test-containerized
unit-test-containerized: export ZENKIT_BUILD_VERSION ?= $(ZENKIT_BUILD_VERSION)
unit-test-containerized: $(DOCKER)
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

.PHONY: mocks
mocks:
	mockery
