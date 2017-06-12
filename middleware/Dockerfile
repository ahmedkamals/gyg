FROM alpine:latest
MAINTAINER Ahmed Kamal <me.ahmed.kamal+gyg@gmaial.com>

RUN apk add --no-cache ca-certificates

ENV GOLANG_VERSION 1.7.5
ENV GOLANG_SRC_URL https://golang.org/dl/go$GOLANG_VERSION.src.tar.gz
ENV GOLANG_SRC_SHA256 4e834513a2079f8cbbd357502cccaac9507fd00a1efe672375798858ff291815

RUN apk update \
    && apk add bash

# https://golang.org/issue/14851
COPY snippets/no-pic.patch /

RUN set -ex \
	&& apk add --no-cache --virtual .build-deps \
		gcc \
		musl-dev \
		openssl \
		go \
	\
	&& export GOROOT_BOOTSTRAP="$(go env GOROOT)" \
	\
	&& wget -q "$GOLANG_SRC_URL" -O golang.tar.gz \
	&& echo "$GOLANG_SRC_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz \
	&& cd /usr/local/go/src \
	&& patch -p2 -i /no-pic.patch \
	&& ./make.bash \
	\
	&& rm -rf /*.patch \
	&& apk del .build-deps

# Environment
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV WORK_DIR $GOPATH/src/gyg/middleware

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

COPY snippets/go-wrapper /usr/local/bin/

ADD . $WORKDIR

WORKDIR $WORK_DIR

RUN go version \
    && set -ex \
    	&& apk add --no-cache --virtual .build-deps \
    		git \
    		make \
    	\
    	&& make default \
    	&& rm -rf .git \
    	&& apk del .build-deps

RUN ls $WORK_DIR/bin
