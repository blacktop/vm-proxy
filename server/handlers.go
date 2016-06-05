package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	vbox "github.com/blacktop/vm-proxy/drivers/virtualbox"
	"github.com/gorilla/mux"
	"github.com/riobard/go-virtualbox"
)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// Index root route
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

// VBoxVersion route returns VBoxManage version
func VBoxVersion(w http.ResponseWriter, r *http.Request) {
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

// VBoxList route lists all VMs
func VBoxList(w http.ResponseWriter, r *http.Request) {
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

// VBoxStatus displays the machine readable status of a VM
func VBoxStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]

	machine, err := virtualbox.GetMachine(nameOrID)
	assert(err)
	if err := json.NewEncoder(w).Encode(machine.State); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}

// VBoxStart router starts a VM
func VBoxStart(w http.ResponseWriter, r *http.Request) {
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

// VBoxStop router stops a VM
func VBoxStop(w http.ResponseWriter, r *http.Request) {
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

// VBoxSnapshotRestore restores a certain snapshot
func VBoxSnapshotRestore(w http.ResponseWriter, r *http.Request) {
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

// VBoxSnapshotRestoreCurrent restores the most resent snapshot
func VBoxSnapshotRestoreCurrent(w http.ResponseWriter, r *http.Request) {
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

// VBoxNicTrace restores a certain snapshot
func VBoxNicTrace(w http.ResponseWriter, r *http.Request) {
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

// VBoxNicTraceFile restores the most resent snapshot
func VBoxNicTraceFile(w http.ResponseWriter, r *http.Request) {
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

// VBoxDumpVM perform a memory core dump (VirtualBox version 5.x)
func VBoxDumpVM(w http.ResponseWriter, r *http.Request) {
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

// VBoxDumpGuest perform a memory core dump (VirtualBox version 4.x)
func VBoxDumpGuest(w http.ResponseWriter, r *http.Request) {
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
