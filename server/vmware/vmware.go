package vmware

import (
	"errors"
	"net/http"

	"github.com/apex/log"
	"github.com/blacktop/vm-proxy/drivers/vmwarefusion"
	"github.com/gorilla/mux"
)

// List route lists all VMs
func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	d := vmwarefusion.NewDriver("", "")
	outPut, err := d.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// Snapshot takes a snapshot of the VM
func Snapshot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	d := vmwarefusion.NewDriver("", "")
	outPut, err := d.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// SnapshotList lists all snapshots for a given vmx
func SnapshotList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	vmxPath := vars["vmx_path"]

	d := vmwarefusion.NewDriver(vmxPath, "")
	outPut, err := d.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// SnapshotRevert reverts a snapshot back to it's previous snapshot
func SnapshotRevert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	vmxPath := vars["vmx_path"]

	d := vmwarefusion.NewDriver(vmxPath, "")
	outPut, err := d.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// SnapshotDelete deletes last snapshot
func SnapshotDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	vmxPath := vars["vmx_path"]

	d := vmwarefusion.NewDriver(vmxPath, "")
	outPut, err := d.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

// Start starts a VM for a given vmx
func Start(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	r.ParseForm()

	if len(r.Form["vmx_path"]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		err := errors.New("bad request - please supply `vmx_path` params")
		w.Write([]byte(err.Error()))
		log.WithError(err).Error("vmware start failed")
		return
	}

	d := vmwarefusion.NewDriver(r.Form["vmx_path"][0], "")

	err := d.Start()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}

// Stop stops a VM for a given vmx
func Stop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	r.ParseForm()

	if len(r.Form["vmx_path"]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		err := errors.New("bad request - please supply `vmx_path` params")
		w.Write([]byte(err.Error()))
		log.WithError(err).Error("vmware start failed")
		return
	}

	d := vmwarefusion.NewDriver(r.Form["vmx_path"][0], "")

	err := d.Stop()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}
