APP_NAME := filer

# install:
# 	cd ui && yarn install
# .PHONY: install

start-server:
	FILER_INBOX_PATH=./example/inbox \
	FILER_DATA_PATH=./example/target \
	FILER_RULES_PATH=./example/rules.yml \
	go run cmd/filer/main.go
.PHONY: start-server

# start-ui:
# 	cd ui && ./node_modules/.bin/webpack serve --config webpack.dev.js --host=0.0.0.0 --port=3000 --open='Google Chrome'
# .PHONY: start-ui

# build: build-server build-ui

build-server:
	GOOS=linux GOARCH=arm GOARM=7 go build -o build/linux/arm/v7/${APP_NAME} cmd/${APP_NAME}/main.go
	GOOS=linux GOARCH=amd64 go build -o build/linux/amd64/${APP_NAME} cmd/${APP_NAME}/main.go
	GOOS=darwin GOARCH=amd64 go build -o build/darwin/amd64/${APP_NAME} cmd/${APP_NAME}/main.go
.PHONY: build-server

# build-ui:
# 	cd ui && NODE_ENV=production ./node_modules/.bin/webpack --config webpack.prod.js
# .PHONY: build-ui

# publish-docker-image: build
# 	@if test -z "$$VERSION"; then echo "Target publish-docker-image requires VERSION env var to be set"; exit 1; fi; \
# 	docker buildx build \
# 		--push \
#     --platform linux/arm/v7,linux/amd64 \
#     --tag docker.olymp.alabor.me/swissmanu/${APP_NAME}:latest \
# 		--tag docker.olymp.alabor.me/swissmanu/${APP_NAME}:${VERSION} \
#     .
# .PHONY: publish-docker-image

clean:
	rm -rf build ui/build
.PHONY: clean