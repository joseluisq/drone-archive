DOCKER_IMAGE_TAG=joseluisq/drone-archive
DRONE_COMMIT_SHA ?= $(shell git rev-parse HEAD)
LABEL_SCHEMA_BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LABEL_SCHEMA_VCS_REF ?= $(shell git rev-parse --short HEAD)
GOOS=linux
BINARY_VERSION=0.1.0

install:
	@go version
	@go get -v golang.org/x/lint/golint
.PHONY: install

test:
	@go version
	@golint -set_exit_status ./...
	@go vet ./...
	@go test $$(go list ./... | grep -v /examples) \
		-v -timeout 30s -race -coverprofile=coverage.txt -covermode=atomic
.PHONY: test

build:
	@env \
		CGO_ENABLED=0 \
		GO111MODULE=on \
		GOOS=$(GOOS) \
			go build -v \
				-ldflags "-s -w -X main.version=$(BINARY_VERSION)" \
				-a -tags netgo \
				-o release/linux/amd64/drone-archive ./cmd
	@du -sh release/linux/amd64/.
.PHONY: build

coverage:
	@bash -c "bash <(curl -s https://codecov.io/bash)"
.PHONY: coverage

image-build:
	@docker build \
		--label org.label-schema.build-date=$(LABEL_SCHEMA_BUILD_DATE) \
		--label org.label-schema.vcs-ref=$(LABEL_SCHEMA_VCS_REF) \
		--file docker/alpine/Dockerfile \
		--tag $(DOCKER_IMAGE_TAG):local .
.PHONY: image-build

image-dryrun:
	@docker run --rm \
		-e PLUGIN_TAG=1 \
		-e PLUGIN_REPO=$(DOCKER_IMAGE_TAG) \
		-e DRONE_COMMIT_SHA=$(DRONE_COMMIT_SHA) \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
		--privileged \
			$(DOCKER_IMAGE_TAG):local \
				--daemon.debug \
				--dockerfile docker/alpine/Dockerfile \
				--dry-run
.PHONY: image-dryrun
