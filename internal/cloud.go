package internal

import (
	"fmt"

	"github.com/ovh/go-ovh/ovh"
)

var instance *ovh.Client

//Client return a new Ovh Client
func Client() *ovh.Client {

	if instance != nil {
		return instance
	}

	oc, err := ovh.NewDefaultClient()
	if err != nil {
		fmt.Printf("Error: %q\n", err)
	}
	return oc
}

// Projects is a list of project IDs
type Projects []string

// Project is a go representation of a Cloud project
type Project struct {
	Name         string `json:"description"`
	ID           string `json:"project_id"`
	Unleash      bool   `json:"unleash"`
	CreationDate string `json:"creationDate"`
	OrderID      int    `json:"orderID"`
	Status       string `json:"status"`
}
