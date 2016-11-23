package cloud

import (
	"fmt"

	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

var sshkeyID string

func init() {
	cmdCloudSSHKeyDelete.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
	cmdCloudSSHKeyDelete.PersistentFlags().StringVarP(&sshkeyID, "sshkeyID", "", "", "Your sshkey ID to delete")

}

var cmdCloudSSHKeyDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete Cloud SSH key: ovhcli cloud sshkey delete",
	Run: func(cmd *cobra.Command, args []string) {

		err := internal.Client.CloudProjectSSHKeyDelete(projectID, sshkeyID)
		internal.Check(err)

		fmt.Printf("Public SSH key %s deleted:\n", sshkeyID)

	},
}
