# Notes

### VBoxManage Commands to Cover

#### Start
```bash
VBoxManage snapshot <LABEL> restore <snapshot>
VBoxManage snapshot <LABEL> restorecurrent
VBoxManage startvm <LABEL> --type <MODE>
```

#### Stop (with timeout)
```bash
VBoxManage controlvm <LABEL> poweroff
```

#### List
```bash
VBoxManage list vms
```

#### Status
```bash
VBoxManage showvminfo <LABEL> --machinereadable
```

#### Dump Memory
```bash
VBoxManage debugvm <LABEL> dumpvmcore --filename <PATH>
VBoxManage debugvm <LABEL> dumpguestcore --filename <PATH>
```

#### Dump PCAP
```bash
VBoxManage controlvm <LABEL> nictracefile1 <PCAP_PATH>
VBoxManage controlvm <LABEL> nictrace1 on
```

#### Version
```bash
VBoxManage -v
```

### vboxwebsrv

#### Install on OSX
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
