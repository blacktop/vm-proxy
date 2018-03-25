![vmware-logo](https://github.com/blacktop/vm-proxy/raw/master/docs/vmware-logo.png)

# blacktop/vmware

[![CircleCI](https://circleci.com/gh/blacktop/vm-proxy.png?style=shield)](https://circleci.com/gh/blacktop/vm-proxy) [![License](https://img.shields.io/badge/licence-Apache%202.0-blue.svg)](LICENSE) [![Docker Stars](https://img.shields.io/docker/stars/blacktop/vmware.svg)](https://hub.docker.com/r/blacktop/vmware/) [![Docker Pulls](https://img.shields.io/docker/pulls/blacktop/vmware.svg)](https://hub.docker.com/r/blacktop/vmware/) [![Docker Image](https://img.shields.io/badge/docker%20image-11MB-blue.svg)](https://hub.docker.com/r/blacktop/vmware/)

> VMware Client for `vm-proxy`

---

## Getting Started

```sh
$ docker run --rm blacktop/vmware --help

vmrun version 1.17.0 build-7520154

Usage:
  vmrun [command]

Available Commands:
  help        Help about any command
  list        List all running VMs
  start       Start a VM
  stop        Stop a VM

Flags:
      --config string   config file (default is $HOME/.vmware.yaml)
  -h, --help            help for vmrun
  -t, --toggle          Help message for toggle

Use "vmrun [command] --help" for more information about a command.
```

### Start VM

```sh
$ docker run --rm \
             --add-host=dockerhost:$(ipconfig getifaddr en0) \
             -v $HOME/.vmproxy:/root/.vmproxy \
             blacktop/vmware start "/Users/blacktop/Documents/Virtual Machines.localized/Ubuntu 64-bit 16.04.vmwarevm/Ubuntu 64-bit 16.04.vmx"
```

### List VMs

```sh
$ docker run --rm \
             --add-host=dockerhost:$(ipconfig getifaddr en0) \
             -v $HOME/.vmproxy:/root/.vmproxy \
             blacktop/vmware list
```

```sh
Total running VMs: 1
/Users/blacktop/Documents/Virtual Machines.localized/Ubuntu 64-bit 16.04.vmwarevm/Ubuntu 64-bit 16.04.vmx
```

### Stop VM

```sh
$ docker run --rm \
             --add-host=dockerhost:$(ipconfig getifaddr en0) \
             -v $HOME/.vmproxy:/root/.vmproxy \
             blacktop/vmware stop "/Users/blacktop/Documents/Virtual Machines.localized/Ubuntu 64-bit 16.04.vmwarevm/Ubuntu 64-bit 16.04.vmx"
```

## Issues

Find a bug? Want more features? Find something missing in the documentation? Let me know! Please don't hesitate to [file an issue](https://github.com/blacktop/vm-proxy/issues/new)

## License

Apache License (Version 2.0) Copyright (c) 2016 - 2018 **blacktop**
