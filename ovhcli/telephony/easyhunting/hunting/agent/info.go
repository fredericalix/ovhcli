package agent

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var agentID float64

func init() {
	cmdEasyHuntingAgentInfo.PersistentFlags().StringVarP(&billingAccount, "billingAccount", "", "", "Billing Account")
	cmdEasyHuntingAgentInfo.PersistentFlags().StringVarP(&serviceName, "serviceName", "", "", "Service Name")
	cmdEasyHuntingAgentInfo.PersistentFlags().Float64VarP(&agentID, "agentID", "", 0.0, "Agent ID")
}

var cmdEasyHuntingAgentInfo = &cobra.Command{
	Use:   "info <easyhunting>",
	Short: "Get info on a easyhunting: ovhcli telephony easyhunting hunting agent info --billingAccount=aa --serviceName=bb --agentID=cc",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewClient()
		common.Check(err)

		d, err := client.TelephonyOvhPabxHuntingAgentInfo(billingAccount, serviceName, agentID)
		common.Check(err)
		common.FormatOutputDef(d)
	},
}