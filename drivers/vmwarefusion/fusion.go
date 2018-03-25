/*
 * Copyright 2014 VMware, Inc.  All rights reserved.  Licensed under the Apache v2 License.
 */

package vmwarefusion

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/blacktop/vm-proxy/drivers"
	"github.com/docker/machine/libmachine/state"
	cryptossh "golang.org/x/crypto/ssh"
)

const (
	defaultSSHUser     = "admin"
	defaultMachineName = "cuckoo"
)

// Driver for VMware Fusion
type Driver struct {
	*drivers.BaseDriver
	Memory         int
	DiskSize       int
	CPU            int
	ISO            string
	Boot2DockerURL string

	SSHPassword    string
	ConfigDriveISO string
	ConfigDriveURL string
	NoShare        bool

	VMXPath string
}

func NewDriver(vmxPath, storePath string) drivers.Driver {
	return &Driver{
		// CPU:      defaultCPU,
		// Memory:   defaultMemory,
		// DiskSize: defaultDiskSize,
		// SSHPassword: defaultSSHPass,
		VMXPath: vmxPath,
		BaseDriver: &drivers.BaseDriver{
			SSHUser:     defaultSSHUser,
			MachineName: defaultMachineName,
			StorePath:   storePath,
		},
	}
}

// List lists all VMs
func (d *Driver) List() (string, error) {
	stdOut, _, err := vmrun("list")
	if err != nil {
		return "", err
	}
	return stdOut, nil
}

func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "vmwarefusion"
}

func (d *Driver) GetIP() (string, error) {
	s, err := d.GetState()
	if err != nil {
		return "", err
	}
	if s != state.Running {
		return "", drivers.ErrHostIsNotRunning
	}

	// determine MAC address for VM
	macaddr, err := d.getMacAddressFromVmx()
	if err != nil {
		return "", err
	}

	// attempt to find the address in the vmnet configuration
	if ip, err := d.getIPfromVmnetConfiguration(macaddr); err == nil {
		return ip, err
	}

	// address not found in vmnet so look for a DHCP lease
	ip, err := d.getIPfromDHCPLease(macaddr)
	if err != nil {
		return "", err
	}

	return ip, nil
}

func (d *Driver) GetState() (state.State, error) {
	// VMRUN only tells use if the vm is running or not
	vmxp, err := filepath.EvalSymlinks(d.VMXPath)
	if err != nil {
		return state.Error, err
	}
	if stdout, _, _ := vmrun("list"); strings.Contains(stdout, vmxp) {
		return state.Running, nil
	}
	return state.Stopped, nil
}

func (d *Driver) Create() error {
	return fmt.Errorf("Create() has not been implimented")
}

func (d *Driver) Start() error {
	vmrun("start", d.VMXPath, "nogui")

	// Do not execute the rest of boot2docker specific configuration, exit here
	if d.ConfigDriveURL != "" {
		log.Debugf("Leaving start sequence early, configdrive found")
		return nil
	}

	// log.Debugf("Mounting Shared Folders...")
	// var shareName, shareDir string // TODO configurable at some point
	// switch runtime.GOOS {
	// case "darwin":
	// 	shareName = "Users"
	// 	shareDir = "/Users"
	// 	// TODO "linux" and "windows"
	// }

	// if shareDir != "" {
	// 	if _, err := os.Stat(shareDir); err != nil && !os.IsNotExist(err) {
	// 		return err
	// 	} else if !os.IsNotExist(err) {
	// 		// create mountpoint and mount shared folder
	// 		command := "[ ! -d " + shareDir + " ]&& sudo mkdir " + shareDir + "; sudo mount --bind /mnt/hgfs/" + shareDir + " " + shareDir + " || [ -f /usr/local/bin/vmhgfs-fuse ]&& sudo /usr/local/bin/vmhgfs-fuse -o allow_other .host:/" + shareName + " " + shareDir + " || sudo mount -t vmhgfs -o uid=$(id -u),gid=$(id -g) .host:/" + shareName + " " + shareDir
	// 		vmrun("-gu", B2DUser, "-gp", B2DPass, "runScriptInGuest", d.VMXPath, "/bin/sh", command)
	// 	}
	// }

	return nil
}

func (d *Driver) Stop() error {
	_, _, err := vmrun("stop", d.VMXPath, "nogui")
	return err
}

