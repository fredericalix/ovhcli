package cloud

import (
	"github.com/admdwrf/ovhcli/internal"

	"github.com/spf13/cobra"
)

func init() {
	cmdCloudSSHKeyCreate.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
	cmdCloudSSHKeyCreate.PersistentFlags().StringVarP(&pubkeyID, "pubkeyID", "", "", "Your sshkey ID to put")
	cmdCloudSSHKeyCreate.PersistentFlags().StringVarP(&name, "name", "", "", "Your sshkey name to put")

}

var cmdCloudSSHKeyCreate = &cobra.Command{
	Use:   "create",
	Short: "Create Cloud ssh key: ovhcli cloud sshkey create",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := internal.Client.CloudProjectSSHKeyCreate(projectID, pubkeyID, name)
		internal.Check(err)
		internal.FormatOutputDef(s)
	},
}
