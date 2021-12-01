## Docker image for Golang Core Data micro service
#ARG BUILDER_BASE=golang:1.16-alpine3.14
#FROM ${BUILDER_BASE} AS builder
#
#WORKDIR /pixiu
#
#COPY . .
#
#RUN go build -o ./cmd/pixiu/pixiu ./cmd/pixiu/

#Next image - Copy built Go binary into new workspace
FROM alpine:3.14 AS result

COPY ./pixiu /
COPY ./samples/dubbogo/simple/triple/pixiu/conf.yaml /

ENTRYPOINT ["/pixiu","gateway", "start", "-c", "/conf.yaml"]


