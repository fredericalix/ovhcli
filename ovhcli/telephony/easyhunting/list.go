package easyhunting

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool
var billingAccount string

func init() {
	cmdEasyHuntingList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display telephony details")
	cmdEasyHuntingList.PersistentFlags().StringVarP(&billingAccount, "billingAccount", "", "", "Billing Account")
}

var cmdEasyHuntingList = &cobra.Command{
	Use:   "list",
	Short: "List all telephony billing account: ovhcli telephony easyhunting list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		services, err := client.TelephonyEasyHuntingList(billingAccount, withDetails)
		common.Check(err)

		common.FormatOutputDef(services)
	},
}
