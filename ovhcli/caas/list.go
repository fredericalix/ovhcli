package caas

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var cmdContainersServicesList = &cobra.Command{
	Use:   "list",
	Short: "List all containers services: ovhcli caas list",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewClient()
		common.Check(err)

		containersservices, err := client.ContainersServicesList()
		common.Check(err)

		common.FormatOutputDef(containersservices)
	},
}
