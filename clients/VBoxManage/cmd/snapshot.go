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

var projectBase string
var userLicense string

// snapshotCmd represents the snapshot command
var snapshotCmd = &cobra.Command{
	Use:   "snapshot <uuid|vmname>",
	Short: "Manage virtualbox snapshots",
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.GetString("server.host")
		port := viper.GetString("server.port")
		// if len(args) == 0 {
		// 	cmd.Help()
		// 	os.Exit(0)
		// }
		// d := virtualbox.NewDriver("", "")
		// if len(args) == 3 {
		// 	if strings.EqualFold("restore", args[1]) {
		// 		outList, err := d.Snapshot(args[0], args[2])
		// 		if err != nil {
		// 			log.Fatal(err)
		// 		}
		// 		fmt.Print(outList)
		// 	}
		// }
		// if strings.EqualFold("restorecurrent", args[1]) {
		// 	outList, err := d.Snapshot(args[0], "")
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Print(outList)
		// }
	},
}

var restoreCmd = &cobra.Command{
	Use:   "restore <uuid|snapname>",
	Short: "Restore snapshot",
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

var restorecurrentCmd = &cobra.Command{
	Use:   "restorecurrent",
	Short: "A brief description of your command",
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
	tmplt := `Usage:{{if .Runnable}}
  {{if .HasAvailableFlags}}{{appendIfNotPresent .UseLine ""}}{{else}}{{.UseLine}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
  {{ .CommandPath}} <uuid|vmname> [command]{{end}}{{if gt .Aliases 0}}
Aliases:
  {{.NameAndAliases}}
{{end}}{{if .HasExample}}

Examples:
{{ .Example }}{{end}}{{ if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}{{ if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimRightSpace}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableSubCommands }}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
	RootCmd.AddCommand(snapshotCmd)
	snapshotCmd.AddCommand(restoreCmd)
	snapshotCmd.AddCommand(restorecurrentCmd)
	snapshotCmd.SetUsageTemplate(tmplt)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// snapshotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// snapshotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
