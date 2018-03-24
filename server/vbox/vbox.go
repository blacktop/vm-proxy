package vbox

import (
	"net/http"

	vbox "github.com/blacktop/vm-proxy/drivers/virtualbox"
	"github.com/gorilla/mux"
)

// Version route returns VBoxManage version
func Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	d := vbox.NewDriver("", "")
	outPut, err := d.Version()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// List route lists all VMs
func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// machines, err := virtualbox.ListMachines()
	// assert(err)
	// for _, machine := range machines {
	// 	fmt.Println(machine.Name)
	// }

	// if err := json.NewEncoder(w).Encode(machines); err != nil {
	// 	panic(err)
	// }
	d := vbox.NewDriver("", "")
	outPut, err := d.ListVMs()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// Status displays the machine readable status of a VM
func Status(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]

	// machine, err := virtualbox.GetMachine(nameOrID)
	// assert(err)
	// if err := json.NewEncoder(w).Encode(machine.State); err != nil {
	// 	panic(err)
	// }

	d := vbox.NewDriver(nameOrID, "")
	outPut, err := d.Status()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// Start router starts a VM
func Start(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]
	startType := vars["startType"]

	d := vbox.NewDriver(nameOrID, "")
	outPut, err := d.StartVM(startType)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// Stop router stops a VM
func Stop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]

	d := vbox.NewDriver(nameOrID, "")
	outPut, err := d.StopVM()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// SnapshotRestore restores a certain snapshot
func SnapshotRestore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]
	snapShot := vars["snapShot"]

	d := vbox.NewDriver(nameOrID, "")
	outPut, err := d.RestoreSnapshot(nameOrID, snapShot)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// SnapshotRestoreCurrent restores the most resent snapshot
func SnapshotRestoreCurrent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]

	d := vbox.NewDriver(nameOrID, "")
	outPut, err := d.RestoreCurrentSnapshot(nameOrID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// NicTrace restores a certain snapshot
func NicTrace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]
	stateOnOff := vars["stateOnOff"]

	d := vbox.NewDriver(nameOrID, "")
	outPut, err := d.NicTrace(stateOnOff)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// NicTraceFile restores the most resent snapshot
func NicTraceFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]
	fileName := vars["fileName"]

	d := vbox.NewDriver(nameOrID, "")
	outPut, err := d.NicTraceFile(fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// DumpVM perform a memory core dump (VirtualBox version 5.x)
func DumpVM(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]
	fileName := vars["fileName"]

	d := vbox.NewDriver(nameOrID, "")
	outPut, err := d.DumpVM(fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// DumpGuest perform a memory core dump (VirtualBox version 4.x)
func DumpGuest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]
	fileName := vars["fileName"]

	d := vbox.NewDriver(nameOrID, "")
	outPut, err := d.DumpGuest(fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}
