GLIDE := $(GOPATH)/bin/glide
GOAGEN := vendor/github.com/goadesign/goa/goagen/goagen

default: $(GOAGEN)

$(GLIDE):
	@curl https://glide.sh/get | sh

glide.lock: $(GLIDE)
	@$(GLIDE) up

$(GOAGEN): glide.lock $(GLIDE)
	@cd $(@D) && go build
