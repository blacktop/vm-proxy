#############################################
## [golang builder]  ########################
#############################################
FROM golang:1.10 as builder

ARG VERSION
ARG GITCOMMIT

COPY . /go/src/github.com/blacktop/vm-proxy/clients/vmware

RUN hack/build/binary

#############################################
## [vmware image] ###########################
#############################################
FROM alpine:3.7

LABEL maintainer "https://github.com/blacktop"

COPY --from=builder /go/src/github.com/blacktop/vm-proxy/clients/vmware/build/vmware /bin/vmrun
COPY .vmware.yaml /root

ENTRYPOINT ["/bin/VBoxManage"]
CMD ["--help"]

#############################################
#############################################
#############################################