FROM golang:1.16 as builder

LABEL MAINTAINER="laurence"

WORKDIR /dubbogo

COPY ./server ./

EXPOSE 20001

ENTRYPOINT ["./server"]
