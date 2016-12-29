package easyhunting

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var serviceName string

func init() {
	cmdEasyHuntingInfo.PersistentFlags().StringVarP(&billingAccount, "billingAccount", "", "", "Billing Account")
	cmdEasyHuntingInfo.PersistentFlags().StringVarP(&serviceName, "serviceName", "", "", "Service Name")
}

var cmdEasyHuntingInfo = &cobra.Command{
	Use:   "info <easyhunting>",
	Short: "Get info on a easyhunting: ovhcli telephony easyhunting info --billingAccount=aa --serviceName=bb",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewClient()
		common.Check(err)

		d, err := client.TelephonyEasyHuntingInfo(billingAccount, serviceName)
		common.Check(err)
		common.FormatOutputDef(d)
	},
}
