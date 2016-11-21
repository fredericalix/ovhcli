package sdk

import (
	"fmt"

	"github.com/ovh/go-ovh/ovh"
)

const (
	// CustomerInterface is the URL of the customer interface, for error messages
	CustomerInterface = "https://www.ovh.com/manager/cloud/index.html"
)

// Project is a go representation of a Cloud project
type Project struct {
	Name         string `json:"description,omitempty"`
	ID           string `json:"project_id"`
	Unleash      bool   `json:"unleash,omitempty"`
	CreationDate string `json:"creationDate,omitempty"`
	OrderID      int    `json:"orderID,omitempty"`
	Status       string `json:"status,omitempty"`
}

// Image is a go representation of a Cloud Image (VM template)
type Image struct {
	Region       string  `json:"region"`
	Name         string  `json:"name"`
	ID           string  `json:"id"`
	OS           string  `json:"type"`
	CreationDate string  `json:"creationDate"`
	Status       string  `json:"status"`
	MinDisk      int     `json:"minDisk"`
	Visibility   string  `json:"visibility"`
	Size         float32 `json:"size"`
	MinRAM       int     `json:"minRam"`
	User         string  `json:"user"`
}

// Flavor is a go representation of Cloud Flavor
type Flavor struct {
	Region            string `json:"region"`
	Name              string `json:"name"`
	ID                string `json:"id"`
	OS                string `json:"osType"`
	Vcpus             int    `json:"vcpus"`
	MemoryGB          int    `json:"ram"`
	DiskSpaceGB       int    `json:"disk"`
	Type              string `json:"type"`
	InboundBandwidth  int    `json:"inboundBandwidth"`
	OutboundBandwidth int    `json:"outboundBandwidth"`
}

// SshkeyReq defines the fields for an SSH Key upload
type SshkeyReq struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
	Region    string `json:"region,omitempty"`
}

// Sshkey is a go representation of Cloud SSH Key
type Sshkey struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	PublicKey   string `json:"publicKey"`
	Fingerprint string `json:"fingerPrint"`
	// TODO Regions     Regions `json:"region"`
}

// Network is a go representation of a Cloud IP address
type Network struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
	VlanID int    `json:"vlanId"`
	//Regions string `json:"regions"`
}

// IP is a go representation of a Cloud IP address
type IP struct {
	IP        string `json:"ip"`
	NetworkID string `json:"networkId"`
	Version   int    `json:"version"`
	Type      string `json:"type"`
}

// InstanceReq defines the fields for a VM creation
type InstanceReq struct {
	Name     string `json:"name"`
	FlavorID string `json:"flavorID"`
	ImageID  string `json:"imageID"`
	Region   string `json:"region"`
	SshkeyID string `json:"sshKeyID"`
}

// Instance is a go representation of Cloud instance
type Instance struct {
	Name           string  `json:"name"`
	ID             string  `json:"id"`
	Status         string  `json:"status"`
	Created        string  `json:"created"`
	Region         string  `json:"region"`
	Image          *Image  `json:"image"`
	Flavor         *Flavor `json:"flavor"`
	Sshkey         *Sshkey `json:"sshKey"`
	IPAddresses    []IP    `json:"ipAddresses"`
	MonthlyBilling *string `json:"monthlyBilling"`
}

// RebootReq defines the fields for a VM reboot
type RebootReq struct {
	Type string `json:"type"`
}

// CloudProjectsList returns a list of string project ID
func (c *Client) CloudProjectsList() ([]Project, error) {
	projects := []Project{}
	ids := []string{}
	e := c.OVHClient.Get("/cloud/project", &ids)
	if e != nil {
		return nil, e
	}
	for _, id := range ids {
		projects = append(projects, Project{ID: id})
	}
	return projects, e
}

// CloudProjectInfoByID return the details of a project given a project id
func (c *Client) CloudProjectInfoByID(projectID string) (*Project, error) {
	project := &Project{}
	path := fmt.Sprintf("/cloud/project/%s", projectID)
	e := c.OVHClient.Get(path, &project)

	return project, e
}

// CloudProjectInfoByName returns the details of a project given its name.
func (c *Client) CloudProjectInfoByName(projectName string) (project *Project, err error) {
	// get project list
	projects, err := c.CloudProjectsList()
	if err != nil {
		return nil, err
	}

	// If projectName is a valid projectID return it.
	for _, p := range projects {
		if p.ID == projectName {
			return c.CloudProjectInfoByID(p.ID)
		}
	}

	// Attempt to find a project matching projectName. This is potentially slow
	for _, p := range projects {
		project, err := c.CloudProjectInfoByID(p.ID)
		if err != nil {
			return nil, err
		}

		if project.Name == projectName {
			return project, nil
		}
	}

	// Ooops
	return nil, fmt.Errorf("Project '%s' does not exist on OVH cloud. To create or rename a project, please visit %s", projectName, CustomerInterface)
}

// CloudGetImages returns a list of images for a given project in a given region
func (c *Client) CloudGetImages(projectID, region string) ([]Image, error) {
	path := fmt.Sprintf("/cloud/project/%s/image?osType=linux&region=%s", projectID, region)
	images := []Image{}
	err := c.OVHClient.Get(path, &images)
	return images, err
}

// CloudGetInstance finds a VM instance given a name or an ID
func (c *Client) CloudGetInstance(projectID, instanceID string) (instance *Instance, err error) {
	path := fmt.Sprintf("/cloud/project/%s/instance/%s", projectID, instanceID)
	err = c.OVHClient.Get(path, &instance)
	return instance, nil
}

// CloudCreateInstance start a new public cloud instance and returns resulting object
func (c *Client) CloudCreateInstance(projectID, name, pubkeyID, flavorID, ImageID, region string) (instance *Instance, err error) {
	var instanceReq InstanceReq
	instanceReq.Name = name
	instanceReq.SshkeyID = pubkeyID
	instanceReq.FlavorID = flavorID
	instanceReq.ImageID = ImageID
	instanceReq.Region = region

	path := fmt.Sprintf("/cloud/project/%s/instance", projectID)
	err = c.OVHClient.Post(path, instanceReq, &instance)
	return instance, err
}

// CloudDeleteInstance stops and destroys a public cloud instance
func (c *Client) CloudDeleteInstance(projectID, instanceID string) (err error) {
	path := fmt.Sprintf("/cloud/project/%s/instance/%s", projectID, instanceID)
	err = c.OVHClient.Delete(path, nil)
	if apierror, ok := err.(*ovh.APIError); ok && apierror.Code == 404 {
		err = nil
	}
	return err
}

// CloudListInstance show cloud instance(s)
func (c *Client) CloudListInstance(projectID string) ([]Instance, error) {
	path := fmt.Sprintf("/cloud/project/%s/instance", projectID)
	instances := []Instance{}
	e := c.OVHClient.Get(path, &instances)

	return instances, e
}

// CloudInfoInstance give info about cloud instance
func (c *Client) CloudInfoInstance(projectID, instanceID string) (*Instance, error) {
	path := fmt.Sprintf("/cloud/project/%s/instance/%s", projectID, instanceID)
	instances := &Instance{}

	e := c.OVHClient.Get(path, &instances)

	return instances, e
}

// CloudInfoNetworkPublic return the list of a public network by given a project id
func (c *Client) CloudInfoNetworkPublic(projectID string) ([]Network, error) {
	path := fmt.Sprintf("/cloud/project/%s/network/public", projectID)
	network := []Network{}

	e := c.OVHClient.Get(path, &network)

	return network, e
}
