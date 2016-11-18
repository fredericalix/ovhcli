package sdk

import (
	"fmt"
)

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

// CloudProjectList ...
func (c *Client) CloudProjectList() (Projects, error) {
	projects := Projects{}
	e := c.OVHClient.Get("/cloud/project", &projects)
	return projects, e
}

// CloudProjectInfo ...
func (c *Client) CloudProjectInfo(projectid string) (*Project, error) {
	project := &Project{}
	path := fmt.Sprintf("/cloud/project/%s", projectid)
	e := c.OVHClient.Get(path, &project)

	return project, e
}
