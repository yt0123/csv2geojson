.PHONY: build fmt setup

BUILDER_IMAGE := golang-build-env
BUILDER_WORK_DIR := /go/src/github.com/ty-edelweiss/csv2geojson
BUILDER_CMD := docker run --rm -e GOPATH=/go -v `pwd`:$(BUILDER_WORK_DIR) -w $(BUILDER_WORK_DIR) $(BUILDER_IMAGE)

build:
	$(BUILDER_CMD) make clean build

dep:
	$(BUILDER_CMD) make dep

fmt:
	$(BUILDER_CMD) make fmt

setup:
	docker build -f Dockerfile.builder -t $(BUILDER_IMAGE) .
