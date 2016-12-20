package user

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var userID string

func init() {
	cmdInfo.PersistentFlags().StringVarP(&userID, "user-id", "", "", "User ID")
}

var cmdInfo = &cobra.Command{
	Use:   "info",
	Short: "Get User Info: ovhcli dbaas queue user info (--name=AppName | <--id=appID>) --user-id=userid",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			id = app.ID
		}

		user, err := client.DBaasQueueUserInfo(id, userID)
		common.Check(err)

		common.FormatOutputDef(user)
	},
}
