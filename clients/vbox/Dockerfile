#############################################
## [golang builder]  ########################
#############################################
FROM golang:1.10 as builder

ARG VERSION
ARG GITCOMMIT

COPY . /go/src/github.com/blacktop/vm-proxy/clients/vbox
WORKDIR /go/src/github.com/blacktop/vm-proxy/clients/vbox/

RUN wget --progress=bar:force -O hack/build/vars
RUN wget --progress=bar:force -O hack/build/bin

RUN hack/build/bin

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