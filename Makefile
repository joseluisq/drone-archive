# Configuration

BINARY_VERSION ?= 0.0.0
BINARY_OUTPUT_PATH ?= release/linux/amd64/drone-archive
BINARY_BUILD_DATETIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

DOCKER_IMAGE_NAME ?= joseluisq/drone-archive
DOCKER_FILE ?= docker/alpine/Dockerfile

DRONE_COMMIT_SHA ?= $(shell git rev-parse HEAD)
LABEL_SCHEMA_VCS_REF ?= $(shell git rev-parse --short HEAD)


# Development

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

coverage:
	@bash -c "bash <(curl -s https://codecov.io/bash)"
.PHONY: coverage

build:
	@env \
		CGO_ENABLED=0 \
		GO111MODULE=on \
			go build -v \
				-ldflags "-s -w -X main.version=$(BINARY_VERSION)" \
				-a -o $(BINARY_OUTPUT_PATH) ./cmd
	@du -sh $(BINARY_OUTPUT_PATH)
.PHONY: build

docker.build:
	@docker build \
		--build-arg DRONE_ARCHIVE_VERSION=$(BINARY_VERSION) \
		--label org.label-schema.build-date=$(BINARY_BUILD_DATETIME) \
		--label org.label-schema.vcs-ref=$(LABEL_SCHEMA_VCS_REF) \
		--file $(DOCKER_FILE) \
		--tag $(DOCKER_IMAGE_NAME):local .
.PHONY: docker.build

docker.tar:
	@docker run --rm \
		-e PLUGIN_CHECKSUM=true \
		-e PLUGIN_CHECKSUM_DESTINATION=$(BINARY_OUTPUT_PATH).CHECKSUM.tar.gz.txt \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
			$(DOCKER_IMAGE_NAME):local \
				--src $(BINARY_OUTPUT_PATH) \
				--dest $(BINARY_OUTPUT_PATH).tar.gz
.PHONY: docker.tar

docker.zip:
	@docker run --rm \
		-e PLUGIN_CHECKSUM=true \
		-e PLUGIN_CHECKSUM_DESTINATION=$(BINARY_OUTPUT_PATH).CHECKSUM.zip.txt \
		-e PLUGIN_FORMAT=zip \
		-v $(PWD):$(PWD) \
		-w $(PWD) \
			$(DOCKER_IMAGE_NAME):local \
				--src $(BINARY_OUTPUT_PATH) \
				--dest $(BINARY_OUTPUT_PATH).zip
.PHONY: docker.zip


# Production

prod.release:
	@go version
	@env \
		CGO_ENABLED=0 \
		GO111MODULE=on \
			go build -v \
				-ldflags "\
					-s -w \
					-X 'main.versionNumber=$(BINARY_VERSION)' \
					-X 'main.buildTime=$(BINARY_BUILD_DATETIME)'\
				" \
				-a -o $(BINARY_OUTPUT_PATH) ./cmd
	@du -sh $(BINARY_OUTPUT_PATH)
.PHONY: prod.release

prod.executable:
	@$(BINARY_OUTPUT_PATH) --help
	@$(BINARY_OUTPUT_PATH) --version
.PHONY: prod.executable
