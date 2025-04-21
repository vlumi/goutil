.DEFAULT_GOAL := all

.PHONY:fmt vet build all

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	go build -o bin/ ./...

clean:
	go clean ./...
	rm -Rf bin/

all: fmt vet build
