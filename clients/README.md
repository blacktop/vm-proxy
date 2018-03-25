# clients

## Coverage

### VBoxManage

* [ ] snapshot `label`
* [ ] startvm `label`
* [ ] controlvm `label` nictracefile1 `pcap_path`
* [ ] controlvm `label` nictrace1 on
* [ ] controlvm `label` poweroff
* [ ] list vms
* [ ] showvminfo `label` --machinereadable
* [ ] debugvm `label` `dumpcmd` --filename `path`

### vmrun

* [ ] listSnapshots `vmx_path`
* [ ] revertToSnapshot `vmx_path` `snapshot`
* [ ] deleteSnapshot `vmx_path` memdump
* [ ] snapshot `vmx_path` memdump
* [x] start `vmx_path` headless
* [x] list
* [x] stop `vmx_path` hard
