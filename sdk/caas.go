package sdk

import "fmt"

// ContainersService is a representation of a Containers Service
type ContainersService struct {
	Cluster      string   `json:"cluster"`
	CreatedAt    string   `json:"createdAt"`
	Frameworks   []string `json:"frameworks"`
	LoadBalancer string   `json:"loadBalancer"`
	Metrics      struct {
		Resources struct {
			CPU int `json:"cpu"`
			Mem int `json:"mem"`
		} `json:"resources"`
		UsedResources struct {
			CPU float64 `json:"cpu"`
			Mem int     `json:"mem"`
		} `json:"usedResources"`
	} `json:"metrics"`
	Name      string   `json:"name"`
	Slaves    []string `json:"slaves"`
	State     string   `json:"state"`
	UpdatedAt string   `json:"updatedAt"`
}

// ContainersServicesList ...
func (c *Client) ContainersServicesList() ([]ContainersService, error) {
	containersservices := []ContainersService{}
	e := c.OVHClient.Get("/caas/containers", &containersservices)
	return containersservices, e
}

// ContainersServiceInfo ...
func (c *Client) ContainersServiceInfo(containerservid string) (*ContainersService, error) {
	containersservice := &ContainersService{}
	path := fmt.Sprintf("/caas/containers/%s", containerservid)
	e := c.OVHClient.Get(path, &containersservice)

	return containersservice, e
}
