# Notes

### VBoxManage Commands to Cover

#### Start
```bash
VBoxManage "snapshot", label
VBoxManage "restore", machine.snapshot
VBoxManage "restorecurrent", machine.snapshot
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
