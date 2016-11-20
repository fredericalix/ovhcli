package sdk

import (
	"fmt"
)

// Domain ...
type Domain struct {
	// "Is whois obfuscation supported by this domain name's registry"
	OwoSupported bool `json:"owoSupported,omitempty"`

	// "Does the registry support ipv6 glue record"
	GlueRecordIpv6Supported bool `json:"glueRecordIpv6Supported,omitempty"`

	// "Transfer lock status"
	TransferLockStatus string `json:"transferLockStatus,omitempty"`
	//fullType: "domain.DomainLockStatusEnum"

	// "Domain's offer"
	Offer string `json:"offer,omitempty"`
	//fullType: "domain.OfferEnum"

	// "Contact Owner (you can edit it via /me/contact/<ID>)"
	WhoisOwner string `json:"whoisOwner,omitempty"`

	// "Is DNSSEC implemented for this domain name's tld"
	DnssecSupported bool `json:"dnssecSupported,omitempty"`

	// "Parent service"
	ParentService *string `json:"parentService,omitempty"`
	//fullType: "domain.ParentService"

	// "Domain name"
	Domain string `json:"domain"`

	// "Last update date"
	LastUpdate string `json:"lastUpdate,omitempty"`

	// "Does the registry support multi ip glue record"
	GlueRecordMultiIPSupported bool `json:"glueRecordMultiIpSupported,omitempty"`

	// "Name servers type"
	NameServerType string `json:"nameServerType,omitempty"`
	//fullType: "domain.DomainNsTypeEnum"
}

// DomainList ...
func (c *Client) DomainList() ([]Domain, error) {
	var names []string
	e := c.OVHClient.Get("/domain", &names)
	domains := []Domain{}
	for _, name := range names {
		domains = append(domains, Domain{Domain: name})
	}
	return domains, e
}

// DomainInfo ...
func (c *Client) DomainInfo(domainName string) (*Domain, error) {
	domain := &Domain{}
	err := c.OVHClient.Get(fmt.Sprintf("/domain/%s", domainName), domain)
	return domain, err
}
