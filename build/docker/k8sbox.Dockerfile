ARG GO_VERSION=1.19

### Build
FROM golang:${GO_VERSION} as build
COPY . /boxie
WORKDIR /boxie
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build \
    -mod vendor \
    -o /boxie/bin/boxie \
    -v /boxie/cmd/boxie

### Run
FROM alpine:3.18.0
RUN apk update
RUN apk add helm
COPY --from=build /boxie/bin/boxie /usr/local/bin/boxie

ENTRYPOINT ["boxie"]
