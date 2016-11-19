package domain

import (
	"github.com/admdwrf/ovhcli/internal"
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
			for _, domain := range domains {
				err := internal.Client.DomainInfo(domain)
				internal.Check(err)
			}
		}

		internal.FormatOutputDef(domains)
	},
}
