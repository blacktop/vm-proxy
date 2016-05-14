# vboxmanage-proxy
VBoxManage proxy through vboxwebsrv

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