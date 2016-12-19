package service

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdServiceList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display details")
}

var cmdServiceList = &cobra.Command{
	Use:   "list",
	Short: "List all services: ovhcli dbaas queue service list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		apps, err := client.DBaasQueueAppList(withDetails)
		common.Check(err)

		common.FormatOutputDef(apps)
	},
}
