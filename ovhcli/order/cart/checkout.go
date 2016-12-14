package cart

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var cmdCartCheckoutGet = &cobra.Command{
	Use:   "GetCheckout <cartID>",
	Short: "get checkout cart : ovhcli order cart getCheckout <cartID>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		cartID := args[0]
		client, err := ovh.NewClient()
		common.Check(err)

		d, err := client.OrderGetCheckoutCart(cartID)
		common.Check(err)
		common.FormatOutputDef(d)
	},
}

var cmdCartCheckoutPost = &cobra.Command{
	Use:   "postCheckout <cartID>",
	Short: "post checkout cart : ovhcli order cart postCheckout <cartID>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			common.WrongUsage(cmd)
		}
		cartID := args[0]
		client, err := ovh.NewClient()
		common.Check(err)

		d, err := client.OrderPostCheckoutCart(cartID, true)
		common.Check(err)
		common.FormatOutputDef(d)
	},
}
