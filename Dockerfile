# syntax = docker/dockerfile:1

########################################
## Build Stage
########################################
FROM golang:1.22-bookworm as builder

# add a label to clean up later
LABEL stage=intermediate

# setup the working directory
WORKDIR /go/src

# install dependencies
ENV GO111MODULE=on
COPY ./go.sum ./go.sum
COPY ./go.mod ./go.mod
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download -x

# add source code
COPY . .

# build the source
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o robin-api-linux-amd64

########################################
## Production Stage
########################################
FROM ubuntu:24.04

# set working directory
WORKDIR /root

# copy required files from builder
COPY --from=builder /go/src/robin-api-linux-amd64 ./robin-api-linux-amd64

# add required files from host
COPY ./configs/ ./configs/

ENTRYPOINT ["./robin-api-linux-amd64"]
