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
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all running VMs",
	Run: func(cmd *cobra.Command, args []string) {
		// d := virtualbox.NewDriver("", "")
		// outList, err := d.List()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Print(outList)

		host := viper.GetString("server.host")
		port := viper.GetString("server.port")

		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(errors.Wrap(err, "could not detect users home directory"))
		}
		// Create client
		caCert, err := ioutil.ReadFile(filepath.Join(home, ".vmproxy", "cert.pem"))
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
		req, err := http.NewRequest("GET", "https://"+host+":"+port+"/vmware/list", nil)

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
