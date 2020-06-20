SHELL=/bin/bash

build-compress: build compress
exec: build run

.PHONY: build-compress build-compress-local exec

OS := darwin
ARCH := 
FLAGS := `-a -gcflags=all="-l -B -C" -ldflags="-w -s"`
DIR := sql-runner-redshift
APP := sql-runner-redshift.run

args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

build: ${DIR}/*.go
	@GOOS=${OS} GOARCH=${ARCH} go build -o ${APP} ${DIR}/*.go

compress: *.run
	@upx -9 ${APP}

run:
	@./${APP}

build-docker:
	@docker build -t ${DIR} .
