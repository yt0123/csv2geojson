.PHONY: build clean

GOOS := darwin
GOARCH := amd64

CMD := cmd/csv2geojson

build: clean $(CMD)

clean:
	-rm $(CMD)

dep:
	dep ensure

fmt:
	go fmt ./...

% : %.go
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $@ -ldflags "-X main.version=$(shell git rev-parse HEAD)" $<