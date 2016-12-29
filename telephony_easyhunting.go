package ovh

import (
	"fmt"
)

// TelephonyEasyHunting struct
type TelephonyEasyHunting struct {
	//Max wait time when caller is in queue (in seconds)
	MaxWaitTime float64 `json:"maxWaitTime"`

	// FeatureType
	FeatureType string `json:"featureType"`

	// Strategy : The calls dispatching strategy
	Strategy string `json:"strategy"`

	// QueueSize Max number of callers in queue
	QueueSize float64 `json:"queueSize"`

	// ToneOnHold: Tone played when caller is put on hold
	ToneOnHold float64 `json:"toneOnHold"`

	// ServiceName containers service Name
	ServiceName string `json:"serviceName"`

	// ShowCallerNumber: The presented number when bridging calls
	ShowCallerNumber string `json:"showCallerNumber"`

	// Description ...
	Description string `json:"description"`

	// AnonymousRejection: Reject (hangup) anonymous calls
	AnonymousRejection bool `json:"anonymousRejection"`

	//ToneOnOpening: Tone played when call is picked up
	ToneOnOpening float64 `json:"toneOnOpening"`

	// serviceType
	ServiceType string `json:"serviceType"`

	// Voicemail: The voicemail used by the EasyPABX
	Voicemail string `json:"voicemail"`

	//ToneOnClosing: Tone played just before call is hang up
	ToneOnClosing float64 `json:"toneOnClosing"`
}

// TelephonyOvhPabxHunting struct
type TelephonyOvhPabxHunting struct {
	// The templated url of your CRM, opened by the banner application of your cloudpabx
	CrmUrlTemplate string `json:"crmUrlTemplate"`
	// The name of your callcenter offer
	Name string `json:"name"`
	// Enable G729 codec on your callcenter
	G729 bool `json:"g729"`
}

// TelephonyOvhPabxHuntingAgent ...
type TelephonyOvhPabxHuntingAgent struct {
	// ID of agent
	AgentId float64 `json:"agentId"`
	// The wrap up time (in seconds) after the calls
	WrapUpTime float64 `json:"wrapUpTime"`
	// The number of the agent
	Number string `json:"number"`
	// The waiting timeout (in seconds) before hangup an assigned called
	Timeout float64 `json:"timeout"`
	// The current status of the agent
	Status string `json:"status"`
	// The maximum of simultaneous calls that the agent will receive from the hunting
	SimultaneousLines float64 `json:"simultaneousLines"`
	// The id of the current break status of the agent
	BreakStatus float64 `json:"breakStatus"`
}

// TelephonyEasyHuntingList list all OVH easy calls queues associated with this billing account
// GET /telephony/{billingAccount}/easyHunting
func (c *Client) TelephonyEasyHuntingList(billingAccount string, withDetails bool) ([]TelephonyEasyHunting, error) {
	var names []string
	if err := c.OVHClient.Get(fmt.Sprintf("/telephony/%s/easyHunting", billingAccount), &names); err != nil {
		return nil, err
	}

	services := []TelephonyEasyHunting{}
	for _, name := range names {
		services = append(services, TelephonyEasyHunting{ServiceName: name})
	}

	if !withDetails {
		return services, nil
	}

	servicesChan, errChan := make(chan TelephonyEasyHunting), make(chan error)
	for _, telephonyEasyHunting := range services {
		go func(billingAccount, serviceName string) {
			d, err := c.TelephonyEasyHuntingInfo(billingAccount, serviceName)
			if err != nil {
				errChan <- err
				return
			}
			servicesChan <- *d
		}(billingAccount, telephonyEasyHunting.ServiceName)
	}

	servicesComplete := []TelephonyEasyHunting{}

	for i := 0; i < len(services); i++ {
		select {
		case services := <-servicesChan:
			servicesComplete = append(servicesComplete, services)
		case err := <-errChan:
			return nil, err
		}
	}

	return servicesComplete, nil
}

// TelephonyEasyHuntingInfo retrieve all infos of one easy hunting service
// GET /telephony/{billingAccount}/easyHunting/{serviceName}
func (c *Client) TelephonyEasyHuntingInfo(billingAccount, serviceName string) (*TelephonyEasyHunting, error) {
	telephonyEasyHunting := &TelephonyEasyHunting{}
	err := c.OVHClient.Get(fmt.Sprintf("/telephony/%s/easyHunting/%s", billingAccount, serviceName), telephonyEasyHunting)
	return telephonyEasyHunting, err
}

// TelephonyOvhPabxHuntingList list all OVH Pabx Hunting
// GET /telephony/{billingAccount}/easyHunting/{serviceName}/hunting
func (c *Client) TelephonyOvhPabxHunting(billingAccount, serviceName string) (*TelephonyOvhPabxHunting, error) {
	telephonyOvhPabxHunting := &TelephonyOvhPabxHunting{}
	err := c.OVHClient.Get(fmt.Sprintf("/telephony/%s/easyHunting/%s/hunting", billingAccount, serviceName), telephonyOvhPabxHunting)
	return telephonyOvhPabxHunting, err
}

// TelephonyOvhPabxHuntingAgentList list all OVH easy calls queues associated with this billing account
// GET  /telephony/{billingAccount}/easyHunting/{serviceName}/hunting/agent
func (c *Client) TelephonyOvhPabxHuntingAgentList(billingAccount, serviceName string, withDetails bool) ([]TelephonyOvhPabxHuntingAgent, error) {
	var names []float64
	fmt.Printf("HOp")
	if err := c.OVHClient.Get(fmt.Sprintf("/telephony/%s/easyHunting/%s/hunting/agent", billingAccount, serviceName), &names); err != nil {
		return nil, err
	}

	agents := []TelephonyOvhPabxHuntingAgent{}
	for _, agentID := range names {
		agents = append(agents, TelephonyOvhPabxHuntingAgent{AgentId: agentID})
	}

	if !withDetails {
		return agents, nil
	}

	agentsChan, errChan := make(chan TelephonyOvhPabxHuntingAgent), make(chan error)
	for _, agent := range agents {
		go func(billingAccount, serviceName string, agentID float64) {
			d, err := c.TelephonyOvhPabxHuntingAgentInfo(billingAccount, serviceName, agentID)
			if err != nil {
				errChan <- err
				return
			}
			agentsChan <- *d
		}(billingAccount, serviceName, agent.AgentId)
	}

	agentsComplete := []TelephonyOvhPabxHuntingAgent{}

	for i := 0; i < len(agents); i++ {
		select {
		case agents := <-agentsChan:
			agentsComplete = append(agentsComplete, agents)
		case err := <-errChan:
			return nil, err
		}
	}

	return agentsComplete, nil
}

// TelephonyOvhPabxHuntingAgent list all OVH Pabx Hunting Agent
// GET /telephony/{billingAccount}/easyHunting/{serviceName}/hunting/agent
func (c *Client) TelephonyOvhPabxHuntingAgentInfo(billingAccount, serviceName string, agentID float64) (*TelephonyOvhPabxHuntingAgent, error) {
	telephonyOvhPabxHuntingAgent := &TelephonyOvhPabxHuntingAgent{}
	err := c.OVHClient.Get(fmt.Sprintf("/telephony/%s/easyHunting/%s/hunting/agent/%f", billingAccount, serviceName, agentID), telephonyOvhPabxHuntingAgent)
	return telephonyOvhPabxHuntingAgent, err
}
