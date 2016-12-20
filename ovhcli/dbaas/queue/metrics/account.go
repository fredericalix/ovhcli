package metrics

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

func init() {
}

var cmdAccount = &cobra.Command{
	Use:   "account",
	Short: "Get metrics account: ovhcli dbaas queue metrics account (--name=AppName | <--id=appID>)",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			id = app.ID
		}

		apps, err := client.DBaasQueueMetricsAccount(id)
		common.Check(err)

		common.FormatOutputDef(apps)
	},
}
