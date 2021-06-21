APP_NAME := filer
DOCKER_REPO := docker.olymp.alabor.me/swissmanu/

FILER_UI_PATH ?= ./ui/public
FILER_INBOX_PATH ?= ./example/inbox
FILER_DATA_PATH ?= ./example/target
FILER_RULES_PATH ?= ./example/rules.yml

install:
	cd ui && yarn install
.PHONY: install

start-server:
	FILER_UI_PATH=$(FILER_UI_PATH) \
	FILER_INBOX_PATH=$(FILER_INBOX_PATH) \
	FILER_DATA_PATH=$(FILER_DATA_PATH) \
	FILER_RULES_PATH=$(FILER_RULES_PATH) \
	go run cmd/filer/main.go
.PHONY: start-server

start-ui:
	cd ui && yarn dev
.PHONY: start-ui

build: clean build-server build-ui

build-server:
	GOOS=linux GOARCH=arm GOARM=7 go build -o build/linux/arm/v7/${APP_NAME} cmd/${APP_NAME}/main.go
	GOOS=linux GOARCH=amd64 go build -o build/linux/amd64/${APP_NAME} cmd/${APP_NAME}/main.go
	GOOS=darwin GOARCH=amd64 go build -o build/darwin/amd64/${APP_NAME} cmd/${APP_NAME}/main.go
.PHONY: build-server

build-ui:
	cd ui && yarn build
.PHONY: build-ui

publish-docker-image:
	@if test -z "$$VERSION"; then echo "Target publish-docker-image requires VERSION env var to be set"; exit 1; fi; \
	docker buildx build \
		--push \
    --platform linux/arm/v7,linux/amd64 \
    --tag ${DOCKER_REPO}${APP_NAME}:latest \
		--tag ${DOCKER_REPO}${APP_NAME}:${VERSION} \
    .
.PHONY: publish-docker-image

clean:
	rm -rf build ui/public/build
.PHONY: clean
