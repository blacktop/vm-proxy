# blacktop/vbox

> VirtualBox Client for `vm-proxy`

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

## Issues

Find a bug? Want more features? Find something missing in the documentation? Let me know! Please don't hesitate to [file an issue](https://github.com/blacktop/vm-proxy/issues/new)

## License

Apache License (Version 2.0) Copyright (c) 2016 - 2018 **blacktop**
