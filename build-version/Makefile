VERSION=$(shell git describe --tags --always)
BRANCH=$(shell git symbolic-ref -q --short HEAD)
REVISION=$(shell git )

build:
	go build -ldflags "-X main.Version=$(VERSION) -X main.Branch=$(BRANCH)"

all: build
