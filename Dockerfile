FROM golang:rc-alpine3.12 AS build

ARG DIR=sql-runner-redshift

COPY . /go/src/

WORKDIR /go/src/${DIR}

RUN apk update -q \
    && apk add --no-cache -q \
        g++ \
        upx \
    && go test ./... \
    && go build -a -gcflags=all="-l -B -C" -ldflags="-w -s" -o /root/runner *.go \
    && upx -9 /root/runner

FROM alpine:3.12.0 AS run

ENV AWS_ACCESS_KEY_ID ""
ENV AWS_SECRET_ACCESS_KEY ""
ENV AWS_DEFAULT_REGION "eu-west-1"

ENV DB_HOST "localhost"
ENV DB_PORT 5432
ENV DB_NAME "postgres"
ENV DB_USER "postgres"
ENV DB_PASSWORD "postgres"

WORKDIR /root

COPY --from=build /root/runner .

ENTRYPOINT ["./runner"]
