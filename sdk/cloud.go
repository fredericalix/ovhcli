package sdk

import (
	"fmt"
	"time"

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
	Region       string  `json:"region,omitempty"`
	Name         string  `json:"name,omitempty"`
	ID           string  `json:"id,omitempty"`
	OS           string  `json:"type,omitempty"`
	CreationDate string  `json:"creationDate,omitempty"`
	Status       string  `json:"status,omitempty"`
	MinDisk      int     `json:"minDisk,omitempty"`
	Visibility   string  `json:"visibility,omitempty"`
	Size         float32 `json:"size,omitempty"`
	MinRAM       int     `json:"minRam,omitempty"`
	User         string  `json:"user,omitempty"`
}

// Flavor is a go representation of Cloud Flavor
type Flavor struct {
	Region            string `json:"region,omitempty"`
	Name              string `json:"name,omitempty"`
	ID                string `json:"id,omitempty"`
	OS                string `json:"osType,omitempty"`
	Vcpus             int    `json:"vcpus,omitempty"`
	MemoryGB          int    `json:"ram,omitempty"`
	DiskSpaceGB       int    `json:"disk,omitempty"`
	Type              string `json:"type,omitempty"`
	InboundBandwidth  int    `json:"inboundBandwidth,omitempty"`
	OutboundBandwidth int    `json:"outboundBandwidth,omitempty"`
}

// SshkeyReq defines the fields for an SSH Key upload
type SshkeyReq struct {
	Name      string `json:"name,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
	Region    string `json:"region,omitempty"`
}

// Sshkey is a go representation of Cloud SSH Key
type Sshkey struct {
	Name        string `json:"name,omitempty"`
	ID          string `json:"id,omitempty"`
	PublicKey   string `json:"publicKey,omitempty"`
	Fingerprint string `json:"fingerPrint,omitempty"`
	//Regions     Regions `json:"regions"`
}

// Regions is a go representation of Cloud Regions
type Regions struct {
	Region             string `json:"region,omitempty"`
	Status             string `json:"status,omitempty"`
	ContinentCode      string `json:"continentCode,omitempty"`
	DatacenterLocation string `json:"datacenterLocation,omitempty"`
	Name               string `json:"name"`
	// Services      *string `json:"services"`
}

// Network is a go representation of a Cloud IP address
type Network struct {
	ID      string    `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Status  string    `json:"status,omitempty"`
	Type    string    `json:"type,omitempty"`
	VlanID  int       `json:"vlanId,omitempty"`
	Regions []Regions `json:"regions,omitempty"`
}

// IP is a go representation of a Cloud IP address
type IP struct {
	IP        string `json:"ip,omitempty"`
	NetworkID string `json:"networkId,omitempty"`
	Version   int    `json:"version,omitempty"`
	Type      string `json:"type,omitempty"`
}

// InstanceReq defines the fields for a VM creation
type InstanceReq struct {
	Name     string `json:"name,omitempty"`
	FlavorID string `json:"flavorID,omitempty"`
	ImageID  string `json:"imageID,omitempty"`
	Region   string `json:"region,omitempty"`
	SshkeyID string `json:"sshKeyID,omitempty"`
}

// Instance is a go representation of Cloud instance
type Instance struct {
	Name           string  `json:"name,omitempty"`
	ID             string  `json:"id,omitempty"`
	Status         string  `json:"status,omitempty"`
	Created        string  `json:"created,omitempty"`
	Region         string  `json:"regio,omitemptyn"`
	Image          *Image  `json:"image,omitempty"`
	Flavor         *Flavor `json:"flavor,omitempty"`
	Sshkey         *Sshkey `json:"sshKey,omitempty"`
	IPAddresses    []IP    `json:"ipAddresses,omitempty"`
	MonthlyBilling *string `json:"monthlyBilling,omitempty"`
}

// User is a go representation of Cloud user instance
type User struct {
	CreationDate time.Time `json:"creationDate"`
	Status       string    `json:"status"`
	ID           int       `json:"id"`
	Description  string    `json:"description"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
}

// RebootReq defines the fields for a VM reboot
type RebootReq struct {
	Type string `json:"type,omitempty"`
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

// CloudListRegions return a list of network regions
func (c *Client) CloudListRegions(projectID string) ([]Regions, error) {
	path := fmt.Sprintf("/cloud/project/%s/region", projectID)
	var resultsreq []string
	e := c.OVHClient.Get(path, &resultsreq)
	regions := []Regions{}
	for _, resultreq := range resultsreq {
		regions = append(regions, Regions{Region: resultreq})
	}
	return regions, e
}

// CloudInfoRegion return services status on a region
func (c *Client) CloudInfoRegion(projectID, regionName string) (*Regions, error) {
	region := &Regions{}
	path := fmt.Sprintf("/cloud/project/%s/region/%s", projectID, regionName)
	err := c.OVHClient.Get(path, region)
	return region, err
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
func (c *Client) CloudCreateInstance(projectID, name, pubkeyID, flavorID, imageID, region string) (instance *Instance, err error) {
	var instanceReq InstanceReq
	instanceReq.Name = name
	instanceReq.SshkeyID = pubkeyID
	instanceReq.FlavorID = flavorID
	instanceReq.ImageID = imageID
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

// CloudInfoNetworkPrivate return the list of a private network by given a project id
func (c *Client) CloudInfoNetworkPrivate(projectID string) ([]Network, error) {
	path := fmt.Sprintf("/cloud/project/%s/network/private", projectID)
	network := []Network{}

	e := c.OVHClient.Get(path, &network)

	return network, e
}

// CloudCreateNetworkPrivate create a private network in a vrack
func (c *Client) CloudCreateNetworkPrivate(projectID, name string, regions []Regions, vlanid int) (net *Network, err error) {
	var project Project
	project.ID = projectID
	var network Network
	network.Name = name
	network.VlanID = vlanid
	network.Regions = regions
	path := fmt.Sprintf("/cloud/project/%s/network/private", projectID)
	err = c.OVHClient.Post(path, network, &net)

	return net, err
}

// CloudProjectUsersList return the list of users by given a project id
func (c *Client) CloudProjectUsersList(projectID string) ([]User, error) {
	path := fmt.Sprintf("/cloud/project/%s/user", projectID)
	users := []User{}
	return users, c.OVHClient.Get(path, &users)
}

// CloudProjectUserCreate return the list of users by given a project id
func (c *Client) CloudProjectUserCreate(projectID, description string) (User, error) {
	path := fmt.Sprintf("/cloud/project/%s/user", projectID)
	data := map[string]string{
		"description": description,
	}
	user := User{}
	return user, c.OVHClient.Post(path, data, &user)
}

// CloudProjectRegionList return the list of region by given a project id
func (c *Client) CloudProjectRegionList(projectID string) ([]string, error) {
	path := fmt.Sprintf("/cloud/project/%s/region", projectID)
	regions := []string{}
	return regions, c.OVHClient.Get(path, &regions)
}
