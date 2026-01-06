#
# Zenoss CONFIDENTIAL
# -------------------
#
# This software Copyright (c) Zenoss, Inc. 2025
# All Rights Reserved.
#
# NOTICE:  All information contained herein is, and remains the property of
# Zenoss Incorporated and its suppliers, if any.  The intellectual and
# technical concepts contained herein are owned and proprietary to Zenoss
# Incorporated and its suppliers and may be covered by U.S. and Foreign
# Patents, patents in process, and are protected by U.S. and foreign trade
# secret or copyright law. Dissemination of this information or reproduction of
# any this material herein is strictly forbidden unless prior written
# permission by an authorized officer is obtained from Zenoss Incorporated.
#
include .env

SHELL                := $(shell command -v bash 2> /dev/null)
GO                   := $(shell command -v go 2> /dev/null)
GOLANGCI_LINT        := $(shell command -v golangci-lint 2> /dev/null)
GOVULNCHECK          := $(shell command -v govulncheck 2> /dev/null)
GOCOV                := $(shell command -v gocov 2> /dev/null)
GOCOVXML             := $(shell command -v gocov-xml 2> /dev/null)
MOCKERY              := $(shell command -v mockery 2> /dev/null)
DOCKER               := $(shell command -v docker 2> /dev/null)
LOCAL_USER_ID        := $(shell id -u)
ZENKIT_BUILD_VERSION := 1.25.1
BUILD_IMG            := gcr.io/zing-registry-188222/zenkit-build:$(ZENKIT_BUILD_VERSION)
COVERAGE_DIR         := coverage
DOCKER_PARAMS        := --rm \
                        --volume $(CURDIR):/workspace:rw \
                        --volume /var/run/docker.sock:/var/run/docker.sock \
                        --env LOCAL_USER_ID=$(LOCAL_USER_ID) \
                        --workdir /workspace
DOCKER_COMPOSE_BASE  := $(DOCKER) compose --progress=quiet
ifdef PROJECT_NAME
DOCKER_COMPOSE_BASE  += -p $(PROJECT_NAME)
endif

M = $(shell printf "\033[34;1m▶\033[0m")
RED = $(shell printf "\033[31;1m▶\033[0m")

.PHONY: default
default: check-containerized unit-test-containerized

.PHONY: check-containerized
check-containerized: export ZENKIT_BUILD_VERSION ?= $(ZENKIT_BUILD_VERSION)
check-containerized: $(DOCKER)
	@echo "$(M) check with docker compose…"
	@$(DOCKER_COMPOSE_BASE) --profile=check run $(DOCKER_PARAMS) check make check

.PHONY: unit-test-containerized
unit-test-containerized: export ZENKIT_BUILD_VERSION ?= $(ZENKIT_BUILD_VERSION)
unit-test-containerized: $(DOCKER)
	@echo "$(M) unit-test with docker compose…"
	@$(DOCKER_COMPOSE_BASE) --profile=test run $(DOCKER_PARAMS) test make unit-test

.PHONY: check
check: lint vuln

.PHONY: lint
lint:
ifeq ($(strip $(GOLANGCI_LINT)),)
	@echo "$(RED) Warning: golangci-lint is not available on this system, please install it"
else
	@echo "$(M) golangci-lint: linting…"
	@$(GOLANGCI_LINT) run --config .golangci.yml ./...
endif

.PHONY: vuln
vuln:
ifeq ($(strip $(GOVULNCHECK)),)
	@echo "$(RED) Warning: govulncheck is not available on this system, please install it"
else
	@echo "$(M) govulncheck: scanning for vulnerabilities…"
	@GOMEMLIMIT=512MiB $(GOVULNCHECK) \
		-test=true \
		-tags=integration \
		-format=sarif \
		./... > govulncheck.sarif
endif

.PHONY: test
test: unit-test # integration-test

.PHONY: unit-test
unit-test: COVERAGE_PROFILE := coverprofile.out
unit-test: COVERAGE_HTML    := $(COVERAGE_DIR)/index.html
unit-test: COVERAGE_XML     := $(COVERAGE_DIR)/coverage.xml
unit-test:
	@echo "$(M) ginkgo: running tests…"
	@mkdir -p $(COVERAGE_DIR)
	@$(GO) tool ginkgo \
		run \
		-r \
		--mod vendor \
		--cover \
		--covermode atomic \
		--race \
		--skip-package vendor,mocks \
		--tags integration \
		--junit-report junit.xml
	@$(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	@$(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

.PHONY: mocks
mocks:
ifeq ($(strip $(MOCKERY)),)
	@echo "$(RED) Warning: mockery is not available on this system, please install it"
else
	@echo "$(M) mockery: generating mocks…"
	@$(MOCKERY) --config .mockery.yaml
endif

.PHONY: build
build: export COMMIT_SHA ?= $(shell git rev-parse HEAD)
build: export GIT_BRANCH ?= $(shell git symbolic-ref HEAD | sed -e "s/^refs\/heads\///")
build: export PULL_REQUEST = ${ghprbPullLink}
build: $(DOCKER)
	@echo "$(M) building with docker compose…"
	@$(DOCKER_COMPOSE_BASE) --profile=build build {{Name}}

.PHONY: run
run: $(DOCKER)
	@echo "$(M) running with docker compose…"
	@$(DOCKER_COMPOSE_BASE) --profile=run up --build

.PHONY: down
down: $(DOCKER)
	@echo "$(M) stopping all containers…"
	$(DOCKER_COMPOSE_BASE) --profile "*" down

.PHONY: clean
clean:: down
	@echo "$(M) cleaning up…"
	rm -f golangci-lint.html golangci-lint.sarif golangci-lint.xml
	rm -f govulncheck.sarif
	rm -f junit.xml
	rm -f coverprofile.out
	rm -rf $(COVERAGE_DIR)
