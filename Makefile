GOVERSION:=$(shell go version)
GOOS:=$(word 1,$(subst /, ,$(lastword $(GOVERSION))))
GOARCH:=$(word 2,$(subst /, ,$(lastword $(GOVERSION))))
APPNAME:=short-url-go

.PHONY:build
build:
	@go build -ldflags="-s -X 'github.com/verniyyy/short-url-go/cmd.version=$(shell git describe --tags --abbrev=0)'" -o build/$(GOOS)-$(GOARCH)/bin/$(APPNAME) -v

.PHONY:serve
serve: build
	@build/$(GOOS)-$(GOARCH)/bin/$(APPNAME)