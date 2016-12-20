package key

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var keyID string

func init() {
	cmdInfo.PersistentFlags().StringVarP(&keyID, "key-id", "", "", "Key ID")
}

var cmdInfo = &cobra.Command{
	Use:   "info",
	Short: "Get Key Info: ovhcli dbaas queue key info (--name=AppName | <--id=appID>) --key-id=keyid",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			id = app.ID
		}

		key, err := client.DBaasQueueKeyInfo(id, keyID)
		common.Check(err)

		common.FormatOutputDef(key)
	},
}
