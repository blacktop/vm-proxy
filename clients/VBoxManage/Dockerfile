FROM gliderlabs/alpine:3.3

MAINTAINER blacktop, https://github.com/blacktop

COPY . /go/src/github.com/blacktop/vm-proxy/clients/VBoxManage
COPY .VBoxManage.yaml ~
RUN apk-install -t build-deps go git mercurial \
  && set -x \
  && echo "Building info Go binary..." \
  && cd /go/src/github.com/blacktop/vm-proxy/clients/VBoxManage \
  && export GOPATH=/go \
  && go version \
  && go get \
  && go build -ldflags "-X main.Version=$(cat VERSION) -X main.BuildTime=$(date -u +%Y%m%d)" -o /usr/bin/VBoxManage \
  && rm -rf /go \
  && apk del --purge build-deps

ENV PATH /usr/bin:$PATH

ENTRYPOINT ["VBoxManage"]

CMD ["--help"]