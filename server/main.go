package main

import (
	"net/http"
	"os"
	"path/filepath"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/blacktop/vm-proxy/server/vbox"
	"github.com/blacktop/vm-proxy/server/vmware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	// Version stores the server's version
	Version string
	// BuildTime stores the server's build time
	BuildTime string
	// GitCommit stores server's gitcommit
	GitCommit string
	// Token stores the server api token
	Token string
	// Host server host
	Host string
	// Port server port
	Port string
)

var appHelpTemplate = `Usage: {{.Name}} {{if .Flags}}[OPTIONS] {{end}}COMMAND [arg...]

{{.Usage}}

Version: {{.Version}}{{if or .Author .Email}}

Author:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}
{{if .Flags}}
Options:
  {{range .Flags}}{{.}}
  {{end}}{{end}}
Commands:
  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}
Run '{{.Name}} COMMAND --help' for more information on a command.
`

func main() {

	cli.AppHelpTemplate = appHelpTemplate
	app := cli.NewApp()

	app.Name = "vm-proxy"
	app.Author = "blacktop"
	app.Email = "https://github.com/blacktop"
	app.Version = Version + ", BuildTime: " + BuildTime
	app.Compiled, _ = time.Parse("20060102", BuildTime)
	app.Usage = "VM Proxy Server - Allows hypervisors to be controlled from docker containers"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose, V",
			Usage: "verbose output",
		},
		cli.StringFlag{
			Name:        "host",
			Value:       "dockerhost",
			Usage:       "vm-proxy server host",
			EnvVar:      "VMPROXY_HOST",
			Destination: &Host,
		},
		cli.StringFlag{
			Name:        "port",
			Value:       "3993",
			Usage:       "vm-proxy server port",
			EnvVar:      "VMPROXY_PORT",
			Destination: &Port,
		},
		cli.StringFlag{
			Name:        "token",
			Value:       "",
			Usage:       "webhook token",
			EnvVar:      "VMPROXY_TOKEN",
			Destination: &Token,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "Update certificates",
			Action: func(c *cli.Context) error {
				err := RegenerateCerts("dockerhost")
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
		{
			Name:    "export",
			Aliases: []string{"u"},
			Usage:   "Export Database",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}
	app.Action = func(c *cli.Context) error {

		if c.Bool("verbose") {
			log.SetLevel(log.DebugLevel)
		}

		if len(c.String("token")) == 0 {
			log.Warn("no webhook token set: --token")
		}

		// create http routes
		router := mux.NewRouter().StrictSlash(true)
		// virtualbox routes
		router.HandleFunc("/vbox/version", vbox.Version).Methods("GET")
		router.HandleFunc("/vbox/list", vbox.List).Methods("GET")
		router.HandleFunc("/vbox/status/{nameOrID}", vbox.Status).Methods("GET")
		router.HandleFunc("/vbox/start/{nameOrID}/{startType}", vbox.Start).Methods("GET")
		router.HandleFunc("/vbox/stop/{nameOrID}", vbox.Stop).Methods("GET")
		router.HandleFunc("/vbox/snapshot/restorecurrent/{nameOrID}", vbox.SnapshotRestoreCurrent).Methods("GET")
		router.HandleFunc("/vbox/snapshot/{nameOrID}/restore/{snapShot}", vbox.SnapshotRestore).Methods("GET")
		router.HandleFunc("/vbox/nictracefile1/{nameOrID}/{fileName}", vbox.NicTraceFile).Methods("GET")
		router.HandleFunc("/vbox/nnictrace1/{nameOrID}/{stateOnOff}", vbox.NicTrace).Methods("GET")
		router.HandleFunc("/vbox/debugvm/{nameOrID}/{fileName}", vbox.DumpVM).Methods("GET")
		// vmware routes
		router.HandleFunc("/vmware/list", vmware.List).Methods("GET")
		router.HandleFunc("/vmware/snapshot", vmware.Snapshot).Methods("GET")
		router.HandleFunc("/vmware/snapshot/list/{vmx_path}", vmware.SnapshotList).Methods("GET")
		router.HandleFunc("/vmware/snapshot/revert/{vmx_path}", vmware.SnapshotRevert).Methods("GET")
		router.HandleFunc("/vmware/snapshot/delete/{vmx_path}", vmware.SnapshotDelete).Methods("GET")
		router.HandleFunc("/vmware/start", vmware.Start).Methods("POST")
		router.HandleFunc("/vmware/stop", vmware.Stop).Methods("POST")
		router.HandleFunc("/vmware/info", vmware.Info).Methods("POST")

		err := GenerateCerts("dockerhost")
		if err != nil {
			log.Fatal(err)
		}

		// start microservice
		log.WithFields(log.Fields{
			"host":  Host,
			"port":  Port,
			"token": Token,
		}).Info("vm-proxy service listening")

		home, err := homedir.Dir()
		if err != nil {
			return errors.Wrap(err, "could not detect users home directory")
		}
		certPath := filepath.Join(home, ".vmproxy", "cert.pem")
		keyPath := filepath.Join(home, ".vmproxy", "key.pem")

		loggedRouter := handlers.LoggingHandler(os.Stdout, router)
		log.Fatal(http.ListenAndServeTLS(":"+Port, certPath, keyPath, loggedRouter))

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
