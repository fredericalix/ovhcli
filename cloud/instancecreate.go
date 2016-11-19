package cloud

import (
	"fmt"
	"os"

	"github.com/admdwrf/ovhcli/sdk"
	"github.com/spf13/cobra"
)

var ProjectID string
var ImageID string
var name string
var pubkeyID string
var flavorID string
var region string
var monthlyBilling bool

func init() {
	cmdInstanceCreate.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
	cmdInstanceCreate.PersistentFlags().StringVarP(&instanceID, "instanceID", "", "", "Your Instance ID to delete")

}

var cmdInstanceCreate = &cobra.Command{
	Use:   "create",
	Short: "Create Cloud Public Instance: ovhcli cloud instance create",
	Run: func(cmd *cobra.Command, args []string) {

		c, err := sdk.NewClient()
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}

		c.CloudCreateInstance(ProjectID, name, pubkeyID, flavorID, ImageID, region, monthlyBilling)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
		fmt.Printf("Instance %s deleted\n", instanceID)

	},
}
