ARG GO_VERSION=1.19

### Build
FROM golang:${GO_VERSION} as build
COPY . /k8sbox
WORKDIR /k8sbox
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build \
    -mod vendor \
    -o /k8sbox/bin/k8sbox \
    -v /k8sbox/cmd/k8sbox

### Run
FROM alpine:3.18.0
RUN apk update
RUN apk add helm
COPY --from=build /k8sbox/bin/k8sbox /usr/local/bin/k8sbox

ENTRYPOINT ["k8sbox"]
