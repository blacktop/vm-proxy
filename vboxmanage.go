package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/blacktop/vboxmanage-proxy/client"
// )

// func main() {
// 	url := "http://127.0.0.1:18083"
// 	vb := client.NewVirtualBox("", "", url)
// 	if err := vb.Logon(); err != nil {
// 		log.Fatal("Logon failed:", err)
// 	}
// 	machines, err := vb.GetMachines()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, machine := range machines {
// 		fmt.Println(machine.GetName())
// 	}
// 	m, err := vb.FindMachine("default")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = vb.PopulateMachineInfo(m)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	name := m.GetName()
// 	id := m.GetID()
// 	fmt.Println(id, ": ", name)

// 	return
// }

// import (
// 	"fmt"
// 	"log"
// 	"strings"

// 	"github.com/riobard/go-virtualbox"
// )

// func main() {
// 	machines, err := virtualbox.ListMachines()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, m := range machines {
// 		if strings.Compare(m.Name, "default") == 0 {
// 			_ = m.Start()
// 		} else {
// 			_ = m.Stop()
// 		}
// 		fmt.Println(m.Name, ": ", m.State)
// 	}
// }

// package main

import (
	"fmt"
	"log"
	"os"

	client "github.com/appropriate/go-virtualboxclient/virtualboxclient"
)

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	url := "http://127.0.0.1:18083"
	if len(os.Args) >= 2 {
		url = os.Args[1]
	}

	client := client.New("", "", url, false, "")
	if err := client.Logon(); err != nil {
		log.Fatalf("Unable to log on to vboxwebsrv: %v\n", err)
	}

	myDefault, err := client.FindMachine("default")
	assert(err)

	assert(myDefault.Stop())

	machines, err := client.GetMachines()
	assert(err)

	sp, err := client.GetSystemProperties()
	assert(err)

	defer sp.Release()

	for _, m := range machines {
		defer m.Release()

		chipset, err := m.GetChipsetType()
		assert(err)

		mna, err := sp.GetMaxNetworkAdapters(chipset)
		assert(err)

		for i := uint32(0); i < mna; i++ {
			na, err := m.GetNetworkAdapter(i)
			assert(err)

			mac, err := na.GetMACAddress()
			assert(err)

			id, err := m.GetID()
			assert(err)
			m.ID = id

			name, err := m.GetName()
			assert(err)
			m.Name = name

			fmt.Printf("%s: %s, %s\n", m.ID, m.Name, mac)
		}
	}

	return
}