func (d *Driver) Restart() error {
	// Stop VM gracefully
	if err := d.Stop(); err != nil {
		return err
	}
	// Start it again and mount shared folder
	if err := d.Start(); err != nil {
		return err
	}
	return nil
}

func (d *Driver) Kill() error {
	_, _, err := vmrun("stop", d.VMXPath, "hard nogui")
	return err
}

func (d *Driver) Remove() error {
	s, _ := d.GetState()
	if s == state.Running {
		if err := d.Kill(); err != nil {
			return fmt.Errorf("Error stopping VM before deletion")
		}
	}
	log.Infof("Deleting %s...", d.MachineName)
	vmrun("deleteVM", d.VMXPath, "nogui")
	return nil
}

func (d *Driver) Upgrade() error {
	return fmt.Errorf("VMware Fusion does not currently support the upgrade operation")
}

func (d *Driver) vmxPath() (string, error) {
	if _, err := os.Stat(d.VMXPath); os.IsNotExist(err) {
		return "", fmt.Errorf("vmx:%s file does not exist", d.VMXPath)
	}
	return d.VMXPath, nil
}

func (d *Driver) vmdkPath() string {
	return d.ResolveStorePath(fmt.Sprintf("%s.vmdk", d.MachineName))
}

func (d *Driver) getMacAddressFromVmx() (string, error) {
	var vmxfh *os.File
	var vmxcontent []byte
	var err error

	if vmxfh, err = os.Open(d.VMXPath); err != nil {
		return "", err
	}
	defer vmxfh.Close()

	if vmxcontent, err = ioutil.ReadAll(vmxfh); err != nil {
		return "", err
	}

	// Look for generatedAddress as we're passing a VMX with addressType = "generated".
	var macaddr string
	vmxparse := regexp.MustCompile(`^ethernet0.generatedAddress\s*=\s*"(.*?)"\s*$`)
	for _, line := range strings.Split(string(vmxcontent), "\n") {
		if matches := vmxparse.FindStringSubmatch(line); matches == nil {
			continue
		} else {
			macaddr = strings.ToLower(matches[1])
		}
	}

	if macaddr == "" {
		return "", fmt.Errorf("couldn't find MAC address in VMX file %s", d.VMXPath)
	}

	log.Debugf("MAC address in VMX: %s", macaddr)

	return macaddr, nil
}

func (d *Driver) getIPfromVmnetConfiguration(macaddr string) (string, error) {

	// DHCP lease table for NAT vmnet interface
	confFiles, _ := filepath.Glob("/Library/Preferences/VMware Fusion/vmnet*/dhcpd.conf")
	for _, conffile := range confFiles {
		log.Debugf("Trying to find IP address in configuration file: %s", conffile)
		if ipaddr, err := d.getIPfromVmnetConfigurationFile(conffile, macaddr); err == nil {
			return ipaddr, err
		}
	}

	return "", fmt.Errorf("IP not found for MAC %s in vmnet configuration files", macaddr)
}

func (d *Driver) getIPfromVmnetConfigurationFile(conffile, macaddr string) (string, error) {
	var conffh *os.File
	var confcontent []byte

	var currentip string
	var lastipmatch string
	var lastmacmatch string

	var err error

	if conffh, err = os.Open(conffile); err != nil {
		return "", err
	}
	defer conffh.Close()

	if confcontent, err = ioutil.ReadAll(conffh); err != nil {
		return "", err
	}

	// find all occurrences of 'host .* { .. }' and extract
	// out of the inner block the MAC and IP addresses

	// key = MAC, value = IP
	m := make(map[string]string)

	// Begin of a host block, that contains the IP, MAC
	hostbegin := regexp.MustCompile(`^host (.+?) {`)
	// End of a host block
	hostend := regexp.MustCompile(`^}`)

	// Get the IP address.
	ip := regexp.MustCompile(`^\s*fixed-address (.+?);$`)
	// Get the MAC address associated.
	mac := regexp.MustCompile(`^\s*hardware ethernet (.+?);$`)

	// we use a block depth so that just in case inner blocks exists
	// we are not being fooled by them
	blockdepth := 0
	for _, line := range strings.Split(string(confcontent), "\n") {

		if matches := hostbegin.FindStringSubmatch(line); matches != nil {
			blockdepth = blockdepth + 1
			continue
		}

		// we are only in interested in endings if we in a block. Otherwise we will count
		// ending of non host blocks as well
		if matches := hostend.FindStringSubmatch(line); blockdepth > 0 && matches != nil {
			blockdepth = blockdepth - 1

			if blockdepth == 0 {
				// add data
				m[lastmacmatch] = lastipmatch

				// reset all temp var holders
				lastipmatch = ""
				lastmacmatch = ""
			}

			continue
		}

		// only if we are within the first level of a block
		// we are looking for addresses to extract
		if blockdepth == 1 {
			if matches := ip.FindStringSubmatch(line); matches != nil {
				lastipmatch = matches[1]
				continue
			}

			if matches := mac.FindStringSubmatch(line); matches != nil {
				lastmacmatch = strings.ToLower(matches[1])
				continue
			}
		}
	}

	log.Debugf("Following IPs found %s", m)

	// map is filled to now lets check if we have a MAC associated to an IP
	currentip, ok := m[strings.ToLower(macaddr)]

	if !ok {
		return "", fmt.Errorf("IP not found for MAC %s in vmnet configuration", macaddr)
	}

	log.Debugf("IP found in vmnet configuration file: %s", currentip)

	return currentip, nil

}

