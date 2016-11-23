package domain

import (
	"sync"

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
	var wg sync.WaitGroup
	wg.Add(len(domains))

	domainsChan := make(chan ovh.Domain)
	domainsComplete := []ovh.Domain{}
	for _, domain := range domains {
		go func(domain ovh.Domain) {
			d, err := client.DomainInfo(domain.Domain)
			common.Check(err)
			domainsChan <- *d
		}(domain)
	}
	go func(wg *sync.WaitGroup) {
		for d := range domainsChan {
			domainsComplete = append(domainsComplete, d)
			wg.Done()
		}
	}(&wg)
	wg.Wait()
	return domainsComplete
}
