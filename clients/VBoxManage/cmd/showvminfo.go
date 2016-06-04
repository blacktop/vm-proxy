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
	"log"
	"os"

	"github.com/blacktop/vm-proxy/drivers/virtualbox"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// showvminfoCmd represents the showvminfo command
var showvminfoCmd = &cobra.Command{
	Use:   "showvminfo <uuid|vmname>",
	Short: "Display VM info",
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.GetString("server.host")
		port := viper.GetString("server.port")

		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		d := virtualbox.NewDriver("", "")
		outList, err := d.Status(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(outList)
	},
}

func init() {
	RootCmd.AddCommand(showvminfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showvminfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showvminfoCmd.Flags().BoolP("machinereadable", "", false, "Display machine readable output")

}
