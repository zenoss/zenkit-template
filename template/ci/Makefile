DOCKER_COMPOSE_FILES ?= docker-compose.yml hooks/docker-compose.yml ci/docker-compose.yml

include Makefile

IMAGE_NAME   := ${SERVICE_IMAGE}:${IMAGE_TAG}

# This is only used for publishing to GCR from jenkins
SERVICE_NAME := $(subst zenosszing/,,${SERVICE_IMAGE})
ifdef ZING_REPOSITORY
ALT_SERVICE_IMAGE := $(subst zenosszing,${ZING_REPOSITORY},${SERVICE_IMAGE})
REMOTE_IMAGE := ${REGISTRY}/${ALT_SERVICE_IMAGE}:$(if ${REMOTE_TAG},${REMOTE_TAG},${IMAGE_TAG})
else
REMOTE_IMAGE := ${REGISTRY}/${SERVICE_IMAGE}:$(if ${REMOTE_TAG},${REMOTE_TAG},${IMAGE_TAG})
endif

# FIXME: This can be probably be removed once zenkit no longer uses the auth_key.secrets
define CI_DOCKER_COMPOSE_YAML
version: "3.3"
secrets:
    auth_key:
        file: $(ROOTDIR)/dev.secret
endef

ci/docker-compose.yml: export CI_DOCKER_COMPOSE_YAML:=$(CI_DOCKER_COMPOSE_YAML)
ci/docker-compose.yml:
	@echo "$$CI_DOCKER_COMPOSE_YAML" > $@


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

.PHONY: gpush
gpush: REMOTE_IMAGE_VALID IMAGE_NAME_VALID GCR_CREDS_VALID
	gcloud auth activate-service-account ${GCR_ACCOUNT} --key-file ${GCR_KEY_FILE}
	docker tag $(IMAGE_NAME) $(REMOTE_IMAGE)
	gcloud docker -- push $(REMOTE_IMAGE)

.PHONY: GCR_CREDS_VALID
GCR_CREDS_VALID:
	@test ${GCR_ACCOUNT} || (echo "GCR_ACCOUNT must be defined"; false)
	@test ${GCR_KEY_FILE} || (echo "GCR_KEY_FILE must be defined"; false)
	@test ${ZING_REPOSITORY} || (echo "ZING_REPOSITORY must be defined"; false)

define VERSION_ENV
IMAGE_NAME=${SERVICE_NAME}:${IMAGE_TAG}
endef

ci/version.env: export VERSION_ENV:=$(VERSION_ENV)
ci/version.env: IMAGE_NAME_VALID
	@echo "$$VERSION_ENV" > $@

.PHONY: clean
clean::
	rm -f ci/version.env
	rm -f ci/docker-compose.yml
