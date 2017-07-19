JENKINS_WORKSPACE ?=
CI_PROJECT_NAME := ci/.project_name
CI_IMAGE_TAG    := ci/.image_tag

ifdef JENKINS_WORKSPACE
DOCKER_CMD := docker run --rm -t \
					--volumes-from $(shell hostname) \
					-e SYMLINKS=$(JENKINS_WORKSPACE):/go/src/$(PACKAGE) \
					-e LOCAL_USER_ID=$(LOCAL_USER_ID) \
					-e IN_DOCKER=1 \
					-w /go/src/$(PACKAGE) \
					$(BUILD_IMG)
endif

ifneq ($(wildcard $(CI_PROJECT_NAME)),)
	PROJECT_NAME := -p $(shell cat $(CI_PROJECT_NAME))
endif

ifneq ($(wildcard $(CI_IMAGE_TAG)),)
	IMAGE_NAME := zenoss/zing-{{Name}}:$(shell cat $(CI_IMAGE_TAG))
endif

.PHONY: build
build: export IMAGE = $(IMAGE_NAME)
build: export COMMIT_SHA = $(shell git rev-parse HEAD)
build: $(CI_PROJECT_NAME) $(CI_IMAGE_TAG) $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE) $(PROJECT_NAME) build {{Name}}

.PHONY: unit-test
ifndef JENKINS_WORKSPACE
unit-test: test
else
unit-test:
	$(DOCKER_CMD) make test
endif

.PHONY: api-test
api-test: $(CI_PROJECT_NAME) $(CI_IMAGE_TAG) $(DOCKER_COMPOSE)
	@echo "Not implemented"

.PHONY: push
push: $(CI_IMAGE_TAG)
ifndef REMOTE_IMAGE
	@echo "REMOTE_IMAGE not set" 1>&2; exit 2
else
	@docker tag $(IMAGE_NAME) $(REMOTE_IMAGE)
	@docker push $(REMOTE_IMAGE)
endif

version.yaml:
ifndef REMOTE_IMAGE
	@echo "REMOTE_IMAGE not set" 1>&2; exit 2
else
	@cat ci/version-template.yaml | REMOTE_IMAGE=$(REMOTE_IMAGE) envsubst > version.yaml
endif

.PHONY: ci-clean
ci-clean:
	$(DOCKER_COMPOSE) $(PROJECT_NAME) down

.PHONY: ci-mrclean
ci-mrclean: ci-clean
	rm -f version.yaml
	rm -f $(CI_PROJECT_NAME) $(CI_IMAGE_TAG)
