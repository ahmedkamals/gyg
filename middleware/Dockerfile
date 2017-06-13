FROM golang:alpine

MAINTAINER Ahmed Kamal <me.ahmed.kamal+gyg@gmaial.com>

RUN apk update \
    && apk add git bash

COPY . /go/src/gyg/middleware

WORKDIR /go/src/gyg/middleware
RUN go get -d -v
RUN go install -v

ENTRYPOINT ["middleware"]
