![vbox-logo](https://github.com/blacktop/vm-proxy/raw/master/docs/vbox-logo.png)

# blacktop/vbox

[![CircleCI](https://circleci.com/gh/blacktop/vm-proxy.png?style=shield)](https://circleci.com/gh/blacktop/vm-proxy) [![License](https://img.shields.io/badge/licence-Apache%202.0-blue.svg)](https://github.com/blacktop/vm-proxy/blob/master/LICENSE) [![Docker Stars](https://img.shields.io/docker/stars/blacktop/vbox.svg)](https://hub.docker.com/r/blacktop/vbox/) [![Docker Pulls](https://img.shields.io/docker/pulls/blacktop/vbox.svg)](https://hub.docker.com/r/blacktop/vbox/) [![Docker Image](https://img.shields.io/badge/docker%20image-11MB-blue.svg)](https://hub.docker.com/r/blacktop/vbox/)

> VirtualBox Client for [vm-proxy](https://github.com/blacktop/vm-proxy)

---

## Getting Started

```sh
$ docker run --rm \
             --add-host=dockerhost:$(ipconfig getifaddr en0) \
             -v $HOME/.vmproxy:/root/.vmproxy \
             blacktop/vbox --help

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

### List `VirtualBox` VMs

```sh
$ docker run --rm \
             --add-host=dockerhost:$(ipconfig getifaddr en0) \
             -v $HOME/.vmproxy:/root/.vmproxy \
             blacktop/vbox list vms

"win-test_default_1456716033001_71487" {f11be617-b053-4a0f-b22c-59887290ec96}
"malice_dev" {cdb35dc9-31f6-469f-aebf-6f69830f7864}
"vagrant-golang-master_default_1458098432288_42734" {8bca67fa-03b9-45dd-9436-53f1877e1608}
"go-malice-test_default_1458098825435_9154" {208244e8-b320-41a8-b037-7127cbc9d09d}
"default" {6e94d53e-5f78-4366-9aa8-a5725ac6dbfb}
```

## Issues

Find a bug? Want more features? Find something missing in the documentation? Let me know! Please don't hesitate to [file an issue](https://github.com/blacktop/vm-proxy/issues/new)

## License

Apache License (Version 2.0) Copyright (c) 2016 - 2018 **blacktop**
