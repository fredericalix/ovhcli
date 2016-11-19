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

// Image is a go representation of a Cloud Image (VM template)
type Image struct {
	Region       string `json:"region"`
	Name         string `json:"name"`
	ID           string `json:"id"`
	OS           string `json:"type"`
	CreationDate string `json:"creationDate"`
	Status       string `json:"status"`
	MinDisk      int    `json:"minDisk"`
	Visibility   string `json:"visibility"`
}

// Images is a list of Images
type Images []Image

// Flavor is a go representation of Cloud Flavor
type Flavor struct {
	Region      string `json:"region"`
	Name        string `json:"name"`
	ID          string `json:"id"`
	OS          string `json:"osType"`
	Vcpus       int    `json:"vcpus"`
	MemoryGB    int    `json:"ram"`
	DiskSpaceGB int    `json:"disk"`
	Type        string `json:"type"`
}

// Flavors is a list flavors
type Flavors []Flavor

// Regions is a list of Cloud Region names
type Regions []string

// SshkeyReq defines the fields for an SSH Key upload
type SshkeyReq struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
	Region    string `json:"region,omitempty"`
}

// Sshkey is a go representation of Cloud SSH Key
type Sshkey struct {
	Name        string  `json:"name"`
	ID          string  `json:"id"`
	PublicKey   string  `json:"publicKey"`
	Fingerprint string  `json:"fingerPrint"`
	Regions     Regions `json:"region"`
}

// Sshkeys is a list of Sshkey
type Sshkeys []Sshkey

// IP is a go representation of a Cloud IP address
type IP struct {
	IP   string `json:"ip"`
	Type string `json:"type"`
}

// IPs is a list of IPs
type IPs []IP

// InstanceReq defines the fields for a VM creation
type InstanceReq struct {
	Name           string `json:"name"`
	FlavorID       string `json:"flavorID"`
	ImageID        string `json:"imageID"`
	Region         string `json:"region"`
	SshkeyID       string `json:"sshKeyID"`
	MonthlyBilling bool   `json:"monthlyBilling"`
}

// Instance is a go representation of Cloud instance
type Instance struct {
	Name           string `json:"name"`
	ID             string `json:"id"`
	Status         string `json:"status"`
	Created        string `json:"created"`
	Region         string `json:"region"`
	Image          Image  `json:"image"`
	Flavor         Flavor `json:"flavor"`
	Sshkey         Sshkey `json:"sshKey"`
	IPAddresses    IPs    `json:"ipAddresses"`
	MonthlyBilling bool   `json:"monthlyBilling"`
}

// RebootReq defines the fields for a VM reboot
type RebootReq struct {
	Type string `json:"type"`
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

// CloudGetImages returns a list of images for a given project in a given region
func (c *Client) CloudGetImages(projectid, region string) (images Images, err error) {
	url := fmt.Sprintf("/cloud/project/%s/image?osType=linux&region=%s", projectid, region)
	err = c.OVHClient.Get(url, &images)
	return images, err
}
