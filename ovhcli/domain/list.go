package domain

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdDomainList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display domain details")
}

var cmdDomainList = &cobra.Command{
	Use:   "list",
	Short: "List all domains: ovhcli domain list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		domains, err := client.DomainList()
		common.Check(err)

		if withDetails {
			domains = getDetailledDomainList(client, domains)
		}

		common.FormatOutputDef(domains)
	},
}

func getDetailledDomainList(client *ovh.Client, domains []ovh.Domain) []ovh.Domain {

	domainsChan, errChan := make(chan ovh.Domain), make(chan error)
	for _, domain := range domains {
		go func(domain ovh.Domain) {
			d, err := client.DomainInfo(domain.Domain)
			if err != nil {
				errChan <- err
				return
			}
			domainsChan <- *d
		}(domain)
	}

	domainsComplete := []ovh.Domain{}

	for i := 0; i < len(domains); i++ {
		select {
		case domains := <-domainsChan:
			domainsComplete = append(domainsComplete, domains)
		case err := <-errChan:
			common.Check(err)
		}
	}

	return domainsComplete
}
