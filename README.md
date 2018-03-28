# vm-proxy [WIP] :construction:

[![Circle CI](https://circleci.com/gh/blacktop/vm-proxy.png?style=shield)](https://circleci.com/gh/blacktop/vm-proxy) [![GitHub release](https://img.shields.io/github/release/blacktop/vm-proxy.svg)](https://github.com/https://github.com/blacktop/vm-proxy/releases/releases) [![License](https://img.shields.io/badge/licence-Apache%202.0-blue.svg)](LICENSE)

> **VBoxManage/vmrun** proxy to allow communication from within a VM to the hypervisor running the VM.

---

## Why?

This allows you to communicate with hypervisors from within docker containers.

The main use case I am working towards for my _MVP_ is to support the local hypervisor machinery that the [cuckoo sandbox](https://github.com/cuckoosandbox/cuckoo/tree/master/cuckoo/machinery) uses so that my project [docker-cuckoo](https://github.com/blacktop/docker-cuckoo) can work with **VMware/VirtualBox/KVM** etc.

## How?

`vm-proxy` works by creating a secure local webhook to _proxy_ `VBoxManage` or `vmrun` out the the host running docker. So from the container's perspective it is using the real tools locally, but they are instead using a small golang binary that securely communicates to `vm-proxy`.

`vm-proxy` also creates SSL certs and a token to secure communications between the container and the hypervisor. Also I will only expose a minimal set of hypervisor functionality at first to prevent malicious actors from trying to harm your host or VMs. I will also sanitize input sent via the clients to the server.

Others have created solutions where containers can `ssh` to the host and run **ANY** commands, which I believe is not safe (think `rm -rf /`). Or you can leverage APIs exposed by the hypervisors, but then you have to maintain your middleware to talk to them. You also will need to setup and start the API servers locally.

My solution (targeting cuckoo) requires **NO** changes to cuckoo as it thinks it is talking to the real `VBoxManage`/`vmrun` binaries, making it easier to maintain in the long term and requiring no changes on cuckoo's side.

## Client Docker Images

* [blacktop/vbox](https://github.com/blacktop/vm-proxy/blob/master/clients/vbox/README.md)
* [blacktop/vmware](https://github.com/blacktop/vm-proxy/blob/master/clients/vmware/README.md)

## Getting Started _(macOS)_

### Install

```sh
$ brew install blacktop/tap/vm-proxy
```

### Start `vm-proxy` brew background service

```sh
$ brew services start blacktop/tap/vm-proxy
```

### Manually run `vm-proxy` server

```sh
$ vm-proxy --help
```

```sh
Usage: vm-proxy [OPTIONS] COMMAND [arg...]

VMProxy Server - allows hypervisors to be controlled from docker containers

Version: , BuildTime:

Author:
  blacktop - <https://github.com/blacktop>

Options:
  --verbose, -V  verbose output
  --host value   microservice host (default: "127.0.0.1") [$VMPROXY_HOST]
  --port value   microservice port (default: "3993") [$VMPROXY_PORT]
  --token value  webhook token [$VMPROXY_TOKEN]
  --help, -h     show help
  --version, -v  print the version

Commands:
  update  Update images
  export  Export Database
  help    Shows a list of commands or help for one command

Run 'vm-proxy COMMAND --help' for more information on a command.
```

### Manually start `vm-proxy` server

```sh
$ vm-proxy

WARN[0000] no webhook token set: --token
2018/03/19 15:58:04 written cert.pem
2018/03/19 15:58:04 written key.pem
INFO[0000] vm-proxy service listening                    host=127.0.0.1 port=3993 token=
```

## Use a Hypervisor Client

### Start `VirtualBox` client within Docker

> See docs [here](https://github.com/blacktop/vm-proxy/blob/master/clients/vbox/README.md)

### Start `VMware` client within Docker

> See docs [here](https://github.com/blacktop/vm-proxy/blob/master/clients/vmware/README.md)

## API

> See docs [here](https://github.com/blacktop/vm-proxy/blob/master/docs/api.md)

## TODO

* [ ] Add version check to debugvm calls
* [ ] vmrun
* [x] create homebrew installer for vm-proxy-server
* [x] build small base images with VBoxManage in them
* [ ] figure out filesystem translation for dropping PCAP or memory dumps so container can see them (using volumes?)
* [x] auto-create certs on first run
* [ ] standardize on a log provider (apex/logrus)

## Issues

Find a bug? Want more features? Find something missing in the documentation? Let me know! Please don't hesitate to [file an issue](https://github.com/blacktop/vm-proxy/issues/new)

## License

Apache License (Version 2.0) Copyright (c) 2016 - 2018 **blacktop**
