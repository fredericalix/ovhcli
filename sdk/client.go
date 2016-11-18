package sdk

import (
	"fmt"

	"github.com/ovh/go-ovh/ovh"
)

var instance *Client

// Client ...
type Client struct {
	OVHClient *ovh.Client
}

//NewClient initialize a client
func NewClient() (*Client, error) {
	if instance != nil {
		return instance, nil
	}

	c, err := ovh.NewDefaultClient()
	if err != nil {
		return nil, fmt.Errorf("Error while creating OVH Client: %s", err)
	}
	instance.OVHClient = c

	return instance, nil
}
