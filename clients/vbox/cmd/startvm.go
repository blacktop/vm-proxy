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

// Type vm start mode
var Type string

// startvmCmd represents the startvm command
var startvmCmd = &cobra.Command{
	Use:   "startvm",
	Short: "Start VMs",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Help()
			os.Exit(0)
		}

		var startType string
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

		switch Type {
		case "gui":
			startType = Type
		case "headless":
			startType = Type
		case "separate":
			startType = Type
		default:
			startType = "headless"
		}

		// Create request
		req, err := http.NewRequest("GET", "http://"+host+":"+port+"/virtualbox/start/"+args[0]+"/"+startType, nil)
		assert(err)

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
	RootCmd.AddCommand(startvmCmd)
	startvmCmd.PersistentFlags().StringVarP(&Type, "type", "", "headless", "gui|headless|separate")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startvmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startvmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
