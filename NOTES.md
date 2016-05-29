# Notes

### VBoxManage Commands to Cover

#### Start
```bash
VBoxManage snapshot <label> restore <snapshot>
VBoxManage snapshot <label> restorecurrent
```

#### Stop (with timeout)
```bash
VBoxManage controlvm <label> poweroff
```

#### List
```bash
VBoxManage list vms
```

#### Status
```bash
VBoxManage showvminfo <label> --machinereadable
```

#### Dump Memory
```bash
VBoxManage debugvm <label> dumpvmcore --filename <path>
VBoxManage debugvm <label> dumpguestcore --filename <path>
```

#### Dump PCAP
```bash
VBoxManage controlvm <label> nictracefile1 <pcap_path>
VBoxManage controlvm <label> nictrace1 on
```

#### Version
```bash
VBoxManage -v
```

### vboxwebsrv

## Install on OSX
```bash
$ edit $HOME/Library/LaunchAgents/org.virtualbox.vboxwebsrv.plist
```
Change
```xml
  <key>Disabled</key>
  <true/>
```
 to
```xml
  <key>Disabled</key>
  <false/>
```
Change **vboxwebsrv** auth to **null**
```bash
$ VBoxManage setproperty websrvauthlibrary null
```
Manually start **vboxwebsrv**
```bash
launchctl load ~/Library/LaunchAgents/org.virtualbox.vboxwebsrv.plist
```

```bash
git clone https://github.com/blacktop/vboxmanage-proxy.git
cd vboxmanage-proxy
GOOS=linux go build -o VBoxManage
```
