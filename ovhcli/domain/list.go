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
			domainsComplete := []ovh.Domain{}
			for _, domain := range domains {
				d, err := client.DomainInfo(domain.Domain)
				common.Check(err)
				domainsComplete = append(domainsComplete, *d)
			}
			domains = domainsComplete
		}

		common.FormatOutputDef(domains)
	},
}
