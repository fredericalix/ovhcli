package caas

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

var containerservid string

func init() {
	cmdContainersServiceInfo.PersistentFlags().StringVarP(&containerservid, "serviceName", "", "", "Containers Service Name")
}

var cmdContainersServiceInfo = &cobra.Command{
	Use:   "info",
	Short: "Info about a project: ovhcli caas info",
	Run: func(cmd *cobra.Command, args []string) {
		containersserviceinfo, err := internal.Client.ContainersServiceInfo(containerservid)
		internal.Check(err)
		internal.FormatOutputDef(containersserviceinfo)
	},
}
