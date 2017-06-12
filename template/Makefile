{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
GLIDE                := $(GOPATH)/bin/glide
GOAGEN               := vendor/github.com/goadesign/goa/goagen/goagen
DESIGNPKG            := {{$pkg}}/design
DESIGN 	             := $(shell find design -name \*.go)
DOCKER_COMPOSE       := /usr/local/bin/docker-compose

DOCKER_COMPOSE_VERSION := 1.13.0

.PHONY: default
default: app run

.PHONY: docker-compose
docker-compose: $(DOCKER_COMPOSE)

$(DOCKER_COMPOSE):
	@if [ ! -w $(@D) ]; then echo 'No docker-compose found. Please run "sudo make docker-compose" to install it.'; exit 2; else true; fi
	@curl -L https://github.com/docker/compose/releases/download/$(DOCKER_COMPOSE_VERSION)/docker-compose-`uname -s`-`uname -m` > $@
	@chmod +x $@

$(GLIDE):
	@curl https://glide.sh/get | sh

glide.lock: $(GLIDE)
	@$(GLIDE) up

$(GOAGEN): glide.lock $(GLIDE)
	@cd $(@D) && go build

swagger/swagger.%: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN) swagger -d $(DESIGNPKG)

resources/controller_reg.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN) gen -o resources --pkg-path=github.com/zenoss/zenkit/generate -d $(DESIGNPKG)

resources/app/%.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN) app -o resources -d $(DESIGNPKG)

resources/%.go: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN) controller --regen --pkg resources -o resources -d $(DESIGNPKG)

.PHONY: app
app: swagger/swagger.json resources/controller_reg.go resources/metrics.go resources/app/controllers.go

run: glide.lock $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE) up

start: glide.lock $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE) up -d

refresh: glide.lock $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE) up -d --build --no-deps {{Name}}

.PHONY: clean
clean:
	rm -f resources/controller_reg.go
	rm -rf resources/app
	rm -rf swagger

.PHONY: mrclean
mrclean: clean
	rm -rf glide.lock
	rm -rf vendor
