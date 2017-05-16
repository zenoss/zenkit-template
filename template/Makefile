GLIDE     := $(GOPATH)/bin/glide
GOAGEN    := vendor/github.com/goadesign/goa/goagen/goagen
DESIGNPKG := {{Package}}/design
DESIGN 	  := $(shell find design -name \*.go)

default: app

$(GLIDE):
	@curl https://glide.sh/get | sh

glide.lock: $(GLIDE)
	@$(GLIDE) up

$(GOAGEN): glide.lock $(GLIDE)
	@cd $(@D) && go build

app: $(DESIGN) | $(GOAGEN)
	@$(GOAGEN) app -d $(DESIGNPKG)
	@$(GOAGEN) controller -d $(DESIGNPKG)
	@$(GOAGEN) swagger -d $(DESIGNPKG)
