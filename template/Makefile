{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
GLIDE                := $(GOPATH)/bin/glide
GINKGO               := $(GOPATH)/bin/ginkgo
GO_BINDATA           := $(GOPATH)/bin/go-bindata
GOAGEN               := vendor/github.com/goadesign/goa/goagen/goagen
DESIGNPKG            := {{$pkg}}/design
DESIGN 	             := $(shell find design -name \*.go)
DOCKER_COMPOSE       := /usr/local/bin/docker-compose
GIT_FLOW 			 := $(shell git flow version &>/dev/null && echo git flow)

DOCKER_COMPOSE_VERSION := 1.13.0

.PHONY: default
default: app init

.PHONY: docker-compose
docker-compose: $(DOCKER_COMPOSE)

$(DOCKER_COMPOSE):
	@if [ ! -w $(@D) ]; then echo 'No docker-compose found. Please run "sudo make docker-compose" to install it.'; exit 2; else true; fi
	@curl -L https://github.com/docker/compose/releases/download/$(DOCKER_COMPOSE_VERSION)/docker-compose-`uname -s`-`uname -m` > $@
	@chmod +x $@

$(GLIDE):
	@curl https://glide.sh/get | sh

$(GINKGO):
	@go get -u github.com/onsi/ginkgo/ginkgo

$(GO_BINDATA):
	@go get -u github.com/jteeuwen/go-bindata/go-bindata

glide.lock: $(GLIDE)
	@$(GLIDE) up

$(GOAGEN): glide.lock $(GLIDE)
	@cd $(@D) && go build

swagger/swagger.go: $(DESIGN) | $(GOAGEN) $(GO_BINDATA)
	@$(GOAGEN) swagger -d $(DESIGNPKG)
	@$(GO_BINDATA) -ignore='swagger\.go' -pkg swagger -o $@ swagger/

resources/controller_reg.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN) gen -o resources --pkg-path=github.com/zenoss/zenkit/generate -d $(DESIGNPKG)

resources/app/%.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN) app -o resources -d $(DESIGNPKG)

resources/%.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN) controller --regen --pkg resources -o resources -d $(DESIGNPKG)

.PHONY: app
app: swagger/swagger.go resources/controller_reg.go resources/metrics.go resources/app/controllers.go

.PHONY: run
run: glide.lock $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE) up --build

.PHONY: build
build: glide.lock $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE) build {{Name}}

.PHONY: test
test: app $(GINKGO)
	@$(GINKGO) -r -race -cover -skipPackage=vendor

.PHONY: clean
clean:
	rm -f resources/controller_reg.go
	rm -rf resources/app
	rm -rf swagger
	$(DOCKER_COMPOSE) down

.PHONY: mrclean
mrclean: clean
	rm -rf glide.lock
	rm -rf vendor

.git:
ifndef GIT_FLOW
	@echo 'No git flow found. Please run "sudo apt install -y git-flow" to install it.'; exit 2
endif
	@git init
	@git add .; git commit -m "Initial commit"
	@git flow init -d

.PHONY: init
init: .git

