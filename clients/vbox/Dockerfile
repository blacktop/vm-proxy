#############################################
## [golang builder]  ########################
#############################################
FROM golang as builder

ARG VERSION
ARG GITCOMMIT

RUN go get -d github.com/blacktop/vm-proxy/...

WORKDIR /go/src/github.com/blacktop/vm-proxy/

RUN VERSION=$(cat VERSION) CLIENT=vbox hack/build/bin

#############################################
## [vbox image] ###########################
#############################################
FROM alpine:3.7

LABEL maintainer "https://github.com/blacktop"

COPY --from=builder /go/src/github.com/blacktop/vm-proxy/clients/vbox/build/vbox /bin/VBoxManage
COPY vbox.yaml /root/.vbox.yaml

ENTRYPOINT ["/bin/VBoxManage"]
CMD ["--help"]

#############################################
#############################################
#############################################