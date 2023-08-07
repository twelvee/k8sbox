ARG GO_VERSION=1.19

### Build go app
FROM golang:${GO_VERSION} as build
COPY . /boxie
WORKDIR /boxie
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build \
    -mod vendor \
    -o /boxie/bin/boxie \
    -v /boxie/cmd/boxie

### Build web app
FROM node:18-alpine as build-web
COPY ./web/app /boxie
WORKDIR /boxie
RUN npm install
RUN npm run build

### Run
FROM debian:stable-slim
COPY --from=build /boxie/bin/boxie /usr/local/bin/boxie
COPY --from=build-web /boxie/dist /boxie
