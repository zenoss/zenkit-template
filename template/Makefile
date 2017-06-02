{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
GLIDE     := $(GOPATH)/bin/glide
GOAGEN    := vendor/github.com/goadesign/goa/goagen/goagen
DESIGNPKG := {{$pkg}}/design
DESIGN 	  := $(shell find design -name \*.go)

default: app

$(GLIDE):
	@curl https://glide.sh/get | sh

glide.lock: $(GLIDE)
	@$(GLIDE) up

$(GOAGEN): glide.lock $(GLIDE)
	@cd $(@D) && go build

app: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN) gen -o resources --pkg-path=github.com/zenoss/zenkit/generate -d $(DESIGNPKG)
	@$(GOAGEN) app -o resources -d $(DESIGNPKG)
	@$(GOAGEN) controller --pkg resources -o resources -d $(DESIGNPKG)
	@$(GOAGEN) swagger -d $(DESIGNPKG)
