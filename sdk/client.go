package sdk

import "github.com/ovh/go-ovh/ovh"

var instance *ovh.Client

//Client return a new Ovh Client
func Client() (*ovh.Client, error) {

	if instance != nil {
		return instance, nil
	}

	return ovh.NewDefaultClient()
}
