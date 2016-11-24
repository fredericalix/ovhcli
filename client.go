package ovh

import (
	"fmt"

	govh "github.com/ovh/go-ovh/ovh"
)

var instance *Client

// Client ...
type Client struct {
	OVHClient *govh.Client
}

//NewClient initialize a client
func NewClient() (*Client, error) {
	if instance != nil {
		return instance, nil
	}

	c, err := govh.NewDefaultClient()
	if err != nil {
		return nil, fmt.Errorf("Error while creating OVH Client: %s\nYou need to create an application; please visite this page https://eu.api.ovh.com/createApp/ and create your $HOME/ovh.conf file\n\t[default]\n\t; general configuration: default endpoint\n\tendpoint=ovh-eu\n\n\t[ovh-eu]\n\t; configuration specific to 'ovh-eu' endpoint\n\tapplication_key=my_app_key", err)
	}
	instance := &Client{}
	instance.OVHClient = c

	return instance, nil
}
