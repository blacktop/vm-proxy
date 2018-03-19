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
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all VMs",
	Run: func(cmd *cobra.Command, args []string) {
		// d := virtualbox.NewDriver("", "")
		// outList, err := d.List()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Print(outList)

		host := viper.GetString("server.host")
		port := viper.GetString("server.port")

		// Create client
		client := &http.Client{}

		// Create request
		req, err := http.NewRequest("GET", "http://"+host+":"+port+"/virtualbox/list", nil)

		// Fetch Request
		resp, err := client.Do(req)
		assert(err)

		// Read Response Body
		respBody, _ := ioutil.ReadAll(resp.Body)

		// Display Results
		fmt.Print(string(respBody))
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
