FROM alpine:3.4
MAINTAINER Daisuke Fujita <dtanshi45@gmail.com> (@dtan4)

ENV GOPATH /go
COPY . /go/src/github.com/dtan4/apig-sample

RUN apk add --no-cache --update --virtual=build-deps g++ git go mercurial \
    && cd /go/src/github.com/dtan4/apig-sample \
    && go get -v ./... \
    && go build -ldflags="-s -w" \
    && cp apig-sample /apig-sample \
    && cd / \
    && rm -rf /go \
    && apk del build-deps

EXPOSE 8080

CMD ["/apig-sample"]
