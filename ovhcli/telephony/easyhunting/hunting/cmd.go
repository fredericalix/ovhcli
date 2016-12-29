package hunting

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"
	"github.com/admdwrf/ovhcli/ovhcli/telephony/easyhunting/hunting/agent"

	"github.com/spf13/cobra"
)

var serviceName string
var withDetails bool
var billingAccount string

func init() {
	Cmd.AddCommand(agent.Cmd)

	Cmd.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display telephony details")
	Cmd.PersistentFlags().StringVarP(&billingAccount, "billingAccount", "", "", "Billing Account")
	Cmd.PersistentFlags().StringVarP(&serviceName, "serviceName", "", "", "Service Name")
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "hunting",
	Short: "Hunting commands: ovhcli telephony easyhunting hunting [--help] [--billingAccount=<billingAccount>]  [--serviceName=<serviceName>]",
	Long:  `Hunting commands: ovhcli telephony easyhunting hunting <command>`,
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		services, err := client.TelephonyOvhPabxHunting(billingAccount, serviceName)
		common.Check(err)

		common.FormatOutputDef(services)
	},
}
