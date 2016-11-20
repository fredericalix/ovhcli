package domain

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/admdwrf/ovhcli/sdk"
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

		domains, err := internal.Client.DomainList()
		internal.Check(err)

		if withDetails {
			domainsComplete := []sdk.Domain{}
			for _, domain := range domains {
				d, err := internal.Client.DomainInfo(domain.Domain)
				internal.Check(err)
				domainsComplete = append(domainsComplete, *d)
			}
			domains = domainsComplete
		}

		internal.FormatOutputDef(domains)
	},
}
