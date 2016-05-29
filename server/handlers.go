package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

// VMList route lists all VMs
func VMList(w http.ResponseWriter, r *http.Request) {
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

// VMStart router starts a VM
func VMStart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// var nameOrID string
	var err error
	nameOrID := vars["nameOrID"]
	machine, err := virtualbox.GetMachine(nameOrID)
	assert(err)
	assert(machine.Start())
}

// VMStop router stops a VM
func VMStop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// var nameOrID string
	var err error
	nameOrID := vars["nameOrID"]
	machine, err := virtualbox.GetMachine(nameOrID)
	assert(err)
	assert(machine.Stop())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
