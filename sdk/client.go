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
