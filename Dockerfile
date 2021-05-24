FROM golang:alpine as builder
ENV GO111MODULE="on"
ENV APP_HOME /usr/src/app

RUN apk update && \
    apk --no-cache add libaio libnsl libc6-compat curl git yarn gcc musl-dev

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY . $APP_HOME
RUN go mod vendor
RUN go build -mod=vendor $APP_HOME/src/main.go

# Runtime image with v container
FROM oraclelinux:7-slim as production

ARG release=19
ARG update=3
ARG VERSION
ENV VERSION_APP=$VERSION
ENV APP_HOME /usr/src/app

# Oracle and Oracle Instant Client installation
RUN yum -y install oracle-release-el7 && \
    yum-config-manager --enable ol7_oracle_instantclient && \
    yum -y install oracle-instantclient${release}.${update}-basic \
    oracle-instantclient${release}.${update}-devel \
    oracle-instantclient${release}.${update}-sqlplus && \
    rm -rf /var/cache/yum

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder $APP_HOME/package.json /package.json
COPY --from=builder $APP_HOME/main /src/
RUN ls -alth

ENTRYPOINT ["/src/main"]
