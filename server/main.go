package main

import (
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/blacktop/vm-proxy/server/vbox"
	"github.com/blacktop/vm-proxy/server/vmware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
			Value:       "127.0.0.1",
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
			Usage:   "Update images",
			Action: func(c *cli.Context) error {
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
		router.HandleFunc("/vbox/list", vbox.List).Methods("GET")
		// router.HandleFunc("/vbox/{source:(?:giphy|xkcd|dilbert|default|contrib)}/{file}", updateImageKeywords).Methods("PATCH")
		// vmware routes
		router.HandleFunc("/vmware/list", vmware.List).Methods("GET")
		// router.HandleFunc("/vmware/{source:(?:giphy|xkcd|dilbert|default|contrib)}/{file}", updateImageKeywords).Methods("PATCH")

		err := GenerateCerts(Host + ":" + Port)
		if err != nil {
			log.Fatal(err)
		}

		// start microservice
		log.WithFields(log.Fields{
			"host":  Host,
			"port":  Port,
			"token": Token,
		}).Info("vm-proxy service listening")

		loggedRouter := handlers.LoggingHandler(os.Stdout, router)
		log.Fatal(http.ListenAndServeTLS(":"+Port, "cert.pem", "key.pem", loggedRouter))

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
