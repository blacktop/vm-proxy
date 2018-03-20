# vm-proxy [WIP] :construction:

[![Circle CI](https://circleci.com/gh/blacktop/vm-proxy.png?style=shield)](https://circleci.com/gh/blacktop/vm-proxy) [![GitHub release](https://img.shields.io/github/release/blacktop/vm-proxy.svg)](https://github.com/https://github.com/blacktop/vm-proxy/releases/releases) [![License](https://img.shields.io/badge/licence-Apache%202.0-blue.svg)](LICENSE)

> **VBoxManage/vmrun** proxy to allow communication from within a VM to the hypervisor running the VM.

---

## Client Docker Images

* blacktop/vbox
* blacktop/vmware

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

## Issues

Find a bug? Want more features? Find something missing in the documentation? Let me know! Please don't hesitate to [file an issue](https://github.com/blacktop/vm-proxy/issues/new)

## License

Apache License (Version 2.0) Copyright (c) 2016 - 2018 **blacktop**
