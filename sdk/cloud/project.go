package cloud

import (
	"fmt"

	"github.com/admdwrf/ovhcli/sdk"
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

func ProjectList() (Projects, error) {
	c, err := sdk.Client()
	if err != nil {
		return nil, err
	}
	projects := Projects{}
	e := c.Get("/cloud/project", &projects)
	return projects, e
}

func ProjectInfo() (Projects, error) {
	var projectid string
	c, err := sdk.Client()
	if err != nil {
		return nil, err
	}
	project := Project{}
	path := fmt.Sprintf("/cloud/project/%s", projectid)
	e := c.Get(path, &project)

	return project, e

	/*		c := internal.Client()
			if c == nil {
				os.Exit(1)
			}
			project := internal.Project{}
			path := fmt.Sprintf("/cloud/project/%s", projectid)
			if err := c.Get(path, &project); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Project Name: %s\n", project.Name)
			fmt.Printf("Project Status: %s\n", project.Status)
			fmt.Printf("Creation Date: %s\n", project.CreationDate)
	*/
}
