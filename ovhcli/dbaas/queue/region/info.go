package region

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var regionID string

func init() {
	cmdInfo.PersistentFlags().StringVarP(&regionID, "region-id", "", "", "Region ID")
}

var cmdInfo = &cobra.Command{
	Use:   "info",
	Short: "Get Application Info: ovhcli dbaas queue region info (--name=AppName | <--id=appID>) --region-id=regionid",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			id = app.ID
		}

		region, err := client.DBaasQueueRegionInfo(id, regionID)
		common.Check(err)

		common.FormatOutputDef(region)
	},
}
