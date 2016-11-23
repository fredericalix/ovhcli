package cloud

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/admdwrf/ovhcli/sdk"
	"github.com/spf13/cobra"
)

func init() {
	cmdCloudSSHKeyList.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")

}

var cmdCloudSSHKeyList = &cobra.Command{
	Use:   "list",
	Short: "List all ssk keys: ovhcli cloud sshkey list",
	Run: func(cmd *cobra.Command, args []string) {
		sshkeys, err := internal.Client.CloudProjectSSHKeyList(projectID)
		internal.Check(err)

		sshkeysComplete := []sdk.Sshkey{}
		for _, sshkey := range sshkeys {
			s, err := internal.Client.CloudProjectSSHKeyInfo(projectID, sshkey.ID)
			internal.Check(err)
			sshkeysComplete = append(sshkeysComplete, *s)
		}

		sshkeys = sshkeysComplete

		internal.Check(err)
		internal.FormatOutputDef(sshkeys)
	},
}
