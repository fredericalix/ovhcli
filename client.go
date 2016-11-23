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
		return nil, fmt.Errorf("Error while creating OVH Client: %s", err)
	}
	instance := &Client{}
	instance.OVHClient = c

	return instance, nil
}
