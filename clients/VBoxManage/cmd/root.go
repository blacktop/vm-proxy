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
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var version bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "VBoxManage",
	Short: "Oracle VM VirtualBox Command Line Management Interface Version 5.0.20",
	Long: `Oracle VM VirtualBox Command Line Management Interface Version 5.0.20
(C) 2005-2016 Oracle Corporation
All rights reserved.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			host := viper.GetString("server.host")
			port := viper.GetString("server.port")
			// Create client
			client := &http.Client{}

			// Create request
			req, err := http.NewRequest("GET", "http://"+host+":"+port+"/virtualbox/version", nil)

			// Fetch Request
			resp, err := client.Do(req)
			assert(err)

			// Read Response Body
			respBody, _ := ioutil.ReadAll(resp.Body)

			// Display Results
			// fmt.Println("response Status : ", resp.Status)
			// fmt.Println("response Headers : ", resp.Header)
			// fmt.Println("response Body : ", string(respBody))
			fmt.Print(string(respBody))
		} else {
			cmd.Help()
			os.Exit(0)
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.VBoxManage.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "print version number and exit")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".VBoxManage") // name of config file (without extension)
	viper.AddConfigPath("$HOME")       // adding home directory as first search path
	viper.AutomaticEnv()               // read in environment variables that match

	viper.SetDefault("server.host", "localhost") // Set default host for vm-proxy server
	viper.SetDefault("server.port", 5000)        // Set default port for vm-proxy server

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
