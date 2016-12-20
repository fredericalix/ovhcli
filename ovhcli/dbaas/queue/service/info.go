package service

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var cmdInfo = &cobra.Command{
	Use:   "info",
	Short: "Get Application Info: ovhcli dbaas queue service info (--name=AppName | <--id=appID>)",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			common.FormatOutputDef(app)
		} else {
			app, errInfo := client.DBaasQueueAppInfo(id)
			common.Check(errInfo)
			common.FormatOutputDef(app)
		}

	},
}
