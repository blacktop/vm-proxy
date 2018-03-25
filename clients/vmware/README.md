![vmware-logo](https://github.com/blacktop/vm-proxy/raw/master/docs/vmware-logo.png)

# blacktop/vmware

> VMware Client for `vm-proxy`

---

## Getting Started

```sh
$ docker run --rm blacktop/vmware --help
```

```

```

### Start VM

```sh
$ docker run --rm \
             --add-host=dockerhost:$(ipconfig getifaddr en0) \
             -v $HOME/.vmproxy:/root/.vmproxy \
             blacktop/vmware start "/Users/blacktop/Documents/Virtual Machines.localized/Ubuntu 64-bit 16.04.vmwarevm/Ubuntu 64-bit 16.04.vmx"
```

### List `VirtualBox` VMs

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
