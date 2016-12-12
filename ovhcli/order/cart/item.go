package cart

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var itemsWithDetails bool
var cartID string

func init() {
	CmdCartListItems.PersistentFlags().StringVarP(&cartID, "cartID", "", "", "id of your cart")

	CmdCartListItems.PersistentFlags().BoolVarP(&itemsWithDetails, "withDetails", "", false, "Display domain details")
}

//CmdCartListItems list all item of a cart
var CmdCartListItems = &cobra.Command{
	Use:   "listItems",
	Short: "List all items of a cart: ovhcli order cart list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		items, err := client.OrderCartItemList(cartID)
		common.Check(err)

		if itemsWithDetails {
			items = getDetailledItemList(client, items)
		}

		common.FormatOutputDef(items)
	},
}

func getDetailledItemList(client *ovh.Client, items []ovh.OrderCartItem) []ovh.OrderCartItem {

	itemsChan, errChan := make(chan ovh.OrderCartItem), make(chan error)
	for _, item := range items {
		go func(item ovh.OrderCartItem) {
			i, err := client.OrderCartItemInfo(cartID, item.ItemID)
			if err != nil {
				errChan <- err
				return
			}
			itemsChan <- *i
		}(item)
	}

	itemsComplete := []ovh.OrderCartItem{}

	for i := 0; i < len(items); i++ {
		select {
		case item := <-itemsChan:
			itemsComplete = append(itemsComplete, item)
		case err := <-errChan:
			common.Check(err)
		}
	}

	return itemsComplete
}
