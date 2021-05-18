FROM golang:alpine AS base
ENV GO111MODULE="on"
RUN apk add --no-cache ca-certificates git

FROM base AS development
ENV APP_HOME /usr/src/app
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY . $APP_HOME
EXPOSE 3000
CMD ["make", "start"]

# Go image for building the project
FROM development as builder
ENV APP_HOME /usr/src/app
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' $APP_HOME/src/main.go

# Runtime image with v container
FROM scratch as production
ARG VERSION
ENV VERSION_APP=$VERSION
ENV APP_HOME /usr/src/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder $APP_HOME/package.json /package.json
COPY --from=builder $APP_HOME/main /src/

ENTRYPOINT ["/src/main"]
