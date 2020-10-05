BUILD_IMAGE ?= golang:1.15.2-buster
DOCKER_ARGS ?= --rm -v $(PWD):/build -w /build
BASENAME ?= dotrr
VERSION ?= SNAPSHOT
COMPILER_ARGS ?= -ldflags="-X github.com/btisdall/dotrr/v2/app/cmd.appVersion=$(VERSION)" \
	-o build/$(BASENAME)-$(GOOS)-$(GOARCH) main.go
GOARCH = amd64

.PHONY: format
format:
	goformat -w .

.PHONY: test
test:
	docker run $(DOCKER_ARGS) $(BUILD_IMAGE) go test ./...

.PHONY: test_local
test_local:
	go test ./...

.PHONY: build_darwin_local
build_darwin_local: GOOS = darwin
build_darwin_local:
	go build $(COMPILER_ARGS)

.PHONY: build_darwin
build_darwin: GOOS = darwin
build_darwin:
	docker run \
	-e GOOS=$(GOOS) \
	-e GOARCH=$(GOARCH) \
	$(DOCKER_ARGS) $(BUILD_IMAGE) go build $(COMPILER_ARGS)

.PHONY: build_windows
build_windows: GOOS = windows
build_windows:
	docker run \
	-e GOOS=$(GOOS) \
	-e GOARCH=$(GOARCH) \
	$(DOCKER_ARGS) $(BUILD_IMAGE) go build $(COMPILER_ARGS)

.PHONY: build_linux
build_linux: GOOS = linux
build_linux:
	docker run \
	-e GOOS=$(GOOS) \
	-e GOARCH=$(GOARCH) \
	$(DOCKER_ARGS) $(BUILD_IMAGE) go build $(COMPILER_ARGS)

.PHONY: build_all
build_all: build_darwin build_windows build_linux