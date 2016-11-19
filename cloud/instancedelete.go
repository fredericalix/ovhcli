package cloud

import (
	"fmt"
	"os"

	"github.com/admdwrf/ovhcli/sdk"
	"github.com/spf13/cobra"
)

var instanceID string

func init() {
	cmdInstanceDelete.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
	cmdInstanceDelete.PersistentFlags().StringVarP(&instanceID, "instanceID", "", "", "Your Instance ID to delete")

}

var cmdInstanceDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete Cloud Public Instance: ovhcli cloud instance delete",
	Run: func(cmd *cobra.Command, args []string) {

		c, err := sdk.NewClient()
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}

		c.CloudDeleteInstance(projectID, instanceID)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
		fmt.Printf("Instance %s deleted\n", instanceID)

	},
}
