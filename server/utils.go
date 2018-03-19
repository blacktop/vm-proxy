package main

import (
	"path/filepath"

	"github.com/kabukky/httpscerts"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

// GenerateCerts generates SSL certs for vm-proxy server
func GenerateCerts(host string) error {
	home, err := homedir.Dir()
	if err != nil {
		return errors.Wrap(err, "could not detect users home directory")
	}

	certPath := filepath.Join(home, ".vmproxy", "cert.pem")
	keyPath := filepath.Join(home, ".vmproxy", "key.pem")

	// Check if the cert files are available.
	err = httpscerts.Check(certPath, keyPath)
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate(certPath, keyPath, host)
		if err != nil {
			return errors.Wrap(err, "could not create https certs")
		}
	}
	return nil
}
