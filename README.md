# vm-proxy
![License][license]  

**VBoxManage/vmrun** proxy to allow communication from within a VM to the hypervisor running the VM.

[godoc]: https://godoc.org/github.com/blacktop/vm-proxy?status.svg
[license]: https://img.shields.io/github/license/blacktop/vm-proxy.svg

### Getting Started (OSX)

#### Install:
 - [Docker for Mac](https://beta.docker.com/)
 - [homebrew](http://brew.sh/)  

#### Now run:
Start `vm-proxy-server`
```bash
$ brew install https://raw.githubusercontent.com/blacktop/vm-proxy/master/homebrew/Formula/vm-proxy-server.rb
$ vm-proxy-server
```
To list all VirtualBox VMs
```bash
$ docker run --rm --add-host=dockerhost:$(ipconfig getifaddr en0) blacktop/vbox list vms
```
```bash
"win-test_default_1456716033001_71487" {f11be617-b053-4a0f-b22c-59887290ec96}
"malice_dev" {cdb35dc9-31f6-469f-aebf-6f69830f7864}
"vagrant-golang-master_default_1458098432288_42734" {8bca67fa-03b9-45dd-9436-53f1877e1608}
"go-malice-test_default_1458098825435_9154" {208244e8-b320-41a8-b037-7127cbc9d09d}
"default" {6e94d53e-5f78-4366-9aa8-a5725ac6dbfb}
```

### Downloads
I will be releasing binaries of **VBoxManage** and **vmrun** soon.

### Docker Images
VBoxManage
```bash
$ docker pull blacktop/vbox
```

### ToDo
 - [ ] Add version check to debugvm calls
 - [ ] vmrun
 - [x] create homebrew installer for vm-proxy-server
 - [x] build small base images with VBoxManage in them
 - [ ] figure out filesystem translation for dropping PCAP or memory dumps so container can see them (using volumes?)  
