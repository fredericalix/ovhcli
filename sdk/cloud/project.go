package cloud

import "github.com/admdwrf/ovhcli/sdk"

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

func ProjectList() (Projects, error) {
	c, err := sdk.Client()
	if err != nil {
		return nil, err
	}
	projects := Projects{}
	e := c.Get("/cloud/project", &projects)
	return projects, e
}
