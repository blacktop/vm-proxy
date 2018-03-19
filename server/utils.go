package main

import (
	"github.com/kabukky/httpscerts"
	"github.com/pkg/errors"
)

// GenerateCerts generates SSL certs for vm-proxy server
func GenerateCerts(host string) error {
	// Check if the cert files are available.
	err := httpscerts.Check("cert.pem", "key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", host)
		if err != nil {
			return errors.Wrap(err, "could not create https certs")
		}
	}
	return nil
}
