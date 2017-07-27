DOCKER_COMPOSE_FILES = docker-compose.yml api/docker-compose.yml ci/docker-compose.yml

include .env
include Makefile

IMAGE_NAME   := ${SERVICE_IMAGE}:${IMAGE_TAG}
REMOTE_IMAGE := ${REGISTRY}/${SERVICE_IMAGE}:$(if ${REMOTE_TAG},${REMOTE_TAG},${IMAGE_TAG})

define DOCKER_COMPOSE_YAML
version: "3.3"
secrets:
    auth_key:
        file: $(ROOTDIR)/dev.secret
endef

ci/docker-compose.yml: export DOCKER_COMPOSE_YAML:=$(DOCKER_COMPOSE_YAML)
ci/docker-compose.yml:
	@echo "$$DOCKER_COMPOSE_YAML" > $@

.PHONY: IMAGE_NAME_VALID
IMAGE_NAME_VALID:
	@test ${SERVICE_IMAGE} || (echo "SERVICE_IMAGE must be defined"; false)
	@test ${IMAGE_TAG} || (echo "IMAGE_TAG must be defined"; false)

.PHONY: REMOTE_IMAGE_VALID
REMOTE_IMAGE_VALID:
	@test ${REGISTRY} || (echo "REGISTRY must be defined"; false)
	@test ${SERVICE_IMAGE} || (echo "SERVICE_IMAGE must be defined"; false)
	@test ${REMOTE_TAG}${IMAGE_TAG} || (echo "REMOTE_TAG or IMAGE_TAG must be defined"; false)

.PHONY: push
push: REMOTE_IMAGE_VALID IMAGE_NAME_VALID
	@docker tag $(IMAGE_NAME) $(REMOTE_IMAGE)
	@docker push $(REMOTE_IMAGE)

define VERSION_YAML
version: "3.2"
services:
    {{Name}}:
        image: $(REMOTE_IMAGE)
endef

version.yaml: export VERSION_YAML:=$(VERSION_YAML)
version.yaml: REMOTE_IMAGE_VALID
	@$(DOCKER_COMPOSE_BASE) -f <(echo "$$VERSION_YAML") config > $@

.PHONY: ci-clean
ci-clean:
	$(DOCKER_COMPOSE_BASE) down

.PHONY: ci-mrclean
ci-mrclean: ci-clean
	rm -f version.yaml
	rm -f ci/docker-compose.yml
