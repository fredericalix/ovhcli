package caas

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

var cmdContainersServicesList = &cobra.Command{
	Use:   "list",
	Short: "List all containers services: ovhcli caas list",
	Run: func(cmd *cobra.Command, args []string) {
		containersservices, err := internal.Client.ContainersServicesList()
		internal.Check(err)
		internal.FormatOutputDef(containersservices)
	},
}
