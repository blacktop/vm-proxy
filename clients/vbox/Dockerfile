#############################################
## [golang builder]  ########################
#############################################
FROM golang:1.10 as builder

ARG VERSION
ARG GITCOMMIT

COPY . /go/src/github.com/blacktop/vm-proxy/clients/vbox

RUN hack/build/binary

#############################################
## [vbox image] ###########################
#############################################
FROM alpine:3.7

LABEL maintainer "https://github.com/blacktop"

COPY --from=builder /go/src/github.com/blacktop/vm-proxy/clients/vbox/build/vbox /bin/VBoxManage
COPY .vbox.yaml /root

ENTRYPOINT ["/bin/VBoxManage"]
CMD ["--help"]

#############################################
#############################################
#############################################