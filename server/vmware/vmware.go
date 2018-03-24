package vmware

import (
	"net/http"

	vmware "github.com/blacktop/vm-proxy/drivers/vmwarefusion"
)

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
	d := vmware.NewDriver("", "")
	outPut := d.DriverName()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// }
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(outPut))
}
