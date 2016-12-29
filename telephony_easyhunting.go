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

// TelephonyEasyHuntingList list all OVH easy calls queues associated with this billing account
func (c *Client) TelephonyEasyHuntingList(billingAccount string, withDetails bool) ([]TelephonyEasyHunting, error) {
	var names []string
	if err := c.OVHClient.Get(fmt.Sprintf("/telephony/%s/easyHunting/", billingAccount), &names); err != nil {
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
func (c *Client) TelephonyEasyHuntingInfo(billingAccount, serviceName string) (*TelephonyEasyHunting, error) {
	telephonyEasyHunting := &TelephonyEasyHunting{}
	err := c.OVHClient.Get(fmt.Sprintf("/telephony/%s/easyHunting/%s", billingAccount, serviceName), telephonyEasyHunting)
	return telephonyEasyHunting, err
}
