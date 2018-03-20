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
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// snapshotCmd represents the snapshot command
var snapshotCmd = &cobra.Command{
	Use:   "snapshot <uuid|vmname>",
	Short: "Manage VirtualBox Snapshots",
	Run: func(cmd *cobra.Command, args []string) {

		var req *http.Request
		var err error

		if len(args) < 2 {
			cmd.Help()
			os.Exit(0)
		}

		host := viper.GetString("server.host")
		port := viper.GetString("server.port")

		// Create client
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		// Create client
		caCert, err := ioutil.ReadFile(filepath.Join(usr.HomeDir, ".vmproxy", "cert.pem"))
		if err != nil {
			log.Fatal(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		// cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
		// if err != nil {
		// 	log.Fatal(err)
		// }

		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
					// Certificates: []tls.Certificate{cert},
				},
			},
		}

		// Create request
		switch args[1] {
		case "restore":
			req, err = http.NewRequest("GET", "http://"+host+":"+port+"/virtualbox/snapshot/"+args[0]+"/restore/"+args[2], nil)
			assert(err)
		case "restorecurrent":
			req, err = http.NewRequest("GET", "http://"+host+":"+port+"/virtualbox/snapshot/restorecurrent/"+args[0], nil)
			assert(err)
		}

		if req != nil {
			// Fetch Request
			resp, err := client.Do(req)
			assert(err)

			// Read Response Body
			respBody, _ := ioutil.ReadAll(resp.Body)

			// Display Results
			fmt.Print(string(respBody))
		} else {
			cmd.Help()
			os.Exit(1)
		}
	},
}

var restoreCmd = &cobra.Command{
	Use:   "restore <uuid|snapname>",
	Short: "Restore snapshot",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var restorecurrentCmd = &cobra.Command{
	Use:   "restorecurrent",
	Short: "Restore most recent snapshot",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(snapshotCmd)
	snapshotCmd.AddCommand(restoreCmd)
	snapshotCmd.AddCommand(restorecurrentCmd)
	snapshotCmd.SetUsageTemplate(snapshotTmplt)
	restoreCmd.SetUsageTemplate(restoreTmplt)
	restorecurrentCmd.SetUsageTemplate(restoreCurrentTmplt)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// snapshotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// snapshotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
