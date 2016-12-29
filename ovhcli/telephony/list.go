package telephony

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdServiceList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display telephony details")
}

var cmdServiceList = &cobra.Command{
	Use:   "list",
	Short: "List all telephony billing account: ovhcli telephony list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		services, err := client.TelephonyListBillingAccount(withDetails)
		common.Check(err)

		common.FormatOutputDef(services)
	},
}
