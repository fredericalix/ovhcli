package cloud

import (
	"fmt"

	"github.com/admdwrf/ovhcli/internal"
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

		err := internal.Client.CloudDeleteInstance(projectID, instanceID)
		internal.Check(err)

		fmt.Printf("Instance %s deleted:\n", instanceID)

	},
}
