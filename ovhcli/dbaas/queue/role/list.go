package role

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdRoleList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display details")
}

var cmdRoleList = &cobra.Command{
	Use:   "list",
	Short: "List all roles on a service: ovhcli dbaas queue role list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			id = app.ID
		}

		apps, err := client.DBaasQueueRoleList(id, withDetails)
		common.Check(err)

		common.FormatOutputDef(apps)
	},
}
