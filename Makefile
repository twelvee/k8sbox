override APP_NAME=k8sbox
override GO_VERSION=1.18

GOOS?=$(shell go env GOOS || echo linux)
GOARCH?=$(shell go env GOARCH || echo amd64)
CGO_ENABLED?=0

DOCKER_REGISTRY?=registry.github.com
DOCKER_IMAGE?=${DOCKER_REGISTRY}/k8s-box/${APP_NAME}
DOCKER_TAG?=current

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