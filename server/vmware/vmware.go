package vmware

import (
	"net/http"

	"github.com/blacktop/vm-proxy/drivers/vmwarefusion"
	"github.com/gorilla/mux"
)

// List route lists all VMs
func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	d := vmwarefusion.NewDriver("", "")
	outPut := d.List()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// }
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

func Snapshot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	d := vmwarefusion.NewDriver("", "")
	outPut := d.List()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// }
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

func SnapshotList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	vmxPath := vars["vmx_path"]

	d := vmwarefusion.NewDriver("", "")
	outPut := d.List()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// }
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

func SnapshotRevert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	vmxPath := vars["vmx_path"]

	d := vmwarefusion.NewDriver("", "")
	outPut := d.List()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// }
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

func SnapshotDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	vmxPath := vars["vmx_path"]

	d := vmwarefusion.NewDriver("", "")
	outPut := d.List()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// }
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

func Start(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	vmxPath := vars["vmx_path"]

	d := vmwarefusion.NewDriver("", "")
	outPut := d.List()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// }
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}

func Stop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	vars := mux.Vars(r)
	vmxPath := vars["vmx_path"]

	d := vmwarefusion.NewDriver("", "")
	outPut := d.List()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// }
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}
