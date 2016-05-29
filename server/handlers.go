package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

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
	var nameOrID string
	var err error
	if nameOrID, err = strconv.Atoi(vars["nameOrID"]); err != nil {
		panic(err)
	}
	machine := GetMachine(nameOrID)
	assert(machine.Start())
}

// VMStop router stops a VM
func VMStop(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
