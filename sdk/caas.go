package sdk

import "fmt"

// ContainersServices is a list of available services
type ContainersServices []string

// ContainersService is a go representation of a Containers Service
type ContainersService struct {
	Cluster      string   `json:"cluster"`
	Name         string   `json:"name"`
	LoadBalancer string   `json:"loadBalancer"`
	CreationDate string   `json:"createdAt"`
	Slave        []string `json:"slaves"`
}

// ContainersServicesList ...
func (c *Client) ContainersServicesList() (ContainersServices, error) {
	containersservices := ContainersServices{}
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
