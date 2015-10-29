PACKAGES = $(shell go list ./...)

all: build

deps:
	go get -u "github.com/stretchr/testify/assert"

build:
	go build

test:
	go test -cover $(PACKAGES)

