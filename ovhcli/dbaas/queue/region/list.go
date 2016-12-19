package region

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdRegionList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display details")
}

var cmdRegionList = &cobra.Command{
	Use:   "list",
	Short: "List all regions on a service: ovhcli dbaas queue region list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			id = app.ID
		}

		apps, err := client.DBaasQueueRegionList(id, withDetails)
		common.Check(err)

		common.FormatOutputDef(apps)
	},
}
