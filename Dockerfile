FROM golang:1.16.5-alpine3.13 as base
ENV GO111MODULE="on"
RUN apk update && \
    apk --no-cache add libaio libnsl libc6-compat curl git yarn gcc musl-dev

FROM base as develoment
ENV APP_HOME /usr/src/app
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY . $APP_HOME
EXPOSE 3000

# Runtime image with v container
FROM oraclelinux:7-slim as production

RUN yum -y install oracle-golang-release-el7 oracle-release-el7 && \
    yum -y install install git gcc golang && \
    yum -y install oracle-instantclient19.3-basiclite && \
    rm -rf /var/cache/yum

RUN echo /usr/lib/oracle/19.3/client64/lib > /etc/ld.so.conf.d/oracle-instantclient.conf

ENV APP_HOME /usr/src/app
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY . $APP_HOME

RUN CGO_ENABLED=1 GOOS=linux go build -mod=vendor $APP_HOME/src/main.go

ENTRYPOINT ["./main"]
