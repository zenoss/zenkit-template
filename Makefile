GLIDE := glide
GLIDE_PATH := $(shell command -v $(GLIDE) 2> /dev/null)
GOAGEN := ./vendor/github.com/goadesign/goa/goagen/goagen

default: $(GOAGEN)

.PHONY: $(GLIDE)
$(GLIDE):
ifndef GLIDE_PATH
	curl https://glide.sh/get | sh
endif

glide.lock: $(GLIDE)
	$(GLIDE) install

$(GOAGEN): glide.lock $(GLIDE)
	$(GLIDE) up
	cd $(@D) && go build
