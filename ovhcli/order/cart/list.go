package cart

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdCartList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display order cart details")
}

var cmdCartList = &cobra.Command{
	Use:   "list",
	Short: "List all carts: ovhcli cart list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		carts, err := client.CartList()
		common.Check(err)

		if withDetails {
			carts = getDetailledCartsList(client, carts)
		}

		common.FormatOutputDef(carts)
	},
}

func getDetailledCartsList(client *ovh.Client, carts []ovh.Cart) []ovh.Cart {

	cartsChan, errChan := make(chan ovh.Cart), make(chan error)
	for _, cart := range carts {
		go func(cart ovh.Cart) {
			c, err := client.CartInfo(cart.CartID)
			if err != nil {
				errChan <- err
				return
			}
			cartsChan <- *c
		}(cart)
	}

	cartsComplete := []ovh.Cart{}

	for i := 0; i < len(carts); i++ {
		select {
		case cart := <-cartsChan:
			cartsComplete = append(cartsComplete, cart)
		case err := <-errChan:
			common.Check(err)
		}
	}

	return cartsComplete
}
