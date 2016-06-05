# vm-proxy
![License][license]  

**VBoxManage/vmrun** proxy to allow communication from within a VM to the hypervisor running the VM.

[godoc]: https://godoc.org/github.com/blacktop/vm-proxy?status.svg
[license]: https://img.shields.io/github/license/blacktop/vm-proxy.svg

### Getting Started (OSX)

#### Install:
 - [Docker for Mac](https://beta.docker.com/)
 - [jq](https://stedolan.github.io/jq/)  

#### Now run:
```bash
$ go get -v github.com/docker/machine
$ go get -v github.com/blacktop/vm-proxy
$ cd $GOPATH/src/github.com/blacktop/vm-proxy/server
$ go run *.go &
# To list all VirtualBox VMs
$ docker run --rm --add-host=dockerhost:$(ipconfig getifaddr en0) blacktop/vbox list vms
```

### Downloads
I will be releasing binaries of **VBoxManage** and **vmrun** soon.

### ToDo
 - [ ] Add version check to debugvm calls
 - [ ] vmrun
 - [ ] create homebrew installer for vm-proxy-server
 - [ ] build small base images with VBoxManage in them
 - [ ] figure out filesystem translation for dropping PCAP or memory dumps so comtainer can see them (using volumes?)