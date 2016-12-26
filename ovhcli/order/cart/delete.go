package cart

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var cmdCartDelete = &cobra.Command{
	Use:   "delete <cartID>",
	Short: "Delete cart",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		cartID := args[0]

		client, err := ovh.NewClient()
		common.Check(err)

		err = client.OrderDeleteCart(cartID)
		common.Check(err)
	},
}
