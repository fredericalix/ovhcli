package sdk

import "fmt"

// Vrack ...
type Vrack struct {

	// "Vrack name"
	Vrack string `json:"name"`
}

// VrackList ...
func (c *Client) VrackList() ([]Vrack, error) {
	var ids []string
	e := c.OVHClient.Get("/vrack", &ids)
	vracks := []Vrack{}
	for _, id := range ids {
		vracks = append(vracks, Vrack{Vrack: id})
	}
	return vracks, e
}

// VrackInfo ...
func (c *Client) VrackInfo(vrackName string) (*Vrack, error) {
	vrack := &Vrack{}
	err := c.OVHClient.Get(fmt.Sprintf("/vrack/%s", vrackName), vrack)
	return vrack, err
}
