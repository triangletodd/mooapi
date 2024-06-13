#vars
IMAGENAME=mooapi

.PHONY: fmt test build run all

.DEFAULT_GOAL := all

fmt:
	@ gofmt -w -s .

test:
	@go test -v

build:
	@docker build --pull -t ${IMAGENAME} .

run:
	@docker run --rm  -it -p 8080:8080 ${IMAGENAME}

all: test build
