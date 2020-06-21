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

run-redshift:
	@./${APP} \
		-bucket-sql com.dkisler.meta \
		-path-sql sql/test.sql \
		-sql-parameters "{\"col1\": \"current_date\", 
		\"bucket_path\": \"net.dkisler.data.test/sql-runner-test/unload\"}"

run-local:
	@./${APP} \
		-bucket-sql com.dkisler.meta \
		-path-sql sql/test-local-pg.sql \
		-sql-parameters '{"col1": "current_date"}'

build-docker:
	@docker build -t ${DIR} .

db-start:
	@docker run -d \
		-p 15439:5432 \
		-e POSTGRES_DB=test \
		-e POSTGRES_USER=admin \
		-e POSTGRES_PASSWORD=qwe123asD \
		-v ${PWD}/data:/data \
		--name db-test \
		-t postgres:9.6-alpine
	
db-up:
	@docker start db-test

db-down:
	@docker stop db-test

db-rm:
	@docker rm -f db-test
