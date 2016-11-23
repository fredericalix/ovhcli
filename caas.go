package ovh

import "fmt"

// ContainersService is a representation of a Containers Service
type ContainersService struct {
	Cluster      string   `json:"cluster,omitempty"`
	CreatedAt    string   `json:"createdAt,omitempty"`
	Frameworks   []string `json:"frameworks,omitempty"`
	LoadBalancer string   `json:"loadBalancer,omitempty"`
	Metrics      struct {
		Resources struct {
			CPU int `json:"cpu,omitempty"`
			Mem int `json:"mem,omitempty"`
		} `json:"resources,omitempty"`
		UsedResources struct {
			CPU float64 `json:"cpu,omitempty"`
			Mem int     `json:"mem,omitempty"`
		} `json:"usedResources,omitempty"`
	} `json:"metrics,omitempty"`
	Name      string   `json:"name,omitempty"`
	Slaves    []string `json:"slaves,omitempty"`
	State     string   `json:"state,omitempty"`
	UpdatedAt string   `json:"updatedAt,omitempty"`
}

// ContainersServicesList ...
func (c *Client) ContainersServicesList() ([]string, error) {
	path := fmt.Sprintf("/caas/containers")
	containersservices := []string{}

	e := c.OVHClient.Get(path, &containersservices)

	return containersservices, e
}

// ContainersServiceInfo ...
func (c *Client) ContainersServiceInfo(containerservid string) (*ContainersService, error) {
	containersservice := &ContainersService{}
	path := fmt.Sprintf("/caas/containers/%s", containerservid)
	e := c.OVHClient.Get(path, &containersservice)

	return containersservice, e
}
