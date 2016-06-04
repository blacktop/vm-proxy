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

// VBoxList route lists all VMs
func VBoxList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	machines, err := virtualbox.ListMachines()
	assert(err)
	for _, machine := range machines {
		fmt.Println(machine.Name)
	}

	if err := json.NewEncoder(w).Encode(machines); err != nil {
		panic(err)
	}
}

// VBoxStatus displays the machine readable status of a VM
func VBoxStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	// var nameOrID string
	// var err error
	nameOrID := vars["nameOrID"]
	machine, err := virtualbox.GetMachine(nameOrID)
	assert(err)
	if err := json.NewEncoder(w).Encode(machine.State); err != nil {
		panic(err)
	}
}

// VBoxStart router starts a VM
func VBoxStart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// var nameOrID string
	// var err error
	nameOrID := vars["nameOrID"]
	machine, err := virtualbox.GetMachine(nameOrID)
	assert(err)
	assert(machine.Start())
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// VBoxStop router stops a VM
func VBoxStop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]

	d := vbox.NewDriver(nameOrID, "")
	assert(d.Stop())
}

// VBoxSnapshotRestore restores a certain snapshot
func VBoxSnapshotRestore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]
	snapShot := vars["snapShot"]
	fmt.Println(nameOrID)
	fmt.Println(snapShot)
	d := vbox.NewDriver(nameOrID, "")
	outPut, err := d.RestoreSnapshot(nameOrID, snapShot)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(outPut))
}

// VBoxSnapshotRestoreCurrent restores the most resent snapshot
func VBoxSnapshotRestoreCurrent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nameOrID := vars["nameOrID"]

	d := vbox.NewDriver(nameOrID, "")
	outPut, err := d.RestoreCurrentSnapshot(nameOrID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(outPut))
}
