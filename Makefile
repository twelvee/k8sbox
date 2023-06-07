override APP_NAME=k8sbox
override GO_VERSION=1.19

GOOS?=$(shell go env GOOS || echo linux)
GOARCH?=$(shell go env GOARCH || echo amd64)
CGO_ENABLED?=0

DOCKER_IMAGE?=twelvee/${APP_NAME}
DOCKER_TAG?=v1.2.0

ifeq (, $(shell which docker))
$(error "Binary docker not found in $(PATH)")
endif

.PHONY: all
all: cleanup vendor

.PHONY: cleanup
cleanup:
	@rm ${PWD}/bin/${APP_NAME} || true
	@rm -r ${PWD}/vendor || true

.PHONY: vendor
vendor:
	@rm -r ${PWD}/vendor || true
	@docker run --rm \
		-v ${PWD}:/project \
		-w /project \
		golang:${GO_VERSION} \
			go mod tidy
	@docker run --rm \
		-v ${PWD}:/project \
		-w /project \
		golang:${GO_VERSION} \
			go mod vendor

.PHONY: build
build:
	@rm ${PWD}/bin/${APP_NAME} || true
	@docker run --rm \
		-v ${PWD}:/project \
		-w /project \
		-e GOOS=${GOOS} \
		-e GOARCH=${GOARCH} \
		-e CGO_ENABLED=${CGO_ENABLED} \
		-e GO111MODULE=on \
		golang:${GO_VERSION} \
			go build \
				-mod vendor \
				-o /project/bin/${APP_NAME} \
				-v /project/cmd/${APP_NAME}

.PHONY: docker-build
docker-build:
	@docker build \
		-f ${PWD}/build/docker/${APP_NAME}.Dockerfile \
		-t ${DOCKER_IMAGE}:${DOCKER_TAG} .

.PHONY: docker-push
docker-push:
	@docker login -u ${DOCKER_USER} -p ${DOCKER_PASSWORD}
	@docker push ${DOCKER_IMAGE}:${DOCKER_TAG}