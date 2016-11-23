package cloud

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

//var project string

func init() {
	cmdCloudNetworkPrivateShow.PersistentFlags().StringVarP(&project, "project", "", "", "Your ID Project")
}

// cmdCloudNetworkPrivateShow show Public network ID of a project
var cmdCloudNetworkPrivateShow = &cobra.Command{
	Use:   "show",
	Short: "Show the private network ID of your project: ovhcli cloud network private show",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewClient()
		common.Check(err)

		netpub, err := client.CloudInfoNetworkPrivate(project)

		common.Check(err)
		common.FormatOutputDef(netpub)
	},
}
