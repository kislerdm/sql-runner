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

build-docker:
	@docker build -t ${DIR} .

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

run-docker:
	@docker run \
		-e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
		-e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
		-e AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} \
		-e DB_HOST=${DB_HOST} \
		-e DB_PORT=${DB_PORT} \
		-e DB_NAME=${DB_NAME} \
		-e DB_USER=${DB_USER} \
		-e DB_PASSWORD=${DB_PASSWORD} \
		--network="host" \
		-t ${DIR} \
			-bucket-sql com.dkisler.meta \
			-path-sql sql/test-local-pg.sql \
			-sql-parameters '{"col1": "current_date"}'

db-start:
	@docker run -d \
		-p ${DB_PORT}:5432 \
		-e POSTGRES_DB=${DB_NAME} \
		-e POSTGRES_USER=${DB_USER} \
		-e POSTGRES_PASSWORD=${DB_PASSWORD} \
		-v ${PWD}/data:/data \
		-t postgres:9.6-alpine
	
db-up:
	@docker start db-test

db-down:
	@docker stop db-test

db-rm:
	@docker rm -f db-test
