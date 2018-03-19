# vm-proxy [WIP] :construction:

[![Circle CI](https://circleci.com/gh/blacktop/vm-proxy.png?style=shield)](https://circleci.com/gh/blacktop/vm-proxy) [![GitHub release](https://img.shields.io/github/release/blacktop/vm-proxy.svg)](https://github.com/https://github.com/blacktop/vm-proxy/releases/releases) [![License](https://img.shields.io/badge/licence-Apache%202.0-blue.svg)](LICENSE)

> **VBoxManage/vmrun** proxy to allow communication from within a VM to the hypervisor running the VM.

---

## Install _(macOS)_

```sh
$ brew install blacktop/tap/vm-proxy
```

## Client Docker Images

* blacktop/vbox
* blacktop/vmware

## Getting Started _(macOS)_

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

### Start `vm-proxy` server

```sh
$ vm-proxy

WARN[0000] no webhook token set: --token
2018/03/19 15:58:04 written cert.pem
2018/03/19 15:58:04 written key.pem
INFO[0000] vm-proxy service listening                    host=127.0.0.1 port=3993 token=
2018/03/19 15:58:43 http: TLS handshake error from 127.0.0.1:64801: EOF
```

### Start `VirtualBox` client within Docker

```sh
$ docker run --rm --add-host=dockerhost:$(ipconfig getifaddr en0) blacktop/vbox --help

Oracle VM VirtualBox Command Line Management Interface Version 5.0.20
(C) 2005-2016 Oracle Corporation
All rights reserved.

Usage:
  VBoxManage [flags]
  VBoxManage [command]

Available Commands:
  controlvm   Control VM
  debugvm     Introspection and guest debugging
  list        List all VMs
  showvminfo  Display VM info
  snapshot    Manage VirtualBox Snapshots
  startvm     Start VMs

Flags:
      --config string   config file (default is $HOME/.VBoxManage.yaml)
  -v, --version         print version number and exit

Use "VBoxManage [command] --help" for more information about a command.
```

### List VirtualBox VMs

```sh
$ docker run --rm --add-host=dockerhost:$(ipconfig getifaddr en0) blacktop/vbox list vms

"win-test_default_1456716033001_71487" {f11be617-b053-4a0f-b22c-59887290ec96}
"malice_dev" {cdb35dc9-31f6-469f-aebf-6f69830f7864}
"vagrant-golang-master_default_1458098432288_42734" {8bca67fa-03b9-45dd-9436-53f1877e1608}
"go-malice-test_default_1458098825435_9154" {208244e8-b320-41a8-b037-7127cbc9d09d}
"default" {6e94d53e-5f78-4366-9aa8-a5725ac6dbfb}
```

### API

#### List VirtualBox VMs _(via API)_

```sh
 $ http --verify=no https://127.0.0.1:3993/vbox/list
```

```http
HTTP/1.1 500 Internal Server Error
Content-Length: 85
Content-Type: application/json; charset=UTF-8
Date: Mon, 19 Mar 2018 22:03:43 GMT

VBoxManage not found. Make sure VirtualBox is installed and VBoxManage is in the path
```

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
