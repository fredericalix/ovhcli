package cloud

import (
	"fmt"

	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

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

		client, err := ovh.NewClient()
		common.Check(err)

		err = client.CloudDeleteInstance(projectID, instanceID)
		common.Check(err)

		fmt.Printf("Instance %s deleted:\n", instanceID)

	},
}
