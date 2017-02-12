OUT := goSoundbench
PKG := github.com/dh1tw/goSoundbench
COMMITID := $(shell git describe --always --long --dirty)
COMMIT := $(shell git rev-parse --short HEAD)
VERSION := $(shell git describe --tags)

PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)
all: build

build:
	go build -v -o ${OUT} -ldflags="-X github.com/dh1tw/goSoundbench/cmd.commitHash=${COMMIT} \
		-X github.com/dh1tw/goSoundbench/cmd.version=${VERSION}"

# strip off dwraf table - used for travis CI
dist: 
	go build -v -o ${OUT} -ldflags="-w -X github.com/dh1tw/goSoundbench/cmd.commitHash=${COMMIT} \
		-X github.com/dh1tw/goSoundbench/cmd.version=${VERSION}"

# test:
# 	@go test -short ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

install: 
	go install -v -ldflags="-w -X github.com/dh1tw/goSoundbench/cmd.commitHash=${COMMIT} \
		-X github.com/dh1tw/goSoundbench/cmd.version=${VERSION}"

# static: vet lint
# 	go build -i -v -o ${OUT}-v${VERSION} -tags netgo -ldflags="-extldflags \"-static\" -w -s -X main.version=${VERSION}" ${PKG}

run: build
	./${OUT} server mqtt

clean:
	-@rm ${OUT} ${OUT}-v*

.PHONY: build run install vet lint clean