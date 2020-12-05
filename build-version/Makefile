VERSION=$(shell git describe --tags --always)
BRANCH=$(shell git symbolic-ref -q --short HEAD)
REVISION=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date +%FT%T%Z)

VERSION_PKG=github.com/tonybase/golang-best-practices/build-version/version

build:
	go build -ldflags "-X $(VERSION_PKG).Version=$(VERSION) -X $(VERSION_PKG).Branch=$(BRANCH) -X $(VERSION_PKG).Revision=$(REVISION) -X $(VERSION_PKG).BuildDate=$(BUILD_DATE)"

all: build
