package sdk

import "fmt"

// Vrack ...
type Vrack struct {

	// "Vrack name"
	Name string `json:"name"`

	// "Vrack decription"
	Description string `json:"description"`
}

// Ids of vracks
type Ids []string

// VrackList ...
func (c *Client) VrackList() ([]Vrack, error) {
	ids := Ids{}
	e := c.OVHClient.Get("/vrack", &ids)
	vracks := []Vrack{}
	for _, id := range ids {
		vracks = append(vracks, Vrack{Name: id})

	}
	return vracks, e
}

// VrackInfo ...
func (c *Client) VrackInfo(vrackName string) (*Vrack, error) {
	vrack := &Vrack{}
	err := c.OVHClient.Get(fmt.Sprintf("/vrack/%s", vrackName), vrack)
	return vrack, err
}
