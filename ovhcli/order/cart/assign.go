package cart

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var cmdCartAssign = &cobra.Command{
	Use:   "assign <cartID>",
	Short: "assign cart to connected user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		cartID := args[0]

		client, err := ovh.NewClient()
		common.Check(err)

		err = client.AssignCart(cartID)
		common.Check(err)
	},
}
