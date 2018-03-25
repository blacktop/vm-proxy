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
	"net/url"
	"path/filepath"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a VM",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a VMX path")
		}
		// if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		// 	return fmt.Errorf("vmx:%s does not exist", args[0])
		// }
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

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

		v := url.Values{}
		v.Set("vmx_path", args[0])

		// Create request
		req, err := http.NewRequest("POST", "https://"+host+":"+port+"/vmware/stop", strings.NewReader(v.Encode()))
		assert(err)

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
	RootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
