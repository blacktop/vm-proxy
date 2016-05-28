# Notes

### VBoxManage Commands to Cover

#### Start
 - VBoxManage "snapshot", label  
 - VBoxManage "restore", machine.snapshot  
 - VBoxManage "restorecurrent", machine.snapshot  


 #### Stop (with timeout)
  - VBoxManage controlvm <label> poweroff  


 #### List
  - VBoxManage list vms  


 #### Status
  - VBoxManage showvminfo <label> --machinereadable  


 #### Dump Memory
 - VBoxManage debugvm <label> dumpvmcore --filename <path>  
 - VBoxManage debugvm <label> dumpguestcore --filename <path>  


 #### Dump PCAP
  - VBoxManage controlvm <label> nictracefile1 <pcap_path>    
  - VBoxManage controlvm <label> nictrace1 on  

#### Version
 - VBoxManage -v
