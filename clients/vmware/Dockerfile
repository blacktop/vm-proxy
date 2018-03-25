#############################################
## [golang builder]  ########################
#############################################
FROM golang as builder

ARG VERSION
ARG GITCOMMIT

RUN go get -d github.com/blacktop/vm-proxy/...

WORKDIR /go/src/github.com/blacktop/vm-proxy/

RUN VERSION=$(cat VERSION) CLIENT=vmware hack/build/bin

#############################################
## [vmware image] ###########################
#############################################
FROM alpine:3.7

LABEL maintainer "https://github.com/blacktop"

COPY --from=builder /go/src/github.com/blacktop/vm-proxy/clients/vmware/build/vmware /bin/vmrun
COPY vmware.yaml /root/.vmware.yaml

ENTRYPOINT ["/bin/vmrun"]
CMD ["--help"]

#############################################
#############################################
#############################################