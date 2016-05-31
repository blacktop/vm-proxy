# vm-proxy
[![GoDoc][godoc]](https://godoc.org/github.com/blacktop/vm-proxy)
![License][license]  

**VBoxManage/vmrun** proxy to allow communication from within a VM to the hypervisor running the VM.

[godoc]: https://godoc.org/github.com/blacktop/vm-proxy?status.svg
[license]: https://img.shields.io/github/license/blacktop/vm-proxy.svg

### Getting Started (OSX)

Install: 
 - Docker for Mac
 - [jq](https://stedolan.github.io/jq/)  
 
 
 Now run:
```bash
$ git clone https://github.com/blacktop/vm-proxy.git
$ cd vm-proxy/server
$ go run *.go &
# To list all VirtualBox VMs
$ docker run --rm --add-host=dockerhost:$(ipconfig getifaddr en0) alpine wget -qO- dockerhost:5000/vms | jq .
```