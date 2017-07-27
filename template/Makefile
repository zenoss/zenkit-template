{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
SHELL                := /bin/bash
ROOTDIR              ?= $(CURDIR)
ZENKIT_VERSION       := 1.3
GINKGO               := $(shell command -v ginkgo 2> /dev/null)
GOAGEN               := vendor/github.com/goadesign/goa/goagen/goagen
PACKAGE              := {{$pkg}}
DESIGNPKG            := $(PACKAGE)/design
DESIGN 	             := $(shell find design -name \*.go)
DOCKER_COMPOSE       := /usr/local/bin/docker-compose
LOCAL_USER_ID        := $(shell id -u)
BUILD_IMG            := zenoss/zenkit-build:$(ZENKIT_VERSION)
COVERAGE_DIR         := coverage

DOCKER_PARAMS        := --rm -v $(ROOTDIR):/go/src/$(PACKAGE):rw \
							-e LOCAL_USER_ID=$(LOCAL_USER_ID) \
							-w /go/src/$(PACKAGE) 
DOCKER_CMD           := docker run -t $(DOCKER_PARAMS) $(BUILD_IMG)

DOCKER_COMPOSE_BASE  := $(DOCKER_COMPOSE)
ifdef PROJECT_NAME
DOCKER_COMPOSE_BASE  += -p $(PROJECT_NAME)
endif
DOCKER_COMPOSE_FILES ?= docker-compose.yml api/docker-compose.yml

GOAGEN_CMD           := $(DOCKER_CMD) $(GOAGEN)
GLIDE                := $(shell command -v glide 2> /dev/null)
ifndef GLIDE
GLIDE                := $(DOCKER_CMD) glide
endif
GOBINDATA            := $(shell command -v go-bindata 2> /dev/null)
ifndef GOBINDATA
GOBINDATA            := $(DOCKER_CMD) go-bindata
endif

.PHONY: default
default: app init

.PHONY: docker-compose
docker-compose: $(DOCKER_COMPOSE)

$(DOCKER_COMPOSE): DOCKER_COMPOSE_VERSION := 1.14.0
$(DOCKER_COMPOSE):
	@if [ ! -w $(@D) ]; then echo 'No docker-compose found. Please run "sudo make docker-compose" to install it.'; exit 2; else true; fi
	@curl -L https://github.com/docker/compose/releases/download/$(DOCKER_COMPOSE_VERSION)/docker-compose-`uname -s`-`uname -m` > $@
	@chmod +x $@

glide.lock: glide.yaml
	@mkdir -p .glide
	@$(GLIDE) update --strip-vendor

vendor: glide.lock
	@mkdir -p .glide
	@$(GLIDE) install --strip-vendor

$(GOAGEN): vendor
	@$(DOCKER_CMD) /bin/bash -c 'cd $(@D) && go build'

swagger/swagger.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN_CMD) swagger -d $(DESIGNPKG)
	@$(GOBINDATA) -ignore='swagger\.go' -pkg swagger -o $@ swagger/

resources/controller_reg.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN_CMD) gen -o resources --pkg-path=github.com/zenoss/zenkit/generate -d $(DESIGNPKG)

resources/app/%.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN_CMD) app -o resources -d $(DESIGNPKG)

resources/%.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN_CMD) controller --regen --pkg resources -o resources -d $(DESIGNPKG)

ifndef GINKGO
api/%_hooks: api/%_hooks.go
	@$(DOCKER_CMD) go build -o $@ $<
else
api/%_hooks: api/%_hooks.go
	@go build -o $@ $<
endif

.PHONY: app
app: swagger/swagger.go resources/controller_reg.go resources/admin.go resources/app/controllers.go

.PHONY: build
build: export COMMIT_SHA ?= $(shell git rev-parse HEAD)
build: export GIT_BRANCH ?= $(shell git symbolic-ref HEAD | sed -e "s/^refs\/heads\///")
build: export PULL_REQUEST = ${ghprbPullLink}
build: $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE_BASE) build {{Name}}

.PHONY: run
run: vendor $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE_BASE) up --build

.PHONY: unit-test-containerized
unit-test-containerized:
	@$(DOCKER_CMD) make unit-test

.PHONY: unit-test
ifndef GINKGO
unit-test: unit-test-containerized
else
unit-test: COVERAGE_PROFILE := $(COVERAGE_DIR)/profile.out
unit-test: COVERAGE_HTML    := $(COVERAGE_DIR)/index.html
unit-test: COVERAGE_XML     := $(COVERAGE_DIR)/coverage.xml
unit-test: 
	@mkdir -p $(COVERAGE_DIR)
	@ginkgo -r \
		-cover \
		-covermode=count \
		--skipPackage vendor
	@gocovmerge $$(find . -name \*.coverprofile) > $(COVERAGE_PROFILE)
	@go tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	@gocov convert $(COVERAGE_PROFILE) | gocov-xml > $(COVERAGE_XML)
endif

.PHONY: api-test-containerized
api-test-containerized: DOCKER_CMD:=$(DOCKER_COMPOSE_BASE) $(addprefix -f ,$(DOCKER_COMPOSE_FILES)) run $(DOCKER_PARAMS) api
api-test-containerized: $(DOCKER_COMPOSE_FILES) api/{{Name}}_hooks
	@$(DOCKER_CMD) dredd ./swagger/swagger.yaml {{Name}}:{{Port}} \
		-a go -f api/{{Name}}_hooks \
		-r junit -o api/junit.xml

.PHONY: api-test
api-test: DOCKER_COMPOSE_FILES:=docker-compose.yml api/docker-compose.yml
api-test: api-test-containerized

.PHONY: test
test: unit-test api-test

.PHONY: clean
clean:
	rm -f resources/controller_reg.go
	rm -rf resources/app
	rm -rf swagger
	rm -rf $(COVERAGE_DIR)
	$(DOCKER_COMPOSE_BASE) down

.PHONY: mrclean
mrclean: clean
	rm -rf glide.lock
	rm -rf vendor

.git:
	@git init
	@git add .; git commit -m "Initial commit"

.PHONY: init
init: .git

local-dev:
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega
	go get github.com/wadey/gocovmerge
	go get github.com/axw/gocov/gocov
	go get github.com/AlekSi/gocov-xml
	curl https://glide.sh/get | sh
