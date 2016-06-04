// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// debugvmCmd represents the debugvm command
var debugvmCmd = &cobra.Command{
	Use:   "debugvm <uuid|vmname>",
	Short: "Introspection and guest debugging",
	Run: func(cmd *cobra.Command, args []string) {
		// VBoxManage debugvm <LABEL> dumpvmcore --filename <PATH>
		// VBoxManage debugvm <LABEL> dumpguestcore --filename <PATH>
		host := viper.GetString("server.host")
		port := viper.GetString("server.port")

		// if len(args) == 0 {
		// 	cmd.Help()
		// 	os.Exit(0)
		// }
		// d := virtualbox.NewDriver("", "")
		// if len(args) == 3 {
		// 	if strings.EqualFold("dumpvmcore", args[1]) {
		// 		outList, err := d.Snapshot(args[0], args[2])
		// 		if err != nil {
		// 			log.Fatal(err)
		// 		}
		// 		fmt.Print(outList)
		// 	}
		// }
		// if strings.EqualFold("dumpguestcore", args[1]) {
		// 	outList, err := d.Snapshot(args[0], "")
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Print(outList)
		// }
	},
}

var dumpvmcoreCmd = &cobra.Command{
	Use:   "dumpvmcore",
	Short: "dump memory (VirtualBox version 5.x)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		// d := virtualbox.NewDriver("", "")
		// outList, err := d.Snapshot(args[0], args[2])
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Print(outList)
	},
}

var dumpguestcoreCmd = &cobra.Command{
	Use:   "dumpguestcore",
	Short: "dump memory (VirtualBox version 4.x)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		// d := virtualbox.NewDriver("", "")
		// outList, err := d.Snapshot(args[0], "")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Print(outList)
	},
}

func init() {
	RootCmd.AddCommand(debugvmCmd)
	debugvmCmd.AddCommand(dumpvmcoreCmd)
	debugvmCmd.AddCommand(dumpguestcoreCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	debugvmCmd.Flags().String("filename", "", "filesystem path of memory dump")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// debugvmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
