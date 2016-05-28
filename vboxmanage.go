package main

import (
	"fmt"
	"log"

	"github.com/riobard/go-virtualbox"
)

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ListMachines prints out a list of machines
func ListMachines() {
	machines, err := virtualbox.ListMachines()
	assert(err)
	for _, machine := range machines {
		fmt.Println(machine.Name)
	}
}

func StartVM(nameOrID string) {
	machine := GetMachine(nameOrID)
	assert(machine.Start())
}

func StopVM(nameOrID string) {
	machine := GetMachine(nameOrID)
	assert(machine.Stop())
}

func GetMachine(nameOrID string) *virtualbox.Machine {
	machine, err := virtualbox.GetMachine(nameOrID)
	assert(err)
	// fmt.Println()
	// fmt.Println("Name: ", machine.Name)
	// fmt.Println("UUID: ", machine.UUID)
	// fmt.Println("State: ", machine.State)
	// fmt.Println("CPUs: ", machine.CPUs)
	// fmt.Println("Memory: ", machine.Memory)
	// fmt.Println("VRAM: ", machine.VRAM)
	// fmt.Println("CfgFile: ", machine.CfgFile)
	// fmt.Println("BaseFolder: ", machine.BaseFolder)
	// fmt.Println("OSType: ", machine.OSType)
	// fmt.Println("Flag: ", machine.Flag)
	// fmt.Println("BootOrder: ", machine.BootOrder)

	return machine
}

func main() {
	ListMachines()
	fmt.Println("Starting VM...")
	StartVM("default")
	fmt.Println("Stoping VM...")
	StopVM("default")
}
