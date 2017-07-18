CI_PROJECT_NAME := ci/.project_name
CI_IMAGE_TAG    := ci/.image_tag

ifneq ($(wildcard $(CI_PROJECT_NAME)),)
	PROJECT_NAME := -p $(shell cat $(CI_PROJECT_NAME))
endif

ifneq ($(wildcard $(CI_IMAGE_NAME)),)
	IMAGE_NAME := zenoss/zing-{{Name}}:$(shell cat $(CI_IMAGE_TAG))
	CI_IMAGE_NAME := zenoss/ci-{{Name}}:$(shell cat $(CI_IMAGE_TAG))
else
	IMAGE_NAME := zenoss/zing-{{Name}}
	CI_IMAGE_NAME := zenoss/ci-{{Name}}
endif

.PHONY: build
build: export IMAGE = $(IMAGE_NAME)
build: export COMMIT_SHA = $(shell git rev-parse HEAD)
build: $(CI_PROJECT_NAME) $(CI_IMAGE_TAG) $(DOCKER_COMPOSE)
	@$(DOCKER_COMPOSE) $(PROJECT_NAME) build {{Name}}

.PHONY: unit-test
unit-test: export IMAGE = $(CI_IMAGE_NAME)
	@$(DOCKER_COMPOSE) $(PROJECT_NAME) -f ci/docker-compose.yml build
	@$(DOCKER_COMPOSE) $(PROJECT_NAME) -f ci/docker-compose.yml run {{Name}} make test

.PHONY: api-test
api-test: export IMAGE = $(IMAGE_NAME)
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
	rm -f $(CI_PROJECT_NAME) $(CI_IMAGE_TAG)
	$(DOCKER_COMPOSE) $(PROJECT_NAME) down

.PHONY: ci-mrclean
ci-mrclean: ci-clean
	rm -f version.yaml
	docker rmi -f $(CI_IMAGE_NAME)