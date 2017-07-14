{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
ZENKIT_VERSION       := 1.0
GINKGO               := $(shell command -v ginkgo 2> /dev/null)
GOAGEN               := vendor/github.com/goadesign/goa/goagen/goagen
PACKAGE              := {{$pkg}}
DESIGNPKG            := $(PACKAGE)/design
DESIGN 	             := $(shell find design -name \*.go)
DOCKER_COMPOSE       := /usr/local/bin/docker-compose
LOCAL_USER_ID        := $(shell id -u)
BUILD_IMG            := zenoss/zenkit-build:$(ZENKIT_VERSION)
COVERAGE_DIR         := coverage
DOCKER_CMD           := docker run --rm -t \
							-v $(CURDIR):/go/src/$(PACKAGE):rw \
							-v $(HOME)/.glide:/home/user/.glide:rw \
							-e LOCAL_USER_ID=$(LOCAL_USER_ID) \
							-w /go/src/$(PACKAGE) \
							$(BUILD_IMG)
GOAGEN_CMD           := $(DOCKER_CMD) $(GOAGEN)
GLIDE                := $(shell command -v glide 2> /dev/null)
ifndef GLIDE
GLIDE                := $(DOCKER_CMD) glide
endif
GOBINDATA            := $(shell command -v go-bindata 2> /dev/null)
ifndef GOBINDATA
GOBINDATA            := $(DOCKER_CMD) go-bindata
endif
IMAGE_NAME			 := zenoss/{{Name}}

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
	@mkdir -p $(HOME)/.glide
	@$(GLIDE) update --strip-vendor

vendor: glide.lock
	@mkdir -p $(HOME)/.glide
	@$(GLIDE) install --strip-vendor

$(GOAGEN): vendor
	@$(DOCKER_CMD) /bin/sh -c 'cd $(@D) && go build'

swagger/swagger.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN_CMD) swagger -d $(DESIGNPKG)
	@$(GOBINDATA) -ignore='swagger\.go' -pkg swagger -o $@ swagger/

resources/controller_reg.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN_CMD) gen -o resources --pkg-path=github.com/zenoss/zenkit/generate -d $(DESIGNPKG)

resources/app/%.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN_CMD) app -o resources -d $(DESIGNPKG)

resources/%.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN_CMD) controller --regen --pkg resources -o resources -d $(DESIGNPKG)

.PHONY: app
app: swagger/swagger.go resources/controller_reg.go resources/admin.go resources/app/controllers.go

.PHONY: run
run: vendor $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE) up --build

.PHONY: build
build: COMMIT_SHA := $(shell git rev-parse HEAD)
ifeq("$(wildcard tag.id)", "")
build: IMAGE := $(IMAGE_NAME)
else
build: IMAGE := $(IMAGE_NAME):$(shell cat tag.id)
endif
build: vendor $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE) build {{Name}}

.PHONY: test-containerized
test-containerized:
	@$(DOCKER_CMD) make test

.PHONY: test
ifndef GINKGO
test: test-containerized
else
test: COVERAGE_PROFILE := $(COVERAGE_DIR)/profile.out
test: COVERAGE_HTML    := $(COVERAGE_DIR)/index.html
test: COVERAGE_XML     := $(COVERAGE_DIR)/coverage.xml
test: app vendor
	@mkdir -p $(COVERAGE_DIR)
	@ginkgo -r \
		-cover \
		-covermode=count \
		--skipPackage vendor
	@gocovmerge $$(find . -name \*.coverprofile) > $(COVERAGE_PROFILE)
	@go tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	@gocov convert $(COVERAGE_PROFILE) | gocov-xml > $(COVERAGE_XML)
endif

.PHONY: api-test
api-test:
	@echo "not implemented"

.PHONY: version
ifndef REMOTE_IMAGE
version:
	@echo "REMOTE_IMAGE not set" 1>&2
	@exit 1
else
version:
	@cat version-template.yaml | envsubst > version.yaml
endif

.PHONY: clean
clean:
	rm -f resources/controller_reg.go
	rm -rf resources/app
	rm -rf swagger
	rm -rf $(COVERAGE_DIR)
	rm -f tag.id version.yaml
	$(DOCKER_COMPOSE) down

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

