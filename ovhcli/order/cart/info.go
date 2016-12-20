package cart

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var cmdCartInfo = &cobra.Command{
	Use:   "info <cartID>",
	Short: "Retrieve cart info",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		cartID := args[0]

		client, err := ovh.NewClient()
		common.Check(err)

		d, err := client.CartInfo(cartID)
		common.Check(err)
		common.FormatOutputDef(d)
	},
}
