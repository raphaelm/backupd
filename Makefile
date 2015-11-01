.PHONY: all test build backupd backupexporter vet check

PACKAGES = $(shell go list ./...)

all: build

deps:
	go get -u "gopkg.in/yaml.v2"
	go get -u "github.com/stretchr/testify/assert"
	go get -u "github.com/gorilla/mux"

backupd:
	cd backupd && go build 

backupexporter:
	cd backupexporter && go build 

build: backupd backupexporter

test:
	go test -cover $(PACKAGES)

vet:
	go vet ./...

check: test vet
