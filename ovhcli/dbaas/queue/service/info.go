package service

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

func init() {

}

var cmdServiceServiceinfo = &cobra.Command{
	Use:   "serviceinfo",
	Short: "Get Service Serviceinfo: ovhcli dbaas queue service serviceinfo [--name=AppName] [--id=appID]]",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			id = app.ID
		}

		apps, err := client.DBaasQueueAppServiceInfo(id)
		common.Check(err)

		common.FormatOutputDef(apps)
	},
}
