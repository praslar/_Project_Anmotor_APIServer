PROJECT_NAME=anmotor
BUILD_VERSION=$(shell cat VERSION)
DOCKER_IMAGE=$(PROJECT_NAME):$(BUILD_VERSION)
GO_BUILD_ENV=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on


.SILENT:

all: mod_tidy fmt vet test install 

build:
	$(GO_BUILD_ENV) go build -v -o $(PROJECT_NAME)-$(BUILD_VERSION).bin .

install:
	$(GO_BUILD_ENV) go install

vet:
	$(GO_BUILD_ENV) go vet $(GO_FILES)

fmt:
	$(GO_BUILD_ENV) go fmt $(GO_FILES)

mod_tidy:
	$(GO_BUILD_ENV) go mod tidy

test:
	$(GO_BUILD_ENV) go test $(GO_FILES) -cover -v

compose_prod: docker
	cd deployment/docker && BUILD_VERSION=$(BUILD_VERSION) docker-compose up

docker_prebuild: vet build
	mkdir -p deployment/docker/configs
	mv $(PROJECT_NAME)-$(BUILD_VERSION).bin deployment/docker/$(PROJECT_NAME).bin; \
	cp -R configs deployment/docker/;

docker_build:
	cd deployment/docker; \
	docker build -t $(DOCKER_IMAGE) .;

docker_postbuild:
	cd deployment/docker; \
	rm -rf $(PROJECT_NAME).bin 2> /dev/null;\
	rm -rf configs 2> /dev/null;

docker: docker_prebuild docker_build docker_postbuild