ifndef VERSION
	VERSION="dev"
endif
PKG=$(shell grep module go.mod | sed "s/module //")
LDFLAGS=-w -X '$(PKG)/pkg/version.BuildTime=$(shell TZ='UTC' date)' -X '$(PKG)/pkg/version.GitCommit=$(shell git rev-parse --short HEAD)' -X '$(PKG)/pkg/version.Version=$(VERSION)' -X '$(PKG)/pkg/version.GoVersion=$(shell go version | awk '{print $$3}')'
COMPILER=go build -ldflags="$(LDFLAGS)"

.PHONY: build
build:
	$(COMPILER) -o bin/zdv cmd/zdv/main.go

.PHONY: clean
clean:
	rm -rf bin