func (d *Driver) getIPfromDHCPLease(macaddr string) (string, error) {

	// DHCP lease table for NAT vmnet interface
	leasesFiles, _ := filepath.Glob("/var/db/vmware/*.leases")
	for _, dhcpfile := range leasesFiles {
		log.Debugf("Trying to find IP address in leases file: %s", dhcpfile)
		if ipaddr, err := d.getIPfromDHCPLeaseFile(dhcpfile, macaddr); err == nil {
			return ipaddr, err
		}
	}

	return "", fmt.Errorf("IP not found for MAC %s in DHCP leases", macaddr)
}

func (d *Driver) getIPfromDHCPLeaseFile(dhcpfile, macaddr string) (string, error) {

	var dhcpfh *os.File
	var dhcpcontent []byte
	var lastipmatch string
	var currentip string
	var lastleaseendtime time.Time
	var currentleadeendtime time.Time
	var err error

	if dhcpfh, err = os.Open(dhcpfile); err != nil {
		return "", err
	}
	defer dhcpfh.Close()

	if dhcpcontent, err = ioutil.ReadAll(dhcpfh); err != nil {
		return "", err
	}

	// Get the IP from the lease table.
	leaseip := regexp.MustCompile(`^lease (.+?) {$`)
	// Get the lease end date time.
	leaseend := regexp.MustCompile(`^\s*ends \d (.+?);$`)
	// Get the MAC address associated.
	leasemac := regexp.MustCompile(`^\s*hardware ethernet (.+?);$`)

	for _, line := range strings.Split(string(dhcpcontent), "\n") {

		if matches := leaseip.FindStringSubmatch(line); matches != nil {
			lastipmatch = matches[1]
			continue
		}

		if matches := leaseend.FindStringSubmatch(line); matches != nil {
			lastleaseendtime, _ = time.Parse("2006/01/02 15:04:05", matches[1])
			continue
		}

		if matches := leasemac.FindStringSubmatch(line); matches != nil && matches[1] == macaddr && currentleadeendtime.Before(lastleaseendtime) {
			currentip = lastipmatch
			currentleadeendtime = lastleaseendtime
		}
	}

	if currentip == "" {
		return "", fmt.Errorf("IP not found for MAC %s in DHCP leases", macaddr)
	}

	log.Debugf("IP found in DHCP lease table: %s", currentip)

	return currentip, nil
}

// execute command over SSH with user / password authentication
func executeSSHCommand(command string, d *Driver) error {
	log.Debugf("Execute executeSSHCommand: %s", command)

	config := &cryptossh.ClientConfig{
		User: d.SSHUser,
		Auth: []cryptossh.AuthMethod{
			cryptossh.Password(d.SSHPassword),
		},
	}

	client, err := cryptossh.Dial("tcp", fmt.Sprintf("%s:%d", d.IPAddress, d.SSHPort), config)
	if err != nil {
		log.Debugf("Failed to dial:", err)
		return err
	}

	session, err := client.NewSession()
	if err != nil {
		log.Debugf("Failed to create session: " + err.Error())
		return err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b

	if err := session.Run(command); err != nil {
		log.Debugf("Failed to run: " + err.Error())
		return err
	}
	log.Debugf("Stdout from executeSSHCommand: %s", b.String())

	return nil
}
